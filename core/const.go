package core

type TokenType uint64
type State uint64
type AstType uint64

const (
	STATE_DOUBLE    State = 0xa001 //
	STATE_HEX       State = 0xa002
	STATE_INT       State = 0xa003
	STATE_INT_0     State = 0xa004
	STATE_INT_OCTAL State = 0xa005
	STATE_INIT      State = 0xaffe // 状态机初始状态
	STATE_END       State = 0xafff // 状态机结束状态
)

const (
	AST_INVALID   AstType = 0x00
	AST_INT       AstType = 0x010
	AST_STRING    AstType = 0x020
	AST_DOUBLE    AstType = 0x030
	AST_BOOL      AstType = 0x040
	AST_FUNC      AstType = 0x050
	AST_CLASS     AstType = 0x060
	AST_BREAK     AstType = 0x070
	AST_RETURN    AstType = 0x080
	AST_CONTINUE  AstType = 0x090
	AST_IF        AstType = 0x0a0
	AST_FOR       AstType = 0x0b0
	AST_FOREACH   AstType = 0x0c0
	AST_NIL       AstType = 0x0c0
	AST_RESULT    AstType = 0x0d0
	AST_ASSIGN    AstType = 0x0e0
	AST_STATEMENT AstType = 0x0f0
	AST_EXPR      AstType = 0x100
	AST_LIST      AstType = 0x110
	AST_TUPLE     AstType = 0x120
	AST_DICT      AstType = 0x130
	AST_BIN_OP    AstType = 0x140
	AST_VAR       AstType = 0x150
	AST_FUNC_CALL AstType = 0x160
)

const (
	INVALID       TokenType = 0xf000
	INT           TokenType = 0xf001 // int
	STRING        TokenType = 0xf002 // string
	PLUS          TokenType = 0xf003 // +
	MINUS         TokenType = 0xf004 // -
	MULTI         TokenType = 0xf005 // *
	DIV           TokenType = 0xf006 // /
	DOUBLE        TokenType = 0xf007 // double
	FUNC          TokenType = 0xf008 // fn
	LINE_COMMENT  TokenType = 0xf009 // #
	KEY           TokenType = 0xf00a // .... 标识符
	MOD           TokenType = 0xf00b // %
	KEY_FUNC      TokenType = 0xf00c // keyword func
	KEY_STRING    TokenType = 0xf00d // keyword string（不用）
	KEY_DOUBLE    TokenType = 0xf00e // keyword double（不用）
	KEY_INT       TokenType = 0xf00f // keyword int （不用）
	LESS          TokenType = 0xf010 // <
	GREAT         TokenType = 0xf011 // >
	LEQ           TokenType = 0xf012 // <=
	GEQ           TokenType = 0xf013 // >=
	EQUAL         TokenType = 0xf014 // ==
	ASSIGN        TokenType = 0xf015 // =
	BLOCK_COMMENT TokenType = 0xf016 // <% %>
	LBRK          TokenType = 0xf017 // [
	RBRK          TokenType = 0xf018 // ]
	CHAR          TokenType = 0xf019 // '?'
	KEY_CHAR      TokenType = 0xf01a // '
	PLUS_PLUS     TokenType = 0xf01b // ++
	MINUS_MINUS   TokenType = 0xf01c // --
	KEY_ELIF      TokenType = 0xf01d // keyword elif
	KEY_ELSE      TokenType = 0xf01e // keyword else
	KEY_IF        TokenType = 0xf01f // keyword if
	LPRNTH        TokenType = 0xf020 // (
	RPRNTH        TokenType = 0xf021 // )
	LBRCS         TokenType = 0xf022 // {
	RBRCS         TokenType = 0xf023 // }
	KEY_FOR       TokenType = 0xf024 // keyword for
	KEY_RETURN    TokenType = 0xf025 // keyword return
	HEX_INT       TokenType = 0xf026 // 0x??
	REF           TokenType = 0xf027 // &
	AND           TokenType = 0xf028 // &&
	BITOR         TokenType = 0xf029 // |
	OR            TokenType = 0xf02a // ||
	NOT           TokenType = 0xf02b // !
	COMMA         TokenType = 0xf02c // ,
	KEY_CLASS     TokenType = 0xf02d // keyword class
	KEY_START     TokenType = 0xf02e // keyword start of program
	KEY_GLOBAL    TokenType = 0xf02f // keyword global data
	KEY_IMPORT    TokenType = 0xf030 // keyword import
	KEY_TRUE      TokenType = 0xf031 // keyword true bool
	KEY_FALSE     TokenType = 0xf032 // keyword false bool
	EOF           TokenType = 0xf033 // end
	QUOTE         TokenType = 0xf034 // . 对象引用
	REVERSE       TokenType = 0xf035 // ~
	PLUS_EQ       TokenType = 0xf036 // +=
	MINUS_EQ      TokenType = 0xf037 // -=
	MULTI_EQ      TokenType = 0xf038 // *=
	DIV_EQ        TokenType = 0xf039 // /=
	MOD_EQ        TokenType = 0xf03a // %=
	NULL          TokenType = 0xf03b // nil 空值
	THIS          TokenType = 0xf03c // this
	OCT_INT       TokenType = 0xf03d // 八进制
	NOT_EQ        TokenType = 0xf03e // !=
	BOOLEAN       TokenType = 0xf03f // boolean
	SEMICOLON     TokenType = 0xf040 // ;
	XOR           TokenType = 0xf041 // ^
	LSHIFT        TokenType = 0xf042 // <<
	RSHIFT        TokenType = 0xf043 // >>
	COLON         TokenType = 0xf044 // :
	KEY_FOREACH   TokenType = 0xf045 // foreach
	KEY_BREAK     TokenType = 0xf046 // break
	KEY_CONTINUE  TokenType = 0xf047 // continue
	INHERIT       TokenType = 0xf048 // 类继承标识
)

var typ = [...]string{
	INVALID & 0x0fff:       "INVALID",
	INT & 0x0fff:           "INT",
	STRING & 0x0fff:        "STRING",
	PLUS & 0x0fff:          "+",
	MINUS & 0x0fff:         "-",
	MULTI & 0x0fff:         "*",
	DIV & 0x0fff:           "/",
	DOUBLE & 0x0fff:        "DOUBLE",
	FUNC & 0x0fff:          "FUNC",
	LINE_COMMENT & 0x0fff:  "#",
	KEY & 0x0fff:           "KEY",
	MOD & 0x0fff:           "%",
	KEY_FUNC & 0x0fff:      "func",
	KEY_STRING & 0x0fff:    "KEY STRING",
	KEY_DOUBLE & 0x0fff:    "KEY DOUBLE",
	KEY_INT & 0x0fff:       "KEY INT",
	LESS & 0x0fff:          "<",
	GREAT & 0x0fff:         ">",
	LEQ & 0x0fff:           "<=",
	GEQ & 0x0fff:           ">=",
	EQUAL & 0x0fff:         "==",
	ASSIGN & 0x0fff:        "=",
	BLOCK_COMMENT & 0x0fff: "<% %>",
	LBRK & 0x0fff:          "[",
	RBRK & 0x0fff:          "]",
	CHAR & 0x0fff:          "'",
	KEY_CHAR & 0x0fff:      "KEY CHAR",
	PLUS_PLUS & 0x0fff:     "++",
	MINUS_MINUS & 0x0fff:   "--",
	KEY_ELIF & 0x0fff:      "elif",
	KEY_ELSE & 0x0fff:      "else",
	KEY_IF & 0x0fff:        "if",
	LPRNTH & 0x0fff:        "(",
	RPRNTH & 0x0fff:        ")",
	LBRCS & 0x0fff:         "{",
	RBRCS & 0x0fff:         "}",
	KEY_FOR & 0x0fff:       "for",
	KEY_RETURN & 0x0fff:    "return",
	HEX_INT & 0x0fff:       "HEX INT",
	REF & 0x0fff:           "&",
	AND & 0x0fff:           "&&",
	BITOR & 0x0fff:         "|",
	OR & 0x0fff:            "||",
	NOT & 0x0fff:           "!",
	COMMA & 0x0fff:         ",",
	KEY_CLASS & 0x0fff:     "class",
	KEY_START & 0x0fff:     "start",
	KEY_GLOBAL & 0x0fff:    "global",
	KEY_IMPORT & 0x0fff:    "import",
	KEY_TRUE & 0x0fff:      "true",
	KEY_FALSE & 0x0fff:     "false",
	EOF & 0x0fff:           "EOF",
	QUOTE & 0x0fff:         ".",
	REVERSE & 0x0fff:       "~",
	PLUS_EQ & 0x0fff:       "+=",
	MINUS_EQ & 0x0fff:      "-=",
	MULTI_EQ & 0x0fff:      "*=",
	DIV_EQ & 0x0fff:        "/=",
	MOD_EQ & 0x0fff:        "%=",
	NULL & 0x0fff:          "nil",
	THIS & 0x0fff:          "this",
	OCT_INT & 0x0fff:       "OCTAL NUMBER",
	NOT_EQ & 0x0fff:        "!=",
	BOOLEAN & 0x0fff:       "Boolean",
	SEMICOLON & 0x0fff:     ";",
	XOR & 0x0fff:           "^",
	LSHIFT & 0x0fff:        "<<",
	RSHIFT & 0x0fff:        ">>",
	COLON & 0x0fff:         ":",
	KEY_FOREACH & 0x0fff:   "foreach",
	KEY_BREAK & 0x0fff:     "break",
	KEY_CONTINUE & 0x0fff:  "continue",
	INHERIT & 0x0fff:       "@",
}

func (t TokenType) String() string {
	return typ[t&0x0fff]
}
