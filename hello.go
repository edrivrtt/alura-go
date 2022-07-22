package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 5
const delay = 5

func main() {

	showIntro()

	for {
		showMenu()
		option := readOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Loading Logs...")
		case 0:
			fmt.Println("Exiting program")
			os.Exit(0)
		default:
			fmt.Println("Unknown option")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Edric"
	version := 1.1

	fmt.Println("Hi,", name)
	fmt.Println("This program is on version", version)
}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("You choose", option)

	return option
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://random-status-code.herokuapp.com/",
		"https://random-status-code.herokuapp.com/", "https://random-status-code.herokuapp.com/"}

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing instance", i, ":", site)
			testingInstance(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testingInstance(site string) {
	monitoringReturn, _ := http.Get(site)

	if monitoringReturn.StatusCode == 200 {
		fmt.Println(site, " is up and running!")
	} else {
		fmt.Println(site, " is down! Status code:", monitoringReturn.StatusCode)
	}
}
