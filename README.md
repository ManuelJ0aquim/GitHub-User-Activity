# GitHub-User-Activity

> Um aplicativo de linha de comando (CLI) desenvolvido em **Go** para buscar e exibir a atividade recente de um usuário do GitHub.

## Descrição

GitHub User Activity é uma aplicação de linha de comando projetada para visualizar informações e atividade de qualquer usuário do GitHub. Construído com Go, oferece uma maneira simples de acompanhar commits, pull requests, issues, repositórios, seguidores e outras informações de usuários na plataforma.

Este projeto foi criado com o objetivo de **praticar Go**, consumo de APIs REST, manipulação de JSON, organização de código e arquitetura básica de aplicações CLI.

## Funcionalidades

- **Validação de usuário** - Verifica se o usuário existe antes de fazer requisições
- **Eventos** - Buscar e exibir eventos recentes de um usuário do GitHub:
  - **PushEvent** - Commits enviados para repositórios
  - **PullRequestEvent** - Pull requests criados ou atualizados
  - **IssuesEvent** - Issues abertas
  - **IssueCommentEvent** - Comentários em issues
  - **WatchEvent** - Repositórios marcados com estrela
- **Seguidores** - Listar todos os seguidores de um usuário
- **Seguindo** - Listar todos os usuários que um usuário está seguindo
- **Repositórios** - Listar todos os repositórios de um usuário
- Integração com a API pública do GitHub

## Estruturas de Dados

### Evento

Cada evento possui as seguintes propriedades:

- **type** – Tipo do evento (PushEvent, PullRequestEvent, etc.)
- **repo** – Informações do repositório (nome)
- **payload** – Dados específicos do evento (commits, issues, pull requests, etc.)

### Usuário

Cada usuário possui:

- **login** – Nome de usuário do GitHub

### Repositório

Cada repositório possui:

- **name** – Nome do repositório

Os dados são obtidos diretamente da API pública do GitHub.

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

### Sintaxe

```bash
./githubUserActivity <github-username> <comando>
```

### Comandos Disponíveis

- `events` - Exibe os eventos recentes do usuário
- `followers` - Lista os seguidores do usuário
- `following` - Lista quem o usuário está seguindo
- `repos` - Lista os repositórios do usuário

### Exemplos de Uso

#### Visualizar eventos de um usuário

```bash
./githubUserActivity ManuelJ0aquim events
```

#### Listar seguidores

```bash
./githubUserActivity ManuelJ0aquim followers
```

#### Listar quem o usuário está seguindo

```bash
./githubUserActivity ManuelJ0aquim following
```

#### Listar repositórios

```bash
./githubUserActivity ManuelJ0aquim repos
```

## Exemplos de Saída

### Comando `events`

#### PushEvent

```
Events:
Pushed 3 commits to ManuelJ0aquim/Hello-World
```

#### PullRequestEvent

```
Events:
Pull request #42 in ManuelJ0aquim/Hello-World
```

#### IssuesEvent

```
Events:
Opened issue #15 in ManuelJ0aquim/Hello-World
```

#### IssueCommentEvent

```
Events:
Commented on issue #15 in ManuelJ0aquim/Hello-World
```

#### WatchEvent

```
Events:
Starred ManuelJ0aquim/Hello-World
```

### Comando `followers`

```
Followers:
- user1
- user2
- user3
```

### Comando `following`

```
Following:
- userA
- userB
- userC
```

### Comando `repos`

```
Repos:
- Hello-World
- Spoon-Knife
- git-consortium
```

## Estrutura do Projeto

```
.
├── cmd/
│   └── githubUserActivity/
│       ├── main.go              # Ponto de entrada da aplicação (CLI)
│       └── printEvent.go       # Função para imprimir eventos formatados
├── internal/
│   ├── github/
│   │   ├── client.go            # Cliente para buscar eventos
│   │   ├── followers.go         # Função para buscar seguidores
│   │   ├── following.go         # Função para buscar quem está seguindo
│   │   ├── repos.go             # Função para buscar repositórios
│   │   └── userExist.go         # Função para validar se usuário existe
│   └── model/
│       └── event.go             # Definição das estruturas de dados
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
