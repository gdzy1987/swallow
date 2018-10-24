package interpreter

import (
	"fmt"
	"strconv"
)

type Double struct {
	Ast
	token *Token
	value float64
}

func NewDouble(token *Token) *Double {
	num := &Double{token: token}

	if v, err := strconv.ParseFloat(token.value, 64); err != nil {
		g_error.error(fmt.Sprintf("传入无效数字类型：%v", token.value))
	} else {
		num.value = v
	}
	num.v = num
	return num
}

func (d *Double) ofToken() *Token {
	return d.token
}

func (n *Double) visit() (interface{}, error) {
	return n, nil
}

func (n *Double) eval() interface{} {
	return n.value
}

func (n *Double) String() string {
	if g_is_debug {
		return fmt.Sprintf("({type=%v}, {value=%f})", n.token.valueType, n.value)
	}
	return fmt.Sprintf("%f", n.value)
}

func (d Double) neg() interface{} {
	d.value = -d.value
	return d
}

func (n *Double) add(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.add(val.result[0])
	case *Integer:
		return NewDouble(&Token{value: fmt.Sprintf("%f", n.value+float64(val.value)), valueType: INT, pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v+%v", n.token, ast))
	case *Double:
		return NewDouble(&Token{value: fmt.Sprintf("%f", n.value+val.value), valueType: DOUBLE, pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v+%v", n.token, ast))
	}
	return nil
}

func (n *Double) minus(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.minus(val.result[0])
	case *Integer:
		return NewDouble(&Token{value: fmt.Sprintf("%f", n.value-float64(val.value)), valueType: INT, pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v-%v", n.token, ast))
	case *Double:
		return NewDouble(&Token{value: fmt.Sprintf("%f", n.value-val.value), valueType: STRING, pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v-%v", n.token, ast))
	}
	return nil
}

func (n *Double) multi(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.multi(val.result[0])
	case *Integer:
		return NewDouble(&Token{value: fmt.Sprintf("%f", n.value*float64(val.value)), valueType: INT, pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v*%v", n.token, ast))
	case *Double:
		return NewDouble(&Token{value: fmt.Sprintf("%f", n.value*val.value), valueType: STRING, pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v*%v", n.token, ast))
	}
	return nil
}

func (n *Double) div(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.div(val.result[0])
	case *Integer:
		return NewDouble(&Token{value: fmt.Sprintf("%d", n.value/float64(val.value)), valueType: INT, pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v/%v", n.token, val))
	case *Double:
		return NewDouble(&Token{value: fmt.Sprintf("%d", n.value/val.value), valueType: INT, pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v/%v", n.token, ast))
	}
	return nil
}

func (n *Double) great(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.great(val.result[0])
	case *Integer:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value > float64(val.value)), pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v>%v", n.token, val))
	case *Double:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value > float64(val.value)), pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v>%v", n.token, ast))
	}
	return nil
}

func (n *Double) less(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.less(val.result[0])
	case *Integer:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value < float64(val.value)), pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v<%v", n.token, val))
	case *Double:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value < val.value), pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v<%v", n.token, ast))
	}
	return nil
}

func (n *Double) geq(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.geq(val.result[0])
	case *Integer:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value >= float64(val.value)), pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v>=%v", n.token, val))
	case *Double:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value >= val.value), pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v>=%v", n.token, ast))
	}
	return nil
}

func (n *Double) leq(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.leq(val.result[0])
	case *Integer:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value <= float64(val.value)), pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v<=%v", n.token, val))
	case *Double:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value <= val.value), pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v<=%v", n.token, ast))
	}
	return nil
}

func (n *Double) equal(ast AstNode) interface{} {
	switch val := ast.(type) {
	case *Result:
		if val.num != 1 {
			g_error.error(fmt.Sprintf("右操作数个数应为1，但为%v", val.num))
		}
		return n.equal(val.result[0])
	case *Integer:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value == float64(val.value)), pos: n.token.pos, line: n.token.line, file: n.token.file})
	case *String:
		g_error.error(fmt.Sprintf("不支持%v==%v", n.token, val))
	case *Double:
		return NewBoolean(&Token{valueType: BOOLEAN, value: BoolToString(n.value == val.value), pos: n.token.pos, line: n.token.line, file: n.token.file})
	default:
		g_error.error(fmt.Sprintf("不支持%v==%v", n.token, ast))
	}
	return nil
}
