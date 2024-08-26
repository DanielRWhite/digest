package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Feed struct {
	bun.BaseModel `json:"feed" bun:"feeds"`

	Site        string        `json:"site" bun:"site,rel:belongs-to,join:site_id"`
	SiteID      int64         `json:"site_id" bun:"site_id,notnull"`
	Name        string        `json:"name" bun:"name,notnull"`
	Link        string        `json:"link" bun:"link,notnull"`
	Enabled     bool          `json:"enabled" bun:"enabled,boolean,notnull,default:true"`
	Cooldown    time.Duration `json:"cooldown" bun:"cooldown,interval,notnull,default:1 minute"`
	LastUpdated time.Time     `json:"last_updated" bun:"last_updated,timestamp"`
	LastCrawled time.Time     `json:"last_crawled" bun:"last_crawled,timestamp"`
	Tags        []string      `json:"tags" bun:"tags"`
}
