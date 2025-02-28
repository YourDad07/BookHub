package controller

import (
	dbs "Bookhub/db"
	"Bookhub/models"
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	db := dbs.Connect()
	query := "SELECT bookpath,imgpath,subjectname,semester,universityname,branch,bookauthor FROM bookinfo"
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Error in getting bookinfo", err)
	}

	bookinfo := []models.Book{}

	defer db.Close()
	defer rows.Close()

	var book models.Book

	for rows.Next() {
		fmt.Println("book ")

		err = rows.Scan(&book.Bookfile, &book.Bannerimage, &book.Subjectname, &book.Semnumber, &book.Universityname, &book.Branch, &book.Authorname)
		if err != nil {
			fmt.Println("Error in scanning bookinfo", err)
		}

		bookinfo = append(bookinfo, book)

		fmt.Println("bookinfo ", bookinfo)
	}

	t, err := template.ParseFiles("./views/index.html")

	if err != nil {
		fmt.Println("Error in parsing home.html", err)
	}

	t.Execute(w, bookinfo)
}
