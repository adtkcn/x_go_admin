package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func Generate() (id string, b64s string, err error) {

	var driver = base64Captcha.DefaultDriverDigit
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return id, b64s, err
	//create base64 encoding captcha
	// switch param.CaptchaType {
	// case "audio":
	// 	driver = param.DriverAudio
	// case "string":
	// 	driver = param.DriverString.ConvertFonts()
	// case "math":
	// 	driver = param.DriverMath.ConvertFonts()
	// case "chinese":
	// 	driver = param.DriverChinese.ConvertFonts()
	// default:
	// 	driver = param.DriverDigit
	// }

}

// / verifyCode 为前端提交的验证码值
// verifyID 为前端提交的验证码ID
func Verify(verifyID, verifyCode string) bool {
	if verifyID == "" || verifyCode == "" {
		return false
	}
	isCorrect := store.Verify(verifyID, verifyCode, true)
	// if !isCorrect {
	// 	// 验证码不正确
	// 	return false
	// } else {
	// 	// 验证码正确
	// 	return true
	// }
	return isCorrect
}
