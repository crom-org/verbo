// Package csv implementa funções de leitura e escrita de arquivos CSV
// para a linguagem de programação Verbo.
// Parte da BibVerbo (Biblioteca Padrão do Verbo).
//
// Uso em Verbo:
//
//	Incluir Csv.
//	O linhas é Ler de Csv com (texto).
//	O saida é Escrever de Csv com (linhas).
package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/juanxto/crom-verbo/pkg/stdlib/arquivo"
)

// Ler parseia uma string CSV e retorna uma lista de listas (uma lista por linha).
// Cada linha é representada como uma lista de strings.
// Lança pânico se o CSV for inválido.
//
// Exemplo em Verbo:
//
//	O linhas é Ler de Csv com (texto).
func Ler(texto string) []interface{} {
	leitor := csv.NewReader(strings.NewReader(texto))
	var resultado []interface{}

	for {
		registro, err := leitor.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("Erro ao ler CSV: " + err.Error())
		}
		linha := make([]interface{}, len(registro))
		for i, campo := range registro {
			linha[i] = campo
		}
		resultado = append(resultado, linha)
	}

	return resultado
}

// LerArquivo lê um arquivo CSV do disco e retorna uma lista de listas.
// Cada linha é representada como uma lista de strings.
// Lança pânico se o arquivo não puder ser lido.
//
// Exemplo em Verbo:
//
//	O dados é LerArquivo de Csv com ("dados.csv").
func LerArquivo(caminho string) []interface{} {
	conteudo := lerTexto(caminho)
	return Ler(conteudo)
}

// Escrever converte uma lista de listas em uma string CSV formatada.
// Cada elemento da lista externa representa uma linha.
// Cada linha deve ser uma lista de strings.
//
// Exemplo em Verbo:
//
//	O saida é Escrever de Csv com (linhas).
func Escrever(linhas []interface{}) string {
	var sb strings.Builder
	escritor := csv.NewWriter(&sb)

	for _, linhaRaw := range linhas {
		linha := toStringSlice(linhaRaw)
		if err := escritor.Write(linha); err != nil {
			panic("Erro ao escrever linha CSV: " + err.Error())
		}
	}

	escritor.Flush()
	if err := escritor.Error(); err != nil {
		panic("Erro ao finalizar CSV: " + err.Error())
	}
	return sb.String()
}

// EscreverArquivo grava uma lista de listas em um arquivo CSV no disco.
// Cada elemento da lista externa representa uma linha.
// Cada linha deve ser uma lista de strings.
//
// Exemplo em Verbo:
//
//	EscreverArquivo de Csv com ("saida.csv", linhas).
func EscreverArquivo(caminho string, linhas []interface{}) {
	conteudo := Escrever(linhas)
	arquivo.EscreverTexto(caminho, conteudo)
}

// Formatar converte uma lista de listas em string CSV (alias para Escrever).
//
// Exemplo em Verbo:
//
//	O texto é Formatar de Csv com (linhas).
func Formatar(linhas []interface{}) string {
	return Escrever(linhas)
}

// toStringSlice converte um interface{} para []string para escrita CSV.
func toStringSlice(v interface{}) []string {
	switch val := v.(type) {
	case []string:
		return val
	case []interface{}:
		result := make([]string, len(val))
		for i, item := range val {
			result[i] = fmt.Sprintf("%v", item)
		}
		return result
	case string:
		return []string{val}
	default:
		return []string{fmt.Sprintf("%v", val)}
	}
}

func lerTexto(caminho string) string {
	return arquivo.LerTexto(caminho)
}

func escreverTexto(caminho, conteudo string) {
	arquivo.EscreverTexto(caminho, conteudo)
}