package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	_ "github.com/joho/godotenv/autoload"
	"github.com/otiai10/twistream"
	"os"
	"regexp"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
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

		if isContainSushi(status.Text) {
			rt, err := api.Favorite(status.Id)
			fmt.Println(rt)
			fmt.Println(err)
		}
		fmt.Println(status.Text)
	}
}

func isContainSushi(text string) (b bool) {
	if m, _ := regexp.MatchString("寿司|スシ|鮨|寿し|🍣|[sS][uU][sS][hH][iI]", text); !m {
		return false
	}
	return true
}

func isContainMot(text string) (b bool) {
	if m, _ := regexp.MatchString("MOT|mot", text); !m {
		return false
	}
	return true
}
