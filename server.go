package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"github.com/martini-contrib/binding"
	"fmt"
	"github.com/martini-contrib/render"
	"net/http"
)

var (
	//db            gorm.DB
	sqlConnection string
)

type User struct {
	User_id  int64  `form:"user_id"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	//var err error

	sqlConnection = "root:pass123@/ad1_t1"

	db, err := gorm.Open("mysql", sqlConnection)

	if err != nil {
		panic(err)
		return
	}

	m := martini.Classic()

	m.Use(render.Renderer())
	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})

	m.Get("/users", func(r render.Render) {
		var retData struct {
			Users []User
		}

		db.Find(&retData.Users)

		r.JSON(http.StatusOK, retData)
	})

	m.Get("/user/:id", func(r render.Render, p martini.Params) {
		var retData struct {
			user User
		}
		db.Where("user_id = ?", p["id"]).Find(&retData.user)
		fmt.Println(retData)

		r.JSON(http.StatusOK, retData.user)
	})

	m.Get("/user/remove/:id", func(r render.Render, p martini.Params) {
		var user User
		db.Where("user_id = ?", p["id"]).Delete(&user)
		r.Redirect("/")
	})

	m.Get("/**", func(r render.Render) {
		r.Redirect("/")
	})
	/*
		m.Get("/item/add", func(r render.Render) {
			var retData struct {
				Item Item
			}
			r.HTML(200, "item_edit", retData)
		})
		m.Post("/item/save", binding.Bind(Item{}), func(r render.Render, i Item) {
			db.Save(&i)
			r.Redirect("/")
		})
		m.Get("/item/edit/:id", func(r render.Render, p martini.Params) {
			var retData struct {
				Item Item
			}
			db.Where("id = ?", p["id"]).Find(&retData.Item)
			r.HTML(200, "item_edit", retData)
		})
		m.Get("/item/remove/:id", func(r render.Render, p martini.Params) {
			var item Item
			db.Where("id = ?", p["id"]).Delete(&item)
			r.Redirect("/")
		})
	*/
	m.Run()
}
