import * as vscode from 'vscode';

export class VerboCompletionProvider implements vscode.CompletionItemProvider {
  public provideCompletionItems(
    document: vscode.TextDocument,
    position: vscode.Position
  ): vscode.ProviderResult<vscode.CompletionItem[] | vscode.CompletionList> {
    const items: vscode.CompletionItem[] = [];

    // Artigos (Mutabilidade)
    items.push(
      this.criarItem('O', vscode.CompletionItemKind.Keyword, 'Artigo definido masculino (constante imutável)', 'O ${1:nome} é ${2:valor}.'),
      this.criarItem('A', vscode.CompletionItemKind.Keyword, 'Artigo definido feminino (constante imutável)', 'A ${1:nome} é ${2:valor}.'),
      this.criarItem('Um', vscode.CompletionItemKind.Keyword, 'Artigo indefinido masculino (variável mutável)', 'Um ${1:nome} está ${2:valor}.'),
      this.criarItem('Uma', vscode.CompletionItemKind.Keyword, 'Artigo indefinido feminino (variável mutável)', 'Uma ${1:nome} está ${2:valor}.')
    );

    // Palavras-chave de controle
    items.push(
      this.criarItem('Para', vscode.CompletionItemKind.Keyword, 'Declaração de função', 'Para ${1:Nome} usando (${2:param}: ${3:Texto}):\n\t$0\n.'),
      this.criarItem('usando', vscode.CompletionItemKind.Keyword, 'Parâmetros de função', 'usando (${1:param}: ${2:Texto})'),
      this.criarItem('com', vscode.CompletionItemKind.Keyword, 'Passagem de argumentos', 'com (${1:args})'),
      this.criarItem('Retorne', vscode.CompletionItemKind.Keyword, 'Retorno de função', 'Retorne ${1:valor}.'),
      this.criarItem('Exibir', vscode.CompletionItemKind.Keyword, 'Saída padrão', 'Exibir com (${1:expressao}).'),
      this.criarItem('Se', vscode.CompletionItemKind.Keyword, 'Estrutura condicional', 'Se ${1:condicao} for ${2|igual,menor que,maior que|} ${3:valor}, então:\n\t$0\n.'),
      this.criarItem('Senão', vscode.CompletionItemKind.Keyword, 'Alternativa condicional', 'Senão:\n\t$0\n.'),
      this.criarItem('Repita', vscode.CompletionItemKind.Keyword, 'Repetição N vezes', 'Repita ${1:10} vezes:\n\t$0\n.'),
      this.criarItem('Repita para cada', vscode.CompletionItemKind.Keyword, 'Repetição sobre lista', 'Repita para cada ${1:item} em ${2:lista}:\n\t$0\n.'),
      this.criarItem('Enquanto', vscode.CompletionItemKind.Keyword, 'Loop condicional', 'Enquanto ${1:condicao}:\n\t$0\n.'),
      this.criarItem('Dado', vscode.CompletionItemKind.Keyword, 'Cláusula de guarda / premissa', 'Dado ${1:condicao}.'),
      this.criarItem('Incluir', vscode.CompletionItemKind.Keyword, 'Importar pacote padrão', 'Incluir ${1|Texto,Matematica,Arquivo,Html|}.')
    );

    // V2 Concorrência e Estruturas
    items.push(
      this.criarItem('Entidade', vscode.CompletionItemKind.Class, 'Declaração de estrutura de dados', 'A entidade ${1:Nome} contendo (${2:campo}: ${3:Texto}).'),
      this.criarItem('novo', vscode.CompletionItemKind.Keyword, 'Instanciação de entidade', 'um novo ${1:Entidade} contendo (${2:valores})'),
      this.criarItem('Simultaneamente', vscode.CompletionItemKind.Keyword, 'Execução paralela', 'Simultaneamente:\n\t$0\n.'),
      this.criarItem('Aguarde', vscode.CompletionItemKind.Keyword, 'Sincronização', 'Aguarde.'),
      this.criarItem('Canal', vscode.CompletionItemKind.TypeParameter, 'Canal concorrente', 'Canal de ${1|Texto,Inteiro,Decimal|}'),
      this.criarItem('Enviar', vscode.CompletionItemKind.Keyword, 'Enviar dado para canal', 'Enviar ${1:dado} para ${2:canal}.'),
      this.criarItem('Receber', vscode.CompletionItemKind.Keyword, 'Receber dado de canal', 'Receber de ${1:canal}'),
      this.criarItem('Tente', vscode.CompletionItemKind.Keyword, 'Tratamento de exceções', 'Tente:\n\t$1\nCapture ${2:erro}:\n\t$0\n.'),
      this.criarItem('Sinalize', vscode.CompletionItemKind.Keyword, 'Lançar pânico / erro crítico', 'Sinalize "${1:mensagem}".')
    );

    // V3 Servidor Web
    items.push(
      this.criarItem('Servidor', vscode.CompletionItemKind.Class, 'Configuração do Servidor Web', 'Um ${1:app} está Servidor com (local, ${2:5000}).'),
      this.criarItem('rota', vscode.CompletionItemKind.Function, 'Rota HTTP', '${1:app} rota ${2|GET,POST,PUT,DELETE|} em "${3:/}":\n\t$0\n.'),
      this.criarItem('iniciar', vscode.CompletionItemKind.Method, 'Iniciar servidor web', '${1:app} iniciar.'),
      this.criarItem('rodar', vscode.CompletionItemKind.Method, 'Rodar servidor web', '${1:app} rodar.')
    );

    // Tipos primitivos
    const tipos = ['Texto', 'Inteiro', 'Decimal', 'Lógico', 'Lista'];
    for (const t of tipos) {
      const item = new vscode.CompletionItem(t, vscode.CompletionItemKind.TypeParameter);
      item.detail = `Tipo primitivo ${t}`;
      items.push(item);
    }

    // Literais
    items.push(
      this.criarItem('Verdadeiro', vscode.CompletionItemKind.Value, 'Valor booleano verdadeiro'),
      this.criarItem('Falso', vscode.CompletionItemKind.Value, 'Valor booleano falso'),
      this.criarItem('Nulo', vscode.CompletionItemKind.Value, 'Valor nulo / ausência de valor')
    );

    // Biblioteca Padrão
    this.adicionarStdlib(items);

    return items;
  }

  private adicionarStdlib(items: vscode.CompletionItem[]) {
    // Módulo Texto
    items.push(
      this.criarItem('Tamanho de Texto', vscode.CompletionItemKind.Function, 'Tamanho(s: Texto) -> Inteiro', 'Tamanho de Texto com (${1:string})'),
      this.criarItem('Maiusculas de Texto', vscode.CompletionItemKind.Function, 'Maiusculas(s: Texto) -> Texto', 'Maiusculas de Texto com (${1:string})'),
      this.criarItem('Minusculas de Texto', vscode.CompletionItemKind.Function, 'Minusculas(s: Texto) -> Texto', 'Minusculas de Texto com (${1:string})'),
      this.criarItem('Contem de Texto', vscode.CompletionItemKind.Function, 'Contem(s, substr: Texto) -> Lógico', 'Contem de Texto com (${1:string}, ${2:busca})'),
      this.criarItem('Substituir de Texto', vscode.CompletionItemKind.Function, 'Substituir(s, velho, novo: Texto) -> Texto', 'Substituir de Texto com (${1:string}, ${2:antigo}, ${3:novo})'),
      this.criarItem('Dividir de Texto', vscode.CompletionItemKind.Function, 'Dividir(s, sep: Texto) -> Lista', 'Dividir de Texto com (${1:string}, ${2:separador})')
    );

    // Módulo Matematica
    items.push(
      this.criarItem('Absoluto de Matematica', vscode.CompletionItemKind.Function, 'Absoluto(x: Decimal) -> Decimal', 'Absoluto de Matematica com (${1:numero})'),
      this.criarItem('Teto de Matematica', vscode.CompletionItemKind.Function, 'Teto(x: Decimal) -> Decimal', 'Teto de Matematica com (${1:numero})'),
      this.criarItem('Piso de Matematica', vscode.CompletionItemKind.Function, 'Piso(x: Decimal) -> Decimal', 'Piso de Matematica com (${1:numero})'),
      this.criarItem('Maximo de Matematica', vscode.CompletionItemKind.Function, 'Maximo(a, b: Decimal) -> Decimal', 'Maximo de Matematica com (${1:a}, ${2:b})'),
      this.criarItem('Minimo de Matematica', vscode.CompletionItemKind.Function, 'Minimo(a, b: Decimal) -> Decimal', 'Minimo de Matematica com (${1:a}, ${2:b})'),
      this.criarItem('Potencia de Matematica', vscode.CompletionItemKind.Function, 'Potencia(base, exp: Decimal) -> Decimal', 'Potencia de Matematica com (${1:base}, ${2:expoente})'),
      this.criarItem('Raiz de Matematica', vscode.CompletionItemKind.Function, 'Raiz(x: Decimal) -> Decimal', 'Raiz de Matematica com (${1:numero})')
    );

    // Módulo Arquivo
    items.push(
      this.criarItem('LerTexto de Arquivo', vscode.CompletionItemKind.Function, 'LerTexto(caminho: Texto) -> Texto', 'LerTexto de Arquivo com (${1:caminho})'),
      this.criarItem('EscreverTexto de Arquivo', vscode.CompletionItemKind.Function, 'EscreverTexto(caminho, conteudo: Texto)', 'EscreverTexto de Arquivo com (${1:caminho}, ${2:conteudo})')
    );

    // Módulo Html
    items.push(
      this.criarItem('CriarPagina de Html', vscode.CompletionItemKind.Function, 'CriarPagina(titulo, corpo: Texto) -> Texto', 'CriarPagina de Html com (${1:titulo}, ${2:corpo})'),
      this.criarItem('CriarElemento de Html', vscode.CompletionItemKind.Function, 'CriarElemento(tag, conteudo: Texto) -> Texto', 'CriarElemento de Html com (${1:tag}, ${2:conteudo})'),
      this.criarItem('CriarElementoComAtributos de Html', vscode.CompletionItemKind.Function, 'CriarElementoComAtributos(tag, attrs, conteudo: Texto) -> Texto', 'CriarElementoComAtributos de Html com (${1:tag}, ${2:atributos}, ${3:conteudo})')
    );
  }

  private criarItem(
    label: string,
    kind: vscode.CompletionItemKind,
    detail: string,
    snippet?: string
  ): vscode.CompletionItem {
    const item = new vscode.CompletionItem(label, kind);
    item.detail = detail;
    if (snippet) {
      item.insertText = new vscode.SnippetString(snippet);
    }
    return item;
  }
}

