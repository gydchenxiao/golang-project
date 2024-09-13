package model

type Course struct {
	BaseModel[int]
	Tilte string `json:"tilte"`
}
