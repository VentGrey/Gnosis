// Gnosis is a simple Go CLI tool for generating definition files for
// typescrtipt from PocketBase json input.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gnosis/lib"
)

func main() {
	var inputFile string
	var outputFile string

	// Flag parsing
	flag.StringVar(&inputFile, "input", "", "Input file to parse")
	flag.StringVar(&inputFile, "i", "", "Input file to parse (shorthand)")
	flag.StringVar(&outputFile, "output", "", "Output file to write to")
	flag.StringVar(&outputFile, "o", "", "Output file to write to (shorthand)")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("Example:")
		fmt.Printf("  %s -i input.json -o output.d.ts\n", os.Args[0])
		fmt.Printf("%s prints the typescript definition to stdout if no output file is specified.\n", os.Args[0])
	}
	flag.Parse()


	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", inputFile)
		fmt.Println(err)
		os.Exit(1)
	}

	var data map[string]interface{}

	err = json.Unmarshal(file, &data)

	if err != nil {
		fmt.Println("Error parsing json")
		fmt.Println(err)
		os.Exit(1)
	}

	items, ok := data["items"].([]interface{})

	if !ok || len(items) == 0 {
		fmt.Println("Invalid item structure")
		os.Exit(1)
	}

	firstItem, ok := items[0].(map[string]interface{})

	if !ok {
		fmt.Println("Invalid item structure")
		os.Exit(1)
	}

	collectionName, ok := firstItem["@collectionName"].(string)

	if !ok {
		fmt.Println("Missing or invalid @collectionName")
		os.Exit(1)
	}

	tsInterface := lib.GenerateTypeScriptInterface(firstItem, collectionName)

	if outputFile == "" {
		fmt.Println(tsInterface)
	} else {
		err = ioutil.WriteFile(outputFile, []byte(tsInterface), 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", outputFile)
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
