package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nxm/crtsh-checker/crtlog"
	"github.com/nxm/crtsh-checker/parser"
)

func showUsage() {
	fmt.Println("This tool shows the result of crt.sh")
	fmt.Println("Option:")
	fmt.Println("  -q Query")
	fmt.Println("  -i Min Cert ID")
	fmt.Println("  -a Alive domains")
	fmt.Println("  -cn Common Name")
	fmt.Printf("Usage: %s -q example.com\n", os.Args[0])
	os.Exit(0)
}

func run() error {
	var (
		query          string
		certID         int
		onlyDomainFlag bool
		commonName     string
		onlyAlive      bool
	)
	flag.StringVar(&query, "q", "", "Query String")
	flag.BoolVar(&onlyDomainFlag, "o", false, "Print only domains")
	flag.IntVar(&certID, "i", 0, "Min Cert ID")
	flag.StringVar(&commonName, "cn", "", "Query string for common name")
	flag.BoolVar(&onlyAlive, "a", false, "Print only alive domains")
	flag.Parse()
	if query == "" && certID == 0 && commonName == "" {
		showUsage()
	}

	if query != "" {
		err := crtlog.QueryCrt(query, onlyDomainFlag, onlyAlive)
		if err != nil {
			return err
		}
	} else if certID != 0 {
		fmt.Printf("CertID: %d\n", certID)

		err := parser.ParseCTLog(certID)
		if err != nil {
			return err
		}
	} else {
		err := crtlog.SearchCommon(commonName, onlyDomainFlag)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
