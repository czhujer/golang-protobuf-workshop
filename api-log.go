package main

import (
	"bufio"
	"czhujer-golang-protobuf-workshop-II/grpcTest/api"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func writeItem(w io.Writer, r *api.Request) {
	fmt.Fprintln(w, "Question:", r.Question)
}

func listRequests(w io.Writer, requests *api.Requests) {
	for _, i := range requests.Items {
		writeItem(w, i)
	}
}

func promptForRequest(r io.Reader) (*api.Request, error) {
	// A protocol buffer can be created like any struct.
	a := &api.Request{}

	rd := bufio.NewReader(r)
	fmt.Print("Enter question: ")
	question, err := rd.ReadString('\n')
	if err != nil {
		return a, err
	}
	a.Question = strings.TrimSpace(question)

	return a, nil
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s API_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	fmt.Println("Using file: ", fname)

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	// [START marshal_proto]
	requests := &api.Requests{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, requests); err != nil {
		log.Fatalln("Failed to parse api log:", err)
	}

	// Add request
	r, err := promptForRequest(os.Stdin)
	if err != nil {
		log.Fatalln("Error with question:", err)
	}
	requests.Items = append(requests.Items, r)
	// [END_EXCLUDE]

	// Write the new requests back to disk.
	out, err := proto.Marshal(requests)
	if err != nil {
		log.Fatalln("Failed to encode api logs:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write api logs:", err)
	}
	// [END marshal_proto]

    // listing all requests
	fmt.Println("")
    fmt.Println("listing all requests...")
	listRequests(os.Stdout, requests)
}
