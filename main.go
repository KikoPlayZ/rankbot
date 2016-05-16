package main

import (
	"flag"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nhooyr/color/log"
)

var (
	email   = flag.String("email", "", "account email")
	pass    = flag.String("pass", "", "account password")
	guild   = flag.String("guild", "", "guild (server) to join")
	channel = flag.String("chan", "", "channel to join")
	message = flag.String("msg", "_", "message to be sent")
)

func main() {
	flag.Parse()
	if *email == "" || *pass == "" {
		log.Fatalln("please provide an email and password")
	}
	s, err := discordgo.New(*email, *pass)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("logged in")
	g := findGuild(s)
	if g == nil {
		log.Fatalln("could not find guild")
	}
	id := findChannel(s, g)
	if id == "" {
		log.Fatalln("could not find channel")
	}
	sendLoop(s, id)
}

func findGuild(s *discordgo.Session) *discordgo.Guild {
	gs, err := s.UserGuilds()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got guilds")
	for _, g := range gs {
		if g.Name == *guild {
			log.Println("found guild")
			return g
		}
	}
	return nil
}

func findChannel(s *discordgo.Session, g *discordgo.Guild) string {
	chs, err := s.GuildChannels(g.ID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got channels")
	for _, ch := range chs {
		if ch.Name == *channel {
			log.Println("found channel")
			return ch.ID
		}
	}
	return ""
}

func sendLoop(s *discordgo.Session, id string) {
	for t := time.Tick(time.Minute); ; <-t {
		if _, err := s.ChannelMessageSend(id, *message); err != nil {
			log.Println(err)
		} else {
			log.Println("sent message")
		}
	}
}
