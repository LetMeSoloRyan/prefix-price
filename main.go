package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PrefixPrices map[string]float64
type CompanyPrices map[string]PrefixPrices
type Response struct {
	Company string
	Prefix  string
	Price   float64
}

func main() {
	var companyPrices CompanyPrices := make(CompanyPrices)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">>> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		processCommand(strings.TrimSpace(cmdString), &companyPrices)
	}
}

func processCommand(cmdString string, companyPrices *CompanyPrices) {
	args := strings.Fields(cmdString)

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "INSERT":
		if len(args) != 4 {
			fmt.Println("INSERT requires exactly 3 arguments")
			return
		}

		err := insert(args[1], args[2], args[3], companyPrices)
		if err != nil {
			fmt.Println(err.Error())
		}

	case "QUERY":
		if len(args) != 2 {
			fmt.Println("QUERY requires exactly 1 argument")
			return
		}

		result, err := query(args[1], companyPrices)
		if err != nil {
			fmt.Printf("%s NA\n", args[1])
		} else {
			fmt.Printf("%s %s %s %.2f\n", args[1], result.Company, result.Prefix, result.Price)
		}

	default:
		fmt.Println("Please enter INSERT or QUERY commands only")
	}
}

func insert(company, prefix, priceStr string, companyPrices *CompanyPrices) error {
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return errors.New("price must be a valid floating-point number")
	}

	if _, exists := (*companyPrices)[company]; !exists {
		(*companyPrices)[company] = make(PrefixPrices)
	}

	(*companyPrices)[company][prefix] = price
	return nil
}

func query(phone string, companyPrices *CompanyPrices) (*Response, error) {
	var bestResponse *Response

	for companyName, prefixPrices := range *companyPrices {
		for prefix, price := range prefixPrices {
			if strings.HasPrefix(phone, prefix) {
				if bestResponse == nil || len(prefix) > len(bestResponse.Prefix) || (len(prefix) == len(bestResponse.Prefix) && price < bestResponse.Price) {
					bestResponse = &Response{Company: companyName, Prefix: prefix, Price: price}
				}
			}
		}
	}

	if bestResponse == nil {
		return nil, errors.New("no matching prefix found")
	}

	return bestResponse, nil
}
