# Módulo Criptografia e Aleatório

Hashes, HMAC, Base64/Hex, AES-256-GCM, cifras didáticas e geração aleatória (segura e estatística).

```verbo
Incluir Criptografia.
```

Decodificação inválida e falha de autenticação AES **sinalizam** pânico.

---

## Hashes

| Função | Assinatura | Notas |
| :--- | :--- | :--- |
| `Sha256` | `(texto: Texto) → Texto` | Hex minúsculo (recomendado) |
| `Sha512` | `(texto: Texto) → Texto` | Hex minúsculo |
| `Sha1` | `(texto: Texto) → Texto` | Legado |
| `Md5` | `(texto: Texto) → Texto` | Legado |
| `HmacSha256` | `(chave: Texto, texto: Texto) → Texto` | HMAC-SHA-256 em hex |
| `CompararHash` | `(a: Texto, b: Texto) → Lógico` | Comparação em tempo constante |

```verbo
Incluir Criptografia.

O mensagem é "Linguagem Verbo".
O hash256 é Sha256 de Criptografia com (mensagem).
O token_hmac é HmacSha256 de Criptografia com ("minha_chave_secreta", mensagem).
O valido é CompararHash de Criptografia com (hash256, hash256).
```

`Md5` e `Sha1` existem para interoperabilidade; não use em senhas novas.

---

## Codificação

| Função | Descrição |
| :--- | :--- |
| `Base64Codificar` / `Base64Decodificar` | Base64 padrão |
| `Base64UrlCodificar` / `Base64UrlDecodificar` | Base64 URL-safe (RFC 4648) |
| `HexCodificar` / `HexDecodificar` | Hexadecimal |

```verbo
O b64 é Base64Codificar de Criptografia com ("Olá Mundo").
O original é Base64Decodificar de Criptografia com (b64).
```

---

## Cifras

### AES-256-GCM (produção)

`CifrarAES(chave, texto)` / `DecifrarAES(chave, dadosBase64)`:

- Chave de 32 bytes é usada direto; qualquer outro tamanho vira SHA-256 da chave.
- Nonce aleatório prefixado; saída em Base64.
- Chave errada ou dados adulterados → pânico.

```verbo
O chave é "chave-mestra-2026".
O texto_cifrado é CifrarAES de Criptografia com (chave, "Informação confidencial").
O texto_decifrado é DecifrarAES de Criptografia com (chave, texto_cifrado).
```

### Didáticas

| Função | Descrição |
| :--- | :--- |
| `CifraCesar(texto, deslocamento)` | Desloca A–Z / a–z; demais caracteres intactos |
| `Rot13(texto)` | César com deslocamento 13 |

Não use César/ROT-13 para dados reais.

---

## Aleatório criptograficamente seguro (`crypto/rand`)

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `GerarUUID` | `() → Texto` | UUID v4 (RFC 4122) |
| `GerarToken` | `(tamanho: Inteiro) → Texto` | Token hex com N caracteres |
| `GerarTextoAleatorio` | `(tamanho: Inteiro) → Texto` | Senha `[a-zA-Z0-9]` |
| `GerarBytesHex` | `(quantidade: Inteiro) → Texto` | N bytes em hex |
| `GerarNumeroSeguro` | `(maximo: Inteiro) → Inteiro` | Inteiro em `[0, maximo)` |

```verbo
O id é GerarUUID de Criptografia com ().
O token é GerarToken de Criptografia com (32).
O senha é GerarTextoAleatorio de Criptografia com (16).
```

---

## Aleatório estatístico (jogos / sorteios)

Não use para tokens, senhas ou chaves.

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `NumeroAleatorio` | `(minimo: Inteiro, maximo: Inteiro) → Inteiro` | Intervalo inclusivo |
| `DecimalAleatorio` | `() → Decimal` | `[0.0, 1.0)` |
| `BooleanoAleatorio` | `() → Lógico` | 50% / 50% |
| `EscolherAleatorio` | `(lista: Lista) → qualquer` | Pânico se a lista estiver vazia |
| `Embaralhar` | `(lista: Lista) → Lista` | Cópia Fisher–Yates; original intacto |

```verbo
O dado é NumeroAleatorio de Criptografia com (1, 6).
O cara_ou_coroa é BooleanoAleatorio de Criptografia com ().
O sorteado é EscolherAleatorio de Criptografia com (["Ana", "Bruno", "Caio"]).
```
