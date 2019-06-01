package taskservice

import "fmt"

func ExampledbHost() {
	port := 27017
	for _, c := range [][]string{
		[]string{"mongo-0.mongo"},
		[]string{"mongo-0.mongo", "mongo-1.mongo"},
		[]string{"localhost", "127.0.0.1", "192.0.0.1"},
	} {
		fmt.Printf("%#v ==> %#v\n", c, dbHost(port, c...))
	}

	// Output:
	// "[]string{"mongo-0.mongo"}" ==> "mongodb://mongo-0.mongo:27017"
	// "[]string{"mongo-0.mongo", "mongo-1.mongo"}" ==> "mongodb://mongo-0.mongo,mongo-1.mongo:27017"
	// "[]string{"localhost", "127.0.0.1", "192.0.0.1"}," ==> "mongodb://localhost,127.0.0.1,192.0.0.1:27017"
}
