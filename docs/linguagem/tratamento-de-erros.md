# Tratamento de Erros

O Verbo 2.0 modela falhas com três palavras-chave: `Tente` (try), `Capture` (recover) e `Sinalize` (panic). O transpiler gera o padrão `defer`/`recover` do Go.

---

## `Sinalize`

Interrompe a execução atual e propaga um pânico:

```verbo
Sinalize com ("algo deu errado!").
Sinalize "Saldo insuficiente para a operação.".
```

Equivalente a `panic(...)` em Go. Funções da BibVerbo (Arquivo, Internet, Criptografia) também sinalizam internamente em falhas de I/O, rede ou decodificação.

---

## `Tente` / `Capture`

Envolve um bloco arriscado e captura o pânico:

```verbo
Tente:
    Exibir com ("Executando operação arriscada...").
    Sinalize com ("algo deu errado!").
    Exibir com ("Esta linha nunca será executada.").
Capture erro:
    Exibir com ("Erro capturado com sucesso!").
    Exibir com (erro).
.
```

- O identificador após `Capture` (`erro`) recebe o valor do pânico.
- Após o bloco `Capture`, o programa **continua normalmente**.
- O fechamento do bloco usa o ponto final `.`.

---

## Código Go gerado

```go
func() {
    defer func() {
        if erro := recover(); erro != nil {
            fmt.Println("Erro capturado com sucesso!")
            fmt.Println(erro)
        }
    }()
    fmt.Println("Executando operação arriscada...")
    panic("algo deu errado!")
}()
```

---

## Padrão com I/O

```verbo
Incluir Arquivo.

Tente:
    O conteudo é LerTexto de Arquivo com ("dados.txt").
    Exibir com (conteudo).
Capture erro:
    Exibir "Falha ao ler o arquivo:".
    Exibir erro.
.
```

---

## Regras

1. `Sinalize` sem `Tente` encerra o programa (pânico não recuperado).
2. `Capture` é opcional; `Tente:` sozinho ainda envolve o bloco, mas o pânico sobe.
3. Blocos `Tente` podem ser aninhados em funções, laços e rotas HTTP.
