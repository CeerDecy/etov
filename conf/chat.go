package conf

import "time"

type Cache struct {
	TTL  time.Duration `json:"ttl"`
	Size int           `json:"size"`
}
