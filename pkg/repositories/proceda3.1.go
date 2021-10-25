package repositories

import (
	"fmt"
	"log"

	"math/rand"

	"github.com/SavioAraujoPagung/edi-break-file/pkg/ocoren"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OcorenRepository interface {
	InsertProceda(ocoren *ocoren.OccurrencesFile) *ocoren.OccurrencesFile
}

type OcorenRepositoryDb struct {
	Db *gorm.DB
}

func (repo *OcorenRepositoryDb) ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=break_file_db_dev port=5412 sslmode=disable"
	var err error
	repo.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
}

func (repo *OcorenRepositoryDb) SaveProceda(fileProceda ocoren.OccurrenceProceda) {
	sql := fmt.Sprintf("INSERT INTO occurrence_files (id, file_name, amountredeployment, amountoccurrences) values (%d, '%s', %d, %d)", fileProceda.OccurrenceFile.ID, fileProceda.OccurrenceFile.FileName, fileProceda.OccurrenceFile.AmountRedeployment, fileProceda.OccurrenceFile.AmountOccurrences)
	repo.Db.Exec(sql)

	idHeadCarrier := repo.saveHead(fileProceda)
	for i := 0; i < fileProceda.OccurrenceFile.AmountOccurrences; i++ {
		idInvoices := getId()
		idRedeployments := getId()
		idOccurrences := getId()
		sql = fmt.Sprintf("INSERT INTO invoices (id, registered_number_invoice, nfe_series, nfe_number) values (%d, '%s', %d, %d)", idInvoices, fileProceda.Carrier.Occurrences[i].Invoice.RegisteredNumberInvoice, fileProceda.Carrier.Occurrences[i].Invoice.Series, fileProceda.Carrier.Occurrences[i].Invoice.Number)
		repo.Db.Exec(sql)

		sql = fmt.Sprintf("INSERT INTO occurrences (id, occurrence_code_id, invoice_id, occurrence_record_identifier, observation_code, text_occurrence, filler_occurrence, carrier_id) values (%d, %d, %d, %d, %d, '%s', '%s', %d)", idOccurrences, fileProceda.Carrier.Occurrences[i].OccurrenceCode.Code, idInvoices, fileProceda.Carrier.Occurrences[i].OccurrenceRecordIdentifier, fileProceda.Carrier.Occurrences[i].ObservationCode, fileProceda.Carrier.Occurrences[i].Text, fileProceda.Carrier.Occurrences[i].FillerOccurrence, idHeadCarrier)
		repo.Db.Exec(sql)

		if len(fileProceda.Carrier.Occurrences[i].Redeployment) > 0 {
			sql = fmt.Sprintf("INSERT INTO redeployments (id, redeployment_record_identifier, registered_number_cte, contracting_carrier, cte_series, cte_number, occurrence_id) values (%d, %d, '%s', '%s', %d, %d, %d)", idRedeployments, fileProceda.Carrier.Occurrences[i].Redeployment[0].RedeploymentRecordIdentifier, fileProceda.Carrier.Occurrences[i].Redeployment[0].RegisteredNumberCte, fileProceda.Carrier.Occurrences[i].Redeployment[0].ContractingCarrier, fileProceda.Carrier.Occurrences[i].Redeployment[0].Series, fileProceda.Carrier.Occurrences[i].Redeployment[0].Number, idOccurrences)
			repo.Db.Exec(sql)
		}

	}
}

func (repo *OcorenRepositoryDb) saveHead(fileProceda ocoren.OccurrenceProceda) (idHeadCarrier int) {
	idHeadFile := getId()
	sql := fmt.Sprintf("INSERT INTO head_files (id, head_file_record_identifier, sender_name, recipient_name) values (%d, %d, '%s', '%s')", idHeadFile, fileProceda.HeadFile.HeadFileRecordIdentifier, fileProceda.HeadFile.SenderName, fileProceda.HeadFile.RecipientName)
	repo.Db.Exec(sql)
	idHeadFileTwo := getId()
	sql = fmt.Sprintf("INSERT INTO head_file_twos (id, head_file_two_record_identifier, file_identifier, filler_head_file_two) values (%d, %d, '%s', '%s')", idHeadFileTwo, fileProceda.HeadFileTwo.HeadFileTwoRecordIdentifier, fileProceda.HeadFileTwo.FileIdentifier, fileProceda.HeadFileTwo.FillerHeadFileTwo)
	repo.Db.Exec(sql)
	idHeadCarrier = getId()
	sql = fmt.Sprintf("INSERT INTO carriers (id, carrier_record_identifier, registered_number_carrier, carrier_name, filler_carrier) values (%d, %d, '%s', '%s', '%s')", idHeadCarrier, fileProceda.Carrier.CarrierRecordIdentifier, fileProceda.Carrier.RegisteredNumberCarrier, fileProceda.Carrier.Name, fileProceda.Carrier.FillerCarrier)
	repo.Db.Exec(sql)
	idOccurrence := getId()
	sql = fmt.Sprintf("INSERT INTO occurrence_procedas (id, head_file_id, head_file_two_id, carrier_id) values (%d, %d, %d, %d)", idOccurrence, idHeadFile, idHeadFileTwo, idHeadCarrier)
	repo.Db.Exec(sql)
	return idHeadCarrier
}

func (repo *OcorenRepositoryDb) FindAllOccurrences() []ocoren.OccurrenceCode {
	var occurrences []ocoren.OccurrenceCode
	id := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		id = append(id, i)
	}
	repo.Db.Find(&occurrences, id)
	return occurrences
}

func getId() (id int) {
	for i := 0; i < rand.Intn(50); i++ {
		id = rand.Intn(1000000)
	}
	return id
}
