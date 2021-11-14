package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ghibranalj/jago/jago_bot/handlers"
	"github.com/bwmarrin/discordgo"
)

var (
	token = os.Getenv("DISCORD")
)

func init() {
	if token == "" {
		fmt.Fprintln(os.Stderr, "$TOKEN not set")
		os.Exit(1)
	}
}

func main() {

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error while creating Discord session : %s \n", err.Error())
		os.Exit(1)
	}
	dg.AddHandler(handlers.MessageHandler)

	dg.Identify.Intents = discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening connection : %s \n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}
