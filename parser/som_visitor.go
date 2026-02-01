// Code generated from parser/Som.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Som

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SomParser.
type SomVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SomParser#classdef.
	VisitClassdef(ctx *ClassdefContext) interface{}

	// Visit a parse tree produced by SomParser#superclass.
	VisitSuperclass(ctx *SuperclassContext) interface{}

	// Visit a parse tree produced by SomParser#instanceFields.
	VisitInstanceFields(ctx *InstanceFieldsContext) interface{}

	// Visit a parse tree produced by SomParser#classFields.
	VisitClassFields(ctx *ClassFieldsContext) interface{}

	// Visit a parse tree produced by SomParser#method.
	VisitMethod(ctx *MethodContext) interface{}

	// Visit a parse tree produced by SomParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by SomParser#unaryPattern.
	VisitUnaryPattern(ctx *UnaryPatternContext) interface{}

	// Visit a parse tree produced by SomParser#binaryPattern.
	VisitBinaryPattern(ctx *BinaryPatternContext) interface{}

	// Visit a parse tree produced by SomParser#keywordPattern.
	VisitKeywordPattern(ctx *KeywordPatternContext) interface{}

	// Visit a parse tree produced by SomParser#methodBlock.
	VisitMethodBlock(ctx *MethodBlockContext) interface{}

	// Visit a parse tree produced by SomParser#unarySelector.
	VisitUnarySelector(ctx *UnarySelectorContext) interface{}

	// Visit a parse tree produced by SomParser#binarySelector.
	VisitBinarySelector(ctx *BinarySelectorContext) interface{}

	// Visit a parse tree produced by SomParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SomParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by SomParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by SomParser#blockContents.
	VisitBlockContents(ctx *BlockContentsContext) interface{}

	// Visit a parse tree produced by SomParser#localDefs.
	VisitLocalDefs(ctx *LocalDefsContext) interface{}

	// Visit a parse tree produced by SomParser#blockBody.
	VisitBlockBody(ctx *BlockBodyContext) interface{}

	// Visit a parse tree produced by SomParser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by SomParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SomParser#assignation.
	VisitAssignation(ctx *AssignationContext) interface{}

	// Visit a parse tree produced by SomParser#assignments.
	VisitAssignments(ctx *AssignmentsContext) interface{}

	// Visit a parse tree produced by SomParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SomParser#evaluation.
	VisitEvaluation(ctx *EvaluationContext) interface{}

	// Visit a parse tree produced by SomParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by SomParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SomParser#messages.
	VisitMessages(ctx *MessagesContext) interface{}

	// Visit a parse tree produced by SomParser#unaryMessage.
	VisitUnaryMessage(ctx *UnaryMessageContext) interface{}

	// Visit a parse tree produced by SomParser#binaryMessage.
	VisitBinaryMessage(ctx *BinaryMessageContext) interface{}

	// Visit a parse tree produced by SomParser#binaryOperand.
	VisitBinaryOperand(ctx *BinaryOperandContext) interface{}

	// Visit a parse tree produced by SomParser#keywordMessage.
	VisitKeywordMessage(ctx *KeywordMessageContext) interface{}

	// Visit a parse tree produced by SomParser#formula.
	VisitFormula(ctx *FormulaContext) interface{}

	// Visit a parse tree produced by SomParser#nestedTerm.
	VisitNestedTerm(ctx *NestedTermContext) interface{}

	// Visit a parse tree produced by SomParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by SomParser#literalArray.
	VisitLiteralArray(ctx *LiteralArrayContext) interface{}

	// Visit a parse tree produced by SomParser#literalNumber.
	VisitLiteralNumber(ctx *LiteralNumberContext) interface{}

	// Visit a parse tree produced by SomParser#literalDecimal.
	VisitLiteralDecimal(ctx *LiteralDecimalContext) interface{}

	// Visit a parse tree produced by SomParser#negativeDecimal.
	VisitNegativeDecimal(ctx *NegativeDecimalContext) interface{}

	// Visit a parse tree produced by SomParser#literalInteger.
	VisitLiteralInteger(ctx *LiteralIntegerContext) interface{}

	// Visit a parse tree produced by SomParser#literalDouble.
	VisitLiteralDouble(ctx *LiteralDoubleContext) interface{}

	// Visit a parse tree produced by SomParser#literalSymbol.
	VisitLiteralSymbol(ctx *LiteralSymbolContext) interface{}

	// Visit a parse tree produced by SomParser#literalString.
	VisitLiteralString(ctx *LiteralStringContext) interface{}

	// Visit a parse tree produced by SomParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by SomParser#keywordSelector.
	VisitKeywordSelector(ctx *KeywordSelectorContext) interface{}

	// Visit a parse tree produced by SomParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by SomParser#nestedBlock.
	VisitNestedBlock(ctx *NestedBlockContext) interface{}

	// Visit a parse tree produced by SomParser#blockPattern.
	VisitBlockPattern(ctx *BlockPatternContext) interface{}

	// Visit a parse tree produced by SomParser#blockArguments.
	VisitBlockArguments(ctx *BlockArgumentsContext) interface{}
}
