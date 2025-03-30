package types

import "reflect"

type Abstract reflect.Type
type PtrConcreteType any
type ContructorConcreteFunc func() any

type ConcreteTypeNode struct {
	PtrConcreteType             PtrConcreteType
	ConstructorConcreteTypeFunc ContructorConcreteFunc
}
