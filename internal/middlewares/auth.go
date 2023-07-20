package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

// these middlewares will be responsible for checking token and if the token is valid

func Authorize() gin.HandlerFunc {

	return func(context *gin.Context) {

		// We will be checking for the token in here and if the token is valid then only we will allow further procrssing

		// Extracting the token from header
		tokenString := context.GetHeader("Authorization")
		if len(tokenString) > 0 {
			// We will verify it
			// We have to validate the token as well that whether its valid or not

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {

				context.JSON(http.StatusUnauthorized, gin.H{
					"message": "Invalid Token",
					"err":     err.Error(),
				})

				context.Abort()
				return

			}

			context.Next() // It will pass the control  from middleware to the controller
			fmt.Println(token)

		} else {
			// We can send the response accordingly

			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide a valid token",
			})

			context.Abort()

		}

	}

}
