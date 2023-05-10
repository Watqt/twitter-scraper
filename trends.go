package twitterscraper

import (
	"fmt"
)

type Trend struct {
	Name            string `json:"name"`
	Url             string `json:"url"`
	PromotedContent string `json:"promoted_content"`
	Query           string `json:"query"`
	TweetVolume     int    `json:"tweet_volume"`
}
type Trends struct {
	Trends []Trend `json:"trends"`
}

// GetTrends return list of trends.
func (s *Scraper) GetTrends() ([]Trend, error) {
	req, err := s.newRequest("GET", "https://api.twitter.com/1.1/trends/place.json")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("id", "23424819") // Paris / France
	req.URL.RawQuery = q.Encode()

	var jsn Trends
	s.setBearerToken(bearerToken2)
	err = s.RequestAPI(req, &jsn)
	s.setBearerToken(bearerToken)
	if err != nil {
		return nil, err
	}

	if len(jsn.Trends) < 2 {
		return nil, fmt.Errorf("no trend entries found")
	}

	return jsn.Trends, nil
}

// Deprecated: GetTrends wrapper for default Scraper
func GetTrends() ([]Trend, error) {
	return defaultScraper.GetTrends()
}
