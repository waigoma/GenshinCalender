package yml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var talentBookCounts TalentBookCounts

type TalentBookCounts struct {
	TalentBookCounts []TalentBookCount
}

type TalentBookCount struct {
	NtoM      string
	BookCount map[string]int
}

func LoadTalentBookCounts() TalentBookCounts {
	s, _ := ioutil.ReadFile("data/talentBookCount.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return TalentBookCounts{}
	}

	for key, value := range yamlData {
		var bookCount = make(map[string]int)

		for k, val := range value.(map[interface{}]interface{}) {
			bookCount[k.(string)] = val.(int)
		}

		talentBookCounts.TalentBookCounts = append(talentBookCounts.TalentBookCounts, TalentBookCount{NtoM: key, BookCount: bookCount})
	}

	return talentBookCounts
}
