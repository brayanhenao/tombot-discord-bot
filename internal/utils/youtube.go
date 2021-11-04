package utils

import "fmt"

type stream map[string]string
type youtube struct {
	streamList []stream
	videoId string
	videoInfo string
}


func GetYoutubeURL(query string)(string, string, error) {
	yInstance := new(youtube)
	fmt.Println("%s %s",query,yInstance.videoId)
	return yInstance.videoInfo, yInstance.videoId, nil
}
