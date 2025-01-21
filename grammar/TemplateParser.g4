parser grammar TemplateParser;

options {
    tokenVocab = TemplateLexer;
}

template
  : node* EOF
  ;

node 
  : expression | statement | text
  ;

expression
  : OPEN_EXPR expr CLOSE_EXPR
  ;

statement
  : assignment | forLoop | ifStatement
  ;

assignment
  : OPEN_STMT var+ CLOSE_STMT
  ;
  
var 
  : name EQ expr SEMICOLON
  ;

forLoop
  : startFor node* endFor
  ;

startFor
  : OPEN_STMT FOR ( name | index WS* COMMA WS* name ) WS+ IN expr CLOSE_STMT_EXPR
  ;

endFor
  : OPEN_STMT ENDFOR CLOSE_STMT
  ;

ifStatement
  : startIf elseIf* else? endIf
  ;

startIf
  : OPEN_STMT IF expr CLOSE_STMT_EXPR node*
  ;

elseIf
  : OPEN_STMT ELSEIF expr CLOSE_STMT_EXPR node*
  ;

else 
  : OPEN_STMT ELSE CLOSE_STMT node*
  ;

endIf
  : OPEN_STMT ENDIF CLOSE_STMT
  ;

name 
  : NAME 
  ;

index
  : NAME
  ;

expr
  : EXPR | STMT_EXPR | ASSIGN_EXPR
  ;

text
  : TEXT
  ;