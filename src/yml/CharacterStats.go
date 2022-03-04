package yml

type CharacterStats struct {
	Character Character
	Talent    []Talent
}

type Talent struct {
	Name  string
	Count int
}
