package transformers

type UserVerifyTransformer struct {
	UserId           uint   `json:"user_id" binding:"required"`
	VerificationCode string `json:"verification_code" binding:"required"`
}
