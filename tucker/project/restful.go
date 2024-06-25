package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	s := e.Group("/students")
	s.GET("", func(c echo.Context) error {
		jsonStr, _ := json.Marshal(students)
		return c.JSON(http.StatusOK, string(jsonStr))
	})
	s.GET("/:id", func(c echo.Context) error {
		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)
		student, _ := findById(id)
		jsonStr, _ := json.Marshal(student)
		return c.JSON(http.StatusOK, string(jsonStr))
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

var students = Students{
	Student{
		Id:    1,
		Name:  "홍길동",
		Age:   20,
		Score: 80,
	},
	Student{
		Id:    2,
		Name:  "유관순",
		Age:   30,
		Score: 70,
	},
}

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}
type Students []Student

func findById(id int) (Student, error) {
	s := Student{}
	for _, stu := range students {
		if stu.Id == id {
			return stu, nil
		}
	}
	return s, nil
}
