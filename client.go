package quartz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	defaultMaxItem = 0
	defaultURL     = "https://hacker-news.firebaseio.com/v0"
)

// Client interface to be implemented
type Client interface {
	MaxItem() (int, error)
	GetItem(itemID int) (Item, error)
}

// HTTPClient implements a http support
type HTTPClient struct {
	BaseURL string
}

// Item is type for article
type Item struct {
	Id     int
	Author string
	Score  int
	Url    string
	Title  string
}

// MaxItem returns max item
func (c HTTPClient) MaxItem() (int, error) {
	response, err := http.Get(c.BaseURL + "/maxitem.json")
	if err != nil {
		return defaultMaxItem, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return defaultMaxItem, err
	}

	return strconv.Atoi(string(body))
}

// GetItem by id
func (c HTTPClient) GetItem(itemID int) (Item, error) {
	itemURL := fmt.Sprintf("%s/item/%d.json", c.BaseURL, itemID)
	response, err := http.Get(itemURL)
	if err != nil {
		return Item{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Item{}, errors.New("Invalid status code")
	}

	var item Item
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&item)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

// NewHTTPClient creates a new client
func NewHTTPClient(url string) Client {
	return HTTPClient{url}
}

// DefaultHTTPClient creates a new client
func DefaultHTTPClient() Client {
	return HTTPClient{defaultURL}
}
