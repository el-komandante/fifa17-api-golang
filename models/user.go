package models

type User struct {
  Model
  Name string `json:"name"`
  Score int `json:"score"`
}

func (user User) GetData() interface{} {
  data := make(map[string]interface{})
  data["name"] = user.Name
  data["id"] = user.ID
  data["score"] = user.Score

  // Get # of wins and losses
  var wins, losses int
  DB.Model(&Game{}).Where("winner_id = ?", user.ID).Count(&wins)
  DB.Model(&Game{}).Where("loser_id = ?", user.ID).Count(&losses)

  data["wins"] = wins
  data["losses"] = losses
  return data
}

func newUser(name string) User {
  user := User{}
  user.Name = name
  user.Score = 1200
  DB.Create(&user)

  return user
}
