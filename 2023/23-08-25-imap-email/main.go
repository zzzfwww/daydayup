package main

import (
	"fmt"
	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"log"
)

func main() {

	// Replace with your email and password
	email := "email"
	password := "password"

	// Connect to the server
	c, err := client.DialTLS("imap.exmail.qq.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	// Login
	if err := c.Login(email, password); err != nil {
		fmt.Println("Error logging in:", err)
		return
	}

	// Select mailbox
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		fmt.Println("Error selecting mailbox:", err)
		return
	}
	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	// Create search criteria for unseen messages
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag, imap.DraftFlag}
	// Search for unseen messages
	ids, err := c.Search(criteria)
	if err != nil {
		log.Fatal(err)
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(ids...)

	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Fatal("Server didn't returned message")
	}
	if msg.Envelope.From[0].HostName != "tapd.cn" {
		println("this email not tapd statistical")
		return
	}

	r := msg.GetBody(&section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// Print some info about the message
	header := mr.Header
	if date, err := header.Date(); err == nil {
		log.Println("Date:", date)
	}
	if from, err := header.AddressList("From"); err == nil {
		log.Println("From:", from)
	}
	if to, err := header.AddressList("To"); err == nil {
		log.Println("To:", to)
	}
	if subject, err := header.Subject(); err == nil {
		log.Println("Subject:", subject)
	}

	// Process each message's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(p.Body)
			strB := string(b)
			log.Println("Got text:", strB)
		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()
			log.Println("Got attachment: %v", filename)
		}
	}
}
