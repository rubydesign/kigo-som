// Code generated from parser/Som.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Som

import "github.com/antlr4-go/antlr/v4"

// SomListener is a complete listener for a parse tree produced by SomParser.
type SomListener interface {
	antlr.ParseTreeListener

	// EnterClassdef is called when entering the classdef production.
	EnterClassdef(c *ClassdefContext)

	// EnterSuperclass is called when entering the superclass production.
	EnterSuperclass(c *SuperclassContext)

	// EnterInstanceFields is called when entering the instanceFields production.
	EnterInstanceFields(c *InstanceFieldsContext)

	// EnterClassFields is called when entering the classFields production.
	EnterClassFields(c *ClassFieldsContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterUnaryPattern is called when entering the unaryPattern production.
	EnterUnaryPattern(c *UnaryPatternContext)

	// EnterBinaryPattern is called when entering the binaryPattern production.
	EnterBinaryPattern(c *BinaryPatternContext)

	// EnterKeywordPattern is called when entering the keywordPattern production.
	EnterKeywordPattern(c *KeywordPatternContext)

	// EnterMethodBlock is called when entering the methodBlock production.
	EnterMethodBlock(c *MethodBlockContext)

	// EnterUnarySelector is called when entering the unarySelector production.
	EnterUnarySelector(c *UnarySelectorContext)

	// EnterBinarySelector is called when entering the binarySelector production.
	EnterBinarySelector(c *BinarySelectorContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterBlockContents is called when entering the blockContents production.
	EnterBlockContents(c *BlockContentsContext)

	// EnterLocalDefs is called when entering the localDefs production.
	EnterLocalDefs(c *LocalDefsContext)

	// EnterBlockBody is called when entering the blockBody production.
	EnterBlockBody(c *BlockBodyContext)

	// EnterResult is called when entering the result production.
	EnterResult(c *ResultContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignation is called when entering the assignation production.
	EnterAssignation(c *AssignationContext)

	// EnterAssignments is called when entering the assignments production.
	EnterAssignments(c *AssignmentsContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterEvaluation is called when entering the evaluation production.
	EnterEvaluation(c *EvaluationContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterMessages is called when entering the messages production.
	EnterMessages(c *MessagesContext)

	// EnterUnaryMessage is called when entering the unaryMessage production.
	EnterUnaryMessage(c *UnaryMessageContext)

	// EnterBinaryMessage is called when entering the binaryMessage production.
	EnterBinaryMessage(c *BinaryMessageContext)

	// EnterBinaryOperand is called when entering the binaryOperand production.
	EnterBinaryOperand(c *BinaryOperandContext)

	// EnterKeywordMessage is called when entering the keywordMessage production.
	EnterKeywordMessage(c *KeywordMessageContext)

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterNestedTerm is called when entering the nestedTerm production.
	EnterNestedTerm(c *NestedTermContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralArray is called when entering the literalArray production.
	EnterLiteralArray(c *LiteralArrayContext)

	// EnterLiteralNumber is called when entering the literalNumber production.
	EnterLiteralNumber(c *LiteralNumberContext)

	// EnterLiteralDecimal is called when entering the literalDecimal production.
	EnterLiteralDecimal(c *LiteralDecimalContext)

	// EnterNegativeDecimal is called when entering the negativeDecimal production.
	EnterNegativeDecimal(c *NegativeDecimalContext)

	// EnterLiteralInteger is called when entering the literalInteger production.
	EnterLiteralInteger(c *LiteralIntegerContext)

	// EnterLiteralDouble is called when entering the literalDouble production.
	EnterLiteralDouble(c *LiteralDoubleContext)

	// EnterLiteralSymbol is called when entering the literalSymbol production.
	EnterLiteralSymbol(c *LiteralSymbolContext)

	// EnterLiteralString is called when entering the literalString production.
	EnterLiteralString(c *LiteralStringContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterKeywordSelector is called when entering the keywordSelector production.
	EnterKeywordSelector(c *KeywordSelectorContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterNestedBlock is called when entering the nestedBlock production.
	EnterNestedBlock(c *NestedBlockContext)

	// EnterBlockPattern is called when entering the blockPattern production.
	EnterBlockPattern(c *BlockPatternContext)

	// EnterBlockArguments is called when entering the blockArguments production.
	EnterBlockArguments(c *BlockArgumentsContext)

	// ExitClassdef is called when exiting the classdef production.
	ExitClassdef(c *ClassdefContext)

	// ExitSuperclass is called when exiting the superclass production.
	ExitSuperclass(c *SuperclassContext)

	// ExitInstanceFields is called when exiting the instanceFields production.
	ExitInstanceFields(c *InstanceFieldsContext)

	// ExitClassFields is called when exiting the classFields production.
	ExitClassFields(c *ClassFieldsContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitUnaryPattern is called when exiting the unaryPattern production.
	ExitUnaryPattern(c *UnaryPatternContext)

	// ExitBinaryPattern is called when exiting the binaryPattern production.
	ExitBinaryPattern(c *BinaryPatternContext)

	// ExitKeywordPattern is called when exiting the keywordPattern production.
	ExitKeywordPattern(c *KeywordPatternContext)

	// ExitMethodBlock is called when exiting the methodBlock production.
	ExitMethodBlock(c *MethodBlockContext)

	// ExitUnarySelector is called when exiting the unarySelector production.
	ExitUnarySelector(c *UnarySelectorContext)

	// ExitBinarySelector is called when exiting the binarySelector production.
	ExitBinarySelector(c *BinarySelectorContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitBlockContents is called when exiting the blockContents production.
	ExitBlockContents(c *BlockContentsContext)

	// ExitLocalDefs is called when exiting the localDefs production.
	ExitLocalDefs(c *LocalDefsContext)

	// ExitBlockBody is called when exiting the blockBody production.
	ExitBlockBody(c *BlockBodyContext)

	// ExitResult is called when exiting the result production.
	ExitResult(c *ResultContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignation is called when exiting the assignation production.
	ExitAssignation(c *AssignationContext)

	// ExitAssignments is called when exiting the assignments production.
	ExitAssignments(c *AssignmentsContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitEvaluation is called when exiting the evaluation production.
	ExitEvaluation(c *EvaluationContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitMessages is called when exiting the messages production.
	ExitMessages(c *MessagesContext)

	// ExitUnaryMessage is called when exiting the unaryMessage production.
	ExitUnaryMessage(c *UnaryMessageContext)

	// ExitBinaryMessage is called when exiting the binaryMessage production.
	ExitBinaryMessage(c *BinaryMessageContext)

	// ExitBinaryOperand is called when exiting the binaryOperand production.
	ExitBinaryOperand(c *BinaryOperandContext)

	// ExitKeywordMessage is called when exiting the keywordMessage production.
	ExitKeywordMessage(c *KeywordMessageContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitNestedTerm is called when exiting the nestedTerm production.
	ExitNestedTerm(c *NestedTermContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralArray is called when exiting the literalArray production.
	ExitLiteralArray(c *LiteralArrayContext)

	// ExitLiteralNumber is called when exiting the literalNumber production.
	ExitLiteralNumber(c *LiteralNumberContext)

	// ExitLiteralDecimal is called when exiting the literalDecimal production.
	ExitLiteralDecimal(c *LiteralDecimalContext)

	// ExitNegativeDecimal is called when exiting the negativeDecimal production.
	ExitNegativeDecimal(c *NegativeDecimalContext)

	// ExitLiteralInteger is called when exiting the literalInteger production.
	ExitLiteralInteger(c *LiteralIntegerContext)

	// ExitLiteralDouble is called when exiting the literalDouble production.
	ExitLiteralDouble(c *LiteralDoubleContext)

	// ExitLiteralSymbol is called when exiting the literalSymbol production.
	ExitLiteralSymbol(c *LiteralSymbolContext)

	// ExitLiteralString is called when exiting the literalString production.
	ExitLiteralString(c *LiteralStringContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitKeywordSelector is called when exiting the keywordSelector production.
	ExitKeywordSelector(c *KeywordSelectorContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitNestedBlock is called when exiting the nestedBlock production.
	ExitNestedBlock(c *NestedBlockContext)

	// ExitBlockPattern is called when exiting the blockPattern production.
	ExitBlockPattern(c *BlockPatternContext)

	// ExitBlockArguments is called when exiting the blockArguments production.
	ExitBlockArguments(c *BlockArgumentsContext)
}
