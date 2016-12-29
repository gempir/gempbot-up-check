package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"os"
	"strings"
	"time"
)

// Bot 123
type Bot struct {
	server        string
	port          string
	nick          string
	user          string
	channel       string
	pass          string
	pread, pwrite chan string
	conn          net.Conn
}

// NewBot xD
func NewBot() *Bot {
	return &Bot{
		server:  "irc.chat.twitch.tv",
		port:    "6667",
		nick:    "gempir",
		channel: "#gempir",
		pass:    oauth,
		conn:    nil,
		user:    "gempir",
	}
}

// Connect asd
func (bot *Bot) Connect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", bot.server+":"+bot.port)
	if err != nil {
		log.Fatal("unable to connect to IRC server ", err)
	}
	bot.conn = conn
	log.Printf("Connected to IRC server %s (%s)\n", bot.server, bot.conn.RemoteAddr())
	return bot.conn, nil
}

func main() {
	ircbot := NewBot()
	conn, _ := ircbot.Connect()

	fmt.Fprintf(conn, "PASS %s\r\n", ircbot.pass)
	fmt.Fprintf(conn, "NICK %s\r\n", ircbot.nick)
	fmt.Fprintf(conn, "JOIN %s\r\n", ircbot.channel)

	defer conn.Close()

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	for {
		line, err := tp.ReadLine()
		if err != nil {
			break // break loop on errors
		}
		if strings.Contains(line, ":tmi.twitch.tv 001") {
			fmt.Fprintf(conn, "PRIVMSG %s :!status\r\n", ircbot.channel)
			time.AfterFunc(3*time.Second, exitWithError)
		}
		if strings.Contains(line, ":gempbot!gempbot@gempbot.tmi.twitch.tv PRIVMSG #gempir :gempir, uptime:") {
			os.Exit(0)
		}
	}

}

func exitWithError() {
	os.Exit(1)
}
