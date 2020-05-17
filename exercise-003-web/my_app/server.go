package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// View struct contains information submitted on the form
type View struct {
	NameFrequencies map[string]int
}

var (
	homeT           *template.Template
	nameFrequencies map[string]int
	record          []string
)

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/my_app/home.html"))
	nameFrequencies = make(map[string]int)
	readCsvFile(dir + "/my_app/usernames.csv")
}

// Create a listener that saves names and frequencies upon receiving an interrupt from the OS.
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		writeCsvFile("./my_app/usernames.csv")
		os.Exit(0)
	}()
}

func readCsvFile(filePath string) {
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	record := make([]string, 0)
	for {
		record, err = r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		nameFrequencies[strings.ToLower(record[0])], _ = strconv.Atoi(record[1])
	}
}

func writeCsvFile(fname string) {
	// empty file contents
	err := os.Truncate(fname, 0)

	if err != nil {
		log.Fatal(err)
	}

	// open the file
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)

	for name, frequency := range nameFrequencies {
		w.Write([]string{name, strconv.Itoa(frequency)})
		w.Flush()
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")

		nameFrequencies[strings.ToLower(username)]++
	}
	v := View{NameFrequencies: nameFrequencies}
	homeT.Execute(w, &v)
}

func main() {
	setup(".")
	setupCloseHandler()
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
