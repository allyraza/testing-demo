package hn_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"workshop-starter/pkg/hn"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient(t *testing.T) {
	t.Run("MaxItem", func(t *testing.T) {
		t.Run("500", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Invalid response body", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte{})
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("No response", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Invalid URL", func(t *testing.T) {
			item, err := hn.NewHTTPClientFor("").MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Returning a string", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				_, _ = w.Write([]byte("whoops"))
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Success", func(t *testing.T) {
			retVal := 732183618
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				_, _ = w.Write([]byte(strconv.Itoa(retVal)))
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).MaxItem()

			assert.NoError(t, err)
			assert.Equal(t, item, retVal)
		})
	})

	t.Run("GetItem", func(t *testing.T) {
		t.Run("500", func(t *testing.T) {
			itemId := 123
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).GetItem(itemId)

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Success", func(t *testing.T) {
			t.Skip()
			file, err := ioutil.ReadFile("testdata/item.json")

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				_, _ = w.Write(file)
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).GetItem(123)
			expectedItem := hn.Item{
				Id:     8863,
				Author: "dhouston",
				Score:  104,
				Url:    "http://www.getdropbox.com/u/2/screencast.html",
				Title:  "My YC app: Dropbox - Throw away your USB drive",
			}

			assert.NoError(t, err)
			assert.Equal(t, expectedItem, item)
		})
	})
}

func TestHTTPClient_Integration(t *testing.T) {
	t.Run("MaxItem integration test", func(t *testing.T) {
		if testing.Short() {
			t.SkipNow()
		}

		item, err := hn.NewHTTPClient().MaxItem()

		assert.NoError(t, err)
		assert.NotEmpty(t, item)
	})

	t.Run("GetItem integration test", func(t *testing.T) {
		if testing.Short() {
			t.SkipNow()
		}

		item, err := hn.NewHTTPClient().GetItem(8863)

		assert.NoError(t, err)
		assert.Equal(t, item.Author, "dhouston")
	})
}
