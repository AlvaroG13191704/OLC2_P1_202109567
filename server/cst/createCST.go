package cst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// make a request to the server to create a new CST

// url -> http://lab.antlr.org/parse/ - POST

// grammar:"parser grammar ExprParser;\noptions { tokenVocab=ExprLexer; }\n\nprogram: block EOF;\n\nblock: (stmts)* // Return a slice \n    ;\n\n\nstmts: declaration (SEMICOLON)?\n     | assignment  (SEMICOLON)?\n     | embbededFunc (SEMICOLON)?\n     | ifstmt\n     | switchStmt\n     | whileStmt\n     | forStmt\n     | guardStmt\n     | transferStmt\n     | functionStmt\n     | printstmt\n     | callFunctionStmt (SEMICOLON)?\n     | callBack (SEMICOLON)?\n     | structStmt\n     ;\n\ntransferStmt: BREAK (SEMICOLON)?          #BreakStmt\n            | CONTINUE (SEMICOLON)?       #ContinueStmt\n            | RETURN (expr)? (SEMICOLON)? #ReturnStmt\n            ;\n\n\n// struct\nstructStmt: STRUCT ID_PRIMITIVE LBRACE structBlock RBRACE ;\n\nstructBlock: (structStmts)* ;\n\nstructStmts: declarationStructs (SEMICOLON)?\n           | functionStructs \n           ;\n\ndeclarationStructs :type_declaration  ID_PRIMITIVE COLON type IS_ expr #StructDeclarationWithValueAndType\n           | type_declaration  ID_PRIMITIVE COLON type #StructDeclarationWithoutValue\n           | type_declaration  ID_PRIMITIVE IS_ expr #StructDeclarationImplicitValue\n           | type_declaration  ID_PRIMITIVE COLON LBRACKET type RBRACKET IS_ (LBRACKET exprList RBRACKET | ID_PRIMITIVE ) #StructDeclarationVector\n           ;\n\nfunctionStructs:MUTATING? FUNC ID_PRIMITIVE LPAREN  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE                  #StructFunctionWithoutParams\n            |MUTATING? FUNC ID_PRIMITIVE LPAREN listFunctionParams  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE  #StructFunctionWithParams\n            ;\n           \nstructCallList: ID_PRIMITIVE COLON expr (COMMA ID_PRIMITIVE COLON expr)* ;\n\n\n\n\ndeclaration  // var fruta = Fruta(nombre: \"pera\", precio: 10) -> Struct creation\n           : type_declaration  ID_PRIMITIVE IS_ ID_PRIMITIVE LPAREN structCallList RPAREN #StructCreation\n           | type_declaration  ID_PRIMITIVE COLON type IS_ expr          #TypeValueDeclaration // var value: String = \"Hola\"\n           | type_declaration  ID_PRIMITIVE COLON type QUESTION_MARK     #TypeOptionalValueDeclaration // var value: String?\n           | type_declaration  ID_PRIMITIVE IS_ expr                     #ValueDeclaration // var value = 10\n           | type_declaration  ID_PRIMITIVE COLON LBRACKET type RBRACKET IS_ (LBRACKET exprList RBRACKET | ID_PRIMITIVE ) #VectorDeclaration // var value: [Int] = [1,2,3]\n           ;\n\ntype_declaration: DECLARATION_VAR | DECLARATION_LET ;\n\n\nassignment: ID_PRIMITIVE IS_ expr         #ValueAssignment // value = 10\n          | ID_PRIMITIVE PLUS_IS expr     #PlusAssignment // var += 10\n          | ID_PRIMITIVE MINUS_IS expr    #MinusAssignment // var -= 10\n          | ID_PRIMITIVE LBRACKET expr RBRACKET IS_ expr #VectorAssignment // var[0] = expr\n          | ID_PRIMITIVE LBRACKET expr RBRACKET MINUS_IS expr #VectorMinusAssignment // var[0] -= expr\n          | ID_PRIMITIVE LBRACKET expr RBRACKET PLUS_IS expr  #VectorPlusAssignment // var[0] += expr\n          ;\n\n// if\nifstmt : IF expr LBRACE block RBRACE (ELSE LBRACE block RBRACE)?  #IfElseStmt\n       | IF expr LBRACE block RBRACE ELSE ifstmt                  #ElseIfStmt\n       ;\n\n\n// switch\nswitchStmt : SWITCH expr LBRACE (caseBlock)* (defaultBlock)? RBRACE ;\n\ncaseBlock : CASE expr COLON block   ;\n\ndefaultBlock : DEFAULT COLON block  ;\n\n// while\nwhileStmt : WHILE expr LBRACE block RBRACE ;\n\n\n// for\nforStmt :FOR ID_PRIMITIVE IN forRange LBRACE block RBRACE #ForRangeExpr\n        |FOR ID_PRIMITIVE IN expr LBRACE block RBRACE     #ForExpr\n        ;\n\nforRange: left=expr DOT DOT DOT right=expr ;\n\n\n// guard\nguardStmt : GUARD expr ELSE LBRACE block RBRACE ;\n\n// functions\nfunctionStmt:FUNC ID_PRIMITIVE LPAREN  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE                        #FunctionWithoutParams\n            |FUNC ID_PRIMITIVE LPAREN listFunctionParams  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE     #FunctionWithParams\n            ;\n\nlistFunctionParams: ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? type (COMMA ID_PRIMITIVE ID_PRIMITIVE COLON type)* #listFunctionParamsEI\n                  | NOT_PARAM ID_PRIMITIVE COLON INOUT? type (COMMA NOT_PARAM ID_PRIMITIVE COLON type)*       #listFunctionParamsNEI\n                  | ID_PRIMITIVE COLON INOUT?  type (COMMA ID_PRIMITIVE COLON INOUT? type)*                          #listFunctionParamsBEI\n                  | ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )* #listFunctionParamsEIVector\n                  | NOT_PARAM ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA NOT_PARAM ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )*       #listFunctionParamsNEIVector\n                  | ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )*                          #listFunctionParamsBEIVector\n                  ;\n\n\ncallFunctionStmt: ID_PRIMITIVE LPAREN  RPAREN                     #CallFunctionWithoutParams\n                | ID_PRIMITIVE LPAREN listCallFunctionStmt RPAREN #CallFunctionWithParams\n                ;\n                \nlistCallFunctionStmt: REFERENCE? ID_PRIMITIVE COLON expr (COMMA REFERENCE? ID_PRIMITIVE COLON expr)*   #listCallFunctionStmtEI\n                    | REFERENCE? expr (COMMA REFERENCE? expr)*                                         #listCallFunctionStmtNEI\n                    ;\n\n\ncallBack: ID_PRIMITIVE DOT APPEND LPAREN expr LPAREN            #AppendVector // vector.append(10)\n        | ID_PRIMITIVE DOT REMOVELAST LPAREN  RPAREN            #RemoveLastVector // vector.removeLast()\n        | ID_PRIMITIVE DOT REMOVE LPAREN AT COLON expr RPAREN   #RemoveAtVector // vector.remove(at: 0)\n        | ID_PRIMITIVE DOT ISEMPTY LPAREN  RPAREN               #IsEmptyVector // vector.isEmpty()\n        | ID_PRIMITIVE DOT COUNT                                #CountVector // vector.count\n        | ID_PRIMITIVE LBRACKET expr RBRACKET                   #AccessVector // let value = vector[0]\n        | SELF DOT ID_PRIMITIVE (IS_ expr)?                     #SelfFunction // self.function(10) // TODO: PENDING IMPLEMENTATION\n        | ID_PRIMITIVE DOT ID_PRIMITIVE (IS_ expr)?             #StructAttribute // struct.value = 10 //TODO: PENDING IMPLEMENTATION\n        | ID_PRIMITIVE DOT ID_PRIMITIVE LPAREN listFunctionParams RPAREN #StructFunction // struct.function(params) // TODO: PENDING IMPLEMENTATION\n        ;\n\n// Embedded functions\nembbededFunc  \n            : intstmt\n            | floatstmt\n            | stringstmt\n            ;\n\nprintstmt: PRINT LPAREN exprList RPAREN ;\n\nexprList : expr (COMMA expr)* ;\n\nintstmt: INT LPAREN expr RPAREN ;\n\nfloatstmt: FLOAT LPAREN expr RPAREN ;\n\nstringstmt: STRING LPAREN expr RPAREN ;\n\n\nexpr: NEGATION_OPERATOR right=expr                                      #NotExpr\n    | MINUS right=expr                                                  #NegExpr\n    | left=expr op=(DIVIDE|MULTIPLY|MODULO) right=expr                  #ArithmeticOperationExpr\n    | left=expr op=(PLUS|MINUS) right=expr                              #ArithmeticOperationExpr\n    // add comparison operators\n    | left=expr op=(EQUALS|NOT_EQUALS) right=expr                       #ComparationOperationExpr\n    // add relational operators\n    | left=expr op=(GREATER|GREATER_EQUALS|LESS|LESS_EQUALS) right=expr #RelationalOperationExpr\n    // add logical operators\n    | left=expr op=(AND|OR) right=expr                                  #LogicalOperationExpr\n    | LPAREN expr RPAREN                                                #ParenExpr\n    // struct call\n    | ID_PRIMITIVE LPAREN structCallList RPAREN                         #StructCallExpr //TODO: PENDING IMPLEMENTATION \n    // function call\n    | callFunctionStmt  (SEMICOLON)?                                    #CallFunctionExpr\n    // callbacks\n    | callBack (SEMICOLON)?                                             #CallBackExpr\n    // emmbeded functions\n    | embbededFunc                                                      #EmbeddedFunctionExpr\n    // Primitives\n    | DIGIT_PRIMITIVE                                                   #DigitExpr\n    | STRING_PRIMITIVE                                                  #StringExpr\n    | ID_PRIMITIVE                                                      #IdExpr\n    | NIL                                                               #NilExpr\n    | (TRU|FAL)                                                         #BooleanExpr\n    ;\n\n\ntype: (INT|FLOAT|STRING|BOOL|CHAR|ID_PRIMITIVE) ;\n"
// input:"var num3: Int = 10.1\nvar num2: Int = false\nvar num2: Int = false"
// lexgrammar:"// DELETE THIS CONTENT IF YOU PUT COMBINED GRAMMAR IN Parser TAB\nlexer grammar ExprLexer;\n\n// tokens\n// -- Types\nINT: 'Int';\nFLOAT: 'Float';\nSTRING: 'String';\nBOOL: 'Bool';\nCHAR: 'Character';\n\n// -- Keywords\nIF: 'if';\nSTRUCT: 'struct';\nMUTATING: 'mutating';\nAPPEND: 'append';\nCOUNT: 'count';\nREMOVELAST: 'removeLast';\nREMOVE: 'remove';\nISEMPTY: 'isEmpty';\nSELF: 'self';\nAT: 'at';\nELSE: 'else';\nSWITCH: 'switch';\nCASE: 'case';\nDEFAULT: 'default';\nBREAK: 'break';\nCONTINUE: 'continue';\nRETURN: 'return';\nWHILE: 'while';\nFOR: 'for';\nFUNC: 'func';\nARROW_FUNCTION: '->';\nIN: 'in';\nDOT: '.';\nGUARD: 'guard';\nPRINT: 'print';\nTRU: 'true';\nFAL: 'false';\nNIL: 'nil';\nDECLARATION_VAR: 'var';\nDECLARATION_LET: 'let';\nREFERENCE: '&';\nINOUT: 'inout';\nNOT_PARAM: '_';\n\n\n// RE\nDIGIT_PRIMITIVE: [0-9]+ ('.'[0-9]+)?;\nSTRING_PRIMITIVE: '\"' ( '\\\\' [nrt\"\\\\] | ~[\\n\\r\"])* '\"'; // : '\"' (~[\"\\r\\n] | '\"\"')* '\"' ;   '\"'~[\"]*'\"'\nID_PRIMITIVE: ([a-zA-Z_])[a-zA-Z0-9_]*;\n\n// -- Symbols\nNEGATION_OPERATOR: '!';\nLPAREN: '(';\nRPAREN: ')';\nLBRACE: '{';\nRBRACE: '}';\nLBRACKET: '[';\nRBRACKET: ']';\nCOLON: ':';\nCOMMA: ',';\nSEMICOLON: ';';\nIS_:'=';\nPLUS_IS: '+=';\nMINUS_IS: '-=';\nQUESTION_MARK: '?';\n\n// -- Operators\nPLUS : '+';\nMINUS : '-';\nMULTIPLY : '*';\nDIVIDE : '/';\nMODULO : '%';\n\n// -- Comparators\nEQUALS : '==';\nNOT_EQUALS : '!=';\n\n// -- Relational\nGREATER : '>';\nGREATER_EQUALS : '>=';\nLESS : '<';\nLESS_EQUALS : '<=';\n\n// -- Logical\nAND : '&&';\nOR : '||';\n\n// skip\nWHITESPACE: [ \\\\\\r\\n\\t]+ -> skip;\nMULTI_COMMENT : '/*' .*? '*/' -> skip;\nLINE_COMMENT : '//' ~[\\r\\n]* -> skip;\n\nfragment // fragment means that this rule is not a token\nESC_SEQ\n    :   '\\\\' ('\\\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')\n    ;"
// start:"program"

func CreateCST(input string) string {

	// replace the " with \"
	input = strings.ReplaceAll(input, "\"", "\\\"")

	url := "http://lab.antlr.org/parse/"

	// payload
	payload := []byte(`{"grammar": "parser grammar ExprParser;\noptions { tokenVocab=ExprLexer; }\n\nprogram: block EOF;\n\nblock: (stmts)* // Return a slice \n    ;\n\n\nstmts: declaration (SEMICOLON)?\n     | assignment  (SEMICOLON)?\n     | embbededFunc (SEMICOLON)?\n     | ifstmt\n     | switchStmt\n     | whileStmt\n     | forStmt\n     | guardStmt\n     | transferStmt\n     | functionStmt\n     | printstmt\n     | callFunctionStmt (SEMICOLON)?\n     | callBack (SEMICOLON)?\n     | structStmt\n     ;\n\ntransferStmt: BREAK (SEMICOLON)?          #BreakStmt\n            | CONTINUE (SEMICOLON)?       #ContinueStmt\n            | RETURN (expr)? (SEMICOLON)? #ReturnStmt\n            ;\n\n\n// struct\nstructStmt: STRUCT ID_PRIMITIVE LBRACE structBlock RBRACE ;\n\nstructBlock: (structStmts)* ;\n\nstructStmts: declarationStructs (SEMICOLON)?\n           | functionStructs \n           ;\n\ndeclarationStructs :type_declaration  ID_PRIMITIVE COLON type IS_ expr #StructDeclarationWithValueAndType\n           | type_declaration  ID_PRIMITIVE COLON type #StructDeclarationWithoutValue\n           | type_declaration  ID_PRIMITIVE IS_ expr #StructDeclarationImplicitValue\n           | type_declaration  ID_PRIMITIVE COLON LBRACKET type RBRACKET IS_ (LBRACKET exprList RBRACKET | ID_PRIMITIVE ) #StructDeclarationVector\n           ;\n\nfunctionStructs:MUTATING? FUNC ID_PRIMITIVE LPAREN  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE                  #StructFunctionWithoutParams\n            |MUTATING? FUNC ID_PRIMITIVE LPAREN listFunctionParams  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE  #StructFunctionWithParams\n            ;\n           \nstructCallList: ID_PRIMITIVE COLON expr (COMMA ID_PRIMITIVE COLON expr)* ;\n\n\n\n\ndeclaration  // var fruta = Fruta(nombre: \"pera\", precio: 10) -> Struct creation\n           : type_declaration  ID_PRIMITIVE IS_ ID_PRIMITIVE LPAREN structCallList RPAREN #StructCreation\n           | type_declaration  ID_PRIMITIVE COLON type IS_ expr          #TypeValueDeclaration // var value: String = \"Hola\"\n           | type_declaration  ID_PRIMITIVE COLON type QUESTION_MARK     #TypeOptionalValueDeclaration // var value: String?\n           | type_declaration  ID_PRIMITIVE IS_ expr                     #ValueDeclaration // var value = 10\n           | type_declaration  ID_PRIMITIVE COLON LBRACKET type RBRACKET IS_ (LBRACKET exprList RBRACKET | ID_PRIMITIVE ) #VectorDeclaration // var value: [Int] = [1,2,3]\n           ;\n\ntype_declaration: DECLARATION_VAR | DECLARATION_LET ;\n\n\nassignment: ID_PRIMITIVE IS_ expr         #ValueAssignment // value = 10\n          | ID_PRIMITIVE PLUS_IS expr     #PlusAssignment // var += 10\n          | ID_PRIMITIVE MINUS_IS expr    #MinusAssignment // var -= 10\n          | ID_PRIMITIVE LBRACKET expr RBRACKET IS_ expr #VectorAssignment // var[0] = expr\n          | ID_PRIMITIVE LBRACKET expr RBRACKET MINUS_IS expr #VectorMinusAssignment // var[0] -= expr\n          | ID_PRIMITIVE LBRACKET expr RBRACKET PLUS_IS expr  #VectorPlusAssignment // var[0] += expr\n          ;\n\n// if\nifstmt : IF expr LBRACE block RBRACE (ELSE LBRACE block RBRACE)?  #IfElseStmt\n       | IF expr LBRACE block RBRACE ELSE ifstmt                  #ElseIfStmt\n       ;\n\n\n// switch\nswitchStmt : SWITCH expr LBRACE (caseBlock)* (defaultBlock)? RBRACE ;\n\ncaseBlock : CASE expr COLON block   ;\n\ndefaultBlock : DEFAULT COLON block  ;\n\n// while\nwhileStmt : WHILE expr LBRACE block RBRACE ;\n\n\n// for\nforStmt :FOR ID_PRIMITIVE IN forRange LBRACE block RBRACE #ForRangeExpr\n        |FOR ID_PRIMITIVE IN expr LBRACE block RBRACE     #ForExpr\n        ;\n\nforRange: left=expr DOT DOT DOT right=expr ;\n\n\n// guard\nguardStmt : GUARD expr ELSE LBRACE block RBRACE ;\n\n// functions\nfunctionStmt:FUNC ID_PRIMITIVE LPAREN  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE                        #FunctionWithoutParams\n            |FUNC ID_PRIMITIVE LPAREN listFunctionParams  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE     #FunctionWithParams\n            ;\n\nlistFunctionParams: ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? type (COMMA ID_PRIMITIVE ID_PRIMITIVE COLON type)* #listFunctionParamsEI\n                  | NOT_PARAM ID_PRIMITIVE COLON INOUT? type (COMMA NOT_PARAM ID_PRIMITIVE COLON type)*       #listFunctionParamsNEI\n                  | ID_PRIMITIVE COLON INOUT?  type (COMMA ID_PRIMITIVE COLON INOUT? type)*                          #listFunctionParamsBEI\n                  | ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )* #listFunctionParamsEIVector\n                  | NOT_PARAM ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA NOT_PARAM ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )*       #listFunctionParamsNEIVector\n                  | ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )*                          #listFunctionParamsBEIVector\n                  ;\n\n\ncallFunctionStmt: ID_PRIMITIVE LPAREN  RPAREN                     #CallFunctionWithoutParams\n                | ID_PRIMITIVE LPAREN listCallFunctionStmt RPAREN #CallFunctionWithParams\n                ;\n                \nlistCallFunctionStmt: REFERENCE? ID_PRIMITIVE COLON expr (COMMA REFERENCE? ID_PRIMITIVE COLON expr)*   #listCallFunctionStmtEI\n                    | REFERENCE? expr (COMMA REFERENCE? expr)*                                         #listCallFunctionStmtNEI\n                    ;\n\n\ncallBack: ID_PRIMITIVE DOT APPEND LPAREN expr LPAREN            #AppendVector // vector.append(10)\n        | ID_PRIMITIVE DOT REMOVELAST LPAREN  RPAREN            #RemoveLastVector // vector.removeLast()\n        | ID_PRIMITIVE DOT REMOVE LPAREN AT COLON expr RPAREN   #RemoveAtVector // vector.remove(at: 0)\n        | ID_PRIMITIVE DOT ISEMPTY LPAREN  RPAREN               #IsEmptyVector // vector.isEmpty()\n        | ID_PRIMITIVE DOT COUNT                                #CountVector // vector.count\n        | ID_PRIMITIVE LBRACKET expr RBRACKET                   #AccessVector // let value = vector[0]\n        | SELF DOT ID_PRIMITIVE (IS_ expr)?                     #SelfFunction // self.function(10) // TODO: PENDING IMPLEMENTATION\n        | ID_PRIMITIVE DOT ID_PRIMITIVE (IS_ expr)?             #StructAttribute // struct.value = 10 //TODO: PENDING IMPLEMENTATION\n        | ID_PRIMITIVE DOT ID_PRIMITIVE LPAREN listFunctionParams RPAREN #StructFunction // struct.function(params) // TODO: PENDING IMPLEMENTATION\n        ;\n\n// Embedded functions\nembbededFunc  \n            : intstmt\n            | floatstmt\n            | stringstmt\n            ;\n\nprintstmt: PRINT LPAREN exprList RPAREN ;\n\nexprList : expr (COMMA expr)* ;\n\nintstmt: INT LPAREN expr RPAREN ;\n\nfloatstmt: FLOAT LPAREN expr RPAREN ;\n\nstringstmt: STRING LPAREN expr RPAREN ;\n\n\nexpr: NEGATION_OPERATOR right=expr                                      #NotExpr\n    | MINUS right=expr                                                  #NegExpr\n    | left=expr op=(DIVIDE|MULTIPLY|MODULO) right=expr                  #ArithmeticOperationExpr\n    | left=expr op=(PLUS|MINUS) right=expr                              #ArithmeticOperationExpr\n    // add comparison operators\n    | left=expr op=(EQUALS|NOT_EQUALS) right=expr                       #ComparationOperationExpr\n    // add relational operators\n    | left=expr op=(GREATER|GREATER_EQUALS|LESS|LESS_EQUALS) right=expr #RelationalOperationExpr\n    // add logical operators\n    | left=expr op=(AND|OR) right=expr                                  #LogicalOperationExpr\n    | LPAREN expr RPAREN                                                #ParenExpr\n    // struct call\n    | ID_PRIMITIVE LPAREN structCallList RPAREN                         #StructCallExpr //TODO: PENDING IMPLEMENTATION \n    // function call\n    | callFunctionStmt  (SEMICOLON)?                                    #CallFunctionExpr\n    // callbacks\n    | callBack (SEMICOLON)?                                             #CallBackExpr\n    // emmbeded functions\n    | embbededFunc                                                      #EmbeddedFunctionExpr\n    // Primitives\n    | DIGIT_PRIMITIVE                                                   #DigitExpr\n    | STRING_PRIMITIVE                                                  #StringExpr\n    | ID_PRIMITIVE                                                      #IdExpr\n    | NIL                                                               #NilExpr\n    | (TRU|FAL)                                                         #BooleanExpr\n    ;\n\n\ntype: (INT|FLOAT|STRING|BOOL|CHAR|ID_PRIMITIVE) ;\n",
	"input": "` + input + `", 
	"lexgrammar": "// DELETE THIS CONTENT IF YOU PUT COMBINED GRAMMAR IN Parser TAB\nlexer grammar ExprLexer;\n\n// tokens\n// -- Types\nINT: 'Int';\nFLOAT: 'Float';\nSTRING: 'String';\nBOOL: 'Bool';\nCHAR: 'Character';\n\n// -- Keywords\nIF: 'if';\nSTRUCT: 'struct';\nMUTATING: 'mutating';\nAPPEND: 'append';\nCOUNT: 'count';\nREMOVELAST: 'removeLast';\nREMOVE: 'remove';\nISEMPTY: 'isEmpty';\nSELF: 'self';\nAT: 'at';\nELSE: 'else';\nSWITCH: 'switch';\nCASE: 'case';\nDEFAULT: 'default';\nBREAK: 'break';\nCONTINUE: 'continue';\nRETURN: 'return';\nWHILE: 'while';\nFOR: 'for';\nFUNC: 'func';\nARROW_FUNCTION: '->';\nIN: 'in';\nDOT: '.';\nGUARD: 'guard';\nPRINT: 'print';\nTRU: 'true';\nFAL: 'false';\nNIL: 'nil';\nDECLARATION_VAR: 'var';\nDECLARATION_LET: 'let';\nREFERENCE: '&';\nINOUT: 'inout';\nNOT_PARAM: '_';\n\n\n// RE\nDIGIT_PRIMITIVE: [0-9]+ ('.'[0-9]+)?;\nSTRING_PRIMITIVE: '\"' ( '\\\\' [nrt\"\\\\] | ~[\\n\\r\"])* '\"'; // : '\"' (~[\"\\r\\n] | '\"\"')* '\"' ;   '\"'~[\"]*'\"'\nID_PRIMITIVE: ([a-zA-Z_])[a-zA-Z0-9_]*;\n\n// -- Symbols\nNEGATION_OPERATOR: '!';\nLPAREN: '(';\nRPAREN: ')';\nLBRACE: '{';\nRBRACE: '}';\nLBRACKET: '[';\nRBRACKET: ']';\nCOLON: ':';\nCOMMA: ',';\nSEMICOLON: ';';\nIS_:'=';\nPLUS_IS: '+=';\nMINUS_IS: '-=';\nQUESTION_MARK: '?';\n\n// -- Operators\nPLUS : '+';\nMINUS : '-';\nMULTIPLY : '*';\nDIVIDE : '/';\nMODULO : '%';\n\n// -- Comparators\nEQUALS : '==';\nNOT_EQUALS : '!=';\n\n// -- Relational\nGREATER : '>';\nGREATER_EQUALS : '>=';\nLESS : '<';\nLESS_EQUALS : '<=';\n\n// -- Logical\nAND : '&&';\nOR : '||';\n\n// skip\nWHITESPACE: [ \\\\\\r\\n\\t]+ -> skip;\nMULTI_COMMENT : '/*' .*? '*/' -> skip;\nLINE_COMMENT : '//' ~[\\r\\n]* -> skip;\n\nfragment // fragment means that this rule is not a token\nESC_SEQ\n    :   '\\\\' ('\\\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')\n    ;",
	"start": "program"}`)

	// request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json") // Set the appropriate content type

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	// parse the response body to json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}

	// create a map to store the json
	var data map[string]interface{}

	// unmarshal the json
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}

	// fmt.Println("Response Body:", data)

	result := data["result"].(map[string]interface{})

	// fmt.Println("Response Body:", result["svgtree"])

	// create the file
	err = os.WriteFile("cst.svg", []byte(result["svgtree"].(string)), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}

	fmt.Println("File created successfully")

	return result["svgtree"].(string)

}
