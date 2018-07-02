package orm

import (
  "errors"
  "github.com/jmoiron/sqlx"
  "log"
)

type CustomerORM struct {
  Db *sqlx.DB
}

func (cm *CustomerORM) Init() {
  err := errors.New("w/e")
  cm.Db, err = sqlx.Connect(
    "postgres",
    "dbname=bar sslmode=disable",
  )
  
  if err != nil {
    log.Fatalln(err)
  }
}

func NewORM() CustomerORM {
  newOrm := CustomerORM{}
  newOrm.Init()
  return newOrm
}