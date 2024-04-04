package lib

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strings"
	"time"
)

type JWTClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type JWTMachineClaims struct {
	MachineID uuid.UUID `json:"machine_id"`
	jwt.RegisteredClaims
}

func GenerateMachineJWT(machineID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := JWTMachineClaims{
		MachineID: machineID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyMachineJWT(tokenString string) (*JWTMachineClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTMachineClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTMachineClaims); ok && token.Valid {
		if time.Now().After(claims.ExpiresAt.Time) {
			return nil, errors.New("token is expired")
		}

		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func GenerateJWT(username string, role string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := JWTClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		if time.Now().After(claims.ExpiresAt.Time) {
			return nil, errors.New("token is expired")
		}

		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func JWTAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if strings.HasPrefix(tokenString, bearerPrefix) {
			tokenString = strings.TrimPrefix(tokenString, bearerPrefix)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		claims, err := VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func JWTMachineAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if strings.HasPrefix(tokenString, bearerPrefix) {
			tokenString = strings.TrimPrefix(tokenString, bearerPrefix)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		claims, err := VerifyMachineJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		c.Set("machine_id", claims.MachineID)
		c.Next()
	}
}
