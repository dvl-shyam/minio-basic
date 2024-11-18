package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/minio/minio-go/v7"
)

type response struct {
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

func respondJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response{
		Message: message,
		Data:    data,
	})
}

func getPathParam(url string, prefix string) (string, error) {
	if !strings.HasPrefix(url, prefix) {
		return "", fmt.Errorf("invalid path, must start with %s", prefix)
	}
	return strings.TrimPrefix(url, prefix), nil
}

// create bucket
func createBucketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bucketName")
	bucketName, err := getPathParam(r.URL.Path, "/bucket/create/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	fmt.Println(bucketName)
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Error creating bucket", err.Error())
		return
	}
	respondJSON(w, http.StatusOK, "Bucket created successfully", nil)
}

// list all buckets
func listBucketsHandler(w http.ResponseWriter, r *http.Request) {
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Error listing buckets", err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "Buckets listed successfully", buckets)
}

// upload file
// func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
// 	params, err := getPathParam(r.URL.Path, "/object/upload/")
// 	if err != nil {
// 		respondJSON(w, http.StatusBadRequest, err.Error(), nil)
// 		return
// 	}
// 	parts := strings.SplitN(params, "/", 2)
// 	if len(parts) < 2 {
// 		respondJSON(w, http.StatusBadRequest, "URL must include bucketName/objectName", nil)
// 		return
// 	}
// 	bucketName := parts[0]
// 	objectName := parts[1]
// 	fmt.Println(bucketName, objectName)

// 	//file path is provided in the body
// 	filePath := r.URL.Query().Get("filePath")
// 	if filePath == "" {
// 		respondJSON(w, http.StatusBadRequest, "filePath is required", nil)
// 		return
// 	}

// 	_, err = minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
// 	if err != nil {
// 		respondJSON(w, http.StatusInternalServerError, "Error uploading file", err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, "File uploaded successfully", nil)
// }
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	params, err := getPathParam(r.URL.Path, "/object/upload/")
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	parts := strings.SplitN(params, "/", 2)
	if len(parts) < 2 {
		http.Error(w, "URL must include bucketName/objectName", http.StatusBadRequest)
		return
	}

	bucketName := parts[0]
	objectName := parts[1]
	
	fileData := r.Body
	defer fileData.Close()

	uploadInfo, err := minioClient.PutObject(context.Background(), bucketName, objectName, fileData, -1, minio.PutObjectOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading object: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	response := map[string]interface{}{
		"message":     "File uploaded successfully",
		"bucketName":  bucketName,
		"objectName":  objectName,
		"uploadedSize": uploadInfo.Size,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler to list objects in a bucket
// func listObjectsHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		respondJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
// 		return
// 	}

// 	bucketName, err := getPathParam(r.URL.Path, "/object/list/")
// 	if err != nil {
// 		respondJSON(w, http.StatusBadRequest, err.Error(), nil)
// 		return
// 	}

// 	objectCh := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
// 		Recursive: true,
// 	})

// 	objects := []string{}
// 	for object := range objectCh {
// 		if object.Err != nil {
// 			respondJSON(w, http.StatusInternalServerError, "Error listing objects", object.Err.Error())
// 			return
// 		}
// 		objects = append(objects, object.Key)
// 	}
// 	respondJSON(w, http.StatusOK, "Objects listed successfully", objects)
// }

// Handler to download an object
func downloadObjectHandler(w http.ResponseWriter, r *http.Request) {

	params, err := getPathParam(r.URL.Path, "/object/download/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	parts := strings.SplitN(params, "/", 2)
	if len(parts) < 2 {
		respondJSON(w, http.StatusBadRequest, "URL must include bucketName/objectName", nil)
		return
	}

	bucketName := parts[0]
	objectName := parts[1]

	object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Error downloading object", err.Error())
		return
	}
	defer object.Close()

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", objectName))
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = io.Copy(w, object)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Error writing object to response", err.Error())
	}
}

// delete object
func deleteObjectHandler(w http.ResponseWriter, r *http.Request) {

	params, err := getPathParam(r.URL.Path, "/object/delete/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	parts := strings.SplitN(params, "/", 2)
	if len(parts) < 2 {
		respondJSON(w, http.StatusBadRequest, "URL must include bucketName/objectName", nil)
		return
	}

	bucketName := parts[0]
	objectName := parts[1]

	err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "Error deleting object", err.Error())
		return
	}
	respondJSON(w, http.StatusOK, "Object deleted successfully", nil)
}
