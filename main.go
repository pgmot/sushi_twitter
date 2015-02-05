package main

import (
	"encoding/json"
	"fmt"
	"github.com/otiai10/twistream"
	"io/ioutil"
)

type Config struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)

	CONSUMERKEY := config.ConsumerKey
	CONSUMERSECRET := config.ConsumerSecret
	ACCESSTOKEN := config.AccessToken
	ACCESSTOKENSECRET := config.AccessTokenSecret
	timeline, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		CONSUMERKEY,
		CONSUMERSECRET,
		ACCESSTOKEN,
		ACCESSTOKENSECRET,
	)

	for {
		status := <-timeline.Listen()
		fmt.Println(status)
	}
}
