package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
	"time"
)

type MyClaims struct {
	jwt.StandardClaims
	Email    string
	Password string
	Role     string
}

func GenerateJWT(email string, pass string, isAdmin string) (string, error) {
	var sampleSecretKey = []byte(os.Getenv("SECRET_KEY"))

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
		Email:    email,
		Password: pass,
		Role:     isAdmin,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, invalid token format",
		})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized, token verification failed",
			})
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, error parsing JWT",
		})
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, invalid token",
		})
		return
	}

	if token.Valid {
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, invalid token",
		})
		return
	}
}

func ExtractClaims(ctx *gin.Context) (string, string, string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized, token verification failed",
			})
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		email := claims["Email"].(string)
		password := claims["Password"].(string)
		role := claims["Role"].(string)
		return email, password, role, nil
	} else {
		err := errors.New("unable to extract claims")
		return "", "", "", err
	}
}

func BackOffice(ctx *gin.Context) {
	_, _, role, err := ExtractClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "forbidden access to API",
		})
		return
	}
	return
}
