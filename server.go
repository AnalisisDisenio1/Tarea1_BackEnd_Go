package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/martini-contrib/cors"
	//"fmt"
)

var (
	sqlConnection string
)

type User struct {
	User_id  int64  `form:"user_id"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {

	sqlConnection = "root:NexeR2995!!@/ad1_t1"

	db, err := gorm.Open("mysql", sqlConnection)

	if err != nil {
		panic(err)
		return
	}

	m := martini.Classic()

	allowCORSHandler := (cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	m.Use(render.Renderer())
	m.Get("/", allowCORSHandler,  func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
		})

	m.Get("/users", allowCORSHandler, func(r render.Render) {
		var retData struct {
			Users []User
		}

		db.Find(&retData.Users)

		r.JSON(200, retData.Users)
	})

	m.Get("/user/:id", allowCORSHandler, func(r render.Render, p martini.Params) {
		var retData struct {
			user User
		}
		db.Where("user_id = ?", p["id"]).Find(&retData.user)

		r.JSON(http.StatusOK, retData.user)
	})

	m.Get("/user/remove/:id", allowCORSHandler, func(r render.Render, p martini.Params) {
		var user User
		db.Where("user_id = ?", p["id"]).Delete(&user)
		r.JSON(http.StatusOK, "User deleted")
	})

	m.Post("/user/save", allowCORSHandler, binding.Bind(User{}), func(r render.Render, u User) {
		db.Save(&u)
		r.JSON(http.StatusOK, "User created")
	})

	m.Get("/**", func(r render.Render) {
		r.Redirect("/")
	})
	
	m.Post("/user/edit", allowCORSHandler, binding.Bind(User{}), func(r render.Render, u User){
		db.Where("user_id = ?", u.User_id).Model(&u).Updates(User{Username: u.Username, Password: u.Password})
		r.JSON(http.StatusOK, "User edited")
	})

	m.Run()
}
