package mainfunctions

import (
	"context"

	"github.com/Ali-Farhadnia/clientGRPC/config"
	"github.com/Ali-Farhadnia/clientGRPC/models/book"
	"github.com/Ali-Farhadnia/clientGRPC/models/modelpb"
)

//all the functions get string as input then pars that string and send it to grpc server

func InsertOneBook(input string) (string, error) {
	var book book.Book
	err := book.UnarshalString(input)
	if err != nil {
		return "unvalid input", err
	}
	grpcbooklist := []*modelpb.Book{}
	grpcbooklist = append(grpcbooklist, &modelpb.Book{Name: book.Name,
		Author: book.Author, Pagescount: book.Pagecount, Inventory: book.Inventory})
	var grpcbooks = modelpb.Books{
		Books: grpcbooklist,
	}
	Status, err := config.App.Grpc_CRUD_client.InsertBook(context.Background(), &grpcbooks)
	if err != nil {
		return "somthing whent wrong", err
	}

	if Status.Status == "no" {
		return "somthing whent wrong:" + Status.Description, err
	}
	return Status.Description, nil
}
func InsertManyBooks(input string) (string, error) {
	var books book.Books
	err := books.UnarshalString(input)
	if err != nil {
		return "unvalid input", err
	}
	grpcbooklist := []*modelpb.Book{}
	for _, book := range books.Books {

		grpcbooklist = append(grpcbooklist, &modelpb.Book{Name: book.Name,
			Author: book.Author, Pagescount: book.Pagecount, Inventory: book.Inventory})
	}
	var grpcbooks = modelpb.Books{
		Books: grpcbooklist,
	}
	Status, err := config.App.Grpc_CRUD_client.InsertBook(context.Background(), &grpcbooks)
	if err != nil {
		return "somthing whent wrong", err
	}
	if Status.Status == "no" {
		return "somthing whent wrong:" + Status.Description, err
	}
	return Status.Description, nil
}
func UpdateBook(input string) (string, error) {
	var book book.Book
	err := book.UnarshalString(input)
	if err != nil {
		return "unvalid input", err
	}
	rec := modelpb.UpdateRequest{
		Book: &modelpb.Book{Name: book.Name,
			Author:     book.Author,
			Pagescount: book.Pagecount,
			Inventory:  book.Inventory}, Id: book.ID,
	}

	Status, err := config.App.Grpc_CRUD_client.UpdateBook(context.Background(), &rec)
	if err != nil {
		return "somthing whent wrong", err
	}
	if Status.Status == "no" {
		return "somthing whent wrong:" + Status.Description, err
	}
	return Status.Description, nil

}
func DeleteBook(input string) (string, error) {
	rec := modelpb.BookID{Id: input}
	Status, err := config.App.Grpc_CRUD_client.DeleteBook(context.Background(), &rec)
	if err != nil {
		return "somthing whent wrong", err
	}
	if Status.Status == "no" {
		return "somthing whent wrong:" + Status.Description, err
	}
	return Status.Description, nil
}
func FindBookByID(input string) (string, error) {
	rec := modelpb.BookID{Id: input}
	res, err := config.App.Grpc_CRUD_client.FindBookById(context.Background(), &rec)
	if err != nil {
		return "somthing whent wrong", err
	}
	if res.Status.Status == "no" {
		return "somthing whent wrong:" + res.Status.Description, err
	}
	resbook := book.Book{
		ID:        res.Book.Id,
		Name:      res.Book.Name,
		Author:    res.Book.Author,
		Pagecount: res.Book.Pagescount,
		Inventory: res.Book.Inventory,
	}

	sres, err := resbook.ToString()
	if err != nil {
		return "somthing whent wrong", err
	}
	return sres, nil

}

//each time  Help(string) called it returns help string
func Help(string) (string, error) {
	help := `
	Functions:

		InsertOneBook:		e.g.#InsertOneBook {"name":string,"author":string,"pagecount": integer,"Inventory":integer}

		InsertManyBooks:        e.g.#InsertManyBooks {"books"[{"name":string,"author":string,"pagecount": integer,"Inventory":integer},
								{"name":string,"author":string,"pagecount": integer,"Inventory":integer}]}

		UpdateBook:		e.g.#UpdateBook {"id":string,"name":string,"author":string,"pagecount": integer,"Inventory":integer}

		DeleteBook:       	e.g.#DeleteBook string

		FindBookByID:           e.g.#FindBookByID string
	`
	return help, nil

}
