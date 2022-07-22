package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Check : domainName,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error : Can't read from input: %v\n", err)
	}
}
                                      
func checkDomain(domain_name string) {
	var has_MX, has_SPF, has_DMARC bool
	var spf_Record, dmarc_Record string

	mxRecords, err := net.LookupMX(domain_name)

	if err != nil {
		log.Printf("Error : %v\n",err)
	}

	if len(mxRecords) > 0 {
		has_MX = true
	}

	txtRecords, err := net.LookupTXT(domain_name)

	if err != nil {  
		log.Printf("Error : %v\n, err")
	}


	for _, record := range txtRecords {
		
		if strings.HasPrefix(record,"v = spf1") {
			has_SPF =  true
			spf_Record = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain_name)
 
	if err != nil {
		log.Printf("Error %v\n, err")
	}

	for _, record := range dmarcRecords {

		if strings.HasPrefix(record,"v = DMARC1") {
			has_DMARC = true
			dmarc_Record =  record
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain_name, has_MX, has_SPF, spf_Record , has_DMARC, dmarc_Record)
 }