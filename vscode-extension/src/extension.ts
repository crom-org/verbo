import * as vscode from 'vscode';
import { VerboDiagnostics } from './diagnostics';
import { VerboHoverProvider } from './hover';
import { VerboCompletionProvider } from './completion';
import { VerboDocumentSymbolProvider } from './symbols';
import { VerboCommands } from './commands';

export function activate(context: vscode.ExtensionContext) {
  console.log('Extensão oficial do Verbo ativada.');

  const selector: vscode.DocumentSelector = { language: 'verbo', scheme: 'file' };

  // 1. Diagnósticos / Validação Sintática
  const diagnostics = new VerboDiagnostics();
  context.subscriptions.push(diagnostics);

  // 2. Comandos
  const commands = new VerboCommands(diagnostics);
  context.subscriptions.push(
    vscode.commands.registerCommand('verbo.executar', (uri?: vscode.Uri) => commands.executarArquivo(uri)),
    vscode.commands.registerCommand('verbo.compilar', (uri?: vscode.Uri) => commands.compilarArquivo(uri)),
    vscode.commands.registerCommand('verbo.servir', (uri?: vscode.Uri) => commands.servirArquivo(uri)),
    vscode.commands.registerCommand('verbo.verificar', (uri?: vscode.Uri) => commands.verificarArquivo(uri))
  );

  // 3. Provedor de Hover
  context.subscriptions.push(
    vscode.languages.registerHoverProvider(selector, new VerboHoverProvider())
  );

  // 4. Provedor de Autocompletação
  context.subscriptions.push(
    vscode.languages.registerCompletionItemProvider(
      selector,
      new VerboCompletionProvider(),
      ' ', '.', '(', ':'
    )
  );

  // 5. Provedor de Símbolos do Documento (Outline)
  context.subscriptions.push(
    vscode.languages.registerDocumentSymbolProvider(selector, new VerboDocumentSymbolProvider())
  );

  // 6. Listeners de eventos de documento
  context.subscriptions.push(
    vscode.workspace.onDidSaveTextDocument((doc) => {
      const config = vscode.workspace.getConfiguration('verbo');
      if (config.get<boolean>('verificarAoSalvar', true)) {
        diagnostics.verificarArquivo(doc);
      }
    }),
    vscode.workspace.onDidOpenTextDocument((doc) => {
      diagnostics.verificarArquivo(doc);
    }),
    vscode.workspace.onDidCloseTextDocument((doc) => {
      diagnostics.limpar(doc.uri);
    })
  );

  // Validar arquivo atualmente aberto ao inicializar
  if (vscode.window.activeTextEditor && vscode.window.activeTextEditor.document.languageId === 'verbo') {
    diagnostics.verificarArquivo(vscode.window.activeTextEditor.document);
  }
}

export function deactivate() {}

