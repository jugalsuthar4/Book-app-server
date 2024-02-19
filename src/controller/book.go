package controller

import (
	"book-app-server/src/data"
	"book-app-server/src/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const INVALID_BOOK_ID = "Invalid Book Id"
const INVALID_REQUEST_BODY = "invalid request body"
const BOOK_NOT_FOUND = "Book not found"

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newBook)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBook.ID = len(data.Books) + 1

	data.Books = append(data.Books, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		sendErrorResponse(w, INVALID_BOOK_ID, http.StatusBadRequest)
		return
	}

	result := findBookByID(id)
	if !result.Success {
		sendErrorResponse(w, result.Error, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		sendErrorResponse(w, INVALID_BOOK_ID, http.StatusBadRequest)
		return
	}

	result := findBookIndexByID(id)
	if !result.Success {
		sendErrorResponse(w, result.Error, http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&updatedBook)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.Books[result.Data.(int)] = updatedBook

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		sendErrorResponse(w, INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	result := findBookIndexByID(id)
	if !result.Success {
		sendErrorResponse(w, result.Error, http.StatusNotFound)
		return
	}

	bookIndex := result.Data.(int)
	data.Books = append(data.Books[:bookIndex], data.Books[bookIndex+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

func findBookByID(id int) models.Result {
	for _, book := range data.Books {
		if book.ID == id {
			return models.Result{Success: true, Data: book}
		}
	}
	return models.Result{Success: false, Error: BOOK_NOT_FOUND}
}

func findBookIndexByID(id int) models.Result {
	for i, book := range data.Books {
		if book.ID == id {
			return models.Result{Success: true, Data: i}
		}
	}
	return models.Result{Success: false, Error: BOOK_NOT_FOUND}
}

func sendErrorResponse(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := models.ErrorResponse{Error: errorMessage}
	json.NewEncoder(w).Encode(errorResponse)
}
