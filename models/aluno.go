package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	NOME string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
