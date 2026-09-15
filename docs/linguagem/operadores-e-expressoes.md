# Operadores e Expressões

O Verbo aceita tanto os símbolos matemáticos tradicionais quanto palavras reservadas em português, tornando expressões complexas fáceis de ler e entender.

---

## Operadores Aritméticos

| Operador Simbólico | Alternativa em Português | Operação | Exemplo |
| :---: | :---: | :--- | :--- |
| `+` | `mais` / `soma` | Adição / Concatenação | `10 + 5` ou `10 mais 5` |
| `-` | `menos` / `subtrai` | Subtração | `20 - 4` ou `20 menos 4` |
| `*` | `multiplica` | Multiplicação | `3 * 7` ou `3 multiplica 7` |
| `/` | `divide` | Divisão | `50 / 2` ou `50 divide 2` |
| `%` | `módulo` / `resto` | Resto da divisão | `10 % 3` ou `10 módulo 3` |

### Concatenação de Texto
O operador `+` também concatena strings e converte valores quando apropriado:
```verbo
O nome é "Maria".
O ola é "Olá, " + nome + "!".
```

---

## Operadores de Comparação / Relacionais

| Operador Simbólico | Frase em Português | Significado | Exemplo |
| :---: | :---: | :--- | :--- |
| `==` | `igual a` | Igualdade | `x == 10` ou `x igual a 10` |
| `!=` | `diferente de` | Desigualdade | `x != 0` ou `x diferente de 0` |
| `<` | `menor que` | Menor que | `x < y` ou `x menor que y` |
| `>` | `maior que` | Maior que | `x > y` ou `x maior que y` |
| `<=` | `menor ou igual a` | Menor ou igual | `x <= 100` |
| `>=` | `maior ou igual a` | Maior ou igual | `x >= 18` |

---

## Operadores Lógicos

| Operador | Significado | Exemplo |
| :---: | :--- | :--- |
| `e` | Conjunção lógica (AND) | `Se idade >= 18 e ativo == Verdadeiro então:` |
| `ou` | Disjunção lógica (OR) | `Se erro ou cancelado então:` |
| `não` *(ou `nao`)* | Negação lógica (NOT) | `Se não valido então:` |

---

## Precedência de Operadores

A avaliação de expressões segue as regras matemáticas e de precedência canônicas:

1. **Agrupamento**: Expressões entre parênteses `( ... )`.
2. **Negação / Unário**: `não`, `-`.
3. **Multiplicativos**: `*`, `/`, `%`.
4. **Aditivos**: `+`, `-`.
5. **Relacionais**: `<`, `>`, `<=`, `>=`, `igual`, `diferente`.
6. **Lógicos**: `e`, `ou`.

### Exemplo:
```verbo
O resultado é (10 + 5) * 2.
Exibir com (resultado). // 30
```

