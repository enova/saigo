package orm

import (
  "errors"
  "github.com/jmoiron/sqlx"
  "log"
)
type GenericORM struct {
  Db *sqlx.DB
}

type CustomerORM struct {
  GenericORM
}

type OrderORM struct {
  GenericORM
}

func (cm *GenericORM) Init() {
  err := errors.New("w/e")
  cm.Db, err = sqlx.Connect(
    "postgres",
    "dbname=bar sslmode=disable",
  )
  
  if err != nil {
    log.Fatalln(err)
  }
}

func ForCustomer() CustomerORM {
  newOrm := CustomerORM{}
  newOrm.Init()
  return newOrm
}

func ForOrder() OrderORM {
  newOrm := OrderORM{}
  newOrm.Init()
  return newOrm
}