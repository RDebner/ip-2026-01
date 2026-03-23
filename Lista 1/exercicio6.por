programa {

  inclua biblioteca Matematica --> mat
  
  funcao inicio() {
    inteiro totalTemperaturas
    escreva("Quantas temperaturas deseja inserir? ")
    leia(totalTemperaturas)
    
    inteiro temperaturasFharenheit[totalTemperaturas]
    real temperaturasCelcius[totalTemperaturas]
    real temperaturaCelcius
    
    para (inteiro i = 0; i < totalTemperaturas; i++) {
      escreva("Digite a temperatura ", i+1, " em Fharenheit: ")
      leia(temperaturasFharenheit[i])
      real temperaturaCelcius = 5*(temperaturasFharenheit[i] - 32) / 9
      temperaturasCelcius[i] = temperaturaCelcius
    }


    para (inteiro i = 0; i < totalTemperaturas; i++) {
      escreva(temperaturasFharenheit[i]," FHARENHEIT EQUIVALE A ", mat.arredondar(temperaturasCelcius[i], 2)," CELSIUS \n")
    }
  }

  
}

