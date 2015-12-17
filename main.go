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
	grocsCategories := readFile("grocery_categories.txt", product.ParseCategoriesFile)
	if grocsCategories != nil {
		fmt.Println("Start================Testing Groceries Categories---------\n\n")
		grocsCategories.PrintValues()
		fmt.Println("================Testing Groceries Categories---------End\n\n")
	}
	/*
		ideasCategories := readFile("ideas_categories.txt", product.ParseCategoriesFile)
		if ideasCategories != nil {
			fmt.Println("Start================Testing Ideas Categories---------\n\n")
			ideasCategories.PrintValues()
			fmt.Println("================Testing Ideas Categories---------End\n\n")
		}
	*/

	if products := readFile("product_data.txt", product.ParseProductFile); products != nil {
		fmt.Println("Start===========Testing Products-----------\n\n")
		// products.PrintValues()
		prods, ok := products.(product.Products)
		if ok {
			groc_categories_map, ok := grocsCategories.(product.MapCategories)
			if ok {

				prods.FlattenGroceriesCategories(groc_categories_map)
			}
		}

		// prods.PrintValues()
		prods.PrintJSON()
		fmt.Println("===========Testing Products-----------End\n\n")
	}

	// product.PrintCategoryMap(cat_map)
}
