package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	OpenPeopleConnection()
	defer ClosePeopleConnection()

	People.DeleteAll()
	People.Insert("Bruce Leroy", 111223333)
	People.Insert("John Doe", 912049182)

	susan, _ := People.Insert("Susan Litner", 119282461)
	People.Delete(susan)

	hale, _ := People.Insert("Hale Bernt", 444556666)
	error := People.Update(hale, 192837981273918273)
	if error != nil {
		fmt.Println("There was an error:", error)
	}

	for _, person := range People.All() {
		fmt.Printf("Person %s", person.ToString())
	}
}
