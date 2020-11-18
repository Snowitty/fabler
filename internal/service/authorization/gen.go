package authorization

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/snowitty/fabler/conf"
	"github.com/snowitty/fabler/internal/model"
	"net/url"
	"strings"
	"time"
)

func Gen(user model.User) (tokenString string, err error){
	SecretString := conf.Config().Secretkey

	secret := []byte(SecretString)

	claims := Claims{
		user.ID,
		user.Type,
		user.ProfileID,
		compatibleJSEncodeURIComponent(url.QueryEscape(user.Profile.Name)),
		 compatibleJSEncodeURIComponent(url.QueryEscape(user.Profile.Desc)),
		jwt.StandardClaims{
			ExpiresAt: time.Time.AddDate(time.Now(), 0,0,7).Unix(),
			Issuer: "localhost",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = token.SignedString(secret)

	return
}

func compatibleJSEncodeURIComponent(str string) (resultStr string){
	resultStr = str
	resultStr = strings.Replace(resultStr, "+", "%20", -1)
	resultStr = strings.Replace(resultStr, "%21", "!", -1)
	resultStr = strings.Replace(resultStr, "%27", "'", -1)
	resultStr = strings.Replace(resultStr, "%28", "(", -1)
	resultStr = strings.Replace(resultStr, "%29", ")", -1)
	resultStr = strings.Replace(resultStr, "%2A", "*", -1)
	return
}
