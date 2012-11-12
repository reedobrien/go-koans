package go_koans

import (
	"bytes"
	// "fmt"
	"io"
	"log"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		_, err := io.Copy(out, in)
		if err != nil {
			log.Fatal(err)
		}
		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		_, err := io.CopyN(out, in, 5)
		if err != nil {
			log.Fatal(err)
		}

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
