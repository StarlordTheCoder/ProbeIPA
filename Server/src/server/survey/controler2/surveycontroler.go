package controler2

import (
	"database/sql"
	"encoding/json"
	"server/survey/model"
	"server/survey/model/answers"
	"server/survey/model/survey"
)

func CreateSurvey(newSurvey *survey.Survey, idUser int, db *sql.DB) (*[]byte, error) {
	idSurvey, err := model.CreateSurvey(newSurvey, idUser, db)
	if err != nil {
		return nil, err
	}
	idSurveyJSON, err := json.Marshal(idSurvey)
	if err != nil {
		return nil, err
	}
	return &idSurveyJSON, nil
}

func GetSurvey(idSurvey int, db *sql.DB) (*[]byte, error) {
	survey, err := model.GetSurvey(idSurvey, db)
	if err != nil {
		return nil, err
	}
	surveyJSON, err := json.Marshal(survey)
	if err != nil {
		return nil, err
	}
	return &surveyJSON, nil
}

func GetSurveyByUser(idUser int, db *sql.DB) (*[]byte, error) {
	surveys, err := model.GetSurvey(idUser, db)
	if err != nil {
		return nil, err
	}
	surveysJSON, err := json.Marshal(surveys)
	if err != nil {
		return nil, err
	}
	return &surveysJSON, nil
}

func GetAnswers(idSurvey int, db *sql.DB) (*[]byte, error) {
	answers, err := model.GetAnswers(idSurvey, db)
	if err != nil {
		return nil, err
	}
	answersJSON, err := json.Marshal(answers)
	if err != nil {
		return nil, err
	}
	return &answersJSON, nil
}

func SetAnswers(newAnswer *answers.NewAnswers, db *sql.DB) error {
	err := model.AnswerSurvey(newAnswer, db)
	if err != nil {
		return err
	}
	return nil
}
