package yml

import (
	"github.com/waigoma/GenshinCalender/src/genshin/character"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
)

var characters []character.Character

func LoadCharacters() []character.Character {
	s, _ := ioutil.ReadFile("data/characters.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return []character.Character{}
	}

	for key, value := range yamlData {
		var nm string
		var tb string
		for k, v := range value.(map[interface{}]interface{}) {
			switch k {
			case "name":
				nm = v.(string)
			case "talentBook":
				tb = v.(string)
			}
		}

		characters = append(characters, character.Character{ENName: key, JPName: nm, TalentBook: tb})
	}

	sortCharacters()

	return characters
}

func sortCharacters() {
	sort.Slice(characters, func(i, j int) bool { return characters[i].ENName < characters[j].ENName })
}
