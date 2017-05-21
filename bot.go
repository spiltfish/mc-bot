package main

import (
	"github.com/bwmarrin/discordgo"
	"errors"
	"fmt"
	"strings"
	"strconv"
	"github.com/spiltfish/mc-worker-sdk"
	"os"
)

var(
	Token    string
	BotId    string
        WorkerUrl string
)


func init() {
	Token = os.Getenv("TOKEN")
	BotId = os.Getenv("BOT_ID")
	WorkerUrl = os.Getenv("WORKER_URL")
}



func main() {
	fmt.Println("Bot starting up!")
	mc_worker_sdk.SetServer(WorkerUrl)
	Token = strings.Replace(Token, "\n","",-1)

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Unable to create Discord Session: ", err)
	}

	//Handlers
	dg.AddHandler(messageCreate)

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	<-make(chan struct{})
	return

}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate){
	if message.Author.ID == BotId {
		return
	}

	if strings.HasPrefix(message.Content,"!mc start") {
		startMessage(session, message)
	}

	if strings.HasPrefix(message.Content , "!mc stop"){
		stopMessage(session, message)
	}

	if strings.HasPrefix(message.Content , "!mc ip"){
		ipMessage(session, message)
	}

	if strings.HasPrefix(message.Content , "!mc status"){
		statusMessage(session, message)
	}

	if strings.HasPrefix(message.Content , "!mc donate"){
		donateMessage(session, message)
	}

	if strings.HasPrefix(message.Content , "!mc help"){
		helpMessage(session, message)
	}

	if strings.HasPrefix(message.Content, "!mc new minecraft"){
		createNewServer(session, message)
	}

}


func createNewServer(session *discordgo.Session, message *discordgo.MessageCreate){
	required_parameters := 5
	words, err := checkParameters(message, required_parameters)
	if err != nil{
		session.ChannelMessageSend(message.ChannelID, "Not enough paramerters. Requires " + strconv.Itoa(required_parameters) + " parameters.")
		return
	}
	serverName := words[4]
	serverVersion := words[5]
	session.ChannelMessageSend(message.ChannelID, "Creating server" + " \"" + serverName + "\" " + serverVersion)
	mc_worker_sdk.CreateMinecraftServer(serverName)
	session.ChannelMessageSend(message.ChannelID, "Created server.")
}

func startMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	required_parameters := 3
	words, err := checkParameters(message, required_parameters)
	if err != nil{
		session.ChannelMessageSend(message.ChannelID, "Not enough paramerters. Requires " + strconv.Itoa(required_parameters) + " parameters.")
		return
	}
	serverName := words[3]
	mc_worker_sdk.PowerOnServer(serverName)
	session.ChannelMessageSend(message.ChannelID, "Server startup completed.")
}

func stopMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	required_parameters := 3
	words, err := checkParameters(message, required_parameters)
	if err != nil{
		session.ChannelMessageSend(message.ChannelID, "Not enough paramerters. Requires " + strconv.Itoa(required_parameters) + " parameters.")
	}
	serverName := words[3]
	fmt.Println("Shutting down server: " + serverName )
	mc_worker_sdk.PowerOffServer(serverName)
	session.ChannelMessageSend(message.ChannelID, "Server Shutting down.")
}

func ipMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	words := strings.Fields(message.Content)
	required_parameters := 3
	words, err := checkParameters(message, required_parameters)
	if err != nil{
		session.ChannelMessageSend(message.ChannelID, "Not enough paramerters. Requires " + strconv.Itoa(required_parameters) + " parameters.")
	}
	serverName := words[3]
	result := mc_worker_sdk.GetMinecraftServerIp(serverName)
	session.ChannelMessageSend(message.ChannelID, string(result))
}

func statusMessage(session *discordgo.Session, message *discordgo.MessageCreate){
	required_parameters := 3
	words, err := checkParameters(message, required_parameters)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Not enough paramerters. Requires " + strconv.Itoa(required_parameters) + " parameters.")
	}
	serverName := words[3]
	result := mc_worker_sdk.GetMinecraftServerStatus(serverName)
	_, err = session.ChannelMessageSend(message.ChannelID, string(result))
	if err != nil {
		fmt.Println(err.Error())
	}
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

func checkParameters(message *discordgo.MessageCreate, req_parameters int)(words []string, err error) {
	words = strings.Fields(message.Content)
	if len(words) < req_parameters {
		err = errors.New("Too few parameters.")
	}
	return words, err
}
