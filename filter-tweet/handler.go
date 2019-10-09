package function

// Copyright Alex Ellis 2019
// Source: https://github.com/openfaas/social-functions/blob/master/filter-tweets/handler.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {

	currentTweet := tweet{}

	unmarshalErr := json.Unmarshal(req, &currentTweet)

	if unmarshalErr != nil {
		return fmt.Sprintf("Unable to unmarshal event: %s\n", unmarshalErr.Error())
	}

	if strings.Contains(currentTweet.Text, "RT") || strings.Contains(currentTweet.Username, os.Getenv("username")) {
		return fmt.Sprintf("Filtered out RT\n")
	}

	if val, ok := os.LookupEnv("username"); ok && len(val) > 0 && strings.Contains(currentTweet.Username, val) {
		return fmt.Sprintf("Filtered out own tweet from %s\n", currentTweet.Username)
	}

	slackURL := readSecret("incoming-webhook-url")

	slackMsg := slackMessage{
		Text:     "@" + currentTweet.Username + ": " + currentTweet.Text + " (via " + currentTweet.Link + ")",
		Username: "@" + currentTweet.Username,
	}

	bodyBytes, _ := json.Marshal(slackMsg)
	httpReq, _ := http.NewRequest(http.MethodPost, slackURL, bytes.NewReader(bodyBytes))
	res, resErr := http.DefaultClient.Do(httpReq)
	if resErr != nil {
		fmt.Fprintf(os.Stderr, "resErr: %s\n", resErr)
		os.Exit(1)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	return fmt.Sprintf("Tweet sent, with statusCode: %d\n", res.StatusCode)
}

// tweet in following format from IFTTT:
// { "text": "<<<{{Text}}>>>", "username": "<<<{{UserName}}>>>", "link": "<<<{{LinkToTweet}}>>>" }
type tweet struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	Link     string `json:"link"`
}

type slackMessage struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

func readSecret(name string) string {
	res, err := ioutil.ReadFile(path.Join("/var/openfaas/secrets/", name))
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}
	return string(res)
}
