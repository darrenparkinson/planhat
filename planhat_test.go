package planhat

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPlanhat_AddOptions(t *testing.T) {
	type TestOptions struct {
		Limit  *int    `url:"limit,omitempty"`
		Offset *int    `url:"offset,omitempty"`
		Sort   *string `url:"sort,omitempty"`
		Select *string `url:"select,omitempty"`
	}

	t.Run("options as expected", func(t *testing.T) {
		to := TestOptions{
			Limit:  Int(100),
			Offset: Int(200),
			Sort:   String("-name"),
			Select: String("one,two,three"),
		}
		want := fmt.Sprintf("%s?limit=100&offset=200&select=%s&sort=-name", "myurl.com/endpoint", url.QueryEscape("one,two,three"))
		got, err := addOptions("myurl.com/endpoint", to)
		if err != nil {
			t.Errorf("didn't expect error adding options: %v", err)
		}
		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})
	t.Run("missing options don't appear", func(t *testing.T) {
		to := TestOptions{
			Limit: Int(100),
			Sort:  String("name"),
		}
		want := "myurl.com/endpoint?limit=100&sort=name"
		got, err := addOptions("myurl.com/endpoint", to)
		if err != nil {
			t.Errorf("didn't expect error adding options: %v", err)
		}
		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})
}
