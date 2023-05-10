package twitterscraper

import (
	"fmt"
)

type Trend struct {
	Name            string      `json:"name"`
	PromotedContent interface{} `json:"promoted_content"`
	Query           string      `json:"query"`
	TweetVolume     interface{} `json:"tweet_volume"`
	URL             string      `json:"url"`
}

type Location struct {
	Name  string  `json:"name"`
	Woeid float64 `json:"woeid"`
}

type TrendData struct {
	AsOf      string     `json:"as_of"`
	CreatedAt string     `json:"created_at"`
	Locations []Location `json:"locations"`
	Trends    []Trend    `json:"trends"`
}

// GetTrends return list of trends.
func (s *Scraper) GetTrends(WOEID string) ([]Trend, error) {
	req, err := s.newRequest("GET", "https://api.twitter.com/1.1/trends/place.json")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("id", WOEID) // 23424819 Paris / France
	req.URL.RawQuery = q.Encode()

	var jsn []TrendData
	s.setBearerToken(bearerToken2)
	err = s.RequestAPI(req, &jsn)
	s.setBearerToken(bearerToken)
	if err != nil {
		return nil, err
	}

	if len(jsn) == 0 || len(jsn[0].Trends) == 0 {
		return nil, fmt.Errorf("no trend entries found")
	}

	return jsn[0].Trends, nil
}

// Deprecated: GetTrends wrapper for default Scraper
func GetTrends(WOEID string) ([]Trend, error) {
	return defaultScraper.GetTrends(WOEID)
}
