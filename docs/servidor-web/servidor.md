# Servidor HTTP Nativo

O Verbo 3.0 inclui um servidor HTTP estilo Flask: declare a instância, registre rotas e inicie. O CLI `verbo servir` transpila e sobrescreve host/porta via ambiente.

---

## Criar o servidor

```verbo
Um app está Servidor com (local, 8080).
Um app está Servidor com (externo, 5000).
Um app está Servidor com (endereço: local, porta: 8080).
```

| Palavra | IP |
| :--- | :--- |
| `local` | `127.0.0.1` (só esta máquina) |
| `externo` | `0.0.0.0` (todas as interfaces) |

A instância precisa ser **iniciada** (`app iniciar.` ou `app rodar.`). Sem isso, o transpiler omite o código do servidor (evita variáveis não usadas).

---

## Rotas

```verbo
app rota GET em "/":
    Exibir "Olá, Mundo!".
.

app rota POST em "/api/echo":
    Exibir "{\"ok\":true}".
.

app rota PUT em "/api/item":
    Exibir "atualizado".
.

app rota DELETE em "/api/item":
    Exibir "removido".
.
```

Atalho com a palavra-chave `Servidor` (usa a instância chamada `servidor`):

```verbo
Um servidor está Servidor com (local, 8080).

Servidor rota GET em "/saude":
    Exibir "ok".
.

servidor iniciar.
```

Métodos: `GET`, `POST`, `PUT`, `DELETE`. Método diferente do declarado responde **405**.

Dentro do handler, `Exibir` escreve no `http.ResponseWriter` (`fmt.Fprint`), não no stdout.

---

## Arquivos estáticos

O mux sempre registra:

| URL | Disco |
| :--- | :--- |
| `/static/*` | `./site/static/` |
| `GET /` (se não houver rota `"/"`) | `./site/index.html` |

Coloque CSS/JS/imagens em `site/static/` e a home em `site/index.html`.

---

## Iniciar

```verbo
app iniciar.
app rodar.
```

Equivalente a `http.ListenAndServe(addr, mux)`.

---

## CLI `verbo servir`

```bash
verbo servir app.vrb
verbo servir app.vrb --host 0.0.0.0 --porta 5000
```

O comando define `VERBO_HOST` e `VERBO_PORTA`. O binário gerado lê essas variáveis e **sobrescreve** o host/porta do código.

Padrão: `127.0.0.1:5000`.

---

## Exemplo completo

```verbo
Incluir Html.

Um app está Servidor com (local, 8099).

app rota GET em "/":
    O corpo é CriarElemento de Html com ("h1", "Verbo no ar").
    O pagina é CriarPagina de Html com ("Início", corpo).
    Exibir pagina.
.

app rota GET em "/saude":
    Exibir "{\"status\":\"ok\"}".
.

app iniciar.
```

```bash
verbo servir app.vrb --porta 8099
```
