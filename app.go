package main

import (
	"encoding/json"
    "encoding/hex"
	"math/rand"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

const (
    host     = "db"
    port     = 5432
    user     = "admin"
    password = "admin"
    dbname   = "admin"
)

// Nasabah Struct(Model)
type Nasabah struct {
	Id			int  	`json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Token		string 	`json:"token"`
	Tabungan	int 	`json:"tabungan"`
}
// Report Struct(Model)
type Report struct {
	dataType 	string `json:"type"`
	Message		string `json:"message"`
}

//Login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
		return
	}
	var nasabah Nasabah
	// unmarshal
	if err := json.Unmarshal(reqData, &nasabah); err != nil {
		panic(err)
	}
	CheckError(err)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)

	selectUser, _, _ := goqu.From("nasabah").Where(goqu.Ex{"username": nasabah.Username, "password": nasabah.Password}).ToSQL()
	rows, err := db.Query(selectUser)
	CheckError(err)

	var report []Report

	defer rows.Close()
	for rows.Next() {
		var id int
		var username string
		var password string
		var token string
		var tabungan int

	
		err = rows.Scan(&id, &username, &password, &token, &tabungan)
		CheckError(err)
		if token == "" {
			goQuery, _, _ := goqu.Update("nasabah").Set(goqu.Record{"token": GenerateSecureToken(5)}).Where(goqu.Ex{"username": nasabah.Username}).ToSQL()
			_, err := db.Query(goQuery)
			CheckError(err)
			report = append(report, Report{dataType:"login", Message: "sukses"})
		} else {
			report = append(report, Report{dataType:"login", Message: "user already login"})
		}
	}
	if report == nil {
		report = append(report, Report{dataType:"login", Message: "no data"})
	}
	json.NewEncoder(w).Encode(report)
}

//Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
		return
	}
	var nasabah Nasabah
	// unmarshal
	if err := json.Unmarshal(reqData, &nasabah); err != nil {
		panic(err)
	}
	CheckError(err)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)

	selectUser := fmt.Sprintf("SELECT * FROM nasabah WHERE username= '%s' AND password='%s'", nasabah.Username, nasabah.Password)
	rows, err := db.Query(selectUser)
	CheckError(err)

	var report []Report

	defer rows.Close()
	for rows.Next() {
		var id int
		var username string
		var password string
		var token string
		var tabungan int

	
		err = rows.Scan(&id, &username, &password, &token, &tabungan)
		CheckError(err)
		if token != "" {
			goQuery, _, _ := goqu.Update("nasabah").Set(goqu.Record{"token": ""}).Where(goqu.Ex{"username": nasabah.Username}).ToSQL()
			_, err := db.Query(goQuery)
			CheckError(err)
			report = append(report, Report{dataType:"logout", Message: "success"})
		} else {
			report = append(report, Report{dataType:"logout", Message: "user already logout"})
		}
	}
	if report == nil {
		report = append(report, Report{dataType:"logout", Message: "no data"})
	}
	json.NewEncoder(w).Encode(report)
}

func CheckError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

func GenerateSecureToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func main() {
	//init router
	r:= mux.NewRouter()

	r.HandleFunc("/api/login", Login).Methods("GET")
	r.HandleFunc("/api/logout", Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
