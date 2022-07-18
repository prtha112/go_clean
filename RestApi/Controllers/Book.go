package Controllers

import (
	"go_clean/RestApi/Config"
	"go_clean/RestApi/Models"
	"go_clean/RestApi/Utility"

	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBook(&book)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func AddNewBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.AddNewBook(&book)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models.PutOneBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func ListBookCustom(c *gin.Context) {
	var result = []struct {
		Id   int
		Name string
	}{}

	rows, err := Config.DB.Raw("select id, name from book").Rows()
	defer rows.Close()
	for rows.Next() {
		Config.DB.ScanRows(rows, &result)
	}
	if err != nil {
		Utility.JSON(c, 500, result)
	} else {
		Utility.JSON(c, 200, result)
	}
}

func GetOneBookCustom(c *gin.Context) {
	id := c.Params.ByName("id")

	var result = struct {
		Id   int
		Name string
	}{}

	rows, err := Config.DB.Raw("select id, name from book where id = ?", id).Rows()
	defer rows.Close()
	for rows.Next() {
		Config.DB.ScanRows(rows, &result)
	}
	if err != nil {
		Utility.JSON(c, 500, result)
	} else {
		Utility.JSON(c, 200, result)
	}
}