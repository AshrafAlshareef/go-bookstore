package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	response, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in the GetBookById function")
	}
	bookDetails, _ := models.GetBookById(ID)
	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Context-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	response, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in the DeleteBook function")
	}
	book := models.DeleteBook(ID)
	response, _ := json.Marshal(book)
	w.Header().Set("Context-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in the UpdateBook function")
	}
	bookDetails, db := models.GetBookById(ID)
	if bookDetails.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if bookDetails.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if bookDetails.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Context-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
