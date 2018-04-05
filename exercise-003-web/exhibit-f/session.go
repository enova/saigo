package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	sessionFile = "/tmp/exhibit-f-session"
	sessionLock = sync.RWMutex{}
)

// Session is a key / value pair of entered
// names to the number of times each name was
// entered
var Session = make(map[string]int)

func sessionEntries() []string {
	contents, _ := ioutil.ReadFile(sessionFile)
	return strings.Split(string(contents), "\n")
}

// AddUser increments the count of a username
// in a thread-safe way.
func AddUser(username string) {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	Session[username]++
}

// LoadSession loads an existing session from
// a file located at sessionFile.
func LoadSession() {
	for _, s := range sessionEntries() {
		if len(s) == 0 {
			continue
		}
		entry := strings.Split(s, ` `)
		Session[entry[0]], _ = strconv.Atoi(entry[1])
	}
}

// SaveSession saves Session to the file
// located at sessionFile.
func SaveSession() {
	f, _ := os.Create(sessionFile)
	defer func() {
		f.Sync()
		f.Close()
	}()

	for str, count := range Session {
		outStr := fmt.Sprintf("%s %d\n", str, count)
		f.WriteString(outStr)
	}
}
