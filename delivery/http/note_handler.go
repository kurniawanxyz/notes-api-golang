package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"notes/domain"
	"notes/usecase"
)

type NoteHandler struct {
	uc *usecase.NoteUsecase
}

func NewNoteHandler(uc *usecase.NoteUsecase) *NoteHandler {
	return &NoteHandler{uc}
}

func (h *NoteHandler) Get(res http.ResponseWriter, req *http.Request) {
	books, err := h.uc.Get()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(books)
}

func (h *NoteHandler) Show(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	book, err := h.uc.Show(uint(id))
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(res).Encode(book)
}

func (h *NoteHandler) Store(res http.ResponseWriter, req *http.Request) {
	var note domain.Note
	json.NewDecoder(req.Body).Decode(&note)
	if err := h.uc.Store(&note); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(note)
}

func (h *NoteHandler) Update(res http.ResponseWriter, req *http.Request) {
	var note domain.Note
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	_, err := h.uc.Show(uint(id))
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	json.NewDecoder(req.Body).Decode(&note)

	if err := h.uc.Update(&note); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(&note)
}

func (h *NoteHandler) Delete(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	if err := h.uc.Delete(uint(id)); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode("Book Deleted")
}
