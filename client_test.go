package quartz_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/allyraza/quartz"
	"github.com/stretchr/testify/assert"
)

func TestHttpClient(t *testing.T) {
	t.Run("MaxTime", func(t *testing.T) {
		t.Run("500", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))
			defer ts.Close()

			item, err := quartz.NewHTTPClient(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Invalid body response", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte{})
			}))
			defer ts.Close()

			item, err := quartz.NewHTTPClient(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("No response", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			defer ts.Close()

			item, err := quartz.NewHTTPClient(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Invalid URL", func(t *testing.T) {
			item, err := quartz.NewHTTPClient("").MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("String response", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("foo"))
			}))
			defer ts.Close()

			item, err := quartz.NewHTTPClient(ts.URL).MaxItem()

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Valid response", func(t *testing.T) {
			value := 101
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(strconv.Itoa(value)))
			}))
			defer ts.Close()

			item, err := quartz.NewHTTPClient(ts.URL).MaxItem()

			assert.NoError(t, err)
			assert.Equal(t, item, value)
		})
	})

	t.Run("GetItem", func(t *testing.T) {
		t.Run("500", func(t *testing.T) {
			itemID := 101
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}))
			defer ts.Close()

			item, err := quartz.NewHTTPClient(ts.URL).GetItem(itemID)

			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("Success", func(t *testing.T) {
			itemID := 101

			blob, err := ioutil.ReadFile("testdata/item.json")
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(blob)
			}))
			defer ts.Close()

			item, err := quartz.NewHTTPClient(ts.URL).GetItem(itemID)
			expectedItem := quartz.Item{
				Id:     1,
				Author: "johndoe",
				Score:  100,
				Url:    "http://google.com",
				Title:  "Google",
			}

			assert.NoError(t, err)
			assert.Equal(t, item, expectedItem)
		})
	})
}

func TestHttpClient_Integration(t *testing.T) {
	t.Run("MaxItem integration test", func(t *testing.T) {
		if testing.Short() {
			t.SkipNow()
		}

		item, err := quartz.DefaultHTTPClient().MaxItem()

		assert.NoError(t, err)
		assert.NotEmpty(t, item)
	})

	t.Run("GetItem integration test", func(t *testing.T) {
		if testing.Short() {
			t.SkipNow()
		}

		itemID := 8863
		item, err := quartz.DefaultHTTPClient().GetItem(itemID)

		assert.NoError(t, err)
		assert.Equal(t, item.Author, "johndoe")
	})
}
