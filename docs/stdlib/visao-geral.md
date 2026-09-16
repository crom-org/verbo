# Visão Geral da BibVerbo

A **BibVerbo** é a biblioteca padrão oficial. Os módulos vivem em `pkg/stdlib/` e são importados com `Incluir`.

---

## Importação

```verbo
Incluir Matematica.
Incluir Texto.
Incluir Arquivo.
Incluir Csv.
Incluir Json.
Incluir Html.
Incluir Internet.
Incluir Criptografia.
Incluir Sistema.
```

O nome do pacote é case-insensitive (`Incluir matematica.` também funciona). O transpiler gera o import Go correspondente:

```go
import "github.com/juanxto/crom-verbo/pkg/stdlib/matematica"
```

Pacotes desconhecidos caem no fallback `import "nome"` (útil para módulos externos no futuro).

---

## Chamada de funções

O padrão é `Funcao de Pacote com (args)`:

```verbo
O cubo é Potencia de Matematica com (2.0, 3.0).
O hash é Sha256 de Criptografia com ("senha").
O pagina é CriarPagina de Html com ("Título", "<h1>Olá</h1>").
```

Funções sem argumentos usam `com ()`:

```verbo
O id é GerarUUID de Criptografia com ().
O ip é ObterIPLocal de Internet com ().
```

Métodos em objetos retornados (sockets, listeners) usam o mesmo padrão com a instância:

```verbo
EnviarLinha de cliente com ("Olá servidor!").
Fechar de cliente com ().
```

---

## Módulos disponíveis

| Módulo | Arquivo | Responsabilidade |
| :--- | :--- | :--- |
| [Matematica](matematica.md) | `pkg/stdlib/matematica/` | Absoluto, potência, raiz, teto, piso, min/máx |
| [Texto](texto.md) | `pkg/stdlib/texto/` | Tamanho, maiúsculas, contém, dividir, substituir |
| [Arquivo](arquivo.md) | `pkg/stdlib/arquivo/` | Ler e escrever arquivos de texto |
| [Csv](csv.md) | `pkg/stdlib/csv/` | Ler e escrever dados no formato CSV |
| [Json](json.md) | `pkg/stdlib/json/` | Codificação e decodificação de JSON |
| [Html](html.md) | `pkg/stdlib/html/` | Páginas, tags, tabelas, links e imagens |
| [Internet](internet.md) | `pkg/stdlib/internet/` | HTTP, sockets TCP/UDP, DNS |
| [Criptografia](criptografia.md) | `pkg/stdlib/criptografia/` | Hashes, Base64/Base32/Hex, AES-GCM, UUID, aleatório |
| [Sistema](sistema.md) | `pkg/stdlib/sistema/` | Argumentos, variáveis de ambiente, execução de comandos, informações do SO |

---

## Erros

Funções de I/O, rede e decodificação **sinalizam** (`panic`) em falha. Envolva com `Tente` / `Capture`:

```verbo
Incluir Arquivo.

Tente:
    O dados é LerTexto de Arquivo com ("config.txt").
Capture erro:
    Exibir "Não foi possível ler o arquivo.".
.
```

---

## Exemplo combinando módulos

```verbo
Incluir Matematica.
Incluir Texto.
Incluir Criptografia.

O nome é Maiusculas de Texto com ("ada").
O potencia é Potencia de Matematica com (2.0, 8.0).
O token é GerarUUID de Criptografia com ().

Exibir nome.
Exibir potencia.
Exibir token.
```
