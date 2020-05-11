package libknary

import (
	"bytes"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func sendMsg(msg string) {
	// closes https://github.com/sudosammy/knary/issues/20
	re := regexp.MustCompile(`\r?\n`)
	msg = re.ReplaceAllString(msg, "\\n")

	if os.Getenv("SLACK_WEBHOOK") != "" {
		jsonMsg := []byte(`{"username":"knary","icon_emoji":":bird:","text":"` + msg + `"}`)
		_, err := http.Post(os.Getenv("SLACK_WEBHOOK"), "application/json", bytes.NewBuffer(jsonMsg))

		if err != nil {
			Printy(err.Error(), 2)
		}
	}

	if os.Getenv("PUSHOVER_TOKEN") != "" && os.Getenv("PUSHOVER_USER") != "" {
		jsonMsg := []byte(`{"token":"` + os.Getenv("PUSHOVER_TOKEN") + `","user":"` + os.Getenv("PUSHOVER_USER") + `","message":"` + msg + `"}`)
		_, err := http.Post("https://api.pushover.net/1/messages.json/", "application/json", bytes.NewBuffer(jsonMsg))

		if err != nil {
			Printy(err.Error(), 2)
		}
	}

	if os.Getenv("DISCORD_WEBHOOK") != "" {
		jsonMsg := []byte(`{"username":"knary","text":"` + msg + `"}`)
		_, err := http.Post(os.Getenv("DISCORD_WEBHOOK")+"/slack", "application/json", bytes.NewBuffer(jsonMsg))

		if err != nil {
			Printy(err.Error(), 2)
		}
	}

	if os.Getenv("TEAMS_WEBHOOK") != "" {
		// swap ``` with <pre> for MS teams :face-with-rolling-eyes:
		msg = strings.Replace(msg, "```", "</pre>", 2)
		msg = strings.Replace(msg, "</pre>", "<pre>", 1)

		jsonMsg := []byte(`{"text":"` + msg + `"}`)
		_, err := http.Post(os.Getenv("TEAMS_WEBHOOK"), "application/json", bytes.NewBuffer(jsonMsg))

		if err != nil {
			Printy(err.Error(), 2)
		}
	}

	// should be simple enough to add support for other webhooks here
}
