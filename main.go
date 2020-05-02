package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := discord.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID := u.ID

	fmt.Println(BotID)
	discord.AddHandler(messageHandler)

	err = discord.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//	fmt.Println(m.Content)
	res := strings.Split(m.Content, " ")
	if len(res) == 2 {
		if res[0] == "stock" {
			url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + res[1] + "&apikey="
			fmt.Println(url)
			resp, err := http.Get(url)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "wrong symbol entered")
				return
			}
			defer resp.Body.Close()
			contents, err := ioutil.ReadAll(resp.Body)
			in := []byte(string(contents))
			var raw map[string]interface{}
			if err := json.Unmarshal(in, &raw); err != nil {
				panic(err)
			}
			out, _ := json.Marshal(raw)
			_, _ = s.ChannelMessageSend(m.ChannelID, string(out))

		}
	}
}
