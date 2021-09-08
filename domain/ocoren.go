package domain

import (
	"errors"
	"log"
	"time"
)

type OcorenFile struct {
	Edi
	AmountOccurrences int `gorm: "type: varchar(255)"`
	AmountFile        int `gorm: "type: varchar(255)"`
	CreateAt          time.Timer
}

func (ocoren *OcorenFile) Prepare() error {
	var err error
	err = ocoren.validateAmountOccurrences()
	ckeckErr(err)
	err = ocoren.validateAmountFile()
	ckeckErr(err)
	return err
}

func (ocoren *OcorenFile) validateAmountOccurrences() error {
	if ocoren.AmountOccurrences < 0 {
		return errors.New("The amount occurrences is invalidade")
	}
	return nil
}

func (ocoren *OcorenFile) validateAmountFile() error {
	if ocoren.AmountFile < 0 {
		return errors.New("The amount file is invalidade")
	}
	return nil
}

func ckeckErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
