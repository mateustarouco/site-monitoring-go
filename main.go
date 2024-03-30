package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const cMonitoring = 5
const cDalay = 5

// main is the entry point of the program.
//
// It initializes the name and version variables, prints an introduction message,
// and enters a loop to handle user commands. The function does not take any
// parameters and does not return anything.
func main() {
	name := "Mateus"
	version := "1.0.0"

	fmt.Println("Ola Sr. ", name, "este programa esta na versão ", version)
	introduction()
	for {

		command := commandRead()

		switch command {
		case 1:
			fmt.Println("Iniciando Monitoramento")
			monitoring()

		case 2:
			fmt.Println("Exibindo Logs")
			printLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("não conheço esse comando ")
			os.Exit(-1)
		}
	}
}

// introduction prints a menu of options for monitoring and displaying logs.
//
// No parameters.
// No return value.
func introduction() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")

}

// commandRead reads an integer from the user input and returns it.
//
// This function does not take any parameters.
// It returns an integer representing the user input.
func commandRead() int {
	var command_ int
	fmt.Scan(&command_)
	return command_
}

// monitoring is a Go function that performs monitoring.
//
// It does not take any parameters and does not return any values.
func monitoring() {
	fmt.Println("Monitorando ...")
	sites := readFile()

	for i := 0; i < cMonitoring; i++ {
		clearTerminal()
		for i, site := range sites {
			testingSites(site, i)

		}
		time.Sleep(cDalay * time.Second)
	}
	clearTerminal()

	finalMessage()
}

// finalMessage prints a message to the console and calls the introduction function.
//
// This function does not take any parameters.
// It does not return any values.
func finalMessage() {
	fmt.Println("Comando executado, o quer deseja fazer agora?")
	introduction()
}

// testingSites is a Go function that takes a site string and an integer i as parameters.
// It makes an HTTP GET request to the given site, checks the response status code, and logs the result.
func testingSites(site string, i int) {

	resp, err := http.Get(site)

	if err == nil {
		if resp.StatusCode == 200 {
			fmt.Println("Site", i+1, ":", site, "foi carregado com sucesso!")
			logRegister(site, true)
		} else {
			fmt.Println("Site", i+1, ":", site, "não foi carregado!, Status: ", resp.StatusCode)
			logRegister(site, false)
		}
	} else {
		fmt.Println("Site", i+1, ":", "não foi possivel concluir a solicitação! erro: ", err)
		logRegister(site, false)
	}
}

// clearTerminal clears the terminal screen.
//
// No parameters.
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run() //nolint:errcheck
}

// readFile reads the contents of the "sites.txt" file and returns a slice of strings
// containing each line of the file. If there is an error reading the file, an empty
// slice is returned and an error message is logged.
//
// No parameters.
// Returns a slice of strings.
func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		log.Println("ocorreu um erro ao ler o arquivo! erro:", err)
		return sites
	}
	dataFile := bufio.NewReader(file)
	for {
		site, err := dataFile.ReadString('\n')
		site = strings.TrimSpace(site)
		sites = append(sites, site)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("ocorreu um erro ao ler o arquivo! erro:", err)
			return sites
		}
	}
	file.Close()
	return sites

}

// logRegister logs the registration of a site with its status.
//
// site: a string representing the site being registered
// status: a boolean indicating the status of the site
func logRegister(site string, status bool) {
	os.Open("log.txt")
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("ocorreu um erro ao registrar o log! erro:", err)
		return
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

// printLogs prints the logs from a file named "log.txt".
//
// No parameters.
// No return types.
func printLogs() {

	file, err := os.Open("log.txt")
	if err != nil {
		log.Println("ocorreu um erro ao ler o arquivo! erro:", err)
		return
	}
	dataFile := bufio.NewReader(file)
	for {
		logs, err := dataFile.ReadString('\n')
		logs = strings.TrimSpace(logs)
		fmt.Println(logs)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("ocorreu um erro ao ler o arquivo! erro:", err)
			return
		}
	}
	file.Close()

	
	finalMessage()

}
