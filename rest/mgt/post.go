package mgt

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Post struct manage all REST operations related with Post.
type Post struct{}

// NewPost returns a new Post struct.
func NewPost() *Post {
	return &Post{}
}

// GetPost returns a Post.
func (p Post) GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post := struct {
		Title string `json:"title"`
	}{
		Title: fmt.Sprintf("REST exercise %s", params["id"]),
	}

	respondWithJSON(w, http.StatusOK, post)
}
