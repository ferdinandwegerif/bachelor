package types

import (
	"errors"
	"fmt"
)

type StatementList struct {
	List []Statement
}

func NewStatementList() *StatementList {
	return &StatementList{}
}

func (s *StatementList) AddStatement(statement Statement) error {
	var err error
	if statement.isValid() {
		statement.ID = s.getNextID()
		s.List = append(s.List, statement)
	} else {
		err = errors.New("failed to add new statement")
	}
	return err
}

//Returns a statement with the given index
func (s *StatementList) GetStatementByIndex(i int) (*Statement, error) {
	var err error
	if s.IsIndexValid(i) {
		return &s.List[i], err
	} else {
		err = errors.New("index is not valid")
		return nil, err
	}
}

func (s *StatementList) GetStatementByID(id int) (*Statement, error) {
	var err error
	for i, statement := range s.List {
		if statement.ID == id {
			return &s.List[i], err
		}
	}
	err = errors.New("couldn't find the statement with the given ID")
	return nil, err
}

//Checks if statement at index can be retrieved
func (s StatementList) IsIndexValid(i int) bool {
	if i >= 0 && i < len(s.List) {
		return true
	}
	return false
}

func (s *StatementList) AddExampleData() {
	example_data := []Statement{
		{Description: "Would you rather go into the past and meet your ancestors or go into the future and meet your great-great grandchildren?", Answer_1: "Go into the past.", Answer_2: "Go into the future."},
		{Description: "Would you rather have more time or more money?", Answer_1: "More time.", Answer_2: "More money."},
		{Description: "Would you rather have a rewind button or a pause button on your life?", Answer_1: "Rewind button.", Answer_2: "Pause button."},
	}
	for _, statement := range example_data {
		err := s.AddStatement(statement)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *StatementList) AddToStatement(id int, answer int) error {
	statement, err := s.GetStatementByID(id)
	if err != nil {
		return err
	}
	err = statement.AddAnswer(answer)
	return err
}

func (s StatementList) getNextID() int {
	currentID := len(s.List)
	return currentID + 1
}
