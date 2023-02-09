package todo

import "time"

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	ActivityGroupID int `json:"activity_group_id"`
	IsActive bool `json:"is_active"`
	Priority string `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Payload struct{
	Title string `json:"title"`
	ActivityGroupID int `json:"activity_group_id"`
	IsActive bool `json:"is_active"`
	Priority string `json:"priority"`
}