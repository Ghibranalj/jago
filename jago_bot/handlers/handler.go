package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	mc := m.Content
	reg := regexp.MustCompile(`!jago.*`)

	if !strings.HasPrefix(mc, "!jago") {
		return
	}
	// fmt.Println(mc)
	inst := reg.FindString(mc)
	fmt.Println("Head:", inst)

	code := reg.ReplaceAllString(mc, "")

	code = strings.ReplaceAll(code, "```", "")

	fmt.Println("Code:", code)

	inp := strings.ReplaceAll(inst, "!jago", "")

	//TODO spin up container

	out, err := run(code, inp)
	if err != nil {

		if err, ok := err.(net.Error); ok && err.Timeout() {
			messOut := fmt.Sprintf(" ```\nError: %s\n```", "Your code took too long to execute (max 10s)")
			s.ChannelMessageSend(m.ChannelID, messOut)

			//TODO kill container

			return
		}

		fmt.Fprintf(os.Stderr, "%s", err.Error())
		// TODO kill container
		return
	}
	// TODO kill container

	messOut := fmt.Sprintf(" ```\n%s\n```", out)
	s.ChannelMessageSend(m.ChannelID, messOut)

}

func run(code, inp string) (string, error) {

	b, _ := json.Marshal(map[string]string{
		"code": code,
		"inp":  inp,
	})

	reqb := bytes.NewBuffer(b)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Post("http://localhost:1234/code", "application/json", reqb)

	if err != nil {
		return "", err
	}

	respb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respb), nil
}
