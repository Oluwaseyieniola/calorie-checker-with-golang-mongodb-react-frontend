package main

import (
	"github.com/Oluwaseyi/calorie-tracker-go/routes"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	

)


func main(){
	port:= os.Getenv("PORT")


	if port == ""{
		port ="3030"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.EntrybyID)
	router.GET("/ingredient/:ingredient", routes.EntriesByIngredient)


	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	router.Run(":"+port)
}
