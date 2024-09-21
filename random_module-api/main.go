package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Session based request counter
var responseCounter int = 1

// Prepared struct for json output
type RandomNumber struct {
	ID        int    `json:"id"`
	Number    string `json:"number"`
	Timestamp string `json:"timestamp"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Try to open random module device output.
	file, err := os.Open("/dev/randommodule")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read single line from random module output.
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	buffer := scanner.Text()

	file.Close()

	// Request time.
	time_current := time.Now()
	time_string := time_current.Format("2006-01-02 15:04:05")

	// Prepare new slice for json output.
	var randjson RandomNumber
	randjson.ID = responseCounter
	randjson.Number = string(buffer)
	randjson.Timestamp = time_string

	response_byte, _ := json.Marshal(randjson)
	response := string(response_byte)

	// Dump
	log.Println("Random number request (", responseCounter, ") : ", randjson.Number)

	// write json to http responseWriter.
	fmt.Fprint(w, response)

	responseCounter++
}

func main() {

	//
	fmt.Println("\"random_number\" -API-")

	router := httprouter.New()
	router.GET("/number", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
