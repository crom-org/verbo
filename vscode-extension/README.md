# 🇧🇷 Verbo — Extensão para Visual Studio Code

Extensão oficial para suporte à linguagem de programação **Verbo** no Visual Studio Code.

O **Verbo** é uma linguagem de programação transpilada que utiliza a gramática da norma culta do Português Brasileiro como sintaxe lógica.

---

## ✨ Funcionalidades

- 🎨 **Destaque de Sintaxe Completo (TextMate Grammar)**:
  - Artigos como mutabilidade (`O`/`A`/`Os`/`As` para constantes imutáveis; `Um`/`Uma` para variáveis mutáveis)
  - Verbos copulativos de estado (`é` estático vs `está` mutável)
  - Funções, parâmetros e chamadas (`Para ... usando`, `... com (...)`)
  - Estruturas de controle (`Se ... então`, `Senão`, `Repita ... vezes`, `Repita para cada ... em`, `Enquanto`)
  - Concorrência e V2 (`Entidade`, `contendo`, `novo`, `Simultaneamente`, `Aguarde`, `Canal`, `Enviar`, `Receber`, `Tente`, `Capture`, `Sinalize`)
  - Servidor Web V3 (`Servidor`, `local`, `externo`, `rota`, `iniciar`, `rodar`, `GET`, `POST`, `PUT`, `DELETE`)
  - Tipos nativos (`Texto`, `Inteiro`, `Decimal`, `Lógico`, `Lista`, `Canal`)
  - Conectores e operadores em português (`de`, `do`, `da`, `ao`, `no`, `por`, `menor que`, `maior que`, `igual`, `não`, `e`, `mais`, `menos`)
- 🔍 **Validação de Sintaxe e Diagnósticos (Linter)**:
  - Integração com `verbo verificar` para exibir erros sintáticos e léxicos diretamente no editor e no painel de Problemas do VSCode ao salvar.
- 💡 **Dicas Contextuais (Hover Provider)**:
  - Passe o cursor sobre palavras-chave, artigos, operadores ou funções da biblioteca padrão (`Texto`, `Matematica`, `Arquivo`, `Html`) para ver explicações detalhadas em português e exemplos de uso.
- ⚡ **Autocompletação Inteligente (IntelliSense)**:
  - Sugestões para palavras-chave, estruturas de controle, tipos primitivos e funções da biblioteca padrão (BibVerbo).
- 📑 **Estrutura do Documento (Outline / Símbolos)**:
  - Navegue facilmente entre funções (`Para`), estruturas de dados (`Entidade`), rotas web (`rota HTTP`) e variáveis.
- 🚀 **Comandos Rápidos no Editor**:
  - Botão **Executar** (`Play`) na barra de título do editor para rodar programas Verbo em um clique.
  - Comandos integrados no menu de contexto e na paleta de comandos (`Ctrl+Shift+P` / `Cmd+Shift+P`).

---

## ⌨️ Comandos Disponíveis

| Comando | Descrição |
| --- | --- |
| `Verbo: Executar Arquivo Atual` | Transpila, compila e executa o arquivo `.vrb` atual no terminal integrado |
| `Verbo: Compilar Arquivo Atual` | Transpila e gera o binário executável |
| `Verbo: Iniciar Servidor Web` | Inicia o servidor web do arquivo atual (`verbo servir`) |
| `Verbo: Verificar Sintaxe do Arquivo` | Executa a verificação sintática e reporta diagnósticos |

---

## 📝 Snippets (Atalhos de Código)

Digite o prefixo e pressione `Tab`:

- `principal` — Estrutura inicial com função principal
- `para_usando` — Declaração de função com parâmetros tipados
- `para` — Declaração de função simples
- `constante` — Declaração imutável (`O nome é valor.`)
- `variavel` — Declaração mutável (`Um nome está valor.`)
- `exibir` — Impressão na tela (`Exibir com (...).`)
- `se_senao` — Condicional Se / Senão
- `repita_vezes` — Repetição N vezes (`Repita 10 vezes:`)
- `repita_cada` — Iteração sobre coleções (`Repita para cada item em lista:`)
- `enquanto` — Laço condicional (`Enquanto ...:`)
- `entidade` — Definição de entidade (`A entidade Nome contendo (...)`)
- `novo` — Instanciação de entidade (`um novo Nome contendo (...)`)
- `tente` — Tratamento de erros (`Tente: ... Capture erro: ...`)
- `canal` — Declaração de canal concorrente
- `simultaneamente` — Execução em paralelo
- `servidor` — Instância do servidor web
- `rota` — Definição de rota HTTP (`GET`, `POST`, `PUT`, `DELETE`)
- `incluir` — Importação de biblioteca padrão (`Texto`, `Matematica`, `Arquivo`, `Html`)

---

## ⚙️ Configurações

Acesse `Configurações` (`Ctrl+,`) e busque por `Verbo`:

- `verbo.cliPath`: Caminho customizado para o executável `verbo`. Por padrão, procura em `./build/verbo` na raiz do projeto ou no `PATH`.
- `verbo.verificarAoSalvar`: Se ativo (`true`), roda a verificação de sintaxe automaticamente ao salvar o arquivo `.vrb`.
- `verbo.servidorHost`: Host padrão para o servidor web (padrão: `127.0.0.1`).
- `verbo.servidorPorta`: Porta padrão para o servidor web (padrão: `5000`).

---

## 🛠️ Como Desenvolver / Compilar a Extensão

Na pasta `vscode-extension/`:

```bash
# 1. Instalar dependências
npm install

# 2. Compilar a extensão
npm run build

# 3. Modo desenvolvimento contínuo (watch)
npm run watch
```

Para instalar a extensão localmente no seu VSCode:

```bash
npx @vscode/vsce package
code --install-extension verbo-0.1.0.vsix
```

---

## 📜 Licença

MIT License — Desenvolvido com ❤️ no Brasil.

