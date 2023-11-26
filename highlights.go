package koboh

import (
	"database/sql"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Highlight struct {
	Title     string
	Highlight string
}

func OpenDB() (*sql.DB, error) {
	dbFile, err := os.OpenFile("/Volumes/KOBOeReader/.kobo/KoboReader.sqlite", os.O_RDONLY, os.ModePerm)
	if os.IsNotExist(err) {
		return nil, err
	}

	dbTmpFile, err := os.OpenFile("/tmp/kobo.sqlite", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if os.IsNotExist(err) {
		return nil, err
	}

	_, err = io.Copy(dbTmpFile, dbFile)
	if err != nil {
		return nil, err
	}

	return sql.Open("sqlite3", "/tmp/kobo.sqlite")
}

func ExtractHighlights(db *sql.DB) ([]Highlight, error) {
	sqlStmt := `
	select c.booktitle, b.text from bookmark as b, content as c where type == "highlight" and c.contentid == b.contentid order by b.datecreated, b.contentid;
	`
	highlights := []Highlight{}

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		var highlight string
		err = rows.Scan(&title, &highlight)
		if err != nil {
			return nil, err
		}
		highlights = append(highlights, Highlight{Title: title, Highlight: highlight})
	}
	err = rows.Err()
	return highlights, err
}
