package main

import (
	// "github.com/sim4life/groc_cat"
	"github.com/sim4life/product"

	"bufio"
	"fmt"
	"os"
)

func readFile(file_name string, parseFile func(*bufio.Scanner) product.PrintStructVals) product.PrintStructVals {
	prodfile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer prodfile.Close()
	scanner := bufio.NewScanner(prodfile)
	// scanner.Split(bufio.ScanLines)
	//TODO: perform error handling
	return parseFile(scanner)

}

func main() {
	/*
	   c := elastigo.NewConn()
	   log.SetFlags(log.LstdFlags)
	   flag.Parse()

	   // Trace all requests
	   c.RequestTracer = func(method, url, body string) {
	     log.Printf("Requesting %s %s", method, url)
	     log.Printf("Request body: %s", body)
	   }
	*/
	//fmt.Println("host = ", *host)

	//TODO: perform error handling
	// cat_map := readFile("grocery_categories.txt", product.ParseCategoriesFile)
	// readFile("ideas_categories.txt", parseIdeasCategoriesFile)
	if products := readFile("product_data.txt", product.ParseProductFile); products != nil {
		products.PrintValues()
		fmt.Println("Testing")
	}
	// product.PrintCategoryMap(cat_map)
}
