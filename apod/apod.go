package apod

import (
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func IsValidDate(date string) bool {
	// check if date format is YYYY-MM-DD
	data := strings.Split(date, "-")
	if len(data) != 3 {
		return false
	}
	return true
}

type Apod struct {
	Date           string
	Explanation    string
	Hdurl          string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
}

func (*Apod) GetApod(date string) Apod {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&date=" + date)
	if err != nil {
		log.Fatal(err)
		return Apod{}
	}

	return Apod{
		Date:           gjson.Get(string(resp.Body()), "date").String(),
		Explanation:    gjson.Get(string(resp.Body()), "explanation").String(),
		Hdurl:          gjson.Get(string(resp.Body()), "hdurl").String(),
		MediaType:      gjson.Get(string(resp.Body()), "media_type").String(),
		ServiceVersion: gjson.Get(string(resp.Body()), "service_version").String(),
		Title:          gjson.Get(string(resp.Body()), "title").String(),
		Url:            gjson.Get(string(resp.Body()), "url").String(),
	}
}
