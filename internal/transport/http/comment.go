package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guillermogrillo/comments-api/internal/comment"
)

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Could not parse the supplied ID to uint", err)
		return
	}
	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by id", err)
		return
	}
	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error retrieving all comments", err)
		return
	}
	if err = sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed decoding incoming request", err)
		return
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Error retrieving all comments", err)
		return
	}
	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed decoding incoming request", err)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Could not parse the supplied ID to uint", err)
		return
	}
	comment, err = h.Service.UpdateComment(uint(i), comment)
	if err != nil {
		sendErrorResponse(w, "Error retrieving all comments", err)
		return
	}
	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Could not parse the supplied ID to uint", err)
		return
	}
	err = h.Service.DeleteComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by id", err)
		return
	}
	if err = sendOkResponse(w, Response{Message: "Comment deleted succesfully"}); err != nil {
		panic(err)
	}
}
