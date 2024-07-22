/*
Copyright © 2024 Rohit Kumar <kumar1rohit@outlook.com>
*/
package utils

import (
	"fmt"
	"net"
	"net/smtp"
)

func verifyEmailWithSMTP(mxHost, toEmail string) error {
	ips, err := net.LookupIP(mxHost)
	if err != nil {
		return fmt.Errorf("failed to resolve IP for %v: %v", mxHost, err)
	}

	if len(ips) == 0 {
		return fmt.Errorf("no IP addresses found for %v", mxHost)
	}

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:25", ips[0].String()), connectionTimeOut)
	if err != nil {
		return fmt.Errorf("failed to connect to %v: %v", mxHost, err)
	}
	defer conn.Close()

	smtpClient, err := smtp.NewClient(conn, mxHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}
	defer smtpClient.Close()

	err = smtpClient.Hello("example.com")
	if err != nil {
		return fmt.Errorf("HELO failed: %v", err)
	}

	err = smtpClient.Mail(defaultFromEmail)
	if err != nil {
		return fmt.Errorf("MAIL FROM failed: %v", err)
	}

	err = smtpClient.Rcpt(toEmail)
	if err != nil {
		return fmt.Errorf("RCPT TO failed: %v", err)
	}

	// If we get here, the email is valid
	return nil
}
