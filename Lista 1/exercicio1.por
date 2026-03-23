programa {

  inclua biblioteca Matematica --> mat

  real nota1, nota2, nota3

  funcao real media() {
    real somaNotas = nota1 + nota2 + nota3
    real media = somaNotas / 3
    retorne media
  }

  funcao inicio(){

    escreva("Notas:")
    leia(nota1, nota2, nota3)

    escreva("MEDIA = ", mat.arredondar(media(), 2), "\n" )

    se (media() >= 6) 
    {
    escreva("APROVADO")
    }
    senao 
    {
    escreva("REPROVADO")
    }

  }

  
}
