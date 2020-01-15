// Code generated from Clojure.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Clojure

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 258,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 7, 2, 84, 10, 2, 12, 2, 14, 2, 87,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 96, 10, 3, 3, 4,
	7, 4, 99, 10, 4, 12, 4, 14, 4, 102, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 116, 10, 7, 12, 7, 14,
	7, 119, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 142, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 7, 17, 168, 10, 17, 12, 17, 14,
	17, 171, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18,
	180, 10, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 207, 10, 24, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 5, 29, 222, 10, 29, 3, 30, 3, 30, 3, 30, 5, 30, 227, 10, 30, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 5, 35,
	239, 10, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3,
	38, 5, 38, 250, 10, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41,
	2, 2, 42, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 2, 2, 2, 255, 2, 85, 3, 2, 2, 2, 4, 95, 3, 2, 2, 2,
	6, 100, 3, 2, 2, 2, 8, 103, 3, 2, 2, 2, 10, 107, 3, 2, 2, 2, 12, 111, 3,
	2, 2, 2, 14, 122, 3, 2, 2, 2, 16, 141, 3, 2, 2, 2, 18, 143, 3, 2, 2, 2,
	20, 146, 3, 2, 2, 2, 22, 149, 3, 2, 2, 2, 24, 152, 3, 2, 2, 2, 26, 155,
	3, 2, 2, 2, 28, 159, 3, 2, 2, 2, 30, 162, 3, 2, 2, 2, 32, 165, 3, 2, 2,
	2, 34, 174, 3, 2, 2, 2, 36, 181, 3, 2, 2, 2, 38, 184, 3, 2, 2, 2, 40, 188,
	3, 2, 2, 2, 42, 191, 3, 2, 2, 2, 44, 195, 3, 2, 2, 2, 46, 206, 3, 2, 2,
	2, 48, 208, 3, 2, 2, 2, 50, 210, 3, 2, 2, 2, 52, 212, 3, 2, 2, 2, 54, 214,
	3, 2, 2, 2, 56, 221, 3, 2, 2, 2, 58, 226, 3, 2, 2, 2, 60, 228, 3, 2, 2,
	2, 62, 230, 3, 2, 2, 2, 64, 232, 3, 2, 2, 2, 66, 234, 3, 2, 2, 2, 68, 238,
	3, 2, 2, 2, 70, 240, 3, 2, 2, 2, 72, 243, 3, 2, 2, 2, 74, 249, 3, 2, 2,
	2, 76, 251, 3, 2, 2, 2, 78, 253, 3, 2, 2, 2, 80, 255, 3, 2, 2, 2, 82, 84,
	5, 4, 3, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2,
	85, 86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 89, 7,
	2, 2, 3, 89, 3, 3, 2, 2, 2, 90, 96, 5, 46, 24, 2, 91, 96, 5, 8, 5, 2, 92,
	96, 5, 10, 6, 2, 93, 96, 5, 12, 7, 2, 94, 96, 5, 16, 9, 2, 95, 90, 3, 2,
	2, 2, 95, 91, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 94,
	3, 2, 2, 2, 96, 5, 3, 2, 2, 2, 97, 99, 5, 4, 3, 2, 98, 97, 3, 2, 2, 2,
	99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 7,
	3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 104, 7, 3, 2, 2, 104, 105, 5, 6,
	4, 2, 105, 106, 7, 4, 2, 2, 106, 9, 3, 2, 2, 2, 107, 108, 7, 5, 2, 2, 108,
	109, 5, 6, 4, 2, 109, 110, 7, 6, 2, 2, 110, 11, 3, 2, 2, 2, 111, 117, 7,
	7, 2, 2, 112, 113, 5, 4, 3, 2, 113, 114, 5, 4, 3, 2, 114, 116, 3, 2, 2,
	2, 115, 112, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117,
	118, 3, 2, 2, 2, 118, 120, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 121,
	7, 8, 2, 2, 121, 13, 3, 2, 2, 2, 122, 123, 7, 9, 2, 2, 123, 124, 5, 6,
	4, 2, 124, 125, 7, 8, 2, 2, 125, 15, 3, 2, 2, 2, 126, 142, 5, 32, 17, 2,
	127, 142, 5, 34, 18, 2, 128, 142, 5, 44, 23, 2, 129, 142, 5, 36, 19, 2,
	130, 142, 5, 38, 20, 2, 131, 142, 5, 14, 8, 2, 132, 142, 5, 26, 14, 2,
	133, 142, 5, 40, 21, 2, 134, 142, 5, 42, 22, 2, 135, 142, 5, 28, 15, 2,
	136, 142, 5, 18, 10, 2, 137, 142, 5, 20, 11, 2, 138, 142, 5, 22, 12, 2,
	139, 142, 5, 24, 13, 2, 140, 142, 5, 30, 16, 2, 141, 126, 3, 2, 2, 2, 141,
	127, 3, 2, 2, 2, 141, 128, 3, 2, 2, 2, 141, 129, 3, 2, 2, 2, 141, 130,
	3, 2, 2, 2, 141, 131, 3, 2, 2, 2, 141, 132, 3, 2, 2, 2, 141, 133, 3, 2,
	2, 2, 141, 134, 3, 2, 2, 2, 141, 135, 3, 2, 2, 2, 141, 136, 3, 2, 2, 2,
	141, 137, 3, 2, 2, 2, 141, 138, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141,
	140, 3, 2, 2, 2, 142, 17, 3, 2, 2, 2, 143, 144, 7, 10, 2, 2, 144, 145,
	5, 4, 3, 2, 145, 19, 3, 2, 2, 2, 146, 147, 7, 11, 2, 2, 147, 148, 5, 4,
	3, 2, 148, 21, 3, 2, 2, 2, 149, 150, 7, 12, 2, 2, 150, 151, 5, 4, 3, 2,
	151, 23, 3, 2, 2, 2, 152, 153, 7, 13, 2, 2, 153, 154, 5, 4, 3, 2, 154,
	25, 3, 2, 2, 2, 155, 156, 7, 14, 2, 2, 156, 157, 5, 4, 3, 2, 157, 158,
	5, 4, 3, 2, 158, 27, 3, 2, 2, 2, 159, 160, 7, 15, 2, 2, 160, 161, 5, 4,
	3, 2, 161, 29, 3, 2, 2, 2, 162, 163, 7, 34, 2, 2, 163, 164, 7, 16, 2, 2,
	164, 31, 3, 2, 2, 2, 165, 169, 7, 17, 2, 2, 166, 168, 5, 4, 3, 2, 167,
	166, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170,
	3, 2, 2, 2, 170, 172, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 173, 7, 4,
	2, 2, 173, 33, 3, 2, 2, 2, 174, 179, 7, 18, 2, 2, 175, 176, 5, 12, 7, 2,
	176, 177, 5, 4, 3, 2, 177, 180, 3, 2, 2, 2, 178, 180, 5, 4, 3, 2, 179,
	175, 3, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180, 35, 3, 2, 2, 2, 181, 182, 7,
	19, 2, 2, 182, 183, 5, 74, 38, 2, 183, 37, 3, 2, 2, 2, 184, 185, 7, 20,
	2, 2, 185, 186, 5, 4, 3, 2, 186, 187, 5, 4, 3, 2, 187, 39, 3, 2, 2, 2,
	188, 189, 7, 21, 2, 2, 189, 190, 5, 4, 3, 2, 190, 41, 3, 2, 2, 2, 191,
	192, 7, 16, 2, 2, 192, 193, 5, 74, 38, 2, 193, 194, 5, 4, 3, 2, 194, 43,
	3, 2, 2, 2, 195, 196, 7, 16, 2, 2, 196, 197, 5, 48, 25, 2, 197, 45, 3,
	2, 2, 2, 198, 207, 5, 48, 25, 2, 199, 207, 5, 56, 29, 2, 200, 207, 5, 58,
	30, 2, 201, 207, 5, 66, 34, 2, 202, 207, 7, 33, 2, 2, 203, 207, 5, 68,
	35, 2, 204, 207, 5, 74, 38, 2, 205, 207, 5, 80, 41, 2, 206, 198, 3, 2,
	2, 2, 206, 199, 3, 2, 2, 2, 206, 200, 3, 2, 2, 2, 206, 201, 3, 2, 2, 2,
	206, 202, 3, 2, 2, 2, 206, 203, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206,
	205, 3, 2, 2, 2, 207, 47, 3, 2, 2, 2, 208, 209, 7, 23, 2, 2, 209, 49, 3,
	2, 2, 2, 210, 211, 7, 25, 2, 2, 211, 51, 3, 2, 2, 2, 212, 213, 7, 26, 2,
	2, 213, 53, 3, 2, 2, 2, 214, 215, 7, 28, 2, 2, 215, 55, 3, 2, 2, 2, 216,
	222, 7, 24, 2, 2, 217, 222, 5, 50, 26, 2, 218, 222, 5, 52, 27, 2, 219,
	222, 5, 54, 28, 2, 220, 222, 7, 27, 2, 2, 221, 216, 3, 2, 2, 2, 221, 217,
	3, 2, 2, 2, 221, 218, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 220, 3, 2,
	2, 2, 222, 57, 3, 2, 2, 2, 223, 227, 5, 60, 31, 2, 224, 227, 5, 64, 33,
	2, 225, 227, 5, 62, 32, 2, 226, 223, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2,
	226, 225, 3, 2, 2, 2, 227, 59, 3, 2, 2, 2, 228, 229, 7, 30, 2, 2, 229,
	61, 3, 2, 2, 2, 230, 231, 7, 31, 2, 2, 231, 63, 3, 2, 2, 2, 232, 233, 7,
	29, 2, 2, 233, 65, 3, 2, 2, 2, 234, 235, 7, 32, 2, 2, 235, 67, 3, 2, 2,
	2, 236, 239, 5, 72, 37, 2, 237, 239, 5, 70, 36, 2, 238, 236, 3, 2, 2, 2,
	238, 237, 3, 2, 2, 2, 239, 69, 3, 2, 2, 2, 240, 241, 7, 22, 2, 2, 241,
	242, 5, 74, 38, 2, 242, 71, 3, 2, 2, 2, 243, 244, 7, 22, 2, 2, 244, 245,
	7, 22, 2, 2, 245, 246, 5, 74, 38, 2, 246, 73, 3, 2, 2, 2, 247, 250, 5,
	78, 40, 2, 248, 250, 5, 76, 39, 2, 249, 247, 3, 2, 2, 2, 249, 248, 3, 2,
	2, 2, 250, 75, 3, 2, 2, 2, 251, 252, 7, 34, 2, 2, 252, 77, 3, 2, 2, 2,
	253, 254, 7, 35, 2, 2, 254, 79, 3, 2, 2, 2, 255, 256, 7, 36, 2, 2, 256,
	81, 3, 2, 2, 2, 14, 85, 95, 100, 117, 141, 169, 179, 206, 221, 226, 238,
	249,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'['", "']'", "'{'", "'}'", "'#{'", "'''", "'`'", "'~'",
	"'~@'", "'^'", "'@'", "'#'", "'#('", "'#^'", "'#''", "'#+'", "'#_'", "':'",
	"", "", "", "", "", "", "", "", "", "'nil'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "STRING", "FLOAT", "HEX", "BIN", "LONG", "BIGN", "CHAR_U",
	"CHAR_NAMED", "CHAR_ANY", "NIL", "BOOLEAN", "SYMBOL", "NS_SYMBOL", "PARAM_NAME",
	"TRASH",
}

var ruleNames = []string{
	"file", "form", "forms", "list", "vector", "map", "set", "reader_macro",
	"quote", "backtick", "unquote", "unquote_splicing", "tag", "deref", "gensym",
	"lambda", "meta_data", "var_quote", "host_expr", "discard", "dispatch",
	"regex", "literal", "string", "hex", "bin", "bign", "number", "character",
	"named_char", "any_char", "u_hex_quad", "nil", "keyword", "simple_keyword",
	"macro_keyword", "symbol", "simple_sym", "ns_symbol", "param_name",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ClojureParser struct {
	*antlr.BaseParser
}

func NewClojureParser(input antlr.TokenStream) *ClojureParser {
	this := new(ClojureParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Clojure.g4"

	return this
}

// ClojureParser tokens.
const (
	ClojureParserEOF        = antlr.TokenEOF
	ClojureParserT__0       = 1
	ClojureParserT__1       = 2
	ClojureParserT__2       = 3
	ClojureParserT__3       = 4
	ClojureParserT__4       = 5
	ClojureParserT__5       = 6
	ClojureParserT__6       = 7
	ClojureParserT__7       = 8
	ClojureParserT__8       = 9
	ClojureParserT__9       = 10
	ClojureParserT__10      = 11
	ClojureParserT__11      = 12
	ClojureParserT__12      = 13
	ClojureParserT__13      = 14
	ClojureParserT__14      = 15
	ClojureParserT__15      = 16
	ClojureParserT__16      = 17
	ClojureParserT__17      = 18
	ClojureParserT__18      = 19
	ClojureParserT__19      = 20
	ClojureParserSTRING     = 21
	ClojureParserFLOAT      = 22
	ClojureParserHEX        = 23
	ClojureParserBIN        = 24
	ClojureParserLONG       = 25
	ClojureParserBIGN       = 26
	ClojureParserCHAR_U     = 27
	ClojureParserCHAR_NAMED = 28
	ClojureParserCHAR_ANY   = 29
	ClojureParserNIL        = 30
	ClojureParserBOOLEAN    = 31
	ClojureParserSYMBOL     = 32
	ClojureParserNS_SYMBOL  = 33
	ClojureParserPARAM_NAME = 34
	ClojureParserTRASH      = 35
)

// ClojureParser rules.
const (
	ClojureParserRULE_file             = 0
	ClojureParserRULE_form             = 1
	ClojureParserRULE_forms            = 2
	ClojureParserRULE_list             = 3
	ClojureParserRULE_vector           = 4
	ClojureParserRULE_map              = 5
	ClojureParserRULE_set              = 6
	ClojureParserRULE_reader_macro     = 7
	ClojureParserRULE_quote            = 8
	ClojureParserRULE_backtick         = 9
	ClojureParserRULE_unquote          = 10
	ClojureParserRULE_unquote_splicing = 11
	ClojureParserRULE_tag              = 12
	ClojureParserRULE_deref            = 13
	ClojureParserRULE_gensym           = 14
	ClojureParserRULE_lambda           = 15
	ClojureParserRULE_meta_data        = 16
	ClojureParserRULE_var_quote        = 17
	ClojureParserRULE_host_expr        = 18
	ClojureParserRULE_discard          = 19
	ClojureParserRULE_dispatch         = 20
	ClojureParserRULE_regex            = 21
	ClojureParserRULE_literal          = 22
	ClojureParserRULE_string           = 23
	ClojureParserRULE_hex              = 24
	ClojureParserRULE_bin              = 25
	ClojureParserRULE_bign             = 26
	ClojureParserRULE_number           = 27
	ClojureParserRULE_character        = 28
	ClojureParserRULE_named_char       = 29
	ClojureParserRULE_any_char         = 30
	ClojureParserRULE_u_hex_quad       = 31
	ClojureParserRULE_nil              = 32
	ClojureParserRULE_keyword          = 33
	ClojureParserRULE_simple_keyword   = 34
	ClojureParserRULE_macro_keyword    = 35
	ClojureParserRULE_symbol           = 36
	ClojureParserRULE_simple_sym       = 37
	ClojureParserRULE_ns_symbol        = 38
	ClojureParserRULE_param_name       = 39
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClojureParserEOF, 0)
}

func (s *FileContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *FileContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *ClojureParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClojureParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ClojureParserT__0)|(1<<ClojureParserT__2)|(1<<ClojureParserT__4)|(1<<ClojureParserT__6)|(1<<ClojureParserT__7)|(1<<ClojureParserT__8)|(1<<ClojureParserT__9)|(1<<ClojureParserT__10)|(1<<ClojureParserT__11)|(1<<ClojureParserT__12)|(1<<ClojureParserT__13)|(1<<ClojureParserT__14)|(1<<ClojureParserT__15)|(1<<ClojureParserT__16)|(1<<ClojureParserT__17)|(1<<ClojureParserT__18)|(1<<ClojureParserT__19)|(1<<ClojureParserSTRING)|(1<<ClojureParserFLOAT)|(1<<ClojureParserHEX)|(1<<ClojureParserBIN)|(1<<ClojureParserLONG)|(1<<ClojureParserBIGN)|(1<<ClojureParserCHAR_U)|(1<<ClojureParserCHAR_NAMED)|(1<<ClojureParserCHAR_ANY)|(1<<ClojureParserNIL)|(1<<ClojureParserBOOLEAN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ClojureParserSYMBOL-32))|(1<<(ClojureParserNS_SYMBOL-32))|(1<<(ClojureParserPARAM_NAME-32)))) != 0) {
		{
			p.SetState(80)
			p.Form()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(ClojureParserEOF)
	}

	return localctx
}

// IFormContext is an interface to support dynamic dispatch.
type IFormContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormContext differentiates from other interfaces.
	IsFormContext()
}

type FormContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormContext() *FormContext {
	var p = new(FormContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_form
	return p
}

func (*FormContext) IsFormContext() {}

func NewFormContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormContext {
	var p = new(FormContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_form

	return p
}

func (s *FormContext) GetParser() antlr.Parser { return s.parser }

func (s *FormContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *FormContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *FormContext) Vector() IVectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *FormContext) Map() IMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapContext)
}

func (s *FormContext) Reader_macro() IReader_macroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReader_macroContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReader_macroContext)
}

func (s *FormContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterForm(s)
	}
}

func (s *FormContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitForm(s)
	}
}

func (p *ClojureParser) Form() (localctx IFormContext) {
	localctx = NewFormContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClojureParserRULE_form)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.List()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Vector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Map()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.Reader_macro()
		}

	}

	return localctx
}

// IFormsContext is an interface to support dynamic dispatch.
type IFormsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormsContext differentiates from other interfaces.
	IsFormsContext()
}

type FormsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormsContext() *FormsContext {
	var p = new(FormsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_forms
	return p
}

func (*FormsContext) IsFormsContext() {}

func NewFormsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormsContext {
	var p = new(FormsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_forms

	return p
}

func (s *FormsContext) GetParser() antlr.Parser { return s.parser }

func (s *FormsContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *FormsContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *FormsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterForms(s)
	}
}

func (s *FormsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitForms(s)
	}
}

func (p *ClojureParser) Forms() (localctx IFormsContext) {
	localctx = NewFormsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ClojureParserRULE_forms)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ClojureParserT__0)|(1<<ClojureParserT__2)|(1<<ClojureParserT__4)|(1<<ClojureParserT__6)|(1<<ClojureParserT__7)|(1<<ClojureParserT__8)|(1<<ClojureParserT__9)|(1<<ClojureParserT__10)|(1<<ClojureParserT__11)|(1<<ClojureParserT__12)|(1<<ClojureParserT__13)|(1<<ClojureParserT__14)|(1<<ClojureParserT__15)|(1<<ClojureParserT__16)|(1<<ClojureParserT__17)|(1<<ClojureParserT__18)|(1<<ClojureParserT__19)|(1<<ClojureParserSTRING)|(1<<ClojureParserFLOAT)|(1<<ClojureParserHEX)|(1<<ClojureParserBIN)|(1<<ClojureParserLONG)|(1<<ClojureParserBIGN)|(1<<ClojureParserCHAR_U)|(1<<ClojureParserCHAR_NAMED)|(1<<ClojureParserCHAR_ANY)|(1<<ClojureParserNIL)|(1<<ClojureParserBOOLEAN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ClojureParserSYMBOL-32))|(1<<(ClojureParserNS_SYMBOL-32))|(1<<(ClojureParserPARAM_NAME-32)))) != 0) {
		{
			p.SetState(95)
			p.Form()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) Forms() IFormsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormsContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *ClojureParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ClojureParserRULE_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(ClojureParserT__0)
	}
	{
		p.SetState(102)
		p.Forms()
	}
	{
		p.SetState(103)
		p.Match(ClojureParserT__1)
	}

	return localctx
}

// IVectorContext is an interface to support dynamic dispatch.
type IVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVectorContext differentiates from other interfaces.
	IsVectorContext()
}

type VectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorContext() *VectorContext {
	var p = new(VectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_vector
	return p
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_vector

	return p
}

func (s *VectorContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorContext) Forms() IFormsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormsContext)
}

func (s *VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitVector(s)
	}
}

func (p *ClojureParser) Vector() (localctx IVectorContext) {
	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ClojureParserRULE_vector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(ClojureParserT__2)
	}
	{
		p.SetState(106)
		p.Forms()
	}
	{
		p.SetState(107)
		p.Match(ClojureParserT__3)
	}

	return localctx
}

// IMapContext is an interface to support dynamic dispatch.
type IMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapContext differentiates from other interfaces.
	IsMapContext()
}

type MapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapContext() *MapContext {
	var p = new(MapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_map
	return p
}

func (*MapContext) IsMapContext() {}

func NewMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapContext {
	var p = new(MapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_map

	return p
}

func (s *MapContext) GetParser() antlr.Parser { return s.parser }

func (s *MapContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *MapContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *MapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterMap(s)
	}
}

func (s *MapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitMap(s)
	}
}

func (p *ClojureParser) Map() (localctx IMapContext) {
	localctx = NewMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ClojureParserRULE_map)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(ClojureParserT__4)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ClojureParserT__0)|(1<<ClojureParserT__2)|(1<<ClojureParserT__4)|(1<<ClojureParserT__6)|(1<<ClojureParserT__7)|(1<<ClojureParserT__8)|(1<<ClojureParserT__9)|(1<<ClojureParserT__10)|(1<<ClojureParserT__11)|(1<<ClojureParserT__12)|(1<<ClojureParserT__13)|(1<<ClojureParserT__14)|(1<<ClojureParserT__15)|(1<<ClojureParserT__16)|(1<<ClojureParserT__17)|(1<<ClojureParserT__18)|(1<<ClojureParserT__19)|(1<<ClojureParserSTRING)|(1<<ClojureParserFLOAT)|(1<<ClojureParserHEX)|(1<<ClojureParserBIN)|(1<<ClojureParserLONG)|(1<<ClojureParserBIGN)|(1<<ClojureParserCHAR_U)|(1<<ClojureParserCHAR_NAMED)|(1<<ClojureParserCHAR_ANY)|(1<<ClojureParserNIL)|(1<<ClojureParserBOOLEAN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ClojureParserSYMBOL-32))|(1<<(ClojureParserNS_SYMBOL-32))|(1<<(ClojureParserPARAM_NAME-32)))) != 0) {
		{
			p.SetState(110)
			p.Form()
		}
		{
			p.SetState(111)
			p.Form()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(ClojureParserT__5)
	}

	return localctx
}

// ISetContext is an interface to support dynamic dispatch.
type ISetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetContext differentiates from other interfaces.
	IsSetContext()
}

type SetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetContext() *SetContext {
	var p = new(SetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_set
	return p
}

func (*SetContext) IsSetContext() {}

func NewSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetContext {
	var p = new(SetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_set

	return p
}

func (s *SetContext) GetParser() antlr.Parser { return s.parser }

func (s *SetContext) Forms() IFormsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormsContext)
}

func (s *SetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterSet(s)
	}
}

func (s *SetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitSet(s)
	}
}

func (p *ClojureParser) Set() (localctx ISetContext) {
	localctx = NewSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ClojureParserRULE_set)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(ClojureParserT__6)
	}
	{
		p.SetState(121)
		p.Forms()
	}
	{
		p.SetState(122)
		p.Match(ClojureParserT__5)
	}

	return localctx
}

// IReader_macroContext is an interface to support dynamic dispatch.
type IReader_macroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReader_macroContext differentiates from other interfaces.
	IsReader_macroContext()
}

type Reader_macroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReader_macroContext() *Reader_macroContext {
	var p = new(Reader_macroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_reader_macro
	return p
}

func (*Reader_macroContext) IsReader_macroContext() {}

func NewReader_macroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reader_macroContext {
	var p = new(Reader_macroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_reader_macro

	return p
}

func (s *Reader_macroContext) GetParser() antlr.Parser { return s.parser }

func (s *Reader_macroContext) Lambda() ILambdaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *Reader_macroContext) Meta_data() IMeta_dataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeta_dataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeta_dataContext)
}

func (s *Reader_macroContext) Regex() IRegexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegexContext)
}

func (s *Reader_macroContext) Var_quote() IVar_quoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_quoteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_quoteContext)
}

func (s *Reader_macroContext) Host_expr() IHost_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHost_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHost_exprContext)
}

func (s *Reader_macroContext) Set() ISetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetContext)
}

func (s *Reader_macroContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Reader_macroContext) Discard() IDiscardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDiscardContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDiscardContext)
}

func (s *Reader_macroContext) Dispatch() IDispatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDispatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDispatchContext)
}

func (s *Reader_macroContext) Deref() IDerefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDerefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDerefContext)
}

func (s *Reader_macroContext) Quote() IQuoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuoteContext)
}

func (s *Reader_macroContext) Backtick() IBacktickContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBacktickContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBacktickContext)
}

func (s *Reader_macroContext) Unquote() IUnquoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnquoteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnquoteContext)
}

func (s *Reader_macroContext) Unquote_splicing() IUnquote_splicingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnquote_splicingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnquote_splicingContext)
}

func (s *Reader_macroContext) Gensym() IGensymContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGensymContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGensymContext)
}

func (s *Reader_macroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reader_macroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Reader_macroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterReader_macro(s)
	}
}

func (s *Reader_macroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitReader_macro(s)
	}
}

func (p *ClojureParser) Reader_macro() (localctx IReader_macroContext) {
	localctx = NewReader_macroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ClojureParserRULE_reader_macro)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Lambda()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Meta_data()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Regex()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Var_quote()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.Host_expr()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Set()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(130)
			p.Tag()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(131)
			p.Discard()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(132)
			p.Dispatch()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(133)
			p.Deref()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(134)
			p.Quote()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(135)
			p.Backtick()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(136)
			p.Unquote()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(137)
			p.Unquote_splicing()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(138)
			p.Gensym()
		}

	}

	return localctx
}

// IQuoteContext is an interface to support dynamic dispatch.
type IQuoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuoteContext differentiates from other interfaces.
	IsQuoteContext()
}

type QuoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoteContext() *QuoteContext {
	var p = new(QuoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_quote
	return p
}

func (*QuoteContext) IsQuoteContext() {}

func NewQuoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuoteContext {
	var p = new(QuoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_quote

	return p
}

func (s *QuoteContext) GetParser() antlr.Parser { return s.parser }

func (s *QuoteContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *QuoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterQuote(s)
	}
}

func (s *QuoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitQuote(s)
	}
}

func (p *ClojureParser) Quote() (localctx IQuoteContext) {
	localctx = NewQuoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ClojureParserRULE_quote)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(ClojureParserT__7)
	}
	{
		p.SetState(142)
		p.Form()
	}

	return localctx
}

// IBacktickContext is an interface to support dynamic dispatch.
type IBacktickContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBacktickContext differentiates from other interfaces.
	IsBacktickContext()
}

type BacktickContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBacktickContext() *BacktickContext {
	var p = new(BacktickContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_backtick
	return p
}

func (*BacktickContext) IsBacktickContext() {}

func NewBacktickContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BacktickContext {
	var p = new(BacktickContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_backtick

	return p
}

func (s *BacktickContext) GetParser() antlr.Parser { return s.parser }

func (s *BacktickContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *BacktickContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BacktickContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BacktickContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterBacktick(s)
	}
}

func (s *BacktickContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitBacktick(s)
	}
}

func (p *ClojureParser) Backtick() (localctx IBacktickContext) {
	localctx = NewBacktickContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ClojureParserRULE_backtick)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(ClojureParserT__8)
	}
	{
		p.SetState(145)
		p.Form()
	}

	return localctx
}

// IUnquoteContext is an interface to support dynamic dispatch.
type IUnquoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnquoteContext differentiates from other interfaces.
	IsUnquoteContext()
}

type UnquoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquoteContext() *UnquoteContext {
	var p = new(UnquoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_unquote
	return p
}

func (*UnquoteContext) IsUnquoteContext() {}

func NewUnquoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnquoteContext {
	var p = new(UnquoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_unquote

	return p
}

func (s *UnquoteContext) GetParser() antlr.Parser { return s.parser }

func (s *UnquoteContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *UnquoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnquoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnquoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterUnquote(s)
	}
}

func (s *UnquoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitUnquote(s)
	}
}

func (p *ClojureParser) Unquote() (localctx IUnquoteContext) {
	localctx = NewUnquoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ClojureParserRULE_unquote)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(ClojureParserT__9)
	}
	{
		p.SetState(148)
		p.Form()
	}

	return localctx
}

// IUnquote_splicingContext is an interface to support dynamic dispatch.
type IUnquote_splicingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnquote_splicingContext differentiates from other interfaces.
	IsUnquote_splicingContext()
}

type Unquote_splicingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquote_splicingContext() *Unquote_splicingContext {
	var p = new(Unquote_splicingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_unquote_splicing
	return p
}

func (*Unquote_splicingContext) IsUnquote_splicingContext() {}

func NewUnquote_splicingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unquote_splicingContext {
	var p = new(Unquote_splicingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_unquote_splicing

	return p
}

func (s *Unquote_splicingContext) GetParser() antlr.Parser { return s.parser }

func (s *Unquote_splicingContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *Unquote_splicingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unquote_splicingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unquote_splicingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterUnquote_splicing(s)
	}
}

func (s *Unquote_splicingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitUnquote_splicing(s)
	}
}

func (p *ClojureParser) Unquote_splicing() (localctx IUnquote_splicingContext) {
	localctx = NewUnquote_splicingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ClojureParserRULE_unquote_splicing)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(ClojureParserT__10)
	}
	{
		p.SetState(151)
		p.Form()
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *TagContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *ClojureParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ClojureParserRULE_tag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(ClojureParserT__11)
	}
	{
		p.SetState(154)
		p.Form()
	}
	{
		p.SetState(155)
		p.Form()
	}

	return localctx
}

// IDerefContext is an interface to support dynamic dispatch.
type IDerefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDerefContext differentiates from other interfaces.
	IsDerefContext()
}

type DerefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDerefContext() *DerefContext {
	var p = new(DerefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_deref
	return p
}

func (*DerefContext) IsDerefContext() {}

func NewDerefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DerefContext {
	var p = new(DerefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_deref

	return p
}

func (s *DerefContext) GetParser() antlr.Parser { return s.parser }

func (s *DerefContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *DerefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DerefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterDeref(s)
	}
}

func (s *DerefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitDeref(s)
	}
}

func (p *ClojureParser) Deref() (localctx IDerefContext) {
	localctx = NewDerefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ClojureParserRULE_deref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(ClojureParserT__12)
	}
	{
		p.SetState(158)
		p.Form()
	}

	return localctx
}

// IGensymContext is an interface to support dynamic dispatch.
type IGensymContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGensymContext differentiates from other interfaces.
	IsGensymContext()
}

type GensymContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGensymContext() *GensymContext {
	var p = new(GensymContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_gensym
	return p
}

func (*GensymContext) IsGensymContext() {}

func NewGensymContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GensymContext {
	var p = new(GensymContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_gensym

	return p
}

func (s *GensymContext) GetParser() antlr.Parser { return s.parser }

func (s *GensymContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ClojureParserSYMBOL, 0)
}

func (s *GensymContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GensymContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GensymContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterGensym(s)
	}
}

func (s *GensymContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitGensym(s)
	}
}

func (p *ClojureParser) Gensym() (localctx IGensymContext) {
	localctx = NewGensymContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ClojureParserRULE_gensym)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(ClojureParserSYMBOL)
	}
	{
		p.SetState(161)
		p.Match(ClojureParserT__13)
	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *LambdaContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (p *ClojureParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ClojureParserRULE_lambda)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(ClojureParserT__14)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ClojureParserT__0)|(1<<ClojureParserT__2)|(1<<ClojureParserT__4)|(1<<ClojureParserT__6)|(1<<ClojureParserT__7)|(1<<ClojureParserT__8)|(1<<ClojureParserT__9)|(1<<ClojureParserT__10)|(1<<ClojureParserT__11)|(1<<ClojureParserT__12)|(1<<ClojureParserT__13)|(1<<ClojureParserT__14)|(1<<ClojureParserT__15)|(1<<ClojureParserT__16)|(1<<ClojureParserT__17)|(1<<ClojureParserT__18)|(1<<ClojureParserT__19)|(1<<ClojureParserSTRING)|(1<<ClojureParserFLOAT)|(1<<ClojureParserHEX)|(1<<ClojureParserBIN)|(1<<ClojureParserLONG)|(1<<ClojureParserBIGN)|(1<<ClojureParserCHAR_U)|(1<<ClojureParserCHAR_NAMED)|(1<<ClojureParserCHAR_ANY)|(1<<ClojureParserNIL)|(1<<ClojureParserBOOLEAN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ClojureParserSYMBOL-32))|(1<<(ClojureParserNS_SYMBOL-32))|(1<<(ClojureParserPARAM_NAME-32)))) != 0) {
		{
			p.SetState(164)
			p.Form()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
		p.Match(ClojureParserT__1)
	}

	return localctx
}

// IMeta_dataContext is an interface to support dynamic dispatch.
type IMeta_dataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeta_dataContext differentiates from other interfaces.
	IsMeta_dataContext()
}

type Meta_dataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeta_dataContext() *Meta_dataContext {
	var p = new(Meta_dataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_meta_data
	return p
}

func (*Meta_dataContext) IsMeta_dataContext() {}

func NewMeta_dataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Meta_dataContext {
	var p = new(Meta_dataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_meta_data

	return p
}

func (s *Meta_dataContext) GetParser() antlr.Parser { return s.parser }

func (s *Meta_dataContext) Map() IMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapContext)
}

func (s *Meta_dataContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *Meta_dataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Meta_dataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Meta_dataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterMeta_data(s)
	}
}

func (s *Meta_dataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitMeta_data(s)
	}
}

func (p *ClojureParser) Meta_data() (localctx IMeta_dataContext) {
	localctx = NewMeta_dataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ClojureParserRULE_meta_data)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(ClojureParserT__15)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(173)
			p.Map()
		}
		{
			p.SetState(174)
			p.Form()
		}

	case 2:
		{
			p.SetState(176)
			p.Form()
		}

	}

	return localctx
}

// IVar_quoteContext is an interface to support dynamic dispatch.
type IVar_quoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_quoteContext differentiates from other interfaces.
	IsVar_quoteContext()
}

type Var_quoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_quoteContext() *Var_quoteContext {
	var p = new(Var_quoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_var_quote
	return p
}

func (*Var_quoteContext) IsVar_quoteContext() {}

func NewVar_quoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_quoteContext {
	var p = new(Var_quoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_var_quote

	return p
}

func (s *Var_quoteContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_quoteContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Var_quoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_quoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_quoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterVar_quote(s)
	}
}

func (s *Var_quoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitVar_quote(s)
	}
}

func (p *ClojureParser) Var_quote() (localctx IVar_quoteContext) {
	localctx = NewVar_quoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ClojureParserRULE_var_quote)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(ClojureParserT__16)
	}
	{
		p.SetState(180)
		p.Symbol()
	}

	return localctx
}

// IHost_exprContext is an interface to support dynamic dispatch.
type IHost_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHost_exprContext differentiates from other interfaces.
	IsHost_exprContext()
}

type Host_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHost_exprContext() *Host_exprContext {
	var p = new(Host_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_host_expr
	return p
}

func (*Host_exprContext) IsHost_exprContext() {}

func NewHost_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Host_exprContext {
	var p = new(Host_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_host_expr

	return p
}

func (s *Host_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Host_exprContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *Host_exprContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *Host_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Host_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Host_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterHost_expr(s)
	}
}

func (s *Host_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitHost_expr(s)
	}
}

func (p *ClojureParser) Host_expr() (localctx IHost_exprContext) {
	localctx = NewHost_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ClojureParserRULE_host_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(ClojureParserT__17)
	}
	{
		p.SetState(183)
		p.Form()
	}
	{
		p.SetState(184)
		p.Form()
	}

	return localctx
}

// IDiscardContext is an interface to support dynamic dispatch.
type IDiscardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiscardContext differentiates from other interfaces.
	IsDiscardContext()
}

type DiscardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiscardContext() *DiscardContext {
	var p = new(DiscardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_discard
	return p
}

func (*DiscardContext) IsDiscardContext() {}

func NewDiscardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiscardContext {
	var p = new(DiscardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_discard

	return p
}

func (s *DiscardContext) GetParser() antlr.Parser { return s.parser }

func (s *DiscardContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *DiscardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiscardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiscardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterDiscard(s)
	}
}

func (s *DiscardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitDiscard(s)
	}
}

func (p *ClojureParser) Discard() (localctx IDiscardContext) {
	localctx = NewDiscardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ClojureParserRULE_discard)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(ClojureParserT__18)
	}
	{
		p.SetState(187)
		p.Form()
	}

	return localctx
}

// IDispatchContext is an interface to support dynamic dispatch.
type IDispatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDispatchContext differentiates from other interfaces.
	IsDispatchContext()
}

type DispatchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDispatchContext() *DispatchContext {
	var p = new(DispatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_dispatch
	return p
}

func (*DispatchContext) IsDispatchContext() {}

func NewDispatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DispatchContext {
	var p = new(DispatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_dispatch

	return p
}

func (s *DispatchContext) GetParser() antlr.Parser { return s.parser }

func (s *DispatchContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *DispatchContext) Form() IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *DispatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DispatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DispatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterDispatch(s)
	}
}

func (s *DispatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitDispatch(s)
	}
}

func (p *ClojureParser) Dispatch() (localctx IDispatchContext) {
	localctx = NewDispatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ClojureParserRULE_dispatch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(ClojureParserT__13)
	}
	{
		p.SetState(190)
		p.Symbol()
	}
	{
		p.SetState(191)
		p.Form()
	}

	return localctx
}

// IRegexContext is an interface to support dynamic dispatch.
type IRegexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexContext differentiates from other interfaces.
	IsRegexContext()
}

type RegexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexContext() *RegexContext {
	var p = new(RegexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_regex
	return p
}

func (*RegexContext) IsRegexContext() {}

func NewRegexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexContext {
	var p = new(RegexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_regex

	return p
}

func (s *RegexContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexContext) String() IStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *RegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterRegex(s)
	}
}

func (s *RegexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitRegex(s)
	}
}

func (p *ClojureParser) Regex() (localctx IRegexContext) {
	localctx = NewRegexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ClojureParserRULE_regex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(ClojureParserT__13)
	}
	{
		p.SetState(194)
		p.String()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) String() IStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *LiteralContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *LiteralContext) Character() ICharacterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterContext)
}

func (s *LiteralContext) Nil() INilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilContext)
}

func (s *LiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ClojureParserBOOLEAN, 0)
}

func (s *LiteralContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *LiteralContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *LiteralContext) Param_name() IParam_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParam_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParam_nameContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *ClojureParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ClojureParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ClojureParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.String()
		}

	case ClojureParserFLOAT, ClojureParserHEX, ClojureParserBIN, ClojureParserLONG, ClojureParserBIGN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Number()
		}

	case ClojureParserCHAR_U, ClojureParserCHAR_NAMED, ClojureParserCHAR_ANY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(198)
			p.Character()
		}

	case ClojureParserNIL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)
			p.Nil()
		}

	case ClojureParserBOOLEAN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(200)
			p.Match(ClojureParserBOOLEAN)
		}

	case ClojureParserT__19:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(201)
			p.Keyword()
		}

	case ClojureParserSYMBOL, ClojureParserNS_SYMBOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(202)
			p.Symbol()
		}

	case ClojureParserPARAM_NAME:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(203)
			p.Param_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_string
	return p
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClojureParserSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *ClojureParser) String() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ClojureParserRULE_string)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(ClojureParserSTRING)
	}

	return localctx
}

// IHexContext is an interface to support dynamic dispatch.
type IHexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexContext differentiates from other interfaces.
	IsHexContext()
}

type HexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexContext() *HexContext {
	var p = new(HexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_hex
	return p
}

func (*HexContext) IsHexContext() {}

func NewHexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexContext {
	var p = new(HexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_hex

	return p
}

func (s *HexContext) GetParser() antlr.Parser { return s.parser }

func (s *HexContext) HEX() antlr.TerminalNode {
	return s.GetToken(ClojureParserHEX, 0)
}

func (s *HexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterHex(s)
	}
}

func (s *HexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitHex(s)
	}
}

func (p *ClojureParser) Hex() (localctx IHexContext) {
	localctx = NewHexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ClojureParserRULE_hex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(ClojureParserHEX)
	}

	return localctx
}

// IBinContext is an interface to support dynamic dispatch.
type IBinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinContext differentiates from other interfaces.
	IsBinContext()
}

type BinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinContext() *BinContext {
	var p = new(BinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_bin
	return p
}

func (*BinContext) IsBinContext() {}

func NewBinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinContext {
	var p = new(BinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_bin

	return p
}

func (s *BinContext) GetParser() antlr.Parser { return s.parser }

func (s *BinContext) BIN() antlr.TerminalNode {
	return s.GetToken(ClojureParserBIN, 0)
}

func (s *BinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterBin(s)
	}
}

func (s *BinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitBin(s)
	}
}

func (p *ClojureParser) Bin() (localctx IBinContext) {
	localctx = NewBinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ClojureParserRULE_bin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(ClojureParserBIN)
	}

	return localctx
}

// IBignContext is an interface to support dynamic dispatch.
type IBignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBignContext differentiates from other interfaces.
	IsBignContext()
}

type BignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBignContext() *BignContext {
	var p = new(BignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_bign
	return p
}

func (*BignContext) IsBignContext() {}

func NewBignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BignContext {
	var p = new(BignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_bign

	return p
}

func (s *BignContext) GetParser() antlr.Parser { return s.parser }

func (s *BignContext) BIGN() antlr.TerminalNode {
	return s.GetToken(ClojureParserBIGN, 0)
}

func (s *BignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterBign(s)
	}
}

func (s *BignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitBign(s)
	}
}

func (p *ClojureParser) Bign() (localctx IBignContext) {
	localctx = NewBignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ClojureParserRULE_bign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(ClojureParserBIGN)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ClojureParserFLOAT, 0)
}

func (s *NumberContext) Hex() IHexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHexContext)
}

func (s *NumberContext) Bin() IBinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinContext)
}

func (s *NumberContext) Bign() IBignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBignContext)
}

func (s *NumberContext) LONG() antlr.TerminalNode {
	return s.GetToken(ClojureParserLONG, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *ClojureParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ClojureParserRULE_number)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ClojureParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(ClojureParserFLOAT)
		}

	case ClojureParserHEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Hex()
		}

	case ClojureParserBIN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Bin()
		}

	case ClojureParserBIGN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.Bign()
		}

	case ClojureParserLONG:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
			p.Match(ClojureParserLONG)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICharacterContext is an interface to support dynamic dispatch.
type ICharacterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterContext differentiates from other interfaces.
	IsCharacterContext()
}

type CharacterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterContext() *CharacterContext {
	var p = new(CharacterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_character
	return p
}

func (*CharacterContext) IsCharacterContext() {}

func NewCharacterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterContext {
	var p = new(CharacterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_character

	return p
}

func (s *CharacterContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterContext) Named_char() INamed_charContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamed_charContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamed_charContext)
}

func (s *CharacterContext) U_hex_quad() IU_hex_quadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IU_hex_quadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IU_hex_quadContext)
}

func (s *CharacterContext) Any_char() IAny_charContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_charContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_charContext)
}

func (s *CharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterCharacter(s)
	}
}

func (s *CharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitCharacter(s)
	}
}

func (p *ClojureParser) Character() (localctx ICharacterContext) {
	localctx = NewCharacterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ClojureParserRULE_character)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(224)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ClojureParserCHAR_NAMED:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Named_char()
		}

	case ClojureParserCHAR_U:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.U_hex_quad()
		}

	case ClojureParserCHAR_ANY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Any_char()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamed_charContext is an interface to support dynamic dispatch.
type INamed_charContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamed_charContext differentiates from other interfaces.
	IsNamed_charContext()
}

type Named_charContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamed_charContext() *Named_charContext {
	var p = new(Named_charContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_named_char
	return p
}

func (*Named_charContext) IsNamed_charContext() {}

func NewNamed_charContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Named_charContext {
	var p = new(Named_charContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_named_char

	return p
}

func (s *Named_charContext) GetParser() antlr.Parser { return s.parser }

func (s *Named_charContext) CHAR_NAMED() antlr.TerminalNode {
	return s.GetToken(ClojureParserCHAR_NAMED, 0)
}

func (s *Named_charContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Named_charContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Named_charContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterNamed_char(s)
	}
}

func (s *Named_charContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitNamed_char(s)
	}
}

func (p *ClojureParser) Named_char() (localctx INamed_charContext) {
	localctx = NewNamed_charContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ClojureParserRULE_named_char)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(ClojureParserCHAR_NAMED)
	}

	return localctx
}

// IAny_charContext is an interface to support dynamic dispatch.
type IAny_charContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAny_charContext differentiates from other interfaces.
	IsAny_charContext()
}

type Any_charContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAny_charContext() *Any_charContext {
	var p = new(Any_charContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_any_char
	return p
}

func (*Any_charContext) IsAny_charContext() {}

func NewAny_charContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Any_charContext {
	var p = new(Any_charContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_any_char

	return p
}

func (s *Any_charContext) GetParser() antlr.Parser { return s.parser }

func (s *Any_charContext) CHAR_ANY() antlr.TerminalNode {
	return s.GetToken(ClojureParserCHAR_ANY, 0)
}

func (s *Any_charContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Any_charContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Any_charContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterAny_char(s)
	}
}

func (s *Any_charContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitAny_char(s)
	}
}

func (p *ClojureParser) Any_char() (localctx IAny_charContext) {
	localctx = NewAny_charContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ClojureParserRULE_any_char)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(ClojureParserCHAR_ANY)
	}

	return localctx
}

// IU_hex_quadContext is an interface to support dynamic dispatch.
type IU_hex_quadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsU_hex_quadContext differentiates from other interfaces.
	IsU_hex_quadContext()
}

type U_hex_quadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyU_hex_quadContext() *U_hex_quadContext {
	var p = new(U_hex_quadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_u_hex_quad
	return p
}

func (*U_hex_quadContext) IsU_hex_quadContext() {}

func NewU_hex_quadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *U_hex_quadContext {
	var p = new(U_hex_quadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_u_hex_quad

	return p
}

func (s *U_hex_quadContext) GetParser() antlr.Parser { return s.parser }

func (s *U_hex_quadContext) CHAR_U() antlr.TerminalNode {
	return s.GetToken(ClojureParserCHAR_U, 0)
}

func (s *U_hex_quadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *U_hex_quadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *U_hex_quadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterU_hex_quad(s)
	}
}

func (s *U_hex_quadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitU_hex_quad(s)
	}
}

func (p *ClojureParser) U_hex_quad() (localctx IU_hex_quadContext) {
	localctx = NewU_hex_quadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ClojureParserRULE_u_hex_quad)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(ClojureParserCHAR_U)
	}

	return localctx
}

// INilContext is an interface to support dynamic dispatch.
type INilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilContext differentiates from other interfaces.
	IsNilContext()
}

type NilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilContext() *NilContext {
	var p = new(NilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_nil
	return p
}

func (*NilContext) IsNilContext() {}

func NewNilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilContext {
	var p = new(NilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_nil

	return p
}

func (s *NilContext) GetParser() antlr.Parser { return s.parser }

func (s *NilContext) NIL() antlr.TerminalNode {
	return s.GetToken(ClojureParserNIL, 0)
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterNil(s)
	}
}

func (s *NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitNil(s)
	}
}

func (p *ClojureParser) Nil() (localctx INilContext) {
	localctx = NewNilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ClojureParserRULE_nil)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(ClojureParserNIL)
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Macro_keyword() IMacro_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMacro_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMacro_keywordContext)
}

func (s *KeywordContext) Simple_keyword() ISimple_keywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_keywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_keywordContext)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *ClojureParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ClojureParserRULE_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Macro_keyword()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Simple_keyword()
		}

	}

	return localctx
}

// ISimple_keywordContext is an interface to support dynamic dispatch.
type ISimple_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_keywordContext differentiates from other interfaces.
	IsSimple_keywordContext()
}

type Simple_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_keywordContext() *Simple_keywordContext {
	var p = new(Simple_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_simple_keyword
	return p
}

func (*Simple_keywordContext) IsSimple_keywordContext() {}

func NewSimple_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_keywordContext {
	var p = new(Simple_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_simple_keyword

	return p
}

func (s *Simple_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_keywordContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Simple_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterSimple_keyword(s)
	}
}

func (s *Simple_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitSimple_keyword(s)
	}
}

func (p *ClojureParser) Simple_keyword() (localctx ISimple_keywordContext) {
	localctx = NewSimple_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ClojureParserRULE_simple_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(ClojureParserT__19)
	}
	{
		p.SetState(239)
		p.Symbol()
	}

	return localctx
}

// IMacro_keywordContext is an interface to support dynamic dispatch.
type IMacro_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacro_keywordContext differentiates from other interfaces.
	IsMacro_keywordContext()
}

type Macro_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacro_keywordContext() *Macro_keywordContext {
	var p = new(Macro_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_macro_keyword
	return p
}

func (*Macro_keywordContext) IsMacro_keywordContext() {}

func NewMacro_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Macro_keywordContext {
	var p = new(Macro_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_macro_keyword

	return p
}

func (s *Macro_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Macro_keywordContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Macro_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Macro_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Macro_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterMacro_keyword(s)
	}
}

func (s *Macro_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitMacro_keyword(s)
	}
}

func (p *ClojureParser) Macro_keyword() (localctx IMacro_keywordContext) {
	localctx = NewMacro_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ClojureParserRULE_macro_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(ClojureParserT__19)
	}
	{
		p.SetState(242)
		p.Match(ClojureParserT__19)
	}
	{
		p.SetState(243)
		p.Symbol()
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) Ns_symbol() INs_symbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INs_symbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INs_symbolContext)
}

func (s *SymbolContext) Simple_sym() ISimple_symContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_symContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_symContext)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *ClojureParser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ClojureParserRULE_symbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ClojureParserNS_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Ns_symbol()
		}

	case ClojureParserSYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Simple_sym()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimple_symContext is an interface to support dynamic dispatch.
type ISimple_symContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_symContext differentiates from other interfaces.
	IsSimple_symContext()
}

type Simple_symContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_symContext() *Simple_symContext {
	var p = new(Simple_symContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_simple_sym
	return p
}

func (*Simple_symContext) IsSimple_symContext() {}

func NewSimple_symContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_symContext {
	var p = new(Simple_symContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_simple_sym

	return p
}

func (s *Simple_symContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_symContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ClojureParserSYMBOL, 0)
}

func (s *Simple_symContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_symContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_symContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterSimple_sym(s)
	}
}

func (s *Simple_symContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitSimple_sym(s)
	}
}

func (p *ClojureParser) Simple_sym() (localctx ISimple_symContext) {
	localctx = NewSimple_symContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ClojureParserRULE_simple_sym)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(ClojureParserSYMBOL)
	}

	return localctx
}

// INs_symbolContext is an interface to support dynamic dispatch.
type INs_symbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNs_symbolContext differentiates from other interfaces.
	IsNs_symbolContext()
}

type Ns_symbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNs_symbolContext() *Ns_symbolContext {
	var p = new(Ns_symbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_ns_symbol
	return p
}

func (*Ns_symbolContext) IsNs_symbolContext() {}

func NewNs_symbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ns_symbolContext {
	var p = new(Ns_symbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_ns_symbol

	return p
}

func (s *Ns_symbolContext) GetParser() antlr.Parser { return s.parser }

func (s *Ns_symbolContext) NS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(ClojureParserNS_SYMBOL, 0)
}

func (s *Ns_symbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ns_symbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ns_symbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterNs_symbol(s)
	}
}

func (s *Ns_symbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitNs_symbol(s)
	}
}

func (p *ClojureParser) Ns_symbol() (localctx INs_symbolContext) {
	localctx = NewNs_symbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ClojureParserRULE_ns_symbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(ClojureParserNS_SYMBOL)
	}

	return localctx
}

// IParam_nameContext is an interface to support dynamic dispatch.
type IParam_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParam_nameContext differentiates from other interfaces.
	IsParam_nameContext()
}

type Param_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_nameContext() *Param_nameContext {
	var p = new(Param_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClojureParserRULE_param_name
	return p
}

func (*Param_nameContext) IsParam_nameContext() {}

func NewParam_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_nameContext {
	var p = new(Param_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClojureParserRULE_param_name

	return p
}

func (s *Param_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_nameContext) PARAM_NAME() antlr.TerminalNode {
	return s.GetToken(ClojureParserPARAM_NAME, 0)
}

func (s *Param_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.EnterParam_name(s)
	}
}

func (s *Param_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClojureListener); ok {
		listenerT.ExitParam_name(s)
	}
}

func (p *ClojureParser) Param_name() (localctx IParam_nameContext) {
	localctx = NewParam_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ClojureParserRULE_param_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(ClojureParserPARAM_NAME)
	}

	return localctx
}
