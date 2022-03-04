package character

type Stats struct {
	Character Character
	Talent    []Talent
}

type Talent struct {
	Type  string
	Name  string
	Count int
}
