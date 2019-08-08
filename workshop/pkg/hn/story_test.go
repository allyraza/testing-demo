package hn_test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"workshop-starter/pkg/hn"
	"workshop-starter/pkg/hn/mock"
)

func TestStoryBuilder(t *testing.T) {
	t.Run("error getting root item", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rootItemID := 4324234

		client := mock.NewMockClient(ctrl)
		client.EXPECT().GetItem(rootItemID).Return(hn.Item{}, fmt.Errorf("Item not found"))

		storyBuilder := hn.NewStoryBuilder(client)
		story, err := storyBuilder.Build(rootItemID)

		assert.Empty(t, story)
		assert.Error(t, err)
	})

	t.Run("success with no children", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storyItemID := 4324234
		client := mock.NewMockClient(ctrl)

		rootItem := getItemFromTestData(t, "item_no_children")
		client.EXPECT().GetItem(storyItemID).Return(rootItem, nil)

		storyBuilder := hn.NewStoryBuilder(client)
		story, err := storyBuilder.Build(storyItemID)
		assert.NoError(t, err)

		expectedStory := hn.Story{
			Id:     8863,
			Author: "dhouston",
			Title:  "My YC app: Dropbox - Throw away your USB drive",
		}
		assert.Equal(t, expectedStory, story)
	})

	t.Run("success, children not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storyItemID := 4324234
		client := mock.NewMockClient(ctrl)

		rootItem := getItemFromTestData(t, "item")
		client.EXPECT().GetItem(storyItemID).Return(rootItem, nil)
		client.EXPECT().GetItem(9224).Return(hn.Item{}, fmt.Errorf("Comment not found"))
		client.EXPECT().GetItem(8917).Return(hn.Item{}, fmt.Errorf("Comment not found"))

		storyBuilder := hn.NewStoryBuilder(client)
		story, err := storyBuilder.Build(storyItemID)
		assert.NoError(t, err)

		expectedStory := hn.Story{
			Id:     8863,
			Author: "dhouston",
			Title:  "My YC app: Dropbox - Throw away your USB drive",
			Comments: []hn.Comment{
				{
					Id:   9224,
					Text: "[[Comment not found]]",
				},
				{
					Id:   8917,
					Text: "[[Comment not found]]",
				},
			},
		}
		assert.Equal(t, expectedStory, story)
	})

	t.Run("success, single child found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storyItemID := 4324234
		client := mock.NewMockClient(ctrl)

		rootItem := getItemFromTestData(t, "item_single_child")
		client.EXPECT().GetItem(storyItemID).Return(rootItem, nil)
		client.EXPECT().GetItem(8917).Return(getItemFromTestData(t, "child_2"), nil)

		storyBuilder := hn.NewStoryBuilder(client)
		story, err := storyBuilder.Build(storyItemID)
		assert.NoError(t, err)

		expectedStory := hn.Story{
			Id:     8863,
			Author: "dhouston",
			Title:  "My YC app: Dropbox - Throw away your USB drive",
			Comments: []hn.Comment{
				{
					Id:     2921984,
					Text:   "Title #2",
					Author: "Wilduck",
				},
			},
		}
		assert.Equal(t, expectedStory, story)
	})

	t.Run("success, multi children recursive", func(t *testing.T) {
		t.Skip()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storyItemID := 4324234
		client := mock.NewMockClient(ctrl)

		rootItem := getItemFromTestData(t, "item")
		client.EXPECT().GetItem(storyItemID).Return(rootItem, nil)
		client.EXPECT().GetItem(9224).Return(getItemFromTestData(t, "child_1"), nil)
		client.EXPECT().GetItem(8917).Return(getItemFromTestData(t, "child_2"), nil)
		client.EXPECT().GetItem(2922097).Return(getItemFromTestData(t, "child_3"), nil)

		storyBuilder := hn.NewStoryBuilder(client)
		story, err := storyBuilder.Build(storyItemID)
		assert.NoError(t, err)

		expectedStory := hn.Story{
			Id:     8863,
			Author: "dhouston",
			Title:  "My YC app: Dropbox - Throw away your USB drive",
			Comments: []hn.Comment{
				{
					Id:     2921983,
					Text:   "Title #1",
					Author: "norvig",
					ChildComments: []hn.Comment{
						{
							Id: 2922097,
						},
					},
				},
				{
					Id:     2921984,
					Text:   "Title #2",
					Author: "Wilduck",
				},
			},
		}
		assert.Equal(t, expectedStory, story)
	})
}

func getItemFromTestData(t *testing.T, filename string) hn.Item {
	t.Helper()
	file, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.json", filename))
	if err != nil {
		t.Errorf(err.Error())
	}

	var item hn.Item
	err = json.Unmarshal(file, &item)
	if err != nil {
		t.Errorf(err.Error())
	}

	return item
}
