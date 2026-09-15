# Seu Primeiro Programa em Verbo

Neste tutorial, criaremos e executaremos seu primeiro programa em Verbo: o tradicional **"Olá, Mundo!"**.

---

## 1. Criando o Arquivo Fonte

Crie um arquivo chamado `ola_mundo.vrb` com o seguinte conteúdo:

```verbo
// Meu primeiro programa em Verbo
O saudacao é "Olá, Mundo!".
Exibir com (saudacao).
```

Note as características da linguagem:
- O comentário de linha única inicia com `//`.
- A constante `saudacao` é definida com o artigo `O` e o verbo `é`.
- A exibição na saída padrão utiliza `Exibir com (...)`.
- Cada instrução encerra com um ponto final obrigatório (`.`).

---

## 2. Executando o Programa Diretamente

O comando `executar` transpila o código Verbo para Go em memória e o roda imediatamente:

```bash
verbo executar ola_mundo.vrb
```

Saída:
```
🚀 Executando programa Verbo...
────────────────────────────────────────
Olá, Mundo!
────────────────────────────────────────
✅ Programa finalizado com sucesso.
```

---

## 3. Compilando para um Binário Nativo

Para gerar um binário executável independente (que não precisa de nada além do sistema operacional para rodar):

```bash
verbo compilar ola_mundo.vrb
```

Isso gera:
- `ola_mundo_verbo.go`: O código Go correspondente.
- `ola_mundo`: O binário executável compilado.

Você pode rodar o binário diretamente:
```bash
./ola_mundo
```

---

## 4. Verificando a Sintaxe (Linter)

Se quiser apenas checar se seu código está sintaticamente correto sem compilar nem executar:

```bash
verbo verificar ola_mundo.vrb
```

Saída:
```
📊 Tokens encontrados: 12
🌳 Declarações na AST: 2
✅ Arquivo 'ola_mundo.vrb' está sintaticamente correto!
```

