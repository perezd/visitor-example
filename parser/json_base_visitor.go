// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JSON
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseJSONVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJSONVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitArr(ctx *ArrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
