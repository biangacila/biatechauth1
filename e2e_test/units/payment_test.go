package units

import (
	"fmt"
	"testing"
)

type Salary[W Working] interface {
	Pay(W)
}
type Working struct {
	Rate   float64
	Value  float64
	Amount float64
}

func Pay[T Salary[W], W Working](s T, work Working) {
	salary := work.Rate * work.Value
	work.Amount += salary
	s.Pay(W(work))
}

type Wage[W Working] struct {
	Name string
}
type Permanent[W Working] struct {
	Name string
}

func (w Wage[S]) Pay(sal S) {
	fmt.Printf("%s pays %.2f\n", w.Name, sal)
}
func (p Permanent[S]) Pay(sal S) {
	fmt.Printf("%s pays %.2f\n", p.Name, sal)
}

func TestSalary(t *testing.T) {
	w := Wage[Working]{Name: "Kasonga"}
	p := Permanent[Working]{Name: "Merveilleux"}
	wageWorked := Working{
		Rate:  20,
		Value: 44,
	}
	permanentWorked := Working{
		Rate:  20,
		Value: 44,
	}
	Pay(w, wageWorked)
	Pay(p, permanentWorked)
}
