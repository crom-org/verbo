# Módulo Html

Geração de HTML tipada. Útil para páginas estáticas, respostas de rotas e e-mails.

```verbo
Incluir Html.
```

---

## Funções

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `CriarElemento` | `(tag: Texto, conteudo: Texto) → Texto` | `<tag>conteudo</tag>` |
| `CriarElementoComAtributos` | `(tag: Texto, atributos: Texto, conteudo: Texto) → Texto` | Tag com atributos |
| `CriarPagina` | `(titulo: Texto, corpo: Texto) → Texto` | Documento HTML5 completo (`lang="pt-BR"`) |
| `CriarPaginaComEstilo` | `(titulo: Texto, estilo: Texto, corpo: Texto) → Texto` | Página com `<style>` |
| `Atributo` | `(chave: Texto, valor: Texto) → Texto` | `chave="valor"` |
| `ListaElementos` | `(...elementos: Texto) → Texto` | Concatena com quebra de linha |
| `CriarLista` | `(...itens: Texto) → Texto` | `<ul><li>...</li></ul>` |
| `CriarTabela` | `(headers, linhas) → Texto` | Tabela com `<thead>` / `<tbody>` |
| `CriarLink` | `(url: Texto, texto: Texto) → Texto` | `<a href="...">` |
| `CriarImagem` | `(src: Texto, alt: Texto) → Texto` | `<img src="..." alt="...">` |

`CriarTabela` aceita listas Verbo (`[]interface{}`) tanto nos cabeçalhos quanto nas linhas.

---

## Exemplos

```verbo
Incluir Html.

O pagina é CriarPagina de Html com ("Meu Portfólio", "<h1>Bem-vindo ao Verbo!</h1>").
Exibir com (pagina).

O botao é CriarElemento de Html com ("button", "Clique aqui").
O link é CriarLink de Html com ("https://github.com/MrJc01/crom-verbo", "GitHub do Verbo").
O img é CriarImagem de Html com ("/static/logo.png", "Logo Verbo").

Exibir botao.
Exibir link.
Exibir img.
```

### Página com estilo e lista

```verbo
Incluir Html.

O css é "body { font-family: sans-serif; }".
O lista é CriarLista de Html com ("Instalar", "Compilar", "Executar").
O corpo é "<h1>Passos</h1>" + lista.
O pagina é CriarPaginaComEstilo de Html com ("Guia", css, corpo).
Exibir pagina.
```

### Tabela

```verbo
Incluir Html.

Uma cabecalhos está ["Produto", "Preço"].
Uma linhas está [
    ["Café", "15"],
    ["Chá", "8"]
].
O tabela é CriarTabela de Html com (cabecalhos, linhas).
Exibir tabela.
```

---

## Integração com o servidor web

Dentro de uma rota, `Exibir` escreve no `http.ResponseWriter`. Combine Html + rotas:

```verbo
Incluir Html.

Um app está Servidor com (local, 8080).

app rota GET em "/":
    O corpo é CriarElemento de Html com ("h1", "Olá do Verbo").
    O pagina é CriarPagina de Html com ("Início", corpo).
    Exibir pagina.
.

app iniciar.
```
