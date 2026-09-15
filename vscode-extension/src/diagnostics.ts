import * as vscode from 'vscode';
import { execFile } from 'child_process';
import * as path from 'path';
import * as fs from 'fs';

export class VerboDiagnostics {
  private diagnosticCollection: vscode.DiagnosticCollection;

  constructor() {
    this.diagnosticCollection = vscode.languages.createDiagnosticCollection('verbo');
  }

  public getCollection(): vscode.DiagnosticCollection {
    return this.diagnosticCollection;
  }

  public dispose() {
    this.diagnosticCollection.dispose();
  }

  public async verificarArquivo(document: vscode.TextDocument): Promise<void> {
    if (document.languageId !== 'verbo' || document.uri.scheme !== 'file') {
      return;
    }

    const config = vscode.workspace.getConfiguration('verbo');
    const cliPath = this.resolverCaminhoCli(document.uri);

    if (cliPath && fs.existsSync(cliPath)) {
      this.executarVerificacaoCli(cliPath, document);
    } else {
      this.executarVerificacaoNativa(document);
    }
  }

  private resolverCaminhoCli(resourceUri: vscode.Uri): string | null {
    const config = vscode.workspace.getConfiguration('verbo');
    const customCli = config.get<string>('cliPath');
    if (customCli && customCli.trim() !== '') {
      return customCli.trim();
    }

    const workspaceFolder = vscode.workspace.getWorkspaceFolder(resourceUri);
    if (workspaceFolder) {
      const candidatoBuild = path.join(workspaceFolder.uri.fsPath, 'build', 'verbo');
      if (fs.existsSync(candidatoBuild)) {
        return candidatoBuild;
      }
    }

    // Procura no PATH
    return 'verbo';
  }

  private executarVerificacaoCli(cliPath: string, document: vscode.TextDocument) {
    const filePath = document.uri.fsPath;
    const cwd = vscode.workspace.getWorkspaceFolder(document.uri)?.uri.fsPath || path.dirname(filePath);

    execFile(cliPath, ['verificar', filePath], { cwd }, (error, stdout, stderr) => {
      const diagnostics: vscode.Diagnostic[] = [];
      const output = (stderr || '') + '\n' + (stdout || '');

      // Padrão de saída do CLI do Verbo:
      // linha 12, coluna 8: esperava ...
      // linha 14: texto não fechado ...
      const regexErro = /(?:linha|Linha)\s+(\d+)(?:,\s*coluna\s+(\d+))?:\s*(.+)/g;
      let match: RegExpExecArray | null;

      while ((match = regexErro.exec(output)) !== null) {
        const linha = Math.max(0, parseInt(match[1], 10) - 1);
        const coluna = match[2] ? Math.max(0, parseInt(match[2], 10) - 1) : 0;
        const mensagem = match[3].trim();

        const docLine = document.lineCount > linha ? document.lineAt(linha) : null;
        let range: vscode.Range;

        if (docLine) {
          const endCol = Math.max(coluna + 1, docLine.text.length);
          range = new vscode.Range(linha, Math.min(coluna, docLine.text.length), linha, endCol);
        } else {
          range = new vscode.Range(linha, coluna, linha, coluna + 1);
        }

        const diagnostic = new vscode.Diagnostic(
          range,
          `Verbo: ${mensagem}`,
          vscode.DiagnosticSeverity.Error
        );
        diagnostic.source = 'verbo';
        diagnostics.push(diagnostic);
      }

      this.diagnosticCollection.set(document.uri, diagnostics);
    });
  }

  // Verificação nativa leve quando o binário 'verbo' ainda não foi compilado
  private executarVerificacaoNativa(document: vscode.TextDocument) {
    const diagnostics: vscode.Diagnostic[] = [];
    const texto = document.getText();
    const linhas = texto.split(/\r?\n/);

    for (let i = 0; i < linhas.length; i++) {
      const linha = linhas[i];
      const linhaSemComentario = linha.replace(/\/\/.*$/, '').trimEnd();

      if (linhaSemComentario === '') {
        continue;
      }

      // Verifica aspas não fechadas
      let inString = false;
      for (let c = 0; c < linhaSemComentario.length; c++) {
        if (linhaSemComentario[c] === '"' && (c === 0 || linhaSemComentario[c - 1] !== '\\')) {
          inString = !inString;
        }
      }
      if (inString) {
        const range = new vscode.Range(i, 0, i, linha.length);
        const diag = new vscode.Diagnostic(
          range,
          'Verbo: Texto não fechado (faltou aspas de fechamento)',
          vscode.DiagnosticSeverity.Error
        );
        diag.source = 'verbo-linter';
        diagnostics.push(diag);
      }
    }

    this.diagnosticCollection.set(document.uri, diagnostics);
  }

  public limpar(uri: vscode.Uri) {
    this.diagnosticCollection.delete(uri);
  }
}

