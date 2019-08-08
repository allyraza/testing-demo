package hn

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeps(t *testing.T) {

	//go get github.com/golang/mock/gomock
	//go install github.com/golang/mock/mockgen
	t.Run("gomock", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		assert.NotNil(t, ctrl)
	})

	t.Run("spew.Dump", func(t *testing.T) {
		assert.NotEmpty(t, spew.Sprint(t))
	})
}
