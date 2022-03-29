package character

type Stats struct {
	Character Character
	Talent    []Talent
	Mora      int
	Day       []string
}

type Talent struct {
	Type  string
	Name  string
	Count int
}
