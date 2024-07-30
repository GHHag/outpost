package outpost

import (
	pb "outpost/outpostrpc"
)

type TextItem struct {
	Text      string
	RefTag    string
	Timestamp string
	Category  string
}

type TextItemPersister interface {
	Insert(TextItem) error
	Retrieve() ([]*pb.TextItem, error)
	RetrieveOnRefTag(string) ([]*pb.TextItem, error)
	// update(string) error
	// delete(string) error
}

type TextItemCollector interface {
	Collect() ([]*pb.TextItem, error)
	Persist() error
}
