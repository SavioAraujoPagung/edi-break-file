package domain

type OccurrencesFile interface {
	OpenFile() (fileOcoren []byte, err error)
	readHead(fileOcoren []string)
	carrierDatas(fileOcoren []string)
	dispacherDatas(fileOcoren []string)
	readOccurrences(fileOcoren []string)
	ReadFile(fileName string) (err error)
}
