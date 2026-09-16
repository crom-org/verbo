package sistema

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestArgumentos(t *testing.T) {
	// Em contexto de teste, os.Args é preenchido pelo framework de testes.
	// Apenas verificamos que as funções executam sem pânico e retornam tipos corretos.
	args := Argumentos()
	if args == nil {
		t.Fatal("Argumentos() não deve retornar nil")
	}

	n := NumeroDeArgumentos()
	if n < 0 {
		t.Fatalf("NumeroDeArgumentos() deve ser >= 0, obteve %d", n)
	}

	if n != len(args) {
		t.Fatalf("NumeroDeArgumentos() (%d) != len(Argumentos()) (%d)", n, len(args))
	}
}

func TestArgumento(t *testing.T) {
	// Garantir pânico ao acessar índice inválido
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("esperava pânico ao acessar Argumento com índice inválido")
		}
	}()
	// os.Args[1:] durante testes tem ao menos alguns flags do go test,
	// então usamos um índice muito grande para forçar pânico
	Argumento(99999)
}

func TestVariaveisDeAmbiente(t *testing.T) {
	const nome = "VERBO_TESTE_VAR"
	const valor = "Verbo123"

	// Definir variável
	DefinirVariavelDeAmbiente(nome, valor)

	// Ler variável
	lido := VariavelDeAmbiente(nome)
	if lido != valor {
		t.Fatalf("esperava %q para variável %s, obteve %q", valor, nome, lido)
	}

	// Variável deve aparecer em VariaveisDeAmbiente
	todas := VariaveisDeAmbiente()
	encontrou := false
	prefixo := nome + "=" + valor
	for _, v := range todas {
		if s, ok := v.(string); ok && strings.HasPrefix(s, prefixo) {
			encontrou = true
			break
		}
	}
	if !encontrou {
		t.Fatalf("variável %s=%s não encontrada em VariaveisDeAmbiente()", nome, valor)
	}

	// Remover variável
	RemoverVariavelDeAmbiente(nome)
	depois := VariavelDeAmbiente(nome)
	if depois != "" {
		t.Fatalf("esperava string vazia após remover variável %s, obteve %q", nome, depois)
	}
}

func TestExecutarComando(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Pulando teste de execução de comando no Windows")
	}
	saida := ExecutarComando("echo VerboTest")
	if !strings.Contains(saida, "VerboTest") {
		t.Fatalf("esperava 'VerboTest' na saída do echo, obteve %q", saida)
	}
}

func TestExecutarComandoComArgumentos(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Pulando teste de execução de comando no Windows")
	}
	saida := ExecutarComandoComArgumentos("echo", []interface{}{"OláVerbo"})
	if !strings.Contains(saida, "OláVerbo") {
		t.Fatalf("esperava 'OláVerbo' na saída, obteve %q", saida)
	}
}

func TestExecutarComandoComEntrada(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Pulando teste de execução de comando no Windows")
	}
	saida := ExecutarComandoComEntrada("cat", "texto de entrada")
	if saida != "texto de entrada" {
		t.Fatalf("esperava 'texto de entrada', obteve %q", saida)
	}
}

func TestDiretorio(t *testing.T) {
	dir := DiretorioAtual()
	if dir == "" {
		t.Fatal("DiretorioAtual() retornou string vazia")
	}

	// Mudar para um diretório temporário e voltar
	original := dir
	temp := os.TempDir()
	MudarDiretorio(temp)

	novoDir := DiretorioAtual()
	if novoDir == "" {
		t.Fatal("DiretorioAtual() retornou vazio após MudarDiretorio()")
	}

	// Voltar ao original
	MudarDiretorio(original)
}

func TestInformacoesSistema(t *testing.T) {
	so := NomeSistemaOperacional()
	if so == "" {
		t.Fatal("NomeSistemaOperacional() retornou string vazia")
	}

	arq := ArquiteturaProcessador()
	if arq == "" {
		t.Fatal("ArquiteturaProcessador() retornou string vazia")
	}

	hostname := NomeDoComputador()
	if hostname == "" {
		t.Fatal("NomeDoComputador() retornou string vazia")
	}

	pid := PID()
	if pid <= 0 {
		t.Fatalf("PID() deve ser > 0, obteve %d", pid)
	}

	cpus := NumeroDeCPUs()
	if cpus <= 0 {
		t.Fatalf("NumeroDeCPUs() deve ser >= 1, obteve %d", cpus)
	}
}
