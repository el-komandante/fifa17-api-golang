package routes

import (
  "encoding/json"
  "net/http"
  "strconv"
  "strings"
  "encoding/base64"

  "github.com/gorilla/mux"

  "github.com/el-komandante/fifa17-api-golang/models"
)

func addUserRoutes(r *mux.Router) {
  // Get all songs
  r.HandleFunc("/users/{user_id}", getUserByIDHandler).Methods("GET")

  // Create new user
  r.HandleFunc("/users", createUserHandler).Methods("POST")

  // Get all users
  r.HandleFunc("/users", getAllUsersHandler).Methods("GET")
}

func getUserByIDHandler(w http.ResponseWriter, req *http.Request) {
  user_id, _ := strconv.Atoi(mux.Vars(req)["user_id"])

  var user models.User
  // Get songs
  models.FromID(&user, user_id)

  // Create response
  setData(req, user.GetData())
}

func createUserHandler(w http.ResponseWriter, req *http.Request) {
  //Basic auth
  w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

  s := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
  if len(s) != 2 {
    http.Error(w, "Not authorized", 401)
    return
  }

  b, decodeErr := base64.StdEncoding.DecodeString(s[1])
  if decodeErr != nil {
    http.Error(w, decodeErr.Error(), 401)
    return
  }

  pair := strings.SplitN(string(b), ":", 2)
  if len(pair) != 2 {
    http.Error(w, "Not authorized", 401)
    return
  }

  if pair[0] != "the_supervisor" || pair[1] != "el_secreto" {
    http.Error(w, "Not authorized", 401)
    return
  }

  // Get JSON request data
  decoder := json.NewDecoder(req.Body)

  // Decode json into user model
  var user models.User
  err := decoder.Decode(&user)
  if err != nil {
    panic(HttpError{400, err.Error()})
    return
  }


  // Save User
  user.Score = 1200
  models.DB.Create(&user)

  setData(req, user.GetData())
}

func getAllUsersHandler(w http.ResponseWriter, req *http.Request) {
  var users []models.User
  models.DB.Order("score desc").Find(&users)

  data := make([]interface{}, len(users))
  for i, user := range users {
    data[i] = user.GetData()
  }

  setData(req, data)
}
