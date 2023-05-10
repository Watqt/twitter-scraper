package twitterscraper_test

import (
	"testing"

	twitterscraper "github.com/Watqt/twitter-scraper"
)

func TestGetTrends(t *testing.T) {
	scraper := twitterscraper.New()
	trends, err := scraper.GetTrends("23424819")
	if err != nil {
		t.Error(err)
	}

	if len(trends) != 20 {
		t.Errorf("Expected 20 trends, got %d: %#v", len(trends), trends)
	}

	for _, trend := range trends {
		if trend.Name == "" {
			t.Error("Expected trend is empty")
		}
	}
}
