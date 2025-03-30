package andi

import (
	"reflect"

	"github.com/JorgeGorrito/anise-dependency-injection/andi/errors"
	"github.com/JorgeGorrito/anise-dependency-injection/andi/port/in"
	"github.com/JorgeGorrito/anise-dependency-injection/andi/types"
)

func GetAbstractType[T any]() types.Abstract {
	var abstractType reflect.Type = reflect.TypeOf((*T)(nil)).Elem()
	if abstractType.Kind() != reflect.Interface {
		panic(errors.ErrAbstractMustBeInterface)
	}
	return abstractType
}

func Inject[T any](dependencyResolver in.Resolver) T {
	return dependencyResolver.Resolve(GetAbstractType[T]()).(T)
}
