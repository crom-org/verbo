# Módulo Matematica

Funções numéricas sobre `Decimal` (`float64`). Todas recebem e devolvem ponto flutuante.

```verbo
Incluir Matematica.
```

---

## Funções

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `Absoluto` | `(x: Decimal) → Decimal` | Valor absoluto |
| `Teto` | `(x: Decimal) → Decimal` | Arredonda para cima (`ceil`) |
| `Piso` | `(x: Decimal) → Decimal` | Arredonda para baixo (`floor`) |
| `Maximo` | `(a: Decimal, b: Decimal) → Decimal` | Maior entre dois valores |
| `Minimo` | `(a: Decimal, b: Decimal) → Decimal` | Menor entre dois valores |
| `Potencia` | `(base: Decimal, exp: Decimal) → Decimal` | `base ^ exp` |
| `Raiz` | `(x: Decimal) → Decimal` | Raiz quadrada |

---

## Exemplos

```verbo
Incluir Matematica.

O abs é Absoluto de Matematica com (-42.5).
O cima é Teto de Matematica com (3.14).
O baixo é Piso de Matematica com (3.14).
O maior é Maximo de Matematica com (10.0, 7.5).
O menor é Minimo de Matematica com (10.0, 7.5).
O cubo é Potencia de Matematica com (2.0, 8.0).
O hipotenusa é Raiz de Matematica com (25.0).

Exibir abs.
Exibir cubo.
Exibir hipotenusa.
```

---

## Equivalente Go

| Verbo | Go (`math`) |
| :--- | :--- |
| `Absoluto` | `math.Abs` |
| `Teto` | `math.Ceil` |
| `Piso` | `math.Floor` |
| `Maximo` / `Minimo` | `math.Max` / `math.Min` |
| `Potencia` | `math.Pow` |
| `Raiz` | `math.Sqrt` |
