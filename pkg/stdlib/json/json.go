// Package json implementa funções de codificação e decodificação JSON
// para a linguagem de programação Verbo.
// Parte da BibVerbo (Biblioteca Padrão do Verbo).
//
// Uso em Verbo:
//
//	Incluir Json.
//	O texto é Codificar de Json com (dados).
//	O obj é Decodificar de Json com (texto).
package json

import (
	"encoding/json"
	"fmt"
)

// Codificar converte um valor (map, lista, string, número, bool) em uma string JSON.
//
// Exemplo em Verbo:
//
//	O texto é Codificar de Json com (dados).
func Codificar(valor interface{}) string {
	b, err := json.Marshal(valor)
	if err != nil {
		panic("Erro ao codificar JSON: " + err.Error())
	}
	return string(b)
}

// Decodificar parseia uma string JSON e retorna um valor genérico (map/lista/escalar).
// Lança pânico se a string não for JSON válido.
//
// Exemplo em Verbo:
//
//	O obj é Decodificar de Json com (texto).
func Decodificar(texto string) interface{} {
	var resultado interface{}
	if err := json.Unmarshal([]byte(texto), &resultado); err != nil {
		panic("Erro ao decodificar JSON: " + err.Error())
	}
	return resultado
}

// ObterCampo extrai o valor de um campo de um mapa JSON decodificado.
// Retorna nil se o campo não existir.
//
// Exemplo em Verbo:
//
//	O nome é ObterCampo de Json com (obj, "nome").
func ObterCampo(obj interface{}, campo string) interface{} {
	if m, ok := obj.(map[string]interface{}); ok {
		return m[campo]
	}
	panic(fmt.Sprintf("Objeto não é um mapa JSON válido (tipo %T)", obj))
}

// ListaChaves retorna as chaves de um mapa JSON decodificado como lista de strings.
//
// Exemplo em Verbo:
//
//	O chaves é ListaChaves de Json com (obj).
func ListaChaves(obj interface{}) []interface{} {
	if m, ok := obj.(map[string]interface{}); ok {
		res := make([]interface{}, 0, len(m))
		for k := range m {
			res = append(res, k)
		}
		return res
	}
	panic(fmt.Sprintf("Objeto não é um mapa JSON válido (tipo %T)", obj))
}

// Tamanho retorna o número de elementos em um mapa ou lista JSON.
//
// Exemplo em Verbo:
//
//	O n é Tamanho de Json com (obj).
func Tamanho(obj interface{}) int {
	switch v := obj.(type) {
	case map[string]interface{}:
		return len(v)
	case []interface{}:
		return len(v)
	case string:
		return len(v)
	default:
		return 0
	}
}