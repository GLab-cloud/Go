package req

type ReqSignUp struct {
	//UserId    int
	FullName string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
