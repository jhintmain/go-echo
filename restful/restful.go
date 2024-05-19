package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = map[string]*User{}

// 유저 생성
// curl -X POST -H "Content-Type: application/json" -d '{"id":"comoozi","name":"jiheapark","email":"jhintmain@gmail.com"}' http://localhost:8080/users
// curl -X POST -H "Content-Type: application/json" -d '{"id":"leemoozi","name":"leemjihyen","email":"leeintmain@gmail.com"}' http://localhost:8080/users
func createUser(c echo.Context) error {
	user := &User{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	users[user.Id] = user
	return c.JSON(http.StatusCreated, user)

}

// 특정 유저 정보 가져오기
func getUser(c echo.Context) error {
	id := c.Param("id")
	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)

}

// 유저 정보 갱신
func updateUser(c echo.Context) error {
	updateUser := &User{}
	if err := c.Bind(updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id := c.Param("id")
	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	users[user.Id] = updateUser
	return c.JSON(http.StatusOK, updateUser)
}

// curl -X DELETE http://localhost:8080/users/leemoozi
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	if _, exists := users[id]; !exists {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	delete(users, id)
	return c.JSON(http.StatusOK, "user deleted")
}

// http://localhost:8080/users
func listUsers(c echo.Context) error {
	userList := []*User{}

	for _, user := range users {
		userList = append(userList, user)
	}

	return c.JSON(http.StatusOK, userList)

}

func main() {
	e := echo.New()

	// 미들웨어 설정
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.GET("/users", listUsers)

	e.Start(":8080")
}
