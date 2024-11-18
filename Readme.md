```bash
Install dependencies
curl 'go mod tidy'
```

```bash
Run the Code
curl 'go run main.go'
```

```bash
create bucket
curl --location --request POST 'http://localhost:8000/bucket/create/{bucketname}'
```

```bash
List bucket
curl --location 'http://localhost:8000/bucket/list'
```

```bash
Put Object
curl --location 'http://localhost:8000/object/upload/{BucketName}/{FileName}' \
--header 'Content-Type: text/plain' \
--data-binary '@/C:/Users/ShyamKuntal/Desktop/ss.txt'
```

```bash
Get Object
curl --location 'http://localhost:8000/object/download/{BucketName}/{FileName}'
```

```bash
DELETE Object
curl --location 'http://localhost:8000/object/delete/{BucketName}/{FileName}'
```
