package yml

import (
	"github.com/waigoma/GenshinCalender/internal/genshin/talent"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadTalentBooks(isLocal bool) []talent.Book {
	filename := "assets/data/talentBooks.yml"
	if !isLocal {
		filename = "/app/" + filename
	}
	s, _ := ioutil.ReadFile(filename)

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

func LoadTalentBookCounts(isLocal bool) []talent.BookCount {
	filename := "assets/data/talentCount.yml"
	if !isLocal {
		filename = "/app/" + filename
	}
	s, _ := ioutil.ReadFile(filename)

	var yamlData map[string]interface{}
	err := yaml.Unmarshal(s, &yamlData)

	if err != nil {
		return []talent.BookCount{}
	}

	var talentCounts []talent.BookCount

	for key, value := range yamlData {
		var mora int
		var bookCount = make(map[string]int)

		for k, val := range value.(map[interface{}]interface{}) {
			if k.(string) == "mora" {
				mora = val.(int)
				continue
			}

			bookCount[k.(string)] = val.(int)
			println(k.(string), val.(int))
		}

		talentCounts = append(talentCounts, talent.BookCount{NtoM: key, Mora: mora, BookCount: bookCount})
	}

	return talentCounts
}
