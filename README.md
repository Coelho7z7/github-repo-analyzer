# GitHub Repo Analyzer

Uma pequena ferramenta de linha de comando escrita em Go que consulta a GitHub REST API e exibe informações de um repositório.

## O que o projeto demonstra

- Integração real com a GitHub REST API
- Requisições HTTP com `net/http`
- Leitura e decodificação de JSON
- Structs em Go
- Tratamento de erros
- Uso opcional de token por variável de ambiente
- Separação básica entre cliente da API e apresentação dos dados

## Requisitos

- Go 1.26 ou superior
- Internet

## Como executar

Clone o projeto e entre na pasta:

```bash
git clone https://github.com/SEU-USUARIO/github-repo-analyzer.git
cd github-repo-analyzer
```

Execute:

```bash
go run . Coelho7z7/GoStock
```

Você pode analisar qualquer repositório público no formato:

```text
owner/repository
```

## Token opcional

Repositórios públicos podem ser consultados sem token.

Para aumentar os limites da API, você pode configurar um token em `GITHUB_TOKEN`.

PowerShell:

```powershell
$env:GITHUB_TOKEN="seu_token"
go run . Coelho7z7/GoStock
```

Nunca coloque tokens diretamente no código ou no repositório.

## Exemplo de saída

```text
GitHub Repo Analyzer
====================
Repositório: Coelho7z7/GoStock
Descrição:   ...
URL:         https://github.com/Coelho7z7/GoStock
Linguagem:   Go
Branch:      main
Stars:       0
Forks:       0
Issues:      0
Watchers:    0
Privado:     false
Arquivado:   false
Criado em:   ...
Atualizado:  ...
```

## GitHub Developer Program

Este projeto foi criado como uma integração independente que utiliza a GitHub REST API para desenvolver uma aplicação sobre a plataforma GitHub.

Antes de solicitar participação no GitHub Developer Program, consulte os requisitos atuais na documentação oficial:

https://docs.github.com/en/integrations/concepts/github-developer-program

A participação no programa e a concessão de benefícios ou badges dependem da análise e dos critérios atuais do GitHub. Este projeto não garante aprovação.

## Contato

Antes de publicar o projeto como uma integração, substitua este campo pelo seu e-mail de suporte:

`Matheushclope@gmail.com`

## Licença

MIT
