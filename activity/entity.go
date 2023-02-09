package activity

import "time"

type Activity struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Payload struct{
	Title string `json:"title"`
	Email string `json:"email"`
}