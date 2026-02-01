// Code generated from parser/Som.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Som

import "github.com/antlr4-go/antlr/v4"

// BaseSomListener is a complete listener for a parse tree produced by SomParser.
type BaseSomListener struct{}

var _ SomListener = &BaseSomListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSomListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSomListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSomListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSomListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClassdef is called when production classdef is entered.
func (s *BaseSomListener) EnterClassdef(ctx *ClassdefContext) {}

// ExitClassdef is called when production classdef is exited.
func (s *BaseSomListener) ExitClassdef(ctx *ClassdefContext) {}

// EnterSuperclass is called when production superclass is entered.
func (s *BaseSomListener) EnterSuperclass(ctx *SuperclassContext) {}

// ExitSuperclass is called when production superclass is exited.
func (s *BaseSomListener) ExitSuperclass(ctx *SuperclassContext) {}

// EnterInstanceFields is called when production instanceFields is entered.
func (s *BaseSomListener) EnterInstanceFields(ctx *InstanceFieldsContext) {}

// ExitInstanceFields is called when production instanceFields is exited.
func (s *BaseSomListener) ExitInstanceFields(ctx *InstanceFieldsContext) {}

// EnterClassFields is called when production classFields is entered.
func (s *BaseSomListener) EnterClassFields(ctx *ClassFieldsContext) {}

// ExitClassFields is called when production classFields is exited.
func (s *BaseSomListener) ExitClassFields(ctx *ClassFieldsContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseSomListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseSomListener) ExitMethod(ctx *MethodContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseSomListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseSomListener) ExitPattern(ctx *PatternContext) {}

// EnterUnaryPattern is called when production unaryPattern is entered.
func (s *BaseSomListener) EnterUnaryPattern(ctx *UnaryPatternContext) {}

// ExitUnaryPattern is called when production unaryPattern is exited.
func (s *BaseSomListener) ExitUnaryPattern(ctx *UnaryPatternContext) {}

// EnterBinaryPattern is called when production binaryPattern is entered.
func (s *BaseSomListener) EnterBinaryPattern(ctx *BinaryPatternContext) {}

// ExitBinaryPattern is called when production binaryPattern is exited.
func (s *BaseSomListener) ExitBinaryPattern(ctx *BinaryPatternContext) {}

// EnterKeywordPattern is called when production keywordPattern is entered.
func (s *BaseSomListener) EnterKeywordPattern(ctx *KeywordPatternContext) {}

// ExitKeywordPattern is called when production keywordPattern is exited.
func (s *BaseSomListener) ExitKeywordPattern(ctx *KeywordPatternContext) {}

// EnterMethodBlock is called when production methodBlock is entered.
func (s *BaseSomListener) EnterMethodBlock(ctx *MethodBlockContext) {}

// ExitMethodBlock is called when production methodBlock is exited.
func (s *BaseSomListener) ExitMethodBlock(ctx *MethodBlockContext) {}

// EnterUnarySelector is called when production unarySelector is entered.
func (s *BaseSomListener) EnterUnarySelector(ctx *UnarySelectorContext) {}

// ExitUnarySelector is called when production unarySelector is exited.
func (s *BaseSomListener) ExitUnarySelector(ctx *UnarySelectorContext) {}

// EnterBinarySelector is called when production binarySelector is entered.
func (s *BaseSomListener) EnterBinarySelector(ctx *BinarySelectorContext) {}

// ExitBinarySelector is called when production binarySelector is exited.
func (s *BaseSomListener) ExitBinarySelector(ctx *BinarySelectorContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSomListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSomListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseSomListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseSomListener) ExitKeyword(ctx *KeywordContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseSomListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseSomListener) ExitArgument(ctx *ArgumentContext) {}

// EnterBlockContents is called when production blockContents is entered.
func (s *BaseSomListener) EnterBlockContents(ctx *BlockContentsContext) {}

// ExitBlockContents is called when production blockContents is exited.
func (s *BaseSomListener) ExitBlockContents(ctx *BlockContentsContext) {}

// EnterLocalDefs is called when production localDefs is entered.
func (s *BaseSomListener) EnterLocalDefs(ctx *LocalDefsContext) {}

// ExitLocalDefs is called when production localDefs is exited.
func (s *BaseSomListener) ExitLocalDefs(ctx *LocalDefsContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseSomListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseSomListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterResult is called when production result is entered.
func (s *BaseSomListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseSomListener) ExitResult(ctx *ResultContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSomListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSomListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignation is called when production assignation is entered.
func (s *BaseSomListener) EnterAssignation(ctx *AssignationContext) {}

// ExitAssignation is called when production assignation is exited.
func (s *BaseSomListener) ExitAssignation(ctx *AssignationContext) {}

// EnterAssignments is called when production assignments is entered.
func (s *BaseSomListener) EnterAssignments(ctx *AssignmentsContext) {}

// ExitAssignments is called when production assignments is exited.
func (s *BaseSomListener) ExitAssignments(ctx *AssignmentsContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSomListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSomListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterEvaluation is called when production evaluation is entered.
func (s *BaseSomListener) EnterEvaluation(ctx *EvaluationContext) {}

// ExitEvaluation is called when production evaluation is exited.
func (s *BaseSomListener) ExitEvaluation(ctx *EvaluationContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseSomListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseSomListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSomListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSomListener) ExitVariable(ctx *VariableContext) {}

// EnterMessages is called when production messages is entered.
func (s *BaseSomListener) EnterMessages(ctx *MessagesContext) {}

// ExitMessages is called when production messages is exited.
func (s *BaseSomListener) ExitMessages(ctx *MessagesContext) {}

// EnterUnaryMessage is called when production unaryMessage is entered.
func (s *BaseSomListener) EnterUnaryMessage(ctx *UnaryMessageContext) {}

// ExitUnaryMessage is called when production unaryMessage is exited.
func (s *BaseSomListener) ExitUnaryMessage(ctx *UnaryMessageContext) {}

// EnterBinaryMessage is called when production binaryMessage is entered.
func (s *BaseSomListener) EnterBinaryMessage(ctx *BinaryMessageContext) {}

// ExitBinaryMessage is called when production binaryMessage is exited.
func (s *BaseSomListener) ExitBinaryMessage(ctx *BinaryMessageContext) {}

// EnterBinaryOperand is called when production binaryOperand is entered.
func (s *BaseSomListener) EnterBinaryOperand(ctx *BinaryOperandContext) {}

// ExitBinaryOperand is called when production binaryOperand is exited.
func (s *BaseSomListener) ExitBinaryOperand(ctx *BinaryOperandContext) {}

// EnterKeywordMessage is called when production keywordMessage is entered.
func (s *BaseSomListener) EnterKeywordMessage(ctx *KeywordMessageContext) {}

// ExitKeywordMessage is called when production keywordMessage is exited.
func (s *BaseSomListener) ExitKeywordMessage(ctx *KeywordMessageContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseSomListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseSomListener) ExitFormula(ctx *FormulaContext) {}

// EnterNestedTerm is called when production nestedTerm is entered.
func (s *BaseSomListener) EnterNestedTerm(ctx *NestedTermContext) {}

// ExitNestedTerm is called when production nestedTerm is exited.
func (s *BaseSomListener) ExitNestedTerm(ctx *NestedTermContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSomListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSomListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralArray is called when production literalArray is entered.
func (s *BaseSomListener) EnterLiteralArray(ctx *LiteralArrayContext) {}

// ExitLiteralArray is called when production literalArray is exited.
func (s *BaseSomListener) ExitLiteralArray(ctx *LiteralArrayContext) {}

// EnterLiteralNumber is called when production literalNumber is entered.
func (s *BaseSomListener) EnterLiteralNumber(ctx *LiteralNumberContext) {}

// ExitLiteralNumber is called when production literalNumber is exited.
func (s *BaseSomListener) ExitLiteralNumber(ctx *LiteralNumberContext) {}

// EnterLiteralDecimal is called when production literalDecimal is entered.
func (s *BaseSomListener) EnterLiteralDecimal(ctx *LiteralDecimalContext) {}

// ExitLiteralDecimal is called when production literalDecimal is exited.
func (s *BaseSomListener) ExitLiteralDecimal(ctx *LiteralDecimalContext) {}

// EnterNegativeDecimal is called when production negativeDecimal is entered.
func (s *BaseSomListener) EnterNegativeDecimal(ctx *NegativeDecimalContext) {}

// ExitNegativeDecimal is called when production negativeDecimal is exited.
func (s *BaseSomListener) ExitNegativeDecimal(ctx *NegativeDecimalContext) {}

// EnterLiteralInteger is called when production literalInteger is entered.
func (s *BaseSomListener) EnterLiteralInteger(ctx *LiteralIntegerContext) {}

// ExitLiteralInteger is called when production literalInteger is exited.
func (s *BaseSomListener) ExitLiteralInteger(ctx *LiteralIntegerContext) {}

// EnterLiteralDouble is called when production literalDouble is entered.
func (s *BaseSomListener) EnterLiteralDouble(ctx *LiteralDoubleContext) {}

// ExitLiteralDouble is called when production literalDouble is exited.
func (s *BaseSomListener) ExitLiteralDouble(ctx *LiteralDoubleContext) {}

// EnterLiteralSymbol is called when production literalSymbol is entered.
func (s *BaseSomListener) EnterLiteralSymbol(ctx *LiteralSymbolContext) {}

// ExitLiteralSymbol is called when production literalSymbol is exited.
func (s *BaseSomListener) ExitLiteralSymbol(ctx *LiteralSymbolContext) {}

// EnterLiteralString is called when production literalString is entered.
func (s *BaseSomListener) EnterLiteralString(ctx *LiteralStringContext) {}

// ExitLiteralString is called when production literalString is exited.
func (s *BaseSomListener) ExitLiteralString(ctx *LiteralStringContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseSomListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseSomListener) ExitSelector(ctx *SelectorContext) {}

// EnterKeywordSelector is called when production keywordSelector is entered.
func (s *BaseSomListener) EnterKeywordSelector(ctx *KeywordSelectorContext) {}

// ExitKeywordSelector is called when production keywordSelector is exited.
func (s *BaseSomListener) ExitKeywordSelector(ctx *KeywordSelectorContext) {}

// EnterString is called when production string is entered.
func (s *BaseSomListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseSomListener) ExitString(ctx *StringContext) {}

// EnterNestedBlock is called when production nestedBlock is entered.
func (s *BaseSomListener) EnterNestedBlock(ctx *NestedBlockContext) {}

// ExitNestedBlock is called when production nestedBlock is exited.
func (s *BaseSomListener) ExitNestedBlock(ctx *NestedBlockContext) {}

// EnterBlockPattern is called when production blockPattern is entered.
func (s *BaseSomListener) EnterBlockPattern(ctx *BlockPatternContext) {}

// ExitBlockPattern is called when production blockPattern is exited.
func (s *BaseSomListener) ExitBlockPattern(ctx *BlockPatternContext) {}

// EnterBlockArguments is called when production blockArguments is entered.
func (s *BaseSomListener) EnterBlockArguments(ctx *BlockArgumentsContext) {}

// ExitBlockArguments is called when production blockArguments is exited.
func (s *BaseSomListener) ExitBlockArguments(ctx *BlockArgumentsContext) {}
