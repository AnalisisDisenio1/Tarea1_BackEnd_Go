package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	//"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
    "net/http"
)

/*var (
	//db            gorm.DB
	sqlConnection string
)*/

type Item struct {
	Id          int64  `form:"id"`
	Title       string `form:"title"`
	Description string `form:"description"`
	UserName    string `form:"user_name"`
}

func main() {
	//var err error

	/*sqlConnection = "root:NexeR2995!!@/ad1_t1"

	db, err := gorm.Open("mysql", sqlConnection)

	if err != nil {
		panic(err)
		return
	}
    */
	m := martini.Classic()

	m.Use(render.Renderer())
	m.Get("/", func(r render.Render) {
		/*var retData struct {
			Items []Item
		}

		db.Find(&retData.Items)
        */
		r.HTML(http.StatusOK, "index", nil)
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
