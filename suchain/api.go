package suchain

import (
	"github.com/assetsadapterstore/suchain-adapter/sdk/client"
	"net/url"
)

const (
	baseURLPath = "/api"
)

type Api struct {
	Client *client.Client
}

func NewApi(baseUrl string) *Api {
	client2 := client.NewClient(nil)
	URL, _ := url.Parse(baseUrl + baseURLPath + "/")
	client2.BaseURL = URL
	Api :=&Api{
		Client: client2,
	}
	return Api
}
