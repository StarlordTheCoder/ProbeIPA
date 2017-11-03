package main 

import (
	"net/http"
    //"github.com/gorilla/mux"
    "server/database"
    "fmt"
)

func homePage(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "index.html")
}

func main() {
    db := database.Database{}
    err := db.Connect("survey_db","root","")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()
    http.HandleFunc("/", homePage)
    http.ListenAndServe(":8080", nil)
}