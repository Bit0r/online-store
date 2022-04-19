package model

import (
	"database/sql"
	"log"
)

type Book struct {
	ID           uint64
	ISBN         string
	Name, Author string
	Price        float64
	Intro        string
}

type Books = []Book

func GetCategories() (names []string) {
	query := `select distinct name
	from category`
	rs, _ := db.Query(query)
	defer rs.Close()

	var name string
	for rs.Next() {
		err := rs.Scan(&name)
		if err != nil {
			log.Fatal(err)
		} else {
			names = append(names, name)
		}
	}

	return
}

func GetBooks(category string, offset, row_count uint64) (rows Books) {
	var query string
	var rs *sql.Rows

	// query := "select id, isbn, name, author, price, intro"
	info := ""
	switch {
	case category == "" && info == "":
		query = `select id, isbn, name, author, price, intro
		from book
		where not deleted
		limit ?, ?`
		rs, _ = db.Query(query, offset, row_count)
	case category != "" && info == "":
		query = `select id, isbn, book.name, author, price, intro
		from book, category
		where id = category.book_id
			and category.name = ?
			and not deleted
		limit ?, ?`
		rs, _ = db.Query(query, category, offset, row_count)
	case category == "" && info != "":
		query = `select id, isbn, name, author, price, intro
		from book
		where not deleted and (isbn = ? or name regex ? or author regex ?)
		limit ?, ?`
		rs, _ = db.Query(query, offset, row_count)
	}

	defer rs.Close()

	book := Book{}

	for rs.Next() {
		err := rs.Scan(&book.ID, &book.ISBN, &book.Name, &book.Author, &book.Price, &book.Intro)
		if err != nil {
			log.Fatal(err)
		} else {
			rows = append(rows, book)
		}
	}

	return
}

func CountBooks(category string) (count uint64) {
	var query string
	if category == "" {
		query = `select count(*)
		from book
		where not deleted`
		db.QueryRow(query).Scan(&count)
	} else {
		query = `select count(*)
		from book, category
		where id = category.book_id
			and category.name = ?
			and not deleted`
		db.QueryRow(query, category).Scan(&count)
	}
	return
}
