package main

import (
	"html/template"
	"net/http"
	"fmt"
	"strings"
	"os"
	//"bufio"
)

var homeT = template.Must(template.ParseFiles("webserver/web.html"))
var names = make(map[string]int)  // A map of names and their frequencies
//var firstTime int = 0


// home starts the server
func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)

	// Building the filepath
	// cwd, _ := os.Getwd()
	// filepath := cwd + "/webserver/names.txt"

	// Transfer file info to the web page table -- Optional Piece
	// if firstTime == 0 {
	// 	file, err := os.Open(filepath)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer file.Close()
	// 	scanner := bufio.NewScanner(file)
	// 	scanner.Split(bufio.ScanWords)
	// 	for scanner.Scan() {
	// 		names[strings.ToLower(scanner.Text())] += 1
	// 	}
	// 	firstTime = 1
	// }

	// Grab the user input
	text := r.PostFormValue("textbox")
	text = strings.ToLower(text)
	fmt.Println("Hi there ", text)

	if text == "" {
		return
	}

	// Add to the map
	names[text] += 1

	// Print message to screen
	for k, v := range names {
		msg := k + " " + fmt.Sprintf("%d<br>", v)
		fmt.Fprintf(w, msg)
	}

	// Append to the file
	// f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// f.WriteString(text)
	// f.WriteString(" ")
	// f.WriteString(fmt.Sprintf("%d\n", names[text]))
	// f.Close()
}


// Main function 
func main() {
	// Start server
	http.HandleFunc("/", home)

	fmt.Println("Listening..")
	http.ListenAndServe(":8080", nil)
}
