/*
Copyright © 2024 Rohit Kumar <kumar1rohit@outlook.com>
*/
package utils

import (
	"fmt"
	"net"
	"net/mail"
	"strings"
	"time"
)

const (
	connectionTimeOut = 10 * time.Second
	defaultFromEmail  = "test@example.com"
)

func CheckEmail(toEmail string) {

	_, err := mail.ParseAddress(toEmail)

	if err != nil {
		fmt.Printf("❌ Not a valid email address: %v\n", err)
		return
	}

	domain := strings.Split(toEmail, "@")[1]

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("❌ Failed to lookup MX records, %v\n", err)
		return
	}

	if len(mxRecords) == 0 {
		fmt.Println("❌ No MX records found, email domain doesn't exist")
		return
	}

	err = verifyEmailWithSMTP(mxRecords[0].Host, toEmail)
	if err != nil {
		fmt.Printf("❌ %v doesn't exist. Typo?\n", toEmail)
		return
	}
	fmt.Printf("✅ %v is reachable. All set!\n", toEmail)
}
