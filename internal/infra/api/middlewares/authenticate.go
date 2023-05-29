package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		//STEP 1 get JWT from headers
		token := getJwt(context)

		//STEP 2 PARSE Token
		jwtToken, err := parseToken(token)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
		}

		//SETEP 3 VALID TOKEN
		if !isValidToken(jwtToken) {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		//STEP 4 get clains
		setClains(context, jwtToken)

	}
}

func getJwt(context *gin.Context) string {
	authorizationHeader := context.Request.Header.Get("Authorization")
	jwt := strings.TrimPrefix(authorizationHeader, "Bearer ")
	return jwt
}

func parseToken(token string) (*jwt.Token, error) {

	jwtToken, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		return []byte("CLAVESUPERSEGURA"), nil
	})
	if err != nil {
		return nil, errors.New("failed unmarshall jwt")
	}
	return jwtToken, nil

}

func isValidToken(jwtToken *jwt.Token) bool {
	return jwtToken.Valid
}

func setClains(context *gin.Context, jwtToken *jwt.Token) {
	//vamos hace cast de un intreface
	clains := jwtToken.Claims.(jwt.MapClaims)
	context.Set("userID", clains["userID"]) //podemosponer solo estso  otodos
	context.Set("user", clains["user"])
}
