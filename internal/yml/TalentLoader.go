package yml

import (
	talent2 "github.com/waigoma/GenshinCalender/internal/genshin/talent"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadTalentBooks() []talent2.Book {
	s, _ := ioutil.ReadFile("assets/data/talentBooks.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return []talent2.Book{}
	}

	var talentBooks []talent2.Book

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

		talentBooks = append(talentBooks, talent2.Book{ENName: key, Day: day, RarityName: rarityName})
	}

	return talentBooks
}

func LoadTalentBookCounts() []talent2.BookCount {
	s, _ := ioutil.ReadFile("assets/data/talentBookCount.yml")

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return []talent2.BookCount{}
	}

	var talentBookCounts []talent2.BookCount

	for key, value := range yamlData {
		var bookCount = make(map[string]int)

		for k, val := range value.(map[interface{}]interface{}) {
			bookCount[k.(string)] = val.(int)
		}

		talentBookCounts = append(talentBookCounts, talent2.BookCount{NtoM: key, BookCount: bookCount})
	}

	return talentBookCounts
}
