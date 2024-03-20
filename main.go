package main

import (
	"blog/routes"
	"blog/utils"
)

func main(){
	utils.LoadEnv()
	utils.LoadDatabase()
	routes.ServeApplication()
}