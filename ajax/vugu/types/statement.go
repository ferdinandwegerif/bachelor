package types

import "errors"

type Statement struct {
	ID          int
	Description string
	Answer_1    string
	Answer_2    string
	/* The amount of votes on this statement/question */
	cntAns1 int
	cntAns2 int
}

func NewStatement(id int, desc string, ans1 string, ans2 string) *Statement {
	return &Statement{
		ID:          id,
		Description: desc,
		Answer_1:    ans1,
		Answer_2:    ans2,
		cntAns1:     0,
		cntAns2:     0,
	}
}

//Checks if a statement has all required fields.
func (s Statement) isValid() bool {
	if (s.Description == "") || (s.Answer_1 == "") || (s.Answer_2 == "") {
		return false
	}
	return true
}

func (s *Statement) AddAnswer(nr int) error {
	var err error
	switch nr {
	case 1:
		s.cntAns1++
	case 2:
		s.cntAns2++
	default:
		err = errors.New("not valid data, only 2 available questions")
	}
	return err
}
