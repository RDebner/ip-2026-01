programa {

  const real popular = 1
  const real geral = 5
  const real arquibancadas = 10
  const real cadeiras = 20

  inteiro numeroTestes
  inteiro numJogo = 1
  inteiro numeroPessoas
  real categPopular
  real categGeral
  real categArquibancada
  real categCadeiras

  funcao inicio() {
    escreva("Número de casos de testes: ")
    leia(numeroTestes)
    inteiro vetor[numeroTestes]
    inteiro i = 0
    enquanto (i < numeroTestes) {
      escreva("Número de pessoas jogo ", i + 1 , ": ")
      leia(numeroPessoas)
      escreva("Porcentagem de pessoas que compraram ingresso na categoria Popular: ")
      leia(categPopular)
      escreva("Porcentagem de pessoas que compraram ingresso na categoria Geral: ")
      leia(categGeral)
      escreva("Porcentagem de pessoas que compraram ingresso na categoria Arquibancada: ")
      leia(categArquibancada)
      escreva("Porcentagem de pessoas que compraram ingresso na categoria Cadeiras: ")
      leia(categCadeiras)

      vetor[i] = calcularValor()

      i = i + 1
    }

    i = 0

    enquanto (i < numeroTestes) {
      escreva("A RENDA DO JOGO ", i + 1, " E: ", vetor[i], "\n")
      i = i + 1
    }

  }

  funcao real calcularValor () {
    real qntPop = numeroPessoas*(categPopular / 100)
    real qntGeral = numeroPessoas*(categGeral / 100)
    real qntArq = numeroPessoas*(categArquibancada / 100)
    real qntCad = numeroPessoas*(categCadeiras / 100)

    real valor = qntPop*popular + qntGeral*geral + qntArq*arquibancadas + qntCad*cadeiras

    retorne valor

  }

}

// • O número de pessoas que compraram ingresso para o jogo correspondente ao caso de teste.
// • A percentagem de pessoas que compraram ingresso na categoria Popular.
// • A percentagem de pessoas que compraram ingresso na categoria Geral.
// • A percentagem de pessoas que compraram ingresso na categoria Arquibancada.
// • A percentagem de pessoas que compraram ingresso na categoria Cadeiras.

//55000 20.2 50.4 30.2 10.2 428010
//49732 15.2 53.4 20.24 11.16 352003.09
//67890 30.0 42.20 23.8 4.0 379505.09
