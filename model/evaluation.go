package model

import (
	"time"
)

type Option struct {
	Id int64     `bson:"_id"`
	Title string `bson:"title"`
	Value int64  `bson:"value"`
}

type Question struct {
	Id int64         `bson:"_id"`
	Title string     `bson:"title"`
	Options []Option `bson:"options"`
}

type Evaluation struct {
	Id string            `bson:"_id"`
	Name string          `bson:"name"`
	Description string   `bson:"description"`
	Questions []Question `bson:"questions"`
	CreatedAt time.Time  `bson:"created_at"`
	CreatedBy string     `bson:"created_by"`
}

type OptionJSON struct {
	Id int64     `json:"id"`
	Title string `json:"title"`
	Value int64  `json:"value"`
}

type QuestionJSON struct {
	Id int64         `json:"id"`
	Title string     `json:"title"`
	Options []OptionJSON `json:"options"`
}

type EvaluationJSON struct {
	Id string            `json:"id"`
	Name string          `json:"name"`
	Description string   `json:"description"`
	Questions []QuestionJSON `json:"questions"`
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
}

type CreateEvaluationRequestJSON struct {
	Name string          `json:"name"`
	Description string   `json:"description"`
	Questions []Question `json:"questions"`
}
