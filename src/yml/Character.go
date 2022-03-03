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
	ENName string
	JPName string
	Type   int
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
		var t int
		for k, v := range value.(map[interface{}]interface{}) {
			if k == "name" {
				nm = v.(string)
			} else if k == "type" {
				t = v.(int)
			}
		}

		characters.Characters = append(characters.Characters, Character{ENName: key, JPName: nm, Type: t})
	}

	sortCharacters()

	return characters
}

func sortCharacters() {
	sort.Slice(characters.Characters, func(i, j int) bool { return characters.Characters[i].Type < characters.Characters[j].Type })
}
