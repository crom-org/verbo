# Módulo Csv

Leitura e escrita de dados no formato CSV (_Comma-Separated Values_). Falhas de parsing ou I/O **sinalizam** pânico.

```verbo
Incluir Csv.
```

---

## Funções

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `Ler` | `(texto: Texto) → Lista` | Parseia uma string CSV e retorna uma lista de listas (uma por linha) |
| `LerArquivo` | `(caminho: Texto) → Lista` | Lê um arquivo CSV do disco e retorna uma lista de listas |
| `Escrever` | `(linhas: Lista) → Texto` | Converte uma lista de listas em string CSV formatada |
| `EscreverArquivo` | `(caminho: Texto, linhas: Lista)` | Grava uma lista de listas como arquivo CSV no disco |
| `Formatar` | `(linhas: Lista) → Texto` | Alias de `Escrever` — converte lista de listas em string CSV |

---

## Exemplos

### Parsear CSV de uma string

```verbo
Incluir Csv.

O texto é "nome,idade\nAna,30\nBeto,25".
O linhas é Ler de Csv com (texto).
Exibir linhas.
```

### Ler arquivo CSV do disco

```verbo
Incluir Csv.

O dados é LerArquivo de Csv com ("pessoas.csv").
Exibir dados.
```

### Escrever CSV para uma string

```verbo
Incluir Csv.

O cabecalho é ["nome", "idade"].
O linha1 é ["Ana", "30"].
O linha2 é ["Beto", "25"].
O linhas é [cabecalho, linha1, linha2].

O saida é Escrever de Csv com (linhas).
Exibir saida.
```

### Gravar arquivo CSV no disco

```verbo
Incluir Csv.

O linhas é [["produto", "preco"], ["Café", "12.50"], ["Chá", "8.00"]].
EscreverArquivo de Csv com ("produtos.csv", linhas).
```

### Com tratamento de erro

```verbo
Incluir Csv.

Tente:
    O dados é LerArquivo de Csv com ("relatorio.csv").
    Exibir dados.
Capture erro:
    Exibir "Falha ao ler CSV.".
    Exibir erro.
.
```

Mensagens de pânico geradas pelo módulo:

- `Erro ao ler CSV: ...`
- `Erro ao escrever linha CSV: ...`
- `Erro ao finalizar CSV: ...`

---

## Equivalente Go

| Verbo | Go (`encoding/csv`) |
| :--- | :--- |
| `Ler` | `csv.NewReader` + `ReadAll` |
| `LerArquivo` | `os.ReadFile` + `csv.NewReader` + `ReadAll` |
| `Escrever` / `Formatar` | `csv.NewWriter` + `WriteAll` |
| `EscreverArquivo` | `csv.NewWriter` + `WriteAll` + `os.WriteFile` |

