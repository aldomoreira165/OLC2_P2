grammar SwiftGrammar; 
import SwiftLexer; 

s: block EOF;

block: (stmt PTCOMA?)*
     ;

blockFunc: (stmt PTCOMA?)*
     ;

stmt: printstmt
    | typedDeclstmt
    | untypedDeclstmt
    | optionalTypedDeclstmt
    | asignstmt
    | ifstmt
    | switchstmt
    | whilestmt
    | forstmt
    | guardstmt
    | opasignstmt
    | funcdclstmt
    | accfuncstm
    | breakstmt
    | continuestmt
    | returnstmt
    | declvectorstmt
    | accesovectorstmt
    | appendvectorstmt
    | removelastvectorstmt
    | removeatvectorstmt
    | asignvectorstmt
    | declmatrizstmt
    | asignmatrizstmt
    | defstructstmt
    | asignstructstmt
    | struct_expr
    | vector_struct_stmt
    | append_vector_struct_stmt
    ;

// struct
defstructstmt
    : STRUCT ID LLAVEIZQ lista_atributos* LLAVEDER #DefStruct
    ;

lista_atributos
    : (LET|VAR) ID (DOSPUNTOS (tipo|ID))? (IG expr)? #AtributoStruct
    ;

struct_expr
    : ST (VAR|LET) ID (DOSPUNTOS ID)? IG valor_struct_expr #StructExpr
    | ST (VAR|LET) ID (DOSPUNTOS ID)? IG ID? #StructExprID 
    ;

valor_struct_expr
    : ST ID ( PARIZQ l_dupla PARDER)? #ValorStructExpr
    ;

l_dupla
    : ID DOSPUNTOS expr (COMA ID DOSPUNTOS expr)* #Duplastruct
    ;

accesostructstmt
    : ID (PUNTO ID)+ #AccesoStruct
    ;

asignstructstmt
    : ID (PUNTO ID)+ IG expr #AsignStruct
    ;

vector_struct_stmt
    : VAR ID IG CORCHIZQ ID CORCHDER PARIZQ PARDER #VectorStruct
    ;

acceso_vector_struct_stmt
    : ID CORCHIZQ expr CORCHDER PUNTO ID #AccesoVectorStruct
    ;

append_vector_struct_stmt
    : ST ID PUNTO APPEND PARIZQ valor_struct_expr PARDER #AppendVectorStruct
    ;

//vectores

declvectorstmt
    : VAR ID DOSPUNTOS CORCHIZQ tipo CORCHDER defvectorstmt
    ;

defvectorstmt
    : IG CORCHIZQ listaexpresiones? CORCHDER #DefVector
    | IG ID                                 #DefVectorID
    ;

listaexpresiones
    : expr (COMA expr)*
    ;

accesovectorstmt
    : ID CORCHIZQ expr CORCHDER
    ;

asignvectorstmt
    : ID CORCHIZQ expr CORCHDER IG expr
    ;

appendvectorstmt
    : ID PUNTO APPEND PARIZQ expr PARDER
    ;

removeatvectorstmt
    : ID PUNTO REMOVE PARIZQ AT DOSPUNTOS expr PARDER
    ;

removelastvectorstmt
    : ID PUNTO REMOVE_LAST PARIZQ PARDER
    ;

countvectorstmt
    : ID PUNTO COUNT
    ;

isemptyvectorstmt
    : ID PUNTO ISEMPTY
    ;

//matrices

declmatrizstmt
    : VAR ID (DOSPUNTOS tipomatriz)? IG listavaloresmatriz #declmatrizstmt2
    | VAR ID (DOSPUNTOS tipomatriz)? IG listavaloresmatriz3 #declmatrizstmt3
    ;

tipomatriz
    : CORCHIZQ CORCHIZQ CORCHIZQ tipo CORCHDER CORCHDER CORCHDER #tipomatriz3
    | CORCHIZQ CORCHIZQ tipo CORCHDER CORCHDER #tipomatriz2
    ;

listavaloresmatriz
    : CORCHIZQ CORCHIZQ listaexpresiones CORCHDER (COMA CORCHIZQ listaexpresiones CORCHDER)+ CORCHDER #listavaloresmatriz2
    ;

listavaloresmatriz3
    : CORCHIZQ listavaloresmatriz (COMA listavaloresmatriz)* CORCHDER
    ;

accesomatriz
    : ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER #accesomatriz2
    | ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER #accesomatriz3
    ;

asignmatrizstmt
    : ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER IG expr #asignmatrizstmt2
    | ID CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER CORCHIZQ expr CORCHDER IG expr #asignmatrizstmt3
    ;

//de flujo

returnstmt
    : RETURN expr PTCOMA
    ;

breakstmt
    : BREAK
    ;

continuestmt
    : CONTINUE
    ;

//funciones embebidas

printstmt
    : PRINT PARIZQ listaexpresiones PARDER
    ;

intstmt
    : INT PARIZQ expr PARDER
    ;

floatstmt
    : FLOAT PARIZQ expr PARDER
    ;

stringstmt
    : PSTRING PARIZQ expr PARDER
    ;

//==============

//funciones

funcdclstmt
    : FUNC ID PARIZQ parametros? PARDER LLAVEIZQ blockFunc LLAVEDER #FuncionNormal
    | FUNC ID PARIZQ parametros? PARDER SUB MAYOR tipo LLAVEIZQ blockFunc LLAVEDER #FuncionRetorno
    ;

accfuncstm
    : ID PARIZQ parametroscall? PARDER
    ;

parametros
    : ID ID DOSPUNTOS INOUT? tipo (COMA ID ID DOSPUNTOS INOUT? tipo)*
    ;

parametroscall
    : ID DOSPUNTOS AMPERSON? expr (COMA ID DOSPUNTOS AMPERSON? expr)*
    ;

//=============================

//sentencias de control de flujo

ifstmt  
    : IF expr LLAVEIZQ blockFunc LLAVEDER (elseifstmt)* (ELSE LLAVEIZQ blockFunc LLAVEDER)?
    ;

elseifstmt: ELSE IF expr LLAVEIZQ blockFunc LLAVEDER;

switchstmt: SWITCH expr LLAVEIZQ  caseStmt+ defaultCase? LLAVEDER;

caseStmt: CASE expr DOSPUNTOS blockFunc;

defaultCase: DEFAULT DOSPUNTOS blockFunc;

typedDeclstmt
    : (VAR|LET) ID DOSPUNTOS tipo IG expr 
    ;

untypedDeclstmt
    : (VAR|LET) ID IG expr                         
    ;

optionalTypedDeclstmt
    : (VAR|LET) ID DOSPUNTOS tipo INTERROGACION                          
    ;

asignstmt
    : ID IG expr
    ;

whilestmt
    : WHILE expr LLAVEIZQ blockFunc LLAVEDER
    ;

forstmt
    : FOR ID IN expr PUNTO PUNTO PUNTO expr LLAVEIZQ blockFunc LLAVEDER #ForRange
    ;

guardstmt
    : GUARD expr ELSE LLAVEIZQ blockFunc LLAVEDER
    ;

opasignstmt
    : expr ADD IG expr                           #Incremento
    | expr SUB IG expr                           #Decremento  
    ;

expr
    : PARIZQ expr PARDER                        # ParExpr
    | SUB expr                                  # UnariaExpr
    | op=NOT left=expr                                  # OpExpr
    | left=expr op=(DIV|MOD|MUL) right=expr     # OpExpr
    | left=expr op=(ADD|SUB) right=expr         # OpExpr
    | left=expr op=(MAY_IG|MAYOR) right=expr    # OpExpr
    | left=expr op=(MEN_IG|MENOR) right=expr    # OpExpr
    | left=expr op=(IG_IG|DIF) right=expr       # OpExpr
    | left=expr op=AND right=expr               # OpExpr
    | left=expr op=OR right=expr                # OpExpr
    | NUMBER                                    # NumExpr
    | ID                                        # IdExpr
    | STRING                                    # StrExpr    
    | (TRU | FAL)                               # BoolExpr
    | NIL                                       # NilExpr
    | accfuncstm                                # AccFuncExpr
    | intstmt                                   # IntExpr
    | floatstmt                                 # FloatExpr
    | stringstmt                                # StringExpr
    | accesovectorstmt                          # AccesoVectorExpr
    | countvectorstmt                           # CountVectorExpr
    | isemptyvectorstmt                         # IsEmptyVectorExpr
    | accesomatriz                              # AccesoMatrizExpr
    | accesostructstmt                          # AccesoStructExpr
    | valor_struct_expr                         # AccesoValorStructExpr
    | acceso_vector_struct_stmt                 # AccesoVectorStructExpr
    ;

tipo
    : INT                                      
    | FLOAT                                     
    | BOOL                                      
    | CHARACTER                                 
    | PSTRING
    | tipo_vector  
    | tipo_matriz2
    | tipo_matriz3                               
    ;

tipo_vector
    : CORCHIZQ INT CORCHDER                                      
    | CORCHIZQ FLOAT CORCHDER                                     
    | CORCHIZQ BOOL CORCHDER                                      
    | CORCHIZQ CHARACTER CORCHDER                                 
    | CORCHIZQ PSTRING CORCHDER                                  
    ;

tipo_matriz2
    : CORCHIZQ CORCHIZQ INT CORCHDER CORCHDER                                      
    | CORCHIZQ CORCHIZQ FLOAT CORCHDER CORCHDER                                     
    | CORCHIZQ CORCHIZQ BOOL CORCHDER CORCHDER                                      
    | CORCHIZQ CORCHIZQ CHARACTER CORCHDER CORCHDER                                 
    | CORCHIZQ CORCHIZQ PSTRING CORCHDER CORCHDER                                  
    ;

tipo_matriz3
    : CORCHIZQ CORCHIZQ CORCHIZQ INT CORCHDER CORCHDER CORCHDER                                      
    | CORCHIZQ CORCHIZQ CORCHIZQ FLOAT CORCHDER CORCHDER CORCHDER                                     
    | CORCHIZQ CORCHIZQ CORCHIZQ BOOL CORCHDER CORCHDER CORCHDER                                      
    | CORCHIZQ CORCHIZQ CORCHIZQ CHARACTER CORCHDER CORCHDER CORCHDER                                 
    | CORCHIZQ CORCHIZQ CORCHIZQ PSTRING CORCHDER CORCHDER CORCHDER                                  
    ;