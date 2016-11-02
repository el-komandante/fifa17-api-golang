package models

import (
  "os"
  "strconv"

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
  DB, err = gorm.Open("postgres", "user=" + os.Getenv("PG_USER") + " dbname=fifa17 password=" + os.Getenv("PG_PASSWORD") + " sslmode=disable")
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
