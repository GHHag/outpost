package main

type TextItem struct {
	// TODO: Do this struct need to define json fields?
	text      string
	id        string
	timestamp string
	category  string
}

type TextItemPersister interface {
	insert(TextItem)
	retrieve() []TextItem
	// update(string)
	// delete(string)
}
