// Package internet implementa recursos de rede, incluindo cliente HTTP e sockets TCP/UDP
// para a linguagem de programação Verbo.
// Parte da BibVerbo (Biblioteca Padrão do Verbo).
//
// Uso em Verbo:
//
//	Incluir Internet.
//	O resposta é Obter de Internet com ("https://api.exemplo.com/dados").
//	Exibir resposta.
package internet

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	clienteHttpPadrao = &http.Client{
		Timeout: 30 * time.Second,
	}
	muCliente sync.RWMutex
)

// DefinirTempoLimite define o tempo limite padrão em segundos para requisições HTTP.
//
// Exemplo em Verbo:
//
//	DefinirTempoLimite de Internet com (10).
func DefinirTempoLimite(segundos int) {
	muCliente.Lock()
	defer muCliente.Unlock()
	clienteHttpPadrao = &http.Client{
		Timeout: time.Duration(segundos) * time.Second,
	}
}

func obterCliente() *http.Client {
	muCliente.RLock()
	defer muCliente.RUnlock()
	return clienteHttpPadrao
}

// -----------------------------------------------------------------------------
// Cliente HTTP
// -----------------------------------------------------------------------------

// Obter faz uma requisição HTTP GET para a URL especificada e retorna o corpo como texto.
// Lança pânico caso ocorra falha de conexão ou erro HTTP.
//
// Exemplo em Verbo:
//
//	O resposta é Obter de Internet com ("https://jsonplaceholder.typicode.com/todos/1").
func Obter(url string) string {
	resp, err := obterCliente().Get(url)
	if err != nil {
		panic("Erro ao executar requisição GET em " + url + ": " + err.Error())
	}
	defer resp.Body.Close()

	corpo, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Erro ao ler resposta da requisição GET em " + url + ": " + err.Error())
	}

	return string(corpo)
}

// Get é um alias em inglês para Obter.
func Get(url string) string {
	return Obter(url)
}

// Postar faz uma requisição HTTP POST para a URL com o corpo e tipo de conteúdo (Content-Type).
// Se tipoConteudo for vazio, assume "application/json".
//
// Exemplo em Verbo:
//
//	O resp é Postar de Internet com ("https://httpbin.org/post", "{\"nome\":\"Verbo\"}", "application/json").
func Postar(url, corpo, tipoConteudo string) string {
	if tipoConteudo == "" {
		tipoConteudo = "application/json"
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(corpo))
	if err != nil {
		panic("Erro ao preparar requisição POST em " + url + ": " + err.Error())
	}
	req.Header.Set("Content-Type", tipoConteudo)

	resp, err := obterCliente().Do(req)
	if err != nil {
		panic("Erro ao executar requisição POST em " + url + ": " + err.Error())
	}
	defer resp.Body.Close()

	dados, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Erro ao ler resposta da requisição POST em " + url + ": " + err.Error())
	}

	return string(dados)
}

// Post é um alias para Postar enviando Content-Type "application/json".
func Post(url, corpo string) string {
	return Postar(url, corpo, "application/json")
}

// Put faz uma requisição HTTP PUT para a URL especificada com corpo e tipo de conteúdo.
//
// Exemplo em Verbo:
//
//	O resp é Put de Internet com ("https://httpbin.org/put", "{\"id\":1}", "application/json").
func Put(url, corpo, tipoConteudo string) string {
	if tipoConteudo == "" {
		tipoConteudo = "application/json"
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBufferString(corpo))
	if err != nil {
		panic("Erro ao preparar requisição PUT em " + url + ": " + err.Error())
	}
	req.Header.Set("Content-Type", tipoConteudo)

	resp, err := obterCliente().Do(req)
	if err != nil {
		panic("Erro ao executar requisição PUT em " + url + ": " + err.Error())
	}
	defer resp.Body.Close()

	dados, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Erro ao ler resposta da requisição PUT em " + url + ": " + err.Error())
	}

	return string(dados)
}

// Deletar faz uma requisição HTTP DELETE para a URL especificada.
//
// Exemplo em Verbo:
//
//	O resp é Deletar de Internet com ("https://httpbin.org/delete").
func Deletar(url string) string {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic("Erro ao preparar requisição DELETE em " + url + ": " + err.Error())
	}

	resp, err := obterCliente().Do(req)
	if err != nil {
		panic("Erro ao executar requisição DELETE em " + url + ": " + err.Error())
	}
	defer resp.Body.Close()

	dados, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Erro ao ler resposta da requisição DELETE em " + url + ": " + err.Error())
	}

	return string(dados)
}

// Requisicao executa uma requisição HTTP genérica (GET, POST, PUT, DELETE, PATCH, etc.).
//
// Exemplo em Verbo:
//
//	O resp é Requisicao de Internet com ("PATCH", "https://api.exemplo.com/item/1", "{\"ativo\":true}", "application/json").
func Requisicao(metodo, url, corpo, tipoConteudo string) string {
	metodo = strings.ToUpper(metodo)
	var bodyReader io.Reader
	if corpo != "" {
		bodyReader = bytes.NewBufferString(corpo)
	}

	req, err := http.NewRequest(metodo, url, bodyReader)
	if err != nil {
		panic("Erro ao criar requisição " + metodo + " em " + url + ": " + err.Error())
	}
	if tipoConteudo != "" {
		req.Header.Set("Content-Type", tipoConteudo)
	}

	resp, err := obterCliente().Do(req)
	if err != nil {
		panic("Erro ao executar requisição " + metodo + " em " + url + ": " + err.Error())
	}
	defer resp.Body.Close()

	dados, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Erro ao ler resposta da requisição " + metodo + " em " + url + ": " + err.Error())
	}

	return string(dados)
}

// Status retorna o código de resposta HTTP (como 200, 404, 500) para uma URL e método.
//
// Exemplo em Verbo:
//
//	O codigo é Status de Internet com ("GET", "https://google.com").
func Status(metodo, url string) int {
	metodo = strings.ToUpper(metodo)
	req, err := http.NewRequest(metodo, url, nil)
	if err != nil {
		panic("Erro ao criar requisição em " + url + ": " + err.Error())
	}

	resp, err := obterCliente().Do(req)
	if err != nil {
		panic("Erro ao executar requisição em " + url + ": " + err.Error())
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

// Baixar faz o download do conteúdo da URL e grava diretamente em um arquivo no disco.
// Retorna o caminho de destino em caso de sucesso.
//
// Exemplo em Verbo:
//
//	Baixar de Internet com ("https://exemplo.com/foto.png", "foto.png").
func Baixar(url, caminhoDestino string) string {
	resp, err := obterCliente().Get(url)
	if err != nil {
		panic("Erro ao baixar " + url + ": " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		panic(fmt.Sprintf("Erro ao baixar %s: status HTTP %d", url, resp.StatusCode))
	}

	arq, err := os.Create(caminhoDestino)
	if err != nil {
		panic("Erro ao criar arquivo de destino " + caminhoDestino + ": " + err.Error())
	}
	defer arq.Close()

	_, err = io.Copy(arq, resp.Body)
	if err != nil {
		panic("Erro ao gravar dados no arquivo " + caminhoDestino + ": " + err.Error())
	}

	return caminhoDestino
}

// -----------------------------------------------------------------------------
// Sockets TCP e Conexões
// -----------------------------------------------------------------------------

// Conexao encapsula uma conexão de rede (como TCP).
type Conexao struct {
	conn    net.Conn
	leitor  *bufio.Reader
	fechada bool
	mu      sync.Mutex
}

func novaConexao(c net.Conn) *Conexao {
	return &Conexao{
		conn:   c,
		leitor: bufio.NewReader(c),
	}
}

// Enviar envia uma mensagem de texto pela conexão. Retorna a quantidade de bytes enviados.
func (c *Conexao) Enviar(mensagem string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.fechada {
		panic("Tentativa de enviar dados em conexão já fechada.")
	}
	n, err := c.conn.Write([]byte(mensagem))
	if err != nil {
		panic("Erro ao enviar dados pela conexão: " + err.Error())
	}
	return n
}

// EnviarLinha envia uma mensagem seguida de quebra de linha (\n).
func (c *Conexao) EnviarLinha(mensagem string) int {
	return c.Enviar(mensagem + "\n")
}

// Receber lê até tamanhoMaximo bytes da conexão e retorna como texto.
func (c *Conexao) Receber(tamanhoMaximo int) string {
	if tamanhoMaximo <= 0 {
		tamanhoMaximo = 4096
	}
	buf := make([]byte, tamanhoMaximo)
	n, err := c.conn.Read(buf)
	if err != nil && err != io.EOF {
		panic("Erro ao ler da conexão: " + err.Error())
	}
	return string(buf[:n])
}

// LerLinha lê da conexão até encontrar uma quebra de linha (\n).
func (c *Conexao) LerLinha() string {
	linha, err := c.leitor.ReadString('\n')
	if err != nil && err != io.EOF {
		panic("Erro ao ler linha da conexão: " + err.Error())
	}
	return strings.TrimRight(linha, "\r\n")
}

// Fechar encerra a conexão de rede.
func (c *Conexao) Fechar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.fechada {
		c.fechada = true
		_ = c.conn.Close()
	}
}

// EnderecoRemoto retorna o endereço IP e porta do ponto remoto.
func (c *Conexao) EnderecoRemoto() string {
	return c.conn.RemoteAddr().String()
}

// EnderecoLocal retorna o endereço IP e porta local da conexão.
func (c *Conexao) EnderecoLocal() string {
	return c.conn.LocalAddr().String()
}

// EstaFechada retorna verdadeiro se a conexão já foi fechada.
func (c *Conexao) EstaFechada() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.fechada
}

// ConectarTCP estabelece uma conexão cliente TCP com o endereço fornecido (ex: "127.0.0.1:8080").
//
// Exemplo em Verbo:
//
//	O socket é ConectarTCP de Internet com ("127.0.0.1:9000").
//	Enviar de socket com ("Olá servidor!\n").
func ConectarTCP(endereco string) *Conexao {
	conn, err := net.DialTimeout("tcp", endereco, 10*time.Second)
	if err != nil {
		panic("Erro ao conectar TCP em " + endereco + ": " + err.Error())
	}
	return novaConexao(conn)
}

// EnviarTCP é uma função standalone para enviar texto através de uma conexão TCP.
func EnviarTCP(c *Conexao, mensagem string) int {
	if c == nil {
		panic("Conexão nula ao chamar EnviarTCP.")
	}
	return c.Enviar(mensagem)
}

// ReceberTCP é uma função standalone para receber dados de uma conexão TCP.
func ReceberTCP(c *Conexao, tamanhoMaximo int) string {
	if c == nil {
		panic("Conexão nula ao chamar ReceberTCP.")
	}
	return c.Receber(tamanhoMaximo)
}

// LerLinhaTCP é uma função standalone para ler uma linha de uma conexão TCP.
func LerLinhaTCP(c *Conexao) string {
	if c == nil {
		panic("Conexão nula ao chamar LerLinhaTCP.")
	}
	return c.LerLinha()
}

// FecharTCP fecha uma conexão TCP.
func FecharTCP(c *Conexao) {
	if c != nil {
		c.Fechar()
	}
}

// -----------------------------------------------------------------------------
// Servidor TCP / Listener
// -----------------------------------------------------------------------------

// ServidorTCP encapsula um listener TCP para escutar conexões de entrada.
type ServidorTCP struct {
	listener net.Listener
	fechado  bool
	mu       sync.Mutex
}

// OuvirTCP cria e inicia um listener TCP no endereço fornecido (ex: ":9000" ou "127.0.0.1:9000").
//
// Exemplo em Verbo:
//
//	O servidor é OuvirTCP de Internet com (":9000").
//	O cliente é Aceitar de servidor com ().
func OuvirTCP(endereco string) *ServidorTCP {
	ln, err := net.Listen("tcp", endereco)
	if err != nil {
		panic("Erro ao escutar TCP no endereço " + endereco + ": " + err.Error())
	}
	return &ServidorTCP{listener: ln}
}

// Aceitar aguarda e aceita a próxima conexão cliente TCP. Bloqueia a execução até uma conexão chegar.
func (s *ServidorTCP) Aceitar() *Conexao {
	conn, err := s.listener.Accept()
	if err != nil {
		panic("Erro ao aceitar conexão TCP: " + err.Error())
	}
	return novaConexao(conn)
}

// Fechar encerra o listener TCP.
func (s *ServidorTCP) Fechar() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.fechado {
		s.fechado = true
		_ = s.listener.Close()
	}
}

// Endereco retorna a string do endereço em que o servidor está escutando.
func (s *ServidorTCP) Endereco() string {
	return s.listener.Addr().String()
}

// ObterEndereco retorna o endereço em que o servidor está escutando.
func (s *ServidorTCP) ObterEndereco() string {
	return s.Endereco()
}

// ObterEnderecoTCP retorna a string do endereço do servidor TCP fornecido.
func ObterEnderecoTCP(s *ServidorTCP) string {
	if s == nil {
		panic("ServidorTCP nulo ao chamar ObterEnderecoTCP.")
	}
	return s.Endereco()
}

// AceitarTCP é a função standalone para aceitar cliente em um servidor TCP.
func AceitarTCP(s *ServidorTCP) *Conexao {
	if s == nil {
		panic("ServidorTCP nulo ao chamar AceitarTCP.")
	}
	return s.Aceitar()
}

// FecharServidorTCP encerra o servidor TCP fornecido.
func FecharServidorTCP(s *ServidorTCP) {
	if s != nil {
		s.Fechar()
	}
}

// -----------------------------------------------------------------------------
// Sockets UDP
// -----------------------------------------------------------------------------

// ConexaoUDP encapsula um socket UDP para envio e recepção de datagramas.
type ConexaoUDP struct {
	conn    *net.UDPConn
	fechada bool
	mu      sync.Mutex
}

// OuvirUDP inicia a escuta UDP no endereço especificado (ex: ":9090").
func OuvirUDP(endereco string) *ConexaoUDP {
	addr, err := net.ResolveUDPAddr("udp", endereco)
	if err != nil {
		panic("Erro ao resolver endereço UDP " + endereco + ": " + err.Error())
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic("Erro ao escutar UDP em " + endereco + ": " + err.Error())
	}
	return &ConexaoUDP{conn: conn}
}

// Receber lê até tamanhoMaximo bytes do socket UDP.
func (u *ConexaoUDP) Receber(tamanhoMaximo int) string {
	if tamanhoMaximo <= 0 {
		tamanhoMaximo = 2048
	}
	buf := make([]byte, tamanhoMaximo)
	n, _, err := u.conn.ReadFromUDP(buf)
	if err != nil {
		panic("Erro ao receber pacote UDP: " + err.Error())
	}
	return string(buf[:n])
}

// Fechar encerra o socket UDP.
func (u *ConexaoUDP) Fechar() {
	u.mu.Lock()
	defer u.mu.Unlock()
	if !u.fechada {
		u.fechada = true
		_ = u.conn.Close()
	}
}

// EnderecoLocal retorna o endereço local do socket UDP.
func (u *ConexaoUDP) EnderecoLocal() string {
	return u.conn.LocalAddr().String()
}

// EnviarUDP envia uma mensagem de texto para o endereço UDP de destino (ex: "127.0.0.1:9090").
func EnviarUDP(endereco, mensagem string) int {
	addr, err := net.ResolveUDPAddr("udp", endereco)
	if err != nil {
		panic("Erro ao resolver destino UDP " + endereco + ": " + err.Error())
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic("Erro ao conectar socket UDP em " + endereco + ": " + err.Error())
	}
	defer conn.Close()

	n, err := conn.Write([]byte(mensagem))
	if err != nil {
		panic("Erro ao enviar mensagem UDP para " + endereco + ": " + err.Error())
	}
	return n
}

// ReceberUDP é a função standalone para receber dados de um socket UDP.
func ReceberUDP(u *ConexaoUDP, tamanhoMaximo int) string {
	if u == nil {
		panic("Socket UDP nulo ao chamar ReceberUDP.")
	}
	return u.Receber(tamanhoMaximo)
}

// FecharUDP fecha o socket UDP.
func FecharUDP(u *ConexaoUDP) {
	if u != nil {
		u.Fechar()
	}
}

// -----------------------------------------------------------------------------
// Utilitários de Rede e DNS
// -----------------------------------------------------------------------------

// ResolverHost retorna a lista de endereços IP correspondentes a um domínio/host.
//
// Exemplo em Verbo:
//
//	O ips é ResolverHost de Internet com ("localhost").
func ResolverHost(host string) []interface{} {
	ips, err := net.LookupHost(host)
	if err != nil {
		panic("Erro ao resolver host " + host + ": " + err.Error())
	}
	res := make([]interface{}, len(ips))
	for i, ip := range ips {
		res[i] = ip
	}
	return res
}

// PortaAberta verifica se uma determinada porta TCP em um host está aberta e acessível.
//
// Exemplo em Verbo:
//
//	O aberta é PortaAberta de Internet com ("127.0.0.1", 8080).
func PortaAberta(host string, porta int) bool {
	endereco := fmt.Sprintf("%s:%d", host, porta)
	conn, err := net.DialTimeout("tcp", endereco, 2*time.Second)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

// ObterIPLocal retorna o primeiro endereço IPv4 local não-loopback da máquina.
// Caso não encontre, retorna "127.0.0.1".
//
// Exemplo em Verbo:
//
//	O ip é ObterIPLocal de Internet com ().
func ObterIPLocal() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
