import * as vscode from 'vscode';

interface HoverInfo {
  titulo: string;
  descricao: string;
  exemplo?: string;
  detalhes?: string;
}

const DOCUMENTACAO_VERBO: Record<string, HoverInfo> = {
  // Artigos
  'o': {
    titulo: 'Artigo Definido (Constante)',
    descricao: 'Declara uma constante **imutável** com concordância masculina.',
    exemplo: 'O limite é 100.'
  },
  'a': {
    titulo: 'Artigo Definido (Constante)',
    descricao: 'Declara uma constante **imutável** com concordância feminina.',
    exemplo: 'A mensagem é "Olá, Mundo!".'
  },
  'os': {
    titulo: 'Artigo Definido Plural (Constante)',
    descricao: 'Declara valores **imutáveis** no plural.',
    exemplo: 'Os dados são [1, 2, 3].'
  },
  'as': {
    titulo: 'Artigo Definido Plural (Constante)',
    descricao: 'Declara valores **imutáveis** no plural feminino.',
    exemplo: 'As chaves são ["a", "b"].'
  },
  'um': {
    titulo: 'Artigo Indefinido (Variável)',
    descricao: 'Declara uma variável **mutável** com concordância masculina.',
    exemplo: 'Um contador está 0.'
  },
  'uma': {
    titulo: 'Artigo Indefinido (Variável)',
    descricao: 'Declara uma variável **mutável** com concordância feminina.',
    exemplo: 'Uma taxa está 0.15.'
  },

  // Verbos de Atribuição
  'é': {
    titulo: 'Atribuição Estática / Essência',
    descricao: 'Atribui valor estático ou definição imutável a um identificador.',
    exemplo: 'O pi é 3.14159.'
  },
  'está': {
    titulo: 'Atribuição de Estado / Mutabilidade',
    descricao: 'Atribui um estado temporário ou valor inicial mutável a uma variável.',
    exemplo: 'contador está contador + 1.'
  },

  // Funções e E/S
  'para': {
    titulo: 'Declaração de Função',
    descricao: 'Inicia a declaração de uma função. No Verbo, toda função é nomeada com um verbo no infinitivo.',
    exemplo: 'Para Calcular usando (x: Inteiro):\n    Retorne x * 2.\n.'
  },
  'usando': {
    titulo: 'Parâmetros de Função',
    descricao: 'Introduz a lista de parâmetros tipados de uma função.',
    exemplo: 'Para Saudar usando (nome: Texto):'
  },
  'com': {
    titulo: 'Passagem de Argumentos',
    descricao: 'Conector para passagem de argumentos em chamadas de função ou instâncias.',
    exemplo: 'Saudar com ("Brasil").'
  },
  'retorne': {
    titulo: 'Retorno de Função',
    descricao: 'Encerra a execução da função retornando um valor ou `Nulo`.',
    exemplo: 'Retorne total.'
  },
  'exibir': {
    titulo: 'Saída Padrão (E/S)',
    descricao: 'Imprime a expressão na saída padrão (terminal).',
    exemplo: 'Exibir com ("Olá, Mundo!").'
  },

  // Controle de fluxo
  'se': {
    titulo: 'Estrutura Condicional (Se)',
    descricao: 'Avalia uma condição booleana. Se verdadeira, executa o bloco.',
    exemplo: 'Se idade for maior que 18, então:\n    Exibir com ("Maior de idade").\n.'
  },
  'senão': {
    titulo: 'Alternativa Condicional (Senão)',
    descricao: 'Executa um bloco de código alternativo caso a condição do `Se` seja falsa.',
    exemplo: 'Senão:\n    Exibir com ("Menor de idade").\n.'
  },
  'então': {
    titulo: 'Conector de Condição',
    descricao: 'Marca a conclusão da expressão condicional e o início do bloco.',
    exemplo: 'Se ativo for igual Verdadeiro, então:'
  },
  'for': {
    titulo: 'Subjuntivo Condicional',
    descricao: 'Conector subjuntivo utilizado em comparações lógicas no `Se` e `Enquanto`.',
    exemplo: 'Se a soma for maior que 10, então:'
  },
  'repita': {
    titulo: 'Laço de Repetição',
    descricao: 'Executa um bloco de código repetidamente por contagem (`N vezes`) ou por iteração (`para cada item em lista`).',
    exemplo: 'Repita 5 vezes:\n    Exibir com ("Executando").\n.'
  },
  'vezes': {
    titulo: 'Contador de Repetição',
    descricao: 'Indica a quantidade de iterações em um loop `Repita N vezes:`.',
    exemplo: 'Repita 10 vezes:'
  },
  'cada': {
    titulo: 'Iteração sobre Elementos',
    descricao: 'Usado em conjunto com `Repita para cada elemento em coleção:`.',
    exemplo: 'Repita para cada item em lista:'
  },
  'em': {
    titulo: 'Conector de Coleção',
    descricao: 'Indica a lista ou coleção sobre a qual iterar.',
    exemplo: 'Repita para cada num em numeros:'
  },
  'enquanto': {
    titulo: 'Laço Condicional (Enquanto)',
    descricao: 'Executa o bloco repetidamente enquanto a condição for verdadeira.',
    exemplo: 'Enquanto i for menor que 10:\n    i está i + 1.\n.'
  },

  // Tipos
  'texto': {
    titulo: 'Tipo Texto (String)',
    descricao: 'Cadeia de caracteres delimitada por aspas duplas.',
    exemplo: 'A saudacao é "Bem-vindo".'
  },
  'inteiro': {
    titulo: 'Tipo Inteiro (Int)',
    descricao: 'Número inteiro positivo ou negativo (64-bit).',
    exemplo: 'Um total está 42.'
  },
  'decimal': {
    titulo: 'Tipo Decimal (Float)',
    descricao: 'Número de ponto flutuante com casas decimais.',
    exemplo: 'Uma taxa está 3.14.'
  },
  'lógico': {
    titulo: 'Tipo Lógico (Boolean)',
    descricao: 'Valor booleano: `Verdadeiro` ou `Falso`.',
    exemplo: 'Um ativo está Verdadeiro.'
  },
  'logico': {
    titulo: 'Tipo Lógico (Boolean)',
    descricao: 'Valor booleano: `Verdadeiro` ou `Falso` (sem acento).',
    exemplo: 'Um ativo está Verdadeiro.'
  },
  'lista': {
    titulo: 'Tipo Lista (Slice)',
    descricao: 'Coleção ordenada de elementos entre colchetes.',
    exemplo: 'Uma notas está [8.5, 9.0, 10.0].'
  },

  // Literais
  'verdadeiro': {
    titulo: 'Literal Booleano: Verdadeiro (true)',
    descricao: 'Representa a verdade lógica.',
    exemplo: 'O sucesso é Verdadeiro.'
  },
  'falso': {
    titulo: 'Literal Booleano: Falso (false)',
    descricao: 'Representa a falsidade lógica.',
    exemplo: 'O erro é Falso.'
  },
  'nulo': {
    titulo: 'Ausência de Valor (null/nil)',
    descricao: 'Representa a inexistência ou nulidade de um valor.',
    exemplo: 'Retorne Nulo.'
  },

  // V2 Concorrência e Estruturas
  'entidade': {
    titulo: 'Entidade (Estrutura / Registro)',
    descricao: 'Declara uma estrutura de dados com campos tipados nomeados.',
    exemplo: 'A entidade Usuario contendo (nome: Texto, idade: Inteiro).'
  },
  'contendo': {
    titulo: 'Campos de Entidade',
    descricao: 'Especifica a lista de atributos de uma `Entidade`.',
    exemplo: 'A entidade Ponto contendo (x: Inteiro, y: Inteiro).'
  },
  'novo': {
    titulo: 'Instanciação de Entidade',
    descricao: 'Cria uma nova instância de uma entidade preenchendo seus valores.',
    exemplo: 'O usuario é um novo Usuario contendo ("Maria", 30).'
  },
  'simultaneamente': {
    titulo: 'Execução Concorrente',
    descricao: 'Inicia a execução paralela das instruções do bloco (goroutines do Go).',
    exemplo: 'Simultaneamente:\n    Tarefa1 com ().\n    Tarefa2 com ().\n.'
  },
  'canal': {
    titulo: 'Canal de Concorrência',
    descricao: 'Cria um canal tipado para comunicação segura entre tarefas concorrentes.',
    exemplo: 'Uma via é um Canal de Inteiros.'
  },
  'enviar': {
    titulo: 'Enviar Dado ao Canal',
    descricao: 'Transmite um dado através de um canal.',
    exemplo: 'Enviar 42 para via.'
  },
  'receber': {
    titulo: 'Receber Dado do Canal',
    descricao: 'Lê um dado vindo de um canal concorrente.',
    exemplo: 'Um valor é Receber de via.'
  },
  'tente': {
    titulo: 'Tratamento de Exceção (Try)',
    descricao: 'Inicia um bloco protegido contra falhas em tempo de execução.',
    exemplo: 'Tente:\n    Operacao com ().\nCapture erro:\n    Exibir com (erro).\n.'
  },
  'capture': {
    titulo: 'Captura de Erro (Catch)',
    descricao: 'Recebe o erro capturado durante a execução do bloco `Tente`.',
    exemplo: 'Capture erro:\n    Exibir com (erro).'
  },
  'sinalize': {
    titulo: 'Sinalizar Pânico / Erro Crítico',
    descricao: 'Interrompe a execução e lança um erro crítico.',
    exemplo: 'Sinalize "Saldo insuficiente".'
  },
  'incluir': {
    titulo: 'Importação de Biblioteca',
    descricao: 'Importa um módulo padrão (`Texto`, `Matematica`, `Arquivo`, `Html`).',
    exemplo: 'Incluir Matematica.'
  },

  // V3 Servidor Web
  'servidor': {
    titulo: 'Servidor Web (Estilo Flask)',
    descricao: 'Configura ou manipula um servidor web HTTP nativo do Verbo.',
    exemplo: 'Um app está Servidor com (local, 5000).'
  },
  'rota': {
    titulo: 'Definição de Rota HTTP',
    descricao: 'Vincula um verbo HTTP (`GET`, `POST`, `PUT`, `DELETE`) e um caminho a um bloco de código.',
    exemplo: 'app rota GET em "/api":\n    Retorne "Sucesso".\n.'
  },
  'iniciar': {
    titulo: 'Inicialização do Servidor',
    descricao: 'Coloca o servidor web em execução ouvindo na porta configurada.',
    exemplo: 'app iniciar.'
  },
  'rodar': {
    titulo: 'Inicialização do Servidor (Alias)',
    descricao: 'Açúcar sintático para `iniciar`.',
    exemplo: 'app rodar.'
  },

  // Biblioteca Padrão (BibVerbo)
  'tamanho': {
    titulo: 'Texto.Tamanho(s: Texto) -> Inteiro',
    descricao: 'Retorna a quantidade de caracteres da string.',
    exemplo: 'Um tam é Tamanho de Texto com ("Olá").'
  },
  'maiusculas': {
    titulo: 'Texto.Maiusculas(s: Texto) -> Texto',
    descricao: 'Converte todas as letras da string para maiúsculas.',
    exemplo: 'Um txt é Maiusculas de Texto com ("brasil").'
  },
  'minusculas': {
    titulo: 'Texto.Minusculas(s: Texto) -> Texto',
    descricao: 'Converte todas as letras da string para minúsculas.',
    exemplo: 'Um txt é Minusculas de Texto com ("BRASIL").'
  },
  'contem': {
    titulo: 'Texto.Contem(s: Texto, busca: Texto) -> Lógico',
    descricao: 'Verifica se a string contém a substring informada.',
    exemplo: 'Um achou é Contem de Texto com ("banana", "na").'
  },
  'substituir': {
    titulo: 'Texto.Substituir(s, antigo, novo: Texto) -> Texto',
    descricao: 'Substitui todas as ocorrências de um termo por outro.',
    exemplo: 'Um res é Substituir de Texto com ("olá mundo", "mundo", "amigo").'
  },
  'absoluto': {
    titulo: 'Matematica.Absoluto(x: Decimal) -> Decimal',
    descricao: 'Retorna o valor absoluto positivo de um número.',
    exemplo: 'Um pos é Absoluto de Matematica com (-42.5).'
  },
  'potencia': {
    titulo: 'Matematica.Potencia(base, expoente: Decimal) -> Decimal',
    descricao: 'Eleva a base à potência do expoente.',
    exemplo: 'Um res é Potencia de Matematica com (2.0, 8.0).'
  },
  'raiz': {
    titulo: 'Matematica.Raiz(x: Decimal) -> Decimal',
    descricao: 'Calcula a raiz quadrada de um número.',
    exemplo: 'Um r é Raiz de Matematica com (9.0).'
  },
  'lertexto': {
    titulo: 'Arquivo.LerTexto(caminho: Texto) -> Texto',
    descricao: 'Lê o conteúdo completo de um arquivo de texto no disco.',
    exemplo: 'O dados é LerTexto de Arquivo com ("config.txt").'
  },
  'escrevertexto': {
    titulo: 'Arquivo.EscreverTexto(caminho, conteudo: Texto)',
    descricao: 'Grava conteúdo textual em um arquivo no disco.',
    exemplo: 'EscreverTexto de Arquivo com ("saida.txt", "OK").'
  },
  'criarpagina': {
    titulo: 'Html.CriarPagina(titulo, corpo: Texto) -> Texto',
    descricao: 'Gera uma página HTML completa com doctype, charset UTF-8 e viewport.',
    exemplo: 'O doc é CriarPagina de Html com ("Meu Site", "<h1>Olá!</h1>").'
  },
  'criarelemento': {
    titulo: 'Html.CriarElemento(tag, conteudo: Texto) -> Texto',
    descricao: 'Gera uma tag HTML delimitando o conteúdo.',
    exemplo: 'O titulo é CriarElemento de Html com ("h1", "Bem-vindo").'
  }
};

export class VerboHoverProvider implements vscode.HoverProvider {
  public provideHover(
    document: vscode.TextDocument,
    position: vscode.Position
  ): vscode.ProviderResult<vscode.Hover> {
    const range = document.getWordRangeAtPosition(position, /[a-zA-ZáàâãéêíóôõúçÁÀÂÃÉÊÍÓÔÕÚÇ_]+/);
    if (!range) {
      return null;
    }

    const palavra = document.getText(range).toLowerCase();
    const info = DOCUMENTACAO_VERBO[palavra];

    if (!info) {
      return null;
    }

    const md = new vscode.MarkdownString();
    md.isTrusted = true;
    md.appendMarkdown(`### 🇧🇷 Verbo: **${info.titulo}**\n\n`);
    md.appendMarkdown(`${info.descricao}\n\n`);

    if (info.exemplo) {
      md.appendMarkdown(`**Exemplo:**\n\`\`\`verbo\n${info.exemplo}\n\`\`\`\n`);
    }

    return new vscode.Hover(md, range);
  }
}

