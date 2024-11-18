package main

import (
	"context"
	"fmt"

	// "io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	
	fmt.Println(os.Getenv("ENDPOINT"))
	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESS_KEY")
	secretAccessKey := os.Getenv("SECRET_KEY")
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}
	fmt.Println("MinIO client initialized successfully")

	// Check if bucket exists
	bucketExists, err := minioClient.BucketExists(context.Background(), "samplebucket2")
	if err != nil {
		log.Fatalf("Error checking if bucket exists: %v", err)
	}
	if !bucketExists {
		// Create bucket if it doesn't exist
		err := minioClient.MakeBucket(context.Background(), "samplebucket2", minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			log.Fatalf("Error creating bucket: %v", err)
		}
		fmt.Printf("Bucket %s created successfully\n", "samplebucket")
	}

	// // list bucket
	// buckets, err := minioClient.ListBuckets(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, bucket := range buckets {
	// 	fmt.Println(bucket)
	// }

	// // remove bucket
	// err = minioClient.RemoveBucket(context.Background(), "samplebucket")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("Bucket %s removed successfully\n", "samplebucket")

	// file, err := os.Open("sample.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()

	// fileStat, err := file.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// uploadInfo, err := minioClient.PutObject(context.Background(), "samplebucket", "mineobject", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Successfully uploaded bytes: ", uploadInfo)

	// _, err = minioClient.FPutObject(context.Background(), "samplebucket", "sample.txt", "C:/Users/ShyamKuntal/Desktop/Digivate/minio/sample.txt", minio.PutObjectOptions{})
	// if err != nil {
	// 	log.Fatalf("Error uploading file: %v", err)
	// }
	// fmt.Println("File uploaded successfully!")

	// ctx, cancel := context.WithCancel(context.Background())

	// defer cancel()

	// objectCh := minioClient.ListObjects(ctx, "samplebucket", minio.ListObjectsOptions{
	// 	Prefix:    "myprefix",
	// 	Recursive: true,
	// })
	// for object := range objectCh {
	// 	if object.Err != nil {
	// 		fmt.Println(object.Err)
	// 		return
	// 	}
	// 	fmt.Println(object)
	// }

	// // Get Object
	// object, err := minioClient.GetObject(context.Background(), "samplebucket", "sample.txt", minio.GetObjectOptions{})

	// defer object.Close()

	// localfile, err := os.Create("C:/Users/ShyamKuntal/Desktop/Digivate/sample.txt")
	// if err != nil {
	// 	log.Fatalf("Error creating local file: %v", err)
	// }
	// defer localfile.Close()

	// _, err = io.Copy(localfile, object)
	// if err != nil {
	// 	log.Fatalf("Error downloading file: %v", err)
	// }

	// fmt.Printf("File '%s' downloaded successfully to '%s'\n", "objectName", "localFilePath")

	// // Delete file
	// err = minioClient.RemoveObject(context.Background(), "samplebucket", "mineobject", minio.RemoveObjectOptions{})
	// if err != nil {
	// 	log.Fatalf("Error deleting object: %v", err)
	// }

	// fmt.Printf("File '%s' deleted successfully from bucket '%s'\n", "objectName", "bucketName")

}
