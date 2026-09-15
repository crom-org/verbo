# Tipos de Dados

O Verbo possui um sistema de tipos forte, inferido estaticamente pelo compilador a partir de valores literais e validado nas anotações de parâmetros de funções e campos de entidades.

---

## Tipos Primitivos

| Tipo Verbo | Descrição | Exemplo Literal | Tipo em Go |
| :--- | :--- | :--- | :--- |
| **`Texto`** | Cadeia de caracteres UTF-8 delimitada por aspas duplas | `"Olá, Mundo!"` | `string` |
| **`Inteiro`** | Número inteiro com sinal de 64 bits | `42`, `-10`, `0` | `int` |
| **`Decimal`** | Número em ponto flutuante de precisão dupla (IEEE 754) | `3.14`, `-0.5`, `100.0` | `float64` |
| **`Lógico`** *(ou `Logico`)* | Valor booleano verdadeiro ou falso | `Verdadeiro`, `Falso` | `bool` |
| **`Nulo`** | Ausência deliberada de valor | `Nulo` | `nil` |
| **`Lista`** | Coleção ordenada heterogênea de elementos | `[1, 2, "três", Verdadeiro]` | `[]interface{}` |

---

## Exemplos de Literais

### 1. Texto
Delimitado sempre por aspas duplas:
```verbo
O saudacao é "Olá do Brasil! 🇧🇷".
O multilinhas é "Linha 1\nLinha 2".
```

### 2. Números (Inteiro e Decimal)
```verbo
O ano é 2026.
O pi é 3.1415926535.
```

### 3. Lógico
Os literais booleanos são escritos com a primeira letra maiúscula:
```verbo
O ativo é Verdadeiro.
O cancelado é Falso.
```

### 4. Listas (Verbo 2.0)
Listas são delimitadas por colchetes `[` e `]` com elementos separados por vírgula:
```verbo
Uma notas está [8.5, 9.0, 7.5, 10.0].
O nomes é ["Maria", "João", "Ana"].

// Acesso por índice (base 0)
O primeira_nota é notas[0].
O segundo_nome é nomes[1].
```

---

## Anotações de Tipos

Anotações de tipos são utilizadas na declaração de parâmetros de funções e em campos de Entidades:

```verbo
A entidade Produto contendo (
    Nome: Texto,
    Preco: Decimal,
    EmEstoque: Lógico
).

Para CalcularTotal usando (quantidade: Inteiro, unitario: Decimal) :
    Retorne quantidade * unitario.
.
```

