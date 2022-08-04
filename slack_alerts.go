package slackAlerts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type messageBody struct {
	Text string `json:"text"`
}

func Message(msg string) (string, error) {

	m := messageBody{Text: msg}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return "fail", errors.New("somethings went wrong")
	}
	mB := string(b)
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return "Please provide SLACK_WEBHOOK_URL at .env file.", err
	}

	url := os.Getenv("SLACK_WEBHOOK_URL")
	payload := strings.NewReader(mB)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}
