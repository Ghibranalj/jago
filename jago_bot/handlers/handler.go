package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

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
	out, err := run(code, inp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}

	messOut := fmt.Sprintf(" ```\n%s\n```", out)
	s.ChannelMessageSend(m.ChannelID, messOut)

	// g
}

func run(code, inp string) (string, error) {

	b, _ := json.Marshal(map[string]string{
		"code": code,
		"inp":  inp,
	})

	reqb := bytes.NewBuffer(b)

	resp, err := http.Post("http://localhost:1234/code", "application/json", reqb)

	if err != nil {
		return "", err
	}
	respb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respb), nil
}
