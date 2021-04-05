package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	t "github.com/vugu-examples/ajax/types"
	"github.com/vugu/vugu"
)

type Root struct {
	Statements      t.StatementList `vugu:"data"`
	Statement       t.Statement     `vugu:"data"`
	Curr_statement  int             `vugu:"data"`
	Answer          int             `vugu:"data"`
	Show_submission bool            `vugu:"data"`
}

func (c *Root) Init(ctx vugu.InitCtx) {
	//Initializing the web application variables
	c.Show_submission = false
	c.Answer = 0
	c.Curr_statement = 0
	go func() {
		statement_list, err := getStatements()
		if err != nil {
			panic(err)
		}
		ee := ctx.EventEnv()
		ee.Lock()
		c.Statements = statement_list
		ptr, err := c.Statements.GetStatementByIndex(c.Curr_statement)
		c.Statement = *ptr
		if err != nil {
			fmt.Println(err.Error())
		}
		ee.UnlockRender()
	}()
}

func (c *Root) toggleForm(event vugu.DOMEvent) {
	c.Show_submission = !c.Show_submission
}

func (c *Root) New(step int) {
	var err error
	var next_statement *t.Statement
	next_statement, err = c.Statements.GetStatementByIndex(c.Curr_statement + step)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.Answer = 0
	c.Statement = *next_statement
	c.Curr_statement = c.Curr_statement + step
}

func (c *Root) Reveal(answer int) {
	if !c.Answered() {
		c.Answer = answer
		c.Statements.AddToStatement(c.Statement.ID, answer)
		go c.submitToServer(answer)
	}
}

func (c Root) Answered() bool {
	return c.Answer != 0
}

func (c *Root) submitToServer(answer int) {
	fmt.Println("Submitting answer to the server...")
	err := postAnswer(c.Curr_statement, answer)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getStatements() (t.StatementList, error) {
	var err error
	var statement_list t.StatementList
	response, err := http.Get("/api/allStatements")
	if err != nil {
		return statement_list, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&statement_list.List)
	if err != nil {
		return statement_list, err
	}
	return statement_list, err
}

func postAnswer(id int, answer int) error {
	var err error
	data := map[string]int{"ID": id, "Answer": answer}
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	response, err := http.Post("/api/submitAnswer", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	posted_data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Successfully posted %v to the server", posted_data)
	return err
}
