package parse

import (
	"errors"
	"fmt"
	"github.com/fatih/camelcase"
	"ospp_rawsql/pkg/curd/code"
	"ospp_rawsql/pkg/curd/extract"
	"strings"
)

// InterfaceOperation is used to store the parsing results of the structure and for use by codegen packages.
type InterfaceOperation struct {
	BelongedToStruct *extract.IdlExtractStruct
	Operations       []Operation
}

const (
	Insert      = "Insert"
	Find        = "Find"
	Update      = "Update"
	Delete      = "Delete"
	Count       = "Count"
	Transaction = "Transaction"
	Bulk        = "Bulk"
)

const (
	One  = "One"
	Many = "Many"
)

type OperateMode int

const (
	OperateOne  = OperateMode(0)
	OperateMany = OperateMode(1)
)

func HandleOperations(structs []*extract.IdlExtractStruct) (result []*InterfaceOperation, err error) {
	for _, st := range structs {
		ifo := newInterfaceOperation()
		if err = ifo.parseInterfaceMethod(st); err != nil {
			return nil, err
		}
		result = append(result, ifo)
	}
	return
}

func newInterfaceOperation() *InterfaceOperation {
	return &InterfaceOperation{Operations: []Operation{}}
}

func (ifo *InterfaceOperation) parseInterfaceMethod(extractStruct *extract.IdlExtractStruct) error {
	for _, method := range extractStruct.InterfaceInfo.Methods {
		tokens := camelcase.Split(method.ParsedTokens)
		switch tokens[0] {
		case Insert:
			curParamIndex := new(int)
			*curParamIndex = 0
			ip := newInsertParse()
			if err := ip.parseInsert(method, curParamIndex, false); err != nil {
				return err
			}
			ifo.BelongedToStruct = extractStruct
			ifo.Operations = append(ifo.Operations, ip)

		case Find:

		case Update:

		case Delete:

		case Count:

		case Transaction:

		case Bulk:

		default:
			return newMethodSyntaxError(method.Name, "wrong operation name, should be Insert, Find, "+
				"Update, Delete, Count, Transaction, Bulk")
		}
	}

	return nil
}

// getFieldNameType is used to get field names and types in the specified structure.
//
//	input params description:
//	tokens: parsed tokens
//	extractStruct: the structure to which tokens belong
//	curIndex: point to the next token to be parsed
//	isFirst: if it is called in recursion
func getFieldNameType(tokens []string, extractStruct *extract.IdlExtractStruct, curIndex *int, isFirst bool) (names []string,
	types []code.Type, err error,
) {
	if len(tokens) == 0 {
		return nil, nil, errors.New("the length of the field name requested for parsing is empty")
	}

	for i := 0; i < len(tokens); i++ {
		flag := 0
		for _, field := range extractStruct.StructFields {
			if field.Name == tokens[i] || strings.Index(field.Name, tokens[i]) == 0 {
				s := ""
				hasFieldFlag := 0
				for j := i; j < len(tokens); j++ {
					s += tokens[j]
					if s == field.Name {
						hasFieldFlag = 1
						i = j
						*curIndex += i + 1
						break
					}
				}
				if hasFieldFlag == 0 {
					return nil, nil, fmt.Errorf("partially equal but unable to fully locate field name in %v", tokens[i:])
				}

				flag = 1
				if !field.IsBelongedToStruct {
					names = append(names, field.Tag.Get("bson"))
					types = append(types, field.Type)
					break
				} else {
					r, t, err := getFieldNameType(tokens[i+1:], field.BelongedToStruct, curIndex, false)
					// The final result of the structural field
					if err != nil {
						names = append(names, field.Tag.Get("bson"))
						types = append(types, field.Type)
						break
					}
					if len(r) != 1 {
						return nil, nil, fmt.Errorf("no field name corresponding to %v found", tokens[i:])
					}
					i += *curIndex
					names = append(names, field.Tag.Get("bson")+"."+r[0])
					types = append(types, t[0])
					break
				}
			}
		}
		if !isFirst && flag == 1 {
			break
		}
		if flag == 0 {
			return nil, nil, fmt.Errorf("no field name corresponding to %v found", tokens[i:])
		}
	}

	return
}
