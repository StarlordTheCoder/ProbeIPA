package model

import (
	"database/sql"
	"server/survey/model/answers"
)

func GetAnswers(idSurvey int, db *sql.DB) (*answers.GivenAnswers, error) {
	stmt, err := db.Prepare("SELECT Choices.idChoices, Choices.choice FROM Survey INNER JOIN Choices ON Survey.IdSurvey = Choices.Survey_idSurvey WHERE Survey.IdSurvey=?")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query(idSurvey)
	if err != nil {
		return nil, err
	}
	newanswers := []answers.GivenAnswer{}
	for res.Next() {
		var amount int
		var idChoice int
		var choice string
		err = res.Scan(&idChoice, &choice)
		if err != nil {
			return nil, err
		}
		stmt, err := db.Prepare("SELECT COUNT(Answers.Choices_idChoices) FROM Answers where Answers.Choices_idChoices")
		if err != nil {
			return nil, err
		}
		res2, err := stmt.Query(idChoice)
		if err != nil {
			return nil, err
		}
		err = res2.Scan(&amount)
		if err != nil {
			return nil, err
		}
		newanswers = append(newanswers, answers.GivenAnswer{IdChoice: idChoice, Choice: choice, Amount: amount})
	}
	return &answers.GivenAnswers{newanswers}, nil
}

func AnswerSurvey(newAnswer *answers.NewAnswers, db *sql.DB) error {
	tx, _ := db.Begin()
	for _, answer := range newAnswer.NewAnswers {
		stmt, err := db.Prepare("INSERT Result SET Choices_idChoices=?")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmt.Exec(answer.IDChoice)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
