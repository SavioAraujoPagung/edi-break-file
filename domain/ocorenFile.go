package domain

import (
	"errors"
	"log"
)

type OcorenFile struct {
	Occurrence        []int
	AmountOccurrences int `gorm: "type: int"`
}

func (ocoren *OcorenFile) Prepare() error {
	var err error
	err = ocoren.validateAmountOccurrences()
	ckeckErr(err)
	err = ocoren.validateAmountFile()
	ckeckErr(err)
	return err
}

func (ocoren *OcorenFile) setOccurrence(fileOcoren []string) error {
	//TODO: mapear as ocorrências
	return nil
}

func (ocoren *OcorenFile) validateAmountOccurrences() error {
	if ocoren.AmountOccurrences < 0 {
		return errors.New("The amount occurrences is invalidade")
	}
	return nil
}

func (ocoren *OcorenFile) validateAmountFile() error {
	if ocoren.AmountOccurrences < 0 {
		return errors.New("The amount file is invalidade")
	}
	return nil
}

func ckeckErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

/**/
