package dtos

// ForgetPasswordSendDto is used when sending the OTP for password reset
type ForgetPasswordSendDto struct {
	SystemName string `json:"system_name" required:"required,min=3"`
	Email      string `json:"email" required:"required,email"`
}

// ForgetPasswordVerifyDto is used when verifying the OTP for password reset
type ForgetPasswordVerifyDto struct {
	Email string `json:"email" required:"required,email"`
	Otp   string `json:"otp" required:"required,number,min=6,max=6"`
}

// ForgetPasswordResetDto is used when resetting the password after OTP verification
type ForgetPasswordResetDto struct {
	SystemName string `json:"system_name"`
	Email      string `json:"email" required:"required,email"`
	Otp        string `json:"otp" required:"required,number,min=6,max=6"`
	Password   string `json:"password" required:"required,password,min=3,max=130"`
}
