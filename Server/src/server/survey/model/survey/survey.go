package survey

type Survey struct {
	Id       int
	Question string
	Choices  []Choice
}

type Choice struct {
	Id     int
	Choice string
}
