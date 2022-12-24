package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain, hasMX, hasSPF, SPFRecord, hasDMARC, DMARCRecord\n ")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal("Error: could not read from input : %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPFRecord, DMARCRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true
			SPFRecord = record
			break
		}
	}

	dmarcRecord, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			DMARCRecord = record
			break
		}
	}

	fmt.Printf("%v %v %v %v %v %v", domain, hasMX, hasSPF, SPFRecord, hasDMARC, DMARCRecord)
}