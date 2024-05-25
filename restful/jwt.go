package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	jwtMiddleware "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWTCustomClaims defines the custom claims for the JWT token
type JWTCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// Handler for the unprotected route
func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Handler for the login route
func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Check username and password
	if username == "user" && password == "password" {
		// Set custom claims
		claims := &JWTCustomClaims{
			Name:  "John Doe",
			Admin: true,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)),
			},
		}

		// Create token

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		log.Println("SigningMethodHS256 : ", *jwt.SigningMethodHS256)

		// Generate encoded token and send it as response
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// Handler for the protected route
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Unauthenticated route
	e.GET("/", accessible)

	// Login route
	e.POST("/login", login)

	// Restricted group

	r := e.Group("/restricted")
	r.Use(jwtMiddleware.WithConfig(jwtMiddleware.Config{
		SigningKey: []byte("secret"),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
	}))
	r.GET("", restricted)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
