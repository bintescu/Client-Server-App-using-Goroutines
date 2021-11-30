package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

// User struct which contains a name
// a type and a list of social links
type Config struct {
	Length int `json:"length"`
}

func processArray(array []string, channel chan string) {
	var result string
	result = ""

	var sb strings.Builder

	sb.WriteString(result)

	fmt.Println(sb.String())

	for i := 0; i < len(array); i++ {
		sb.WriteByte(array[i][0])
	}

	channel <- sb.String()

}

func handleRequest(conexiune net.Conn, maxStringArr int) {
	mesaj, _ := bufio.NewReader(conexiune).ReadString('\n')

	fmt.Print("Clientul ", conexiune.RemoteAddr(), " s-a conectat \n")
	fmt.Print("Clientul ", conexiune.RemoteAddr(), " a facut request cu datele : ", string(mesaj))

	fmt.Print(conexiune.LocalAddr(), " Serverul a primit requestul \n")
	fmt.Print(conexiune.LocalAddr(), " Serverul proceseaza datele \n")

	res2 := strings.Split(mesaj, " ")

	if len(res2) > maxStringArr {
		res2 = res2[0 : maxStringArr-1]
	}

	channel := make(chan string)

	go processArray(res2, channel)
	select {
	case msg1 := <-channel:
		fmt.Println("Serverul trimite ", msg1, "catre clientul :", conexiune.RemoteAddr())
		fmt.Fprint(conexiune, msg1+"\n")
	}
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Nu am putut gasi fisierul de configurare config.json")
	} else {
		fmt.Println("Am gasit fisierul de configurare config.json")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var configuration Config

	json.Unmarshal(byteValue, &configuration)

	asculta, _ := net.Listen("tcp", ":45623")

	fmt.Println("Am pornit serverul ..")

	var maxStringArr int = 5

	if configuration.Length != 0 {
		fmt.Print("Am setat dimensiunea array-ului la : ", configuration.Length, "\n")
		maxStringArr = configuration.Length
	} else {
		fmt.Print("Dimensiunea array-ului a ramas default la valoarea de : ", maxStringArr, "\n")
	}

	for {

		conn, err := asculta.Accept()
		if err != nil {
			log.Println("Error accepting request:", err)
		}

		go handleRequest(conn, maxStringArr)

	}

}
