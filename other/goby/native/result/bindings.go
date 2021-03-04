package result

import (
	"fmt"

	vm "github.com/goby-lang/goby/vm"
	"github.com/goby-lang/goby/vm/classes"
	errors "github.com/goby-lang/goby/vm/errors"
)

// DO NOT EDIT THIS FILE MANUALLY
// This code has been generated by github.com/goby-lang/goby/cmd/binder

func init() {
	vm.RegisterExternalClass(
		"result", vm.NewExternalClassLoader(
			"Result",
			"result.gb",
			map[string]vm.Method{
				"empty": bindingResultEmpty,
				"new":   bindingResultNew,
			},
			map[string]vm.Method{
				"method_missing": bindingResultMethodMissing,
				"or":             bindingResultOr,
			}))
}

var staticResult = new(Result)

// bindingResultNew is a class method binding for Result.New
func bindingResultNew(receiver vm.Object, line int, t *vm.Thread, args []vm.Object) vm.Object {
	r := staticResult
	if len(args) != 2 {
		return t.VM().InitErrorObject(errors.ArgumentError, line, errors.WrongNumberOfArgument, 2, len(args))
	}
	arg0, ok := args[0].(Object)
	if !ok {
		return t.VM().InitErrorObject(errors.TypeError, line, errors.WrongArgumentTypeFormat, classes.ObjectClass, args[0].Class().Name)
	}

	arg1, ok := args[1].(Object)
	if !ok {
		return t.VM().InitErrorObject(errors.TypeError, line, errors.WrongArgumentTypeFormat, classes.ObjectClass, args[1].Class().Name)
	}

	return r.New(t, arg0, arg1)
}

// bindingResultEmpty is a class method binding for Result.Empty
func bindingResultEmpty(receiver vm.Object, line int, t *vm.Thread, args []vm.Object) vm.Object {
	r := staticResult
	if len(args) != 0 {
		return t.VM().InitErrorObject(errors.ArgumentError, line, errors.WrongNumberOfArgument, 0, len(args))
	}
	return r.Empty(t)
}

// bindingResultMethodMissing is an instance method binding for *Result.MethodMissing
func bindingResultMethodMissing(receiver vm.Object, line int, t *vm.Thread, args []vm.Object) vm.Object {
	r, ok := receiver.(*Result)
	if !ok {
		panic(fmt.Sprintf("Impossible receiver type. Wanted Result got %s", receiver))
	}
	if len(args) != 1 {
		return t.VM().InitErrorObject(errors.ArgumentError, line, errors.WrongNumberOfArgument, 1, len(args))
	}
	arg0, ok := args[0].(Object)
	if !ok {
		return t.VM().InitErrorObject(errors.TypeError, line, errors.WrongArgumentTypeFormat, classes.ObjectClass, args[0].Class().Name)
	}

	return r.MethodMissing(t, arg0)
}

// bindingResultOr is an instance method binding for *Result.Or
func bindingResultOr(receiver vm.Object, line int, t *vm.Thread, args []vm.Object) vm.Object {
	r, ok := receiver.(*Result)
	if !ok {
		panic(fmt.Sprintf("Impossible receiver type. Wanted Result got %s", receiver))
	}
	if len(args) != 0 {
		return t.VM().InitErrorObject(errors.ArgumentError, line, errors.WrongNumberOfArgument, 0, len(args))
	}
	return r.Or(t)
}