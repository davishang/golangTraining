//Create a web app that uses a google cloud storage query (storage.Query) and
//demonstrates the functionality of the query's delimeter field

package main

import (
	"fmt"
	"net/http"
	//"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

func handler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	//[START get_default_bucket]
	ctx := appengine.NewContext(res)
	if bucket == "" {
		var err error
		if bucket, err = file.DefaultBucketName(ctx); err != nil {
			log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
			return
		}
	}
	//[END get_default_bucket]
	//Make Client
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
		return
	}
	defer client.Close()
	//Construct our object to interface with the cloud storage.

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(res, "Demo GCS Application running from Version: %v\n", appengine.VersionID(ctx))
	fmt.Fprintf(res, "Using bucket name: %v\n\n", bucket)

	d := &demo{
		w:      res,			//The response writer so it knows what to write to later if retrieving.
		ctx:    ctx,			//Current Context
		client: client,			//Current Client
		bucket: client.Bucket(bucket),	//Relevant Bucket, bucket is a global string that contains the default bucket location.
	}
	//The name of the file we will be creating.
	n := "demo-testfile-go"

	//Create the File and store in google cloud.
	d.createFile(n)
	d.listBucket()
}

func retrieve(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
		return
	}
	d := &demo{
		w:      res,
		ctx:    ctx,
		client: client,
		bucket: client.Bucket(bucket),
	}
	//The name of the file we will be grabbing and reading from.
	n := "demo-testfile-go"

	//Read the file from google cloud.
	d.readFile(n)
	d.listBucket()
}

//Empty base webpage. 
func index(res http.ResponseWriter, req *http.Request) {}

func init(){
	http.HandleFunc("/", index)
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/retrieve", retrieve)
}