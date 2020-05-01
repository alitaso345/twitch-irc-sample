package main

import (
	"crypto/tls"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	irc "github.com/thoj/go-ircevent"
)

const channel = "#tkch_taisyou"
const serverssl = "irc.chat.twitch.tv:6697"

type Config struct {
	Nick     string
	Password string
}

func main() {
	var config Config
	envconfig.Process("TWITCH", &config)

	nick := config.Nick
	con := irc.IRC(nick, nick)
	con.Password = config.Password
	con.UseTLS = true
	con.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	con.AddCallback("001", func(e *irc.Event) { con.Join(channel) })
	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		fmt.Printf("[twitch]%s: %s\n", e.User, e.Arguments[1])
	})
	err := con.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	con.Loop()
}
