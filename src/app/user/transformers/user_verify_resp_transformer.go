package transformers

type UserVerifyRespTransformer struct {
	UserId   uint `json:"user_id" binding:"required"`
	Verified bool `json:"verified" binding:"required"`
}
