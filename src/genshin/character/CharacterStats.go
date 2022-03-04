package character

type Stats struct {
	Character Character
	Talent    []Talent
	Day       []string
}

type Talent struct {
	Type  string
	Name  string
	Count int
}
