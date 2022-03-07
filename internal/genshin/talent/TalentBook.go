package talent

var talentBooks []Book

type Book struct {
	ENName     string
	Day        []string
	RarityName map[string]string
}

func InitializeBook(talentBooksList []Book) {
	talentBooks = talentBooksList
}

func GetTalentBook(enName string) Book {
	for _, talentBook := range talentBooks {
		if talentBook.ENName == enName {
			return talentBook
		}
	}

	return Book{}
}

func GetAllTalentBooks() []Book {
	return talentBooks
}
