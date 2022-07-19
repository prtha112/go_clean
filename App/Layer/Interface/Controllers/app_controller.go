package Controllers

type AppController struct {
	User interface{ UserController }
}
