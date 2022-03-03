package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/web"
	"github.com/waigoma/GenshinCalender/src/yml"
)

var characters yml.Characters

func main() {
	characters = yml.LoadCharacters()

	//for _, chara := range characters.Characters {
	//	fmt.Println(chara.ENName, chara.JPName, chara.Type)
	//}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	web.RegisterIndexHandler(router, characters)

	err := router.Run()

	if err != nil {
		return
	}

}
