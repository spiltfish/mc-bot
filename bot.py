import discord
from commands import execute_command

client = discord.Client()


@client.event
async def on_message(message):
    if message.author != 'mc-hammer#8736':
        if message.content.startswith("!mc "):
            response = execute_command(message.content[4:])
            await client.send_message(message.channel, response)

client.run('MjA5Mzg2ODkxOTYxODI3MzM4.Cn_3wQ._AevNMOXIhcxO0i1_auoz16udr0')