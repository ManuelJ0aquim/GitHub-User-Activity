# GitHub-User-Activity

> Um aplicativo de linha de comando (CLI) desenvolvido em **Go** para buscar e exibir a atividade recente de um usuário do GitHub.

## Descrição

GitHub User Activity é uma aplicação de linha de comando projetada para visualizar a atividade recente de qualquer usuário do GitHub. Construído com Go, oferece uma maneira simples de acompanhar commits, pull requests, issues e outras ações realizadas por usuários na plataforma.

Este projeto foi criado com o objetivo de **praticar Go**, consumo de APIs REST, manipulação de JSON, organização de código e arquitetura básica de aplicações CLI.

## Funcionalidades

- Buscar eventos recentes de um usuário do GitHub
- Exibir diferentes tipos de eventos:
  - **PushEvent** - Commits enviados para repositórios
  - **PullRequestEvent** - Pull requests criados ou atualizados
  - **IssuesEvent** - Issues abertas
  - **IssueCommentEvent** - Comentários em issues
  - **WatchEvent** - Repositórios marcados com estrela
- Integração com a API pública do GitHub

## Estrutura do Evento

Cada evento possui as seguintes propriedades:

- **type** – Tipo do evento (PushEvent, PullRequestEvent, etc.)
- **repo** – Informações do repositório (nome)
- **payload** – Dados específicos do evento (commits, issues, pull requests, etc.)

Os dados são obtidos diretamente da API do GitHub (`https://api.github.com/users/{username}/events`).

## Requisitos

- Go 1.22+ (ou compatível)
- Sistema Linux, macOS ou WSL (Windows)
- Acesso à internet (para consumir a API do GitHub)

## Instalação

### Pré-requisitos

- Go 1.22 ou superior

### Construir a partir do código fonte

```bash
# Clonar o repositório
git clone https://github.com/ManuelJ0aquim/GitHub-User-Activity.git

# Navegar para o diretório do projeto
cd GitHub-User-Activity

# Compilar o projeto
go build -o githubUserActivity ./cmd/githubUserActivity
```

Isto irá gerar o binário `githubUserActivity`.

## Como Usar

### Comando Básico

```bash
# Executar a aplicação
./githubUserActivity <github-username>
```

### Exemplos de Uso

#### Visualizar atividade de um usuário

```bash
./githubUserActivity octocat
```

## Exemplos de Saída

### PushEvent

```
Pushed 3 commits to octocat/Hello-World
```

### PullRequestEvent

```
Pull request #42 in octocat/Hello-World
```

### IssuesEvent

```
Opened issue #15 in octocat/Hello-World
```

### IssueCommentEvent

```
Commented on issue #15 in octocat/Hello-World
```

### WatchEvent

```
Starred octocat/Hello-World
```

## Estrutura do Projeto

```
.
├── cmd/
│   └── githubUserActivity/
│       └── main.go              # Ponto de entrada da aplicação (CLI)
├── internal/
│   ├── github/
│   │   └── client.go            # Cliente para consumir a API do GitHub
│   └── model/
│       └── event.go             # Definição das estruturas de eventos
├── go.mod
└── README.md
```

## Objetivo do Projeto

Este projeto foi desenvolvido como exercício prático para:

- Aprender Go na prática
- Trabalhar com CLI e argumentos
- Consumir APIs REST externas
- Manipular dados JSON
- Aplicar boas práticas de organização de código
- Criar um projeto real para portfólio

## Autor

**Manuel Joaquim**

- GitHub: [@ManuelJ0aquim](https://github.com/ManuelJ0aquim)

## Referência

Projeto inspirado no desafio **GitHub User Activity CLI** do [roadmap.sh](https://roadmap.sh)
