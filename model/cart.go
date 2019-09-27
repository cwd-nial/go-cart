package model

import "time"

type Cart struct {
	Id         string    `json:"id"`
	CustomerId string    `json:"customer_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
