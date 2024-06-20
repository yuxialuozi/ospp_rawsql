package code

import "fmt"

type Type interface {
	RealName() string
}

type IdentType string

func (st IdentType) RealName() string {
	return string(st)
}

type SelectorExprType struct {
	X   string
	Sel string
}

func (set SelectorExprType) RealName() string {
	return set.X + "." + set.Sel
}

type InterfaceType struct {
	Name    string
	Methods InterfaceMethods
}

func (it InterfaceType) RealName() string {
	return "interface{}"
}

type SliceType struct {
	ElementType Type
}

func (st SliceType) RealName() string {
	return "[]" + st.ElementType.RealName()
}

type MapType struct {
	KeyType   Type
	ValueType Type
}

func (mt MapType) RealName() string {
	return fmt.Sprintf("map[%s]%s", mt.KeyType.RealName(), mt.ValueType.RealName())
}

type StarExprType struct {
	RealType Type
}

func (set StarExprType) RealName() string {
	return "*" + set.RealType.RealName()
}
