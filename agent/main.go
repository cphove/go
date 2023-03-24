package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"example/agent/parser"
	"github.com/gin-gonic/gin"
)

type result struct {
	Code    int         `json:"code"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

type command struct {
	Args     []string `json:"args"`
	Feedback bool     `json:"feedback"`
	Filepath *string  `json:"filepath"`
}

func main() {
	router := gin.Default()
	router.POST("/run", runCommand)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func runCommand(c *gin.Context) {
	var newCommand command
	if err := c.BindJSON(&newCommand); err != nil {
		return
	}
	arg1, args := newCommand.Args[0], newCommand.Args[1:]
	cmd := exec.Command(arg1, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	if !newCommand.Feedback {
		var res = result{
			Code:    200,
			Message: nil,
			Data:    string(output),
		}
		c.IndentedJSON(http.StatusOK, res)
		return
	}
	testsuites := parser.Parse(*newCommand.Filepath)
	fmt.Println(testsuites)
	var res = result{
		Code:    200,
		Message: nil,
		Data:    testsuites,
	}
	c.IndentedJSON(http.StatusOK, res)
}
