package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const HEAD_SIZE = 3
const EXTENSION_FILE = ".txt"
const FOLDER_NAME = "NEW_OCOREN/"
const PERMISSION_FOLDER = 0700
const AMOUNT_LINE_BY_OCOREN = 2
const FIST_ID_FILE_OCOREN = 1
const FILE_NAME = "OCOREN_"

func main() {

	var nameFileOcoren string
	fmt.Print("NOME DO ARQUIVO QUE SERÁ QUEBRADO: ")
	fmt.Scanln(&nameFileOcoren)
	var amount int
	fmt.Print("QUANTOS ARQUIVOS QUEBRAR: ")
	fmt.Scanln(&amount)
	fileOcoren, err := ioutil.ReadFile(nameFileOcoren + EXTENSION_FILE)
	checkErr(err)
	ocoren := string(fileOcoren)
	originalOcorenSplit := strings.Split(ocoren, "\n")
	makeOcoren(amount, originalOcorenSplit)
}

func getAmountOccurrencesFile(amoutFile int, amountOccurrences int) (amountOccurrencesFile int) {
	amountOccurrencesFile = amountOccurrences / amoutFile
	return amountOccurrencesFile
}

func getRestOccurrences(amoutFile int, amountOccurrences int) (restOccurrences int) {
	restOccurrences = amountOccurrences % amoutFile
	return restOccurrences
}

func getLineFinal(amountLineWrite int, lineInit int) (lineFinal int) {
	lineFinal = amountLineWrite + lineInit
	return lineFinal
}

func getAmountLineWrite(amountOccurrencesFile int, idNewFile int, restOccurrences int) (amountLineWrite int) {
	if idNewFile <= restOccurrences {
		return (amountOccurrencesFile + 1) * AMOUNT_LINE_BY_OCOREN
	}
	return amountOccurrencesFile * AMOUNT_LINE_BY_OCOREN
}

func getLineInit(idNewFile int, amountOccurrencesFile int, afterlineInit int) (lineInit int) {
	if idNewFile == FIST_ID_FILE_OCOREN {
		return HEAD_SIZE
	}
	lineInit = afterlineInit + amountOccurrencesFile
	return lineInit
}

func writeHead(newOcoren *os.File, originalOcorenSplit []string) {
	for line := 0; line < HEAD_SIZE; line++ {
		fmt.Fprint(newOcoren, originalOcorenSplit[line])
	}
}

func writeBody(lineInit int, newOcoren *os.File, originalOcorenSplit []string, lineFinal int) {
	for line := lineInit; line < lineFinal; line++ {
		//fmt.Fprint(newOcoren, originalOcorenSplit[line])
		fmt.Fprintln(newOcoren, originalOcorenSplit[line])
	}
}

func createFolder() {
	os.Mkdir(FOLDER_NAME, PERMISSION_FOLDER)
}

func createFile(fileName string) (newOcoren *os.File) {
	newOcoren, err := os.Create(fileName)
	checkErr(err)
	return newOcoren
}

func getAmountOccurrences(originalOcorenSplit []string) (amountOccurrences int) {
	amountOccurrences = (len(originalOcorenSplit) - HEAD_SIZE) / AMOUNT_LINE_BY_OCOREN
	return amountOccurrences
}

func makeOcoren(amout int, originalOcorenSplit []string) {
	createFolder()
	amountOccurrences := getAmountOccurrences(originalOcorenSplit)
	amountOccurrencesFile := getAmountOccurrencesFile(amout, amountOccurrences)
	restOccurrences := getRestOccurrences(amout, amountOccurrences)
	lineInit := 1
	amountLineWrite := getAmountLineWrite(amountOccurrencesFile, FIST_ID_FILE_OCOREN, restOccurrences)
	for idNewFile := 1; idNewFile <= amout; idNewFile++ {
		idNewFileStr := strconv.Itoa(idNewFile)
		newOcoren := createFile(FOLDER_NAME + FILE_NAME + idNewFileStr + EXTENSION_FILE)
		lineInit = getLineInit(idNewFile, amountLineWrite, lineInit)
		amountLineWrite = getAmountLineWrite(amountOccurrencesFile, idNewFile, restOccurrences)
		lineFinal := getLineFinal(amountLineWrite, lineInit)
		writeHead(newOcoren, originalOcorenSplit)
		writeBody(lineInit, newOcoren, originalOcorenSplit, lineFinal)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
