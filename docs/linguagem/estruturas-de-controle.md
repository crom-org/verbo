# Estruturas de Controle de Fluxo

O Verbo oferece estruturas claras e idiomáticas para ramificação condicional e checagens de guarda.

---

## 1. Condicional `Se ... então` e `Senão`

A estrutura condicional básica inicia com `Se`, seguida por uma condição, a palavra-chave `então` (ou `entao`), dois pontos `:` para abrir o bloco e um ponto final `.` para fechar:

```verbo
O nota é 8.5.

Se nota >= 7.0 então :
    Exibir "Parabéns, você foi aprovado!".
.
```

### Com cláusula `Senão`
```verbo
O saldo é 50.0.
O valor_saque é 120.0.

Se saldo >= valor_saque então :
    Exibir "Saque autorizado.".
Senão :
    Exibir "Saldo insuficiente para o saque.".
.
```

### O Subjuntivo `for`
Em português culto, é natural expressar a condição no modo subjuntivo (`Se x for maior que y`). O Verbo aceita opcionalmente a palavra `for`:

```verbo
Se x for maior que 10 então :
    Exibir "x é grande".
.
```

---

## 2. Cláusulas de Guarda (`Dado`)

Inspirado na escrita de especificações matemáticas e testes formais, o comando `Dado` estabelece uma premissa prévia. Se a condição for falsa, o fluxo pode ser interrompido ou desviado:

```verbo
Dado idade >= 18.
Exibir "Acesso permitido para maiores de idade.".
```

---

## 3. Blocos e Escopo

- Todo bloco de código em Verbo é iniciado por dois pontos `:` e delimitado ao final por um ponto final solitário `.`.
- Comentários podem acompanhar o fechamento de blocos para clareza em estruturas aninhadas:

```verbo
Se ativo então :
    Se pontos > 100 então :
        Exibir "Cliente VIP!".
    . // fim do se pontos
. // fim do se ativo
```

