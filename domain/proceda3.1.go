package domain

import (
	"fmt"
	"io/ioutil"
	"time"
)

//Cabeçalho do arquivo - "000"
type HeadFile struct {
	RecordIdentifier int       `json: "identificador"`
	CarrierName      string    `json: "transportadora"`
	ShipperName      string    `json: "embarcador"`
	CreatedAt        time.Time `json: "data_criacao"`
	Filler           string    `json: _`
}

//Cabeçalho dois - "340"
type HeadFileTwo struct {
	RecordIdentifier int    `json: "identificador"`
	FileIdentifier   string `json: "identificador_arquivo"`
	Filler           string `json: _`
}

//Informação de transportadora - "341"
type Carrier struct {
	RecordIdentifier    int    `json: "identificador"`
	RegisteredNumber    string `json: "cnpj_transportadora"`
	Name                string `json: "nome_transportadora"`
	Filler              string `json: "_"`
	TransportKnowledges []TransportKnowledge
}

//Conhecimento de transporte CT-e - "343"
type TransportKnowledge struct {
	RecordIdentifier   int    `json: "identificador"`
	RegisteredNumber   string `json: "cgc_contratante"`
	ContractingCarrier string `json: "transportadora_contratante"`
	Series             int    `json: "cte_serie"`
	Number             int    `json: "cte_numero"`
	Occurrences        []Occurrence
}

//Nota fiscal - NF-e
type Invoice struct {
	RegisteredNumber string `json: "nfe_cnpj_emitente"`
	Series           int    `json: "nfe_serie"`
	Number           int    `json: "nfe_numero"`
}

//Codigo da ocorrencia - vide tabela de ocorrencias Proceda-3.1
type OccurrenceCode struct {
	Code        int    `json: "codigo_ocorrencia"`
	Description string `json: "nome_ocorrencia"`
}

//Informações sobre uma ocorrencia - "342"
type Occurrence struct {
	Invoice          []Invoice
	OccurrenceCode   []OccurrenceCode
	RecordIdentifier int       `json: "identificador"`
	OccurrenceDate   time.Time `json: "data_ocorencia"`
	ObservationCode  int       `json: "observacao_entrega"`
	Text             string    `json: "texto"`
	Filler           string    `json: _`
}

//PROCEDA-3.1
type OccurrenceProceda struct {
	HeadFile
	HeadFileTwo
	Carrier
	//todo o conteudo do arquivo
	ContentFile string
	FileName    string
}

//read all content file - OCOREN PROCEDA 3.1
func (proceda *OccurrenceProceda) ReadFile(fileName string) (err error) {
	fileOcoren, err := ioutil.ReadFile(fileName)

	proceda.ContentFile = string(fileOcoren)
	err = proceda.readHead(proceda.ContentFile)
	checkError(err, "Err to read head")
	err = proceda.carrierDatas(proceda.ContentFile)
	checkError(err, "Err to read carrier Datas")
	err = proceda.dispacherDatas(proceda.ContentFile)
	checkError(err, "Err to read dispacher Datas")
	err = proceda.readOccurrences(proceda.ContentFile)
	checkError(err, "Err to read Occurrences")

	return nil
}
func (proceda *OccurrenceProceda) readHead(fileOcoren string) (err error) {
	return nil
}
func (proceda *OccurrenceProceda) carrierDatas(fileOcoren string) (err error) {
	return nil
}
func (proceda *OccurrenceProceda) dispacherDatas(fileOcoren string) (err error) {
	return nil
}
func (proceda *OccurrenceProceda) readOccurrences(fileOcoren string) (err error) {
	return nil
}
func checkError(err error, message string) {
	if message == "" {
		if err != nil {
			fmt.Println("Erro ao fazer alguma coisa")
		}
		return
	}
	if err != nil {
		fmt.Println(message)
	}
}
