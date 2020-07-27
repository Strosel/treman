package main

//Rule is a rule
type Rule interface {
	String() string
	Valid(Roll) bool
}

//SumRule is a rule based on the sum of a roll
type SumRule struct {
	Name string
	Sum  int
}

func (sr SumRule) String() string {
	return sr.Name
}

func (sr SumRule) Valid(r Roll) bool {
	return sr.Sum == (r[0] + r[1])
}

//SetRule is a rule based on a specific roll
type SetRule struct {
	Name string
	Set  Roll
}

func (sr SetRule) String() string {
	return sr.Name
}

func (sr SetRule) Valid(r Roll) bool {
	lr := (sr.Set[0] == r[0] && sr.Set[1] == r[1])
	rl := (sr.Set[1] == r[0] && sr.Set[0] == r[1])
	return lr || rl
}

//SingleRule is a rule based on a single dice
type SingleRule struct {
	Name string
	Dice int
}

func (sr SingleRule) String() string {
	return sr.Name
}

func (sr SingleRule) Valid(r Roll) bool {
	return sr.Dice == r[0] || sr.Dice == r[1]
}

type specialRule struct {
	Name string
	Rule func(Roll) bool
}

func (sr specialRule) String() string {
	return sr.Name
}

func (sr specialRule) Valid(r Roll) bool {
	return sr.Rule(r)
}

//default rules
func drules() []Rule {
	return []Rule{
		specialRule{
			Name: "Treman",
			Rule: func(r Roll) bool {
				//Fast än treman dricker på 3,3 (då det är ny treman) ska det inte stå "treman dricker och ny treman"
				return (r[0] == 3 || r[1] == 3) && (r[0] != r[1])
			},
		},
		SetRule{
			Name: "Krig",
			Set:  Roll{1, 1},
		},
		SetRule{
			Name: "Utmaning",
			Set:  Roll{1, 2},
		},
		SetRule{
			Name: "En ferrari",
			Set:  Roll{1, 4},
		},
		SetRule{
			Name: "Ny Treman",
			Set:  Roll{3, 3},
		},
		SetRule{
			Name: "Jag har aldrig sett...",
			Set:  Roll{6, 6},
		},
		SetRule{
			Name: "Dela ut 2+2 klunkar",
			Set:  Roll{2, 2},
		},
		SetRule{
			Name: "Dela ut 4+4 klunkar",
			Set:  Roll{4, 4},
		},
		SetRule{
			Name: "Dela ut 5+5 klunkar",
			Set:  Roll{5, 5},
		},
		SumRule{
			Name: "Seven ahead",
			Sum:  7,
		},
		SumRule{
			Name: "Nine behind",
			Sum:  9,
		},
		SumRule{
			Name: "Finger på näsan",
			Sum:  11,
		},
	}
}
