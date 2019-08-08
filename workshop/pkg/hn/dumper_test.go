package hn_test

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"workshop-starter/pkg/hn"
	"workshop-starter/pkg/hn/mock"
)

type errDump struct {
}

func (errDump) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("Cannot dump items")
}

func TestDumper(t *testing.T) {
	t.Run("writer error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockClient(ctrl)
		client.EXPECT().MaxItem().Return(5, nil)
		client.EXPECT().GetItem(5).Return(getItem(5), nil)

		err := hn.NewDump(client, 3).Dump(errDump{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockClient(ctrl)
		client.EXPECT().MaxItem().Return(5, nil)
		client.EXPECT().GetItem(5).Return(getItem(5), nil)
		client.EXPECT().GetItem(4).Return(getItem(4), nil)
		client.EXPECT().GetItem(3).Return(getItem(3), nil)

		var b bytes.Buffer
		err := hn.NewDump(client, 3).Dump(&b)
		assert.NoError(t, err)
		assert.Equal(t, "Title 5,5\nTitle 4,4\nTitle 3,3\n", b.String())
	})

	t.Run("top stories error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockClient(ctrl)
		client.EXPECT().MaxItem().Return(0, fmt.Errorf("Dupa"))

		var b bytes.Buffer
		err := hn.NewDump(client, 3).Dump(&b)
		assert.Error(t, err)
	})

	t.Run("one of get items fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockClient(ctrl)
		client.EXPECT().MaxItem().Return(5, nil)
		client.EXPECT().GetItem(5).Return(getItem(5), nil)
		client.EXPECT().GetItem(4).Return(getItem(4), fmt.Errorf("Failed"))
		client.EXPECT().GetItem(3).Return(getItem(3), nil)

		var b bytes.Buffer
		err := hn.NewDump(client, 3).Dump(&b)
		assert.NoError(t, err)
		assert.Equal(t, "Title 5,5\nTitle 3,3\n", b.String())
	})
}

func getItem(itemID int) hn.Item {
	items := map[int]hn.Item{
		5: hn.Item{Title: "Title 5", Score: 5},
		4: hn.Item{Title: "Title 4", Score: 4},
		3: hn.Item{Title: "Title 3", Score: 3},
		2: hn.Item{Title: "Title 2", Score: 2},
		1: hn.Item{Title: "Title 1", Score: 1},
	}

	return items[itemID]
}
