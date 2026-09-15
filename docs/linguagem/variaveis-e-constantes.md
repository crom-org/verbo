# Variáveis e Constantes (Ser vs. Estar)

No Verbo, a mutabilidade de dados não é controlada por palavras técnicas estrangeiras, mas pela conjugação verbal da língua portuguesa: a oposição clássica entre **Ser** (essência definitiva) e **Estar** (estado transitório).

---

## Constantes (Imutáveis)

Para declarar uma constante cujo valor nunca muda após a definição, utiliza-se um **artigo definido** (`O`, `A`, `Os`, `As`) combinado com o verbo **`é`**:

```verbo
O pi é 3.14159.
A mensagem é "Bem-vindo ao Verbo".
O limite_maximo é 100.
```

### Regras de Imutabilidade
- Uma constante declarada com `é` **não pode ser reatribuída**.
- O compilador acusa erro caso ocorra uma tentativa de reatribuição:
```verbo
O nome é "Carlos".
nome está "Roberto". // ❌ Erro do compilador: tentativa de mutar constante imutável
```

---

## Variáveis (Mutáveis)

Para valores que podem mudar ao longo do tempo, utiliza-se um **artigo indefinido** (`Um`, `Uma`, `Uns`, `Umas`) combinado com o verbo **`está`**:

```verbo
Um contador está 0.
Uma temperatura está 24.5.
Um status está "iniciando".
```

### Reatribuição de Estado
Para atualizar o valor de uma variável existente, utiliza-se o identificador diretamente acompanhado de **`está`**:

```verbo
Um total está 10.
total está total + 5.
total está 50.
```

---

## Tabela Comparativa

| Tipo de Dado | Artigo | Verbo na Criação | Verbo na Atualização | Equivalente em Go |
| :--- | :--- | :--- | :--- | :--- |
| **Constante** | `O`, `A`, `Os`, `As` | `é` | *(Não permitido)* | `var x = val` (imutável na semântica Verbo) |
| **Variável** | `Um`, `Uma`, `Uns`, `Umas` | `está` | `está` | `x := val` / `x = novo_val` |

---

## Exemplo Completo

```verbo
// Constantes do sistema
A taxa_juros é 0.05.
O nome_banco é "Banco Verbo".

// Variáveis de saldo
Um saldo está 1000.0.
Exibir "Saldo inicial:".
Exibir saldo.

// Depósito
saldo está saldo + 500.0.
Exibir "Saldo após depósito:".
Exibir saldo.
```

