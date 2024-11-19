```bash
Install dependencies
'go mod tidy'
```

```bash
Set the Environment Variables
ENDPOINT="localhost:{PORT}"
SECRET_KEY=""
ACCESS_KEY=""
```

```bash
Run minio Server
'.\minio.exe server {Path/To/Minio.exe} --console-address :9001'
```

```bash
Run the Code
'go run main.go'
```

```bash
create bucket
curl --location --request POST 'http://localhost:8000/bucket/create/{bucketname}'
```

```bash
List bucket
curl --location --request GET 'http://localhost:8000/bucket/list'
```

```bash
Put Object
curl --location --request POST 'http://localhost:8000/object/upload/{BucketName}/{FileName}'
--header 'Content-Type: text/plain'
--data-binary '@/C:/Users/ShyamKuntal/Desktop/ss.txt'
```

```bash
Get Object
curl --location --request GET 'http://localhost:8000/object/download/{BucketName}/{FileName}'
```

```bash
DELETE Object
curl --location --request DELETE 'http://localhost:8000/object/delete/{BucketName}/{FileName}'
```
