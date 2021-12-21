from flask import Flask, Response
from flask_restful import Resource, Api, reqparse
import paramiko
import json
import re

LIST_REGEX = r"\d: (.+) \((.*), last modified (.*)\)$"

HOSTS_FILE = "/app/hosts.conf"
HOSTS = []
with open(HOSTS_FILE) as hosts:
    for host in hosts.readlines():
        if not host in HOSTS:
            HOSTS.append(host)

SSH_CLIENT = paramiko.SSHClient()
SSH_CLIENT.set_missing_host_key_policy(paramiko.AutoAddPolicy())
SSH_KEY = paramiko.RSAKey.from_private_key_file("/root/.ssh/id_rsa")

app = Flask(__name__)
api = Api(app)


def ssh_command(command):

    data = []

    index=0

    for host in HOSTS:
        username, hostname = host.strip().split("@")
        SSH_CLIENT.connect(username=username, hostname=hostname, pkey=SSH_KEY)
        i, o, e = SSH_CLIENT.exec_command(command)

        stdout = ""
        stderr = ""

        for line in o.readlines():
            stdout += line + "\n"
        for line in e.readlines():
            stderr += line + "\n"

        data.append({"host":hostname,"stdout":stdout,"stderr":stderr})

        index += 1

    return data


class Gravity(Resource):
    pass


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

            for line in obj['stdout'].split('\n'):
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

            result.append({
                "host": obj['host'],
                "service_listening": service_listening,
                "udp_ipv4": udp_ipv4,
                "udp_ipv6": udp_ipv6,
                "tcp_ipv4": tcp_ipv4,
                "tcp_ipv6": tcp_ipv6,
                "blocking": blocking,
            })

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
                if len(result['stderr']) > 0:
                    print(result['stderr'])
                    return Response(response=result['stderr'], status=500)
            return Response(response="Success", status=200)

        elif action == "disable":
            disable_results = ssh_command("pihole disable")
            for result in disable_results:
                if len(result['stderr']) > 0:
                    print(result['stderr'])
                    return Response(response=result['stderr'], status=500)
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
            if len(result['stderr']) > 0:
                print(result['stderr'])
                return Response(response=result['stderr'], status=500)
        return Response(response="Success", status=200)


class DNSRecord(Resource):
    pass


class Whitelist(Resource):
    def get(self):
        results = ssh_command("pihole -w -l")
        for result in results:
            if len(result['stderr']) > 0:
                print(result['stderr'])
                return Response(response=result['stderr'], status=500)
        
        lists = []
        for result in results:
            data = []
            for line in result['stdout'].split('\n'):
                match = re.search(LIST_REGEX, line)
                if match:
                    domain = match[1]
                    state = match[2]
                    modified = match[3]
                    data.append(
                        {"domain": domain, "state": state, "modified": modified}
                    )
            lists.append({'host':result['host'],'data':data})
        response = json.dumps(lists)

        return Response(response=response, status=200)


class Blacklist(Resource):
    def get(self):
        results = ssh_command("pihole -b -l")
        for result in results:
            if len(result['stderr']) > 0:
                print(result['stderr'])
                return Response(response=result['stderr'], status=500)
        
        lists = []
        for result in results:
            data = []
            for line in result['stdout'].split('\n'):
                match = re.search(LIST_REGEX, line)
                if match:
                    domain = match[1]
                    state = match[2]
                    modified = match[3]
                    data.append(
                        {"domain": domain, "state": state, "modified": modified}
                    )
            lists.append({'host':result['host'],'data':data})
        response = json.dumps(lists)

        return Response(response=response, status=200)


api.add_resource(Gravity, "/gravity")
api.add_resource(Status, "/status")
api.add_resource(RestartDNS, "/restartdns")
api.add_resource(DNSRecord, "/dnsrecord")
api.add_resource(Whitelist, "/whitelist")
api.add_resource(Blacklist, "/blacklist")

if __name__ == "__main__":
    app.run(host="0.0.0.0")
