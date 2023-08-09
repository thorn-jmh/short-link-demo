package dto

// >>>>>>>>>>>>>>>>>> UserRegister >>>>>>>>>>>>>>>>>>

type UserRegisterReq struct {
	Email    string `json:"email"`    // 用户邮箱
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 用户密码
}

// >>>>>>>>>>>>>>>>>> UserLogin >>>>>>>>>>>>>>>>>>

type UserLoginReq struct {
	CaptchaID    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
	Email        string `json:"email"`    // 用户邮箱
	Password     string `json:"password"` // 用户密码
}

// >>>>>>>>>>>>>>>>>> Captcha >>>>>>>>>>>>>>>>>>

type UserCaptchaResp struct {
	CaptchaID  string `json:"captcha_id"`  // 验证码 ID
	CaptchaURL string `json:"captcha_url"` // 验证码图片地址
}

// >>>>>>>>>>>>>>>>>> GetInfo >>>>>>>>>>>>>>>>>>

type GetUserInfoResp struct {
	Email string `json:"email"` // 用户邮箱
	ID    uint   `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
}

// >>>>>>>>>>>>>>>>>> UpdateInfo >>>>>>>>>>>>>>>>>>

type UpdateUserInfoReq struct {
	Email string `json:"email"` // 用户邮箱
	ID    uint   `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
}

// >>>>>>>>>>>>>>>>>> Password >>>>>>>>>>>>>>>>>>

type UserPasswdReq struct {
	NewPwd string `json:"new_pwd"` // 新密码
	OldPwd string `json:"old_pwd"` // 旧密码
}
