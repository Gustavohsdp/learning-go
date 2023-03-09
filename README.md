## Learning GO

-> Oferece performance
<br>
-> Velocidade na entrega de suas soluções
<br>
-> Economiza com infra, pois o mesmo tem uma redução no poder computacional
<br>
<br>

## Caracteristica

-> Uma linguagem openSource
<br>
-> Utiliza o maximo de recurso computacional e de rede
<br>
-> Extremamente rápido, mesmo utilizando garbage collection
<br>
-> É uma linguagem compilada, com geração apenas de um binário
<br>
-> Multiplataforma (Windows, Linux, Mac OS,...)
<br>
<br>

## Iniciar um projeto em GO

go mod init github.com/xxxxx/xxxxx

<br>
<br>

## Realizar download de pacotes

go mod tidy
<br>
<br>

## Utilização

Para rodar a aplicação executo -> go run main.go
<br>
Para gerar o binario da aplicação -> go buil main.go
<br>
Para gerar o binário para um sistema em específico -> GOOS=windows go build main.go
<br>
<br>

## Utilização das variaveis

a := "Gustavo" -> Quando utilizamos := significa que estou criando e inicializando ao mesmo tempo.
a = "Henrique" -> Não preciso colocar os := pois a mesma já foi criada

Em Go, ele realiza inferencia de tipos, se o tipo ele foi definido, você não consegue mais altera-lo.

Na liguagem Go, não existe por exemplo, while, do while, existe somente for.
<br>
<br>

## Observações

Quando inicio uma função com a letra maiscula eu estou exportando ela, ou seja
consigo utilizar a mesma em outros arquivos, quando coloco a mesma miniscula, não exporto a mesma.
<br>
Sempre que criamos um projeto go e o código não será compartilhado com nenhum outro projeto, deixamos o código na pasta internal.
<br>
Quando fomos precisar compartilhar o código, o mesmo ficará na pasta package
<br>
<br>

## Struct

-> Forma de estruturar dados
-> Sempre quando adiciono dados dentro de uma struct os item dentro da mesma, tem que conter o letra maiuscula no inicio

Exemplo de uma struct

```
   type Order struct {
    Id         string
    Price      float64
    Tax        float64
    FinalPrice float64
  }
```

Para criar uma nova order partindo dessa struct posso fazer

```
  order:= Order {
    Id: "321",
    Price: 10.0,
    Tax: 10.0,
    FinalPrice: 10.0
  }
```

## Ponteiros

a := 10
<br>
b := a
<br>
fmt.Println((b)) -> 10
<br>
b = 20
<br>
fmt.Println((a)) -> 10
<br>
Neste trecho de código acima, criei o a recebendo o valor de 10 e o b recebendo o valor de a,
após o b ter recebido o valor, eles são independentes, cada um por si.
<br>
fmt.Println((b)) -> 20
<br>
<br>

### Outra Situação

a := 10
<br>
b := &a -> Referencia de memoria
<br>
fmt.Println((b)) -> 0xc0000160f8 -> endereço de memoria que o a está apontando
<br>
fmt.Println((\*b)) -> 10 -> Para saber o valor do endereço que o b está apontando tenho que colocar o \*
<br>
Isso significa que o B não vai criar um novo espaço na memoria, o mesmo
vai apontar para o mesmo valor de a, ou seja se o valor de a alterar o b também será
alterado e vice versa.
<br>
<br>
Quando um \* é utilizado em uma linguagem de programação, ele geralmente indica que a variável é um ponteiro para um endereço de memória. Isso significa que, em vez de duplicar o valor da variável na memória, ela armazena um endereço de memória onde o valor da variável é armazenado.

Quando uma variável é um ponteiro, todas as variáveis que têm o mesmo endereço apontado apontam para a mesma localização de memória. Portanto, se uma variável com um endereço de memória altera seu valor, todas as outras variáveis que apontam para esse mesmo endereço de memória verão o valor alterado globalmente.

Em resumo, o uso de ponteiros pode ajudar a economizar espaço de memória, permitir o compartilhamento de dados entre funções e permitir a modificação de valores de variáveis globais a partir de funções. No entanto, é importante ter cuidado ao usar ponteiros para evitar erros como desreferenciar ponteiros nulos ou acessar endereços de memória inválidos.
