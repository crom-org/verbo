package criptografia

import (
	"regexp"
	"strings"
	"testing"
)

func TestHash(t *testing.T) {
	texto := "Verbo Linguagem"

	s256 := Sha256(texto)
	if len(s256) != 64 {
		t.Fatalf("esperava 64 caracteres hex para SHA-256, obteve %d (%s)", len(s256), s256)
	}

	s512 := Sha512(texto)
	if len(s512) != 128 {
		t.Fatalf("esperava 128 caracteres hex para SHA-512, obteve %d", len(s512))
	}

	s1 := Sha1(texto)
	if len(s1) != 40 {
		t.Fatalf("esperava 40 caracteres hex para SHA-1, obteve %d", len(s1))
	}

	m5 := Md5(texto)
	if len(m5) != 32 {
		t.Fatalf("esperava 32 caracteres hex para MD5, obteve %d", len(m5))
	}

	hm := HmacSha256("chave_teste", texto)
	if len(hm) != 64 {
		t.Fatalf("esperava 64 caracteres hex para HMAC-SHA256, obteve %d", len(hm))
	}

	if !CompararHash(s256, s256) {
		t.Fatalf("esperava CompararHash verdadeiro para hashes iguais")
	}
	if CompararHash(s256, "outro_hash") {
		t.Fatalf("esperava CompararHash falso para hashes diferentes")
	}
}

func TestCodificacao(t *testing.T) {
	original := "Mensagem secreta com acentuação: olá mundo! 🇧🇷"

	// Base64
	b64 := Base64Codificar(original)
	decB64 := Base64Decodificar(b64)
	if decB64 != original {
		t.Fatalf("esperava %q após decodificar Base64, obteve %q", original, decB64)
	}

	// Base64 URL
	b64url := Base64UrlCodificar(original)
	decB64url := Base64UrlDecodificar(b64url)
	if decB64url != original {
		t.Fatalf("esperava %q após decodificar Base64Url, obteve %q", original, decB64url)
	}

	// Hex
	hexStr := HexCodificar(original)
	decHex := HexDecodificar(hexStr)
	if decHex != original {
		t.Fatalf("esperava %q após decodificar Hex, obteve %q", original, decHex)
	}
}

func TestCifrarEDecifrarAES(t *testing.T) {
	chave := "senha-super-secreta"
	mensagemOriginal := "Texto confidencial de teste para a BibVerbo."

	cifrado := CifrarAES(chave, mensagemOriginal)
	if cifrado == mensagemOriginal {
		t.Fatalf("esperava texto cifrado diferente do original")
	}

	decifrado := DecifrarAES(chave, cifrado)
	if decifrado != mensagemOriginal {
		t.Fatalf("esperava %q, obteve %q", mensagemOriginal, decifrado)
	}

	// Teste com chave errada (deve disparar pânico)
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("esperava pânico ao decifrar com chave errada")
		}
	}()
	DecifrarAES("senha-errada", cifrado)
}

func TestCifrasClassicas(t *testing.T) {
	texto := "Ola, Mundo!"
	cifrado := CifraCesar(texto, 3)
	if cifrado != "Rod, Pxqgr!" {
		t.Fatalf("esperava 'Rod, Pxqgr!', obteve %q", cifrado)
	}
	decifrado := CifraCesar(cifrado, -3)
	if decifrado != texto {
		t.Fatalf("esperava %q após reverter César, obteve %q", texto, decifrado)
	}

	r13 := Rot13(texto)
	r13Reverso := Rot13(r13)
	if r13Reverso != texto {
		t.Fatalf("esperava duplo ROT13 restaurar %q, obteve %q", texto, r13Reverso)
	}
}

func TestGeracaoAleatoriaSegura(t *testing.T) {
	uuid := GerarUUID()
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(uuid) {
		t.Fatalf("UUID gerado não segue a RFC 4122 v4: %s", uuid)
	}

	token := GerarToken(20)
	if len(token) != 20 {
		t.Fatalf("esperava token de tamanho 20, obteve tamanho %d (%s)", len(token), token)
	}

	textoAleatorio := GerarTextoAleatorio(16)
	if len(textoAleatorio) != 16 {
		t.Fatalf("esperava texto de tamanho 16, obteve %d (%s)", len(textoAleatorio), textoAleatorio)
	}

	bytesHex := GerarBytesHex(8)
	if len(bytesHex) != 16 {
		t.Fatalf("esperava 16 chars hex para 8 bytes, obteve %d (%s)", len(bytesHex), bytesHex)
	}

	numSeguro := GerarNumeroSeguro(10)
	if numSeguro < 0 || numSeguro >= 10 {
		t.Fatalf("esperava numSeguro em [0, 9], obteve %d", numSeguro)
	}
}

func TestGeracaoAleatoriaUtilitaria(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := NumeroAleatorio(10, 20)
		if n < 10 || n > 20 {
			t.Fatalf("NumeroAleatorio fora da faixa [10, 20]: %d", n)
		}

		dec := DecimalAleatorio()
		if dec < 0.0 || dec >= 1.0 {
			t.Fatalf("DecimalAleatorio fora da faixa [0.0, 1.0): %f", dec)
		}
	}

	// BooleanoAleatorio deve produzir ao menos uma alternância em muitas iterações
	viuVerdadeiro, viuFalso := false, false
	for i := 0; i < 100; i++ {
		if BooleanoAleatorio() {
			viuVerdadeiro = true
		} else {
			viuFalso = true
		}
	}
	if !viuVerdadeiro || !viuFalso {
		t.Fatalf("BooleanoAleatorio não variou valores após 100 iterações")
	}

	// EscolherAleatorio
	lista := []interface{}{"A", "B", "C"}
	escolhido := EscolherAleatorio(lista)
	s := escolhido.(string)
	if !strings.Contains("ABC", s) {
		t.Fatalf("elemento escolhido inválido: %v", escolhido)
	}

	// Embaralhar
	copia := Embaralhar(lista)
	if len(copia) != len(lista) {
		t.Fatalf("tamanho de lista embaralhada difere da original")
	}
}
