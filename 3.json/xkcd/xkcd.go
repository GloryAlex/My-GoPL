package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type comic struct {
	Mouth      string
	Num        int
	Link       string
	Year       string
	News       string
	Safe_title string
	Alt        string
	Img        string
	Title      string
	Day        string
}

var comicDictionary = make(map[int]comic)

func GetComicURL(id int) string {
	if _, ok := comicDictionary[id]; !ok {
		resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", id))
		if err != nil {
			return ""
		} else if resp.StatusCode != http.StatusOK {
			return ""
		}
		var result comic
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return ""
		}
		_ = resp.Body.Close()
		comicDictionary[id] = result
	}
	return comicDictionary[id].Img
}
