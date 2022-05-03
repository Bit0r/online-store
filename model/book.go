package model

import (
	"log"
)

type Book struct {
	ID           uint64 // 数据库自增id，添加图书时不用填写
	ISBN         string
	Cover        string
	Name, Author string
	Price        float64
	Intro        string
	Deleted      bool
}

type BooksFilter struct {
	Category   string
	Info       string
	MustExists bool
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

func GetBookCategories(id uint64) (categories []string) {
	query := `select name
	from category
	where book_id = ?`
	rs, _ := db.Query(query, id)
	defer rs.Close()

	var name string
	for rs.Next() {
		err := rs.Scan(&name)
		if err != nil {
			log.Fatal(err)
		} else {
			categories = append(categories, name)
		}
	}

	return
}

func GetBooks(filter BooksFilter, offset, count uint64) (rows Books, err error) {
	query, args := consBookQuery("id,isbn,book.name,author,price,intro,deleted", filter)

	query += " limit ?, ?"
	args = append(args, offset, count)

	rs, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	book := Book{}

	for rs.Next() {
		err := rs.Scan(&book.ID, &book.ISBN, &book.Name, &book.Author, &book.Price, &book.Intro, &book.Deleted)
		if err != nil {
			log.Fatal(err)
		} else {
			rows = append(rows, book)
		}
	}

	return
}

func CountBooks(filter BooksFilter) (count uint64) {
	query, args := consBookQuery("count(*)", filter)
	db.QueryRow(query, args...).Scan(&count)
	return
}

func consBookQuery(cols string, filter BooksFilter) (query string, args []any) {
	query = "select " + cols

	if category := filter.Category; category == "" {
		query += ` from book
		where true`
	} else {
		query += ` from book, category
		where id = category.book_id
			and category.name = ?`
		args = append(args, category)
	}

	if filter.MustExists {
		query += " and not deleted"
	}

	if info := filter.Info; info != "" {
		query += " and (isbn = ? or book.name like ? or author like ?)"
		args = append(args, info, "%"+info+"%", "%"+info+"%")
	}

	query += " order by id desc"

	return
}

func GetBook(id uint64) (book Book, err error) {
	query := `select id, isbn, name, author, coalesce(cover, ''), price, intro, deleted
	from book
	where id = ?`
	err = db.QueryRow(query, id).Scan(&book.ID, &book.ISBN, &book.Name, &book.Author, &book.Cover, &book.Price, &book.Intro, &book.Deleted)
	return
}

func AddBook(book Book) {

}
