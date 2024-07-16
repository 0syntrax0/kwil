# kwil

## Requirements

- Docker
- Port available `:8080`
- Postman (_optional_)

## How To Run

Run make command `make up` to build and start project.

## End Points

When working locally use `http://localhost:8080`, when using postman to test on live site use `https://kwil-production.up.railway.app/`.

### Upload File

- Method `POST`
- Url `https://kwil-production.up.railway.app/file`
- With `form-data` in the body and field `file`

#### Sample Payload

```curl
curl -X POST https://kwil-production.up.railway.app/file \
  -F "file=@/Users/LOCAL_USER_NAME/PATH_TO_FILE" \
  -H "Content-Type: multipart/form-data"
```

#### Sample Response

A `json` object with `.name` containing a `UUIDv4` associated with the file. 
To be used to retreive later.

```json
{
    "name": "4bbc9469-6682-4e0c-9990-9a9dfbfab6dc"
}
```

### Fetch Given File

- Method `GET`
- URL `https://kwil-production.up.railway.app/file/:uuid`

#### Sample Response

The fetched file will be forced downloaded into the browser.

### Fetch All Available Files

- Method `GET`
- Url `https://kwil-production.up.railway.app/file/`

#### Sample Response

```json
[
    {
        "uid": "4bbc9469-6682-4e0c-9990-9a9dfbfab6dc",
        "data": {
            "name": "d5h4ha1.png",
        }
    },
    {
        "uid": "6895edf9-8265-4646-b235-432f2756a3dd",
        "data": {
            "name": "file1.zip",
        }
    },
    {
        "uid": "5cc9890f-888f-425a-b6e6-b459dac1e0dc",
        "data": {
            "name": "file2.jpg",
        }
    }
]
```
