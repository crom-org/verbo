# 🇧🇷 Linguagem Verbo — Documentação Oficial

Bem-vindo à documentação oficial da **Linguagem Verbo**, a linguagem de programação moderna transpilada para Go baseada na gramática da norma culta do Português Brasileiro.

---

## O que é o Verbo?

O **Verbo** foi concebido com uma premissa ousada: transformar a gramática da língua portuguesa em regras semânticas determinísticas para um compilador de alto desempenho.

Em vez de palavras-chave emprestadas do inglês (`let`, `var`, `const`, `function`, `class`), o Verbo utiliza a riqueza sintática do português:

- **Artigos definidos e indefinidos** determinam imutabilidade (`O total é 100.` vs `Um contador está 0.`).
- **Verbos de estado e essência** diferenciam atribuição estática e reatribuição (`é` para constante, `está` para mutável).
- **Preposições e conectivos gramaticais** conectam argumentos (`de`, `com`, `para`, `em`, `por`).
- **Pontuação natural** encerra instruções com o ponto final (`.`).

O código Verbo é transpilado diretamente para Go nativo, combinando **elegância de leitura** com **alta performance de execução**.

---

## Exemplo Rápido

```verbo
// Importação de módulos da Biblioteca Padrão
Incluir Internet.
Incluir Criptografia.

// Definição de Estrutura de Dados
A entidade Usuario contendo (Nome: Texto, Email: Texto, Nivel: Inteiro).

// Função com parâmetros tipados
Para BoasVindas usando (usuario: Usuario) :
    O nome_formatado é Nome de usuario.
    Exibir "Bem-vindo(a), " + nome_formatado + "!".
.

// Bloco principal
O admin é um novo Usuario contendo ("Ada Lovelace", "ada@exemplo.com", 1).
BoasVindas com (admin).

// Criptografia e identificadores seguros
O token_sessao é GerarUUID de Criptografia com ().
Exibir "Token da sessão: " + token_sessao.
```

---

## Navegação Rápida

<table style="width:100%">
  <tr>
    <td width="50%">
      <h3>🚀 <a href="comecando/introducao.md">Começando</a></h3>
      <p>Aprenda a instalar o Verbo, rodar seu primeiro programa e entender a filosofia da linguagem.</p>
      <ul>
        <li><a href="comecando/instalacao.md">Instalação e CLI</a></li>
        <li><a href="comecando/primeiro-programa.md">Primeiro Programa</a></li>
      </ul>
    </td>
    <td width="50%">
      <h3>📘 <a href="linguagem/variaveis-e-constantes.md">Guia da Linguagem</a></h3>
      <p>Domine a sintaxe do Verbo 2.0, variáveis, funções, entidades, concorrência e tratamento de erros.</p>
      <ul>
        <li><a href="linguagem/variaveis-e-constantes.md">Ser vs Estar</a></li>
        <li><a href="linguagem/entidades.md">Entidades (Structs)</a></li>
        <li><a href="linguagem/concorrencia.md">Concorrência e Canais</a></li>
      </ul>
    </td>
  </tr>
  <tr>
    <td width="50%">
      <h3>📦 <a href="stdlib/visao-geral.md">BibVerbo (Stdlib)</a></h3>
      <p>Biblioteca padrão oficial do Verbo com baterias inclusas.</p>
      <ul>
        <li><a href="stdlib/internet.md">Internet (HTTP & Sockets)</a></li>
        <li><a href="stdlib/criptografia.md">Criptografia & Aleatório</a></li>
        <li><a href="stdlib/html.md">Geração de HTML</a></li>
        <li><a href="stdlib/texto.md">Manipulação de Texto</a></li>
      </ul>
    </td>
    <td width="50%">
      <h3>🌐 <a href="servidor-web/servidor.md">Servidor Web & Ferramentas</a></h3>
      <p>Crie APIs e servidores estilo Flask diretamente em Verbo.</p>
      <ul>
        <li><a href="servidor-web/servidor.md">Servidor HTTP Nativo</a></li>
        <li><a href="ferramentas/cli.md">Comandos da CLI</a></li>
        <li><a href="ferramentas/gerenciador-de-pacotes.md">Gerenciador de Pacotes</a></li>
        <li><a href="ferramentas/extensao-vscode.md">Extensão para VS Code</a></li>
      </ul>
    </td>
  </tr>
</table>

---

## Filosofia e Princípios de Design

1. **Legibilidade Máxima**: Código deve poder ser lido como um texto técnico formal e claro.
2. **Gramática como Semântica**: Regras gramaticais substituem palavras-chave artificiais.
3. **Determinismo Estrito**: Sem inteligência artificial no parsing; regras de derivação pura via analisador descendente recursivo.
4. **Performance Nativa**: Transpilação limpa para Go, permitindo compilar binários nativos ultravelozes para Linux, macOS e Windows.

