package main

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"log"
)

const (
	AUTHOR = "hocklo"
	INFO = "INFO"
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
	logInfo("func::handler::start::"+time.Now().Format(time.RFC3339Nano))
	// localhost:8080/{argument} 
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
	logInfo("func::handler::end::"+time.Now().Format(time.RFC3339Nano))
}
/**
 * Output a message about this go program.
 */
func about(w http.ResponseWriter, r *http.Request) {
	logInfo("func::about::start::"+time.Now().Format(time.RFC3339Nano))
	zuluDate:= time.Now().Format(time.RFC3339)
	parsedDate:= time.Now().Format("02-01-2006")
	// Save the message inside m
	m:= Message{"Welcome to the Hocklo API","0.1,", zuluDate, parsedDate, AUTHOR}
	// Marshal "m" inside "b" if some problem occurs the value are writed inside "err".
	b, err := json.Marshal(m)
	// Check the value of "err" if it isn't "nil" output a panic error.
	if err != nil {
		panic(err)
	}
	// Write the value of "b" at ResponseWriter
	w.Write(b)
        logInfo("func::about::end::"+time.Now().Format(time.RFC3339Nano))
}

/**
 * Main function to handle all mappings
 */
func main() {
	logInfo("GoMicroservice! Started!")
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}


/**
 * Print a log INFO.
 */
func logInfo(s string) {
	writeLog(INFO, s)
}

/**
 * Print &buf to log.
 */
func writeLog(l string, s string) {
        var buf bytes.Buffer // Instance buffer
        var logger = log.New(&buf, "logger: ", log.Lshortfile) // Instance log 
        logger.Print(time.Now().Format(time.RFC3339)+": "+l+": "+s) // Print log to buffer
        fmt.Print(&buf) // throw log from buffer to file.
}

