package code

import "fmt"

type Statement interface {
	Code() string
}

type RawStmt string

func (rs RawStmt) Code() string {
	return string(rs)
}

type DeclVarStmt struct {
	Name  string
	Type  Type
	Value Statement
}

func (dvs DeclVarStmt) Code() string {
	if dvs.Value != nil {
		return fmt.Sprintf("var %s = %s", dvs.Name, dvs.Value.Code())
	} else {
		return fmt.Sprintf("var %s %s", dvs.Name, dvs.Type.RealName())
	}
}

type DeclColonStmt struct {
	Left  ListCommaStmt
	Right Statement
}

func (dcs DeclColonStmt) Code() string {
	return fmt.Sprintf("%s := %s", dcs.Left.Code(), dcs.Right.Code())
}

type ReturnStmt struct {
	ListCommaStmt ListCommaStmt
}

func (rs ReturnStmt) Code() string {
	return "return " + rs.ListCommaStmt.Code()
}

type ListCommaStmt []Statement

func (lcs ListCommaStmt) Code() string {
	result := ""
	for index, cs := range lcs {
		if index != len(lcs)-1 {
			result += cs.Code() + ", "
		} else {
			result += cs.Code()
		}
	}
	return result
}

type IfBlockStmt struct {
	Condition []Statement
	Body      Body
}

func (ibt IfBlockStmt) Code() string {
	result := "if "
	for _, c := range ibt.Condition {
		result += c.Code()
	}
	result += fmt.Sprintf("{ \n\t%s\n}", ibt.Body.GetCode())
	return result
}

type ForRangeBlockStmt struct {
	RangeName string
	Key       string
	Value     string
	Body      Body
}

func (rbs ForRangeBlockStmt) Code() string {
	key := "_"
	value := "_"
	if rbs.Key != "" {
		key = rbs.Key
	}
	if rbs.Value != "" {
		value = rbs.Value
	}
	return fmt.Sprintf("for %s, %s := range %s {\n%s\n}",
		key, value, rbs.RangeName, rbs.Body.GetCode())
}

type MapStmt struct {
	Name string
	Pair []MapPair
}

func (ms MapStmt) Code() string {
	result := ms.Name + "{\n"
	for _, m := range ms.Pair {
		result += "\t" + m.Code() + "\n"
	}
	result += "}"
	return result
}

type MapPair struct {
	Key   Statement
	Value Statement
}

func (mp MapPair) Code() string {
	return fmt.Sprintf("\"%s\": %s,", mp.Key.Code(), mp.Value.Code())
}

type SliceStmt struct {
	Name   string
	Values []MapPair
}

func (ss SliceStmt) Code() string {
	result := ss.Name + "{\n"
	for _, value := range ss.Values {
		result += "\t{\n\t" + value.Code() + "\n\t},"
	}
	result += "}"
	return result
}

type CallStmt struct {
	Caller   Statement
	CallName string
	Args     ListCommaStmt
}

func (cs CallStmt) Code() string {
	if cs.Caller == nil {
		return cs.CallName + "(" + cs.Args.Code() + ")"
	} else {
		return cs.Caller.Code() + "." + cs.CallName + "(" + cs.Args.Code() + ")"
	}
}

type ChainStmt []Chain

type Chain struct {
	CallName string
	Args     ListCommaStmt
}

func (cs ChainStmt) ChainCall(chain Chain) ChainStmt {
	if cs == nil {
		cs = make([]Chain, 0, 10)
	}
	cs = append(cs, chain)
	return cs
}

func (cs ChainStmt) Code() string {
	result := ""
	for index, chain := range cs {
		if index != len(cs)-1 {
			result += chain.CallName + "(" + chain.Args.Code() + ")."
		} else {
			result += chain.CallName + "(" + chain.Args.Code() + ")"
		}
	}
	return result
}

type SliceAppendsStmt []SliceAppendStmt

func (sas SliceAppendsStmt) Code() string {
	result := ""
	for index, appendStmt := range sas {
		if index != len(sas)-1 {
			result += appendStmt.Code() + "\n"
		} else {
			result += appendStmt.Code()
		}
	}
	return result
}

type SliceAppendStmt struct {
	SliceName  string
	AppendData Statement
}

func (sas SliceAppendStmt) Code() string {
	return fmt.Sprintf("%s = append(%s, %s)", sas.SliceName, sas.SliceName, sas.AppendData.Code())
}

type AnonymousFuncStmt struct {
	Params  Params
	Returns Returns
	Body    Body
}

func (afs AnonymousFuncStmt) Code() string {
	return fmt.Sprintf("func%s %s {\n%s\n}", afs.Params.GetCode(),
		afs.Returns.GetCode(), afs.Body.GetCode())
}
