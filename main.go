package main

import (
	"fmt"
	"log"

	"go-test/greetings"
)

func main() {
	log.SetPrefix("greeting:")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.Lmsgprefix)

	msg, err := greetings.Hello("zsy")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

	msgs, err := greetings.Hellos([]string{"test", "zsy"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msgs)

}
