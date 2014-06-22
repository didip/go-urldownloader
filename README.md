This is a small HTTP daemon that downloads URL content and reupload to S3


### Why make this?

To test how decoupled a Task can be in [hadou](https://github.com/didip/hadou)


### Development Workflow

* Running daemon locally

    ```
    cd $GOPATH/src/github.com/didip/go-urldownloader
    HTTP_PORT=8080 S3_BUCKET=hadou-test-url-downloader AWS_ACCESS_KEY_ID=abc123 AWS_SECRET_ACCESS_KEY=qwerty go run main.go
    ```

* Submitting URL to scrape

    ```
    # Assume HTTP_PORT is 8080
    curl -d url=http://www.reddit.com/r/python.rss http://localhost:8080/
    ```