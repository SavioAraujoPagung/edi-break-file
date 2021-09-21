package domain

import (
	"time"
)

//Cabeçalho do arquivo - "000"
type HeadFile struct {
	RecordIdentifier int
	CarrierName      string
	ShipperName      string
	CreatedAt        time.Time
	Filler           string
}

//Cabeçalho dois - "340"
type HeadFileTwo struct {
	RecordIdentifier int
	FileIdentifier   string
	Filler           string
}

//Informação de transportadora - "341"
type Carrier struct {
	RecordIdentifier int
	RegisteredNumber string
	Name             string
	Filler           string
}

//Conhecimento de transporte CT-e - "343"
type TransportKnowledge struct {
	RegisteredNumber   string
	ContractingCarrier string
	Series             int
	Number             int
}

//Nota fiscal - NF-e
type Invoice struct {
	RegisteredNumber string
	Series           int
	Number           int
}

//Codigo da ocorrencia - vide tabela de ocorrencias Proceda-3.1
type OccurrenceCode struct {
	Code        int
	Description string
}

//Informações sobre uma ocorrencia - "342"
type Occurrence struct {
	RecordIdentifier int
	Invoice          Invoice
	OccurrenceCode   OccurrenceCode
	OccurrenceDate   time.Time
	ObservationCode  int
	Text             string
	Filler           string
}

type OccurrenceProceda struct {
	HeadFile
	HeadFileTwo
}

func (proceda OccurrenceProceda) OpenFile() (fileOcoren []byte, err error) {
	return nil, nil
}
func (proceda OccurrenceProceda) ReadHead(fileOcoren []string) (err error) {
	return nil
}
func (proceda OccurrenceProceda) CarrierDatas(fileOcoren []string) (err error) {
	return nil
}
func (proceda OccurrenceProceda) DispacherDatas(fileOcoren []string) (err error) {
	return nil
}
func (proceda OccurrenceProceda) ReadOccurrences(fileOcoren []string) (err error) {
	return nil
}
