package types

import (
	"fmt"
	"sort"
	"strings"
)

// -----------------------------------------------------------------
type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Skills  `json:"skills"`
}

func (p Person) String() string {
	return fmt.Sprintf(`
Name: %s,
Address: %s,
Skills: %s,
`, p.Name, p.Address, p.Skills.Sort())
}

type Employee struct {
	Person     Person `json:"person"`
	ID         string `json:"id"`
	Department string `json:"department"`
	Status     Status `json:"status"`
}

func (e Employee) String() string {
	return fmt.Sprintf(`
Person: %s,
ID: %s.
Department: %s,
Status:%s`, e.ID,
		e.Person.String(), e.Department, e.Status.String())
}

type Employess []Employee

func (e Employess) Len() int           { return len(e) }
func (e Employess) Less(i, j int) bool { return strings.Compare(e[i].ID, e[j].ID) == -1 }
func (e Employess) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

var _ sort.Interface = Employess{}

func (e Employess) Sort() Employess {
	sort.Sort(e)
	return e
}

func (e Employess) String() string {
	var a string

	if len(e) == 0 {
		return ""
	}

	for _, employee := range e {
		a = a + "\n" + employee.String()
	}

	return a[:len(e)-1]
}
