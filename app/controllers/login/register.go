package login

import (
	"PJApp/app/models"
	"PJApp/app/services"
	"PJApp/app/routes"
	"PJApp/app/validation/login"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"unsafe"
	"PJApp/app/controllers/common"
)

/* ユーザ登録画面.登録ボタン押下処理 */
func (c Login) SaveUser(user models.User, verifyPassword string) revel.Result {
	common.WriteLog("Register", "pushLoginButton", "Start")

	// バリデーションチェック（jsがオフになっているブラウザ対応）
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).
	Message("Password does not match")
	login.Validate(c.Validation, user)

	if c.Validation.HasErrors() {
		common.WriteLog("Register", "pushLoginButton", "Validate Error")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Login.Register())
	}

	// ユーザ名の重複確認
	if userData := services.GetUser(user.Username); userData != nil {
		common.WriteLogStr("Register", "pushLoginButton", "Input Regestered Username", "userName = " + user.Username)
		c.Flash.Error("UserName Already Registred.")
		return c.Redirect(routes.Login.Register())
	}

	// パスワードのハッシュ化
	hassedPass , _ := bcrypt.GenerateFromPassword(
		*(*[]byte)(unsafe.Pointer(&user.Password)), bcrypt.DefaultCost)
	user.Password = string(hassedPass)

	// 登録処理
	if err := services.InsertUser(user); err != nil {
		common.WriteLogStr("Register", "pushLoginButton", "Insert Error",
			"userName = " + user.Username + " password = " + user.Password)

		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Login.Register())
	}

	c.Session["user"] = user.Username
	c.Flash.Success("登録が完了しました。")
	c.Flash.Out["username"] = user.Username

	common.WriteLog("Register", "pushLoginButton", "End")

	return c.Redirect(routes.Login.Index())
}
