// Code generated from parser/Som.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Som

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SomParser struct {
	*antlr.BaseParser
}

var SomParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func somParserInit() {
	staticData := &SomParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "'primitive'", "", "'='", "", "'('", "')'", "'|'", "','",
		"'-'", "'~'", "'&'", "'*'", "'/'", "'\\'", "'+'", "'>'", "'<'", "'@'",
		"'%'", "", "':'", "'['", "']'", "'#'", "'^'", "'.'", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "Comment", "Whitespace", "Primitive", "Identifier", "Equal", "Separator",
		"NewTerm", "EndTerm", "Or", "Comma", "Minus", "Not", "And", "Star",
		"Div", "Mod", "Plus", "More", "Less", "At", "Per", "OperatorSequence",
		"Colon", "NewBlock", "EndBlock", "Pound", "Exit", "Period", "Assign",
		"Integer", "Double", "Keyword", "KeywordSequence", "STString",
	}
	staticData.RuleNames = []string{
		"classdef", "superclass", "instanceFields", "classFields", "method",
		"pattern", "unaryPattern", "binaryPattern", "keywordPattern", "methodBlock",
		"unarySelector", "binarySelector", "identifier", "keyword", "argument",
		"blockContents", "localDefs", "blockBody", "result", "expression", "assignation",
		"assignments", "assignment", "evaluation", "primary", "variable", "messages",
		"unaryMessage", "binaryMessage", "binaryOperand", "keywordMessage",
		"formula", "nestedTerm", "literal", "literalArray", "literalNumber",
		"literalDecimal", "negativeDecimal", "literalInteger", "literalDouble",
		"literalSymbol", "literalString", "selector", "keywordSelector", "string",
		"nestedBlock", "blockPattern", "blockArguments",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 359, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 102, 8, 0, 10, 0, 12, 0, 105,
		9, 0, 1, 0, 1, 0, 1, 0, 5, 0, 110, 8, 0, 10, 0, 12, 0, 113, 9, 0, 3, 0,
		115, 8, 0, 1, 0, 1, 0, 1, 1, 3, 1, 120, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5,
		2, 126, 8, 2, 10, 2, 12, 2, 129, 9, 2, 1, 2, 3, 2, 132, 8, 2, 1, 3, 1,
		3, 5, 3, 136, 8, 3, 10, 3, 12, 3, 139, 9, 3, 1, 3, 3, 3, 142, 8, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 148, 8, 4, 1, 5, 1, 5, 1, 5, 3, 5, 153, 8, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 4, 8, 163, 8, 8, 11, 8,
		12, 8, 164, 1, 9, 1, 9, 3, 9, 169, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 3, 15, 187, 8, 15, 1, 15, 1, 15, 1, 16, 5, 16, 192, 8, 16, 10, 16,
		12, 16, 195, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 202, 8, 17,
		3, 17, 204, 8, 17, 3, 17, 206, 8, 17, 1, 18, 1, 18, 3, 18, 210, 8, 18,
		1, 19, 1, 19, 3, 19, 214, 8, 19, 1, 20, 1, 20, 1, 20, 1, 21, 4, 21, 220,
		8, 21, 11, 21, 12, 21, 221, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 229,
		8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 235, 8, 24, 1, 25, 1, 25, 1,
		26, 4, 26, 240, 8, 26, 11, 26, 12, 26, 241, 1, 26, 5, 26, 245, 8, 26, 10,
		26, 12, 26, 248, 9, 26, 1, 26, 3, 26, 251, 8, 26, 1, 26, 4, 26, 254, 8,
		26, 11, 26, 12, 26, 255, 1, 26, 3, 26, 259, 8, 26, 1, 26, 3, 26, 262, 8,
		26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 5, 29, 271, 8, 29,
		10, 29, 12, 29, 274, 9, 29, 1, 30, 1, 30, 1, 30, 4, 30, 279, 8, 30, 11,
		30, 12, 30, 280, 1, 31, 1, 31, 5, 31, 285, 8, 31, 10, 31, 12, 31, 288,
		9, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 298,
		8, 33, 1, 34, 1, 34, 1, 34, 5, 34, 303, 8, 34, 10, 34, 12, 34, 306, 9,
		34, 1, 34, 1, 34, 1, 35, 1, 35, 3, 35, 312, 8, 35, 1, 36, 1, 36, 3, 36,
		316, 8, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 40, 3, 40, 328, 8, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 3, 42,
		335, 8, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 3, 45, 343, 8, 45,
		1, 45, 3, 45, 346, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 4, 47, 355, 8, 47, 11, 47, 12, 47, 356, 1, 47, 0, 0, 48, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 92, 94, 0, 3, 2, 0, 5, 5, 9, 22, 1, 0, 3, 4, 1,
		0, 32, 33, 357, 0, 96, 1, 0, 0, 0, 2, 119, 1, 0, 0, 0, 4, 131, 1, 0, 0,
		0, 6, 141, 1, 0, 0, 0, 8, 143, 1, 0, 0, 0, 10, 152, 1, 0, 0, 0, 12, 154,
		1, 0, 0, 0, 14, 156, 1, 0, 0, 0, 16, 162, 1, 0, 0, 0, 18, 166, 1, 0, 0,
		0, 20, 172, 1, 0, 0, 0, 22, 174, 1, 0, 0, 0, 24, 176, 1, 0, 0, 0, 26, 178,
		1, 0, 0, 0, 28, 180, 1, 0, 0, 0, 30, 186, 1, 0, 0, 0, 32, 193, 1, 0, 0,
		0, 34, 205, 1, 0, 0, 0, 36, 207, 1, 0, 0, 0, 38, 213, 1, 0, 0, 0, 40, 215,
		1, 0, 0, 0, 42, 219, 1, 0, 0, 0, 44, 223, 1, 0, 0, 0, 46, 226, 1, 0, 0,
		0, 48, 234, 1, 0, 0, 0, 50, 236, 1, 0, 0, 0, 52, 261, 1, 0, 0, 0, 54, 263,
		1, 0, 0, 0, 56, 265, 1, 0, 0, 0, 58, 268, 1, 0, 0, 0, 60, 278, 1, 0, 0,
		0, 62, 282, 1, 0, 0, 0, 64, 289, 1, 0, 0, 0, 66, 297, 1, 0, 0, 0, 68, 299,
		1, 0, 0, 0, 70, 311, 1, 0, 0, 0, 72, 315, 1, 0, 0, 0, 74, 317, 1, 0, 0,
		0, 76, 320, 1, 0, 0, 0, 78, 322, 1, 0, 0, 0, 80, 324, 1, 0, 0, 0, 82, 329,
		1, 0, 0, 0, 84, 334, 1, 0, 0, 0, 86, 336, 1, 0, 0, 0, 88, 338, 1, 0, 0,
		0, 90, 340, 1, 0, 0, 0, 92, 349, 1, 0, 0, 0, 94, 354, 1, 0, 0, 0, 96, 97,
		5, 4, 0, 0, 97, 98, 5, 5, 0, 0, 98, 99, 3, 2, 1, 0, 99, 103, 3, 4, 2, 0,
		100, 102, 3, 8, 4, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 114, 1, 0, 0, 0, 105, 103,
		1, 0, 0, 0, 106, 107, 5, 6, 0, 0, 107, 111, 3, 6, 3, 0, 108, 110, 3, 8,
		4, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114,
		106, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117,
		5, 8, 0, 0, 117, 1, 1, 0, 0, 0, 118, 120, 5, 4, 0, 0, 119, 118, 1, 0, 0,
		0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 5, 7, 0, 0, 122,
		3, 1, 0, 0, 0, 123, 127, 5, 9, 0, 0, 124, 126, 3, 50, 25, 0, 125, 124,
		1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 130, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 132, 5, 9, 0, 0,
		131, 123, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 5, 1, 0, 0, 0, 133, 137,
		5, 9, 0, 0, 134, 136, 3, 50, 25, 0, 135, 134, 1, 0, 0, 0, 136, 139, 1,
		0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0,
		0, 139, 137, 1, 0, 0, 0, 140, 142, 5, 9, 0, 0, 141, 133, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 7, 1, 0, 0, 0, 143, 144, 3, 10, 5, 0, 144, 147, 5,
		5, 0, 0, 145, 148, 5, 3, 0, 0, 146, 148, 3, 18, 9, 0, 147, 145, 1, 0, 0,
		0, 147, 146, 1, 0, 0, 0, 148, 9, 1, 0, 0, 0, 149, 153, 3, 12, 6, 0, 150,
		153, 3, 16, 8, 0, 151, 153, 3, 14, 7, 0, 152, 149, 1, 0, 0, 0, 152, 150,
		1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 11, 1, 0, 0, 0, 154, 155, 3, 20,
		10, 0, 155, 13, 1, 0, 0, 0, 156, 157, 3, 22, 11, 0, 157, 158, 3, 28, 14,
		0, 158, 15, 1, 0, 0, 0, 159, 160, 3, 26, 13, 0, 160, 161, 3, 28, 14, 0,
		161, 163, 1, 0, 0, 0, 162, 159, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 17, 1, 0, 0, 0, 166, 168, 5,
		7, 0, 0, 167, 169, 3, 30, 15, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0,
		0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 5, 8, 0, 0, 171, 19, 1, 0, 0, 0,
		172, 173, 3, 24, 12, 0, 173, 21, 1, 0, 0, 0, 174, 175, 7, 0, 0, 0, 175,
		23, 1, 0, 0, 0, 176, 177, 7, 1, 0, 0, 177, 25, 1, 0, 0, 0, 178, 179, 5,
		32, 0, 0, 179, 27, 1, 0, 0, 0, 180, 181, 3, 50, 25, 0, 181, 29, 1, 0, 0,
		0, 182, 183, 5, 9, 0, 0, 183, 184, 3, 32, 16, 0, 184, 185, 5, 9, 0, 0,
		185, 187, 1, 0, 0, 0, 186, 182, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		188, 1, 0, 0, 0, 188, 189, 3, 34, 17, 0, 189, 31, 1, 0, 0, 0, 190, 192,
		3, 50, 25, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1,
		0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 33, 1, 0, 0, 0, 195, 193, 1, 0, 0,
		0, 196, 197, 5, 27, 0, 0, 197, 206, 3, 36, 18, 0, 198, 203, 3, 38, 19,
		0, 199, 201, 5, 28, 0, 0, 200, 202, 3, 34, 17, 0, 201, 200, 1, 0, 0, 0,
		201, 202, 1, 0, 0, 0, 202, 204, 1, 0, 0, 0, 203, 199, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 196, 1, 0, 0, 0, 205, 198,
		1, 0, 0, 0, 206, 35, 1, 0, 0, 0, 207, 209, 3, 38, 19, 0, 208, 210, 5, 28,
		0, 0, 209, 208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 37, 1, 0, 0, 0,
		211, 214, 3, 40, 20, 0, 212, 214, 3, 46, 23, 0, 213, 211, 1, 0, 0, 0, 213,
		212, 1, 0, 0, 0, 214, 39, 1, 0, 0, 0, 215, 216, 3, 42, 21, 0, 216, 217,
		3, 46, 23, 0, 217, 41, 1, 0, 0, 0, 218, 220, 3, 44, 22, 0, 219, 218, 1,
		0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0,
		0, 222, 43, 1, 0, 0, 0, 223, 224, 3, 50, 25, 0, 224, 225, 5, 29, 0, 0,
		225, 45, 1, 0, 0, 0, 226, 228, 3, 48, 24, 0, 227, 229, 3, 52, 26, 0, 228,
		227, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 47, 1, 0, 0, 0, 230, 235, 3,
		50, 25, 0, 231, 235, 3, 64, 32, 0, 232, 235, 3, 90, 45, 0, 233, 235, 3,
		66, 33, 0, 234, 230, 1, 0, 0, 0, 234, 231, 1, 0, 0, 0, 234, 232, 1, 0,
		0, 0, 234, 233, 1, 0, 0, 0, 235, 49, 1, 0, 0, 0, 236, 237, 3, 24, 12, 0,
		237, 51, 1, 0, 0, 0, 238, 240, 3, 54, 27, 0, 239, 238, 1, 0, 0, 0, 240,
		241, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 246,
		1, 0, 0, 0, 243, 245, 3, 56, 28, 0, 244, 243, 1, 0, 0, 0, 245, 248, 1,
		0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 250, 1, 0, 0,
		0, 248, 246, 1, 0, 0, 0, 249, 251, 3, 60, 30, 0, 250, 249, 1, 0, 0, 0,
		250, 251, 1, 0, 0, 0, 251, 262, 1, 0, 0, 0, 252, 254, 3, 56, 28, 0, 253,
		252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256,
		1, 0, 0, 0, 256, 258, 1, 0, 0, 0, 257, 259, 3, 60, 30, 0, 258, 257, 1,
		0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 262, 3, 60, 30,
		0, 261, 239, 1, 0, 0, 0, 261, 253, 1, 0, 0, 0, 261, 260, 1, 0, 0, 0, 262,
		53, 1, 0, 0, 0, 263, 264, 3, 20, 10, 0, 264, 55, 1, 0, 0, 0, 265, 266,
		3, 22, 11, 0, 266, 267, 3, 58, 29, 0, 267, 57, 1, 0, 0, 0, 268, 272, 3,
		48, 24, 0, 269, 271, 3, 54, 27, 0, 270, 269, 1, 0, 0, 0, 271, 274, 1, 0,
		0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 59, 1, 0, 0, 0,
		274, 272, 1, 0, 0, 0, 275, 276, 3, 26, 13, 0, 276, 277, 3, 62, 31, 0, 277,
		279, 1, 0, 0, 0, 278, 275, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 278,
		1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 61, 1, 0, 0, 0, 282, 286, 3, 58,
		29, 0, 283, 285, 3, 56, 28, 0, 284, 283, 1, 0, 0, 0, 285, 288, 1, 0, 0,
		0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 63, 1, 0, 0, 0, 288,
		286, 1, 0, 0, 0, 289, 290, 5, 7, 0, 0, 290, 291, 3, 38, 19, 0, 291, 292,
		5, 8, 0, 0, 292, 65, 1, 0, 0, 0, 293, 298, 3, 68, 34, 0, 294, 298, 3, 80,
		40, 0, 295, 298, 3, 82, 41, 0, 296, 298, 3, 70, 35, 0, 297, 293, 1, 0,
		0, 0, 297, 294, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 296, 1, 0, 0, 0,
		298, 67, 1, 0, 0, 0, 299, 300, 5, 26, 0, 0, 300, 304, 5, 7, 0, 0, 301,
		303, 3, 66, 33, 0, 302, 301, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302,
		1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0, 306, 304, 1, 0,
		0, 0, 307, 308, 5, 8, 0, 0, 308, 69, 1, 0, 0, 0, 309, 312, 3, 74, 37, 0,
		310, 312, 3, 72, 36, 0, 311, 309, 1, 0, 0, 0, 311, 310, 1, 0, 0, 0, 312,
		71, 1, 0, 0, 0, 313, 316, 3, 76, 38, 0, 314, 316, 3, 78, 39, 0, 315, 313,
		1, 0, 0, 0, 315, 314, 1, 0, 0, 0, 316, 73, 1, 0, 0, 0, 317, 318, 5, 11,
		0, 0, 318, 319, 3, 72, 36, 0, 319, 75, 1, 0, 0, 0, 320, 321, 5, 30, 0,
		0, 321, 77, 1, 0, 0, 0, 322, 323, 5, 31, 0, 0, 323, 79, 1, 0, 0, 0, 324,
		327, 5, 26, 0, 0, 325, 328, 3, 88, 44, 0, 326, 328, 3, 84, 42, 0, 327,
		325, 1, 0, 0, 0, 327, 326, 1, 0, 0, 0, 328, 81, 1, 0, 0, 0, 329, 330, 3,
		88, 44, 0, 330, 83, 1, 0, 0, 0, 331, 335, 3, 22, 11, 0, 332, 335, 3, 86,
		43, 0, 333, 335, 3, 20, 10, 0, 334, 331, 1, 0, 0, 0, 334, 332, 1, 0, 0,
		0, 334, 333, 1, 0, 0, 0, 335, 85, 1, 0, 0, 0, 336, 337, 7, 2, 0, 0, 337,
		87, 1, 0, 0, 0, 338, 339, 5, 34, 0, 0, 339, 89, 1, 0, 0, 0, 340, 342, 5,
		24, 0, 0, 341, 343, 3, 92, 46, 0, 342, 341, 1, 0, 0, 0, 342, 343, 1, 0,
		0, 0, 343, 345, 1, 0, 0, 0, 344, 346, 3, 30, 15, 0, 345, 344, 1, 0, 0,
		0, 345, 346, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 348, 5, 25, 0, 0, 348,
		91, 1, 0, 0, 0, 349, 350, 3, 94, 47, 0, 350, 351, 5, 9, 0, 0, 351, 93,
		1, 0, 0, 0, 352, 353, 5, 23, 0, 0, 353, 355, 3, 28, 14, 0, 354, 352, 1,
		0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 356, 357, 1, 0, 0,
		0, 357, 95, 1, 0, 0, 0, 40, 103, 111, 114, 119, 127, 131, 137, 141, 147,
		152, 164, 168, 186, 193, 201, 203, 205, 209, 213, 221, 228, 234, 241, 246,
		250, 255, 258, 261, 272, 280, 286, 297, 304, 311, 315, 327, 334, 342, 345,
		356,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SomParserInit initializes any static state used to implement SomParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSomParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SomParserInit() {
	staticData := &SomParserStaticData
	staticData.once.Do(somParserInit)
}

// NewSomParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSomParser(input antlr.TokenStream) *SomParser {
	SomParserInit()
	this := new(SomParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SomParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Som.g4"

	return this
}

// SomParser tokens.
const (
	SomParserEOF              = antlr.TokenEOF
	SomParserComment          = 1
	SomParserWhitespace       = 2
	SomParserPrimitive        = 3
	SomParserIdentifier       = 4
	SomParserEqual            = 5
	SomParserSeparator        = 6
	SomParserNewTerm          = 7
	SomParserEndTerm          = 8
	SomParserOr               = 9
	SomParserComma            = 10
	SomParserMinus            = 11
	SomParserNot              = 12
	SomParserAnd              = 13
	SomParserStar             = 14
	SomParserDiv              = 15
	SomParserMod              = 16
	SomParserPlus             = 17
	SomParserMore             = 18
	SomParserLess             = 19
	SomParserAt               = 20
	SomParserPer              = 21
	SomParserOperatorSequence = 22
	SomParserColon            = 23
	SomParserNewBlock         = 24
	SomParserEndBlock         = 25
	SomParserPound            = 26
	SomParserExit             = 27
	SomParserPeriod           = 28
	SomParserAssign           = 29
	SomParserInteger          = 30
	SomParserDouble           = 31
	SomParserKeyword          = 32
	SomParserKeywordSequence  = 33
	SomParserSTString         = 34
)

// SomParser rules.
const (
	SomParserRULE_classdef        = 0
	SomParserRULE_superclass      = 1
	SomParserRULE_instanceFields  = 2
	SomParserRULE_classFields     = 3
	SomParserRULE_method          = 4
	SomParserRULE_pattern         = 5
	SomParserRULE_unaryPattern    = 6
	SomParserRULE_binaryPattern   = 7
	SomParserRULE_keywordPattern  = 8
	SomParserRULE_methodBlock     = 9
	SomParserRULE_unarySelector   = 10
	SomParserRULE_binarySelector  = 11
	SomParserRULE_identifier      = 12
	SomParserRULE_keyword         = 13
	SomParserRULE_argument        = 14
	SomParserRULE_blockContents   = 15
	SomParserRULE_localDefs       = 16
	SomParserRULE_blockBody       = 17
	SomParserRULE_result          = 18
	SomParserRULE_expression      = 19
	SomParserRULE_assignation     = 20
	SomParserRULE_assignments     = 21
	SomParserRULE_assignment      = 22
	SomParserRULE_evaluation      = 23
	SomParserRULE_primary         = 24
	SomParserRULE_variable        = 25
	SomParserRULE_messages        = 26
	SomParserRULE_unaryMessage    = 27
	SomParserRULE_binaryMessage   = 28
	SomParserRULE_binaryOperand   = 29
	SomParserRULE_keywordMessage  = 30
	SomParserRULE_formula         = 31
	SomParserRULE_nestedTerm      = 32
	SomParserRULE_literal         = 33
	SomParserRULE_literalArray    = 34
	SomParserRULE_literalNumber   = 35
	SomParserRULE_literalDecimal  = 36
	SomParserRULE_negativeDecimal = 37
	SomParserRULE_literalInteger  = 38
	SomParserRULE_literalDouble   = 39
	SomParserRULE_literalSymbol   = 40
	SomParserRULE_literalString   = 41
	SomParserRULE_selector        = 42
	SomParserRULE_keywordSelector = 43
	SomParserRULE_string          = 44
	SomParserRULE_nestedBlock     = 45
	SomParserRULE_blockPattern    = 46
	SomParserRULE_blockArguments  = 47
)

// IClassdefContext is an interface to support dynamic dispatch.
type IClassdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Equal() antlr.TerminalNode
	Superclass() ISuperclassContext
	InstanceFields() IInstanceFieldsContext
	EndTerm() antlr.TerminalNode
	AllMethod() []IMethodContext
	Method(i int) IMethodContext
	Separator() antlr.TerminalNode
	ClassFields() IClassFieldsContext

	// IsClassdefContext differentiates from other interfaces.
	IsClassdefContext()
}

type ClassdefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassdefContext() *ClassdefContext {
	var p = new(ClassdefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_classdef
	return p
}

func InitEmptyClassdefContext(p *ClassdefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_classdef
}

func (*ClassdefContext) IsClassdefContext() {}

func NewClassdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassdefContext {
	var p = new(ClassdefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_classdef

	return p
}

func (s *ClassdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassdefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SomParserIdentifier, 0)
}

func (s *ClassdefContext) Equal() antlr.TerminalNode {
	return s.GetToken(SomParserEqual, 0)
}

func (s *ClassdefContext) Superclass() ISuperclassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuperclassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuperclassContext)
}

func (s *ClassdefContext) InstanceFields() IInstanceFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstanceFieldsContext)
}

func (s *ClassdefContext) EndTerm() antlr.TerminalNode {
	return s.GetToken(SomParserEndTerm, 0)
}

func (s *ClassdefContext) AllMethod() []IMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodContext); ok {
			len++
		}
	}

	tst := make([]IMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodContext); ok {
			tst[i] = t.(IMethodContext)
			i++
		}
	}

	return tst
}

func (s *ClassdefContext) Method(i int) IMethodContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *ClassdefContext) Separator() antlr.TerminalNode {
	return s.GetToken(SomParserSeparator, 0)
}

func (s *ClassdefContext) ClassFields() IClassFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassFieldsContext)
}

func (s *ClassdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterClassdef(s)
	}
}

func (s *ClassdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitClassdef(s)
	}
}

func (p *SomParser) Classdef() (localctx IClassdefContext) {
	localctx = NewClassdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SomParserRULE_classdef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(SomParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(SomParserEqual)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Superclass()
	}
	{
		p.SetState(99)
		p.InstanceFields()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4303355448) != 0 {
		{
			p.SetState(100)
			p.Method()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SomParserSeparator {
		{
			p.SetState(106)
			p.Match(SomParserSeparator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.ClassFields()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4303355448) != 0 {
			{
				p.SetState(108)
				p.Method()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(116)
		p.Match(SomParserEndTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISuperclassContext is an interface to support dynamic dispatch.
type ISuperclassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NewTerm() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsSuperclassContext differentiates from other interfaces.
	IsSuperclassContext()
}

type SuperclassContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuperclassContext() *SuperclassContext {
	var p = new(SuperclassContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_superclass
	return p
}

func InitEmptySuperclassContext(p *SuperclassContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_superclass
}

func (*SuperclassContext) IsSuperclassContext() {}

func NewSuperclassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuperclassContext {
	var p = new(SuperclassContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_superclass

	return p
}

func (s *SuperclassContext) GetParser() antlr.Parser { return s.parser }

func (s *SuperclassContext) NewTerm() antlr.TerminalNode {
	return s.GetToken(SomParserNewTerm, 0)
}

func (s *SuperclassContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SomParserIdentifier, 0)
}

func (s *SuperclassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperclassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuperclassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterSuperclass(s)
	}
}

func (s *SuperclassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitSuperclass(s)
	}
}

func (p *SomParser) Superclass() (localctx ISuperclassContext) {
	localctx = NewSuperclassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SomParserRULE_superclass)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SomParserIdentifier {
		{
			p.SetState(118)
			p.Match(SomParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(121)
		p.Match(SomParserNewTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstanceFieldsContext is an interface to support dynamic dispatch.
type IInstanceFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOr() []antlr.TerminalNode
	Or(i int) antlr.TerminalNode
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext

	// IsInstanceFieldsContext differentiates from other interfaces.
	IsInstanceFieldsContext()
}

type InstanceFieldsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceFieldsContext() *InstanceFieldsContext {
	var p = new(InstanceFieldsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_instanceFields
	return p
}

func InitEmptyInstanceFieldsContext(p *InstanceFieldsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_instanceFields
}

func (*InstanceFieldsContext) IsInstanceFieldsContext() {}

func NewInstanceFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceFieldsContext {
	var p = new(InstanceFieldsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_instanceFields

	return p
}

func (s *InstanceFieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceFieldsContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(SomParserOr)
}

func (s *InstanceFieldsContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(SomParserOr, i)
}

func (s *InstanceFieldsContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *InstanceFieldsContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *InstanceFieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceFieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceFieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterInstanceFields(s)
	}
}

func (s *InstanceFieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitInstanceFields(s)
	}
}

func (p *SomParser) InstanceFields() (localctx IInstanceFieldsContext) {
	localctx = NewInstanceFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SomParserRULE_instanceFields)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(123)
			p.Match(SomParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SomParserPrimitive || _la == SomParserIdentifier {
			{
				p.SetState(124)
				p.Variable()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(130)
			p.Match(SomParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassFieldsContext is an interface to support dynamic dispatch.
type IClassFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOr() []antlr.TerminalNode
	Or(i int) antlr.TerminalNode
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext

	// IsClassFieldsContext differentiates from other interfaces.
	IsClassFieldsContext()
}

type ClassFieldsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassFieldsContext() *ClassFieldsContext {
	var p = new(ClassFieldsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_classFields
	return p
}

func InitEmptyClassFieldsContext(p *ClassFieldsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_classFields
}

func (*ClassFieldsContext) IsClassFieldsContext() {}

func NewClassFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassFieldsContext {
	var p = new(ClassFieldsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_classFields

	return p
}

func (s *ClassFieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassFieldsContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(SomParserOr)
}

func (s *ClassFieldsContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(SomParserOr, i)
}

func (s *ClassFieldsContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *ClassFieldsContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ClassFieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassFieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassFieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterClassFields(s)
	}
}

func (s *ClassFieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitClassFields(s)
	}
}

func (p *SomParser) ClassFields() (localctx IClassFieldsContext) {
	localctx = NewClassFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SomParserRULE_classFields)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(133)
			p.Match(SomParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SomParserPrimitive || _la == SomParserIdentifier {
			{
				p.SetState(134)
				p.Variable()
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(140)
			p.Match(SomParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pattern() IPatternContext
	Equal() antlr.TerminalNode
	Primitive() antlr.TerminalNode
	MethodBlock() IMethodBlockContext

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_method
	return p
}

func InitEmptyMethodContext(p *MethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_method
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *MethodContext) Equal() antlr.TerminalNode {
	return s.GetToken(SomParserEqual, 0)
}

func (s *MethodContext) Primitive() antlr.TerminalNode {
	return s.GetToken(SomParserPrimitive, 0)
}

func (s *MethodContext) MethodBlock() IMethodBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodBlockContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *SomParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SomParserRULE_method)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Pattern()
	}
	{
		p.SetState(144)
		p.Match(SomParserEqual)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserPrimitive:
		{
			p.SetState(145)
			p.Match(SomParserPrimitive)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SomParserNewTerm:
		{
			p.SetState(146)
			p.MethodBlock()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryPattern() IUnaryPatternContext
	KeywordPattern() IKeywordPatternContext
	BinaryPattern() IBinaryPatternContext

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_pattern
	return p
}

func InitEmptyPatternContext(p *PatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_pattern
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) UnaryPattern() IUnaryPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryPatternContext)
}

func (s *PatternContext) KeywordPattern() IKeywordPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordPatternContext)
}

func (s *PatternContext) BinaryPattern() IBinaryPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryPatternContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *SomParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SomParserRULE_pattern)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserPrimitive, SomParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.UnaryPattern()
		}

	case SomParserKeyword:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.KeywordPattern()
		}

	case SomParserEqual, SomParserOr, SomParserComma, SomParserMinus, SomParserNot, SomParserAnd, SomParserStar, SomParserDiv, SomParserMod, SomParserPlus, SomParserMore, SomParserLess, SomParserAt, SomParserPer, SomParserOperatorSequence:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.BinaryPattern()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryPatternContext is an interface to support dynamic dispatch.
type IUnaryPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnarySelector() IUnarySelectorContext

	// IsUnaryPatternContext differentiates from other interfaces.
	IsUnaryPatternContext()
}

type UnaryPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryPatternContext() *UnaryPatternContext {
	var p = new(UnaryPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_unaryPattern
	return p
}

func InitEmptyUnaryPatternContext(p *UnaryPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_unaryPattern
}

func (*UnaryPatternContext) IsUnaryPatternContext() {}

func NewUnaryPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryPatternContext {
	var p = new(UnaryPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_unaryPattern

	return p
}

func (s *UnaryPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryPatternContext) UnarySelector() IUnarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarySelectorContext)
}

func (s *UnaryPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterUnaryPattern(s)
	}
}

func (s *UnaryPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitUnaryPattern(s)
	}
}

func (p *SomParser) UnaryPattern() (localctx IUnaryPatternContext) {
	localctx = NewUnaryPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SomParserRULE_unaryPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.UnarySelector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryPatternContext is an interface to support dynamic dispatch.
type IBinaryPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinarySelector() IBinarySelectorContext
	Argument() IArgumentContext

	// IsBinaryPatternContext differentiates from other interfaces.
	IsBinaryPatternContext()
}

type BinaryPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryPatternContext() *BinaryPatternContext {
	var p = new(BinaryPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binaryPattern
	return p
}

func InitEmptyBinaryPatternContext(p *BinaryPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binaryPattern
}

func (*BinaryPatternContext) IsBinaryPatternContext() {}

func NewBinaryPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryPatternContext {
	var p = new(BinaryPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_binaryPattern

	return p
}

func (s *BinaryPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryPatternContext) BinarySelector() IBinarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinarySelectorContext)
}

func (s *BinaryPatternContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *BinaryPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBinaryPattern(s)
	}
}

func (s *BinaryPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBinaryPattern(s)
	}
}

func (p *SomParser) BinaryPattern() (localctx IBinaryPatternContext) {
	localctx = NewBinaryPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SomParserRULE_binaryPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.BinarySelector()
	}
	{
		p.SetState(157)
		p.Argument()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordPatternContext is an interface to support dynamic dispatch.
type IKeywordPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKeyword() []IKeywordContext
	Keyword(i int) IKeywordContext
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext

	// IsKeywordPatternContext differentiates from other interfaces.
	IsKeywordPatternContext()
}

type KeywordPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordPatternContext() *KeywordPatternContext {
	var p = new(KeywordPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keywordPattern
	return p
}

func InitEmptyKeywordPatternContext(p *KeywordPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keywordPattern
}

func (*KeywordPatternContext) IsKeywordPatternContext() {}

func NewKeywordPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordPatternContext {
	var p = new(KeywordPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_keywordPattern

	return p
}

func (s *KeywordPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordPatternContext) AllKeyword() []IKeywordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeywordContext); ok {
			len++
		}
	}

	tst := make([]IKeywordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeywordContext); ok {
			tst[i] = t.(IKeywordContext)
			i++
		}
	}

	return tst
}

func (s *KeywordPatternContext) Keyword(i int) IKeywordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *KeywordPatternContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *KeywordPatternContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *KeywordPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterKeywordPattern(s)
	}
}

func (s *KeywordPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitKeywordPattern(s)
	}
}

func (p *SomParser) KeywordPattern() (localctx IKeywordPatternContext) {
	localctx = NewKeywordPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SomParserRULE_keywordPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SomParserKeyword {
		{
			p.SetState(159)
			p.Keyword()
		}
		{
			p.SetState(160)
			p.Argument()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodBlockContext is an interface to support dynamic dispatch.
type IMethodBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NewTerm() antlr.TerminalNode
	EndTerm() antlr.TerminalNode
	BlockContents() IBlockContentsContext

	// IsMethodBlockContext differentiates from other interfaces.
	IsMethodBlockContext()
}

type MethodBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodBlockContext() *MethodBlockContext {
	var p = new(MethodBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_methodBlock
	return p
}

func InitEmptyMethodBlockContext(p *MethodBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_methodBlock
}

func (*MethodBlockContext) IsMethodBlockContext() {}

func NewMethodBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodBlockContext {
	var p = new(MethodBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_methodBlock

	return p
}

func (s *MethodBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodBlockContext) NewTerm() antlr.TerminalNode {
	return s.GetToken(SomParserNewTerm, 0)
}

func (s *MethodBlockContext) EndTerm() antlr.TerminalNode {
	return s.GetToken(SomParserEndTerm, 0)
}

func (s *MethodBlockContext) BlockContents() IBlockContentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContentsContext)
}

func (s *MethodBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterMethodBlock(s)
	}
}

func (s *MethodBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitMethodBlock(s)
	}
}

func (p *SomParser) MethodBlock() (localctx IMethodBlockContext) {
	localctx = NewMethodBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SomParserRULE_methodBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(SomParserNewTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20619201176) != 0 {
		{
			p.SetState(167)
			p.BlockContents()
		}

	}
	{
		p.SetState(170)
		p.Match(SomParserEndTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnarySelectorContext is an interface to support dynamic dispatch.
type IUnarySelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsUnarySelectorContext differentiates from other interfaces.
	IsUnarySelectorContext()
}

type UnarySelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnarySelectorContext() *UnarySelectorContext {
	var p = new(UnarySelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_unarySelector
	return p
}

func InitEmptyUnarySelectorContext(p *UnarySelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_unarySelector
}

func (*UnarySelectorContext) IsUnarySelectorContext() {}

func NewUnarySelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnarySelectorContext {
	var p = new(UnarySelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_unarySelector

	return p
}

func (s *UnarySelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnarySelectorContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *UnarySelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarySelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnarySelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterUnarySelector(s)
	}
}

func (s *UnarySelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitUnarySelector(s)
	}
}

func (p *SomParser) UnarySelector() (localctx IUnarySelectorContext) {
	localctx = NewUnarySelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SomParserRULE_unarySelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinarySelectorContext is an interface to support dynamic dispatch.
type IBinarySelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Or() antlr.TerminalNode
	Comma() antlr.TerminalNode
	Minus() antlr.TerminalNode
	Equal() antlr.TerminalNode
	Not() antlr.TerminalNode
	And() antlr.TerminalNode
	Star() antlr.TerminalNode
	Div() antlr.TerminalNode
	Mod() antlr.TerminalNode
	Plus() antlr.TerminalNode
	More() antlr.TerminalNode
	Less() antlr.TerminalNode
	At() antlr.TerminalNode
	Per() antlr.TerminalNode
	OperatorSequence() antlr.TerminalNode

	// IsBinarySelectorContext differentiates from other interfaces.
	IsBinarySelectorContext()
}

type BinarySelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinarySelectorContext() *BinarySelectorContext {
	var p = new(BinarySelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binarySelector
	return p
}

func InitEmptyBinarySelectorContext(p *BinarySelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binarySelector
}

func (*BinarySelectorContext) IsBinarySelectorContext() {}

func NewBinarySelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinarySelectorContext {
	var p = new(BinarySelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_binarySelector

	return p
}

func (s *BinarySelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinarySelectorContext) Or() antlr.TerminalNode {
	return s.GetToken(SomParserOr, 0)
}

func (s *BinarySelectorContext) Comma() antlr.TerminalNode {
	return s.GetToken(SomParserComma, 0)
}

func (s *BinarySelectorContext) Minus() antlr.TerminalNode {
	return s.GetToken(SomParserMinus, 0)
}

func (s *BinarySelectorContext) Equal() antlr.TerminalNode {
	return s.GetToken(SomParserEqual, 0)
}

func (s *BinarySelectorContext) Not() antlr.TerminalNode {
	return s.GetToken(SomParserNot, 0)
}

func (s *BinarySelectorContext) And() antlr.TerminalNode {
	return s.GetToken(SomParserAnd, 0)
}

func (s *BinarySelectorContext) Star() antlr.TerminalNode {
	return s.GetToken(SomParserStar, 0)
}

func (s *BinarySelectorContext) Div() antlr.TerminalNode {
	return s.GetToken(SomParserDiv, 0)
}

func (s *BinarySelectorContext) Mod() antlr.TerminalNode {
	return s.GetToken(SomParserMod, 0)
}

func (s *BinarySelectorContext) Plus() antlr.TerminalNode {
	return s.GetToken(SomParserPlus, 0)
}

func (s *BinarySelectorContext) More() antlr.TerminalNode {
	return s.GetToken(SomParserMore, 0)
}

func (s *BinarySelectorContext) Less() antlr.TerminalNode {
	return s.GetToken(SomParserLess, 0)
}

func (s *BinarySelectorContext) At() antlr.TerminalNode {
	return s.GetToken(SomParserAt, 0)
}

func (s *BinarySelectorContext) Per() antlr.TerminalNode {
	return s.GetToken(SomParserPer, 0)
}

func (s *BinarySelectorContext) OperatorSequence() antlr.TerminalNode {
	return s.GetToken(SomParserOperatorSequence, 0)
}

func (s *BinarySelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinarySelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinarySelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBinarySelector(s)
	}
}

func (s *BinarySelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBinarySelector(s)
	}
}

func (p *SomParser) BinarySelector() (localctx IBinarySelectorContext) {
	localctx = NewBinarySelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SomParserRULE_binarySelector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8388128) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primitive() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Primitive() antlr.TerminalNode {
	return s.GetToken(SomParserPrimitive, 0)
}

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SomParserIdentifier, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *SomParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SomParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SomParserPrimitive || _la == SomParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword() antlr.TerminalNode

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Keyword() antlr.TerminalNode {
	return s.GetToken(SomParserKeyword, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *SomParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SomParserRULE_keyword)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(SomParserKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *SomParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SomParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Variable()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContentsContext is an interface to support dynamic dispatch.
type IBlockContentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BlockBody() IBlockBodyContext
	AllOr() []antlr.TerminalNode
	Or(i int) antlr.TerminalNode
	LocalDefs() ILocalDefsContext

	// IsBlockContentsContext differentiates from other interfaces.
	IsBlockContentsContext()
}

type BlockContentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContentsContext() *BlockContentsContext {
	var p = new(BlockContentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockContents
	return p
}

func InitEmptyBlockContentsContext(p *BlockContentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockContents
}

func (*BlockContentsContext) IsBlockContentsContext() {}

func NewBlockContentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContentsContext {
	var p = new(BlockContentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_blockContents

	return p
}

func (s *BlockContentsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContentsContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *BlockContentsContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(SomParserOr)
}

func (s *BlockContentsContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(SomParserOr, i)
}

func (s *BlockContentsContext) LocalDefs() ILocalDefsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalDefsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalDefsContext)
}

func (s *BlockContentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBlockContents(s)
	}
}

func (s *BlockContentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBlockContents(s)
	}
}

func (p *SomParser) BlockContents() (localctx IBlockContentsContext) {
	localctx = NewBlockContentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SomParserRULE_blockContents)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SomParserOr {
		{
			p.SetState(182)
			p.Match(SomParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.LocalDefs()
		}
		{
			p.SetState(184)
			p.Match(SomParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(188)
		p.BlockBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILocalDefsContext is an interface to support dynamic dispatch.
type ILocalDefsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext

	// IsLocalDefsContext differentiates from other interfaces.
	IsLocalDefsContext()
}

type LocalDefsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalDefsContext() *LocalDefsContext {
	var p = new(LocalDefsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_localDefs
	return p
}

func InitEmptyLocalDefsContext(p *LocalDefsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_localDefs
}

func (*LocalDefsContext) IsLocalDefsContext() {}

func NewLocalDefsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalDefsContext {
	var p = new(LocalDefsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_localDefs

	return p
}

func (s *LocalDefsContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalDefsContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *LocalDefsContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *LocalDefsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalDefsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalDefsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLocalDefs(s)
	}
}

func (s *LocalDefsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLocalDefs(s)
	}
}

func (p *SomParser) LocalDefs() (localctx ILocalDefsContext) {
	localctx = NewLocalDefsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SomParserRULE_localDefs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SomParserPrimitive || _la == SomParserIdentifier {
		{
			p.SetState(190)
			p.Variable()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockBodyContext is an interface to support dynamic dispatch.
type IBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exit() antlr.TerminalNode
	Result() IResultContext
	Expression() IExpressionContext
	Period() antlr.TerminalNode
	BlockBody() IBlockBodyContext

	// IsBlockBodyContext differentiates from other interfaces.
	IsBlockBodyContext()
}

type BlockBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockBodyContext() *BlockBodyContext {
	var p = new(BlockBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockBody
	return p
}

func InitEmptyBlockBodyContext(p *BlockBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockBody
}

func (*BlockBodyContext) IsBlockBodyContext() {}

func NewBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockBodyContext {
	var p = new(BlockBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_blockBody

	return p
}

func (s *BlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockBodyContext) Exit() antlr.TerminalNode {
	return s.GetToken(SomParserExit, 0)
}

func (s *BlockBodyContext) Result() IResultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultContext)
}

func (s *BlockBodyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BlockBodyContext) Period() antlr.TerminalNode {
	return s.GetToken(SomParserPeriod, 0)
}

func (s *BlockBodyContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *BlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBlockBody(s)
	}
}

func (s *BlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBlockBody(s)
	}
}

func (p *SomParser) BlockBody() (localctx IBlockBodyContext) {
	localctx = NewBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SomParserRULE_blockBody)
	var _la int

	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserExit:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.Match(SomParserExit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Result()
		}

	case SomParserPrimitive, SomParserIdentifier, SomParserNewTerm, SomParserMinus, SomParserNewBlock, SomParserPound, SomParserInteger, SomParserDouble, SomParserSTString:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Expression()
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SomParserPeriod {
			{
				p.SetState(199)
				p.Match(SomParserPeriod)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20619200664) != 0 {
				{
					p.SetState(200)
					p.BlockBody()
				}

			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResultContext is an interface to support dynamic dispatch.
type IResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Period() antlr.TerminalNode

	// IsResultContext differentiates from other interfaces.
	IsResultContext()
}

type ResultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultContext() *ResultContext {
	var p = new(ResultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_result
	return p
}

func InitEmptyResultContext(p *ResultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_result
}

func (*ResultContext) IsResultContext() {}

func NewResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultContext {
	var p = new(ResultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_result

	return p
}

func (s *ResultContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ResultContext) Period() antlr.TerminalNode {
	return s.GetToken(SomParserPeriod, 0)
}

func (s *ResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterResult(s)
	}
}

func (s *ResultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitResult(s)
	}
}

func (p *SomParser) Result() (localctx IResultContext) {
	localctx = NewResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SomParserRULE_result)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Expression()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SomParserPeriod {
		{
			p.SetState(208)
			p.Match(SomParserPeriod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignation() IAssignationContext
	Evaluation() IEvaluationContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Assignation() IAssignationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignationContext)
}

func (s *ExpressionContext) Evaluation() IEvaluationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEvaluationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEvaluationContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SomParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SomParserRULE_expression)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Assignation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Evaluation()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignationContext is an interface to support dynamic dispatch.
type IAssignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignments() IAssignmentsContext
	Evaluation() IEvaluationContext

	// IsAssignationContext differentiates from other interfaces.
	IsAssignationContext()
}

type AssignationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignationContext() *AssignationContext {
	var p = new(AssignationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_assignation
	return p
}

func InitEmptyAssignationContext(p *AssignationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_assignation
}

func (*AssignationContext) IsAssignationContext() {}

func NewAssignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignationContext {
	var p = new(AssignationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_assignation

	return p
}

func (s *AssignationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignationContext) Assignments() IAssignmentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentsContext)
}

func (s *AssignationContext) Evaluation() IEvaluationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEvaluationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEvaluationContext)
}

func (s *AssignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterAssignation(s)
	}
}

func (s *AssignationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitAssignation(s)
	}
}

func (p *SomParser) Assignation() (localctx IAssignationContext) {
	localctx = NewAssignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SomParserRULE_assignation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Assignments()
	}
	{
		p.SetState(216)
		p.Evaluation()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentsContext is an interface to support dynamic dispatch.
type IAssignmentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext

	// IsAssignmentsContext differentiates from other interfaces.
	IsAssignmentsContext()
}

type AssignmentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentsContext() *AssignmentsContext {
	var p = new(AssignmentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_assignments
	return p
}

func InitEmptyAssignmentsContext(p *AssignmentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_assignments
}

func (*AssignmentsContext) IsAssignmentsContext() {}

func NewAssignmentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentsContext {
	var p = new(AssignmentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_assignments

	return p
}

func (s *AssignmentsContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentsContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentsContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignmentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterAssignments(s)
	}
}

func (s *AssignmentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitAssignments(s)
	}
}

func (p *SomParser) Assignments() (localctx IAssignmentsContext) {
	localctx = NewAssignmentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SomParserRULE_assignments)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(218)
				p.Assignment()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	Assign() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) Assign() antlr.TerminalNode {
	return s.GetToken(SomParserAssign, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *SomParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SomParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Variable()
	}
	{
		p.SetState(224)
		p.Match(SomParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEvaluationContext is an interface to support dynamic dispatch.
type IEvaluationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	Messages() IMessagesContext

	// IsEvaluationContext differentiates from other interfaces.
	IsEvaluationContext()
}

type EvaluationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEvaluationContext() *EvaluationContext {
	var p = new(EvaluationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_evaluation
	return p
}

func InitEmptyEvaluationContext(p *EvaluationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_evaluation
}

func (*EvaluationContext) IsEvaluationContext() {}

func NewEvaluationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvaluationContext {
	var p = new(EvaluationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_evaluation

	return p
}

func (s *EvaluationContext) GetParser() antlr.Parser { return s.parser }

func (s *EvaluationContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *EvaluationContext) Messages() IMessagesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessagesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessagesContext)
}

func (s *EvaluationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvaluationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvaluationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterEvaluation(s)
	}
}

func (s *EvaluationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitEvaluation(s)
	}
}

func (p *SomParser) Evaluation() (localctx IEvaluationContext) {
	localctx = NewEvaluationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SomParserRULE_evaluation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Primary()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4303355448) != 0 {
		{
			p.SetState(227)
			p.Messages()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	NestedTerm() INestedTermContext
	NestedBlock() INestedBlockContext
	Literal() ILiteralContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *PrimaryContext) NestedTerm() INestedTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedTermContext)
}

func (s *PrimaryContext) NestedBlock() INestedBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedBlockContext)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *SomParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SomParserRULE_primary)
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserPrimitive, SomParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Variable()
		}

	case SomParserNewTerm:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.NestedTerm()
		}

	case SomParserNewBlock:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(232)
			p.NestedBlock()
		}

	case SomParserMinus, SomParserPound, SomParserInteger, SomParserDouble, SomParserSTString:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(233)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *SomParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SomParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessagesContext is an interface to support dynamic dispatch.
type IMessagesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryMessage() []IUnaryMessageContext
	UnaryMessage(i int) IUnaryMessageContext
	AllBinaryMessage() []IBinaryMessageContext
	BinaryMessage(i int) IBinaryMessageContext
	KeywordMessage() IKeywordMessageContext

	// IsMessagesContext differentiates from other interfaces.
	IsMessagesContext()
}

type MessagesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessagesContext() *MessagesContext {
	var p = new(MessagesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_messages
	return p
}

func InitEmptyMessagesContext(p *MessagesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_messages
}

func (*MessagesContext) IsMessagesContext() {}

func NewMessagesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagesContext {
	var p = new(MessagesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_messages

	return p
}

func (s *MessagesContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagesContext) AllUnaryMessage() []IUnaryMessageContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryMessageContext); ok {
			len++
		}
	}

	tst := make([]IUnaryMessageContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryMessageContext); ok {
			tst[i] = t.(IUnaryMessageContext)
			i++
		}
	}

	return tst
}

func (s *MessagesContext) UnaryMessage(i int) IUnaryMessageContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryMessageContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryMessageContext)
}

func (s *MessagesContext) AllBinaryMessage() []IBinaryMessageContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinaryMessageContext); ok {
			len++
		}
	}

	tst := make([]IBinaryMessageContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinaryMessageContext); ok {
			tst[i] = t.(IBinaryMessageContext)
			i++
		}
	}

	return tst
}

func (s *MessagesContext) BinaryMessage(i int) IBinaryMessageContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryMessageContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryMessageContext)
}

func (s *MessagesContext) KeywordMessage() IKeywordMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordMessageContext)
}

func (s *MessagesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterMessages(s)
	}
}

func (s *MessagesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitMessages(s)
	}
}

func (p *SomParser) Messages() (localctx IMessagesContext) {
	localctx = NewMessagesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SomParserRULE_messages)
	var _la int

	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserPrimitive, SomParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SomParserPrimitive || _la == SomParserIdentifier {
			{
				p.SetState(238)
				p.UnaryMessage()
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8388128) != 0 {
			{
				p.SetState(243)
				p.BinaryMessage()
			}

			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SomParserKeyword {
			{
				p.SetState(249)
				p.KeywordMessage()
			}

		}

	case SomParserEqual, SomParserOr, SomParserComma, SomParserMinus, SomParserNot, SomParserAnd, SomParserStar, SomParserDiv, SomParserMod, SomParserPlus, SomParserMore, SomParserLess, SomParserAt, SomParserPer, SomParserOperatorSequence:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8388128) != 0) {
			{
				p.SetState(252)
				p.BinaryMessage()
			}

			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SomParserKeyword {
			{
				p.SetState(257)
				p.KeywordMessage()
			}

		}

	case SomParserKeyword:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.KeywordMessage()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryMessageContext is an interface to support dynamic dispatch.
type IUnaryMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnarySelector() IUnarySelectorContext

	// IsUnaryMessageContext differentiates from other interfaces.
	IsUnaryMessageContext()
}

type UnaryMessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryMessageContext() *UnaryMessageContext {
	var p = new(UnaryMessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_unaryMessage
	return p
}

func InitEmptyUnaryMessageContext(p *UnaryMessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_unaryMessage
}

func (*UnaryMessageContext) IsUnaryMessageContext() {}

func NewUnaryMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryMessageContext {
	var p = new(UnaryMessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_unaryMessage

	return p
}

func (s *UnaryMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryMessageContext) UnarySelector() IUnarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarySelectorContext)
}

func (s *UnaryMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterUnaryMessage(s)
	}
}

func (s *UnaryMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitUnaryMessage(s)
	}
}

func (p *SomParser) UnaryMessage() (localctx IUnaryMessageContext) {
	localctx = NewUnaryMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SomParserRULE_unaryMessage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.UnarySelector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryMessageContext is an interface to support dynamic dispatch.
type IBinaryMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinarySelector() IBinarySelectorContext
	BinaryOperand() IBinaryOperandContext

	// IsBinaryMessageContext differentiates from other interfaces.
	IsBinaryMessageContext()
}

type BinaryMessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryMessageContext() *BinaryMessageContext {
	var p = new(BinaryMessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binaryMessage
	return p
}

func InitEmptyBinaryMessageContext(p *BinaryMessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binaryMessage
}

func (*BinaryMessageContext) IsBinaryMessageContext() {}

func NewBinaryMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryMessageContext {
	var p = new(BinaryMessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_binaryMessage

	return p
}

func (s *BinaryMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryMessageContext) BinarySelector() IBinarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinarySelectorContext)
}

func (s *BinaryMessageContext) BinaryOperand() IBinaryOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperandContext)
}

func (s *BinaryMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBinaryMessage(s)
	}
}

func (s *BinaryMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBinaryMessage(s)
	}
}

func (p *SomParser) BinaryMessage() (localctx IBinaryMessageContext) {
	localctx = NewBinaryMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SomParserRULE_binaryMessage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.BinarySelector()
	}
	{
		p.SetState(266)
		p.BinaryOperand()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOperandContext is an interface to support dynamic dispatch.
type IBinaryOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	AllUnaryMessage() []IUnaryMessageContext
	UnaryMessage(i int) IUnaryMessageContext

	// IsBinaryOperandContext differentiates from other interfaces.
	IsBinaryOperandContext()
}

type BinaryOperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperandContext() *BinaryOperandContext {
	var p = new(BinaryOperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binaryOperand
	return p
}

func InitEmptyBinaryOperandContext(p *BinaryOperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_binaryOperand
}

func (*BinaryOperandContext) IsBinaryOperandContext() {}

func NewBinaryOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperandContext {
	var p = new(BinaryOperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_binaryOperand

	return p
}

func (s *BinaryOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperandContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *BinaryOperandContext) AllUnaryMessage() []IUnaryMessageContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryMessageContext); ok {
			len++
		}
	}

	tst := make([]IUnaryMessageContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryMessageContext); ok {
			tst[i] = t.(IUnaryMessageContext)
			i++
		}
	}

	return tst
}

func (s *BinaryOperandContext) UnaryMessage(i int) IUnaryMessageContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryMessageContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryMessageContext)
}

func (s *BinaryOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBinaryOperand(s)
	}
}

func (s *BinaryOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBinaryOperand(s)
	}
}

func (p *SomParser) BinaryOperand() (localctx IBinaryOperandContext) {
	localctx = NewBinaryOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SomParserRULE_binaryOperand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Primary()
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SomParserPrimitive || _la == SomParserIdentifier {
		{
			p.SetState(269)
			p.UnaryMessage()
		}

		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordMessageContext is an interface to support dynamic dispatch.
type IKeywordMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKeyword() []IKeywordContext
	Keyword(i int) IKeywordContext
	AllFormula() []IFormulaContext
	Formula(i int) IFormulaContext

	// IsKeywordMessageContext differentiates from other interfaces.
	IsKeywordMessageContext()
}

type KeywordMessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordMessageContext() *KeywordMessageContext {
	var p = new(KeywordMessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keywordMessage
	return p
}

func InitEmptyKeywordMessageContext(p *KeywordMessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keywordMessage
}

func (*KeywordMessageContext) IsKeywordMessageContext() {}

func NewKeywordMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordMessageContext {
	var p = new(KeywordMessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_keywordMessage

	return p
}

func (s *KeywordMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordMessageContext) AllKeyword() []IKeywordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeywordContext); ok {
			len++
		}
	}

	tst := make([]IKeywordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeywordContext); ok {
			tst[i] = t.(IKeywordContext)
			i++
		}
	}

	return tst
}

func (s *KeywordMessageContext) Keyword(i int) IKeywordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *KeywordMessageContext) AllFormula() []IFormulaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormulaContext); ok {
			len++
		}
	}

	tst := make([]IFormulaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormulaContext); ok {
			tst[i] = t.(IFormulaContext)
			i++
		}
	}

	return tst
}

func (s *KeywordMessageContext) Formula(i int) IFormulaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormulaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *KeywordMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterKeywordMessage(s)
	}
}

func (s *KeywordMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitKeywordMessage(s)
	}
}

func (p *SomParser) KeywordMessage() (localctx IKeywordMessageContext) {
	localctx = NewKeywordMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SomParserRULE_keywordMessage)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SomParserKeyword {
		{
			p.SetState(275)
			p.Keyword()
		}
		{
			p.SetState(276)
			p.Formula()
		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormulaContext is an interface to support dynamic dispatch.
type IFormulaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinaryOperand() IBinaryOperandContext
	AllBinaryMessage() []IBinaryMessageContext
	BinaryMessage(i int) IBinaryMessageContext

	// IsFormulaContext differentiates from other interfaces.
	IsFormulaContext()
}

type FormulaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulaContext() *FormulaContext {
	var p = new(FormulaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_formula
	return p
}

func InitEmptyFormulaContext(p *FormulaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_formula
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_formula

	return p
}

func (s *FormulaContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulaContext) BinaryOperand() IBinaryOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperandContext)
}

func (s *FormulaContext) AllBinaryMessage() []IBinaryMessageContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinaryMessageContext); ok {
			len++
		}
	}

	tst := make([]IBinaryMessageContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinaryMessageContext); ok {
			tst[i] = t.(IBinaryMessageContext)
			i++
		}
	}

	return tst
}

func (s *FormulaContext) BinaryMessage(i int) IBinaryMessageContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryMessageContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryMessageContext)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (p *SomParser) Formula() (localctx IFormulaContext) {
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SomParserRULE_formula)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.BinaryOperand()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8388128) != 0 {
		{
			p.SetState(283)
			p.BinaryMessage()
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedTermContext is an interface to support dynamic dispatch.
type INestedTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NewTerm() antlr.TerminalNode
	Expression() IExpressionContext
	EndTerm() antlr.TerminalNode

	// IsNestedTermContext differentiates from other interfaces.
	IsNestedTermContext()
}

type NestedTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedTermContext() *NestedTermContext {
	var p = new(NestedTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_nestedTerm
	return p
}

func InitEmptyNestedTermContext(p *NestedTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_nestedTerm
}

func (*NestedTermContext) IsNestedTermContext() {}

func NewNestedTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedTermContext {
	var p = new(NestedTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_nestedTerm

	return p
}

func (s *NestedTermContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedTermContext) NewTerm() antlr.TerminalNode {
	return s.GetToken(SomParserNewTerm, 0)
}

func (s *NestedTermContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedTermContext) EndTerm() antlr.TerminalNode {
	return s.GetToken(SomParserEndTerm, 0)
}

func (s *NestedTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterNestedTerm(s)
	}
}

func (s *NestedTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitNestedTerm(s)
	}
}

func (p *SomParser) NestedTerm() (localctx INestedTermContext) {
	localctx = NewNestedTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SomParserRULE_nestedTerm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(SomParserNewTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Expression()
	}
	{
		p.SetState(291)
		p.Match(SomParserEndTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralArray() ILiteralArrayContext
	LiteralSymbol() ILiteralSymbolContext
	LiteralString() ILiteralStringContext
	LiteralNumber() ILiteralNumberContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) LiteralArray() ILiteralArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralArrayContext)
}

func (s *LiteralContext) LiteralSymbol() ILiteralSymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralSymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralSymbolContext)
}

func (s *LiteralContext) LiteralString() ILiteralStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralStringContext)
}

func (s *LiteralContext) LiteralNumber() ILiteralNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralNumberContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *SomParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SomParserRULE_literal)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.LiteralArray()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.LiteralSymbol()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.LiteralString()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(296)
			p.LiteralNumber()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralArrayContext is an interface to support dynamic dispatch.
type ILiteralArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pound() antlr.TerminalNode
	NewTerm() antlr.TerminalNode
	EndTerm() antlr.TerminalNode
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext

	// IsLiteralArrayContext differentiates from other interfaces.
	IsLiteralArrayContext()
}

type LiteralArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralArrayContext() *LiteralArrayContext {
	var p = new(LiteralArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalArray
	return p
}

func InitEmptyLiteralArrayContext(p *LiteralArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalArray
}

func (*LiteralArrayContext) IsLiteralArrayContext() {}

func NewLiteralArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralArrayContext {
	var p = new(LiteralArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literalArray

	return p
}

func (s *LiteralArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralArrayContext) Pound() antlr.TerminalNode {
	return s.GetToken(SomParserPound, 0)
}

func (s *LiteralArrayContext) NewTerm() antlr.TerminalNode {
	return s.GetToken(SomParserNewTerm, 0)
}

func (s *LiteralArrayContext) EndTerm() antlr.TerminalNode {
	return s.GetToken(SomParserEndTerm, 0)
}

func (s *LiteralArrayContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *LiteralArrayContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteralArray(s)
	}
}

func (s *LiteralArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteralArray(s)
	}
}

func (p *SomParser) LiteralArray() (localctx ILiteralArrayContext) {
	localctx = NewLiteralArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SomParserRULE_literalArray)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(SomParserPound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(SomParserNewTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20468205568) != 0 {
		{
			p.SetState(301)
			p.Literal()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
		p.Match(SomParserEndTerm)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralNumberContext is an interface to support dynamic dispatch.
type ILiteralNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NegativeDecimal() INegativeDecimalContext
	LiteralDecimal() ILiteralDecimalContext

	// IsLiteralNumberContext differentiates from other interfaces.
	IsLiteralNumberContext()
}

type LiteralNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralNumberContext() *LiteralNumberContext {
	var p = new(LiteralNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalNumber
	return p
}

func InitEmptyLiteralNumberContext(p *LiteralNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalNumber
}

func (*LiteralNumberContext) IsLiteralNumberContext() {}

func NewLiteralNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralNumberContext {
	var p = new(LiteralNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literalNumber

	return p
}

func (s *LiteralNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralNumberContext) NegativeDecimal() INegativeDecimalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INegativeDecimalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INegativeDecimalContext)
}

func (s *LiteralNumberContext) LiteralDecimal() ILiteralDecimalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralDecimalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralDecimalContext)
}

func (s *LiteralNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteralNumber(s)
	}
}

func (s *LiteralNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteralNumber(s)
	}
}

func (p *SomParser) LiteralNumber() (localctx ILiteralNumberContext) {
	localctx = NewLiteralNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SomParserRULE_literalNumber)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserMinus:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.NegativeDecimal()
		}

	case SomParserInteger, SomParserDouble:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.LiteralDecimal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralDecimalContext is an interface to support dynamic dispatch.
type ILiteralDecimalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralInteger() ILiteralIntegerContext
	LiteralDouble() ILiteralDoubleContext

	// IsLiteralDecimalContext differentiates from other interfaces.
	IsLiteralDecimalContext()
}

type LiteralDecimalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralDecimalContext() *LiteralDecimalContext {
	var p = new(LiteralDecimalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalDecimal
	return p
}

func InitEmptyLiteralDecimalContext(p *LiteralDecimalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalDecimal
}

func (*LiteralDecimalContext) IsLiteralDecimalContext() {}

func NewLiteralDecimalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralDecimalContext {
	var p = new(LiteralDecimalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literalDecimal

	return p
}

func (s *LiteralDecimalContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralDecimalContext) LiteralInteger() ILiteralIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralIntegerContext)
}

func (s *LiteralDecimalContext) LiteralDouble() ILiteralDoubleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralDoubleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralDoubleContext)
}

func (s *LiteralDecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralDecimalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralDecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteralDecimal(s)
	}
}

func (s *LiteralDecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteralDecimal(s)
	}
}

func (p *SomParser) LiteralDecimal() (localctx ILiteralDecimalContext) {
	localctx = NewLiteralDecimalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SomParserRULE_literalDecimal)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserInteger:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.LiteralInteger()
		}

	case SomParserDouble:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(314)
			p.LiteralDouble()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INegativeDecimalContext is an interface to support dynamic dispatch.
type INegativeDecimalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Minus() antlr.TerminalNode
	LiteralDecimal() ILiteralDecimalContext

	// IsNegativeDecimalContext differentiates from other interfaces.
	IsNegativeDecimalContext()
}

type NegativeDecimalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegativeDecimalContext() *NegativeDecimalContext {
	var p = new(NegativeDecimalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_negativeDecimal
	return p
}

func InitEmptyNegativeDecimalContext(p *NegativeDecimalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_negativeDecimal
}

func (*NegativeDecimalContext) IsNegativeDecimalContext() {}

func NewNegativeDecimalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NegativeDecimalContext {
	var p = new(NegativeDecimalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_negativeDecimal

	return p
}

func (s *NegativeDecimalContext) GetParser() antlr.Parser { return s.parser }

func (s *NegativeDecimalContext) Minus() antlr.TerminalNode {
	return s.GetToken(SomParserMinus, 0)
}

func (s *NegativeDecimalContext) LiteralDecimal() ILiteralDecimalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralDecimalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralDecimalContext)
}

func (s *NegativeDecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeDecimalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NegativeDecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterNegativeDecimal(s)
	}
}

func (s *NegativeDecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitNegativeDecimal(s)
	}
}

func (p *SomParser) NegativeDecimal() (localctx INegativeDecimalContext) {
	localctx = NewNegativeDecimalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SomParserRULE_negativeDecimal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(SomParserMinus)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.LiteralDecimal()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralIntegerContext is an interface to support dynamic dispatch.
type ILiteralIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() antlr.TerminalNode

	// IsLiteralIntegerContext differentiates from other interfaces.
	IsLiteralIntegerContext()
}

type LiteralIntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralIntegerContext() *LiteralIntegerContext {
	var p = new(LiteralIntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalInteger
	return p
}

func InitEmptyLiteralIntegerContext(p *LiteralIntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalInteger
}

func (*LiteralIntegerContext) IsLiteralIntegerContext() {}

func NewLiteralIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralIntegerContext {
	var p = new(LiteralIntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literalInteger

	return p
}

func (s *LiteralIntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralIntegerContext) Integer() antlr.TerminalNode {
	return s.GetToken(SomParserInteger, 0)
}

func (s *LiteralIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralIntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteralInteger(s)
	}
}

func (s *LiteralIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteralInteger(s)
	}
}

func (p *SomParser) LiteralInteger() (localctx ILiteralIntegerContext) {
	localctx = NewLiteralIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SomParserRULE_literalInteger)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(SomParserInteger)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralDoubleContext is an interface to support dynamic dispatch.
type ILiteralDoubleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Double() antlr.TerminalNode

	// IsLiteralDoubleContext differentiates from other interfaces.
	IsLiteralDoubleContext()
}

type LiteralDoubleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralDoubleContext() *LiteralDoubleContext {
	var p = new(LiteralDoubleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalDouble
	return p
}

func InitEmptyLiteralDoubleContext(p *LiteralDoubleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalDouble
}

func (*LiteralDoubleContext) IsLiteralDoubleContext() {}

func NewLiteralDoubleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralDoubleContext {
	var p = new(LiteralDoubleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literalDouble

	return p
}

func (s *LiteralDoubleContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralDoubleContext) Double() antlr.TerminalNode {
	return s.GetToken(SomParserDouble, 0)
}

func (s *LiteralDoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralDoubleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralDoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteralDouble(s)
	}
}

func (s *LiteralDoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteralDouble(s)
	}
}

func (p *SomParser) LiteralDouble() (localctx ILiteralDoubleContext) {
	localctx = NewLiteralDoubleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SomParserRULE_literalDouble)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(SomParserDouble)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralSymbolContext is an interface to support dynamic dispatch.
type ILiteralSymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pound() antlr.TerminalNode
	String_() IStringContext
	Selector() ISelectorContext

	// IsLiteralSymbolContext differentiates from other interfaces.
	IsLiteralSymbolContext()
}

type LiteralSymbolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralSymbolContext() *LiteralSymbolContext {
	var p = new(LiteralSymbolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalSymbol
	return p
}

func InitEmptyLiteralSymbolContext(p *LiteralSymbolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalSymbol
}

func (*LiteralSymbolContext) IsLiteralSymbolContext() {}

func NewLiteralSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralSymbolContext {
	var p = new(LiteralSymbolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literalSymbol

	return p
}

func (s *LiteralSymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralSymbolContext) Pound() antlr.TerminalNode {
	return s.GetToken(SomParserPound, 0)
}

func (s *LiteralSymbolContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *LiteralSymbolContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *LiteralSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteralSymbol(s)
	}
}

func (s *LiteralSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteralSymbol(s)
	}
}

func (p *SomParser) LiteralSymbol() (localctx ILiteralSymbolContext) {
	localctx = NewLiteralSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SomParserRULE_literalSymbol)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(SomParserPound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserSTString:
		{
			p.SetState(325)
			p.String_()
		}

	case SomParserPrimitive, SomParserIdentifier, SomParserEqual, SomParserOr, SomParserComma, SomParserMinus, SomParserNot, SomParserAnd, SomParserStar, SomParserDiv, SomParserMod, SomParserPlus, SomParserMore, SomParserLess, SomParserAt, SomParserPer, SomParserOperatorSequence, SomParserKeyword, SomParserKeywordSequence:
		{
			p.SetState(326)
			p.Selector()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralStringContext is an interface to support dynamic dispatch.
type ILiteralStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IStringContext

	// IsLiteralStringContext differentiates from other interfaces.
	IsLiteralStringContext()
}

type LiteralStringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralStringContext() *LiteralStringContext {
	var p = new(LiteralStringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalString
	return p
}

func InitEmptyLiteralStringContext(p *LiteralStringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_literalString
}

func (*LiteralStringContext) IsLiteralStringContext() {}

func NewLiteralStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralStringContext {
	var p = new(LiteralStringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_literalString

	return p
}

func (s *LiteralStringContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralStringContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *LiteralStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterLiteralString(s)
	}
}

func (s *LiteralStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitLiteralString(s)
	}
}

func (p *SomParser) LiteralString() (localctx ILiteralStringContext) {
	localctx = NewLiteralStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SomParserRULE_literalString)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.String_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinarySelector() IBinarySelectorContext
	KeywordSelector() IKeywordSelectorContext
	UnarySelector() IUnarySelectorContext

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) BinarySelector() IBinarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinarySelectorContext)
}

func (s *SelectorContext) KeywordSelector() IKeywordSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordSelectorContext)
}

func (s *SelectorContext) UnarySelector() IUnarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarySelectorContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *SomParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SomParserRULE_selector)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SomParserEqual, SomParserOr, SomParserComma, SomParserMinus, SomParserNot, SomParserAnd, SomParserStar, SomParserDiv, SomParserMod, SomParserPlus, SomParserMore, SomParserLess, SomParserAt, SomParserPer, SomParserOperatorSequence:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.BinarySelector()
		}

	case SomParserKeyword, SomParserKeywordSequence:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.KeywordSelector()
		}

	case SomParserPrimitive, SomParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(333)
			p.UnarySelector()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordSelectorContext is an interface to support dynamic dispatch.
type IKeywordSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword() antlr.TerminalNode
	KeywordSequence() antlr.TerminalNode

	// IsKeywordSelectorContext differentiates from other interfaces.
	IsKeywordSelectorContext()
}

type KeywordSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordSelectorContext() *KeywordSelectorContext {
	var p = new(KeywordSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keywordSelector
	return p
}

func InitEmptyKeywordSelectorContext(p *KeywordSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_keywordSelector
}

func (*KeywordSelectorContext) IsKeywordSelectorContext() {}

func NewKeywordSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordSelectorContext {
	var p = new(KeywordSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_keywordSelector

	return p
}

func (s *KeywordSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordSelectorContext) Keyword() antlr.TerminalNode {
	return s.GetToken(SomParserKeyword, 0)
}

func (s *KeywordSelectorContext) KeywordSequence() antlr.TerminalNode {
	return s.GetToken(SomParserKeywordSequence, 0)
}

func (s *KeywordSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterKeywordSelector(s)
	}
}

func (s *KeywordSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitKeywordSelector(s)
	}
}

func (p *SomParser) KeywordSelector() (localctx IKeywordSelectorContext) {
	localctx = NewKeywordSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SomParserRULE_keywordSelector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SomParserKeyword || _la == SomParserKeywordSequence) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STString() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STString() antlr.TerminalNode {
	return s.GetToken(SomParserSTString, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *SomParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SomParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(SomParserSTString)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedBlockContext is an interface to support dynamic dispatch.
type INestedBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NewBlock() antlr.TerminalNode
	EndBlock() antlr.TerminalNode
	BlockPattern() IBlockPatternContext
	BlockContents() IBlockContentsContext

	// IsNestedBlockContext differentiates from other interfaces.
	IsNestedBlockContext()
}

type NestedBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedBlockContext() *NestedBlockContext {
	var p = new(NestedBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_nestedBlock
	return p
}

func InitEmptyNestedBlockContext(p *NestedBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_nestedBlock
}

func (*NestedBlockContext) IsNestedBlockContext() {}

func NewNestedBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedBlockContext {
	var p = new(NestedBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_nestedBlock

	return p
}

func (s *NestedBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedBlockContext) NewBlock() antlr.TerminalNode {
	return s.GetToken(SomParserNewBlock, 0)
}

func (s *NestedBlockContext) EndBlock() antlr.TerminalNode {
	return s.GetToken(SomParserEndBlock, 0)
}

func (s *NestedBlockContext) BlockPattern() IBlockPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockPatternContext)
}

func (s *NestedBlockContext) BlockContents() IBlockContentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContentsContext)
}

func (s *NestedBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterNestedBlock(s)
	}
}

func (s *NestedBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitNestedBlock(s)
	}
}

func (p *SomParser) NestedBlock() (localctx INestedBlockContext) {
	localctx = NewNestedBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SomParserRULE_nestedBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(SomParserNewBlock)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SomParserColon {
		{
			p.SetState(341)
			p.BlockPattern()
		}

	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20619201176) != 0 {
		{
			p.SetState(344)
			p.BlockContents()
		}

	}
	{
		p.SetState(347)
		p.Match(SomParserEndBlock)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockPatternContext is an interface to support dynamic dispatch.
type IBlockPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BlockArguments() IBlockArgumentsContext
	Or() antlr.TerminalNode

	// IsBlockPatternContext differentiates from other interfaces.
	IsBlockPatternContext()
}

type BlockPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockPatternContext() *BlockPatternContext {
	var p = new(BlockPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockPattern
	return p
}

func InitEmptyBlockPatternContext(p *BlockPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockPattern
}

func (*BlockPatternContext) IsBlockPatternContext() {}

func NewBlockPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockPatternContext {
	var p = new(BlockPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_blockPattern

	return p
}

func (s *BlockPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockPatternContext) BlockArguments() IBlockArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockArgumentsContext)
}

func (s *BlockPatternContext) Or() antlr.TerminalNode {
	return s.GetToken(SomParserOr, 0)
}

func (s *BlockPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBlockPattern(s)
	}
}

func (s *BlockPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBlockPattern(s)
	}
}

func (p *SomParser) BlockPattern() (localctx IBlockPatternContext) {
	localctx = NewBlockPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SomParserRULE_blockPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.BlockArguments()
	}
	{
		p.SetState(350)
		p.Match(SomParserOr)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockArgumentsContext is an interface to support dynamic dispatch.
type IBlockArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColon() []antlr.TerminalNode
	Colon(i int) antlr.TerminalNode
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext

	// IsBlockArgumentsContext differentiates from other interfaces.
	IsBlockArgumentsContext()
}

type BlockArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockArgumentsContext() *BlockArgumentsContext {
	var p = new(BlockArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockArguments
	return p
}

func InitEmptyBlockArgumentsContext(p *BlockArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SomParserRULE_blockArguments
}

func (*BlockArgumentsContext) IsBlockArgumentsContext() {}

func NewBlockArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockArgumentsContext {
	var p = new(BlockArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SomParserRULE_blockArguments

	return p
}

func (s *BlockArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockArgumentsContext) AllColon() []antlr.TerminalNode {
	return s.GetTokens(SomParserColon)
}

func (s *BlockArgumentsContext) Colon(i int) antlr.TerminalNode {
	return s.GetToken(SomParserColon, i)
}

func (s *BlockArgumentsContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *BlockArgumentsContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *BlockArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.EnterBlockArguments(s)
	}
}

func (s *BlockArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SomListener); ok {
		listenerT.ExitBlockArguments(s)
	}
}

func (p *SomParser) BlockArguments() (localctx IBlockArgumentsContext) {
	localctx = NewBlockArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SomParserRULE_blockArguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SomParserColon {
		{
			p.SetState(352)
			p.Match(SomParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Argument()
		}

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
