package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/joho/godotenv"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	awsPublic := os.Getenv("AWS_PUBLIC_KEY")
	awsSecret := os.Getenv("AWS_SECRET_KEY")

	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			awsPublic,
			awsSecret,
			"",
		),
	})
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(session)
	s3Bucket = "font-gg"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 10)
	for {
		files, err := dir.ReadDir(1)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}) {
	defer wg.Done()

	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)
	file, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Failed to open file %s, %v\n", completeFileName, err)
		<-uploadControl // release the control
		return
	}
	defer file.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Failed to upload file %s, %v\n", completeFileName, err)
		return
	}
	<-uploadControl // release the control
	fmt.Printf("File %s uploaded successfully\n", completeFileName)
}
