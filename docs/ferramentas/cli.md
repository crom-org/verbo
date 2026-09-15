# CLI do Verbo

A ferramenta `verbo` transpila `.vrb` para Go e aciona `go build` / `go run`. Versão atual: **0.1.0**.

---

## Instalação

Veja [Instalação](../comecando/instalacao.md). Atalho a partir do repositório:

```bash
make build
./build/verbo ajuda
```

---

## Comandos

| Comando | Uso | Efeito |
| :--- | :--- | :--- |
| `compilar` | `verbo compilar arquivo.vrb` | Gera `arquivo_verbo.go` e o binário `arquivo` |
| `executar` | `verbo executar arquivo.vrb` | Transpila para temp e `go run` |
| `verificar` | `verbo verificar arquivo.vrb` | Só lexer + parser (linter) |
| `servir` | `verbo servir arquivo.vrb [--host IP] [--porta N]` | Roda o servidor HTTP |
| `pacote` | `verbo pacote <instalar\|listar\|info\|remover>` | Gerenciador de pacotes |
| `versão` | `verbo versão` | Alias: `versao`, `--version`, `-v` |
| `ajuda` | `verbo ajuda` | Alias: `help`, `--help`, `-h` |

Arquivos **devem** ter extensão `.vrb`.

---

## `compilar`

```bash
verbo compilar ola_mundo.vrb
```

Saída:

- `ola_mundo_verbo.go` — código Go gerado
- `ola_mundo` — binário nativo (`go build -o`)

```bash
./ola_mundo
```

---

## `executar`

```bash
verbo executar ola_mundo.vrb
```

Escreve um arquivo temporário, executa com `go run` e remove o temp. Stdout/stderr do programa aparecem entre as linhas `────`.

---

## `verificar`

```bash
verbo verificar ola_mundo.vrb
```

```
📊 Tokens encontrados: 12
🌳 Declarações na AST: 2
✅ Arquivo 'ola_mundo.vrb' está sintaticamente correto!
```

Erros léxicos ou sintáticos saem em stderr e o processo retorna código 1. Não gera Go nem binário.

---

## `servir`

```bash
verbo servir app.vrb --host 0.0.0.0 --porta 5000
```

Flags: `--host`, `--porta` (ou `--port`). Padrão: `127.0.0.1:5000`. Detalhes em [Servidor HTTP Nativo](../servidor-web/servidor.md).

---

## `pacote`

Documentado em [Gerenciador de Pacotes](gerenciador-de-pacotes.md).

```bash
verbo pacote instalar gh:user/repo@main
verbo pacote instalar path:./meu_pacote
verbo pacote listar
verbo pacote info
verbo pacote remover meu_pacote
```

---

## Make (atalhos do repositório)

```bash
make build
make run ARQUIVO=examples/ola_mundo.vrb
make verificar ARQUIVO=examples/ola_mundo.vrb
make test
```
