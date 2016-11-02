package main

import (
  "github.com/cyrusaf/fifa17-api-golang/routes"
  _ "net/http/pprof"
  "net/http"
  "log"
)

func main() {
  go func() {
    log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
  }()
  n  := routes.GetRouter()
  n.Run(":3002")
}
