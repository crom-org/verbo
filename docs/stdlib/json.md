# Módulo Json

Codificação e decodificação de dados no formato JSON. Falhas de parsing **sinalizam** pânico.

```verbo
Incluir Json.
```

---

## Funções

| Função | Assinatura | Descrição |
| :--- | :--- | :--- |
| `Codificar` | `(valor: qualquer) → Texto` | Converte um valor (mapa, lista, texto, número, booleano) em string JSON |
| `Decodificar` | `(texto: Texto) → qualquer` | Parseia uma string JSON e retorna o valor correspondente |
| `ObterCampo` | `(obj: qualquer, campo: Texto) → qualquer` | Extrai o valor de um campo de um mapa JSON decodificado |
| `ListaChaves` | `(obj: qualquer) → Lista` | Retorna as chaves de um mapa JSON como lista de strings |
| `Tamanho` | `(obj: qualquer) → Numero` | Retorna o número de elementos de um mapa, lista ou string JSON |

---

## Exemplos

### Codificar valor como JSON

```verbo
Incluir Json.

O dados é ["nome", "Ana", "idade", 30].
O texto é Codificar de Json com (dados).
Exibir texto.
```

### Decodificar string JSON

```verbo
Incluir Json.

O texto é "{\"nome\":\"Beto\",\"idade\":25}".
O obj é Decodificar de Json com (texto).
Exibir obj.
```

### Acessar campo de um objeto JSON

```verbo
Incluir Json.

O texto é "{\"linguagem\":\"Verbo\",\"versao\":\"1.0\"}".
O obj é Decodificar de Json com (texto).
O linguagem é ObterCampo de Json com (obj, "linguagem").
Exibir linguagem.
```

### Listar chaves de um objeto JSON

```verbo
Incluir Json.

O texto é "{\"a\":1,\"b\":2,\"c\":3}".
O obj é Decodificar de Json com (texto).
O chaves é ListaChaves de Json com (obj).
Exibir chaves.
```

### Obter tamanho de um objeto ou lista

```verbo
Incluir Json.

O texto é "[1,2,3,4,5]".
O lista é Decodificar de Json com (texto).
O n é Tamanho de Json com (lista).
Exibir n.
```

### Com tratamento de erro

```verbo
Incluir Json.

Tente:
    O obj é Decodificar de Json com ("JSON inválido!!!").
    Exibir obj.
Capture erro:
    Exibir "Falha ao decodificar JSON.".
    Exibir erro.
.
```

Mensagens de pânico geradas pelo módulo:

- `Erro ao codificar JSON: ...`
- `Erro ao decodificar JSON: ...`
- `Objeto não é um mapa JSON válido (tipo ...)`

---

## Equivalente Go

| Verbo | Go (`encoding/json`) |
| :--- | :--- |
| `Codificar` | `json.Marshal` |
| `Decodificar` | `json.Unmarshal` |
| `ObterCampo` | Acesso a `map[string]interface{}` |
| `ListaChaves` | Iteração de chaves em `map[string]interface{}` |
| `Tamanho` | `len(...)` |

