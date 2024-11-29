# Minio - Basic Setup

## API Requests


Install dependencies
```bash
'go mod tidy'
```


Set the Environment Variables
```bash
ENDPOINT="localhost:{PORT}"
SECRET_KEY=""
ACCESS_KEY=""
```


Run minio Server
```bash
'.\minio.exe server {Path/To/Minio.exe} --console-address :9001'
```


Run the Code
```bash
'go run main.go'
```


create bucket
```bash
curl --location --request POST 'http://localhost:8000/bucket/create/{bucketname}'
```

List bucket
```bash
curl --location --request GET 'http://localhost:8000/bucket/list'
```


Put Object
```bash
curl --location --request POST 'http://localhost:8000/object/upload/{BucketName}/{FileName}'
--header 'Content-Type: text/plain'
--data-binary '@/C:/Users/ShyamKuntal/Desktop/ss.txt'
```


Get Object
```bash
curl --location --request GET 'http://localhost:8000/object/download/{BucketName}/{FileName}'
```


DELETE Object
```bash
curl --location --request DELETE 'http://localhost:8000/object/delete/{BucketName}/{FileName}'
```
