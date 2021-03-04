package main

import (
	"fmt"
    "strconv"

	"go-stress-testing/model"
	"go-stress-testing/server"
)

type Gmeter struct {
}

// NewTodos attempts to create a new Todo list
func NewGmeter() (*Gmeter, error) {
	// Create new Todos instance
	result := &Gmeter{}
	// Return it
	return result, nil
}

func (t *Gmeter) BackWord(word string) string {
	return fmt.Sprintf("hi, %s", word)
}

func (t *Gmeter) Stress(requestUrl string, currency string, testNumber string) model.Response {
	intcur, _ := strconv.Atoi(currency)
    int64cur := uint64(intcur)
	intNumber, _ := strconv.Atoi(testNumber)
    int64Number := uint64(intNumber)
	var response model.Response
	request, _ := model.NewRequest(requestUrl, "", 0, false, "")
	response = server.Dispose(int64cur, int64Number, request)
	fmt.Println("Stress here")
    fmt.Println(response.MaxTimeFloat)
	return response
}
