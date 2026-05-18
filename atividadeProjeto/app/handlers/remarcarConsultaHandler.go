package handlers

import (
	"atividadeProjeto/app/utils"
	"net/http"
)

func RemarcarConsulta(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	email := request.FormValue("email")
	novaData := request.FormValue("novaData")

	err := utils.RemarcarConsulta(email, novaData)
	if err != nil {
		http.Error(response, "Erro ao mudar data da consulta no banco de dados", http.StatusInternalServerError)
		return
	}

}