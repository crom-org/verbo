# Módulo Sistema

Interação com o sistema operacional: argumentos da linha de comando, variáveis de ambiente, execução de comandos, informações da máquina e controle do processo.

```verbo
Incluir Sistema.
```

Falhas ao executar comandos, acessar diretórios inexistentes ou índices inválidos **sinalizam** pânico.

---

## Argumentos da Linha de Comando

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `Argumentos` | `() → Lista` | Lista de argumentos (sem o nome do executável) |
| `Argumento` | `(indice: Inteiro) → Texto` | Argumento na posição `indice` (base 0); pânico se fora do intervalo |
| `NumeroDeArgumentos` | `() → Inteiro` | Quantidade de argumentos |

```verbo
Incluir Sistema.

O quantidade é NumeroDeArgumentos de Sistema com ().
Exibir "Total de argumentos: " + quantidade.

Se quantidade > 0:
    O primeiro é Argumento de Sistema com (0).
    Exibir "Primeiro argumento: " + primeiro.
.
```

---

## Variáveis de Ambiente

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `VariavelDeAmbiente` | `(nome: Texto) → Texto` | Valor da variável; `""` se não definida |
| `DefinirVariavelDeAmbiente` | `(nome: Texto, valor: Texto)` | Define a variável no processo atual |
| `RemoverVariavelDeAmbiente` | `(nome: Texto)` | Remove a variável do processo |
| `VariaveisDeAmbiente` | `() → Lista` | Todas as variáveis no formato `"NOME=valor"` |

```verbo
Incluir Sistema.

O caminho é VariavelDeAmbiente de Sistema com ("PATH").
Exibir "PATH: " + caminho.

DefinirVariavelDeAmbiente de Sistema com ("APP_MODO", "producao").
O modo é VariavelDeAmbiente de Sistema com ("APP_MODO").
Exibir "Modo: " + modo.
```

---

## Execução de Comandos

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `ExecutarComando` | `(comando: Texto) → Texto` | Executa via shell; retorna stdout sem `\n` final |
| `ExecutarComandoComArgumentos` | `(programa: Texto, args: Lista) → Texto` | Executa diretamente (sem shell); mais seguro |
| `ExecutarComandoComEntrada` | `(comando: Texto, entrada: Texto) → Texto` | Executa via shell com stdin |

> **Atenção**: `ExecutarComando` usa `/bin/sh -c` (Unix) ou `cmd /C` (Windows). Prefira `ExecutarComandoComArgumentos` quando os argumentos forem dinâmicos, pois evita injeção de shell.

```verbo
Incluir Sistema.

O data é ExecutarComando de Sistema com ("date").
Exibir "Data atual: " + data.

O arquivos é ExecutarComandoComArgumentos de Sistema com ("ls", ["-la"]).
Exibir arquivos.

O contagem é ExecutarComandoComEntrada de Sistema com ("wc -w", "Olá Mundo Verbo").
Exibir "Palavras: " + contagem.
```

---

## Diretório de Trabalho

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `DiretorioAtual` | `() → Texto` | Caminho absoluto do diretório atual |
| `MudarDiretorio` | `(caminho: Texto)` | Muda o diretório de trabalho; pânico se inválido |

```verbo
Incluir Sistema.

O dir é DiretorioAtual de Sistema com ().
Exibir "Estou em: " + dir.

MudarDiretorio de Sistema com ("/tmp").
Exibir "Mudei para: " + DiretorioAtual de Sistema com ().
```

---

## Informações do Sistema

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `NomeSistemaOperacional` | `() → Texto` | `"linux"`, `"windows"`, `"darwin"`, etc. |
| `ArquiteturaProcessador` | `() → Texto` | `"amd64"`, `"arm64"`, `"386"`, etc. |
| `NomeDoComputador` | `() → Texto` | Hostname da máquina |
| `PID` | `() → Inteiro` | ID do processo atual |
| `NumeroDeCPUs` | `() → Inteiro` | CPUs lógicas disponíveis |

```verbo
Incluir Sistema.

O so é NomeSistemaOperacional de Sistema com ().
O arq é ArquiteturaProcessador de Sistema com ().
O host é NomeDoComputador de Sistema com ().
O pid é PID de Sistema com ().
O cpus é NumeroDeCPUs de Sistema com ().

Exibir "Sistema: " + so + " (" + arq + ")".
Exibir "Máquina: " + host.
Exibir "PID: " + pid.
Exibir "CPUs: " + cpus.
```

---

## Controle do Processo

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `Sair` | `()` | Encerra o processo com código 0 (sucesso) |
| `SairCom` | `(codigo: Inteiro)` | Encerra com o código fornecido (0 = sucesso) |

```verbo
Incluir Sistema.

O ok é Verdadeiro.
Se ok:
    Exibir "Concluído com sucesso.".
    Sair de Sistema com ().
Senao:
    Exibir "Falhou.".
    SairCom de Sistema com (1).
.
```

---

## Exemplo Completo

```verbo
Incluir Sistema.

O so é NomeSistemaOperacional de Sistema com ().
O host é NomeDoComputador de Sistema com ().
O pid é PID de Sistema com ().

Exibir "Rodando no " + so + " — host: " + host + " — PID: " + pid.

O quantidade é NumeroDeArgumentos de Sistema com ().
Se quantidade > 0:
    O nome_arquivo é Argumento de Sistema com (0).
    O conteudo é ExecutarComando de Sistema com ("cat " + nome_arquivo).
    Exibir conteudo.
Senao:
    Exibir "Nenhum arquivo fornecido como argumento.".
.
```

