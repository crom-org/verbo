// Package sistema implementa funções de interação com o sistema operacional,
// incluindo execução de comandos, leitura de argumentos da linha de comando,
// variáveis de ambiente e informações da máquina.
// Parte da BibVerbo (Biblioteca Padrão do Verbo).
//
// Uso em Verbo:
//
//	Incluir Sistema.
//	O args é Argumentos de Sistema com ().
//	O so é NomeSistemaOperacional de Sistema com ().
package sistema

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// -----------------------------------------------------------------------------
// Argumentos da Linha de Comando
// -----------------------------------------------------------------------------

// Argumentos retorna a lista de argumentos passados ao programa na linha de comando,
// excluindo o próprio nome do executável (os.Args[0]).
//
// Exemplo em Verbo:
//
//	O args é Argumentos de Sistema com ().
func Argumentos() []interface{} {
	args := os.Args[1:]
	res := make([]interface{}, len(args))
	for i, a := range args {
		res[i] = a
	}
	return res
}

// Argumento retorna o argumento da linha de comando na posição indicada (base 0).
// Lança pânico caso o índice esteja fora do intervalo disponível.
//
// Exemplo em Verbo:
//
//	O primeiro é Argumento de Sistema com (0).
func Argumento(indice int) string {
	args := os.Args[1:]
	if indice < 0 || indice >= len(args) {
		panic(fmt.Sprintf("Argumento: índice %d fora do intervalo (total de argumentos: %d).", indice, len(args)))
	}
	return args[indice]
}

// NumeroDeArgumentos retorna a quantidade de argumentos passados ao programa,
// sem contar o nome do executável.
//
// Exemplo em Verbo:
//
//	O quantidade é NumeroDeArgumentos de Sistema com ().
func NumeroDeArgumentos() int {
	return len(os.Args) - 1
}

// -----------------------------------------------------------------------------
// Variáveis de Ambiente
// -----------------------------------------------------------------------------

// VariavelDeAmbiente retorna o valor da variável de ambiente com o nome fornecido.
// Retorna uma string vazia se a variável não estiver definida.
//
// Exemplo em Verbo:
//
//	O caminho é VariavelDeAmbiente de Sistema com ("PATH").
func VariavelDeAmbiente(nome string) string {
	return os.Getenv(nome)
}

// DefinirVariavelDeAmbiente define uma variável de ambiente no processo atual.
// Lança pânico se ocorrer falha ao definir a variável.
//
// Exemplo em Verbo:
//
//	DefinirVariavelDeAmbiente de Sistema com ("MINHA_VAR", "valor").
func DefinirVariavelDeAmbiente(nome, valor string) {
	if err := os.Setenv(nome, valor); err != nil {
		panic("Erro ao definir variável de ambiente '" + nome + "': " + err.Error())
	}
}

// RemoverVariavelDeAmbiente remove uma variável de ambiente do processo atual.
// Lança pânico se ocorrer falha.
//
// Exemplo em Verbo:
//
//	RemoverVariavelDeAmbiente de Sistema com ("MINHA_VAR").
func RemoverVariavelDeAmbiente(nome string) {
	if err := os.Unsetenv(nome); err != nil {
		panic("Erro ao remover variável de ambiente '" + nome + "': " + err.Error())
	}
}

// VariaveisDeAmbiente retorna todas as variáveis de ambiente do processo como
// uma lista de strings no formato "NOME=valor".
//
// Exemplo em Verbo:
//
//	O vars é VariaveisDeAmbiente de Sistema com ().
func VariaveisDeAmbiente() []interface{} {
	env := os.Environ()
	res := make([]interface{}, len(env))
	for i, e := range env {
		res[i] = e
	}
	return res
}

// -----------------------------------------------------------------------------
// Execução de Comandos
// -----------------------------------------------------------------------------

// ExecutarComando executa uma linha de comando no shell padrão do sistema
// (/bin/sh -c no Unix, cmd /C no Windows) e retorna a saída padrão como texto.
// Lança pânico se o comando não puder ser iniciado ou retornar erro.
//
// Exemplo em Verbo:
//
//	O saida é ExecutarComando de Sistema com ("echo Olá Mundo").
func ExecutarComando(comando string) string {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", comando)
	} else {
		cmd = exec.Command("/bin/sh", "-c", comando)
	}
	saida, err := cmd.Output()
	if err != nil {
		panic("Erro ao executar comando '" + comando + "': " + err.Error())
	}
	return strings.TrimRight(string(saida), "\r\n")
}

// ExecutarComandoComArgumentos executa um programa diretamente com a lista de
// argumentos fornecida, sem intermediação de shell. É mais seguro que
// ExecutarComando pois evita injeção de shell.
// Lança pânico se o programa não puder ser iniciado ou retornar erro.
//
// Exemplo em Verbo:
//
//	O saida é ExecutarComandoComArgumentos de Sistema com ("ls", ["-la", "/tmp"]).
func ExecutarComandoComArgumentos(programa string, args []interface{}) string {
	strArgs := make([]string, len(args))
	for i, a := range args {
		strArgs[i] = fmt.Sprintf("%v", a)
	}
	cmd := exec.Command(programa, strArgs...)
	saida, err := cmd.Output()
	if err != nil {
		panic("Erro ao executar '" + programa + "': " + err.Error())
	}
	return strings.TrimRight(string(saida), "\r\n")
}

// ExecutarComandoComEntrada executa um comando shell fornecendo texto como entrada padrão (stdin).
// Retorna a saída padrão como texto.
// Lança pânico se o comando falhar.
//
// Exemplo em Verbo:
//
//	O resultado é ExecutarComandoComEntrada de Sistema com ("wc -w", "Olá Verbo mundo").
func ExecutarComandoComEntrada(comando, entrada string) string {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", comando)
	} else {
		cmd = exec.Command("/bin/sh", "-c", comando)
	}
	cmd.Stdin = strings.NewReader(entrada)
	saida, err := cmd.Output()
	if err != nil {
		panic("Erro ao executar comando '" + comando + "' com entrada: " + err.Error())
	}
	return strings.TrimRight(string(saida), "\r\n")
}

// -----------------------------------------------------------------------------
// Diretório de Trabalho
// -----------------------------------------------------------------------------

// DiretorioAtual retorna o caminho absoluto do diretório de trabalho atual.
// Lança pânico se não for possível obter o diretório.
//
// Exemplo em Verbo:
//
//	O caminho é DiretorioAtual de Sistema com ().
func DiretorioAtual() string {
	dir, err := os.Getwd()
	if err != nil {
		panic("Erro ao obter diretório atual: " + err.Error())
	}
	return dir
}

// MudarDiretorio muda o diretório de trabalho atual do processo para o caminho indicado.
// Lança pânico se o caminho não existir ou não for acessível.
//
// Exemplo em Verbo:
//
//	MudarDiretorio de Sistema com ("/tmp").
func MudarDiretorio(caminho string) {
	if err := os.Chdir(caminho); err != nil {
		panic("Erro ao mudar para o diretório '" + caminho + "': " + err.Error())
	}
}

// -----------------------------------------------------------------------------
// Informações do Sistema
// -----------------------------------------------------------------------------

// NomeSistemaOperacional retorna o nome do sistema operacional em letras minúsculas.
// Valores comuns: "linux", "windows", "darwin".
//
// Exemplo em Verbo:
//
//	O so é NomeSistemaOperacional de Sistema com ().
func NomeSistemaOperacional() string {
	return runtime.GOOS
}

// ArquiteturaProcessador retorna a arquitetura do processador.
// Valores comuns: "amd64", "arm64", "386".
//
// Exemplo em Verbo:
//
//	O arq é ArquiteturaProcessador de Sistema com ().
func ArquiteturaProcessador() string {
	return runtime.GOARCH
}

// NomeDoComputador retorna o nome do host (hostname) da máquina.
// Lança pânico se não for possível obtê-lo.
//
// Exemplo em Verbo:
//
//	O nome é NomeDoComputador de Sistema com ().
func NomeDoComputador() string {
	nome, err := os.Hostname()
	if err != nil {
		panic("Erro ao obter nome do computador: " + err.Error())
	}
	return nome
}

// PID retorna o identificador do processo atual (Process ID).
//
// Exemplo em Verbo:
//
//	O pid é PID de Sistema com ().
func PID() int {
	return os.Getpid()
}

// NumeroDeCPUs retorna o número de CPUs lógicas disponíveis para o processo.
//
// Exemplo em Verbo:
//
//	O cpus é NumeroDeCPUs de Sistema com ().
func NumeroDeCPUs() int {
	return runtime.NumCPU()
}

// -----------------------------------------------------------------------------
// Controle do Processo
// -----------------------------------------------------------------------------

// Sair encerra o processo atual com código de saída 0 (sucesso).
//
// Exemplo em Verbo:
//
//	Sair de Sistema com ().
func Sair() {
	os.Exit(0)
}

// SairCom encerra o processo atual com o código de saída fornecido.
// Convenção: 0 = sucesso, qualquer outro valor = erro.
//
// Exemplo em Verbo:
//
//	SairCom de Sistema com (1).
func SairCom(codigo int) {
	os.Exit(codigo)
}
