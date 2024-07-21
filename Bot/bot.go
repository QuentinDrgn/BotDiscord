package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)


var BotToken string

func CheckNilError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	CheckNilError(err)

	discord.AddHandler(MessageCreate)

	err = discord.Open()
	CheckNilError(err)

	fmt.Println("Bot is running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func MessageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, "!ping") {
		discord.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}