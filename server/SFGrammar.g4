grammar SFGrammar;
import SFLexer;


start: block EOF;

block: (stmts)* // Return a slice 
    ;

stmts: printstmt
    | declaration (SEMICOLON)?
    ;

printstmt: PRINT LPAREN expr RPAREN ;

// Examples:
// var value: String?
// var value = 10
// var valor: Int = 10
// let constante: String = "Hola"
declaration: type_declaration  ID_PRIMITIVE COLON type IS_ expr    #TypeValueDeclaration
           | type_declaration  ID_PRIMITIVE COLON type QUESTION_MARK    #TypeOptionalValueDeclaration
           | type_declaration  ID_PRIMITIVE IS_ expr               #ValueDeclaration
           ;

type_declaration: DECLARATION_VAR | DECLARATION_LET ;

expr: NEGATION_OPERATOR right=expr                        #NotExpr
    | MINUS right=expr                                    #NegExpr
    | left=expr op=(DIVIDE|MULTIPLY|MODULO) right=expr    #OperationExpr
    | left=expr op=(PLUS|MINUS) right=expr                #OperationExpr
    // add comparison operators
    // add relational operators
    // add logical operators
    | LPAREN expr RPAREN                                   #ParenExpr
    // Primitives
    | DIGIT_PRIMITIVE                                      #DigitExpr
    | STRING_PRIMITIVE                                     #StringExpr
    | ID_PRIMITIVE                                         #IdExpr
    | (TRU|FAL)                                            #BooleanExpr
    ;

type: (INT|FLOAT|STRING|BOOL|CHAR) ;