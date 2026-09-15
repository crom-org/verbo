package internet

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestHttpBasico(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("resposta GET"))
		case http.MethodPost:
			corpo, _ := io.ReadAll(r.Body)
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte("recebido POST: " + string(corpo)))
		case http.MethodPut:
			corpo, _ := io.ReadAll(r.Body)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("recebido PUT: " + string(corpo)))
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("deletado com sucesso"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer ts.Close()

	// Teste Obter e Get
	respGet := Obter(ts.URL)
	if respGet != "resposta GET" {
		t.Fatalf("esperava 'resposta GET', obteve %q", respGet)
	}
	respGetAlias := Get(ts.URL)
	if respGetAlias != "resposta GET" {
		t.Fatalf("esperava 'resposta GET', obteve %q", respGetAlias)
	}

	// Teste Postar e Post
	respPost := Postar(ts.URL, "dados de teste", "text/plain")
	if respPost != "recebido POST: dados de teste" {
		t.Fatalf("esperava resposta de POST correta, obteve %q", respPost)
	}
	respPostJson := Post(ts.URL, "{\"ok\":true}")
	if respPostJson != "recebido POST: {\"ok\":true}" {
		t.Fatalf("esperava resposta de Post JSON correta, obteve %q", respPostJson)
	}

	// Teste Put
	respPut := Put(ts.URL, "dados put", "text/plain")
	if respPut != "recebido PUT: dados put" {
		t.Fatalf("esperava resposta de PUT correta, obteve %q", respPut)
	}

	// Teste Deletar
	respDelete := Deletar(ts.URL)
	if respDelete != "deletado com sucesso" {
		t.Fatalf("esperava resposta de DELETE correta, obteve %q", respDelete)
	}

	// Teste Requisicao genérica
	respReq := Requisicao("POST", ts.URL, "meu corpo", "text/plain")
	if respReq != "recebido POST: meu corpo" {
		t.Fatalf("esperava resposta de Requisicao correta, obteve %q", respReq)
	}

	// Teste Status
	status := Status("GET", ts.URL)
	if status != 200 {
		t.Fatalf("esperava status 200, obteve %d", status)
	}
}

func TestBaixarArquivo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("conteudo para download"))
	}))
	defer ts.Close()

	tempDir := t.TempDir()
	destino := filepath.Join(tempDir, "arquivo_baixado.txt")

	resultado := Baixar(ts.URL, destino)
	if resultado != destino {
		t.Fatalf("esperava retorno %q, obteve %q", destino, resultado)
	}

	conteudo, err := os.ReadFile(destino)
	if err != nil {
		t.Fatalf("erro ao ler arquivo baixado: %v", err)
	}
	if string(conteudo) != "conteudo para download" {
		t.Fatalf("conteúdo inesperado: %q", string(conteudo))
	}
}

func TestDefinirTempoLimite(t *testing.T) {
	DefinirTempoLimite(5)
	if obterCliente().Timeout != 5*time.Second {
		t.Fatalf("esperava timeout de 5 segundos, obteve %v", obterCliente().Timeout)
	}
}

func TestSocketsTCP(t *testing.T) {
	// Iniciar servidor TCP em porta livre
	srv := OuvirTCP("127.0.0.1:0")
	defer srv.Fechar()

	endereco := srv.Endereco()
	if !strings.HasPrefix(endereco, "127.0.0.1:") {
		t.Fatalf("esperava endereço iniciando com 127.0.0.1:, obteve %s", endereco)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Goroutine do servidor para aceitar conexão e responder
	go func() {
		defer wg.Done()
		cliente := srv.Aceitar()
		defer cliente.Fechar()

		linha := cliente.LerLinha()
		if linha != "MENSAGEM_CLIENTE" {
			t.Errorf("servidor recebeu linha inesperada: %q", linha)
		}

		cliente.EnviarLinha("ECO: " + linha)
	}()

	// Conectar cliente TCP
	conn := ConectarTCP(endereco)
	defer conn.Fechar()

	if conn.EstaFechada() {
		t.Fatalf("esperava conexão aberta")
	}

	bytesEnviados := conn.EnviarLinha("MENSAGEM_CLIENTE")
	if bytesEnviados == 0 {
		t.Fatalf("esperava bytes enviados > 0")
	}

	resposta := conn.LerLinha()
	if resposta != "ECO: MENSAGEM_CLIENTE" {
		t.Fatalf("esperava 'ECO: MENSAGEM_CLIENTE', obteve %q", resposta)
	}

	wg.Wait()
}

func TestStandaloneTCP(t *testing.T) {
	srv := OuvirTCP("127.0.0.1:0")
	defer FecharServidorTCP(srv)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		c := AceitarTCP(srv)
		defer FecharTCP(c)

		msg := ReceberTCP(c, 1024)
		if msg != "TESTE_STANDALONE" {
			t.Errorf("esperava 'TESTE_STANDALONE', obteve %q", msg)
		}
		_ = EnviarTCP(c, "RESPOSTA_STANDALONE")
	}()

	conn := ConectarTCP(srv.Endereco())
	defer FecharTCP(conn)

	_ = EnviarTCP(conn, "TESTE_STANDALONE")
	resp := ReceberTCP(conn, 1024)
	if resp != "RESPOSTA_STANDALONE" {
		t.Fatalf("esperava 'RESPOSTA_STANDALONE', obteve %q", resp)
	}

	wg.Wait()
}

func TestSocketsUDP(t *testing.T) {
	udpSrv := OuvirUDP("127.0.0.1:0")
	defer udpSrv.Fechar()

	end := udpSrv.EnderecoLocal()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		msg := udpSrv.Receber(1024)
		if msg != "PACOTE_UDP" {
			t.Errorf("esperava 'PACOTE_UDP', obteve %q", msg)
		}
	}()

	n := EnviarUDP(end, "PACOTE_UDP")
	if n != len("PACOTE_UDP") {
		t.Fatalf("esperava envio de %d bytes, enviou %d", len("PACOTE_UDP"), n)
	}

	wg.Wait()
}

func TestUtilitariosRede(t *testing.T) {
	ips := ResolverHost("localhost")
	if len(ips) == 0 {
		t.Fatalf("esperava ao menos um IP para localhost")
	}

	ipLocal := ObterIPLocal()
	if ipLocal == "" {
		t.Fatalf("esperava IP local não vazio")
	}

	srv := OuvirTCP("127.0.0.1:0")
	defer srv.Fechar()

	partes := strings.Split(srv.Endereco(), ":")
	portaStr := partes[len(partes)-1]
	var porta int
	_, _ = fmt.Sscanf(portaStr, "%d", &porta)

	if !PortaAberta("127.0.0.1", porta) {
		t.Fatalf("esperava porta %d aberta", porta)
	}

	if PortaAberta("127.0.0.1", 59999) {
		// porta aleatória improvável de estar aberta
		// Se estiver aberta, ignorar; geralmente fechada
	}
}
