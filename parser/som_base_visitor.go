// Code generated from parser/Som.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Som

import "github.com/antlr4-go/antlr/v4"

type BaseSomVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSomVisitor) VisitClassdef(ctx *ClassdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitSuperclass(ctx *SuperclassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitInstanceFields(ctx *InstanceFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitClassFields(ctx *ClassFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitMethod(ctx *MethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitUnaryPattern(ctx *UnaryPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBinaryPattern(ctx *BinaryPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitKeywordPattern(ctx *KeywordPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitMethodBlock(ctx *MethodBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitUnarySelector(ctx *UnarySelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBinarySelector(ctx *BinarySelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBlockContents(ctx *BlockContentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLocalDefs(ctx *LocalDefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBlockBody(ctx *BlockBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitResult(ctx *ResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitAssignation(ctx *AssignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitAssignments(ctx *AssignmentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitEvaluation(ctx *EvaluationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitMessages(ctx *MessagesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitUnaryMessage(ctx *UnaryMessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBinaryMessage(ctx *BinaryMessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBinaryOperand(ctx *BinaryOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitKeywordMessage(ctx *KeywordMessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitFormula(ctx *FormulaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitNestedTerm(ctx *NestedTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteralArray(ctx *LiteralArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteralNumber(ctx *LiteralNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteralDecimal(ctx *LiteralDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitNegativeDecimal(ctx *NegativeDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteralInteger(ctx *LiteralIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteralDouble(ctx *LiteralDoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteralSymbol(ctx *LiteralSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitLiteralString(ctx *LiteralStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitKeywordSelector(ctx *KeywordSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitNestedBlock(ctx *NestedBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBlockPattern(ctx *BlockPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSomVisitor) VisitBlockArguments(ctx *BlockArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
