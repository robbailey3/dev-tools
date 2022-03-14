package tcp

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"time"
)

var tcp = &cobra.Command{
	Use: "tcp",
}

var scan = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		openPorts := PortScan()
		for _, i := range openPorts {
			fmt.Println(fmt.Sprintf("Port %d is open", i))
		}
	},
}

func Commands() *cobra.Command {
	tcp.AddCommand(scan)

	return tcp
}

func PortScan() []int {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	//var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go checkPorts("robbailey3.co.uk", ports, results)
	}

	go func() {
		for i := 75; i < 92; i++ {
			//wg.Add(1)
			ports <- i
		}
	}()
	for i := range results {
		openPorts = append(openPorts, i)
	}

	return openPorts
}

func checkPorts(path string, ports chan int, results chan int) {
	for port := range ports {
		if checkPort(path, port) {
			results <- port
		}
	}
}

func checkPort(path string, port int) bool {
	address := fmt.Sprintf("%s:%d", path, port)

	fmt.Println(fmt.Sprintf("Checking port %d", port))

	_, err := net.DialTimeout("tcp", address, 5*time.Second)

	if err != nil {
		fmt.Println(fmt.Sprintf("Port %d is closed", port))
		return false
	}
	fmt.Println(fmt.Sprintf("Port %d is open", port))
	return true
}
