from oauth2client.client import GoogleCredentials
from googleapiclient import discovery

credentials = GoogleCredentials.get_application_default()
compute = discovery.build('compute', 'v1', credentials=credentials)

MC_PROJECT = "silent-space-421"
MC_ZONE = "us-central1-a"
MC_INSTANCE = "ftb-infinity-server-2"

HELP = "help"
START = "start"
STOP = "stop"
STATUS = "status"
IP = "ip"


def execute_command(command):
    if command.startswith(HELP):
        response = 'help' \
                   'start' \
                   'stop' \
                   'status' \
                   'ip'
    elif command.startswith(START):
        response = compute.instances().start(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['status']
    elif command.startswith(STOP):
        response = compute.instances().stop(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['status']
    elif command.startswith(STATUS):
        response = compute.instances().get(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['status']
    elif command.startswith(IP):
        response = compute.instances().get(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['networkInterfaces'][0]['accessConfigs'][0]['natIP']
    else:
        response = "Fuck you trever"
    return response


if __name__ == '__main__':
    print(compute.instances().list(project=MC_PROJECT, zone="us-central1-a").execute())