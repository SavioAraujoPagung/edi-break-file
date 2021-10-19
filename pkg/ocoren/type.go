package ocoren

/*
const (
	SENDER_NAME_INIT = 3
	SENDER_NAME_END  = 38

	RECIPIENT_NAME_INIT = 38
	RECIPIENT_NAME_END  = 74

	CREATED_AT_INIT = 74
	CREATED_AT_END  = 85
)
*/

//Cabeçalho do arquivo - "000"
type HeadFile struct {
	HeadFileRecordIdentifier int    `json:"identificador" init:"0" end:"3"`
	SenderName               string `json:"remetente" init:"3" end:"38"`
	RecipientName            string `json:"destinatario" init:"38" end:"74"`
	CreatedAt                string `json:"data_criacao" init:"78" end:"85"`
	FillerHeadFile           string `json:"complemento"`
}
/*
const (
	FILE_IDENTIFIER_INIT = 3
	FILE_IDENTIFIER_END  = 13
)
*/
//Cabeçalho dois - "340"
type HeadFileTwo struct {
	HeadFileTwoRecordIdentifier int    `json:"identificador" init:"0" end:"3"`
	FileIdentifier              string `json:"identificador_arquivo" init:"3" end:"13"`
	FillerHeadFileTwo           string `json:"complemento"`
}

/*
const (
	REGISTERED_NUMBER_CARRIER_INIT = 3
	REGISTERED_NUMBER_CARRIER_END  = 17

	CARRIER_NAME_INIT = 17
	CARRIER_NAME_END  = 57

	FILLER_CARRIER_INIT = 57
	FILLER_CARRIER_END  = 119
)
*/
//Informação de transportadora - "341"
type Carrier struct {
	CarrierRecordIdentifier   int                  `json:"identificador" init:"0" end:"3"`
	RegisteredNumberCarrier   string               `json:"cnpj_transportadora" init:"3" end:"17"`
	Name                      string               `json:"nome_transportadora" init:"17" end:"57"`
	FillerCarrier             string               `json:"complemento" init:"57" end:"119"`
	AmountTransportKnowledges int                  `json:"quantidade_cte"`
	AmountOccurrences         int                  `json:"quantidade_ocorrencia"`
	TransportKnowledges       []TransportKnowledge `json:"ct-e"`
}

//Conhecimento de transporte CT-e - "343"
type TransportKnowledge struct {
	TransportKnowledgeRecordIdentifier int          `json:"identificador" init:"0" end:"3"`
	RegisteredNumberCte                string       `json:"cgc_contratante" init:"3" end:"17"`
	ContractingCarrier                 string       `json:"transportadora_contratante" init:"17" end:"27"`
	AmountOccurrences                  int          `json:"quantidade_ocorrencia"`
	Series                             int          `json:"cte_serie" init:"27" end:"32"`
	Number                             int          `json:"cte_numero" init:"32" end:"44"`
	Occurrences                        []Occurrence `json:"ocorrencias"`
}

//Nota fiscal - NF-e
type Invoice struct {
	RegisteredNumberInvoice string `json:"nfe_cnpj_emitente" init:"3" end:"17"`
	Series                  int    `json:"nfe_serie" init:"17" end:"20"` 
	Number                  int    `json:"nfe_numero" init:"20" end:"28"`
}

//Codigo da ocorrencia - vide tabela de ocorrencias Proceda-3.1
type OccurrenceCode struct {
	Code        int    `json:"codigo_ocorrencia" init:"42" end:"44"`
	Description string `json:"nome_ocorrencia" `
}

//Informações sobre uma ocorrencia - "342"
type Occurrence struct {
	OccurrenceRecordIdentifier int            `json:"identificador" init:"0" end:"3"`
	Invoice                    Invoice        `json:"nf-e"`
	OccurrenceCode             OccurrenceCode `json:"codigo_ocorencia"`
	OccurrenceDate             string         `json:"data_ocorencia" init:"30" end:"42"`
	ObservationCode            int            `json:"observacao_entrega"`
	Text                       string         `json:"texto" init:"44" end:"115"`
	FillerOccurrence           string         `json:"complemento" init:"115" end:"121"`
}

//PROCEDA-3.1
type OccurrenceProceda struct {
	ID          int                     `json:"id"`
	FileName    string                  `json:"nome_do_arquivo"`
	ContentFile string                  `json:"-"`
	HeadFile    `json:"cabecalho"`      //000
	HeadFileTwo `json:"cabecalhoDois"`  //340
	Carrier     `json:"transportadora"` //341
}
