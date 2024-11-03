package req

type ReqSignUp struct {
	//UserId    int
	FullName string `json:"fullname,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"pwd"`
}
