package auth

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

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

	// ลงทะเบียนประเภท User กับ gob
	gob.Register(User{})
}

var store = sessions.NewCookieStore([]byte("secret"))

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedUser User
	var storedHashedPassword string
	// ดึงข้อมูลผู้ใช้และรหัสผ่านที่เข้ารหัสจากฐานข้อมูล
	err = db.QueryRow("SELECT id, name, username, email, password FROM users WHERE username = ?", user.Username).
		Scan(&storedUser.ID, &storedUser.Name, &storedUser.Username, &storedUser.Email, &storedHashedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// ตรวจสอบรหัสผ่านที่ผู้ใช้ส่งมา
	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// สร้างเซสชัน
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["user"] = storedUser
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// สร้างโครงสร้างข้อมูลสำหรับการตอบกลับ (ยกเว้นรหัสผ่าน)
	response := map[string]interface{}{
		"status": "ยืนยัน",
		"user":   storedUser.Name, // ส่งข้อมูลผู้ใช้ (ไม่รวมรหัสผ่าน)
	}

	// ส่งการตอบกลับ JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = createUserInDatabase(user, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ส่งข้อมูลผู้ใช้กลับไปยัง client (ยกเว้นรหัสผ่าน)
	user.Password = "" // ไม่ให้ส่งรหัสผ่านในการตอบกลับ
	json.NewEncoder(w).Encode(user)
}

// hashPassword ใช้สำหรับเข้ารหัสรหัสผ่าน
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// createUserInDatabase ใช้สำหรับบันทึกผู้ใช้ใหม่ลงในฐานข้อมูล
func createUserInDatabase(user User, hashedPassword string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)", user.Name, user.Username, user.Email, hashedPassword)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// Logout ใช้สำหรับลบเซสชันของผู้ใช้ และออกจากระบบ
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ลบข้อมูลผู้ใช้ใน session
	delete(session.Values, "user")

	// บันทึก session หลังจากลบข้อมูล
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ส่งการตอบกลับ JSON ว่า logout สำเร็จ
	response := map[string]string{
		"status": "logout successful",
	}

	// ส่งการตอบกลับ JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
