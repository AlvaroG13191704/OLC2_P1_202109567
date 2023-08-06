grammar SFGrammar;
import SFLexer;


start: block EOF;

block: (stmts)* // Return a slice 
    ;

stmts: printstmt
    ;

printstmt: PRINT LPAREN expr RPAREN ;



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