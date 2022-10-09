package owners

import "bank-go/src/occupations"

type Owner struct {
	Name       string
	CPF        int
	Occupation occupations.Occupation
}
