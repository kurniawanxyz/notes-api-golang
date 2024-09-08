package main

import (
	"log"
	coreHttp "net/http"

	"github.com/gorilla/mux"

	"notes/delivery/http"
	"notes/domain"
	"notes/repository"
	"notes/usecase"
	"notes/utils"
)

func main() {
	db, err := utils.DBConnect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	db.AutoMigrate(&domain.Note{})

	noteRepo := repository.NewMySQLNoteRepository(db)
	noteUC := usecase.NewNoteUsecase(noteRepo)
	bookHandler := http.NewNoteHandler(noteUC)

	r := mux.NewRouter()
	r.HandleFunc("/notes", bookHandler.Get).Methods("GET")
	r.HandleFunc("/notes/{id}", bookHandler.Show).Methods("GET")
	r.HandleFunc("/notes", bookHandler.Store).Methods("POST")
	r.HandleFunc("/notes/{id}", bookHandler.Update).Methods("PUT")
	r.HandleFunc("/notes/{id}", bookHandler.Delete).Methods("DELETE")

	log.Println("SERVER RUNNING ON PORT 8000")
	log.Fatal(coreHttp.ListenAndServe(":8000", r))
}
