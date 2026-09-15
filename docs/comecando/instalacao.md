# Instalação e Configuração

A ferramenta oficial de linha de comando do Verbo (`verbo`) permite compilar, executar, verificar e servir aplicações escritas em Verbo.

---

## Pré-requisitos

- **Go (Golang)**: Versão 1.20 ou superior instalada e acessível no seu `$PATH`.
- **Sistema Operacional**: Linux, macOS ou Windows.

Para verificar se o Go está instalado no seu terminal:
```bash
go version
```

---

## Compilando o CLI Verbo a partir do Código-Fonte

Clone o repositório oficial e faça o build do binário:

```bash
git clone https://github.com/MrJc01/crom-verbo.git
cd crom-verbo

# Compila o binário do CLI
go build -o verbo ./cmd/verbo
```

Para disponibilizar o comando `verbo` globalmente em seu sistema:

### No Linux e macOS:
```bash
sudo mv ./verbo /usr/local/bin/verbo
```

### No Windows:
Mova `verbo.exe` para uma pasta presente nas Variáveis de Ambiente do seu sistema (como `C:\Program Files\Verbo\` ou `%USERPROFILE%\bin`).

---

## Verificando a Instalação

Execute o comando de versão para garantir que o CLI está operacional:

```bash
verbo versão
```

Saída esperada:
```
🇧🇷 Verbo v0.1.0 — Linguagem de Programação em Português
```

Execute a ajuda para conferir os comandos disponíveis:
```bash
verbo ajuda
```

