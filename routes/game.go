package routes

import (
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"

  "github.com/cyrusaf/fifa-api-golang/models"
)

func addGameRoutes(r *mux.Router) {
  // Create new game
  r.HandleFunc("/games", createGameHandler).Methods("POST")

  // Get a users games
  r.HandleFunc("/users/{user_id}/games", getUserGamesHandler).Methods("GET")
}

type CreateGameReq struct {
  WinnerId int `json:"winner_id"`
  LoserId int `json:"loser_id"`
  WinnerGoals int `json:"winner_goals"`
  LoserGoals int `json:"loser_goals"`
}
func createGameHandler(w http.ResponseWriter, req *http.Request) {
  // Get JSON request data
  decoder := json.NewDecoder(req.Body)

  // Decode json into user model
  var body CreateGameReq
  err := decoder.Decode(&body)
  if err != nil {
    panic(HttpError{400, err.Error()})
    return
  }

  game := models.NewGame(body.WinnerId, body.LoserId, body.WinnerGoals, body.LoserGoals)

  setData(req, game.GetData())
}

func getUserGamesHandler(w http.ResponseWriter, req *http.Request) {
  user_id, _ := strconv.Atoi(mux.Vars(req)["user_id"])
  var user models.User
  models.FromID(&user, user_id)

  var games []models.Game
  models.DB.Where("winner_id = ?", user_id).Or("loser_id = ?", user_id).Order("date").Find(&games)

  response := make([]interface{}, len(games))
  for i, game := range games {
    response[i] = game.GetData()
  }

  setData(req, response)
}
