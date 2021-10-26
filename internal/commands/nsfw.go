package commands

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	config "github.com/brayanhenao/tombot-discord-bot/internal/config"
	utils "github.com/brayanhenao/tombot-discord-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
)

type RedditHelper struct {
	RedditResponse []utils.ResponseData
}

var (
	Helper RedditHelper
	err    error
)

func Nsfw(session *discordgo.Session, channelId string) {

	if config.CallNum == -1 {
		Helper, err = RefillImages()
		if err != nil {
			log.Fatalln(err)
		}

		config.CallNum++
	}

	if config.CallNum == len(Helper.RedditResponse) {
		config.CallNum = -1
	} else {
		messageImage := &discordgo.MessageEmbedImage{
			URL: Helper.RedditResponse[config.CallNum].Data["url"].(string),
		}

		_, err := session.ChannelMessageSendEmbed(channelId, &discordgo.MessageEmbed{
			URL:       messageImage.URL,
			Title:     Helper.RedditResponse[config.CallNum].Data["title"].(string),
			Color:     0x1e0f3,
			Image:     messageImage,
			Timestamp: time.Now().Format(time.RFC3339),
		})

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func RefillImages() (RedditHelper, error) {

	var (
		err            error
		response       []utils.ResponseData
		redditResponse utils.RedditResponse
	)

	subredditsUrls := getSubredditsUrls()

	for _, url := range subredditsUrls {
		redditResponse, err = utils.GetJson(url)
		if err != nil {
			return RedditHelper{}, err
		}

		subSlice := redditResponse.Data.Children[2:len(redditResponse.Data.Children)]
		response = append(response, subSlice...)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(response), func(i, j int) { response[i], response[j] = response[j], response[i] })

	return RedditHelper{
		RedditResponse: response,
	}, nil
}

func getSubredditsUrls() []string {
	var subredditsUrls []string

	subreddits := strings.Split(os.Getenv("NSFW_SUBREDDITS"), ",")
	for _, subreddit := range subreddits {
		url := fmt.Sprintf("https://www.reddit.com/r/%s/.json?limit=100", subreddit)
		subredditsUrls = append(subredditsUrls, url)
	}

	return subredditsUrls
}
