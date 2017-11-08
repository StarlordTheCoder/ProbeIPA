package answers

type GivenAnswer struct {
	IdChoice int
	Choice   string
	Amount   int
}

type GivenAnswers struct {
	GivenAnswers []GivenAnswer
}

type NewAnswer struct {
	IDChoice int
}

type NewAnswers struct {
	NewAnswers []NewAnswer
}
