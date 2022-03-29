package yml

import (
	"github.com/waigoma/GenshinCalender/internal/genshin/talent"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadTalentBooks() []talent.Book {
	s, _ := ioutil.ReadFile("assets/data/talentBooks.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return []talent.Book{}
	}

	var talentBooks []talent.Book

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

		talentBooks = append(talentBooks, talent.Book{ENName: key, Day: day, RarityName: rarityName})
	}

	return talentBooks
}

func LoadTalentBookCounts() []talent.BookCount {
	s, _ := ioutil.ReadFile("assets/data/talentCount.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return []talent.BookCount{}
	}

	var talentBookCounts []talent.BookCount

	for key, value := range yamlData {
		var bookCount = make(map[string]int)

		for k, val := range value.(map[interface{}]interface{}) {
			bookCount[k.(string)] = val.(int)
		}

		talentBookCounts = append(talentBookCounts, talent.BookCount{NtoM: key, BookCount: bookCount})
	}

	return talentBookCounts
}
