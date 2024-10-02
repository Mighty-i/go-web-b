package googlelogin

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CheckGoogleUser(w http.ResponseWriter, r *http.Request) {
	var user struct {
		GoogleID string `json:"google_id"`
		Email    string `json:"email"`
		Name     string `json:"username"`
		Image    string `json:"profile_image"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gowebpro")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var id int
	err = db.QueryRow("SELECT id FROM users WHERE google_id = ? OR email = ? OR name = ?", user.GoogleID, user.Email, user.Name).Scan(&id)

	// If no rows are returned, insert new user
	if err == sql.ErrNoRows {
		_, err = db.Exec("INSERT INTO users (google_id, email, name, profile_image, created_at) VALUES (?, ?, ?, ?, ?)",
			user.GoogleID, user.Email, user.Name, user.Image, time.Now())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ยืนยัน"})
		return
	}

	// If an error occurred while querying the database, return the error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "ยืนยัน"})
}
