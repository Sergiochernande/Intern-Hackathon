package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "abar"
	password = "odoo"
	dbname   = "it3"
)

var (
	insertUserStmt = `insert into person ("first_name", "last_name", "email", "phone_number", "organization", "job_title") values($1, $2, $3, $4, $5, $6)`
	fetchUserStmt  = `select * from person where first_name=$1 and last_name=$2`
)

type Server struct {
	router *mux.Router
	db     *sql.DB
}

type Person struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Organization string `json:"organization"`
	JobTitle     string `json:"job_title"`
}

func newServer() *Server {
	r := mux.NewRouter()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	sqlConn, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	server := Server{router: r, db: sqlConn}
	server.addRoutes()
	return &server
}

func (s *Server) addRoutes() {
	s.router.HandleFunc("/users", s.createUser).Methods("POST")
	s.router.HandleFunc("/users", s.getUser).Methods("GET")
	s.router.HandleFunc("/users", s.handleOptions).Methods("OPTIONS")
}

func (s *Server) handleOptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received options request")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(`{"message": "success"}`))
}

func (s *Server) dbInsert(insertStmt string, args ...interface{}) {
	//insertDynStmt := `insert into Students ("name", "roll") values($1, $2)`
	_, e := s.db.Exec(insertStmt, args...)
	CheckError(e)
}

func (s *Server) dbFetch(dbStmt string, args ...interface{}) {
	_, e := s.db.Exec(dbStmt, args...)
	CheckError(e)
	rows, err := s.db.Query(dbStmt, args...)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var roll int

		err = rows.Scan(&name, &roll)
		CheckError(err)

		fmt.Println(name, roll)
	}
}

func main() {

	server := newServer()
	// insert
	// hardcoded
	//insertStmt := `insert into Students ("name", "roll") values('John', 1)`
	//_, e := sqlConn.Exec(insertStmt)
	//CheckError(e)
	//
	//// dynamic
	//insertDynStmt := `insert into Students ("name", "roll") values($1, $2)`
	//_, e = sqlConn.Exec(insertDynStmt, "Jane", 2)
	//CheckError(e)
	srv := &http.Server{
		Handler: server.router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println(srv.ListenAndServe())
	defer server.db.Close()
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var p Person

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	fmt.Printf("Person: %+v", p)
	s.dbInsert(insertUserStmt, p.FirstName, p.LastName, p.Email, p.PhoneNumber, p.Organization, p.JobTitle)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(`{"message": "success"}`))
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	//var p Person
	//values := r.URL.Query()
	//for k, v := range values {
	//	fmt.Println(k, " => ", v)
	//}
	//userData := s.dbFetch(fetchUserStmt, values["fist_name"], values["last_name"])
	//fmt.Printf("%+v\n", userData)

}

func CheckError(err error) {
	if err != nil {
		fmt.Println("error ", err)
	}
}
