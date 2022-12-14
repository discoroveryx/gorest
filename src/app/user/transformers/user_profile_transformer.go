package transformers

import "time"

type UserCProfileTransformer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
}
