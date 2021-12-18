from flask import Flask, Response
from flask_restful import Resource, Api, reqparse
import paramiko
from getpass import getpass
import json

HOSTNAME = input("Pihole Hostname: ")
USERNAME = input("Pihole Username: ")
PASSWORD = getpass("Pihole Password: ")

app = Flask(__name__)
api = Api(app)


def ssh_command(command):
    ssh_client = paramiko.SSHClient()
    ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh_client.connect(hostname=HOSTNAME, username=USERNAME, password=PASSWORD)

    return ssh_client.exec_command(command)


class Gravity(Resource):
    pass


class Status(Resource):
    def get(self):

        stdin, stdout, stderr = ssh_command("pihole status")

        service_listening = False
        udp_ipv4 = False
        udp_ipv6 = False
        tcp_ipv4 = False
        tcp_ipv6 = False
        blocking = False

        for line in stdout.readlines():
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

        data = {
            "service_listening": service_listening,
            "udp_ipv4": udp_ipv4,
            "udp_ipv6": udp_ipv6,
            "tcp_ipv4": tcp_ipv4,
            "tcp_ipv6": tcp_ipv6,
            "blocking": blocking,
        }

        response = json.dumps(data)

        print(response)

        return Response(response=response, mimetype="application/json", status=200)

    def post(self):
        parser = reqparse.RequestParser()

        parser.add_argument("action", required=True)

        args = parser.parse_args()

        action = args["action"]

        if action == "enable":
            en_in, en_out, en_err = ssh_command("pihole enable")
            if len(en_err.readlines()) > 0:
                print(en_err.readlines())
                return Response(response=en_err.readlines(), status=500)
            else:
                return Response(response="Success", status=200)
        elif action == "disable":
            dis_in, dis_out, dis_err = ssh_command("pihole disable")
            if len(dis_err.readlines()) > 0:
                print(dis_err.readlines())
                return Response(response=dis_err.readlines(), status=500)
            else:
                return Response(response="Success", status=200)
        else:
            return Response(
                response="Invalid Request. Your options are 'action=enable' or 'action=disable'",
                status=400,
            )


class RestartDNS(Resource):
    def post(self):
        stdin, stdout, stderr = ssh_command("pihole restartdns")
        if len(stderr.readlines()) > 0:
            print(stderr.readlines())
            return Response(response=stderr.readlines(), status=500)
        else:
            return Response(response="Success", status=200)


class DNSRecord(Resource):
    pass


api.add_resource(Gravity, "/gravity")
api.add_resource(Status, "/status")
api.add_resource(RestartDNS, "/restartdns")
api.add_resource(DNSRecord, "/dnsrecord")

if __name__ == "__main__":
    app.run()
