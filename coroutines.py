from oauth2client.client import GoogleCredentials
from googleapiclient import discovery

from constants import *
import asyncio

credentials = GoogleCredentials.get_application_default()
compute = discovery.build('compute', 'v1', credentials=credentials)


@asyncio.coroutine
async def wait_for_startup(client, channel):
    status = compute.instances().get(project=MC_PROJECT, zone=MC_ZONE, instance=MC_INSTANCE).execute()['status']
    if status == "RUNNING":
        await client.send_message(channel, "Server Started.")
    else:
        await asyncio.sleep(15)
        await wait_for_startup(client, channel)

