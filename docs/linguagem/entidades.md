# Entidades (Estruturas)

Entidades são o equivalente a *structs*: agrupam campos nomeados e tipados em um único tipo. No Verbo 2.0, a declaração usa `entidade` + `contendo`, e o acesso a campos usa a preposição `de`.

---

## Declaração

```verbo
A entidade Pessoa contendo (Nome: Texto, Idade: Inteiro, Saldo: Decimal).
```

Cada campo tem a forma `Nome: Tipo`. Tipos válidos: `Texto`, `Inteiro`, `Decimal`, `Lógico` / `Logico`, `Lista`.

O compilador gera um `struct` Go com campos exportados (primeira letra maiúscula):

```go
type Pessoa struct {
    Nome  string
    Idade int
    Saldo float64
}
```

---

## Instanciação

Há duas formas equivalentes:

### 1. `novo` + `contendo`

```verbo
O cliente é um novo Pessoa contendo ("Ada Lovelace", 36, 1500.50).
```

### 2. Chamada com o nome da entidade

```verbo
Um cafe é Produto com ("Café Premium", 15).
```

Os argumentos são mapeados **na ordem dos campos** declarados.

---

## Acesso a campos

Em vez de `usuario.nome`, o Verbo usa a preposição `de`:

```verbo
O nome_cliente é Nome de cliente.
Exibir com (nome de cafe).
Exibir com (preco de cafe).
```

O transpiler capitaliza o campo (`Nome de cliente` → `cliente.Nome`).

---

## Entidades em funções

Entidades podem ser usadas como tipo de parâmetro:

```verbo
A entidade Usuario contendo (Nome: Texto, Email: Texto, Nivel: Inteiro).

Para BoasVindas usando (usuario: Usuario):
    O nome_formatado é Nome de usuario.
    Exibir "Bem-vindo(a), " + nome_formatado + "!".
.

O admin é um novo Usuario contendo ("Ada Lovelace", "ada@verbo.dev", 1).
BoasVindas com (admin).
```

---

## Exemplo completo

```verbo
A Entidade Produto contendo (nome: Texto, preco: Inteiro).

Um cafe é Produto com ("Café Premium", 15).
Um cha é Produto com ("Chá Verde", 8).

Exibir com ("Cardápio:").
Exibir com (nome de cafe).
Exibir com (preco de cafe).
Exibir com (nome de cha).
Exibir com (preco de cha).
```
