package models

import (
	"github.com/DanielRWhite/digest/rss/pkg/types"
	"github.com/uptrace/bun"
)

type Site struct {
	bun.BaseModel `json:"site" bun:"sites"`

	ID        int64            `json:"id" bun:"id,pk,autoincrement"`
	Name      string           `json:"name" bun:"name,notnull"`
	Website   string           `json:"website" bun:"website,notnull"`
	Languages []types.Language `json:"languages" bun:"languages,notnull"`
	Office    Location         `json:"location" bun:"location"`
}
