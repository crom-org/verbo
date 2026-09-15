# Galeria de Exemplos Práticos

Abaixo estão reunidos exemplos completos em código Verbo cobrindo desde fundamentos até concorrência, APIs e segurança.

---

## 1. Olá Mundo e Variáveis (`ola_mundo.vrb`)

```verbo
// Constante imutável com artigo definido
A mensagem é "Olá, Mundo! Bem-vindo ao Verbo.".
Exibir com (mensagem).

// Variável mutável com artigo indefinido
Um contador está 0.
contador está contador + 1.
Exibir "Contador atual: " + contador.
```

---

## 2. Funções e Tipagem (`calculadora.vrb`)

```verbo
Para Somar usando (x: Inteiro, y: Inteiro) :
    Retorne x + y.
.

Para Multiplicar usando (x: Inteiro, y: Inteiro) :
    Retorne x * y.
.

O resultado_soma é Somar com (15, 25).
O resultado_mult é Multiplicar com (6, 7).

Exibir "15 + 25 = " + resultado_soma.
Exibir "6 * 7 = " + resultado_mult.
```

---

## 3. Entidades e Modelos de Dados (`entidade.vrb`)

```verbo
A Entidade Produto contendo (nome: Texto, preco: Inteiro, ativo: Lógico).

// Instanciação
O cafe é um novo Produto contendo ("Café Premium", 15, Verdadeiro).
O cha é Produto com ("Chá Verde", 8, Verdadeiro).

// Acesso a campos com a preposição "de"
Exibir "Produto: " + nome de cafe + " | Preço: R$ " + preco de cafe.
Exibir "Produto: " + nome de cha + " | Preço: R$ " + preco de cha.
```

---

## 4. Concorrência e Canais (`canais.vrb`)

```verbo
Para ProcessarPacote usando (pacote: Inteiro, canal_envio: Canal_Inteiros) :
    Exibir "Processando pacote número " + pacote + "...".
    O resultado é pacote * 100.
    Enviar resultado para canal_envio.
.

Para principal :
    Um resultados é um Canal de Inteiros.

    // Disparando workers simultâneos
    Simultaneamente :
        ProcessarPacote com (1, resultados).
    .
    Simultaneamente :
        ProcessarPacote com (2, resultados).
    .
    Simultaneamente :
        ProcessarPacote com (3, resultados).
    .

    // Coleta dos resultados
    O r1 é Receber de resultados.
    O r2 é Receber de resultados.
    O r3 é Receber de resultados.

    Exibir "Resultados recebidos:".
    Exibir r1.
    Exibir r2.
    Exibir r3.
.

principal com ().
```

---

## 5. Tratamento de Erros e Pânico (`erros.vrb`)

```verbo
Exibir "Iniciando processamento com proteção de erro...".

Tente :
    Exibir "Realizando operação arriscada...".
    Sinalize "Falha de comunicação com o banco de dados.".
    Exibir "Esta linha não será executada.".
Capture erro :
    Exibir "Exceção interceptada com sucesso!".
    Exibir "Detalhes do erro: " + erro.
.

Exibir "Fluxo principal continuou normalmente.".
```

---

## 6. Servidor Web com Rotas HTTP (`web_index.vrb`)

```verbo
Incluir Html.

Um app está Servidor com (local, 8099).

app rota GET em "/":
    O titulo é CriarElemento de Html com ("h1", "Servidor Verbo 3.0").
    O link é CriarLink de Html com ("/saude", "Verificar Saúde da API").
    O pagina é CriarPagina de Html com ("Página Inicial", titulo + link).
    Exibir pagina.
.

app rota GET em "/saude":
    Exibir "{\"status\":\"operacional\",\"versao\":\"3.0\"}".
.

app iniciar.
```

---

## 7. Cliente HTTP e Sockets TCP (`internet_exemplo.vrb`)

```verbo
Incluir Internet.

// 1. Resolução DNS e IP Local
O ip_local é ObterIPLocal de Internet com ().
Exibir "IP Local: " + ip_local.

// 2. Servidor e Cliente TCP
O ouvinte é OuvirTCP de Internet com ("127.0.0.1:9500").
O cliente é ConectarTCP de Internet com ("127.0.0.1:9500").
O atendente é Aceitar de ouvinte com ().

EnviarLinha de cliente com ("Olá via Socket TCP!").
O mensagem_recebida é LerLinha de atendente com ().
Exibir "Servidor recebeu: " + mensagem_recebida.

Fechar de cliente com ().
Fechar de atendente com ().
Fechar de ouvinte com ().
```

---

## 8. Criptografia AES, Hashes e UUIDs (`criptografia_exemplo.vrb`)

```verbo
Incluir Criptografia.

// Hashing seguro
O hash_senha é Sha256 de Criptografia com ("minha-senha-forte").
Exibir "Hash SHA-256: " + hash_senha.

// Geração de UUID e Tokens
O id_sessao é GerarUUID de Criptografia com ().
O token_auth é GerarToken de Criptografia com (32).
Exibir "UUID v4: " + id_sessao.
Exibir "Token Auth: " + token_auth.

// Criptografia Autenticada AES-256-GCM
O chave_mestra é "segredo-super-protegido-2026".
O cifrado é CifrarAES de Criptografia com (chave_mestra, "Dados Ultraconfidenciais").
Exibir "Cifrado (Base64): " + cifrado.

O decifrado é DecifrarAES de Criptografia com (chave_mestra, cifrado).
Exibir "Decifrado: " + decifrado.
```
