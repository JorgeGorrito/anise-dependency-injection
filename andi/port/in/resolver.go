package in

import "github.com/JorgeGorrito/anise-dependency-injection/andi/types"

type Resolver interface {
	Resolve(abstract types.Abstract) types.PtrConcreteType
}
