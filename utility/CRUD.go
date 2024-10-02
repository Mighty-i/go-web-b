package utility

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/gowebpro")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// ดึง userName จากพารามิเตอร์ของ query
	userName := r.URL.Query().Get("userName")
	if userName == "" {
		http.Error(w, "จำเป็นต้องมี userName", http.StatusBadRequest)
		return
	}

	// ดึง user ID จากตาราง users
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE name = ?", userName).Scan(&userID)
	if err != nil {
		http.Error(w, "ข้อผิดพลาดในการดึง user ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, title, done FROM todos WHERE name_id = ? ", userID)
	if err != nil {
		http.Error(w, "ข้อผิดพลาดในการดึง todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if err != nil {
			http.Error(w, "ข้อผิดพลาดในการสแกน todo: "+err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "ข้อผิดพลาดในการวนลูปแถว: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Title    string
		UserName string
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new Todo instance
	newTodo := Todo{
		Title:  requestBody.Title,
	}

	// Retrieve the user ID from the users table
	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE name = ?", requestBody.UserName).Scan(&userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("INSERT INTO todos (title, done, name_id) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(newTodo.Title, false, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newTodo)
}

func UpdateDone(w http.ResponseWriter, r *http.Request) {
	var updatedTodo Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	stmt, err := db.Prepare("UPDATE todos SET done = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(updatedTodo.Done, updatedTodo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

func Updatetitle(w http.ResponseWriter, r *http.Request) {
	var updatedTodo Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	stmt, err := db.Prepare("UPDATE todos SET title = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(updatedTodo.Title, updatedTodo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var updatedTodo Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = db.Exec("DELETE FROM todos WHERE id = ?", updatedTodo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
