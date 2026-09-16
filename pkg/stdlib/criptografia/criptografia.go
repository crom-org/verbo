// Package criptografia implementa funções essenciais de segurança, criptografia simétrica,
// hashing, codificação (Base64, Hex) e geração aleatória (segura e pseudoaleatória)
// para a linguagem de programação Verbo.
// Parte da BibVerbo (Biblioteca Padrão do Verbo).
//
// Uso em Verbo:
//
//	Incluir Criptografia.
//	O hash é Sha256 de Criptografia com ("minha senha").
//	O token é GerarUUID de Criptografia com ().
package criptografia

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

// -----------------------------------------------------------------------------
// Funções de Hashing (Resumo Criptográfico)
// -----------------------------------------------------------------------------

// Sha256 calcula o hash SHA-256 do texto e o retorna como string hexadecimal em minúsculas.
//
// Exemplo em Verbo:
//
//	O hash é Sha256 de Criptografia com ("Verbo").
func Sha256(texto string) string {
	h := sha256.Sum256([]byte(texto))
	return hex.EncodeToString(h[:])
}

// Sha512 calcula o hash SHA-512 do texto e o retorna como string hexadecimal em minúsculas.
//
// Exemplo em Verbo:
//
//	O hash é Sha512 de Criptografia com ("Texto longo").
func Sha512(texto string) string {
	h := sha512.Sum512([]byte(texto))
	return hex.EncodeToString(h[:])
}

// Sha1 calcula o hash legado SHA-1 do texto e o retorna em formato hexadecimal.
func Sha1(texto string) string {
	h := sha1.Sum([]byte(texto))
	return hex.EncodeToString(h[:])
}

// Md5 calcula o hash legado MD5 do texto e o retorna em formato hexadecimal.
func Md5(texto string) string {
	h := md5.Sum([]byte(texto))
	return hex.EncodeToString(h[:])
}

// HmacSha256 gera um código de autenticação de mensagem HMAC baseado em SHA-256
// usando uma chave secreta e uma mensagem de texto.
//
// Exemplo em Verbo:
//
//	O assinado é HmacSha256 de Criptografia com ("chave_secreta", "mensagem").
func HmacSha256(chave, texto string) string {
	mac := hmac.New(sha256.New, []byte(chave))
	mac.Write([]byte(texto))
	return hex.EncodeToString(mac.Sum(nil))
}

// CompararHash realiza a comparação entre duas strings de hash em tempo constante,
// prevenindo ataques de temporização (timing attacks).
//
// Exemplo em Verbo:
//
//	O valido é CompararHash de Criptografia com (hash1, hash2).
func CompararHash(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}

// -----------------------------------------------------------------------------
// Codificação e Decodificação (Base64 e Hex)
// -----------------------------------------------------------------------------

// Base64Codificar converte uma string de texto em sua representação codificada em Base64 padrão.
//
// Exemplo em Verbo:
//
//	O b64 é Base64Codificar de Criptografia com ("Olá Mundo").
func Base64Codificar(texto string) string {
	return base64.StdEncoding.EncodeToString([]byte(texto))
}

// Base64Decodificar decodifica uma string Base64 padrão de volta para texto legível.
// Lança pânico caso a string de entrada não seja um Base64 válido.
//
// Exemplo em Verbo:
//
//	O original é Base64Decodificar de Criptografia com (b64).
func Base64Decodificar(b64 string) string {
	dados, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic("Erro ao decodificar Base64: " + err.Error())
	}
	return string(dados)
}

// Base64UrlCodificar converte texto para Base64 seguro para URLs (RFC 4648).
func Base64UrlCodificar(texto string) string {
	return base64.URLEncoding.EncodeToString([]byte(texto))
}

// Base64UrlDecodificar decodifica uma string Base64 URL-safe de volta para texto.
func Base64UrlDecodificar(b64 string) string {
	dados, err := base64.URLEncoding.DecodeString(b64)
	if err != nil {
		panic("Erro ao decodificar Base64 URL: " + err.Error())
	}
	return string(dados)
}

// HexCodificar converte uma sequência de texto em sua representação hexadecimal.
func HexCodificar(texto string) string {
	return hex.EncodeToString([]byte(texto))
}

// HexDecodificar decodifica uma string hexadecimal de volta para texto UTF-8.
func HexDecodificar(hexStr string) string {
	dados, err := hex.DecodeString(hexStr)
	if err != nil {
		panic("Erro ao decodificar Hexadecimal: " + err.Error())
	}
	return string(dados)
}

// Base32Codificar converte uma string de texto em sua representação codificada em Base32 padrão (RFC 4648).
// A saída usa o alfabeto maiúsculo A-Z e dígitos 2-7.
//
// Exemplo em Verbo:
//
//	O b32 é Base32Codificar de Criptografia com ("Olá Mundo").
func Base32Codificar(texto string) string {
	return base32.StdEncoding.EncodeToString([]byte(texto))
}

// Base32Decodificar decodifica uma string Base32 padrão de volta para texto legível.
// Lança pânico caso a string de entrada não seja um Base32 válido.
//
// Exemplo em Verbo:
//
//	O original é Base32Decodificar de Criptografia com (b32).
func Base32Decodificar(b32 string) string {
	dados, err := base32.StdEncoding.DecodeString(b32)
	if err != nil {
		panic("Erro ao decodificar Base32: " + err.Error())
	}
	return string(dados)
}

// -----------------------------------------------------------------------------
// Criptografia Simétrica (AES-GCM e Didáticas)
// -----------------------------------------------------------------------------

// normalizarChaveAES garante uma chave de 32 bytes (256 bits) para o AES.
// Se a chave fornecida tiver qualquer outro tamanho, usa o SHA-256 da chave.
func normalizarChaveAES(chave string) []byte {
	chaveBytes := []byte(chave)
	if len(chaveBytes) == 32 {
		return chaveBytes
	}
	hash := sha256.Sum256(chaveBytes)
	return hash[:]
}

// CifrarAES cifra um texto utilizando o algoritmo autenticado AES-256 no modo GCM.
// A chave é automaticamente normalizada para 32 bytes via SHA-256 se necessário.
// Retorna a mensagem cifrada combinada com o nonce, codificada em Base64.
//
// Exemplo em Verbo:
//
//	O segredo é CifrarAES de Criptografia com ("minha_senha_123", "Dados Confidenciais").
func CifrarAES(chave, texto string) string {
	chaveBytes := normalizarChaveAES(chave)

	block, err := aes.NewCipher(chaveBytes)
	if err != nil {
		panic("Erro ao inicializar cifra AES: " + err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("Erro ao inicializar modo GCM: " + err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := crand.Read(nonce); err != nil {
		panic("Erro ao gerar nonce aleatório para AES: " + err.Error())
	}

	cifrado := gcm.Seal(nonce, nonce, []byte(texto), nil)
	return base64.StdEncoding.EncodeToString(cifrado)
}

// DecifrarAES decifra uma mensagem previamente cifrada com CifrarAES.
// Lança pânico se a chave estiver incorreta ou se o texto tiver sido adulterado.
//
// Exemplo em Verbo:
//
//	O texto_original é DecifrarAES de Criptografia com ("minha_senha_123", segredo).
func DecifrarAES(chave, dadosCifradosBase64 string) string {
	chaveBytes := normalizarChaveAES(chave)

	dados, err := base64.StdEncoding.DecodeString(dadosCifradosBase64)
	if err != nil {
		panic("Erro ao decodificar Base64 dos dados cifrados: " + err.Error())
	}

	block, err := aes.NewCipher(chaveBytes)
	if err != nil {
		panic("Erro ao inicializar cifra AES: " + err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("Erro ao inicializar modo GCM: " + err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(dados) < nonceSize {
		panic("Dados cifrados corrompidos ou inválidos (tamanho inferior ao nonce).")
	}

	nonce, textoCifrado := dados[:nonceSize], dados[nonceSize:]
	original, err := gcm.Open(nil, nonce, textoCifrado, nil)
	if err != nil {
		panic("Erro ao decifrar dados: chave incorreta ou dados violados.")
	}

	return string(original)
}

// CifraCesar aplica a clássica cifra de deslocamento de César nas letras do texto.
// Caracteres fora do alfabeto A-Z são mantidos intactos.
//
// Exemplo em Verbo:
//
//	O cifrado é CifraCesar de Criptografia com ("Ola Mundo", 3).
func CifraCesar(texto string, deslocamento int) string {
	deslocamento = deslocamento % 26
	if deslocamento < 0 {
		deslocamento += 26
	}

	var sb strings.Builder
	for _, r := range texto {
		switch {
		case r >= 'a' && r <= 'z':
			sb.WriteRune('a' + (r-'a'+rune(deslocamento))%26)
		case r >= 'A' && r <= 'Z':
			sb.WriteRune('A' + (r-'A'+rune(deslocamento))%26)
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// Rot13 aplica a rotação de 13 posições no texto (equivalente a CifraCesar com deslocamento 13).
func Rot13(texto string) string {
	return CifraCesar(texto, 13)
}

// -----------------------------------------------------------------------------
// Geração Aleatória Segura (crypto/rand)
// -----------------------------------------------------------------------------

// GerarUUID gera um identificador único universal (UUID v4) aleatório segundo a RFC 4122.
// Formato: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
//
// Exemplo em Verbo:
//
//	O id é GerarUUID de Criptografia com ().
func GerarUUID() string {
	var b [16]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("Erro ao gerar bytes aleatórios para UUID: " + err.Error())
	}
	// Configurar bits da versão 4 e variante RFC 4122
	b[6] = (b[6] & 0x0f) | 0x40 // versão 4
	b[8] = (b[8] & 0x3f) | 0x80 // variante RFC 4122

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// GerarToken gera um token hexadecimal criptograficamente seguro com o tamanho especificado.
//
// Exemplo em Verbo:
//
//	O token é GerarToken de Criptografia com (32).
func GerarToken(tamanho int) string {
	if tamanho <= 0 {
		return ""
	}
	bytesQtd := (tamanho + 1) / 2
	b := make([]byte, bytesQtd)
	if _, err := crand.Read(b); err != nil {
		panic("Erro ao gerar token aleatório seguro: " + err.Error())
	}
	str := hex.EncodeToString(b)
	if len(str) > tamanho {
		return str[:tamanho]
	}
	return str
}

// GerarTextoAleatorio gera uma sequência alfanumérica segura contendo letras maiúsculas,
// minúsculas e números (a-z, A-Z, 0-9) com o tamanho requisitado. Útil para senhas seguras.
//
// Exemplo em Verbo:
//
//	O senha é GerarTextoAleatorio de Criptografia com (16).
func GerarTextoAleatorio(tamanho int) string {
	if tamanho <= 0 {
		return ""
	}
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	limite := big.NewInt(int64(len(chars)))

	res := make([]byte, tamanho)
	for i := 0; i < tamanho; i++ {
		n, err := crand.Int(crand.Reader, limite)
		if err != nil {
			panic("Erro ao gerar texto aleatório seguro: " + err.Error())
		}
		res[i] = chars[n.Int64()]
	}
	return string(res)
}

// GerarBytesHex gera uma quantidade N de bytes aleatórios criptograficamente seguros em formato hexadecimal.
func GerarBytesHex(quantidade int) string {
	if quantidade <= 0 {
		return ""
	}
	b := make([]byte, quantidade)
	if _, err := crand.Read(b); err != nil {
		panic("Erro ao gerar bytes aleatórios: " + err.Error())
	}
	return hex.EncodeToString(b)
}

// GerarNumeroSeguro retorna um número inteiro seguro no intervalo [0, maximo - 1].
func GerarNumeroSeguro(maximo int) int {
	if maximo <= 0 {
		return 0
	}
	n, err := crand.Int(crand.Reader, big.NewInt(int64(maximo)))
	if err != nil {
		panic("Erro ao gerar número seguro: " + err.Error())
	}
	return int(n.Int64())
}

// -----------------------------------------------------------------------------
// Geração Pseudoaleatória Utilitária (Estatística e Jogos)
// -----------------------------------------------------------------------------

// NumeroAleatorio retorna um número inteiro aleatório dentro do intervalo inclusivo [minimo, maximo].
//
// Exemplo em Verbo:
//
//	O dado é NumeroAleatorio de Criptografia com (1, 6).
func NumeroAleatorio(minimo, maximo int) int {
	if minimo > maximo {
		minimo, maximo = maximo, minimo
	}
	if minimo == maximo {
		return minimo
	}
	return minimo + mrand.Intn(maximo-minimo+1)
}

// DecimalAleatorio retorna um número decimal aleatório no intervalo [0.0, 1.0).
//
// Exemplo em Verbo:
//
//	O taxa é DecimalAleatorio de Criptografia com ().
func DecimalAleatorio() float64 {
	return mrand.Float64()
}

// BooleanoAleatorio retorna aleatoriamente Verdadeiro ou Falso (50% de probabilidade cada).
//
// Exemplo em Verbo:
//
//	O cara_ou_coroa é BooleanoAleatorio de Criptografia com ().
func BooleanoAleatorio() bool {
	return mrand.Intn(2) == 1
}

// EscolherAleatorio sorteia e retorna um elemento qualquer de uma lista.
// Lança pânico caso a lista esteja vazia.
//
// Exemplo em Verbo:
//
//	O sorteado é EscolherAleatorio de Criptografia com (jogadores).
func EscolherAleatorio(lista []interface{}) interface{} {
	if len(lista) == 0 {
		panic("Tentativa de escolher aleatório em uma lista vazia.")
	}
	idx := mrand.Intn(len(lista))
	return lista[idx]
}

// Embaralhar cria e retorna uma cópia da lista fornecida com seus elementos embaralhados aleatoriamente.
// A lista original não é modificada.
//
// Exemplo em Verbo:
//
//	O baralho_misturado é Embaralhar de Criptografia com (cartas).
func Embaralhar(lista []interface{}) []interface{} {
	tamanho := len(lista)
	copia := make([]interface{}, tamanho)
	copy(copia, lista)

	// Algoritmo Fisher-Yates
	for i := tamanho - 1; i > 0; i-- {
		j := mrand.Intn(i + 1)
		copia[i], copia[j] = copia[j], copia[i]
	}
	return copia
}
