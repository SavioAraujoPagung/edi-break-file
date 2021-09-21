package domain

type OccurrencesFile interface {
	OpenFile() (fileOcoren []byte, err error)
	ReadHead(fileOcoren []string)
	CarrierDatas(fileOcoren []string)
	DispacherDatas(fileOcoren []string)
	ReadOccurrences(fileOcoren []string)
}
