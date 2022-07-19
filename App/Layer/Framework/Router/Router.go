package Router

import (
	"fmt"

	controller "go_clean/App/Layer/Interface/Controllers"
	"log"

	config "go_clean/App/Config"
	datastore "go_clean/App/Layer/Framework/Database"
	registry "go_clean/App/Layer/Interface/Registry"

	"github.com/labstack/echo"
)

func Run() {
	db := datastore.NewDB()
	// db.LogMode(true)
	// defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })

	return e
}
