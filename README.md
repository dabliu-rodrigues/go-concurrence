# Estudo de Concorrência com Golang

Este repositório contém um estudo sobre concorrência em Golang. O objetivo do projeto é capturar uma lista de milhões de domínios, filtrar por domínios brasileiros e verificar o certificado TLS de cada um desses domínios.

## Prólogo

### Concorrência em Golang

Golang é conhecido por seu suporte robusto à concorrência. A linguagem utiliza goroutines, que são funções ou métodos executados concorrentemente com outras goroutines no mesmo endereço de memória. Goroutines são leves e gerenciadas pela runtime do Go, permitindo que milhares de goroutines sejam executadas simultaneamente sem sobrecarregar o sistema.

### Wait Groups

Para sincronizar goroutines e garantir que todas tenham terminado antes de continuar a execução do programa, o Go fornece o pacote `sync`, que inclui o tipo `WaitGroup`. Um `WaitGroup` permite que você espere pela conclusão de um conjunto de goroutines. Você adiciona o número de goroutines ao `WaitGroup` com `Add`, decreta o contador com `Done` e espera todas terminarem com `Wait`.

Para mais informações, consulte a [documentação oficial do Go](https://golang.org/doc/effective_go#concurrency) e a [documentação do pacote sync](https://pkg.go.dev/sync#WaitGroup).

## Descrição do Projeto

O projeto realiza as seguintes etapas:

1. **Download da Lista de Domínios**: Baixa um arquivo CSV contendo uma lista de milhões de domínios do site Majestic.
2. **Filtragem de Domínios Brasileiros**: Filtra a lista para manter apenas os domínios brasileiros (.br).
3. **Verificação de Certificados TLS**: Verifica o certificado TLS de cada domínio brasileiro, utilizando concorrência para acelerar o processo.

## Arquitetura do Código

O código é composto pelas seguintes funções:

- **DownloadMillionDomains**: Faz o download do arquivo CSV com a lista de domínios, se ele não existir no diretório.
- **CreateUrlList**: Lê o arquivo CSV e cria uma lista de domínios brasileiros.
- **CheckSSL**: Verifica o certificado SSL de um domínio específico.
- **main**: Função principal que orquestra o fluxo do programa.

## Execução do Código

Para executar o código, siga os seguintes passos:

1. Clone o repositório:
    ```sh
    git clone https://github.com/seu-usuario/seu-repositorio.git
    ```

2. Navegue até o diretório do projeto:
    ```sh
    cd seu-repositorio
    ```

3. Execute o programa:
    ```sh
    go run main.go
    ```

## Detalhes da Implementação

### Download da Lista de Domínios

A função `DownloadMillionDomains` baixa o arquivo CSV contendo a lista de domínios do site Majestic. O arquivo é salvo localmente com o nome `majestic_million.csv`. Se o arquivo já existir, o download é ignorado.

### Filtragem de Domínios Brasileiros

A função `CreateUrlList` lê o arquivo CSV e filtra apenas os domínios brasileiros (com extensão `.br`). Os domínios são armazenados em uma slice de structs `TopURL`.

### Verificação de Certificados TLS

A função `CheckSSL` estabelece uma conexão TLS com cada domínio e verifica a validade do certificado. O tempo restante para a expiração do certificado é logado.

### Concorrência

A verificação dos certificados TLS é realizada de forma concorrente utilizando goroutines e um `sync.WaitGroup`. Isso permite que múltiplos domínios sejam verificados simultaneamente, aumentando a eficiência do processo.

## Logs

Os resultados das verificações são salvos em um arquivo de log chamado `logs.txt`.
