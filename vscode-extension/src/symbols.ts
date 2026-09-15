import * as vscode from 'vscode';

export class VerboDocumentSymbolProvider implements vscode.DocumentSymbolProvider {
  public provideDocumentSymbols(
    document: vscode.TextDocument
  ): vscode.ProviderResult<vscode.DocumentSymbol[]> {
    const symbols: vscode.DocumentSymbol[] = [];
    const lineCount = document.lineCount;

    for (let i = 0; i < lineCount; i++) {
      const line = document.lineAt(i);
      const text = line.text.trim();

      if (text.startsWith('//') || text === '') {
        continue;
      }

      // 1. Declaração de Função: Para <Nome> usando ...: ou Para <Nome>:
      const matchFuncao = /^Para\s+([A-Za-zÀ-ÖØ-öø-ÿ_][A-Za-zÀ-ÖØ-öø-ÿ0-9_]*)/i.exec(text);
      if (matchFuncao) {
        const nome = matchFuncao[1];
        // Determinar o fim do bloco (procura pelo ponto de fechamento correspondente ou fim do arquivo)
        const fimLinha = this.encontrarFimDoBloco(document, i);
        const range = new vscode.Range(i, 0, fimLinha, document.lineAt(fimLinha).text.length);
        const selectionRange = new vscode.Range(i, line.text.indexOf(nome), i, line.text.indexOf(nome) + nome.length);

        const sym = new vscode.DocumentSymbol(
          nome,
          'Função',
          vscode.SymbolKind.Function,
          range,
          selectionRange
        );
        symbols.push(sym);
        continue;
      }

      // 2. Declaração de Entidade: [A|O] entidade <Nome> contendo ...
      const matchEntidade = /^(?:A|O|Uma|Um)?\s*entidade\s+([A-Za-zÀ-ÖØ-öø-ÿ_][A-Za-zÀ-ÖØ-öø-ÿ0-9_]*)/i.exec(text);
      if (matchEntidade) {
        const nome = matchEntidade[1];
        const range = line.range;
        const col = line.text.indexOf(nome);
        const selectionRange = new vscode.Range(i, col, i, col + nome.length);

        const sym = new vscode.DocumentSymbol(
          nome,
          'Entidade (Struct)',
          vscode.SymbolKind.Struct,
          range,
          selectionRange
        );
        symbols.push(sym);
        continue;
      }

      // 3. Rota Web: <app> rota GET/POST/PUT/DELETE em "<caminho>":
      const matchRota = /([A-Za-zÀ-ÖØ-öø-ÿ_][A-Za-zÀ-ÖØ-öø-ÿ0-9_]*)\s+rota\s+(GET|POST|PUT|DELETE)\s+em\s+("[^"]+")/i.exec(text);
      if (matchRota) {
        const app = matchRota[1];
        const metodo = matchRota[2];
        const caminho = matchRota[3];
        const nome = `${metodo} ${caminho} (${app})`;
        const fimLinha = this.encontrarFimDoBloco(document, i);
        const range = new vscode.Range(i, 0, fimLinha, document.lineAt(fimLinha).text.length);
        const selectionRange = line.range;

        const sym = new vscode.DocumentSymbol(
          nome,
          'Rota HTTP',
          vscode.SymbolKind.Method,
          range,
          selectionRange
        );
        symbols.push(sym);
        continue;
      }

      // 4. Declaração de Constante: O/A <nome> é ...
      const matchConstante = /^(?:O|A)\s+([A-Za-zÀ-ÖØ-öø-ÿ_][A-Za-zÀ-ÖØ-öø-ÿ0-9_]*)\s+é\b/i.exec(text);
      if (matchConstante) {
        const nome = matchConstante[1];
        const col = line.text.indexOf(nome);
        const sym = new vscode.DocumentSymbol(
          nome,
          'Constante imutável',
          vscode.SymbolKind.Constant,
          line.range,
          new vscode.Range(i, col, i, col + nome.length)
        );
        symbols.push(sym);
        continue;
      }

      // 5. Declaração de Variável: Um/Uma <nome> está ...
      const matchVariavel = /^(?:Um|Uma)\s+([A-Za-zÀ-ÖØ-öø-ÿ_][A-Za-zÀ-ÖØ-öø-ÿ0-9_]*)\s+está\b/i.exec(text);
      if (matchVariavel) {
        const nome = matchVariavel[1];
        const col = line.text.indexOf(nome);
        const sym = new vscode.DocumentSymbol(
          nome,
          'Variável mutável',
          vscode.SymbolKind.Variable,
          line.range,
          new vscode.Range(i, col, i, col + nome.length)
        );
        symbols.push(sym);
        continue;
      }
    }

    return symbols;
  }

  private encontrarFimDoBloco(document: vscode.TextDocument, linhaInicio: number): number {
    let profundidade = 0;
    for (let i = linhaInicio; i < document.lineCount; i++) {
      const t = document.lineAt(i).text.trim();
      if (t.endsWith(':')) {
        profundidade++;
      }
      if (t === '.' || t.startsWith('. ') || t.startsWith('.//')) {
        profundidade--;
        if (profundidade <= 0) {
          return i;
        }
      }
    }
    return Math.min(linhaInicio + 20, document.lineCount - 1);
  }
}

