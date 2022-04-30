package model

import (
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

func GetBooks(category, info string, offset, count uint64) (rows Books, err error) {
	query, args := consBookQuery("id,isbn,book.name,author,price,intro",
		category, info)

	query += " limit ?, ?"
	args = append(args, offset, count)

	rs, err := db.Query(query, args...)
	if err != nil {
		return nil, err
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

func CountBooks(category, info string) (count uint64) {
	query, args := consBookQuery("count(*)", category, info)
	db.QueryRow(query, args...).Scan(&count)
	return
}

func consBookQuery(cols, category, info string) (query string, args []any) {
	query = "select " + cols

	if category == "" {
		query += ` from book
		where not deleted`
	} else {
		query += ` from book, category
		where id = category.book_id
			and category.name = ?
			and not deleted`
		args = append(args, category)
	}

	if info != "" {
		query += " and (isbn = ? or book.name like ? or author like ?)"
		args = append(args, info, "%"+info+"%", "%"+info+"%")
	}

	return
}
