package ear

import (
	"bufio"
	"fmt"
	"html"
	"net/http"

	"github.com/santoshbachar/navik-cli/constant"
)

func Serve(args []string) {
	switch args[0] {
	case "start":
		fmt.Println("Starting Navik engines...")
	case "stop":
		fmt.Println("Stopping Navik engines...")
	case "status":
		fmt.Println("Status")
		handleStatus(args[0:])
	case "reload":
		fmt.Println("reload")
		handleReload()

	}
}

func issue(cmd string) {
	resp, err := http.Get(constant.HOME_ADDR + cmd)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("status command issued to Navik")
	fmt.Println(resp.Body)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func handleStatus(args []string) {
	switch args[0] {
	case "test":
		fmt.Println("Will check for the status of test")
	default:
		fmt.Println("Will check all status")
		issue("status")
	}
}

func handleReload() {
	issue("reload")
}

func Test() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})
}

func Reload() {

	resp, err := http.Get("/reload")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("reload command issued to Navik")

}
