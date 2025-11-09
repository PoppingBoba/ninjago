package ninjago

import (
	"fmt"
)

type NinjaGoBody struct {
	Cmd  string
	Desc string
}

type NinjaGoRuleBase struct {
	NinjaGoBody
	Name string
}

func MakeNinjaGoRule(name string, cmd string, desc string) *NinjaGoRuleBase {
	return &NinjaGoRuleBase{
		NinjaGoBody: NinjaGoBody{
			Cmd:  cmd,
			Desc: desc,
		},
		Name: name,
	}
}

type NinjaGoInterface interface {
	String() string
}

func (n NinjaGoRuleBase) String() string {
	var ret = ""

	ret += fmt.Sprintf("rule %v\n", n.Name)
	ret += fmt.Sprintf("	command = %v\n", n.Cmd)
	ret += fmt.Sprintf("	description = %v\n", n.Desc)

	return ret
}
