package utils

import (
	"encoding/json"
	"fmt"
	"github.com/brayanhenao/tombot-discord-bot/internal/config"
	"log"
	"net/http"
	"strings"
	"time"
)

type YoutubeList struct {
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Id   struct {
			Kind    string `json:"kind"`
			VideoId string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string    `json:"title"`
			Thumbnails  struct {
				Default struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

func searchOnYoutube(query string)(YoutubeList,error){

	wordsArray := strings.Fields(query)
	urlSentence := wordsArray[0]

	for i := range wordsArray {
		if i == 0 || strings.Contains(wordsArray[i], "-l") {continue}
		urlSentence = fmt.Sprintf("%s&%s",urlSentence,wordsArray[i])
	}

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&key=%s",urlSentence,config.GoogleApi)
	client := &http.Client{Timeout: 0 * time.Second}
	req, err := http.NewRequest("GET",url,nil )
	if err != nil{
		log.Fatalln(err)
		return YoutubeList{},err
	}

	res, err := client.Do(req)
	if err != nil{
		log.Fatalln(err)
	}

	defer res.Body.Close()

	searchResults := YoutubeList{}
	err = json.NewDecoder(res.Body).Decode(&searchResults)
	if err != nil{
		return YoutubeList{},err
	}

	return searchResults, nil
}


func GetYoutubeURL(query string)(string, string, error) {

	list, error := searchOnYoutube(query)

	if strings.Contains(query,"-l"){

	}

	itemsLength := fmt.Sprintf("%s",list.Items[0].Snippet.Title)
	return itemsLength,"", error
}
