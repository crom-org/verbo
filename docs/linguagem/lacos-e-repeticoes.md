# Laços e Repetições (Loops)

O Verbo possui três formas principais de repetição: repetição com contagem fixa (`Repita N vezes`), iteração sobre listas (`Repita para cada`) e repetição condicional (`Enquanto`).

---

## 1. Repetição por Contagem (`Repita N vezes`)

Para executar um bloco de código um número fixo de vezes:

```verbo
Repita 5 vezes :
    Exibir "Executando iteração...".
.
```

O compilador transpilada isso diretamente para um loop `for i := 0; i < N; i++` otimizado em Go.

---

## 2. Iteração sobre Listas (`Repita para cada`)

No Verbo 2.0, é possível percorrer cada elemento de uma lista facilmente com `Repita para cada <item> em <coleção>`:

```verbo
Uma compras está ["Maçã", "Pão", "Café", "Queijo"].

Repita para cada item em compras :
    Exibir "Item da lista: " + item.
.
```

---

## 3. Repetição Condicional (`Enquanto`)

Executa o bloco de instruções enquanto uma determinada condição for verdadeira:

```verbo
Um contador está 1.

Enquanto contador <= 5 :
    Exibir "Contador: " + contador.
    contador está contador + 1.
.
```

---

## Aninhamento de Laços

Laços podem ser combinados livremente para iteração matricial ou processamento em múltiplos níveis:

```verbo
Uma matriz está [
    [1, 2, 3],
    [4, 5, 6]
].

Repita para cada linha em matriz :
    Repita para cada valor em linha :
        Exibir valor.
    .
.
```

