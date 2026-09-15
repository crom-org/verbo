# Extensão para VS Code

Extensão oficial na pasta `vscode-extension/`. Dá destaque de sintaxe, linter, hover, IntelliSense, outline e comandos para arquivos `.vrb`.

---

## Funcionalidades

- **TextMate grammar**: artigos, `é`/`está`, funções, V2 (entidades, concorrência, erros) e V3 (servidor, rotas HTTP).
- **Linter**: chama `verbo verificar` ao salvar e mostra diagnósticos no painel Problemas.
- **Hover**: explicações em português para palavras-chave e funções da BibVerbo.
- **IntelliSense**: palavras-chave, tipos e funções padrão.
- **Outline**: funções (`Para`), entidades e rotas.
- **Play** na barra do editor: executa o arquivo atual.

---

## Comandos

| Comando | Ação |
| :--- | :--- |
| `Verbo: Executar Arquivo Atual` | `verbo executar` no terminal integrado |
| `Verbo: Compilar Arquivo Atual` | `verbo compilar` |
| `Verbo: Iniciar Servidor Web` | `verbo servir` |
| `Verbo: Verificar Sintaxe do Arquivo` | `verbo verificar` |

Paleta: `Ctrl+Shift+P` / `Cmd+Shift+P`.

---

## Snippets

Prefixo + `Tab`:

| Prefixo | Gera |
| :--- | :--- |
| `principal` | Função principal |
| `para_usando` / `para` | Função com/sem parâmetros |
| `constante` / `variavel` | `O x é` / `Um y está` |
| `exibir` | `Exibir com (...).` |
| `se_senao` | Condicional |
| `repita_vezes` / `repita_cada` / `enquanto` | Laços |
| `entidade` / `novo` | Struct e instanciação |
| `tente` | `Tente` / `Capture` |
| `canal` / `simultaneamente` | Concorrência |
| `servidor` / `rota` | HTTP |
| `incluir` | `Incluir` de módulo |

---

## Configurações

Em Configurações, busque `Verbo`:

| Chave | Padrão | Descrição |
| :--- | :--- | :--- |
| `verbo.cliPath` | `./build/verbo` ou `PATH` | Caminho do binário |
| `verbo.verificarAoSalvar` | `true` | Linter no save |
| `verbo.servidorHost` | `127.0.0.1` | Host do `servir` |
| `verbo.servidorPorta` | `5000` | Porta do `servir` |

---

## Instalar localmente

```bash
cd vscode-extension
npm install
npm run build
npx @vscode/vsce package
code --install-extension verbo-0.1.0.vsix
```

Desenvolvimento contínuo: `npm run watch`.
