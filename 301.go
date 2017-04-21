package main

import (
  "net/http"
  "net/url"
  "fmt"
  "github.com/gorilla/mux"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "os"
)

const AppName = "301 Tracker" //
const Version = "0.1"         //
const DefaultNetwork = ":8080"       // eth listen interface

type DBData struct {
  driver string
  source string
}

var DB DBData

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", Index)
  r.HandleFunc("/link", LinkRedirect)
  DB.driver = os.Getenv("APP_DB_DRIVER")
  DB.source = os.Getenv("APP_DB_SOURCE")
  network := os.Getenv("APP_NETWORK")
  if len(network) == 0 {
    network = DefaultNetwork
    _ = network
  }
  dbStatus := checkDB()
  if  dbStatus == "ok" {
    fmt.Printf("Starting `%s`\n - Network  : %s\n - DB.driver: %s\n - DB.source: %s\n", AppName, network, DB.driver, DB.source)
    checkErr(http.ListenAndServe(network, r))
  } else {
    fmt.Printf("%+v\nExit\n", dbStatus)
  }
}

// Index handle GET / and reply with APP infos
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s [v%s]\n", AppName, Version)
  fmt.Fprintf(w, "db_status: %s\n", checkDB())
}

// LinkRedirect handle GET /link?url=WEBSITE and perform a 301 redirect to WEBSITE
func LinkRedirect(w http.ResponseWriter, r *http.Request) {
  gotoUrl := r.URL.Query().Get("url")
  a := r.URL.Query().Get("a")
  b := r.URL.Query().Get("b")
  c := r.URL.Query().Get("c")
  d := r.URL.Query().Get("d")
  e := r.URL.Query().Get("e")
  if checkUrl(gotoUrl) {
    Store(gotoUrl, a, b, c, d, e)
    http.Redirect(w, r, gotoUrl, 301)
  } else {
    w.WriteHeader(404)
    fmt.Fprintf(w, "404 Not Found")
  }
}


// Store save data on DB
func Store(urlLink, a, b, c, d, e string){
  db, err := sql.Open(DB.driver, DB.source)
  checkErr(err)
  stmt, err := db.Prepare("INSERT INTO redirect (url, a, b, c, d, e, date) VALUES (?, ?, ?, ?, ?, ?, NOW());")
  checkErr(err)
  res, err := stmt.Exec(urlLink, a, b, c, d, e)
  _ = res
  checkErr(err)
}

// checkUrl check if urlLink is a valid website
func checkUrl(urlLink string)(bool){
  _, err := url.ParseRequestURI(urlLink)
  if err != nil {
     return false
  }
  return true
}

// checkDB check if the DB parameters works
func checkDB()(string){
  if DB.driver == "" {
    return "env.APP_DB_DRIVER not set"
  }
  if DB.source == "" {
    return "env.APP_DB_SOURCE not set"
  }
  db, err := sql.Open(DB.driver, DB.source)
  if err != nil {
    checkErr(err)
    return "error, check your connection string"
  }
  stmt, err := db.Prepare("SELECT id FROM redirect LIMIT 0,1")
  if err != nil {
    checkErr(err)
    return "error on prepare, check your db structure"
  }
  res, err := stmt.Exec()
  if err != nil {
    checkErr(err)
    return "error on exec, check your db structure"
  }
  _ = res
  return "ok"
}

// checkErr wrap every error check
func checkErr(err error) {
  if err != nil {
    fmt.Printf("%+v\n", err)
  }
}
