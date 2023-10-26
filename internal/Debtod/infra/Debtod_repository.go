package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/rochaeduardo997/money-back/internal/Debtod/entity"
)

type DebtodRepositoryDB struct {
	db *sql.DB
}

func NewDebtodRepository(db *sql.DB) (debtodRepository *DebtodRepositoryDB) {
	debtodRepository = &DebtodRepositoryDB{db: db}
	return
}

func (dr *DebtodRepositoryDB) Save(d *entity.Debtod) (*entity.Debtod, error) {
	err := d.Validate()
	if err != nil {
		return nil, err
	}

	d.Id = uuid.NewString()

	defer dr.db.Close()

	tx, err := dr.db.Begin()
	if err != nil {
		return nil, err
	}

	var debtod entity.Debtod

	debtodSQL := `
	  INSERT INTO tbl_debtods
		VALUES($1,$2,$3,$4,$5,$6)
		RETURNING *
	`
	tx.QueryRow(debtodSQL, &d.Id, &d.Name, &d.BussinessName, &d.Observation, &d.CPF_CNPJ, &d.PersonType).Scan(&debtod.Id, &debtod.Name, &debtod.Surname, &debtod.BussinessName, &debtod.Observation, &debtod.CPF_CNPJ, &debtod.PersonType)

	tx.Rollback()

	return &debtod, nil
}
