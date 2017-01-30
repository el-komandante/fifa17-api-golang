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

	if pair[0] != " " || pair[1] != "fleming" {
		http.Error(w, "Not authorized", 401)
		return
	}

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
