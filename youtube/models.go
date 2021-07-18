package youtube

type RestResponseYouTube struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id ItemInfo `json:"id"`
}

type ItemInfo struct {
	VideoId string `json:"videoId"`
}