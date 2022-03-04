package character

type Stats struct {
	Character Character
	Talent    []Talent
}

type Talent struct {
	Name  string
	Count int
}
