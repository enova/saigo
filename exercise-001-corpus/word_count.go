package main

import (
	"fmt"
	c "github.com/yorker78/saigo/exercise-001-corpus/corpus"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Bad file opening")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Bad file stats")
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Bad file read")
		return
	}

	str := string(bs)
	c.Analyze(str)
}
