package dto

import "time"

type CreateSubscription struct {
	UserId      string `json:"user_id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
}

type Subscription struct {
	Id          int64     `json:"id"`
	UserId      string    `json:"user_id"`
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type UpdateSubscription struct {
	UserId      string `json:"user_id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
