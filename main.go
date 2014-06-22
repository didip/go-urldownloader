package main

import (
    "os"
    "strings"
    "io/ioutil"
    "net/http"
    "github.com/jcelliott/lumber"
    "launchpad.net/goamz/aws"
    "launchpad.net/goamz/s3"
)

func StripURLProtocol(url string) string {
    url = strings.Replace(url, "https://", "", -1)
    url = strings.Replace(url, "http://", "", -1)

    return url
}

//
// Environment Variables:
// AWS_ACCESS_KEY_ID
// AWS_SECRET_ACCESS_KEY
// S3_BUCKET
// HTTP_PORT
//
func main() {
    auth, err := aws.EnvAuth()
    if err != nil {
        panic(err.Error())
    }

    httpPort     := os.Getenv("HTTP_PORT")
    s3BucketName := os.Getenv("S3_BUCKET")
    s3Conn       := s3.New(auth, aws.USEast)
    s3Bucket     := s3Conn.Bucket(s3BucketName)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            w.Write([]byte("Status: UP"))
        }

        if r.Method == "POST" {
            url    := r.FormValue("url")

            if url != "" {
                response, err := http.Get(url)

                if err != nil {
                    lumber.Error("%s", err)

                } else {
                    defer response.Body.Close()

                    content, err := ioutil.ReadAll(response.Body)

                    if err != nil {
                        lumber.Error("%s", err)

                    } else {
                        contentType := http.DetectContentType(content)

                        w.Header().Set("Content-Type", contentType)
                        w.Write(content)

                        go func(url string, contentType string, content []byte) {
                            s3Item := StripURLProtocol(url)

                            err = s3Bucket.Put(s3Item, content, contentType, s3.PublicRead)

                            if err != nil {
                                lumber.Error("%s", err)
                            }
                        }(url, contentType, content)
                    }
                }

            } else {
                lumber.Error("URL is empty.")
                w.Write([]byte(""))
            }
        }
    })

    http.ListenAndServe(":" + httpPort, nil)
}
