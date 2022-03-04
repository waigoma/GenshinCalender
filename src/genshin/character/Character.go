package character

var characters []Character

type Character struct {
	ENName     string
	JPName     string
	TalentBook string
}

func Initialize(charactersList []Character) {
	characters = charactersList
}

func GetAllCharacters() []Character {
	return characters
}

func GetCharacter(name string) Character {
	for _, c := range characters {
		if c.ENName == name {
			return c
		}
	}

	return Character{}
}
