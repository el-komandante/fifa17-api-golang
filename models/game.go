package models

import (
  "time"
  "math"
  "fmt"
)

type Game struct {
  Model
  Winner User `gorm:"ForeignKey:WinnerId" json:"winner"`
  WinnerId uint `json:"-"`
  Loser User `gorm:"ForeignKey:LoserId" json:"loser"`
  LoserId uint `json:"-"`

  WinnerScore int `json:"loser_score"`
  LoserScore int `json:"winner_score"`

  WinnerGoals int `json:"winner_goals"`
  LoserGoals int `json:"loser_goals"`

  Date int `json:"date"`
}

func (game Game) GetData() interface{} {
  var winner User
  var loser User
  DB.Model(&game).Association("Winner").Find(&winner)
  DB.Model(&game).Association("Loser").Find(&loser)
  game.Winner = winner
  game.Loser = loser
  return game
}

func calcEloDiff(winner_rating, loser_rating, diff int, draw bool) int {
  s := 1.0
  if draw {s = 0.5}
  k := 100.0
  e := 1/(1 + math.Pow(10, float64(loser_rating - winner_rating)/500.0))
  fmt.Println(e)
  fmt.Println(k * (s - e))
  return int(k * (s - e))
}

func NewGame(winner_id, loser_id, winner_goals, loser_goals int) Game {
  // Get winner and loser
  var winner User
  var loser User
  FromID(&winner, winner_id)
  FromID(&loser, loser_id)

  draw := false
  if winner_goals == loser_goals {
    draw = true
  } else if winner_goals < loser_goals {
    panic(HttpError{400, "Winner goals must be greater than or equal to loser goals."})
  }

  elo_diff := calcEloDiff(winner.Score, loser.Score, winner_goals-loser_goals, draw)
  winner.Score += elo_diff
  loser.Score -= elo_diff

  game := Game{}
  game.Date = int(time.Now().Unix())
  game.WinnerId = uint(winner_id)
  game.LoserId = uint(loser_id)
  game.WinnerScore = winner.Score
  game.LoserScore = loser.Score
  game.WinnerGoals = winner_goals
  game.LoserGoals = loser_goals


  DB.Create(&game)
  DB.Save(&winner)
  DB.Save(&loser)

  return game
}
