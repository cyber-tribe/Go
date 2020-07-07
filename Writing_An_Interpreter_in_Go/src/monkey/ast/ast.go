package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string // デバック用のメソッド
	String() string
}

type Statement interface {
	Node
	statementNode() // ダミーメソッド
}

type Expression interface {
	Node
	expressionNode() // ダミーメソッド
}

type Program struct {
	// root Node
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 束縛識別子
	Value Expression  // 値を生成する式
}

type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression  // 返す式
}

type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression  // 式
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token // 前置トークン Example: !
	Operator string      // 演算子
	Right    Expression  // オペランド
}

type InfixExpression struct {
	Token    token.Token // 演算子トークン. Example: +
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean struct {
	Token token.Token // 真偽値トークン
	Value bool        // 真偽値
}

type IfExpression struct {
	Token       token.Token     // 'if' トークン
	Condition   Expression      // 条件式
	Consequence *BlockStatement // ifの処理
	Alternative *BlockStatement // elseの処理
}

type BlockStatement struct {
	Token      token.Token // { トークン
	Statements []Statement // { }の中身
}

type FunctionLiteral struct {
	Token      token.Token     // 'fn' トークン
	Parameters []*Identifier   // 引数
	Body       *BlockStatement // 関数本体
}

func (p *Program) TokenLiteral() string {
	// メソッド
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode()              {}                          // メソッド
func (ls *LetStatement) TokenLiteral() string        { return ls.Token.Literal } // メソッド
func (i *Identifier) expressionNode()                {}                          // メソッド
func (i *Identifier) TokenLiteral() string           { return i.Token.Literal }  // メソッド
func (rs *ReturnStatement) statementNode()           {}                          // メソッド
func (rs *ReturnStatement) TokenLiteral() string     { return rs.Token.Literal } // メソッド
func (es *ExpressionStatement) statementNode()       {}                          // メソッド
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal } // メソッド
func (il *IntegerLiteral) expressionNode()           {}                          // メソッド
func (il *IntegerLiteral) TokenLiteral() string      { return il.Token.Literal } // メソッド
func (pe *PrefixExpression) expressionNode()         {}                          // メソッド
func (pe *PrefixExpression) TokenLiteral() string    { return pe.Token.Literal } // メソッド
func (oe *InfixExpression) expressionNode()          {}                          // メソッド
func (oe *InfixExpression) TokenLiteral() string     { return oe.Token.Literal }
func (b *Boolean) expressionNode()                   {}                          // メソッド
func (b *Boolean) TokenLiteral() string              { return b.Token.Literal }  // メソッド
func (ie *IfExpression) expressionNode()             {}                          // メソッド
func (ie *IfExpression) TokenLiteral() string        { return ie.Token.Literal } // メソッド
func (bs *BlockStatement) expressionNode()           {}                          // メソッド
func (bs *BlockStatement) TokenLiteral() string      { return bs.Token.Literal } // メソッド
func (fl *FunctionLiteral) expressionNode()          {}                          // メソッド
func (fl *FunctionLiteral) TokenLiteral() string     { return fl.Token.Literal } // メソッド

func (p *Program) String() string {
	// メソッド
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ls *LetStatement) String() string {
	// メソッド
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) String() string {
	// メソッド
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	// メソッド
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string { return i.Value } // メソッド

func (il *IntegerLiteral) String() string { return il.Token.Literal } // メソッド

func (pe *PrefixExpression) String() string {
	// メソッド
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (oe *InfixExpression) String() string {
	// メソッド
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (b *Boolean) String() string { return b.Token.Literal } // メソッド

func (ie *IfExpression) String() string {
	// メソッド
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

func (bs *BlockStatement) String() string {
	// メソッド
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
