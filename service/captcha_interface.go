package service

type CaptchaInterface interface {

	// Get 获取验证码
	Get() (map[string]interface{}, error)

	// Check 核对验证码
	Check(token string, pointJson string) error

	// Verification 二次校验验证码(后端)
	Verification(token string, pointJson string) error

	// VerificationByEncryptCode 二次验证，必须要先执行一次验证才能调用二次验证
	VerificationByEncryptCode(encryptCode string) error
}
