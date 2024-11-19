package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func main() {
	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESS_KEY")
	secretAccessKey := os.Getenv("SECRET_KEY")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v %v", err, minioClient)
	}
	fmt.Println("MinIO client initialized successfully")

	http.HandleFunc("POST /bucket/create/", createBucketHandler)
	http.HandleFunc("GET /bucket/list", listBucketsHandler)
	http.HandleFunc("POST /object/upload/", uploadFileHandler)
	http.HandleFunc("GET /object/download/", downloadObjectHandler)
	http.HandleFunc("DELETE /object/delete/", deleteObjectHandler)
   
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        fmt.Fprintln(w, "Minio Project")
    })
	port := "8000"
	fmt.Printf("Server started on port %s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}


