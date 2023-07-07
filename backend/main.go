package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
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
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	Email        *string `json:"email,omitempty"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Organization *string `json:"organization,omitempty"`
	JobTitle     *string `json:"job_title,omitempty"`
}

type PersonResponse struct {
	Applicants []Person `json:"applicants"`
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

func (s *Server) dbFetch(dbStmt string, conditions map[string][]string) []Person {
	//_, e := s.db.Exec(dbStmt, ...args)
	//CheckError(e)
	fetchStmt := `Select * from person`
	var keys = make([]string, 0)
	var args = make([]interface{}, 0)
	if len(conditions) > 0 {
		for key, val := range conditions {
			keys = append(keys, key)
			args = append(args, val[0])
		}
		fetchStmt += " where "
	}
	var lk = len(keys)

	for i, key := range keys {
		if i < lk-1 {
			fetchStmt = fetchStmt + key + "=$" + strconv.Itoa(i+1) + " and "
		} else {
			fetchStmt = fetchStmt + key + "=$" + strconv.Itoa(i+1)
		}
	}
	fmt.Println(fetchStmt)
	rows, err := s.db.Query(fetchStmt, args...)
	CheckError(err)

	var persons = make([]Person, 0)

	defer rows.Close()
	for rows.Next() {
		var person Person
		var ID int
		err = rows.Scan(&ID, &person.FirstName, &person.LastName, &person.Email, &person.PhoneNumber, &person.Organization, &person.JobTitle)
		CheckError(err)
		persons = append(persons, person)
	}
	return persons
}

func main() {

	server := newServer()
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

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Person: %+v", p)
	s.dbInsert(insertUserStmt, p.FirstName, p.LastName, p.Email, p.PhoneNumber, p.Organization, p.JobTitle)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(`{"message": "success"}`))
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}
	userData := s.dbFetch(fetchUserStmt, values)
	fmt.Printf("%+v\n", userData)

	applicants := PersonResponse{
		Applicants: userData,
	}
	userDataBytes, err := json.Marshal(applicants)
	CheckError(err)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(userDataBytes)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("error ", err)
	}
}
