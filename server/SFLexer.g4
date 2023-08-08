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
PRINT: 'print';
TRU: 'true';
FAL: 'false';
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
COLON: ':';
COMMA: ',';
SEMICOLON: ';';
IS_:'=';
QUESTION_MARK: '?';

// -- Operators
PLUS : '+';
MINUS : '-';
MULTIPLY : '*';
DIVIDE : '/';
MODULO : '%';

// -- Comparators



// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
MULTI_COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment // fragment means that this rule is not a token
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;