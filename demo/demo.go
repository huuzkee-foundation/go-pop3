
package main

import (
	//"errors"
	//"fmt"
	//"net"
	"github.com/vonwenm/go-pop3"
	//"strconv"
	//"strings"
	"log"
)	
	
	
func main() {	
	
	address := "mail.domaind.com:110"
	user := "name@domain.com"	
	pass := "pwd"
	
	client, err := pop3.Dial(address)

	if err != nil {
	    log.Fatalf("Error: %v\n", err)
	}

	defer func() {
	    client.Quit()
	    client.Close()
	}()

	if err = client.User(user); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	log.Printf("User: %v\n", user)
	    
	if err = client.Pass(pass); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	var count int
	var size uint64

	if count, size, err = client.Stat(); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	log.Printf("Count: %d, Size: %d\n", count, size)

	if count, size, err = client.List(6); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	log.Printf("Number: %d, Size: %d\n", count, size)

	var mis []pop3.MessageInfo

	if mis, err = client.ListAll(); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	for _, mi := range mis {
	    log.Printf("Number: %d, Size: %d\n", mi.Number, mi.Size)
	}

	var number int
	var uid string

	if number, uid, err = client.Uidl(6); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	log.Printf("Number: %d, Uid: %s\n", number, uid)

	if mis, err = client.UidlAll(); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	for _, mi := range mis {
	    log.Printf("Number: %d, Uid: %s\n", mi.Number, mi.Uid)
	}

	var content string

	if content, err = client.Retr(8); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	log.Printf("Content:\n%s\n", content)

	if err = client.Dele(6); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	if err = client.Noop(); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}

	if err = client.Rset(); err != nil {
	    log.Printf("Error: %v\n", err)
	    return
	}
	
}