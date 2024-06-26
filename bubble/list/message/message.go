package message

import (
	"time"

	"github.com/bensadeh/circumflex/item"
)

type EditorFinishedMsg struct {
	Err error
}

type EnteringCommentSection struct {
	Id           int
	CommentCount int
}

type EnteringReaderMode struct {
	Url    string
	Title  string
	Domain string
}

type ShowStatusMessage struct {
	Message  string
	Duration time.Duration
}

type StatusMessageTimeout struct{}

type FetchingFinished struct {
	Message string
}

type FetchAndChangeToCategory struct {
	Index    int
	Category int
	Cursor   int
}

type Refresh struct {
	CurrentCategory int
	CurrentIndex    int
}

type CategoryFetchingFinished struct {
	Index   int
	Cursor  int
	Message string
}

type AddToFavorites struct {
	Item *item.Item
}
