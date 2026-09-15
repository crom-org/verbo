import * as vscode from 'vscode';
import * as path from 'path';
import * as fs from 'fs';
import { VerboDiagnostics } from './diagnostics';

export class VerboCommands {
  private terminal: vscode.Terminal | null = null;

  constructor(private diagnostics: VerboDiagnostics) {}

  private obterTerminal(): vscode.Terminal {
    if (!this.terminal || this.terminal.exitStatus !== undefined) {
      this.terminal = vscode.window.createTerminal({
        name: 'Verbo',
        iconPath: new vscode.ThemeIcon('terminal')
      });
    }
    return this.terminal;
  }

  private resolverCli(resourceUri: vscode.Uri): string {
    const config = vscode.workspace.getConfiguration('verbo');
    const customCli = config.get<string>('cliPath');
    if (customCli && customCli.trim() !== '') {
      return customCli.trim();
    }

    const workspaceFolder = vscode.workspace.getWorkspaceFolder(resourceUri);
    if (workspaceFolder) {
      const candidatoBuild = path.join(workspaceFolder.uri.fsPath, 'build', 'verbo');
      if (fs.existsSync(candidatoBuild)) {
        return `./build/verbo`;
      }
    }

    return 'verbo';
  }

  private async salvarSeNecessario(document: vscode.TextDocument) {
    if (document.isDirty) {
      await document.save();
    }
  }

  public async executarArquivo(uri?: vscode.Uri) {
    const doc = await this.obterDocumentoAtivo(uri);
    if (!doc) return;

    await this.salvarSeNecessario(doc);
    const cli = this.resolverCli(doc.uri);
    const terminal = this.obterTerminal();
    terminal.show();

    const workspaceFolder = vscode.workspace.getWorkspaceFolder(doc.uri);
    let caminhoRelativo = doc.uri.fsPath;
    if (workspaceFolder) {
      caminhoRelativo = path.relative(workspaceFolder.uri.fsPath, doc.uri.fsPath);
    }

    terminal.sendText(`${cli} executar "${caminhoRelativo}"`);
  }

  public async compilarArquivo(uri?: vscode.Uri) {
    const doc = await this.obterDocumentoAtivo(uri);
    if (!doc) return;

    await this.salvarSeNecessario(doc);
    const cli = this.resolverCli(doc.uri);
    const terminal = this.obterTerminal();
    terminal.show();

    const workspaceFolder = vscode.workspace.getWorkspaceFolder(doc.uri);
    let caminhoRelativo = doc.uri.fsPath;
    if (workspaceFolder) {
      caminhoRelativo = path.relative(workspaceFolder.uri.fsPath, doc.uri.fsPath);
    }

    terminal.sendText(`${cli} compilar "${caminhoRelativo}"`);
  }

  public async servirArquivo(uri?: vscode.Uri) {
    const doc = await this.obterDocumentoAtivo(uri);
    if (!doc) return;

    await this.salvarSeNecessario(doc);
    const cli = this.resolverCli(doc.uri);
    const config = vscode.workspace.getConfiguration('verbo');
    const host = config.get<string>('servidorHost', '127.0.0.1');
    const porta = config.get<string>('servidorPorta', '5000');

    const terminal = this.obterTerminal();
    terminal.show();

    const workspaceFolder = vscode.workspace.getWorkspaceFolder(doc.uri);
    let caminhoRelativo = doc.uri.fsPath;
    if (workspaceFolder) {
      caminhoRelativo = path.relative(workspaceFolder.uri.fsPath, doc.uri.fsPath);
    }

    terminal.sendText(`${cli} servir "${caminhoRelativo}" --host ${host} --porta ${porta}`);
  }

  public async verificarArquivo(uri?: vscode.Uri) {
    const doc = await this.obterDocumentoAtivo(uri);
    if (!doc) return;

    await this.salvarSeNecessario(doc);
    await this.diagnostics.verificarArquivo(doc);

    const collection = this.diagnostics.getCollection();
    const diags = collection.get(doc.uri);
    if (!diags || diags.length === 0) {
      vscode.window.showInformationMessage(`✅ Verbo: Arquivo '${path.basename(doc.fileName)}' está sintaticamente correto!`);
    } else {
      vscode.window.showErrorMessage(`❌ Verbo: Encontrados ${diags.length} erro(s) de sintaxe em '${path.basename(doc.fileName)}'.`);
    }
  }

  private async obterDocumentoAtivo(uri?: vscode.Uri): Promise<vscode.TextDocument | null> {
    if (uri) {
      return await vscode.workspace.openTextDocument(uri);
    }

    const activeEditor = vscode.window.activeTextEditor;
    if (!activeEditor || activeEditor.document.languageId !== 'verbo') {
      vscode.window.showWarningMessage('Nenhum arquivo Verbo (.vrb) ativo no editor.');
      return null;
    }

    return activeEditor.document;
  }
}

