package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	t "github.com/vugu-examples/ajax/types"
	"github.com/vugu/vugu"
)

type Root struct {
	Statements      t.StatementList `vugu:"data"`
	curr_statement  int             `vugu:"data"`
	Show_submission bool            `vugu:"data"`
}

func (c *Root) Init(ctx vugu.InitCtx) {
	//Initializing the web application variables
	c.Show_submission = false
	c.curr_statement = 0
	go func() {
		response, err := http.Get("/api/allStatements")
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		var statement_list t.StatementList
		err = json.NewDecoder(response.Body).Decode(&statement_list.List)
		if err != nil {
			panic(err)
		}
		ee := ctx.EventEnv()
		ee.Lock()
		c.Statements = statement_list
		ee.UnlockRender()
	}()
}

func (c *Root) APITest(event vugu.DOMEvent) {
	fmt.Println("Sending Request...")
}

func (c *Root) toggleForm(event vugu.DOMEvent) {
	c.Show_submission = !c.Show_submission
}

/* func (c *Root) submitAnswer(event vugu.DOMEvent, i int) {
	statement, err := c.Statements.getStatementByIndex(c.curr_statement)
	if err != nil {
		fmt.Println(err)
	}
	url := fmt.Sprintf("/api/submit/%d/%d", c.curr_statement, i)
	go func(){
		http.Post(url, "application/json", json.NewEncoder())
	}

	http.Post(url, i))
}
*/
