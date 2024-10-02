package main

import (
	"net/http"

	"project/auth"
	"project/googlelogin"
	"project/utility"

	"github.com/gorilla/mux"
)

func main() {
	authRouter := mux.NewRouter()
	authRouter.HandleFunc("/register", corsHandler(auth.Register))
	authRouter.HandleFunc("/login", corsHandler(auth.Login))
	authRouter.HandleFunc("/logout", corsHandler(auth.Logout))
	authRouter.HandleFunc("/check-google-user", corsHandler(googlelogin.CheckGoogleUser))
	go http.ListenAndServe(":8088", authRouter)

	todosRouter := mux.NewRouter()
	todosRouter.HandleFunc("/api/todos", corsHandler(utility.GetTodos))
	todosRouter.HandleFunc("/api/CreateTodo", corsHandler(utility.CreateTodo))
	todosRouter.HandleFunc("/api/UpdateDone", corsHandler(utility.UpdateDone))
	todosRouter.HandleFunc("/api/updatetitle", corsHandler(utility.Updatetitle))
	todosRouter.HandleFunc("/api/DeleteTodo", corsHandler(utility.DeleteTodo))
	http.ListenAndServe(":8080", todosRouter)
}

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	}
}
