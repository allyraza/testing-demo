package hn

import (
	"github.com/davecgh/go-spew/spew"
	"io"
	"strconv"
)

type Dump struct {
	client Client
	limit  int
}

func NewDump(client Client, limit int) *Dump {
	return &Dump{client: client, limit: limit}
}

func (d *Dump) Dump(w io.Writer) error {
	maxItem, err := d.client.MaxItem()
	if err != nil {
		spew.Dump(err)
		return err
	}

	for i := 0; i < d.limit; i++ {
		itemID := maxItem - i
		item, err := d.client.GetItem(itemID)
		if err != nil {
			continue
		}
		_, err = io.WriteString(w, item.Title+","+strconv.Itoa(item.Score)+"\n")
		if err != nil {
			return err
		}
	}

	return nil
}
