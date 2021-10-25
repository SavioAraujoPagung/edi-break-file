package ocoren

type OccurrenceFile struct {
	ID                 int    `json:"id"`
	FileName           string `json:"nome_do_arquivo"`
	ContentFile        string `json:"-"`
	AmountRedeployment int    `json:"quantidade_redespacho"`
	AmountOccurrences  int    `json:"quantidade_ocorrencia"`
}

//PROCEDA-3.1
type OccurrenceProceda struct {
	OccurrenceFile OccurrenceFile `json:"dados"`
	HeadFile       HeadFile       `json:"cabecalho"`      //000
	HeadFileTwo    HeadFileTwo    `json:"cabecalhoDois"`  //340
	Carrier        Carrier        `json:"transportadora"` //341
}

//Cabeçalho do arquivo - "000"
type HeadFile struct {
	HeadFileRecordIdentifier int    `json:"identificador" init:"0" end:"3"`
	SenderName               string `json:"remetente" init:"3" end:"38"`
	RecipientName            string `json:"destinatario" init:"38" end:"74"`
	CreatedAt                string `json:"data_criacao" init:"78" end:"85"`
	FillerHeadFile           string `json:"complemento"`
}

//Cabeçalho dois - "340"
type HeadFileTwo struct {
	HeadFileTwoRecordIdentifier int    `json:"identificador" init:"0" end:"3"`
	FileIdentifier              string `json:"identificador_arquivo" init:"3" end:"13"`
	FillerHeadFileTwo           string `json:"complemento"`
}

//Informação de transportadora - "341"
type Carrier struct {
	CarrierRecordIdentifier int          `json:"identificador" init:"0" end:"3"`
	RegisteredNumberCarrier string       `json:"cnpj_transportadora" init:"3" end:"17"`
	Name                    string       `json:"nome_transportadora" init:"17" end:"57"`
	FillerCarrier           string       `json:"complemento" init:"57" end:"119"`
	Occurrences             []Occurrence `json:"ocorrencias"`
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
	Redeployment               []Redeployment `json:"redespacho"`
}

//Redespacho CT-e - "343"
type Redeployment struct {
	RedeploymentRecordIdentifier int    `json:"identificador" init:"0" end:"3"`
	RegisteredNumberCte          string `json:"cgc_contratante" init:"3" end:"17"`
	ContractingCarrier           string `json:"transportadora_contratante" init:"17" end:"27"`
	Series                       int    `json:"cte_serie" init:"27" end:"32"`
	Number                       int    `json:"cte_numero" init:"32" end:"44"`
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
