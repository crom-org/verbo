# Módulo Texto

Manipulação de cadeias UTF-8.

```verbo
Incluir Texto.
```

---

## Funções

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `Tamanho` | `(s: Texto) → Inteiro` | Quantidade de bytes da string |
| `Maiusculas` | `(s: Texto) → Texto` | Converte para maiúsculas |
| `Minusculas` | `(s: Texto) → Texto` | Converte para minúsculas |
| `Contem` | `(s: Texto, substr: Texto) → Lógico` | `Verdadeiro` se `s` contém `substr` |
| `Substituir` | `(s: Texto, velho: Texto, novo: Texto) → Texto` | Substitui todas as ocorrências |
| `Dividir` | `(s: Texto, separador: Texto) → Lista` | Parte a string pelo separador |

---

## Exemplos

```verbo
Incluir Texto.

O frase é "Olá, Mundo Verbo".

O n é Tamanho de Texto com (frase).
O alto é Maiusculas de Texto com (frase).
O baixo é Minusculas de Texto com (frase).
O tem é Contem de Texto com (frase, "Mundo").
O trocado é Substituir de Texto com (frase, "Mundo", "Brasil").
Uma partes é Dividir de Texto com (frase, " ").

Exibir n.
Exibir alto.
Exibir tem.
Exibir trocado.
Exibir partes[0].
```

Uso típico com entidades:

```verbo
Incluir Texto.

A entidade Pessoa contendo (Nome: Texto, Idade: Inteiro).

Para Saudacao usando (pessoa: Pessoa):
    Um nome_maiusculo é Maiusculas de Texto com (Nome de pessoa).
    Exibir "Olá " + nome_maiusculo + "!".
.
```

---

## Equivalente Go

| Verbo | Go (`strings`) |
| :--- | :--- |
| `Tamanho` | `len(s)` |
| `Maiusculas` | `strings.ToUpper` |
| `Minusculas` | `strings.ToLower` |
| `Contem` | `strings.Contains` |
| `Substituir` | `strings.ReplaceAll` |
| `Dividir` | `strings.Split` → `[]interface{}` |
