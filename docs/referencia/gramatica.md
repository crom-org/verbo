# Gramática Formal EBNF (Verbo 2.0 / 3.0)

Esta é a especificação formal completa da gramática da Linguagem Verbo na notação **EBNF** (*Extended Backus-Naur Form*), cobrindo variáveis, entidades, funções, concorrência, tratamento de erros, biblioteca padrão e servidor web.

---

## 1. Estrutura do Programa

```ebnf
programa            = { declaracao_topo } ;

declaracao_topo     = decl_incluir
                    | decl_entidade
                    | decl_funcao
                    | decl_instrucao ;

declaracao_instrucao = decl_variavel
                    | decl_atribuicao
                    | decl_servidor
                    | decl_rota
                    | decl_iniciar_servidor
                    | decl_se
                    | decl_repita
                    | decl_enquanto
                    | decl_simultaneamente
                    | decl_tente
                    | decl_sinalize
                    | decl_enviar_canal
                    | decl_exibir
                    | decl_retorne
                    | decl_expressao ;

bloco               = { declaracao_instrucao } ;
```

---

## 2. Inclusão de Módulos e Entidades

```ebnf
decl_incluir        = "Incluir" IDENTIFICADOR PONTO ;

decl_entidade       = artigo "entidade" IDENTIFICADOR "contendo" "(" lista_campos ")" PONTO ;

lista_campos        = campo { VIRGULA campo } ;
campo               = IDENTIFICADOR DOIS_PONTOS TIPO ;

instanciacao_ent    = [ artigo ] "novo" IDENTIFICADOR "contendo" "(" lista_expressoes ")"
                    | IDENTIFICADOR "com" "(" lista_expressoes ")" ;
```

---

## 3. Variáveis e Atribuições

```ebnf
decl_variavel       = artigo IDENTIFICADOR verbo_atrib expressao PONTO ;

artigo              = ARTIGO_DEFINIDO       (* O | A | Os | As *)
                    | ARTIGO_INDEFINIDO ;   (* Um | Uma | Uns | Umas *)

verbo_atrib         = "é" | "está" ;

decl_atribuicao     = IDENTIFICADOR "está" expressao PONTO ;
```

---

## 4. Funções

```ebnf
decl_funcao         = "Para" IDENTIFICADOR [ "usando" "(" [ lista_parametros ] ")" ] DOIS_PONTOS
                      bloco
                      PONTO ;

lista_parametros    = parametro { VIRGULA parametro } ;
parametro           = IDENTIFICADOR [ DOIS_PONTOS TIPO ] ;

chamada_funcao      = [ IDENTIFICADOR "de" ] IDENTIFICADOR "com" "(" [ lista_expressoes ] ")" ;
```

---

## 5. Controle de Fluxo e Laços

```ebnf
decl_se             = "Se" expressao_condicional [ "então" ] DOIS_PONTOS
                      bloco
                      [ "Senão" DOIS_PONTOS bloco ]
                      PONTO ;

expressao_condicional = [ artigo ] expressao [ "for" ] op_relacional expressao
                      | expressao ;

decl_repita         = "Repita" expressao "vezes" DOIS_PONTOS bloco PONTO
                    | "Repita" "para" "cada" IDENTIFICADOR "em" expressao DOIS_PONTOS bloco PONTO ;

decl_enquanto       = "Enquanto" expressao_condicional DOIS_PONTOS bloco PONTO ;
```

---

## 6. Concorrência, Canais e Tratamento de Erros

```ebnf
decl_simultaneamente = "Simultaneamente" DOIS_PONTOS bloco PONTO ;

expressao_canal     = "Canal" "de" IDENTIFICADOR ;

decl_enviar_canal   = "Enviar" expressao "para" IDENTIFICADOR PONTO ;

expressao_receber   = "Receber" "de" IDENTIFICADOR ;

decl_tente          = "Tente" DOIS_PONTOS
                      bloco
                      [ "Capture" [ IDENTIFICADOR ] DOIS_PONTOS bloco ]
                      PONTO ;

decl_sinalize       = "Sinalize" [ "com" ] ( "(" expressao ")" | expressao ) PONTO ;
```

---

## 7. Servidor Web

```ebnf
decl_servidor       = [ artigo ] IDENTIFICADOR "está" "Servidor" "com" "(" param_servidor VIRGULA param_servidor ")" PONTO ;

param_servidor      = [ "endereço" DOIS_PONTOS | "porta" DOIS_PONTOS ] ( "local" | "externo" | expressao ) ;

decl_rota           = ( IDENTIFICADOR | "Servidor" ) "rota" metodo_http "em" TEXTO DOIS_PONTOS
                      bloco
                      PONTO ;

metodo_http         = "GET" | "POST" | "PUT" | "DELETE" ;

decl_iniciar_servidor = IDENTIFICADOR ( "iniciar" | "rodar" ) PONTO ;
```

---

## 8. Expressões e Operadores

```ebnf
expressao           = expr_ou ;

expr_ou             = expr_e { "ou" expr_e } ;
expr_e              = expr_relacional { "e" expr_relacional } ;
expr_relacional     = expr_aditiva { op_relacional expr_aditiva } ;
expr_aditiva        = expr_multiplicativa { op_aditivo expr_multiplicativa } ;
expr_multiplicativa = expr_unaria { op_multiplicativo expr_unaria } ;

expr_unaria         = ( "não" | "-" ) expr_unaria
                    | expr_primaria ;

expr_primaria       = NUMERO
                    | TEXTO
                    | "Verdadeiro" | "Falso" | "Nulo"
                    | instanciacao_ent
                    | expressao_canal
                    | expressao_receber
                    | expressao_lista
                    | expressao_acesso
                    | chamada_funcao
                    | "(" expressao ")"
                    | [ artigo ] IDENTIFICADOR ;

expressao_lista     = "[" [ expressao { VIRGULA expressao } ] "]" ;

expressao_acesso    = IDENTIFICADOR "de" expressao
                    | expressao "[" expressao "]" ;

op_relacional       = "==" | "!=" | "<" | ">" | "<=" | ">="
                    | "igual" [ "a" ] | "idêntico" [ "a" ] | "diferente" [ "de" ]
                    | "menor" [ "que" | "ou" "igual" "a" ]
                    | "maior" [ "que" | "ou" "igual" "a" ] ;

op_aditivo          = "+" | "-" | "mais" | "soma" | "menos" | "subtrai" ;
op_multiplicativo   = "*" | "/" | "%" | "multiplica" | "divide" | "módulo" | "porcentagem" ;
```

---

## 9. Terminais Léxicos

```ebnf
IDENTIFICADOR       = ( letra | "_" ) { letra | digito | "_" } ;
NUMERO              = digito { digito } [ "." digito { digito } ] ;
TEXTO               = '"' { caractere_unicode } '"' ;
TIPO                = "Texto" | "Inteiro" | "Decimal" | "Logico" | "Lógico" | "Lista" ;

PONTO               = "." ;
DOIS_PONTOS         = ":" ;
VIRGULA             = "," ;

ARTIGO_DEFINIDO     = "O" | "A" | "Os" | "As" | "o" | "a" | "os" | "as" ;
ARTIGO_INDEFINIDO   = "Um" | "Uma" | "Uns" | "Umas" | "um" | "uma" ;
```
