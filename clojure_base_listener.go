// Code generated from Clojure.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Clojure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseClojureListener is a complete listener for a parse tree produced by ClojureParser.
type BaseClojureListener struct{}

var _ ClojureListener = &BaseClojureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClojureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClojureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClojureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClojureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseClojureListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseClojureListener) ExitFile(ctx *FileContext) {}

// EnterForm is called when production form is entered.
func (s *BaseClojureListener) EnterForm(ctx *FormContext) {}

// ExitForm is called when production form is exited.
func (s *BaseClojureListener) ExitForm(ctx *FormContext) {}

// EnterForms is called when production forms is entered.
func (s *BaseClojureListener) EnterForms(ctx *FormsContext) {}

// ExitForms is called when production forms is exited.
func (s *BaseClojureListener) ExitForms(ctx *FormsContext) {}

// EnterList is called when production list is entered.
func (s *BaseClojureListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseClojureListener) ExitList(ctx *ListContext) {}

// EnterVector is called when production vector is entered.
func (s *BaseClojureListener) EnterVector(ctx *VectorContext) {}

// ExitVector is called when production vector is exited.
func (s *BaseClojureListener) ExitVector(ctx *VectorContext) {}

// EnterMap is called when production map is entered.
func (s *BaseClojureListener) EnterMap(ctx *MapContext) {}

// ExitMap is called when production map is exited.
func (s *BaseClojureListener) ExitMap(ctx *MapContext) {}

// EnterSet is called when production set is entered.
func (s *BaseClojureListener) EnterSet(ctx *SetContext) {}

// ExitSet is called when production set is exited.
func (s *BaseClojureListener) ExitSet(ctx *SetContext) {}

// EnterReader_macro is called when production reader_macro is entered.
func (s *BaseClojureListener) EnterReader_macro(ctx *Reader_macroContext) {}

// ExitReader_macro is called when production reader_macro is exited.
func (s *BaseClojureListener) ExitReader_macro(ctx *Reader_macroContext) {}

// EnterQuote is called when production quote is entered.
func (s *BaseClojureListener) EnterQuote(ctx *QuoteContext) {}

// ExitQuote is called when production quote is exited.
func (s *BaseClojureListener) ExitQuote(ctx *QuoteContext) {}

// EnterBacktick is called when production backtick is entered.
func (s *BaseClojureListener) EnterBacktick(ctx *BacktickContext) {}

// ExitBacktick is called when production backtick is exited.
func (s *BaseClojureListener) ExitBacktick(ctx *BacktickContext) {}

// EnterUnquote is called when production unquote is entered.
func (s *BaseClojureListener) EnterUnquote(ctx *UnquoteContext) {}

// ExitUnquote is called when production unquote is exited.
func (s *BaseClojureListener) ExitUnquote(ctx *UnquoteContext) {}

// EnterUnquote_splicing is called when production unquote_splicing is entered.
func (s *BaseClojureListener) EnterUnquote_splicing(ctx *Unquote_splicingContext) {}

// ExitUnquote_splicing is called when production unquote_splicing is exited.
func (s *BaseClojureListener) ExitUnquote_splicing(ctx *Unquote_splicingContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseClojureListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseClojureListener) ExitTag(ctx *TagContext) {}

// EnterDeref is called when production deref is entered.
func (s *BaseClojureListener) EnterDeref(ctx *DerefContext) {}

// ExitDeref is called when production deref is exited.
func (s *BaseClojureListener) ExitDeref(ctx *DerefContext) {}

// EnterGensym is called when production gensym is entered.
func (s *BaseClojureListener) EnterGensym(ctx *GensymContext) {}

// ExitGensym is called when production gensym is exited.
func (s *BaseClojureListener) ExitGensym(ctx *GensymContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseClojureListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseClojureListener) ExitLambda(ctx *LambdaContext) {}

// EnterMeta_data is called when production meta_data is entered.
func (s *BaseClojureListener) EnterMeta_data(ctx *Meta_dataContext) {}

// ExitMeta_data is called when production meta_data is exited.
func (s *BaseClojureListener) ExitMeta_data(ctx *Meta_dataContext) {}

// EnterVar_quote is called when production var_quote is entered.
func (s *BaseClojureListener) EnterVar_quote(ctx *Var_quoteContext) {}

// ExitVar_quote is called when production var_quote is exited.
func (s *BaseClojureListener) ExitVar_quote(ctx *Var_quoteContext) {}

// EnterHost_expr is called when production host_expr is entered.
func (s *BaseClojureListener) EnterHost_expr(ctx *Host_exprContext) {}

// ExitHost_expr is called when production host_expr is exited.
func (s *BaseClojureListener) ExitHost_expr(ctx *Host_exprContext) {}

// EnterDiscard is called when production discard is entered.
func (s *BaseClojureListener) EnterDiscard(ctx *DiscardContext) {}

// ExitDiscard is called when production discard is exited.
func (s *BaseClojureListener) ExitDiscard(ctx *DiscardContext) {}

// EnterDispatch is called when production dispatch is entered.
func (s *BaseClojureListener) EnterDispatch(ctx *DispatchContext) {}

// ExitDispatch is called when production dispatch is exited.
func (s *BaseClojureListener) ExitDispatch(ctx *DispatchContext) {}

// EnterRegex is called when production regex is entered.
func (s *BaseClojureListener) EnterRegex(ctx *RegexContext) {}

// ExitRegex is called when production regex is exited.
func (s *BaseClojureListener) ExitRegex(ctx *RegexContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseClojureListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseClojureListener) ExitLiteral(ctx *LiteralContext) {}

// EnterString is called when production string is entered.
func (s *BaseClojureListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseClojureListener) ExitString(ctx *StringContext) {}

// EnterHex is called when production hex is entered.
func (s *BaseClojureListener) EnterHex(ctx *HexContext) {}

// ExitHex is called when production hex is exited.
func (s *BaseClojureListener) ExitHex(ctx *HexContext) {}

// EnterBin is called when production bin is entered.
func (s *BaseClojureListener) EnterBin(ctx *BinContext) {}

// ExitBin is called when production bin is exited.
func (s *BaseClojureListener) ExitBin(ctx *BinContext) {}

// EnterBign is called when production bign is entered.
func (s *BaseClojureListener) EnterBign(ctx *BignContext) {}

// ExitBign is called when production bign is exited.
func (s *BaseClojureListener) ExitBign(ctx *BignContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseClojureListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseClojureListener) ExitNumber(ctx *NumberContext) {}

// EnterCharacter is called when production character is entered.
func (s *BaseClojureListener) EnterCharacter(ctx *CharacterContext) {}

// ExitCharacter is called when production character is exited.
func (s *BaseClojureListener) ExitCharacter(ctx *CharacterContext) {}

// EnterNamed_char is called when production named_char is entered.
func (s *BaseClojureListener) EnterNamed_char(ctx *Named_charContext) {}

// ExitNamed_char is called when production named_char is exited.
func (s *BaseClojureListener) ExitNamed_char(ctx *Named_charContext) {}

// EnterAny_char is called when production any_char is entered.
func (s *BaseClojureListener) EnterAny_char(ctx *Any_charContext) {}

// ExitAny_char is called when production any_char is exited.
func (s *BaseClojureListener) ExitAny_char(ctx *Any_charContext) {}

// EnterU_hex_quad is called when production u_hex_quad is entered.
func (s *BaseClojureListener) EnterU_hex_quad(ctx *U_hex_quadContext) {}

// ExitU_hex_quad is called when production u_hex_quad is exited.
func (s *BaseClojureListener) ExitU_hex_quad(ctx *U_hex_quadContext) {}

// EnterNil is called when production nil is entered.
func (s *BaseClojureListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production nil is exited.
func (s *BaseClojureListener) ExitNil(ctx *NilContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseClojureListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseClojureListener) ExitKeyword(ctx *KeywordContext) {}

// EnterSimple_keyword is called when production simple_keyword is entered.
func (s *BaseClojureListener) EnterSimple_keyword(ctx *Simple_keywordContext) {}

// ExitSimple_keyword is called when production simple_keyword is exited.
func (s *BaseClojureListener) ExitSimple_keyword(ctx *Simple_keywordContext) {}

// EnterMacro_keyword is called when production macro_keyword is entered.
func (s *BaseClojureListener) EnterMacro_keyword(ctx *Macro_keywordContext) {}

// ExitMacro_keyword is called when production macro_keyword is exited.
func (s *BaseClojureListener) ExitMacro_keyword(ctx *Macro_keywordContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseClojureListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseClojureListener) ExitSymbol(ctx *SymbolContext) {}

// EnterSimple_sym is called when production simple_sym is entered.
func (s *BaseClojureListener) EnterSimple_sym(ctx *Simple_symContext) {}

// ExitSimple_sym is called when production simple_sym is exited.
func (s *BaseClojureListener) ExitSimple_sym(ctx *Simple_symContext) {}

// EnterNs_symbol is called when production ns_symbol is entered.
func (s *BaseClojureListener) EnterNs_symbol(ctx *Ns_symbolContext) {}

// ExitNs_symbol is called when production ns_symbol is exited.
func (s *BaseClojureListener) ExitNs_symbol(ctx *Ns_symbolContext) {}

// EnterParam_name is called when production param_name is entered.
func (s *BaseClojureListener) EnterParam_name(ctx *Param_nameContext) {}

// ExitParam_name is called when production param_name is exited.
func (s *BaseClojureListener) ExitParam_name(ctx *Param_nameContext) {}
