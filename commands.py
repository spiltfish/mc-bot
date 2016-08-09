from oauth2client.client import GoogleCredentials
from googleapiclient import discovery

from constants import *
from coroutines import wait_for_startup
import asyncio

credentials = GoogleCredentials.get_application_default()
compute = discovery.build('compute', 'v1', credentials=credentials)


def get_donation_info():
    info_text = "Please consider donating: {link}".format(link="https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=KZ8YFPXGHKY3W&lc=US&item_name=Mary%27s%20Servers%20and%20Bots&currency_code=USD&bn=PP%2dDonationsBF%3abtn_donate_SM%2egif%3aNonHosted")
    return info_text


def start_server_command(client, channel):
    response = compute.instances().start(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['status']
    response += "\n"
    response += get_donation_info()
    asyncio.ensure_future(wait_for_startup(client, channel), loop=asyncio.get_event_loop())
    return response


def status_command():
    response = compute.instances().get(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['status']
    return response


def execute_command(client, command, channel):
    if command.startswith(HELP):
        response = 'help' \
                   'start' \
                   'stop' \
                   'status' \
                   'ip'
    elif command.startswith(START):
        return start_server_command(client, channel)
    elif command.startswith(STOP):
        response = compute.instances().stop(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['status']
    elif command.startswith(STATUS):
        return status_command()
    elif command.startswith(IP):
        response = compute.instances().get(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['networkInterfaces'][0]['accessConfigs'][0]['natIP']
    else:
        response = "Fuck you trever"
    return response
