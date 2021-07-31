package tryreflect

import (
	"fmt"
	"testing"
)

var global_e = Example{
	String:  "test",
	Number:  12345,
	Array:   [5]int{1, 2, 3, 4, 5},
	Slice:   []int{1, 2, 3, 4, 5},
	Maps:    make(map[string]string, 0),
	Channel: make(chan int, 0),
	Pointer: new(string),
}

func TestTryReflectType(t *testing.T) {
	e := global_e
	printlnType(e)
	printlnType(&e)
	printlnType(e.String)
	printlnType(&e.String)
	printlnType(e.Number)
	printlnType(&e.Number)
	printlnType(e.Array)
	printlnType(&e.Array)
	printlnType(e.Slice)
	printlnType(&e.Slice)
	printlnType(e.Maps)
	printlnType(&e.Maps)
	printlnType(&e.Channel)
	printlnType(e.Pointer)
	printlnType(&e.Pointer)
}

func TestTryReflectTypeElem(t *testing.T) {
	e := global_e
	printlnTypeElem(&e)
	printlnTypeElem(e.Slice)
	printlnTypeElem(e.Array)
	printlnTypeElem(e.Pointer)
	printlnTypeElem(e.Maps)
	printlnTypeElem(e.Channel)
	printlnTypeElem(e.Number)
	printlnTypeElem(e.String)
}

func TestTryReflectValue(t *testing.T) {
	e := global_e
	e.Maps["abc"] = "abc"
	TryReflectValue(e)
	TryReflectValue(&e)
	TryReflectValue(e.String)
	TryReflectValue(&e.String)
	TryReflectValue(e.Number)
	TryReflectValue(&e.Number)
	TryReflectValue(e.Array)
	TryReflectValue(&e.Array)
	TryReflectValue(e.Slice)
	TryReflectValue(&e.Slice)
	TryReflectValue(e.Maps)
	TryReflectValue(&e.Maps)
	TryReflectValue(e.Channel)
	TryReflectValue(&e.Channel)
	TryReflectValue(e.Pointer)
	TryReflectValue(&e.Pointer)
}

func TestTryTraverseStruct(t *testing.T) {
	e := global_e
	e.Maps["abc"] = "abc"
	TryTraverseStruct(e)
}

func TestTryTraverseStructPointer(t *testing.T) {
	e := global_e
	e.Maps["abc"] = "abc"
	TryTraverseStructPointer(&e)
}

func TestTryTraverseStructPointer2(t *testing.T) {
	str := "123456"
	TryTraverseStructPointer(&str)
}

func TestTryTraverseStructModifyElement(t *testing.T) {
	e := global_e
	TryTraverseStructModifyElement(&e)
	TryTraverseStruct(e)
}

func TestTryTraverseStructModifyElement2(t *testing.T) {
	e := global_e
	vBefore := TryGetElementFromStructByName("Number", e)
	if v, ok := vBefore.(int); ok {
		fmt.Println(v)
	}
	TryTraverseStructModifyElement(&e)
	vAfter := TryGetElementFromStructByName("Number", e)
	if v, ok := vAfter.(int); ok {
		fmt.Println(v)
	}
}

func TestTryGetElementFromStructByName(t *testing.T) {
	e := global_e
	v1 := TryGetElementFromStructByName("String", e)
	if v, ok := v1.(string); ok {
		fmt.Println(v)
	}
}

func TestTryGetElementFromStructByName2(t *testing.T) {
	e := global_e
	v1 := TryGetElementFromStructByName("String", &e)
	if v, ok := v1.(string); ok {
		fmt.Println(v)
	}
}

func TestTryGetElementFromStructByName3(t *testing.T) {
	e := global_e
	v1 := TryGetElementFromStructByName("Slice", &e)
	if v, ok := v1.([]int); ok {
		fmt.Println(v)
	}
}
