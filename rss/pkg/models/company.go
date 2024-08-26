package models

import "github.com/uptrace/bun"

type Company struct {
	bun.BaseModel `json:"company" bun:"companies"`

	ID          int64      `json:"id" bun:"id,pk,notnull,autoincrease"`
	Name        string     `json:"name" bun:"name,notnull"`
	Headquaters Location   `json:"headquaters" bun:"headquaters"`
	Offices     []Location `json:"offices" bun:"offices"`
}

type Location struct {
	Latitude  float64 `json:"latitude" bun:"latitude,notnull,double precision"`
	Longitude float64 `json:"longitude" bun:"longitude,notnull,double precision"`
	Timezone  string  `json:"timezone" bun:"timezone"`
}
