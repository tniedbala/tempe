lexer grammar TemplateLexer;

OPEN_EXPR
  : '{{' -> pushMode(Expression)
  ;

OPEN_STMT
  : '{%' WS* -> pushMode(Statement)
  ;

TEXT 
  : ( ~'{' | '{' ~('{'|'%') )+
  ;

// ----------------------------------------------------
// expression: code to be evaluated within double curly braces

mode Expression;

CLOSE_EXPR
  : '}}' -> pushMode(DEFAULT_MODE)
  ;

EXPR
  : ( ~'}' | '}' ~'}' )+
  ;

// ----------------------------------------------------
// statement: assignment, for-loop or if-statement

mode Statement;

CLOSE_STMT
  : WS* '%}' -> mode(DEFAULT_MODE)
  ;

FOR
  : 'for' WS+
  ;

IN 
  : 'in' WS+ -> pushMode(StatementExpression)
  ;

ENDFOR
  : WS* 'endfor' WS*
  ;

IF 
  : 'if' WS+ -> pushMode(StatementExpression)
  ;

ELSEIF 
  : 'elseif' WS+ -> pushMode(StatementExpression)
  ;

ELSE
  : 'else' WS*
  ;

ENDIF
  : 'endif' WS*
  ;

NAME 
  : [a-zA-Z_] [a-zA-Z0-9_]*
  ;

EQ
  : WS* '=' WS* -> pushMode(AssignmentExpression)
  ;

COMMA
  : ','
  ;

// WS
//   : [ \t\r\n]
//   ;

WS
  : N | S 
  ;

fragment N
  : '\r'? '\n'
  ;

fragment S 
  : [ \t]
  ;

// ----------------------------------------------------
// for-loop & if/elseif: capture all text leading up to "%}"

mode StatementExpression;

CLOSE_STMT_EXPR
  : CLOSE_STMT -> mode(DEFAULT_MODE)
  ;

STMT_EXPR
  : ( ~'%' | '%' ~'}' )+
  ;

// ----------------------------------------------------
// variable assignment: capture all text leading up to ";"

mode AssignmentExpression;

SEMICOLON
  : WS* ';' WS* -> popMode
  ;

ASSIGN_EXPR
  : ~';'+
  ;
