package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/hocklo/logger"
)

const (
	AUTHOR = "hocklo"
	HTTP_PORT = ":8080"
)

type Message struct {
	Text		string
	Version		string
	ZuluDate	string
	ParsedDate	string
	Author		string
}

/**
 * Take the first url parameter and say Welcome, "value of argument".
 * e.g : Welcome, hocklo!
 */
func handler(w http.ResponseWriter, r *http.Request) {
	logger.Audit("func::handler::start")
	// localhost:8080/{argument} 
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
	logger.Audit("func::handler::end")
}
/**
 * Output a message about this go program.
 */
func about(w http.ResponseWriter, r *http.Request) {
	logger.Audit("func::about::start")
	zuluDate:= time.Now().Format(time.RFC3339)
	parsedDate:= time.Now().Format("02-01-2006")
	// Save the message inside m
	m:= Message{"Welcome to the Hocklo API","0.1,", zuluDate, parsedDate, AUTHOR}
	// Marshal "m" inside "b" if some problem occurs the value are writed inside "err".
	b, err := json.Marshal(m)
	// Check the value of "err" if it isn't "nil" output a panic error.
	errorManagement(err, "func::about::Json marshall of message failed!")
	// Write the value of "b" at ResponseWriter
	w.Write(b)
        logger.Audit("func::about::end")
}

/**
 * Main function to handle all mappings
 */
func main() {
	logger.Info("func::main::GoMicroservice! Started!")
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	err:= http.ListenAndServe(HTTP_PORT, nil)
	// Check the value of "err" if it isn't nil output a panic error
	//errorManagement(err, "func::main::GoMicroservice! BOOM!!")
	//errorManagement(err)
	errorManagement(err, "func::main::GoMicroservice! BOOM!!","Do commit before run", "Your server is down.")
}


/**
 * Error management
 * Params err Error to throw, m Message to log with ERROR lvl.
 */
func errorManagement(err error, messages ...string) {
	// Check the value of "err" if it isn't nil output a panic error
	if err != nil {
		for i := 0; i< len(messages); i++ {
			logger.Error(messages[i])
		}
                panic(err)
        }
}
