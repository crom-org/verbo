# Guia de Contribuição e Desenvolvimento

Agradecemos o seu interesse em contribuir com a **Linguagem Verbo**! Este guia descreve o fluxo de trabalho, padrões de código e diretrizes de desenvolvimento do projeto.

---

## 1. Pré-requisitos de Desenvolvimento

- **Go (Golang)**: Versão 1.22 ou superior instalada.
- **Git**: Controle de versão.
- **Make**: Automação de tarefas (opcional, mas recomendado).

---

## 2. Configurando o Ambiente Local

```bash
# 1. Clone o repositório
git clone https://github.com/crom-org/verbo.git
cd crom-verbo

# 2. Compile o CLI localmente
make build

# 3. Execute a suíte completa de testes
make test
```

---

## 3. Estrutura do Código-Fonte

| Diretório | Responsabilidade |
| :--- | :--- |
| `cmd/verbo/` | Ponto de entrada do executável CLI, subcomandos e gerenciador de pacotes. |
| `pkg/lexer/` | Tokenizador UTF-8, definições de tokens e palavras-chave. |
| `pkg/parser/` | Analisador sintático descendente recursivo e validação estrutural. |
| `pkg/ast/` | Definição das interfaces e tipos de nós da Árvore de Sintaxe Abstrata. |
| `pkg/transpiler/` | Transpilador AST → Go e gerador de código. |
| `pkg/stdlib/` | Biblioteca padrão oficial (BibVerbo: `matematica`, `texto`, `arquivo`, `csv`, `json`, `html`, `internet`, `criptografia`). |
| `examples/` | Catálogo de programas `.vrb` de demonstração e testes de usuário. |
| `docs/` | Documentação oficial do projeto. |
| `vscode-extension/` | Extensão oficial para o Visual Studio Code. |

---

## 4. Fluxo de Trabalho (Branching & Commits)

1. Crie uma branch a partir de `main` com prefixo descritivo:
   ```bash
   git checkout -b feat/suporte-modulo-banco-dados
   # ou
   git checkout -b fix/parser-se-aninhado
   ```
2. Escreva testes unitários para cobrir suas alterações.
3. Adicione um exemplo funcional na pasta `examples/` se for um novo recurso da linguagem.
4. Siga as convenções de **Conventional Commits**:
   - `feat:` Nova funcionalidade ou instrução da linguagem.
   - `fix:` Correção de bug no compilador ou stdlib.
   - `docs:` Atualizações ou adições na documentação.
   - `test:` Novos testes unitários ou de integração.
   - `refactor:` Melhorias internas sem alteração de comportamento.
5. Abra um **Pull Request** detalhando as motivações e alterações.

---

## 5. Diretrizes de Qualidade e Código

- **Go Canônico**: Todo código Go deve estar formatado segundo o `gofmt`.
- **Nomenclatura**:
  - Nomes de funções e variáveis públicas da **BibVerbo** e da **API da Linguagem** devem ser escritos em **Português**.
  - Código interno de suporte em Go pode utilizar nomes técnicos universais.
- **Cobertura de Testes**: Nenhuma alteração no Lexer, Parser, Transpiler ou Stdlib deve ser aceita sem testes automatizados (`go test ./... -v`).
- **Determinismo**: O compilador nunca deve depender de heurísticas não-determinísticas. A gramática formal é a fonte única da verdade.

---

## 6. Reportando Problemas e Dúvidas

Encontrou um erro no compilador ou tem uma sugestão de melhoria sintática? Abra uma issue no repositório oficial:
👉 [GitHub Issues — crom-verbo](https://github.com/crom-org/verbo/issues)
