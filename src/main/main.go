package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waigoma/GenshinCalender/src/yml"
	"net/http"
)

var characters yml.Characters

func main() {
	characters = yml.LoadCharacters()

	//for _, chara := range characters.Characters {
	//	fmt.Println(chara.ENName, chara.JPName, chara.Type)
	//}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", indexGetHandle)
	router.POST("/", indexPostHandle)

	err := router.Run()
	if err != nil {
		return
	}

}

func indexGetHandle(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"characters": characters.Characters,
		})
}

func indexPostHandle(ctx *gin.Context) {
	selectedCharacters := ctx.PostFormArray("selectCharacter")
	var selectedCharactersList []yml.Character

	// 選択したキャラクター
	for _, sc := range selectedCharacters {
		// データとして用意されたキャラクター
		for _, chara := range characters.Characters {
			// 名前一致で選択キャラリストに追加
			if chara.ENName == sc {
				selectedCharactersList = append(selectedCharactersList, chara)
			}
		}
	}

	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"characters":         characters.Characters,
			"selectedCharacters": selectedCharactersList,
		})
}
