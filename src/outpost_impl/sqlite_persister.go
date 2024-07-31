package main

import (
	"database/sql"
	"hash/fnv"
	"outpost"
	pb "outpost/outpostrpc"

	_ "github.com/mattn/go-sqlite3"
)

type SQLitePersister struct {
	db *sql.DB
}

func NewSQLitePersister(path string) (SQLitePersister, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return SQLitePersister{}, err
	}

	persister := SQLitePersister{
		db: db,
	}

	if err = persister.createReferenceTagsTable(); err != nil {
		return SQLitePersister{}, err
	}

	if err = persister.createTextItemsTable(); err != nil {
		return SQLitePersister{}, err
	}

	if err = persister.createIndices(); err != nil {
		return SQLitePersister{}, err
	}

	return persister, nil
}

func (persister SQLitePersister) createReferenceTagsTable() error {
	referenceTableSQL := `
		CREATE TABLE IF NOT EXISTS reference_tags(
			reference_tag VARCHAR(64) UNIQUE NOT NULL PRIMARY KEY
		);
	`

	_, err := persister.db.Exec(referenceTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func (persister SQLitePersister) createTextItemsTable() error {
	textItemsTableSQL := `
		CREATE TABLE IF NOT EXISTS text_items(
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			reference_tag_fk VARCHAR(64) NOT NULL,
			text TEXT NOT NULL,
			timestamp VARCHAR(32) NOT NULL,
			category VARCHAR(64) NOT NULL,
			unihash INTEGER UNIQUE NOT NULL,
			FOREIGN KEY(reference_tag_fk) REFERENCES references_tags(reference_tag)
		);
	`

	_, err := persister.db.Exec(textItemsTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func (persister SQLitePersister) createIndices() error {
	stmt := `
		CREATE INDEX IF NOT EXISTS reference_tag_idx ON text_items(reference_tag_fk);
		CREATE INDEX IF NOT EXISTS timestamp_idx ON text_items(timestamp);
		CREATE INDEX IF NOT EXISTS category_idx ON text_items(timestamp);
	`
	_, err := persister.db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (persister SQLitePersister) insertRefTag(refTag string) error {
	insertQuery := `
		INSERT OR IGNORE INTO reference_tags(reference_tag)
		VALUES (?);
	`

	tx, err := persister.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(refTag)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (persister SQLitePersister) Insert(textItem outpost.TextItem) error {
	// TODO: Implement batch insert query instead?

	// TODO: Insert the given reference_tag if it doesn't already exists in reference_tags table.
	// Is there a better way to do it?
	if err := persister.insertRefTag(textItem.RefTag); err != nil {
		return err
	}

	insertQuery := `
		INSERT INTO text_items(reference_tag_fk, text, timestamp, category, unihash)
		VALUES (?, ?, ?, ?, ?);
	`

	tx, err := persister.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	unihash := fnv.New64a()
	unihash.Write([]byte(
		textItem.RefTag +
			// textItem.Text +
			textItem.Timestamp +
			textItem.Category))
	hashSum := unihash.Sum64()

	_, err = stmt.Exec(
		textItem.RefTag,
		textItem.Text,
		textItem.Timestamp,
		textItem.Category,
		hashSum,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (persister SQLitePersister) Retrieve() ([]*pb.TextItem, error) {
	selectQuery := `
		SELECT reference_tag_fk, text, timestamp, category
		FROM text_items;
	`

	rows, err := persister.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	textItems := make([]*pb.TextItem, 0)
	for rows.Next() {
		var textItem pb.TextItem
		if err = rows.Scan(
			&textItem.RefTag,
			&textItem.Text,
			&textItem.Timestamp,
			&textItem.Category,
		); err != nil {
			return nil, err
		} else {
			textItems = append(textItems, &textItem)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return textItems, nil
}

// TODO: Abstract the similar simple select queries
func (persister SQLitePersister) RetrieveOnRefTag(refTag string) ([]*pb.TextItem, error) {
	selectQuery := `
		SELECT reference_tag_fk, text, timestamp, category
		FROM text_items
		WHERE reference_tag_fk = ?;
	`

	stmt, err := persister.db.Prepare(selectQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(refTag)
	if err != nil {
		return nil, err
	}

	textItems := make([]*pb.TextItem, 0)
	for rows.Next() {
		var textItem pb.TextItem
		if err = rows.Scan(
			&textItem.RefTag,
			&textItem.Text,
			&textItem.Timestamp,
			&textItem.Category,
		); err != nil {
			return nil, err
		} else {
			textItems = append(textItems, &textItem)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return textItems, nil

}
