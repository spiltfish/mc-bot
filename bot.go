package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"./discord_util"
	"strings"
	"./mc-worker-sdk"
)

var(
	Email    string
	Password string
	Token    string
	BotId    string

)


func init() {
	di := discord.ReadConfig("./bot/discord_config.yml")

	Email = di.Email
	Password = di.Password
	Token = di.Token
	BotId = di.BotId

}



func main() {
	fmt.Println("Bot starting up!")

	dg, err := discordgo.New(Email, Password, Token)
	if err != nil {
		fmt.Println("Unable to create Discord Session: ", err)
	}

	//Handlers
	dg.AddHandler(messageCreate)

	dg.Open()

	<-make(chan struct{})
	return

}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate){

	if message.Author.ID == BotId {
		return
	}

	if message.Content == "!mc" {
		_, _ = session.ChannelMessageSend(message.ChannelID, "NOT IMPLEMENTED")
	}

	if message.Content == "!mc start"{
		startMessage(session, message)
	}

	if message.Content == "!mc stop"{
		stopMessage(session, message)
	}

	if message.Content == "!mc ip"{
		ipMessage(session, message)
	}

	if message.Content == "!mc status"{
		statusMessage(session, message)
	}

	if message.Content == "!mc donate"{
		donateMessage(session, message)
	}

	if message.Content == "!mc help"{
		helpMessage(session, message)
	}

	if strings.HasPrefix(message.Content, "!mc new minecraft"){
		createNewServer(session, message)
	}

}


func createNewServer(session *discordgo.Session, message *discordgo.MessageCreate){
	words := strings.Fields(message.Content)
	serverName := words[3]
	serverVersion := words[4]
	session.ChannelMessageSend(message.ChannelID, "Creating server" + "\"" + serverName + "\" " + serverVersion)
	mc_worker_sdk.CreateMinecraftServer(serverName)
	session.ChannelMessageSend(message.ChannelID, "Created server.")
}

func startMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	words := strings.Fields(message.Content)
	serverName := words[3]
	mc_worker_sdk.PowerOnServer(serverName)
	session.ChannelMessageSend(message.ChannelID, "Server startup completed.")
}

func stopMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	words := strings.Fields(message.Content)
	serverName := words[3]
	mc_worker_sdk.PowerOffServer(serverName)
	session.ChannelMessageSend(message.ChannelID, "Server Shutting down.")
}

func ipMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	words := strings.Fields(message.Content)
	serverName := words[3]
	result := mc_worker_sdk.GetMinecraftServerIp(serverName)
	session.ChannelMessageSend(message.ChannelID, string(result))
}

func statusMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	words := strings.Fields(message.Content)
	serverName := words[3]
	result := mc_worker_sdk.GetMinecraftServerStatus(serverName)
	_, _ = session.ChannelMessageSend(message.ChannelID, string(result))
}

func donateMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	_, _ = session.ChannelMessageSend(message.ChannelID, "Please consider donating: https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=KZ8YFPXGHKY3W&lc=US&item_name=Mary%27s%20Servers%20and%20Bots&currency_code=USD&bn=PP%2dDonationsBF%3abtn_donate_SM%2egif%3aNonHosted")
}

func helpMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	helpString := "ip" +
		"status" +
		"start" +
		"stop" +
		"donate" +
		"new minecraft <name> <version tag>"
	_, _ = session.ChannelMessageSend(message.ChannelID, helpString)
}
