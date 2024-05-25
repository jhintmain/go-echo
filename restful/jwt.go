// package main

// import (
// 	"net/http"
// 	"time"

// 	jwt "github.com/golang-jwt/jwt/v5"
// 	jwtMiddleware "github.com/labstack/echo-jwt/v4"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// // JWTCustomClaims defines the custom claims for the JWT token
// type JWTCustomClaims struct {
// 	Name  string `json:"name"`
// 	Admin bool   `json:"admin"`
// 	jwt.RegisteredClaims
// }

// // Handler for the unprotected route
// func accessible(c echo.Context) error {
// 	c.String(http.StatusOK, "Accessible")
// 	return c.String(http.StatusOK, "Accessible")
// }

// // Handler for the login route
// /*
// curl --location --request POST 'http://localhost:8080/login' \
// --header 'Content-Type: application/x-www-form-urlencoded' \
// --data-urlencode 'username=user' \
// --data-urlencode 'password=password'
// */

// /*
// TOKEN=?
// curl --location --request GET 'http://localhost:8080/restricted' --header "Authorization: Bearer $TOKEN"
// */
// func login(c echo.Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	// Check username and password
// 	if username == "user" && password == "password" {
// 		// Set custom claims
// 		claims := &JWTCustomClaims{
// 			Name:  "John Doe",
// 			Admin: true,
// 			RegisteredClaims: jwt.RegisteredClaims{
// 				Subject:   "login",
// 				IssuedAt:  jwt.NewNumericDate(time.Now()),
// 				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
// 			},
// 		}

// 		// Create token
// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 		// Generate encoded token and send it as response
// 		t, err := token.SignedString([]byte("secret"))
// 		if err != nil {
// 			return err
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{
// 			"token": t,
// 		})
// 	}

// 	return echo.ErrUnauthorized
// }

// // Handler for the protected route
// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*JWTCustomClaims)
// 	name := claims.Name
// 	expiresAt := claims.ExpiresAt.Format(time.DateTime)
// 	// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "Welcome " + name + "!",
// 		"user":    user,
// 		"claims":  claims,
// 		"expires": expiresAt,
// 	})
// }

// func jwtValid(c echo.Context) error {

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "jwtValid",
// 	})
// }
// func main() {
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// Unauthenticated route
// 	e.GET("/", accessible)

// 	// Login route
// 	e.POST("/login", login)

// 	// Restricted group

// 	r := e.Group("/restricted")
// 	r.Use(jwtMiddleware.WithConfig(jwtMiddleware.Config{
// 		SigningKey: []byte("secret"),
// 		NewClaimsFunc: func(c echo.Context) jwt.Claims {
// 			return new(JWTCustomClaims)
// 		},
// 	}))
// 	r.GET("", restricted)

// 	// Start server
// 	e.Logger.Fatal(e.Start(":8080"))
// }
