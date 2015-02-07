package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/otiai10/twistream"
	"io/ioutil"
	"regexp"
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

	consumerKey := config.ConsumerKey
	consumerSecret := config.ConsumerSecret
	accessToken := config.AccessToken
	accessTokenSecret := config.AccessTokenSecret
	timeline, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		consumerKey,
		consumerSecret,
		accessToken,
		accessTokenSecret,
	)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	for {
		status := <-timeline.Listen()

		if IsContainSushi(status.Text) {
			api.Favorite(status.Id)
		}
		fmt.Println(status.Text)
	}
}

func IsContainSushi(text string) (b bool) {
	if m, _ := regexp.MatchString("寿司|スシ|鮨|寿し|🍣|[sS][uU][sS][hH][iI]", text); !m {
		return false
	}
	return true
}
