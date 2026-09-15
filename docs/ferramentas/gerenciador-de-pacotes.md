# Gerenciador de Pacotes

O CLI instala dependências em `./pacotes/<nome>`, registra o manifesto e o lockfile, e opcionalmente executa `installer.vrb`.

---

## Arquivos

| Arquivo | Papel |
| :--- | :--- |
| `verbo.mod.json` | Lista de nomes instalados |
| `verbo.lock.json` | Resolução: fonte, ref e pasta |
| `./pacotes/<nome>/` | Cópia do pacote |

Exemplo de manifesto:

```json
{
  "pacotes": ["exemplo_pacote"]
}
```

---

## Instalar

Formatos:

```bash
# GitHub (zip da branch/ref)
verbo pacote instalar gh:user/repo
verbo pacote instalar gh:user/repo@main

# Pasta local (cópia)
verbo pacote instalar path:./meu_pacote
```

O nome no manifesto é o último segmento (`repo` ou o basename do path). Pacote já listado é ignorado.

---

## Listar, info e remover

```bash
verbo pacote listar
verbo pacote info
verbo pacote remover exemplo_pacote
```

`remover` apaga `pacotes/<nome>` e atualiza manifesto + lockfile.

---

## `installer.vrb`

Se existir na raiz do pacote, o CLI roda `verbo executar installer.vrb` **depois** de copiar os arquivos. Se o stdout for JSON válido, aplica as ações.

### Ações

```json
{
  "dependencias": ["gh:org/lib@main", "path:./pacotes/outro"],
  "criar_pasta": [{ "caminho": "site/static" }],
  "copiar": [{ "de": "static", "para": "site/static" }],
  "remover": [{ "caminho": "site/static/velho.css" }],
  "patch": [
    { "arquivo": "site/index.html", "procurar": "OLÁ", "trocar": "Olá", "limite": 1 }
  ],
  "allowlist": { "executar_comando": ["go", "npm"] },
  "executar_comando": [
    { "comando": "go", "args": ["fmt", "./..."], "cwd": "." }
  ]
}
```

| Campo | Significado |
| :--- | :--- |
| `dependencias` | Specs iguais ao CLI; instaladas **antes** das outras ações |
| `criar_pasta` | `mkdir -p` relativo ao projeto |
| `copiar` | `de` relativo ao pacote; `para` relativo ao projeto |
| `remover` | Apaga caminho relativo do projeto |
| `patch` | Substitui texto (`limite` = máximo de trocas) |
| `executar_comando` | Só comandos listados em `allowlist.executar_comando` |

O installer pode só imprimir texto e **não** emitir JSON — nesse caso nenhuma ação extra é aplicada.

Exemplo mínimo no repositório (`pacotes/exemplo_pacote/installer.vrb`):

```verbo
A acoes é "{\n  \"copiar\": [\n    {\n      \"de\": \"static\",\n      \"para\": \"site/static\"\n    }\n  ]\n}".

Exibir com (acoes).
```

---

## Segurança

1. Caminhos (`de`, `para`, `caminho`, `arquivo`, `cwd`) devem ser **relativos** — sem absoluto e sem `..`.
2. `executar_comando` exige allowlist explícita.
3. `remover` bloqueia `pacotes`, `verbo.mod.json` e `verbo.lock.json`.
4. Dependências cíclicas (mesma pasta na pilha do installer) abortam a instalação.

---

## Ciclo de dependências

`dependencias` usa o mesmo parser (`gh:` / `path:`). Cada dependência é instalada com a pasta atual na pilha; ciclo por pasta gera erro fatal.
