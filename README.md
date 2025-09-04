Simulador de Frete com Serviços Independentes

Este projeto tem como objetivo simular cotações de frete (preço e prazo) de forma modular e escalável, utilizando microserviços independentes com comunicação baseada em Pub/Sub, cache com Redis, observabilidade com Datadog, e armazenamento analítico com Elasticsearch.

Cada microserviço é independente e pode ser executado, testado e implantado isoladamente.

################################################################################################################

Fase 1:

    Infraestrutura local com Docker Compose
    Serviço address-svc criado
    Redis operando como serviço de cache
    Aplicação em Go com build Docker funcional

################################################################################################################

Fase 2:

    - Criação do Dockerfile para o microsserviço address-svc
    - Configuração do build multi-stage (compilação e imagem final leve com Alpine)
    - Docker Compose configurado para orquestrar Redis e o address-svc
    - Utilização de `.env.example` para carregar variáveis no container
    - Teste do `address-svc` via `docker compose up --build`
    - Ajuste da versão do Go utilizada no Dockerfile para compatibilidade com o `go.mod` (`go 1.23.4`)
    - Uso do modo `detached` (`docker compose up -d`) para rodar containers em background

################################################################################################################

Fase 3:

    - Criação do domínio de negócio do address-svc
    - Implementação da função NormalizeCEP com validação e normalização de CEPs
    - Implementação da interface Provider para abstração de provedores de CEP
    - Criação de MockProvider que simula resposta de cidade/UF para qualquer CEP
    - Integração com Redis usando TTL configurável
    - Implementação de lógica de cache: busca no Redis antes de consultar provedor
    - Criação do service.Lookup(cep) que centraliza a lógica do serviço
    - Testes unitários da função NormalizeCEP

################################################################################################################

Fase 4:

    - Implementação do handler HTTP REST no microsserviço address-svc
    - Definição da rota: GET /cep/{cep} para busca de endereço por CEP
    - Criação da função FindCEP no service para centralizar a lógica da rota
    - Utilização de cache Redis e fallback para provedor (mock)
    - Uso do logger estruturado (zap) com tratamento de erro crítico no main
    - Registro de rotas com o pacote gorilla/mux
    - Criação e execução de testes unitários para o handler HTTP com httptest
    - Validação de resposta JSON com status HTTP 200 e corpo esperado
    - Testes executados com make test, cobrindo sucesso e validações

################################################################################################################

Estrutura de diretórios (até o momento - 01/09/2025)

freight-simulator/
│
├── docker-compose.yml           # Orquestra Redis + address-svc
├── .env.example                 # Variáveis de ambiente do address-svc
├── Makefile                     # Comandos utilitários Go
│
├── services/
│   └── address/
│       ├── Dockerfile           # Build do serviço address-svc
│       ├── go.mod               # Módulo Go inicializado para address-svc
│       ├── go.sum               # Resoluções de dependência
│       └── cmd/
│           └── api/
│               └── main.go     # Entry point mínimo para address-svc
│
└── pkg/
    └── (ainda vazio – será usado nas próximas fases)

################################################################################################################

Funcionalidade dos principais arquivos

Makefile
    - Automatiza os comandos de build, run e test do address-svc

.env.example
    - Contém as variáveis de ambiente necessárias para o serviço funcionar localmente ou em container

docker-compose.yml
    - Orquestra os serviços necessários para a aplicação funcionar localmente.

Dockerfile (específico para cada microsserviço)
    - Build de dois estágios: compila o binário Go e gera imagem final minimalista com Alpine.

main.go (dentro do diretório do microsserviço)
    - Arquivo mínimo para inicialização do serviço:

################################################################################################################

Comandos

go mod init freight-simulator/services/address

docker compose up --build