package handlers

import (
	"be-library-go/db"
	"be-library-go/models"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

// get category
func GetBooks(w http.ResponseWriter, r*http.Request){
	rows, err := db.DB.Query("SELECT title,tatgia,category_id,total_copies,available_copies  FROM books")
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
return
	}
	defer rows.Close()
	var books []models.Books
	for rows.Next(){
		var b models.Books
		err := rows.Scan(&b.ID, &b.TITLE, &b.TACGIA, &b.CATEGORYID, &b.TOTAL_COPIES, &b.AVAILABLE_COPIES)
		if err !=nil {
			http.Error(w,err.Error(),  http.StatusInternalServerError)
			return
		}


		books = append(books, b)
	}
	json.NewEncoder(w).Encode(books)
}
func CreateBooks(w http.ResponseWriter, r *http.Request) {

	var b models.Books

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(
		"INSERT INTO categories(title,tatgia,category_id,total_copies,available_copies) VALUES(?,?,?,?,?)",
	
		b.TITLE,
		b.TACGIA,
		b.CATEGORYID,
		b.TOTAL_COPIES,
		b.AVAILABLE_COPIES,


	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b.ID = int(id)

	json.NewEncoder(w).Encode(b)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM categories WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted",
	})
}

func UpdateBooks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var b models.Books

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"UPDATE categories SET ,title=?,tatgia=?,category_id=?,total_copies=?,available_copies=? WHERE id=?",
	
		b.TITLE,
		b.TACGIA,
		b.CATEGORYID,
		b.TOTAL_COPIES,
		b.AVAILABLE_COPIES,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b.ID, _ = strconv.Atoi(id)

	json.NewEncoder(w).Encode(b)
}