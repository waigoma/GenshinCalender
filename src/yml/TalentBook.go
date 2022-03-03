package yml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var talentBooks TalentBooks

type TalentBooks struct {
	TalentBooks []TalentBook
}

type TalentBook struct {
	ENName     string
	Day        []string
	RarityName map[string]string
}

func LoadTalentBooks() TalentBooks {
	s, _ := ioutil.ReadFile("data/talentBooks.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return TalentBooks{}
	}

	for key, value := range yamlData {
		var day []string
		var rarityName = make(map[string]string)

		for k, val := range value.(map[interface{}]interface{}) {
			switch k {
			case "day":
				for _, v := range val.([]interface{}) {
					day = append(day, v.(string))
				}
			default:
				rarityName[k.(string)] = val.(string)
			}
		}

		talentBooks.TalentBooks = append(talentBooks.TalentBooks, TalentBook{ENName: key, Day: day, RarityName: rarityName})
	}

	return talentBooks
}
