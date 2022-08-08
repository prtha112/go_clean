package Gin

import (
	"fmt"
	"log"

	controller "go_clean/App/Layer/Interface/Controllers"

	config "go_clean/App/Layer/Framework/Config"
	datastore "go_clean/App/Layer/Framework/Database"
	registry "go_clean/App/Layer/Interface/Registry"
	model "go_clean/App/Layer/Entity/Model"

	"github.com/labstack/echo"
)

func Run() {
	config.ReadConfig()
	db := datastore.NewDB()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = NewRouter(e, r.NewAppController())

	datastore.NewDB().AutoMigrate(&model.User{})

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	return e
}
