// Code generated from /workspace/grammar/TemplateLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package base

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TemplateLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TemplateLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func templatelexerLexerInit() {
	staticData := &TemplateLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "Expression", "Statement", "StatementExpression", "AssignmentExpression",
	}
	staticData.LiteralNames = []string{
		"", "'{{'", "", "", "'}}'", "", "", "", "", "", "", "", "", "", "",
		"", "','",
	}
	staticData.SymbolicNames = []string{
		"", "OPEN_EXPR", "OPEN_STMT", "TEXT", "CLOSE_EXPR", "EXPR", "CLOSE_STMT",
		"FOR", "IN", "ENDFOR", "IF", "ELSEIF", "ELSE", "ENDIF", "NAME", "EQ",
		"COMMA", "WS", "CLOSE_STMT_EXPR", "STMT_EXPR", "SEMICOLON", "ASSIGN_EXPR",
	}
	staticData.RuleNames = []string{
		"OPEN_EXPR", "OPEN_STMT", "TEXT", "CLOSE_EXPR", "EXPR", "CLOSE_STMT",
		"FOR", "IN", "ENDFOR", "IF", "ELSEIF", "ELSE", "ENDIF", "NAME", "EQ",
		"COMMA", "WS", "N", "S", "CLOSE_STMT_EXPR", "STMT_EXPR", "SEMICOLON",
		"ASSIGN_EXPR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 248, 6, -1, 6, -1, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7,
		1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7,
		7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2,
		13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18,
		7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 61, 8, 1, 10, 1, 12,
		1, 64, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 4, 2, 71, 8, 2, 11, 2, 12, 2,
		72, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 4, 4, 83, 8, 4, 11,
		4, 12, 4, 84, 1, 5, 5, 5, 88, 8, 5, 10, 5, 12, 5, 91, 9, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 103, 8, 6, 11, 6,
		12, 6, 104, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 111, 8, 7, 11, 7, 12, 7, 112,
		1, 7, 1, 7, 1, 8, 5, 8, 118, 8, 8, 10, 8, 12, 8, 121, 9, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 131, 8, 8, 10, 8, 12, 8, 134,
		9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 140, 8, 9, 11, 9, 12, 9, 141, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 154,
		8, 10, 11, 10, 12, 10, 155, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 5, 11, 166, 8, 11, 10, 11, 12, 11, 169, 9, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 178, 8, 12, 10, 12, 12, 12, 181,
		9, 12, 1, 13, 1, 13, 5, 13, 185, 8, 13, 10, 13, 12, 13, 188, 9, 13, 1,
		14, 5, 14, 191, 8, 14, 10, 14, 12, 14, 194, 9, 14, 1, 14, 1, 14, 5, 14,
		198, 8, 14, 10, 14, 12, 14, 201, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 3, 16, 209, 8, 16, 1, 17, 3, 17, 212, 8, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 4, 20, 225,
		8, 20, 11, 20, 12, 20, 226, 1, 21, 5, 21, 230, 8, 21, 10, 21, 12, 21, 233,
		9, 21, 1, 21, 1, 21, 5, 21, 237, 8, 21, 10, 21, 12, 21, 240, 9, 21, 1,
		21, 1, 21, 1, 22, 4, 22, 245, 8, 22, 11, 22, 12, 22, 246, 0, 0, 23, 5,
		1, 7, 2, 9, 3, 11, 4, 13, 5, 15, 6, 17, 7, 19, 8, 21, 9, 23, 10, 25, 11,
		27, 12, 29, 13, 31, 14, 33, 15, 35, 16, 37, 17, 39, 0, 41, 0, 43, 18, 45,
		19, 47, 20, 49, 21, 5, 0, 1, 2, 3, 4, 8, 1, 0, 123, 123, 2, 0, 37, 37,
		123, 123, 1, 0, 125, 125, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 2, 0, 9, 9, 32, 32, 1, 0, 37, 37, 1, 0, 59, 59,
		265, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0,
		0, 0, 1, 13, 1, 0, 0, 0, 2, 15, 1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 2, 19, 1,
		0, 0, 0, 2, 21, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 2, 27,
		1, 0, 0, 0, 2, 29, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 2,
		35, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0,
		4, 47, 1, 0, 0, 0, 4, 49, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 56, 1, 0, 0,
		0, 9, 70, 1, 0, 0, 0, 11, 74, 1, 0, 0, 0, 13, 82, 1, 0, 0, 0, 15, 89, 1,
		0, 0, 0, 17, 97, 1, 0, 0, 0, 19, 106, 1, 0, 0, 0, 21, 119, 1, 0, 0, 0,
		23, 135, 1, 0, 0, 0, 25, 145, 1, 0, 0, 0, 27, 159, 1, 0, 0, 0, 29, 170,
		1, 0, 0, 0, 31, 182, 1, 0, 0, 0, 33, 192, 1, 0, 0, 0, 35, 204, 1, 0, 0,
		0, 37, 208, 1, 0, 0, 0, 39, 211, 1, 0, 0, 0, 41, 215, 1, 0, 0, 0, 43, 217,
		1, 0, 0, 0, 45, 224, 1, 0, 0, 0, 47, 231, 1, 0, 0, 0, 49, 244, 1, 0, 0,
		0, 51, 52, 5, 123, 0, 0, 52, 53, 5, 123, 0, 0, 53, 54, 1, 0, 0, 0, 54,
		55, 6, 0, 0, 0, 55, 6, 1, 0, 0, 0, 56, 57, 5, 123, 0, 0, 57, 58, 5, 37,
		0, 0, 58, 62, 1, 0, 0, 0, 59, 61, 3, 37, 16, 0, 60, 59, 1, 0, 0, 0, 61,
		64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0,
		0, 64, 62, 1, 0, 0, 0, 65, 66, 6, 1, 1, 0, 66, 8, 1, 0, 0, 0, 67, 71, 8,
		0, 0, 0, 68, 69, 5, 123, 0, 0, 69, 71, 8, 1, 0, 0, 70, 67, 1, 0, 0, 0,
		70, 68, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1,
		0, 0, 0, 73, 10, 1, 0, 0, 0, 74, 75, 5, 125, 0, 0, 75, 76, 5, 125, 0, 0,
		76, 77, 1, 0, 0, 0, 77, 78, 6, 3, 2, 0, 78, 12, 1, 0, 0, 0, 79, 83, 8,
		2, 0, 0, 80, 81, 5, 125, 0, 0, 81, 83, 8, 2, 0, 0, 82, 79, 1, 0, 0, 0,
		82, 80, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1,
		0, 0, 0, 85, 14, 1, 0, 0, 0, 86, 88, 3, 37, 16, 0, 87, 86, 1, 0, 0, 0,
		88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1,
		0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 37, 0, 0, 93, 94, 5, 125, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 96, 6, 5, 3, 0, 96, 16, 1, 0, 0, 0, 97, 98, 5,
		102, 0, 0, 98, 99, 5, 111, 0, 0, 99, 100, 5, 114, 0, 0, 100, 102, 1, 0,
		0, 0, 101, 103, 3, 37, 16, 0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0,
		0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 18, 1, 0, 0, 0, 106,
		107, 5, 105, 0, 0, 107, 108, 5, 110, 0, 0, 108, 110, 1, 0, 0, 0, 109, 111,
		3, 37, 16, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 110, 1,
		0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 6, 7, 4,
		0, 115, 20, 1, 0, 0, 0, 116, 118, 3, 37, 16, 0, 117, 116, 1, 0, 0, 0, 118,
		121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122,
		1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5, 101, 0, 0, 123, 124, 5,
		110, 0, 0, 124, 125, 5, 100, 0, 0, 125, 126, 5, 102, 0, 0, 126, 127, 5,
		111, 0, 0, 127, 128, 5, 114, 0, 0, 128, 132, 1, 0, 0, 0, 129, 131, 3, 37,
		16, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 22, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136,
		5, 105, 0, 0, 136, 137, 5, 102, 0, 0, 137, 139, 1, 0, 0, 0, 138, 140, 3,
		37, 16, 0, 139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0,
		0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 6, 9, 4, 0,
		144, 24, 1, 0, 0, 0, 145, 146, 5, 101, 0, 0, 146, 147, 5, 108, 0, 0, 147,
		148, 5, 115, 0, 0, 148, 149, 5, 101, 0, 0, 149, 150, 5, 105, 0, 0, 150,
		151, 5, 102, 0, 0, 151, 153, 1, 0, 0, 0, 152, 154, 3, 37, 16, 0, 153, 152,
		1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 6, 10, 4, 0, 158, 26, 1, 0, 0, 0,
		159, 160, 5, 101, 0, 0, 160, 161, 5, 108, 0, 0, 161, 162, 5, 115, 0, 0,
		162, 163, 5, 101, 0, 0, 163, 167, 1, 0, 0, 0, 164, 166, 3, 37, 16, 0, 165,
		164, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168,
		1, 0, 0, 0, 168, 28, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 171, 5, 101,
		0, 0, 171, 172, 5, 110, 0, 0, 172, 173, 5, 100, 0, 0, 173, 174, 5, 105,
		0, 0, 174, 175, 5, 102, 0, 0, 175, 179, 1, 0, 0, 0, 176, 178, 3, 37, 16,
		0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 30, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 186, 7,
		3, 0, 0, 183, 185, 7, 4, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0,
		0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 32, 1, 0, 0, 0, 188,
		186, 1, 0, 0, 0, 189, 191, 3, 37, 16, 0, 190, 189, 1, 0, 0, 0, 191, 194,
		1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 1, 0,
		0, 0, 194, 192, 1, 0, 0, 0, 195, 199, 5, 61, 0, 0, 196, 198, 3, 37, 16,
		0, 197, 196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203,
		6, 14, 5, 0, 203, 34, 1, 0, 0, 0, 204, 205, 5, 44, 0, 0, 205, 36, 1, 0,
		0, 0, 206, 209, 3, 39, 17, 0, 207, 209, 3, 41, 18, 0, 208, 206, 1, 0, 0,
		0, 208, 207, 1, 0, 0, 0, 209, 38, 1, 0, 0, 0, 210, 212, 5, 13, 0, 0, 211,
		210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214,
		5, 10, 0, 0, 214, 40, 1, 0, 0, 0, 215, 216, 7, 5, 0, 0, 216, 42, 1, 0,
		0, 0, 217, 218, 3, 15, 5, 0, 218, 219, 1, 0, 0, 0, 219, 220, 6, 19, 3,
		0, 220, 44, 1, 0, 0, 0, 221, 225, 8, 6, 0, 0, 222, 223, 5, 37, 0, 0, 223,
		225, 8, 2, 0, 0, 224, 221, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 226,
		1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 46, 1, 0,
		0, 0, 228, 230, 3, 37, 16, 0, 229, 228, 1, 0, 0, 0, 230, 233, 1, 0, 0,
		0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0, 0, 0, 233,
		231, 1, 0, 0, 0, 234, 238, 5, 59, 0, 0, 235, 237, 3, 37, 16, 0, 236, 235,
		1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0,
		0, 0, 239, 241, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 242, 6, 21, 6, 0,
		242, 48, 1, 0, 0, 0, 243, 245, 8, 7, 0, 0, 244, 243, 1, 0, 0, 0, 245, 246,
		1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 50, 1, 0,
		0, 0, 29, 0, 1, 2, 3, 4, 62, 70, 72, 82, 84, 89, 104, 112, 119, 132, 141,
		155, 167, 179, 186, 192, 199, 208, 211, 224, 226, 231, 238, 246, 7, 5,
		1, 0, 5, 2, 0, 5, 0, 0, 2, 0, 0, 5, 3, 0, 5, 4, 0, 4, 0, 0,
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

// TemplateLexerInit initializes any static state used to implement TemplateLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTemplateLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TemplateLexerInit() {
	staticData := &TemplateLexerLexerStaticData
	staticData.once.Do(templatelexerLexerInit)
}

// NewTemplateLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTemplateLexer(input antlr.CharStream) *TemplateLexer {
	TemplateLexerInit()
	l := new(TemplateLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TemplateLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "TemplateLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TemplateLexer tokens.
const (
	TemplateLexerOPEN_EXPR       = 1
	TemplateLexerOPEN_STMT       = 2
	TemplateLexerTEXT            = 3
	TemplateLexerCLOSE_EXPR      = 4
	TemplateLexerEXPR            = 5
	TemplateLexerCLOSE_STMT      = 6
	TemplateLexerFOR             = 7
	TemplateLexerIN              = 8
	TemplateLexerENDFOR          = 9
	TemplateLexerIF              = 10
	TemplateLexerELSEIF          = 11
	TemplateLexerELSE            = 12
	TemplateLexerENDIF           = 13
	TemplateLexerNAME            = 14
	TemplateLexerEQ              = 15
	TemplateLexerCOMMA           = 16
	TemplateLexerWS              = 17
	TemplateLexerCLOSE_STMT_EXPR = 18
	TemplateLexerSTMT_EXPR       = 19
	TemplateLexerSEMICOLON       = 20
	TemplateLexerASSIGN_EXPR     = 21
)

// TemplateLexer modes.
const (
	TemplateLexerExpression = iota + 1
	TemplateLexerStatement
	TemplateLexerStatementExpression
	TemplateLexerAssignmentExpression
)
