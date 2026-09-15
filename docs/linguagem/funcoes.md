# Funções

No Verbo, toda função é um **verbo no infinitivo**. A declaração usa `Para`, os parâmetros entram com `usando` e a chamada usa a preposição `com`.

---

## Declaração

```verbo
Para Somar usando (x: Inteiro, y: Inteiro):
    Retorne x + y.
.
```

| Parte | Papel |
| :--- | :--- |
| `Para` | Inicia a declaração |
| `Somar` | Nome da função (verbo no infinitivo) |
| `usando (x: Inteiro, y: Inteiro)` | Lista de parâmetros tipados (opcional) |
| `:` ... `.` | Corpo do bloco |

Parâmetros sem anotação de tipo também são aceitos:

```verbo
Para Dobrar usando (n):
    Retorne n * 2.
.
```

### Função sem parâmetros

```verbo
Para principal:
    Exibir com ("Programa iniciado.").
.
```

---

## Chamada

A preposição `com` liga o nome da função aos argumentos:

```verbo
Um resultado é Somar com (10, 5).
Exibir com (resultado).

Somar com (3, 7).
principal com ().
```

`Exibir` aceita tanto `Exibir com (valor).` quanto `Exibir valor.`.

---

## Retorno

```verbo
Para Maior usando (a: Inteiro, b: Inteiro):
    Se a > b então:
        Retorne a.
    Senão:
        Retorne b.
    .
.

O vencedor é Maior com (8, 15).
```

`Retorne Nulo.` (ou `Retorne.` sem valor) encerra a função sem produzir resultado.

---

## Chamadas de módulo (BibVerbo)

Funções da biblioteca padrão usam a preposição `de` para indicar o pacote:

```verbo
Incluir Matematica.

O cubo é Potencia de Matematica com (2.0, 3.0).
```

Isso transpilada para `matematica.Potencia(2.0, 3.0)`.

---

## Código Go gerado

| Verbo | Go |
| :--- | :--- |
| `Para Somar usando (x: Inteiro, y: Inteiro):` | `func Somar(x int, y int) interface{} {` |
| `Retorne x + y.` | `return x + y` |
| `Somar com (10, 5).` | `Somar(10, 5)` |

---

## Exemplo completo

```verbo
Para Saudar usando (nome: Texto):
    Exibir com ("Bem-vindo, " + nome + "!").
.

Para CalcularMedia usando (a: Decimal, b: Decimal):
    Retorne (a + b) / 2.
.

Saudar com ("Brasil").

O media é CalcularMedia com (8.5, 9.5).
Exibir com (media).
```
