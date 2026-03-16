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
func GetCategories(w http.ResponseWriter, r*http.Request){
	rows, err := db.DB.Query("SELECT name,decscription  FROM categories")
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
return
	}
	defer rows.Close()
	var categories []models.Categories
	for rows.Next(){
		var c models.Categories
		err := rows.Scan(&c.ID, &c.NAME, &c.DECSCIPTION)
		if err !=nil {
			http.Error(w,err.Error(),  http.StatusInternalServerError)
			return
		}


		categories = append(categories, c)
	}
	json.NewEncoder(w).Encode(categories)
}
func CreateCategories(w http.ResponseWriter, r *http.Request) {

	var c models.Categories

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(
		"INSERT INTO categories(name,description) VALUES(?,?)",
		c.NAME,
		c.DECSCIPTION,
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

	c.ID = int(id)

	json.NewEncoder(w).Encode(c)
}

func DeleteCategories(w http.ResponseWriter, r *http.Request) {

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

func UpdateCategories(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var c models.Categories

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"UPDATE categories SET ,name=?,description=? WHERE id=?",
		
		c.NAME,
		c.DECSCIPTION,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.ID, _ = strconv.Atoi(id)

	json.NewEncoder(w).Encode(c)
}