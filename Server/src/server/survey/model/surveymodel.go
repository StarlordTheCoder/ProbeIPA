package model

import (
	"database/sql"
	"server/survey/model/survey"
	"fmt"
)

func CreateSurvey(newSurvey *survey.Survey, idUser int, db *sql.DB) (*int, error) {
	tx, err := db.Begin()
	stmt, err := db.Prepare("INSERT Survey SET question=?, User_idUser=?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res, err := stmt.Exec(newSurvey.Question, idUser)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	id, _ := res.LastInsertId()
	idint := int(id)
	for _, choice := range newSurvey.Choices {
		stmt, err := db.Prepare("INSERT Choices SET choice=?, Survey_idSurvey=?")
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		_, err = stmt.Exec(choice.Choice, idint)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return &idint, nil
}

func GetSurvey(id int, db *sql.DB) (*survey.Survey, error) {
	stmt, err := db.Prepare("SELECT Survey.IdSurvey, Survey.question, Choices.idChoices, Choices.choice FROM Survey INNER JOIN Choices ON Survey.IdSurvey = Choices.Survey_idSurvey WHERE Survey.IdSurvey=?")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	var sid int
	var question string
	choices := []survey.Choice{}
	for res.Next() {
		var cid int
		var choice string
		err = res.Scan(&sid, &question, &cid, &choice)
		if err != nil {
			return nil, err
		}
		choices = append(choices, survey.Choice{Id: cid, Choice: choice})
	}
	return &survey.Survey{Id: sid, Question: question, Choices: choices}, nil
}

func GetSurveyByUser(idUser int, db *sql.DB) (*[]survey.Survey, error) {
	stmt, err := db.Prepare("SELECT Survey.IdSurvey, Survey.question FROM Survey WHERE Survey.User_idUser=?")
	if err != nil {
		return nil, err
	}
	fmt.Println(idUser)
	res, err := stmt.Query(idUser)
	if err != nil {
		return nil, err
	}
	surveys := []survey.Survey{}
	for res.Next() {
		var sid int
		var question string
		err = res.Scan(&sid, &question)
		if err != nil {
			return nil, err
		}
		surveys = append(surveys, survey.Survey{Id: sid, Question: question, Choices: []survey.Choice{}})
	}
	return &surveys, nil
}
