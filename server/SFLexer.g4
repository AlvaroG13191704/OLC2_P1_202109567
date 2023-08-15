lexer grammar SFLexer;

// tokens
// -- Types
INT: 'Int';
FLOAT: 'Float';
STRING: 'String';
BOOL: 'Bool';
CHAR: 'Character';

// -- Keywords
IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
BREAK: 'break';
CONTINUE: 'continue';
RETURN: 'return';
WHILE: 'while';
FOR: 'for';
FUNC: 'func';
ARROW_FUNCTION: '->';
IN: 'in';
DOT: '.';
GUARD: 'guard';
PRINT: 'print';
TRU: 'true';
FAL: 'false';
NIL: 'nil';
DECLARATION_VAR: 'var';
DECLARATION_LET: 'let';

// RE
DIGIT_PRIMITIVE: [0-9]+ ('.'[0-9]+)?;
STRING_PRIMITIVE: '"'~["]*'"'; // : '"' (~["\r\n] | '""')* '"' ;
ID_PRIMITIVE: ([a-zA-Z_])[a-zA-Z0-9_]*;

// -- Symbols
NEGATION_OPERATOR: '!';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
COLON: ':';
COMMA: ',';
SEMICOLON: ';';
IS_:'=';
PLUS_IS: '+=';
MINUS_IS: '-=';
QUESTION_MARK: '?';

// -- Operators
PLUS : '+';
MINUS : '-';
MULTIPLY : '*';
DIVIDE : '/';
MODULO : '%';

// -- Comparators
EQUALS : '==';
NOT_EQUALS : '!=';

// -- Relational
GREATER : '>';
GREATER_EQUALS : '>=';
LESS : '<';
LESS_EQUALS : '<=';

// -- Logical
AND : '&&';
OR : '||';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
MULTI_COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment // fragment means that this rule is not a token
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;