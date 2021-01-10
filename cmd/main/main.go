package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
)

type Information struct {
	Request *http.Request
	System  System
}

type System struct {
	OS  string
	CPU int
	IP  string
}

func main() {
	// Setup system information
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("Error retrieving network addresses")
	}

	var ip string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}

	sys := System{OS: runtime.GOOS, CPU: runtime.NumCPU(), IP: ip}

	// Setup handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("tpl/index.html")
		if err != nil {
			log.Println("Error occurred when parsing template")
			_, _ = fmt.Fprintf(w, "An unfortunate error occured")
			return
		}

		info := Information{Request: r, System: sys}
		if err = tpl.Execute(w, info); err != nil {
			log.Println("Error occurred when executing template")
			_, _ = fmt.Fprintf(w, "An unfortunate error occured")
			return
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Everything is working as intended"))
	})

	// Start server
	address := "0.0.0.0:" + os.Getenv("PORT")

	log.Println("Starting web server on port 80")
	log.Fatal(http.ListenAndServe(address, nil))
}
