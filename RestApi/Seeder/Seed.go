package Seeder

import (
	"go_clean/RestApi/Config"
	"go_clean/RestApi/Models"
)

func Load() {
	category := Models.Category{
		Name: "Novel",
	}
	Config.DB.Create(&category)
}
