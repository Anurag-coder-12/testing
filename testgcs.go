package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func main() {

	bucketName := `tathaproj-1-bucket-2`
	projectID := "tathaproj-1"
	filename := "ic_sample.csv"

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer client.Close()
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			//return err
		}
		fmt.Println("Name : ", battrs.Name)
	}

	bkt := client.Bucket(bucketName)
	attrs, err := bkt.Attrs(ctx)
	if err != nil {
		// TODO: Handle error.
		fmt.Println(err)
	}
	fmt.Printf("bucket %s, created at %s, is located in %s with storage class %s\n",
		attrs.Name, attrs.Created, attrs.Location, attrs.StorageClass)

	rc, err := bkt.Object(filename).NewReader(ctx)
	if err != nil {
		//return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	//data, err := ioutil.ReadAll(rc)
	if err != nil {
		//return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	//fmt.Printf("%v", string(data[:]))

	r := csv.NewReader(rc)
	var intlist []IIPIntegration

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(record)

		intlist = append(intlist, IIPIntegration{Intno: record[1], Gitloc: record[2], Composite: record[3]})
	}

	intJson, _ := json.Marshal(intlist)

	fmt.Println(string(intJson))
}

// IIPIntegration blah
type IIPIntegration struct {
	Intno     string `json:"intno"`
	Gitloc    string `json:"gitloc"`
	Composite string `json:"composite"`
}
