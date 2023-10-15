package crawler

import (
	"github.com/go-resty/resty/v2"
)

const baseEndpoint = "https://www.instagram.com/api/v1/"

var client = resty.New()

type InstagramWebLayer struct {
}

func (s *InstagramWebLayer) GetProfileInfo(username string) (resty.Response, error) {
	endpoint := baseEndpoint + "users/web_profile_info/"
	response, err := client.R().SetQueryParams(map[string]string{
		"username": username,
	}).SetHeaders(map[string]string{
		"authority":                   "www.instagram.com",
		"accept":                      "*/*",
		"accept-language":             "pt-BR:pt;q=0.9",
		"dpr":                         "1",
		"referer":                     "https://www.instagram.com/igor_deo/",
		"sec-ch-prefers-color-scheme": "dark",
		"sec-ch-ua":                   "\"Chromium\";v=\"118\": \"Google Chrome\";v=\"118\": \"Not=A?Brand\";v=\"99\"",
		"sec-ch-ua-full-version-list": "\"Chromium\";v=\"118.0.5993.70\": \"Google Chrome\";v=\"118.0.5993.70\": \"Not=A?Brand\";v=\"99.0.0.0\"",
		"sec-ch-ua-mobile":            "?0",
		"sec-ch-ua-model":             "\"\"",
		"sec-ch-ua-platform":          "\"Windows\"",
		"sec-ch-ua-platform-version":  "\"15.0.0\"",
		"sec-fetch-dest":              "empty",
		"sec-fetch-mode":              "cors",
		"sec-fetch-site":              "same-origin",
		"user-agent":                  "Mozilla/5.0 (Windows NT 10.0; Win64; x64, AppleWebKit/537.36 (KHTML: like Gecko, Chrome/118.0.0.0 Safari/537.36",
		"viewport-width":              "1278",
		"x-ig-app-id":                 "936619743392459",
	}).Get(endpoint)

	if err != nil {
		panic(err)
	}

	if response.StatusCode() != 200 {
		panic("Status code is not 200")
	}

	return *response, nil
}
