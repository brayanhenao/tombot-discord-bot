package commands

import (
	"fmt"
	"github.com/brayanhenao/tombot-discord-bot/internal/framework"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/brayanhenao/tombot-discord-bot/internal/config"
	"github.com/brayanhenao/tombot-discord-bot/internal/utils"
)

type RedditHelper struct {
	RedditResponse []utils.ResponseData
}

var (
	Helper RedditHelper
	err    error
)

func Nsfw(ctx framework.Context) {
	valid := false

	ignoredUrlProviders := strings.Split(os.Getenv("IGNORED_URL_PROVIDERS"), ",")

	for !valid {
		if config.CallNum == -1 || config.CallNum == len(Helper.RedditResponse) {
			Helper, err = RefillImages()
			if err != nil {
				log.Fatalln(err)
			}

			config.CallNum = 0
		}

		imageUrl := Helper.RedditResponse[config.CallNum].Data["url"].(string)
		if imageProviderIsValid(imageUrl, ignoredUrlProviders) {
			if imageContainsGift(imageUrl) {
				createGiftEmbed(ctx, imageUrl)
			} else {
				createImageEmbed(ctx, imageUrl)
			}

			valid = true
		}
		config.CallNum = config.CallNum + 1
	}
}

func imageContainsGift(url string) bool {
	return strings.Contains(url, "gifv") || strings.Contains(url, "gif")
}

func createImageEmbed(ctx framework.Context, url string) {
	messageImage := &discordgo.MessageEmbedImage{
		URL: url,
	}

	_, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, &discordgo.MessageEmbed{
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

func createGiftEmbed(ctx framework.Context, url string) {
	messageImage := &discordgo.MessageEmbedImage{
		URL: strings.Replace(url, "gifv", "gif", 1),
	}

	_, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, &discordgo.MessageEmbed{
		URL:         url,
		Title:       Helper.RedditResponse[config.CallNum].Data["title"].(string),
		Color:       0x1e0f3,
		Image:       messageImage,
		Timestamp:   time.Now().Format(time.RFC3339),
		Description: "This is a gif, click the title to see it in full size",
	})

	if err != nil {
		log.Fatalln(err)
	}
}

func RefillImages() (RedditHelper, error) {

	var (
		err            error
		data           []utils.ResponseData
		redditResponse utils.RedditResponse
	)

	subredditsUrls := getSubredditsUrls()

	for _, url := range subredditsUrls {
		redditResponse, err = utils.GetRedditResponse(url)
		if err != nil {
			return RedditHelper{}, err
		}

		subSlice := redditResponse.Data.Children[2:len(redditResponse.Data.Children)]
		data = append(data, subSlice...)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })

	return RedditHelper{
		RedditResponse: data,
	}, nil
}

func getSubredditsUrls() []string {
	var subredditsUrls []string

	subreddits := strings.Split(os.Getenv("NSFW_SUBREDDITS"), ",")
	for _, subreddit := range subreddits {
		url := fmt.Sprintf("https://reddit-helper-tombot.herokuapp.com/v1/subreddit/%s", subreddit)
		subredditsUrls = append(subredditsUrls, url)
	}

	return subredditsUrls
}

func imageProviderIsValid(url string, ignoredUrlProviders []string) bool {
	valid := true
	for _, ignoredUrlProvider := range ignoredUrlProviders {
		if strings.Contains(url, ignoredUrlProvider) {
			valid = false
			break
		}
	}

	return valid
}
