package transformers

import "time"

type UserCreateResponseTransformer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
	Verified  bool      `json:"verified"`
}
