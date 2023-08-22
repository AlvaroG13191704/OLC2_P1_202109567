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
     | forStmt
     | guardStmt
     | transferStmt
     | functionStmt
     | printstmt
     | callFunctionStmt (SEMICOLON)?
     | callBack (SEMICOLON)?
     | structStmt
     ;

transferStmt: BREAK (SEMICOLON)?          #BreakStmt
            | CONTINUE (SEMICOLON)?       #ContinueStmt
            | RETURN (expr)? (SEMICOLON)? #ReturnStmt
            ;


// struct
structStmt: STRUCT ID_PRIMITIVE LBRACE structBlock RBRACE ;

structBlock: (structStmts)* ;

structStmts: declarationStructs (SEMICOLON)?
           ;

declarationStructs :type_declaration  ID_PRIMITIVE COLON type IS_ expr #StructDeclarationWithValueAndType
           | type_declaration  ID_PRIMITIVE COLON type #StructDeclarationWithoutValue
           | type_declaration  ID_PRIMITIVE IS_ expr #StructDeclarationImplicitValue
           | type_declaration  ID_PRIMITIVE COLON LBRACKET type RBRACKET IS_ (LBRACKET exprList RBRACKET | ID_PRIMITIVE ) #StructDeclarationVector
           ;

           
structCallList: ID_PRIMITIVE COLON expr (COMMA ID_PRIMITIVE COLON expr)* ;




declaration: type_declaration  ID_PRIMITIVE COLON type IS_ expr          #TypeValueDeclaration // var value: String = "Hola"
           | type_declaration  ID_PRIMITIVE COLON type QUESTION_MARK     #TypeOptionalValueDeclaration // var value: String?
           | type_declaration  ID_PRIMITIVE IS_ expr                     #ValueDeclaration // var value = 10
           | type_declaration  ID_PRIMITIVE COLON LBRACKET type RBRACKET IS_ (LBRACKET exprList RBRACKET | ID_PRIMITIVE ) #VectorDeclaration // var value: [Int] = [1,2,3]
           ;

type_declaration: DECLARATION_VAR | DECLARATION_LET ;


assignment: ID_PRIMITIVE IS_ expr         #ValueAssignment // value = 10
          | ID_PRIMITIVE PLUS_IS expr     #PlusAssignment // var += 10
          | ID_PRIMITIVE MINUS_IS expr    #MinusAssignment // var -= 10
          | ID_PRIMITIVE LBRACKET expr RBRACKET IS_ expr #VectorAssignment // var[0] = expr
          | ID_PRIMITIVE LBRACKET expr RBRACKET MINUS_IS expr #VectorMinusAssignment // var[0] -= expr
          | ID_PRIMITIVE LBRACKET expr RBRACKET PLUS_IS expr  #VectorPlusAssignment // var[0] += expr
          ;

// if
ifstmt : IF expr LBRACE block RBRACE (ELSE LBRACE block RBRACE)?  #IfElseStmt
       | IF expr LBRACE block RBRACE ELSE ifstmt                  #ElseIfStmt
       ;


// switch
switchStmt : SWITCH expr LBRACE (caseBlock)* (defaultBlock)? RBRACE ;

caseBlock : CASE expr COLON block   ;

defaultBlock : DEFAULT COLON block  ;

// while
whileStmt : WHILE expr LBRACE block RBRACE ;


// for
forStmt :FOR ID_PRIMITIVE IN forRange LBRACE block RBRACE #ForRangeExpr
        |FOR ID_PRIMITIVE IN expr LBRACE block RBRACE     #ForExpr
        ;

forRange: left=expr DOT DOT DOT right=expr ;


// guard
guardStmt : GUARD expr ELSE LBRACE block RBRACE ;

// functions
functionStmt:MUTATING? FUNC ID_PRIMITIVE LPAREN  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE                        #FunctionWithoutParams
            |MUTATING? FUNC ID_PRIMITIVE LPAREN listFunctionParams  RPAREN (ARROW_FUNCTION type)? LBRACE block RBRACE     #FunctionWithParams
            ;

listFunctionParams: ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? type (COMMA ID_PRIMITIVE ID_PRIMITIVE COLON type)* #listFunctionParamsEI
                  | NOT_PARAM ID_PRIMITIVE COLON INOUT? type (COMMA NOT_PARAM ID_PRIMITIVE COLON type)*       #listFunctionParamsNEI
                  | ID_PRIMITIVE COLON INOUT?  type (COMMA ID_PRIMITIVE COLON INOUT? type)*                          #listFunctionParamsBEI
                  | ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA ID_PRIMITIVE ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )* #listFunctionParamsEIVector
                  | NOT_PARAM ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA NOT_PARAM ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )*       #listFunctionParamsNEIVector
                  | ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET (COMMA ID_PRIMITIVE COLON INOUT? LBRACKET type RBRACKET )*                          #listFunctionParamsBEIVector
                  ;


callFunctionStmt: ID_PRIMITIVE LPAREN  RPAREN                     #CallFunctionWithoutParams
                | ID_PRIMITIVE LPAREN listCallFunctionStmt RPAREN #CallFunctionWithParams
                ;
                
listCallFunctionStmt: REFERENCE? ID_PRIMITIVE COLON expr (COMMA REFERENCE? ID_PRIMITIVE COLON expr)*   #listCallFunctionStmtEI
                    | REFERENCE? expr (COMMA REFERENCE? expr)*                                         #listCallFunctionStmtNEI
                    ;


callBack: ID_PRIMITIVE DOT APPEND LPAREN expr LPAREN            #AppendVector // vector.append(10)
        | ID_PRIMITIVE DOT REMOVELAST LPAREN  RPAREN            #RemoveLastVector // vector.removeLast()
        | ID_PRIMITIVE DOT REMOVE LPAREN AT COLON expr RPAREN   #RemoveAtVector // vector.remove(at: 0)
        | ID_PRIMITIVE DOT ISEMPTY LPAREN  RPAREN               #IsEmptyVector // vector.isEmpty()
        | ID_PRIMITIVE DOT COUNT                                #CountVector // vector.count
        | ID_PRIMITIVE LBRACKET expr RBRACKET                   #AccessVector // let value = vector[0]
        | SELF DOT ID_PRIMITIVE (IS_ expr)?                     #SelfFunction // self.function(10) // PENDING IMPLEMENTATION
        | ID_PRIMITIVE DOT ID_PRIMITIVE (IS_ expr)?             #StructAttribute // struct.value = 10 // PENDING IMPLEMENTATION
        | ID_PRIMITIVE DOT ID_PRIMITIVE LPAREN listFunctionParams RPAREN #StructFunction // struct.function(params) // PENDING IMPLEMENTATION
        ;

// Embedded functions
embbededFunc  
            : intstmt
            | floatstmt
            | stringstmt
            ;

printstmt: PRINT LPAREN exprList RPAREN ;
exprList : expr (COMMA expr)* ;

intstmt: INT LPAREN expr RPAREN ;

floatstmt: FLOAT LPAREN expr RPAREN ;

stringstmt: STRING LPAREN expr RPAREN ;


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
    // function call
    | callFunctionStmt  (SEMICOLON)?                                    #CallFunctionExpr
    // callbacks
    | callBack (SEMICOLON)?                                             #CallBackExpr
    // emmbeded functions
    | embbededFunc                                                      #EmbeddedFunctionExpr
    // struct call
    | ID_PRIMITIVE LPAREN structCallList RPAREN                         #StructCallExpr // PENDING IMPLEMENTATION
    // Primitives
    | DIGIT_PRIMITIVE                                                   #DigitExpr
    | STRING_PRIMITIVE                                                  #StringExpr
    | ID_PRIMITIVE                                                      #IdExpr
    | NIL                                                               #NilExpr
    | (TRU|FAL)                                                         #BooleanExpr
    ;


type: (INT|FLOAT|STRING|BOOL|CHAR|ID_PRIMITIVE) ;
