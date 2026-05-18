package main

import (
	"fmt"      // Usado para imprimir mensagens no terminal
	"atividadeProjeto/app/handlers"
	"atividadeProjeto/app/utils"
	"log"      // Usado para registrar mensagens de erro ou log
	"net/http" // Usado para criar o servidor HTTP
)

func main() {
	utils.ConnectToDB()

	fileserver := http.FileServer(http.Dir("./static"))

    http.Handle("/", fileserver)

	http.HandleFunc("/marcarConsulta", handlers.MarcarConsultaHandler)

	http.HandleFunc("/desmarcarConsulta", handlers.DesmarcarConsulta)

	http.HandleFunc("/remarcarConsulta", handlers.RemarcarConsulta)

	http.HandleFunc("/verConsultas", handlers.VerConsultaHandler)

	
	fmt.Printf("port running on http://localhost:8081/\n")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}