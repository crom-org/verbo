# Concorrência e Canais

O Verbo 2.0 expõe goroutines e canais do Go com vocabulário em português: `Simultaneamente`, `Canal`, `Enviar` e `Receber`.

> **Nota:** a palavra-chave `Aguarde` está reservada no lexer, mas a sincronização de `Simultaneamente` já é automática (`WaitGroup.Wait()`). Não é necessário escrever `Aguarde` no código atual.

---

## `Simultaneamente`

Cada declaração no bloco roda em uma goroutine. O transpiler espera todas terminarem antes de seguir:

```verbo
Exibir com ("Iniciando tarefas paralelas...").

Simultaneamente:
    Exibir com ("Tarefa 1: Processando dados.").
    Exibir com ("Tarefa 2: Enviando email.").
    Exibir com ("Tarefa 3: Gerando relatório.").
.

Exibir com ("Todas as tarefas concluídas!").
```

Código Go gerado (simplificado):

```go
{
    var wg sync.WaitGroup
    wg.Add(3)
    go func() { defer wg.Done(); fmt.Println("Tarefa 1: ...") }()
    go func() { defer wg.Done(); fmt.Println("Tarefa 2: ...") }()
    go func() { defer wg.Done(); fmt.Println("Tarefa 3: ...") }()
    wg.Wait()
}
```

---

## Canais

### Criação

```verbo
Um resultados é um Canal de Inteiros.
```

Tipos aceitos após `Canal de`: `Inteiros` / `Inteiro` → `chan int`, e equivalentes para os demais primitivos. Transpila para `make(chan int)`.

### Enviar

```verbo
Enviar resultado para canal_envio.
```

Equivalente a `canal_envio <- resultado`.

### Receber

```verbo
O res1 é Receber de resultados.
```

Equivalente a `res1 := <-resultados`. A recepção **bloqueia** até um valor chegar.

---

## Exemplo completo (workers)

```verbo
Para ProcessarPacote usando (pacote: Inteiro, canal_envio: Canal_Inteiros):
    Exibir "Processando pacote...".
    O resultado é pacote * 10.
    Enviar resultado para canal_envio.
.

Um resultados é um Canal de Inteiros.

Simultaneamente:
    ProcessarPacote com (1, resultados).
.
Simultaneamente:
    ProcessarPacote com (2, resultados).
.
Simultaneamente:
    ProcessarPacote com (3, resultados).
.

O res1 é Receber de resultados.
O res2 é Receber de resultados.
O res3 é Receber de resultados.

Exibir "Todos os pacotes processados!".
Exibir res1.
Exibir res2.
Exibir res3.
```

Cada `Simultaneamente` dispara um worker. Os três `Receber` coletam os resultados na ordem de chegada.

---

## Quando usar o quê

| Construto | Uso |
| :--- | :--- |
| `Simultaneamente` | Paralelizar tarefas independentes e esperar o conjunto |
| `Canal` + `Enviar` / `Receber` | Comunicar valores entre goroutines sem memória compartilhada |
