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

            response, err := http.Get(url)

            if err != nil {
                lumber.Error("%s", err)

            } else {
                defer response.Body.Close()

                content, err := ioutil.ReadAll(response.Body)

                if err != nil {
                    lumber.Error("%s", err)

                } else {
                    s3Item      := StripURLProtocol(url)
                    contentType := http.DetectContentType(content)

                    err = s3Bucket.Put(s3Item, content, contentType, s3.PublicRead)

                    if err != nil {
                        lumber.Error("%s", err)
                    }
                }
            }
        }
    })

    http.ListenAndServe(":" + httpPort, nil)
}
