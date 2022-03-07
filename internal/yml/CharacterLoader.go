package yml

import (
	"github.com/waigoma/GenshinCalender/internal/genshin/character"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
)

var characters []character.Character

func LoadCharacters() []character.Character {
	s, _ := ioutil.ReadFile("assets/data/characters.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return []character.Character{}
	}

	for key, value := range yamlData {
		var nation string

		switch key {
		case "mondstadt":
			nation = "モンド"
		case "liyue":
			nation = "璃月"
		case "inazuma":
			nation = "稲妻"
		}

		var enName string
		var jpName string
		var book string

		for ke, val := range value.(map[interface{}]interface{}) {
			enName = ke.(string)
			for k, v := range val.(map[interface{}]interface{}) {
				switch k {
				case "name":
					jpName = v.(string)
				case "talentBook":
					book = v.(string)
				}
			}

			characters = append(characters, character.Character{ENName: enName, JPName: jpName, ENNation: key, JPNation: nation, TalentBook: book})
		}
	}

	sortCharacters()

	return characters
}

func sortCharacters() {
	sort.Slice(characters, func(i, j int) bool { return characters[i].ENName < characters[j].ENName })
}
