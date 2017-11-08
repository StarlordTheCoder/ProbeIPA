package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/authorisation"
	"server/database"
	"server/survey/controler2"
	"server/survey/model/answers"
	"server/survey/model/survey"
	"server/user/controler"
	"server/user/model/user"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var db database.Database

func Register(res http.ResponseWriter, req *http.Request) {
	fmt.Println("d")
	decoder := json.NewDecoder(req.Body)
	var user user.User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("c")
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = controler.RegisterUser(user.Username, user.Password, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("d")
	decoder := json.NewDecoder(req.Body)
	var user user.User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("a")
		http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	_, err = authorisation.Authorisate(user.Username, user.Password, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func CreateSurvey(res http.ResponseWriter, req *http.Request) {
	fmt.Println("d")
	decoder := json.NewDecoder(req.Body)
	var newSurvey survey.Survey
	err := decoder.Decode(&newSurvey)
	if err != nil {
		fmt.Println("a")
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	username := req.Header.Get("username")
	password := req.Header.Get("password")
	id, err := authorisation.Authorisate(username, password, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	result, err := controler2.CreateSurvey(&newSurvey, *id, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(*result)
}

func GetSurvey(res http.ResponseWriter, req *http.Request) {
	fmt.Println("d")
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	result, err := controler2.GetSurvey(id, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(*result)
}

func GetSurveyByUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("d")
	username := req.Header.Get("username")
	password := req.Header.Get("password")
	id, err := authorisation.Authorisate(username, password, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	fmt.Println(*id)
	result, err := controler2.GetSurveyByUser(*id, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(*result)
}

func AnswerSurvey(res http.ResponseWriter, req *http.Request) {
	fmt.Println("d")
	decoder := json.NewDecoder(req.Body)
	var newAnswers answers.NewAnswers
	err := decoder.Decode(&newAnswers)
	if err != nil {
		fmt.Println("a")
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = controler2.SetAnswers(&newAnswers, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func GetAnswers(res http.ResponseWriter, req *http.Request) {
	fmt.Println("d")
	decoder := json.NewDecoder(req.Body)
	var Survey survey.Survey
	err := decoder.Decode(&Survey)
	if err != nil {
		fmt.Println("a")
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	username := req.Header.Get("username")
	password := req.Header.Get("password")
	_, err = authorisation.Authorisate(username, password, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	result, err := controler2.GetAnswers(Survey.Id, db.Db)
	if err != nil {
		fmt.Println(err)
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(*result)
}

func main() {
	db = database.Database{}
	err := db.Connect("survey_db", "root", "")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/login", Login).Methods("POST")
	r.HandleFunc("/api/register", Register).Methods("POST")
	r.HandleFunc("/api/createSurvey", CreateSurvey).Methods("POST")
	r.HandleFunc("/api/getSurvey/{id}", GetSurvey).Methods("GET")
	r.HandleFunc("/api/getSurveyByUser", GetSurveyByUser).Methods("GET")
	r.HandleFunc("/api/answerSurvey", AnswerSurvey).Methods("POST")
	r.HandleFunc("/api/getAnswers", GetAnswers).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"content-type", "username", "password"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{"*"})
	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk)(r))
}
