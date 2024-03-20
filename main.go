package main

import (
	"blog/utils"
)

func main(){
	utils.LoadEnv()
	utils.LoadDatabase()
}