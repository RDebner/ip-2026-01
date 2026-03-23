programa {

  inclua biblioteca Matematica --> mat

  real temperaturaFharenheit
  real polegadas 

  funcao inicio() {
    escreva("Temperatura em Fharenheit: ")
    leia(temperaturaFharenheit)
    escreva("Quantidade de chuva em polegadas: ")
    leia(polegadas)

    real tempCelcius = 5*(temperaturaFharenheit - 32) / 9
    real qntMilimetros = polegadas * 25.4

    escreva("O VALOR EM CELSIUS = ", mat.arredondar(tempCelcius, 2), "\nA QUANTIDADE DE CHUVA E = ", qntMilimetros,"\n")
  }
}
