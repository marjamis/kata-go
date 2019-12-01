package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ddbTableColumns = [9]string{
	"Column1",
	"Column2",
	"Column3",
	"Column4",
	"Column5",
	"Column6",
	"Column7",
	"Column8",
	"Column9",
}

var svc = dynamodb.New(session.New(&aws.Config{
	Region: aws.String(os.Getenv("DDB_REGION")),
}))

func main() {
	data, _ := os.Open(os.Getenv("CSV_FILE"))
	r := csv.NewReader(data)

	// Note: A simple hack to ensure the header fields of the first line are skipped as the csv package doesn't seem to have this functionality in built.
	_, err := r.Read()
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	var records []map[string]*dynamodb.AttributeValue
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		records = append(records, recordToItem(record))
	}

	// for i := range records {
	// 	putItem(records[i])
	// }

	batchWorkflow(records)
}
