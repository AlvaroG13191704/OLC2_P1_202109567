grammar SFGrammar;
import SFLexer;


start: block EOF;

block: (stmts)* // Return a slice 
    ;


stmts: declaration (SEMICOLON)?
     | assignment  (SEMICOLON)?
     | embbededFunc (SEMICOLON)?
     | ifstmt
     | switchStmt
     | whileStmt
     | BREAK (SEMICOLON)?
     ;


printstmt: PRINT LPAREN exprList RPAREN ;


declaration: type_declaration  ID_PRIMITIVE COLON type IS_ expr          #TypeValueDeclaration // var value: String = "Hola"
           | type_declaration  ID_PRIMITIVE COLON type QUESTION_MARK     #TypeOptionalValueDeclaration // var value: String?
           | type_declaration  ID_PRIMITIVE IS_ expr                     #ValueDeclaration // var value = 10
           ;

type_declaration: DECLARATION_VAR | DECLARATION_LET ;


assignment: ID_PRIMITIVE IS_ expr         #ValueAssignment // value = 10
          | ID_PRIMITIVE PLUS_IS expr     #PlusAssignment // var += 10
          | ID_PRIMITIVE MINUS_IS expr    #MinusAssignment // var -= 10
          ;


ifstmt: IF expr LBRACE block RBRACE ELSE ifstmt                      #IfElseStmt
      | IF expr LBRACE block RBRACE (ELSE LBRACE block RBRACE)?      #IfStmt
      ;

// switch
switchStmt : SWITCH expr LBRACE (caseBlock)* (defaultBlock)? RBRACE ;

caseBlock : CASE expr COLON block   ;

defaultBlock : DEFAULT COLON block  ;

// while
whileStmt : WHILE expr LBRACE block RBRACE ;

embbededFunc: printstmt  ;

exprList : expr (COMMA expr)* ;

expr: NEGATION_OPERATOR right=expr                                      #NotExpr
    | MINUS right=expr                                                  #NegExpr
    | left=expr op=(DIVIDE|MULTIPLY|MODULO) right=expr                  #ArithmeticOperationExpr
    | left=expr op=(PLUS|MINUS) right=expr                              #ArithmeticOperationExpr
    // add comparison operators
    | left=expr op=(EQUALS|NOT_EQUALS) right=expr                       #ComparationOperationExpr
    // add relational operators
    | left=expr op=(GREATER|GREATER_EQUALS|LESS|LESS_EQUALS) right=expr #RelationalOperationExpr
    // add logical operators
    | left=expr op=(AND|OR) right=expr                                  #LogicalOperationExpr
    | LPAREN expr RPAREN                                                #ParenExpr
    // Primitives
    | DIGIT_PRIMITIVE                                                   #DigitExpr
    | STRING_PRIMITIVE                                                  #StringExpr
    | ID_PRIMITIVE                                                      #IdExpr
    | NIL                                                               #NilExpr
    | (TRU|FAL)                                                         #BooleanExpr
    ;

type: (INT|FLOAT|STRING|BOOL|CHAR) ;