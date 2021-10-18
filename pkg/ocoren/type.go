package ocoren

//Cabeçalho do arquivo - "000"
type HeadFile struct {
	HeadFileRecordIdentifier int    `json:"identificador"`
	SenderName               string `json:"remetente"`
	RecipientName            string `json:"destinatario"`
	CreatedAt                string `json:"data_criacao"`
	FillerHeadFile           string `json:"complemento"`
}

//Cabeçalho dois - "340"
type HeadFileTwo struct {
	HeadFileTwoRecordIdentifier int    `json:"identificador"`
	FileIdentifier              string `json:"identificador_arquivo"`
	FillerHeadFileTwo           string `json:"complemento"`
}

//Informação de transportadora - "341"
type Carrier struct {
	CarrierRecordIdentifier   int                  `json:"identificador"`
	RegisteredNumberCarrier   string               `json:"cnpj_transportadora"`
	Name                      string               `json:"nome_transportadora"`
	FillerCarrier             string               `json:"complemento"`
	AmountTransportKnowledges int                  `json:"quantidade_cte"`
	AmountOccurrences         int                  `json:"quantidade_ocorrencia"`
	TransportKnowledges       []TransportKnowledge `json:"ct-e"`
}

//Conhecimento de transporte CT-e - "343"
type TransportKnowledge struct {
	TransportKnowledgeRecordIdentifier int          `json:"identificador"`
	RegisteredNumberCte                string       `json:"cgc_contratante"`
	ContractingCarrier                 string       `json:"transportadora_contratante"`
	AmountOccurrences                  int          `json:"quantidade_ocorrencia"`
	Series                             int          `json:"cte_serie"`
	Number                             int          `json:"cte_numero"`
	Occurrences                        []Occurrence `json:"ocorrencias"`
}

//Nota fiscal - NF-e
type Invoice struct {
	RegisteredNumberInvoice string `json:"nfe_cnpj_emitente"`
	Series                  int    `json:"nfe_serie"`
	Number                  int    `json:"nfe_numero"`
}

//Codigo da ocorrencia - vide tabela de ocorrencias Proceda-3.1
type OccurrenceCode struct {
	Code        int    `json:"codigo_ocorrencia"`
	Description string `json:"nome_ocorrencia"`
}

//Informações sobre uma ocorrencia - "342"
type Occurrence struct {
	OccurrenceRecordIdentifier int            `json:"identificador"`
	Invoice                    Invoice        `json:"nf-e"`
	OccurrenceCode             OccurrenceCode `json:"codigo_ocorencia"`
	OccurrenceDate             string         `json:"data_ocorencia"`
	ObservationCode            int            `json:"observacao_entrega"`
	Text                       string         `json:"texto"`
	FillerOccurrence           string         `json:"complemento"`
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
