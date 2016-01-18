package login

import (
	"regexp"
	"PJApp/app/models"
	"github.com/revel/revel"

	"PJApp/app/controllers/common"
)

var userRegex = regexp.MustCompile("^\\w*$")

type Login struct {
	common.PJController
}

func Validate(v *revel.Validation, user models.User) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{18},
		revel.MinSize{2},
		revel.Match{userRegex},
	)

	validatePassword(v, user.Password).Key("user.Password")

}

func validatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{20},
		revel.MinSize{4},
	)
}

