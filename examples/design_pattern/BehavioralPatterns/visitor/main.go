package main

/*
the Visitor Pattern: 访客模式
	> The Visitor Pattern works by separating the algorithm from the object structure.

*/

type Visitor interface {
	VisitManager(manager *Manager)
	VisitEngineer(engineer *Engineer)
	VisitSalesman(salesman *Salesman)
}

type Employee interface {
	Accept(visitor Visitor)
}

type Manager struct {
	Name   string
	Salary float64
}

func (m *Manager) Accept(visitor Visitor) {
	visitor.VisitManager(m)
}

type Engineer struct {
	Name   string
	Salary float64
}

func main() {

}
