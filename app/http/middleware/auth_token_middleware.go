package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/goravel/framework/contracts/http"
	users_repository "goravel/src/Adapter/Driven/Repositories"
)

func AuthToken() http.Middleware {
	return func(ctx http.Context) {
		if ctx.Request().Header("Authorization") == "" {
			ctx.Request().AbortWithStatusJson(422, map[string]interface{}{"error": "Authorization header not found"})
			return
		}

		authToken, err := jwtValidate(ctx.Request().Header("Authorization"))
		fmt.Println(err)
		if err != nil {
			ctx.Request().AbortWithStatusJson(422, map[string]interface{}{"error": "Invalid token format"})
			return
		}

		_, err = users_repository.FindOne(authToken)

		if err != nil {
			ctx.Request().AbortWithStatusJson(422, map[string]interface{}{"error": "User not found"})
			return
		}

		ctx.Request().Next()
	}
}

func jwtValidate(tokenString string) (string, error) {
	fmt.Println("JWT Validation")

	// Your JWT secret key (this should match the one used to sign the token)
	secretKey := "teste"

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	// Check for parsing errors
	if err != nil {
		return "", fmt.Errorf("Error parsing token: %v", err)
	}

	// Check if the token is valid
	if token.Valid {
		// Token is valid, you can access its claims like this:
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println("Valid token!")
		tokenClaim, ok := claims["token"].(string)
		if !ok {
			return "", fmt.Errorf("Token claim is not a string")
		}
		return tokenClaim, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		// Token is not valid, handle the validation error
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", fmt.Errorf("Invalid token format")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "", fmt.Errorf("Token is either expired or not active yet")
		} else {
			return "", fmt.Errorf("Couldn't handle this token: %v", err)
		}
	} else {
		return "", fmt.Errorf("Couldn't handle this token: %v", err)
	}
}
