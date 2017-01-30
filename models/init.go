package models

import (
  "strconv"
  "time"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type HttpError struct {
  Status_code int
  Error       string
}

type Model struct {
  ID uint `gorm:"primary_key" json:"id"`
}



var DB *gorm.DB

func init() {
  var err interface{}
  time.Sleep(time.Duration(10) * time.Second)
  DB, err = gorm.Open("postgres", "postgres://postgres:postgres@postgres/postgres?sslmode=disable")
  if err != nil {
    panic(err)
  }
}


func FromID(model interface{}, id int) {
  DB.First(model, id)

  // Check that model exists
  if DB.NewRecord(model){
    panic(HttpError{404, "Cannot find model with ID=" + strconv.Itoa(id)})
  }
}
