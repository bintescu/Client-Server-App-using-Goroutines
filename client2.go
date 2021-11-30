package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//Pasul 1 - Conectarea la server

	for {
		conexiune, _ := net.Dial("tcp", ":45623")
		//Pasul 2 - ce trimitem

		cititor := bufio.NewReader(os.Stdin)
		fmt.Print("Array de string-uri de trimis la server din ", conexiune.LocalAddr(), ": ")

		mesaj, _ := cititor.ReadString('\n')

		//Pasul 3 - trimitem mesajul catre server

		fmt.Fprint(conexiune, mesaj+"\n")

		//Pasul 4 - asteptam raspuns de la server

		mesajServer, _ := bufio.NewReader(conexiune).ReadString('\n')

		fmt.Print("Clientul  : ", conexiune.LocalAddr(), " a primit raspunsul: ", mesajServer)

	}

}
