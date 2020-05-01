package main

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	irc "github.com/thoj/go-ircevent"
)

const serverssl = "irc.chat.twitch.tv:6697"

type Config struct {
	Nick     string
	Password string
}

// example：https://www.twitch.tv/mogra
// go run main.go #mogra
func main() {
	var config Config
	envconfig.Process("TWITCH", &config)

	nick := config.Nick
	con := irc.IRC(nick, nick)

	con.Password = config.Password
	con.UseTLS = true
	con.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	channel := os.Args[1]
	con.AddCallback("001", func(e *irc.Event) { con.Join(channel) })
	con.AddCallback("PRIVMSG", printMessage)
	err := con.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	con.Loop()
}

func printMessage(e *irc.Event) {
	fmt.Printf("[twitch]%s: %s\n", e.User, e.Arguments[1])
}
