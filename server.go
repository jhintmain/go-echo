package main

// import (
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	e.GET("/users/:id", getUser)
// 	e.GET("/show", show)
// 	e.PUT("/users/:id", updateUser)
// 	e.DELETE("/users/:id", deleteUser)

// 	e.POST("/save", save)

// 	e.Logger.Fatal(e.Start(":1323"))

// }

// func save(c echo.Context) error {
// 	name := c.FormValue("name")
// 	avatar, err := c.FormFile("avatar")
// 	if err != nil {
// 		return err
// 	}

// 	// 파일 open
// 	src, err := avatar.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close() // 파일 닫기

// 	// 파일 저장
// 	dst, err := os.Create(avatar.Filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()

// 	if _, err := io.Copy(dst, src); err != nil {
// 		return err
// 	}

// 	// curl -F "name=Joe Smith" -F "avatar=@/mnt/c/Users/JIHEA/Pictures/comoozi.png" http://localhost:1323/save
// 	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")

// }

// // func save(c echo.Context) error {
// // 	name := c.FormValue("name")
// // 	email := c.FormValue("email")
// // 	return c.String(http.StatusOK, "name: "+name+", email: "+email)
// // }

// func deleteUser(c echo.Context) error {
// 	id := c.Param(("id"))
// 	return c.String(http.StatusOK, id)
// }

// func updateUser(c echo.Context) error {
// 	id := c.Param(("id"))
// 	return c.String(http.StatusOK, id)
// }

// func getUser(c echo.Context) error {
// 	id := c.Param(("id"))
// 	return c.String(http.StatusOK, id)
// }

// // e.GET("/show", show)
// func show(c echo.Context) error {
// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")
// 	return c.String(http.StatusOK, "team:"+team+", member:"+member)
// }
