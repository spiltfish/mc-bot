#!/usr/bin/python3
import asyncio
import discord
from commands import execute_command

client = discord.Client()

with open("discord.key") as discord_key_file:
    discord_key = discord_key_file.readline().rstrip('\n')


@client.event
async def on_message(message):
    if message.author != 'mc-hammer#8736':
        if message.content.startswith("!mc "):
            response = execute_command(client, message.content[4:], message.channel)
            await client.send_message(message.channel, response)


try:
    tasks = [
        client.start(discord_key)
    ]
    client.loop.run_until_complete(asyncio.wait(tasks))
except KeyboardInterrupt:
    client.loop.run_until_complete(client.logout())
    pending = asyncio.Task.all_tasks()
    gathered = asyncio.gather(*pending)
    try:
        gathered.cancel()
        client.loop.run_until_complete(gathered)
        # we want to retrieve any exceptions to make sure that
        # they don't nag us about it being un-retrieved.
        gathered.exception()
    except:
        pass
finally:
    client.loop.close()
