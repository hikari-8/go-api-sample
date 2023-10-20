package response

import (
	"time"

	"github.com/mahiro72/go-api-sample/domain/model"
)

type user struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(m model.User) user {
	return user{
		ID:        m.ID.String(),
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}