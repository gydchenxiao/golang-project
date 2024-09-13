package model

type User struct {
	BaseModel[uint]
	Username string `json:"username"`
}
