package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	//ConstMaxItemsPerBatchWriteItem - Minus one due to usage starting from 0
	ConstMaxItemsPerBatchWriteItem = 25 - 1
)

func recordToItem(record []string) map[string]*dynamodb.AttributeValue {
	entry := make(map[string]*dynamodb.AttributeValue)
	entry["service"] = &dynamodb.AttributeValue{
		S: aws.String(record[0]),
	}
	entry["date"] = &dynamodb.AttributeValue{
		S: aws.String(record[1]),
	}
	for i, stat := range record[2:] {
		entry[ddbTableColumns[i]] = &dynamodb.AttributeValue{
			N: aws.String(stat),
		}
	}

	return entry
}

func putItem(item map[string]*dynamodb.AttributeValue) {
	pi := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("DDB_TABLE")),
	}

	fmt.Println(pi)

	_, err := svc.PutItem(pi)
	if err != nil {
		log.Fatal(err)
	}
}

func batchWorkflow(items []map[string]*dynamodb.AttributeValue) {
	var prs []*dynamodb.PutRequest
	for _, item := range items {
		var pr dynamodb.PutRequest
		pr.Item = item
		prs = append(prs, &pr)
	}

	var wrs []*dynamodb.WriteRequest
	for _, pr := range prs {
		var wr dynamodb.WriteRequest
		wr.PutRequest = pr
		wrs = append(wrs, &wr)
	}

	//Using a count for simplicity to ensure only 25 PutRequests (the maximum) are in the one BatchWriteItem APi call
	count := 0
	wra := make(map[string][]*dynamodb.WriteRequest)
	for i, wr := range wrs {
		wra[os.Getenv("DDB_TABLE")] = append(wra[os.Getenv("DDB_TABLE")], wr)
		if count == ConstMaxItemsPerBatchWriteItem || (len(wrs) == (i + 1)) {
			input := &dynamodb.BatchWriteItemInput{
				RequestItems: wra,
			}
			fmt.Println(input)
			_, err := svc.BatchWriteItem(input)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("API call successfully made.")
			}

			//Resetting the number for the next run for the API BatchWriteItem API Call
			count = 0
			wra[os.Getenv("DDB_TABLE")] = nil
		}
		count++
	}
}
