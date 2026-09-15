# Módulo Internet (HTTP e Sockets)

Cliente HTTP, sockets TCP/UDP, listeners e resolução DNS.

```verbo
Incluir Internet.
```

O cliente HTTP padrão tem timeout de **30 segundos**. Falhas de conexão, status de download fora de 2xx e I/O de socket **sinalizam** pânico.

---

## Cliente HTTP

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `DefinirTempoLimite` | `(segundos: Inteiro)` | Timeout global do cliente HTTP |
| `Obter` / `Get` | `(url: Texto) → Texto` | HTTP GET, devolve o corpo |
| `Postar` | `(url: Texto, corpo: Texto, tipoConteudo: Texto) → Texto` | HTTP POST (`tipoConteudo` vazio → `application/json`) |
| `Post` | `(url: Texto, corpo: Texto) → Texto` | POST com JSON |
| `Put` | `(url: Texto, corpo: Texto, tipoConteudo: Texto) → Texto` | HTTP PUT |
| `Deletar` | `(url: Texto) → Texto` | HTTP DELETE |
| `Requisicao` | `(metodo: Texto, url: Texto, corpo: Texto, tipoConteudo: Texto) → Texto` | Método arbitrário (`PATCH`, etc.) |
| `Status` | `(metodo: Texto, url: Texto) → Inteiro` | Código HTTP (200, 404, 500…) |
| `Baixar` | `(url: Texto, caminhoDestino: Texto) → Texto` | Download para disco; devolve o caminho |

```verbo
Incluir Internet.

DefinirTempoLimite de Internet com (10).

O json é Obter de Internet com ("https://httpbin.org/get").
Exibir json.

O criado é Postar de Internet com ("https://httpbin.org/post", "{\"nome\":\"Verbo\"}", "application/json").
O codigo é Status de Internet com ("GET", "https://httpbin.org/status/200").
Exibir codigo.
```

---

## Sockets TCP (cliente)

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `ConectarTCP` | `(endereco: Texto) → Conexao` | Cliente TCP (`host:porta`), timeout 10s |
| `EnviarTCP` | `(c: Conexao, mensagem: Texto) → Inteiro` | Envia bytes |
| `ReceberTCP` | `(c: Conexao, tamanhoMaximo: Inteiro) → Texto` | Lê até N bytes (padrão 4096 se N ≤ 0) |
| `LerLinhaTCP` | `(c: Conexao) → Texto` | Lê até `\n` |
| `FecharTCP` | `(c: Conexao)` | Fecha a conexão |

Métodos na instância `Conexao`:

| Método | Descrição |
| :--- | :--- |
| `Enviar` / `EnviarLinha` | Envia texto (com `\n` no segundo) |
| `Receber` / `LerLinha` | Lê bytes / linha |
| `Fechar` | Encerra |
| `EnderecoRemoto` / `EnderecoLocal` | `host:porta` |
| `EstaFechada` | `Lógico` |

---

## Servidor TCP (listener)

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `OuvirTCP` | `(endereco: Texto) → ServidorTCP` | Escuta (`:9000` ou `127.0.0.1:9000`) |
| `AceitarTCP` | `(s: ServidorTCP) → Conexao` | Bloqueia até um cliente |
| `ObterEnderecoTCP` | `(s: ServidorTCP) → Texto` | Endereço em escuta |
| `FecharServidorTCP` | `(s: ServidorTCP)` | Encerra o listener |

Métodos em `ServidorTCP`: `Aceitar`, `Fechar`, `Endereco` / `ObterEndereco`.

```verbo
Incluir Internet.

O ouvinte é OuvirTCP de Internet com ("127.0.0.1:9876").
O endereco_tcp é ObterEndereco de ouvinte com ().
Exibir "Servidor TCP iniciado no endereço: " + endereco_tcp.

O cliente é ConectarTCP de Internet com ("127.0.0.1:9876").
O atendente é Aceitar de ouvinte com ().

EnviarLinha de cliente com ("Olá do Verbo!").
O recebido é LerLinha de atendente com ().
Exibir "Servidor TCP recebeu: " + recebido.

EnviarLinha de atendente com ("Mensagem recebida com sucesso!").
O resposta_cliente é LerLinha de cliente com ().
Exibir "Cliente TCP recebeu confirmação: " + resposta_cliente.

Fechar de cliente com ().
Fechar de atendente com ().
Fechar de ouvinte com ().
```

---

## UDP

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `OuvirUDP` | `(endereco: Texto) → ConexaoUDP` | Socket UDP em escuta |
| `EnviarUDP` | `(endereco: Texto, mensagem: Texto) → Inteiro` | Datagrama para `host:porta` |
| `ReceberUDP` | `(u: ConexaoUDP, tamanhoMaximo: Inteiro) → Texto` | Lê até N bytes (padrão 2048) |
| `FecharUDP` | `(u: ConexaoUDP)` | Fecha o socket |

Métodos em `ConexaoUDP`: `Receber`, `Fechar`, `EnderecoLocal`.

---

## DNS e utilitários

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `ResolverHost` | `(host: Texto) → Lista` | IPs do domínio |
| `PortaAberta` | `(host: Texto, porta: Inteiro) → Lógico` | Probe TCP (timeout 2s) |
| `ObterIPLocal` | `() → Texto` | Primeiro IPv4 não-loopback; senão `127.0.0.1` |

```verbo
Incluir Internet.

O ip_local é ObterIPLocal de Internet com ().
O ips é ResolverHost de Internet com ("localhost").
O aberta é PortaAberta de Internet com ("127.0.0.1", 8080).

Exibir ip_local.
Exibir ips.
Exibir aberta.
```
