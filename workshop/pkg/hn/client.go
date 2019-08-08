package hn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client interface {
	MaxItem() (int, error)
	GetItem(itemID int) (Item, error)
}

type HackerNewsClient struct {
	BaseUrl string
}

type Item struct {
	Id     int
	Author string `json:"by"`
	Score  int
	Url    string
	Title  string
	Text   string
	Kids   []int
}

func (s *HackerNewsClient) MaxItem() (int, error) {
	response, err := http.Get(s.BaseUrl + "/maxitem.json")

	if err != nil {
		return 0, err
	}

	parsedResponse, err := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		return 0, fmt.Errorf("Non 200 Status Code")
	}

	return strconv.Atoi(string(parsedResponse))
}

func (s *HackerNewsClient) GetItem(itemId int) (Item, error) {
	targetUrl := s.BaseUrl + "/item/" + strconv.Itoa(itemId) + ".json"
	response, err := http.Get(targetUrl)

	if err != nil {
		return Item{}, err
	}
  defer response.Body.Close()

	if response.StatusCode != 200 {
		return Item{}, fmt.Errorf("Non 200 Status Code")
	}

	var item Item
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&item)

	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func NewHTTPClientFor(url string) *HackerNewsClient {
	return &HackerNewsClient{url}
}

func NewHTTPClient() *HackerNewsClient {
	return &HackerNewsClient{"https://hacker-news.firebaseio.com/v0"}
}
