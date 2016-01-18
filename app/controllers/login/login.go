package login

import (
	"PJApp/app/models"
	"PJApp/app/controllers/common"
	"PJApp/app/services"
	"PJApp/app/routes"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"unsafe"
)

type Login struct {
	common.PJController
}

/* ルートパス遷移処理 */
func (c Login) Index() revel.Result {
	if c.Connected() != nil {
		//		return c.Redirect(routes.PJ.Index())
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}

/* ログイン情報がsessionに保持されているか確認 */
func (c Login) Connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return services.GetUser(username)
	}
	return nil
}

/* ログイン画面.ログインボタン押下処理 */
func (c Login)Login(username, password string, remember bool) revel.Result {
	common.WriteLog("Login", "pushLoginButton", "Start")

	// ユーザ情報取得
	userData := services.GetUser(username)

	// ユーザが見つからない場合
	if userData == nil {
		common.WriteLogStr("Login", "Login", "Not Found Username", "input userName = " + username)
		c.Flash.Out["username"] = username
		c.Flash.Error("Login failed")
		return c.Redirect(routes.Login.Index())
	}

	// パスワードチェック
	err := bcrypt.CompareHashAndPassword(*(*[]byte)(unsafe.Pointer(&userData.Password)), []byte(password))

	// パスワードが間違っている場合
	if err != nil {
		common.WriteLogStr("Login", "Login", "Not Correct Password", "input password = " + password)
		c.Flash.Out["username"] = username
		c.Flash.Error("Login failed")
		return c.Redirect(routes.Login.Index())
	}

	c.Session["user"] = username

	if remember {
		c.Session.SetDefaultExpiration()
	} else {
		c.Session.SetNoExpiration()
	}

	common.WriteLogStr("Login", "Login", "End", "Login by " + userData.Username)

	return c.Redirect(routes.PJ.Index())
}

/* ログイン画面.新規登録ボタン押下処理 */
func (c Login) Register() revel.Result {
	common.WriteLog("Login", "Register", "")
	return c.Render()
}