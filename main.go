package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	replica := 1 //number of camunda deployment pods, when the program starts
	pointa := 0  // instances at point 0
	for {
		//Loop to poll Camunda service
		initial := processcount() // process count when the program starts
		fmt.Println("initial processcount is :", initial)
		processesstarted := initial - pointa
		if processesstarted != 0 {
			procesesstartedperinstance := processesstarted / replica
			if procesesstartedperinstance >= 50 && replica < 4 {
				scaleup() // calling the replica scale up function
			} else if procesesstartedperinstance <= 20 && replica > 1 {
				scaledown() // calling the replica scale down function
			}
		} else {
			pointa = initial
		}

		time.Sleep(10 * time.Second)
	}
}
func processcount() int { // to get process count from camunda service
	type counter struct {
		Count int
	}
	fmt.Println("Starting the application...")
	//polling the service
	response, err := http.Get("http://192.168.99.101:30193/engine-rest/history/process-instance/count")
	fmt.Println("Got instance count", response)
	var processcounter counter
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))
		err := json.Unmarshal(data, &processcounter)
		if err != nil {
			log.Println(err)
		}

	}
	fmt.Println("Printing Count", processcounter.Count)
	return (processcounter.Count)
}

func scaleup() {
	fmt.Println("scaling up")
	cmd, err := exec.Command("/bin/bash", "/Users/siddarth.surana/Downloads/GOworkspace/src/github.com/Camunda-autoscalar/scaleup.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Printf("%s", cmd)

}

func scaledown() {
	cmd, err := exec.Command("/bin/bash", "/Users/siddarth.surana/Downloads/GOworkspace/src/github.com/Camunda-autoscalar/scaledown.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Printf("%s", cmd)
}
