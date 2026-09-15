# Palavras Reservadas e Tokens

Catálogo completo de palavras-chave, artigos, conectores, tipos e delimitadores reconhecidos pelo analisador léxico (`pkg/lexer/token.go`).

---

## 1. Artigos e Semântica de Mutabilidade

| Token | Palavras-chave no Código | Função / Semântica |
| :--- | :--- | :--- |
| `TOKEN_ARTIGO_DEFINIDO` | `O`, `A`, `Os`, `As`, `o`, `a`, `os`, `as` | Declaração de constante imutável (combinado com `é`). |
| `TOKEN_ARTIGO_INDEFINIDO` | `Um`, `Uma`, `um`, `uma` | Declaração de variável mutável (combinado com `está`). |
| `TOKEN_DEMONSTRATIVO` | `Este`, `Esta`, `Aquele`, `Aquela`, `este`, `esta`, `aquele`, `aquela` | Pronomes demonstrativos de referência contextual. |

---

## 2. Verbos e Estruturas de Controle

| Token | Palavras-chave | Uso Sintático |
| :--- | :--- | :--- |
| `TOKEN_E_ACENTO` | `É`, `é` | Atribuição de essência / imutabilidade (`O total é 100.`) |
| `TOKEN_ESTA` | `Está`, `está` | Atribuição e reatribuição de estado (`Um saldo está 0.`, `saldo está 10.`) |
| `TOKEN_PARA` | `Para`, `para` | Início de declaração de função (`Para Somar:`) |
| `TOKEN_USANDO` | `usando` | Declaração de parâmetros de função (`usando (a: Inteiro)`) |
| `TOKEN_RETORNE` | `Retorne`, `retorne` | Retorno de valor de função (`Retorne x + y.`) |
| `TOKEN_EXIBIR` | `Exibir`, `exibir` | Impressão na saída padrão ou resposta HTTP |
| `TOKEN_COM` | `com` | Conector de argumentos de função/exibição (`com (10, 20)`) |
| `TOKEN_SE` | `Se`, `se` | Condicional (`Se x > 0 então:`) |
| `TOKEN_FOR` | `for` | Subjuntivo opcional em condições (`Se x for maior que y`) |
| `TOKEN_ENTAO` | `então`, `Então` | Início do bloco condicional verdadeiro |
| `TOKEN_SENAO` | `Senão`, `senão` | Ramo alternativo da condicional |
| `TOKEN_REPITA` | `Repita`, `repita` | Laço de repetição (`Repita 5 vezes:` / `Repita para cada item em lista:`) |
| `TOKEN_VEZES` | `vezes` | Complemento do laço por contagem fixa |
| `TOKEN_PARA_CADA` | `cada` | Iteração sobre coleções (`para cada`) |
| `TOKEN_EM` | `em` | Conector de pertencimento em loops e rotas HTTP |
| `TOKEN_ENQUANTO` | `Enquanto`, `enquanto` | Laço condicional |
| `TOKEN_DADO` | `Dado`, `dado` | Cláusula de guarda / premissa |
| `TOKEN_INCLUIR` | `Incluir`, `incluir` | Importação de pacote da biblioteca padrão |

---

## 3. Conectores Gramaticais e Regência

| Token | Palavras-chave | Uso Sintático |
| :--- | :--- | :--- |
| `TOKEN_DE` | `de`, `do`, `da`, `dos`, `das` | Acesso a campos (`nome de usuario`) ou pacote (`Potencia de Matematica`) |
| `TOKEN_AO` | `ao`, `aos` | Conector gramatical de destino |
| `TOKEN_NO` | `no`, `na`, `nos`, `nas` | Conector gramatical de posição |
| `TOKEN_PELO` | `pelo`, `pela`, `pelos`, `pelas` | Conector gramatical de passagem |
| `TOKEN_POR` | `por` | Conector gramatical de proporção ou divisão |

---

## 4. Concorrência e Canais (Verbo 2.0)

| Token | Palavras-chave | Uso Sintático |
| :--- | :--- | :--- |
| `TOKEN_SIMULTANEAMENTE` | `Simultaneamente`, `simultaneamente` | Bloco de execução paralela via goroutines (`sync.WaitGroup`) |
| `TOKEN_AGUARDE` | `Aguarde`, `aguarde` | Sincronização explícita |
| `TOKEN_CANAL` | `Canal`, `canal` | Tipo de canal concorrente (`Canal de Inteiros`) |
| `TOKEN_ENVIAR` | `Enviar`, `enviar` | Envio de mensagem para canal (`Enviar x para canal.`) |
| `TOKEN_RECEBER` | `Receber`, `receber` | Leitura de mensagem do canal (`Receber de canal`) |

---

## 5. Entidades e Tratamento de Erros (Verbo 2.0)

| Token | Palavras-chave | Uso Sintático |
| :--- | :--- | :--- |
| `TOKEN_ENTIDADE` | `Entidade`, `entidade` | Declaração de tipo struct (`A entidade Pessoa contendo (...)`) |
| `TOKEN_CONTENDO` | `contendo`, `Contendo` | Delimitador dos campos de uma entidade |
| `TOKEN_NOVO` | `novo`, `Novo` | Instanciação explícita (`um novo Pessoa contendo (...)`) |
| `TOKEN_TENTE` | `Tente`, `tente` | Bloco protegido contra pânico |
| `TOKEN_CAPTURE` | `Capture`, `capture` | Tratamento do pânico recuperado (`Capture erro:`) |
| `TOKEN_SINALIZE` | `Sinalize`, `sinalize` | Emissão de pânico de execução (`Sinalize "mensagem".`) |

---

## 6. Servidor Web (Verbo 3.0)

| Token | Palavras-chave | Uso Sintático |
| :--- | :--- | :--- |
| `TOKEN_SERVIDOR` | `Servidor`, `servidor` | Instanciação e rotas do servidor HTTP |
| `TOKEN_ENDERECO` | `Endereço`, `endereço`, `Endereco`, `endereco` | Parâmetro de bind IP |
| `TOKEN_PORTA` | `Porta`, `porta` | Parâmetro de porta TCP |
| `TOKEN_LOCAL` | `Local`, `local` | Atalho para o IP `127.0.0.1` |
| `TOKEN_EXTERNO` | `Externo`, `externo` | Atalho para o IP `0.0.0.0` |
| `TOKEN_ROTA` | `Rota`, `rota` | Declaração de manipulador HTTP (`app rota GET em "/":`) |
| `TOKEN_INICIAR` / `TOKEN_RODAR` | `Iniciar`, `iniciar`, `Rodar`, `rodar` | Inicialização do listener HTTP |
| `TOKEN_GET` / `POST` / `PUT` / `DELETE` | `GET`, `POST`, `PUT`, `DELETE` | Métodos HTTP suportados nas rotas |

---

## 7. Tipos de Dados Primitivos

| Token | Palavra-chave | Tipo em Go |
| :--- | :--- | :--- |
| `TOKEN_TIPO` | `Texto` | `string` |
| `TOKEN_TIPO` | `Inteiro` | `int` |
| `TOKEN_TIPO` | `Decimal` | `float64` |
| `TOKEN_TIPO` | `Logico`, `Lógico` | `bool` |
| `TOKEN_TIPO` | `Lista` | `[]interface{}` |

---

## 8. Literais e Operadores

| Categoria | Tokens e Símbolos | Alternativas Textuais em Português |
| :--- | :--- | :--- |
| **Booleanos** | `Verdadeiro`, `verdadeiro`, `Falso`, `falso` | — |
| **Nulo** | `Nulo`, `nulo` | — |
| **Aritmética** | `+`, `-`, `*`, `/`, `%` | `mais`, `soma`, `menos`, `subtrai`, `multiplica`, `divide`, `módulo`, `modulo`, `porcentagem` |
| **Comparação** | `==`, `!=`, `<`, `>`, `<=`, `>=` | `igual`, `idêntico`, `diferente`, `menor que`, `maior que` |
| **Lógicos** | `e`, `ou`, `não`, `Não` | — |
| **Delimitadores** | `.`, `:`, `,`, `(`, `)`, `[`, `]`, `{`, `}` | O ponto final `.` é obrigatório ao final de instruções |
