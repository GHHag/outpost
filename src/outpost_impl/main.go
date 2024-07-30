package main

import (
	"outpost"
)

func main() {
	persister, err := NewSQLitePersister("./db/outpost.db")
	if err != nil {
		panic(err)
	}
	defer persister.db.Close()

	outpost.Run(persister)
}
