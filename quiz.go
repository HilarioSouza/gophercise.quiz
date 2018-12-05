package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var filepath string
	var acertos int
	var total int
	var tempo int

	if len(os.Args) == 2 {
		filepath = os.Args[1]
		fmt.Printf(filepath + "\n")
	} else {
		//fmt.Println("O arquivo não foi informado.")
		//os.Exit(1)
		filepath = "quiz.csv"
	}
	if len(os.Args) == 3 {
		tempo, _ = strconv.Atoi(os.Args[2])
	} else {
		tempo = 30
	}
	//fmt.Print(os.Args)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	Reader := csv.NewReader(bufio.NewReader(file))
	for {
		record, err := Reader.Read()
		if err == io.EOF {
			break
		}
		total++
		fmt.Printf("Responda a questão: %s\n", record[0])
		var resp string
		fmt.Print("Digite a resposta: ")
		timer := time.NewTimer(time.Duration(tempo) * time.Second)
		go fmt.Scanln(&resp)
		<-timer.C
		if strings.Trim(resp, " ") == strings.Trim(record[1], " ") {
			acertos++
		}
	}
	fmt.Printf("Resultado: %d de %d", acertos, total)
}
