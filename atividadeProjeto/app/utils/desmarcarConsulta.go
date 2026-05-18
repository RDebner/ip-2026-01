package utils

import (
	"log"
)

func DesmarcarConsulta(nome string) error {
	query := `DELETE FROM consultas WHERE nome = $1`
	_, err := DB.Exec(query, nome)

	if err != nil {
		log.Printf("Erro ao desmarcar consulta: %v", err)
		return err
	}

	log.Println("Consulta desmarcada com sucesso!")
	return nil
}