package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhiongithub/publicationmgmt/internal/models"
	"github.com/abhiongithub/publicationmgmt/internal/services"
	"github.com/gorilla/mux"
)

//Handler- stores pointer to publication service
type Handler struct {
	Router  *mux.Router
	Service *services.Service
}

// Response objecgi
type Response struct {
	Message string
	Error   string
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/book", h.GetAllBooks).Methods("GET")
	h.Router.HandleFunc("/api/book", h.PostBook).Methods("POST")
	h.Router.HandleFunc("/api/book/{id}", h.GetBookById).Methods("GET")
	h.Router.HandleFunc("/api/book/{id}", h.UpdateBook).Methods("PUT")
	h.Router.HandleFunc("/api/book/{id}", h.DeleteBook).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Healthy")
	})
}

// GetBookById -
func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprint(w, "unable to parse UINT from id")
	}

	book, err := h.Service.GetBookById(uint(i))

	if err != nil {
		fmt.Fprint(w, "Error retriving book by id")
	}

	fmt.Fprintf(w, "%+v", book)
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllBooks()
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all books", err)
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

// PostComment - adds a new comment
func (h *Handler) PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
	}

	comment, err := h.Service.PostBook(book)
	if err != nil {
		sendErrorResponse(w, "Failed to post new comment", err)
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// UpdateBook- updates a book by Id
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	bookId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from ID", err)
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
	}

	book, err = h.Service.UpdateBook(uint(bookId), book)
	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
	}
	if err := json.NewEncoder(w).Encode(book); err != nil {
		panic(err)
	}
}

// DeleteBook - deletes a book by Id
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from ID", err)
	}

	err = h.Service.DeleteBook(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete book by book id", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
