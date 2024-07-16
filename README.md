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
            "size": 1234
        }
    },
    {
        "uid": "6895edf9-8265-4646-b235-432f2756a3dd",
        "data": {
            "name": "file1.zip",
            "size": 78386
        }
    },
    {
        "uid": "5cc9890f-888f-425a-b6e6-b459dac1e0dc",
        "data": {
            "name": "file2.jpg",
            "size": 2658743
        }
    }
]
```

## Technical Decisions

I decided to go with a simplistic approach to save time when deploying, which is why I ended up using a small `memcache` library to store saved data.
This was the first time setting up a cloud environment from scratch, so there was some figuring out to do.

I mostly had problems with permissions issues and problems switching between my personal and professional account when pushing my Docker image to GCP.
Otherwise I would've endedup using either PostgreSQL or SurrealDB.

## Requirements

- [x] Provide an HTTP endpoint to upload assets and retrieve a unique identifier for the uploaded file
- [x] Provide an HTTP endpoint to download an asset by its identifier. The original filename should be preserved when downloading
- Pick one of the following:
- - [ ] Add user-based access control to your files such that only the user who originally uploaded the file can access it
- - [ ] Add token-based access control to your files such that instead of the identifier, files can be accessed with a token that expires after a set period of time
- - [x] Add an endpoint that returns a list of all files in the system, their identifier, original filename, and the byte size of the file
- - [ ] Build a web page/app that provides a browser-based mechanism for using your
upload and download endpoints
- - [ ] Automate the setup of all infrastructure (servers, cloud services, code, etc) such that you could easily deploy a second complete, working copy of your app in a command or two 
- [x] Deploy this code to the cloud provider of your choice and provide us a link (it can be password-protected or otherwise authenticated if you’d like) (https://kwil-production.up.railway.app/files)
- [x] Provide a link to your code from a git-based source control platform or bundle your code together in a ZIP so we can review it
- Include a README that details:
- - [x] Which additional requirement you chose
- - [x] How to compile/build/run the code locally
- - [x] Where to access the deployed version of the project
- - [x] All design/architectural/technical decisions
