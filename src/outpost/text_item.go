package outpost

import (
	pb "outpost/outpostrpc"
)

type TextItem struct {
	// TODO: Do this struct need to define json fields?
	Text      string
	Id        string
	Timestamp string
	Category  string
}

type TextItemPersister interface {
	Insert(TextItem) error
	Retrieve() ([]*pb.TextItem, error)
	RetrieveOnId(string) ([]*pb.TextItem, error)
	// update(string) error
	// delete(string) error
}
