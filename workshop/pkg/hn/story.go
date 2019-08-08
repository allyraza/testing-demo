package hn

type Story struct {
	Id       int
	Author   string `json:"by"`
	Title    string
	Comments []Comment
}

type Comment struct {
	Id            int
	Author        string `json:"by"`
	Text          string
	ChildComments []Comment
}

type StoryBuilder struct {
	client Client
}

func NewStoryBuilder(client Client) *StoryBuilder {
	return &StoryBuilder{client: client}
}

func (b *StoryBuilder) Build(itemID int) (Story, error) {
	item, err := b.client.GetItem(itemID)
	if err != nil {
		return Story{}, err
	}

	var comments []Comment
	for _, k := range item.Kids {
		i, e := b.client.GetItem(k)
		comment := Comment{
			Id:     i.Id,
			Text:   i.Text,
			Author: i.Author,
		}
		if e != nil {
			comment.Id = k
			comment.Text = "[[Comment not found]]"
		}

		comments = append(comments, comment)
	}

	return Story{
		Id:       item.Id,
		Author:   item.Author,
		Title:    item.Title,
		Comments: comments,
	}, nil
}
