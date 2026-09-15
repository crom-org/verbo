# Módulo Arquivo

Leitura e escrita de arquivos de texto (UTF-8). Falhas de I/O **sinalizam** pânico.

```verbo
Incluir Arquivo.
```

---

## Funções

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `LerTexto` | `(caminho: Texto) → Texto` | Lê o arquivo inteiro |
| `EscreverTexto` | `(caminho: Texto, conteudo: Texto)` | Grava o conteúdo (modo `0644`), sobrescrevendo |

---

## Exemplos

### Escrever e ler

```verbo
Incluir Arquivo.

EscreverTexto de Arquivo com ("notas.txt", "Primeira linha\nSegunda linha").
O conteudo é LerTexto de Arquivo com ("notas.txt").
Exibir conteudo.
```

### Com tratamento de erro

```verbo
Incluir Arquivo.

Tente:
    O dados é LerTexto de Arquivo com ("config.txt").
    Exibir dados.
Capture erro:
    Exibir "Arquivo ausente ou ilegível.".
    Exibir erro.
.
```

Mensagens de pânico geradas pelo módulo:

- `Erro ao ler arquivo <caminho>: ...`
- `Erro ao escrever arquivo <caminho>: ...`

---

## Equivalente Go

| Verbo | Go (`os`) |
| :--- | :--- |
| `LerTexto` | `os.ReadFile` → `string` |
| `EscreverTexto` | `os.WriteFile(..., 0644)` |
