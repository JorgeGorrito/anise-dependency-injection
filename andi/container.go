package andi

import (
	"github.com/JorgeGorrito/anise-dependency-injection/andi/errors"
	"github.com/JorgeGorrito/anise-dependency-injection/andi/types"
)

type Container struct {
	registry map[types.Abstract]*types.ConcreteTypeNode
}

func NewContainer() *Container {
	return &Container{
		registry: make(map[types.Abstract]*types.ConcreteTypeNode),
	}
}

func (c *Container) Bind(abstract types.Abstract, constructorConcreteTypeFunc types.ContructorConcreteFunc) {
	if _, ok := c.registry[abstract]; ok {
		panic(errors.NewErrTypeAlreadyRegistered(abstract.String()))
	}
	c.registry[abstract] = &types.ConcreteTypeNode{
		PtrConcreteType:             nil,
		ConstructorConcreteTypeFunc: constructorConcreteTypeFunc,
	}
}

func (c *Container) Resolve(abstract types.Abstract) types.PtrConcreteType {
	nodoConcreteType, ok := c.registry[abstract]
	if !ok {
		panic(errors.NewErrTypeNotRegister(abstract.String()))
	}
	if nodoConcreteType.PtrConcreteType == nil {
		nodoConcreteType.PtrConcreteType = nodoConcreteType.ConstructorConcreteTypeFunc()
	}
	return nodoConcreteType.PtrConcreteType
}
