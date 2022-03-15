package tcp

import (
	"fmt"
	"github.com/robbailey3/dev-tools/ui"
	"net"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	path        string
	startPort   int
	endPort     int
	concurrency int
)

var tcp = &cobra.Command{
	Use: "tcp",
}

var scan = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		view := ui.New()
		view.Start()
		view.SetLoading(true)

	},
}

func Commands() *cobra.Command {
	tcp.PersistentFlags().StringVarP(&path, "path", "p", "robbailey3.co.uk", "The path to scan")
	tcp.PersistentFlags().IntVarP(&startPort, "start", "s", 1, "The start port")
	tcp.PersistentFlags().IntVarP(&endPort, "end", "e", 1024, "The end port")
	tcp.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", 10, "The number of concurrent connections")
	tcp.AddCommand(scan)

	return tcp
}

func PortScan() []int {
	ports := make(chan int, concurrency)
	results := make(chan int)
	var openPorts []int

	var wg sync.WaitGroup

	wg.Add(endPort - startPort)

	go func() {
		for i := startPort; i < endPort; i++ {
			ports <- i
		}
		close(ports)
	}()

	for i := 0; i < cap(ports); i++ {
		go func() {
			for port := range ports {
				if checkPort(path, port, &wg) {
					results <- port
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		openPorts = append(openPorts, port)
	}

	return openPorts
}

func checkPort(path string, port int, wg *sync.WaitGroup) bool {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", path, port)

	_, err := net.DialTimeout("tcp", address, 5*time.Second)

	return err == nil
}
