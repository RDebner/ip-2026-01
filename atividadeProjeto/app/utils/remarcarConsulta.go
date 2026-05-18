package utils

import (
	"log"
)

func RemarcarConsulta(email, novaData string) error {
	query := `UPDATE consultas SET data_consulta = $1 WHERE email = $2`
	_, err := DB.Exec(query, novaData, email)

	if err != nil {
		log.Printf("Erro ao remarcar a consulta: %v", err)
		return err
	}

	log.Println("Consulta remarcada com sucesso!")
	return nil
}