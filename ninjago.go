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

type NinjaGoInterface interface {
	toString() string
}

func (n NinjaGoRuleBase) toString() string {
	var ret = ""

	ret += fmt.Sprintf("rule %v\n", n.Name)
	ret += fmt.Sprintf("	command = %v", n.Cmd)
	ret += fmt.Sprintf("	description = %v", n.Desc)

	return ret
}
