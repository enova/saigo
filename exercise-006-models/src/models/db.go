package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB


const DB_URI = "user=postgres dbname=exercise-006-models sslmode=disable"


func GetConn() *sqlx.DB {
	if Conn == nil{
		var err error
		Conn, err = sqlx.Connect("postgres", DB_URI)
		fmt.Println(err)

	}
	return Conn
}

func Truncate(){
	GetConn().MustExec("TRUNCATE customers, orders, products")
}

