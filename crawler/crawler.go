package crawler

import "fmt"

type instagramCrawler struct {
	parser InstagramParser
	web    InstagramWebLayer
}

func InstagramCrawlerFactory() *instagramCrawler {
	return &instagramCrawler{
		parser: InstagramParser{},
		web:    InstagramWebLayer{},
	}
}

func (s *instagramCrawler) GetProfileInfo(username string) (InstagramProfileInfo, error) {

	response, error := s.web.GetProfileInfo(username)

	parsed_response, error := s.parser.ParseProfileInfo(response)

	fmt.Print(parsed_response)

	return parsed_response, error
}
