package yml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
)

var characters Characters

type Characters struct {
	Characters []Character
}

type Character struct {
	ENName     string
	JPName     string
	TalentBook string
}

func LoadCharacters() Characters {
	s, _ := ioutil.ReadFile("data/characters.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return Characters{}
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

		characters.Characters = append(characters.Characters, Character{ENName: key, JPName: nm, TalentBook: tb})
	}

	sortCharacters()

	return characters
}

func sortCharacters() {
	sort.Slice(characters.Characters, func(i, j int) bool { return characters.Characters[i].ENName < characters.Characters[j].ENName })
}
