
# Mini Image Processor

  

A web application that does three things through its three APIs

  

-  **Convert: Given a file, check if it is a valid PNG file. If it is a valid PNG file convert it to JPEG.**

-  **Resize: Given a file and some parameters, check if it is a valid PNG file. Resize according to the parameters if it is.**

-  **Compress: Given a file, check if it is a valid PNG file. Compress it with compression level 9 if it is.**

# Tech Stack

- Golang 1.19

- Echo

- GoCV

- OpenCV 4.7.0

  

# Code Structure

```
├── Makefile
├── README.md
├── api
│   └── http
│       ├── errorjson.go
│       ├── httpapis.go
│       ├── httpapis_common_test.go
│       └── httpapis_test.go
├── cmd
│   └── app
│       ├── config.yaml
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── conf
│   │   ├── config.go
│   │   └── init.go
│   ├── data
│   │   └── model
│   │       └── resize.go
│   └── pkg
│       ├── core
│       │   ├── adapter
│       │   │   ├── commonadapter
│       │   │   │   └── commonadapter.go
│       │   │   ├── gocvadapter
│       │   │   │   └── gocvadapter.go
│       │   │   ├── imageadapter
│       │   │   │   └── imageadapter.go
│       │   │   ├── ioadapter
│       │   │   │   └── ioadapter.go
│       │   │   └── osadapter
│       │   │       └── osadapter.go
│       │   └── service
│       │       └── imageservice
│       │           ├── compress.go
│       │           ├── compress_test.go
│       │           ├── convert.go
│       │           ├── convert_test.go
│       │           ├── imageservice.go
│       │           ├── imageservice_common_test.go
│       │           ├── removefile.go
│       │           ├── resize.go
│       │           └── resize_test.go
│       ├── mocks
│       │   ├── mockcommonadapter
│       │   │   ├── mockcommonadapter.go
│       │   │   └── mockfileopener.go
│       │   ├── mockecho
│       │   │   └── mockcontext.go
│       │   ├── mockgocvadapter
│       │   │   └── mockgocvadapter.go
│       │   ├── mockimageadapter
│       │   │   └── mockimageadapter.go
│       │   ├── mockioadapter
│       │   │   └── mockioadapter.go
│       │   ├── mockosadapter
│       │   │   └── mockosadapter.go
│       │   └── mockother
│       │       └── mockwritecloser.go
│       ├── platform
│       │   ├── common
│       │   │   ├── common.go
│       │   │   └── common_test.go
│       │   ├── errors
│       │   │   └── errors.go
│       │   ├── gocv
│       │   │   └── gocv.go
│       │   ├── io
│       │   │   └── io.go
│       │   └── os
│       │       └── os.go
│       └── transport
│           └── http.go
├── main
├── mini-image-converter-apis.postman_collection.json
└── ubersnap_mini_image_processor_service.pdf

```

  

# Setup

  

## Installing Dependencies

  

- This is a Golang app. Download Golang here (https://go.dev/)

- This app uses GoCV. In order to use GoCV, you must install OpenCV 4.7.0 on your system. Follow the instructions in this link for more info: https://gocv.io/getting-started/

## Running the app

Once the dependencies have been installed, run this command on the terminal:

  

- make all

  

The app will be run locally on the port specified in the cmd/app/config.yaml file. The default port is 8008.

  

## Running the UnitTests

  

There are two options to run the unit test. Simply use one of these commands on the terminal:

  

- make test

- Runs the test without coverage

- make cover=true test

- Runs the test with coverage

  

## Additional Info

- To make it easy to call the APIs, a postman collection (mini-image-converter-apis.postman_collection.json) is provided. To find out how to import it, you can visit this link: https://learning.postman.com/docs/getting-started/importing-and-exporting/importing-and-exporting-overview/. Additionally, to find out more about postman, you can visit this link: https://www.postman.com/

- The tech design document is also available within this repository (ubersnap_mini_image_processor_service.pdf).

