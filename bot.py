#!/usr/bin/python3
import discord
from commands import execute_command

client = discord.Client()

with open("discord.key") as discord_key_file:
   discord_key = discord_key_file.readline().rstrip('\n')

@client.event
async def on_message(message):
    if message.author != 'mc-hammer#8736':
        if message.content.startswith("!mc "):
            response = execute_command(message.content[4:])
            await client.send_message(message.channel, response)

client.run(discord_key) 
