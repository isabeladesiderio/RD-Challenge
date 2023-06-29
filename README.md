
# Customer Success Balancing

O projeto  é responsável por equilibrar o atendimento de clientes entre os Customer Success (CS) disponíveis. Ela determina o CS mais adequado para acompanhar estratégicamente cada cliente com base em suas pontuações.

## Pré-requisitos

Antes de usar a função `CustomerSuccessBalancing`, é necessário ter os seguintes elementos:

- Uma lista de entidades de Customer Success representadas pelo tipo `Entity`, contendo informações sobre o ID e a pontuação de cada CS.
- Uma lista de entidades de clientes, também representadas pelo tipo `Entity`, contendo informações sobre o ID e a pontuação de cada cliente.
- Uma lista de IDs de Customer Success ausentes, representada por um array de inteiros, caso haja CS que não estejam disponíveis para atendimento.

## Configuração

- Necessário versão mínima 1.16 da linguagem Go.

## Estrutura do Projeto

- Na pasta `cmd` está o arquivo `main.go` que possui a função `main()`, responsável pela execução do projeto.
- Na pasta `internal` cntém o pacote `csbalancing`, onde está todo o código presente nesse projeto dividido em 5 arquivos:
    - `csbalancing_test.go` é um arquivo onde contém cenários de testes.
    - `csbalancing.go` é um arquivo onde se encontra todas as funções necessárias para a execução do projeto.
    - `customers.go` é um arquivo onde contém funções voltadas para o cliente.
    - `customersuccess.go` é  um arquivo onde contém funções voltadas para o Customer Success (CS).
    - `successaway.go` é  um arquivo onde contém funções voltadas para o Customer Success ausentes.

## Testes

- Para rodar os testes execute o comando `go test -v ./...` no terminal. Esse comando serve para exibir informações detalhadas durante a execução dos testes em todos os pacotes e subpacotes do diretório atual.
- Para rodar a função main() basta apenas executar o comando `go run cmd/main.go`. Ela retornará apenas os IDs dos CSs que atendem mais clientes.

## Descrevendo as funções

A função `CustomerSuccessBalancing` segue os seguintes passos:

1.  A função `CheckCustomerSuccessEntities` tem como objetivo verificar se as entidades de Customer Success fornecidas estão corretas e dentro dos limites estabelecidos.
    
2.  A função `CheckCustomerEntities` tem como objetivo verificar se as entidades de clientes fornecidas estão corretas e dentro dos limites estabelecidos.
    
3.  A função `CheckDuplicateCSScore` é empregada para verificar se há Customer Success com pontuações duplicadas e removê-los, mantendo apenas o CS com a pontuação mais alta.
    
4.  A função `RemoveCustomerSuccessAway` é utilizada para remover da lista de Customer Success aqueles que estão ausentes, com base nos IDs fornecidos na lista `customerSuccessAway`.
    
5.  A função `CountCustomersPerCustomerSuccess` é usada para contar o número de clientes atendidos por cada Customer Success. Ela percorre a lista de clientes e verifica qual Customer Success é o mais adequado para atendê-los com base em suas pontuações.
    
6.  A função `FindClosestCSToCustomer` é utilizada para encontrar o Customer Success mais próximo de cada cliente com base nas suas pontuações. 
    
7.  A função `FindCSWithHighestCustomers` é empregada para determinar o Customer Success que atendeu ao maior número de clientes. 

8. A função `CheckMultipleMaxCustomers` verifica se há mais de um Customer Success com o mesmo número máximo de clientes, caso houver retorna o valor 0.
    
9. A função `FindMaxCustomerSuccessID` é utilizada para obter o ID do Customer Success com o maior número de clientes atendidos.
    
10. A função `FindCustomerSuccessIndexByID` é usada para encontrar o índice do Customer Success selecionado na lista original de Customer Success.
    
11. A função `CustomerSuccessBalancing` retorna o índice do CS que atende mais clientes.


## Autor

- Nome: Isabela Desiderio
- E-mail: isabelaaraujodesiderio@gmail.com
- GitHub: isabeladesiderio
- LinkedIn: https://www.linkedin.com/in/isabela-araujo-desiderio/