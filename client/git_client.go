package client

import (
	"github.com/go-resty/resty"
)

func GetReposGitByUser(url string) (*resty.Response, error) {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url)

	return resp, err
}
