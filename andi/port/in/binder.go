package in

import "github.com/JorgeGorrito/anise-dependency-injection/andi/types"

type Binder interface {
	Bind(abstract types.Abstract, constructorConcreteTypeFunc types.ContructorConcreteFunc)
}
