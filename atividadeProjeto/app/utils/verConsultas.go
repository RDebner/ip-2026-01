package utils

import (
	"log"
)

type Consultas struct {
	Nome string
	Email string
	Data string
}

func VerConsulta(email string) (*Consultas, error){
	query := `SELECT nome, email, data_consulta FROM consultas WHERE email = $1`
	var cons Consultas
	err := DB.QueryRow(query, email).Scan(&cons.Nome, &cons.Email, &cons.Data)
	if err != nil {
		log.Printf("Erro ao buscar consulta no banco de dados: %v", err)
        return nil, err
	}

	return &cons, nil
}

