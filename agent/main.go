package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
)

type result struct {
	Code    int         `json:"code"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

type command struct {
	Args []string `json:"args"`
}

func main() {
	router := gin.Default()
	router.POST("/run", runCommand)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf("localhost:%s", port))
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
		fmt.Println(err)
		return
	}
	var res = result{
		Code:    200,
		Message: nil,
		Data:    string(output),
	}
	c.IndentedJSON(http.StatusOK, res)
}
