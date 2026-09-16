# 🏗️ Arquitetura do Compilador Verbo


## Visão Geral

O compilador Verbo é um **transpilador determinístico** que converte código-fonte `.vrb` em código Go nativo de alto desempenho.

```mermaid
graph TD
    A["Código Fonte .vrb"] --> B["Lexer (pkg/lexer)"]
    B --> C["Stream de Tokens UTF-8"]
    C --> D["Parser (pkg/parser)"]
    D --> E["AST Tipada (pkg/ast)"]
    E --> F["Transpiler (pkg/transpiler)"]
    F --> G["Código Fonte Go Nativo"]
    G --> H["Toolchain Go (go build / go run)"]
    H --> I["Binário de Máquina Executável"]
```

Para detalhes dos nós da AST, camadas do compilador e arquitetura da BibVerbo, veja [Arquitetura Completa](referencia/arquitetura.md).
