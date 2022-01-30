from flask import Flask, Response
from flask_restful import Resource, Api, reqparse
import paramiko
import json
import re

LIST_REGEX = r"\d: (.+) \((.*), last modified (.*)\)$"

CUSTOM_DOMAIN_FILE = "/etc/dnsmasq.d/05-pihole-custom-cname.conf"

HOSTS_FILE = "/app/hosts.conf"
HOSTS = []
with open(HOSTS_FILE) as hosts:
    for host in hosts.readlines():
        if not host in HOSTS:
            HOSTS.append(host)

SSH_CLIENT = paramiko.SSHClient()
SSH_CLIENT.set_missing_host_key_policy(paramiko.AutoAddPolicy())
SSH_KEY = paramiko.RSAKey.from_private_key_file("/root/.ssh/id_rsa")

ADLIST_ID = 0
ADLIST_ADDRESS = 1
ADLIST_COMMENT = 2

app = Flask(__name__)
api = Api(app)


def ssh_command(command):

    data = []

    index = 0

    for host in HOSTS:
        try:
            username, hostname = host.strip().split("@")
            print(f"Running {command} on {hostname}")
            SSH_CLIENT.connect(username=username, hostname=hostname, pkey=SSH_KEY)
            i, o, e = SSH_CLIENT.exec_command(command)

            stdout = ""
            stderr = ""

            for line in o.readlines():
                stdout += line #+ "\n"
            for line in e.readlines():
                stderr += line #+ "\n"

            data.append({"host": hostname, "stdout": stdout, "stderr": stderr})

            index += 1
        finally:  # Ensure the connection is closed, even if there is an error during the execution above
            SSH_CLIENT.close()

    return data


class Gravity(Resource):
    def post(self):  # Updates Gravity
        results = ssh_command("pihole -g")

        for result in results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)
        return Response(response="Success", status=200)

    def delete(self):
        pass

    def patch(self):
        pass

    def get(self):
        results = ssh_command('/usr/bin/sqlite3 /etc/pihole/gravity.db "SELECT id, address, comment FROM adlist;"')

        for result in results:
            if len(result['stderr']) > 0:
                print(result['stderr'])
                return Response(response=result['stderr'], status=500)
        
        data = {}

        for result in results:
            host_result = []
            for entry in result['stdout'].split('\n'):
                columns = entry.split("|")
                if len(columns) == 3:
                    host_result.append({'id':columns[ADLIST_ID], 'comment':columns[ADLIST_COMMENT], 'address':columns[ADLIST_ADDRESS]})
            data[result['host']] = host_result

        response = json.dumps(data)

        return Response(response=response,mimetype='application/json',status=200)


class Status(Resource):
    def get(self):
        result = []

        data = ssh_command("pihole status")

        for obj in data:
            service_listening = False
            udp_ipv4 = False
            udp_ipv6 = False
            tcp_ipv4 = False
            tcp_ipv6 = False
            blocking = False

            for line in obj["stdout"].split("\n"):
                # Check for True conditions
                if "✓" in line:
                    if "DNS service" in line:
                        service_listening = True
                    elif "UDP (IPv4)" in line:
                        udp_ipv4 = True
                    elif "UDP (IPv6)" in line:
                        udp_ipv6 = True
                    elif "TCP (IPv4)" in line:
                        tcp_ipv4 = True
                    elif "TCP (IPv6)" in line:
                        tcp_ipv6 = True
                    elif "Pi-hole blocking" in line:
                        blocking = True

            result.append(
                {
                    "host": obj["host"],
                    "service_listening": service_listening,
                    "udp_ipv4": udp_ipv4,
                    "udp_ipv6": udp_ipv6,
                    "tcp_ipv4": tcp_ipv4,
                    "tcp_ipv6": tcp_ipv6,
                    "blocking": blocking,
                }
            )

        response = json.dumps(result)

        return Response(response=response, mimetype="application/json", status=200)

    def post(self):
        parser = reqparse.RequestParser()

        parser.add_argument("action", required=True)

        args = parser.parse_args()

        action = args["action"]

        if action == "enable":
            enable_results = ssh_command("pihole enable")
            for result in enable_results:
                if len(result["stderr"]) > 0:
                    print(result["stderr"])
                    return Response(response=result["stderr"], status=500)
            return Response(response="Success", status=200)

        elif action == "disable":
            disable_results = ssh_command("pihole disable")
            for result in disable_results:
                if len(result["stderr"]) > 0:
                    print(result["stderr"])
                    return Response(response=result["stderr"], status=500)
            return Response(response="Success", status=200)
        else:
            return Response(
                response="Invalid Request. Your options are 'action=enable' or 'action=disable'",
                status=400,
            )


class RestartDNS(Resource):
    def post(self):
        results = ssh_command("pihole restartdns")
        for result in results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)
        return Response(response="Success", status=200)


class CNAMERecord(Resource):
    def get_records(self):
        results = ssh_command(f"cat {CUSTOM_DOMAIN_FILE}")
        for result in results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)

        domains = []
        for result in results:
            for line in result["stdout"].split("\n"):
                if not line in domains and line != "":
                    domains.append(line)

        resp = json.dumps({"domains": domains})

        return resp

    def get(self):
        return Response(response=self.get_records(), status=200)

    def put(self):
        parser = reqparse.RequestParser()

        parser.add_argument("cname", required=True)
        parser.add_argument("host", required=True)

        args = parser.parse_args()

        cname = args["cname"]
        host = args["host"]

        domains = json.loads(self.get_records()).get("domains")
        domains.append(f"cname={cname},{host}")

        domain_str = ""
        for domain in domains:
            domain_str += domain + "\n"

        update_results = ssh_command(
            f"echo '{domain_str.strip()}' > /tmp/domains.tmp; sudo cp /tmp/domains.tmp {CUSTOM_DOMAIN_FILE}"
        )
        for result in update_results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)

        restart_results = ssh_command("pihole restartdns")
        for result in restart_results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)

        return Response(response="Success", status=200)

    def delete(self):
        parser = reqparse.RequestParser()

        parser.add_argument("cname", required=True)
        parser.add_argument("host", required=True)

        args = parser.parse_args()

        cname = args["cname"]
        host = args["host"]

        domains = json.loads(self.get_records()).get("domains")

        domains.remove(f"cname={cname},{host}")

        domain_str = ""
        for domain in domains:
            domain_str += domain + "\n"

        update_results = ssh_command(
            f"echo '{domain_str.strip()}' > /tmp/domains.tmp; sudo cp /tmp/domains.tmp {CUSTOM_DOMAIN_FILE}"
        )
        for result in update_results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)

        restart_results = ssh_command("pihole restartdns")
        for result in restart_results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)

        return Response(response="Success", status=200)


class Whitelist(Resource):
    def get(self):
        results = ssh_command("pihole -w -l")
        for result in results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)

        lists = []
        for result in results:
            data = []
            for line in result["stdout"].split("\n"):
                match = re.search(LIST_REGEX, line)
                if match:
                    domain = match[1]
                    state = match[2]
                    modified = match[3]
                    data.append(
                        {"domain": domain, "state": state, "modified": modified}
                    )
            lists.append({"host": result["host"], "data": data})
        response = json.dumps(lists)

        return Response(response=response, status=200)


class Blacklist(Resource):
    def get(self):
        results = ssh_command("pihole -b -l")
        for result in results:
            if len(result["stderr"]) > 0:
                print(result["stderr"])
                return Response(response=result["stderr"], status=500)

        lists = []
        for result in results:
            data = []
            for line in result["stdout"].split("\n"):
                match = re.search(LIST_REGEX, line)
                if match:
                    domain = match[1]
                    state = match[2]
                    modified = match[3]
                    data.append(
                        {"domain": domain, "state": state, "modified": modified}
                    )
            lists.append({"host": result["host"], "data": data})
        response = json.dumps(lists)

        return Response(response=response, status=200)


api.add_resource(Gravity, "/gravity")
api.add_resource(Status, "/status")
api.add_resource(RestartDNS, "/restartdns")
api.add_resource(CNAMERecord, "/cnamerecord")
api.add_resource(Whitelist, "/whitelist")
api.add_resource(Blacklist, "/blacklist")

if __name__ == "__main__":
    ssh_command(f"sudo touch {CUSTOM_DOMAIN_FILE}")
    app.run(host="0.0.0.0")
