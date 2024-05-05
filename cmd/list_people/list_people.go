package main

import (
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/viktor-titov/protobuf-simple-example/addressbookpb"
	"google.golang.org/protobuf/proto"
)

// Main reads the entire address book from a file and prints all the
// information inside.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// [START unmarshal_proto]
	// Read the existing address book.
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	// [END unmarshal_proto]

	listPeople(os.Stdout, book)
}

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.PhoneType_PHONE_TYPE_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case pb.PhoneType_PHONE_TYPE_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case pb.PhoneType_PHONE_TYPE_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		default:
			fmt.Fprint(w, "  Mobile phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}
