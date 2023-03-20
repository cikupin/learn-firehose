package payload

import (
	"time"
)

type Identity struct {
	ID        string    `json:"id"  faker:"-"`
	FirstName string    `json:"first_name" faker:"first_name"`
	LastName  string    `json:"last_name" faker:"last_name"`
	CreatedAt time.Time `json:"created_at"  faker:"-"`
}
