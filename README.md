# GoLang RESTful API Practise
## Welcome!
This simple RESTful API is for me to practise my GoLang skills and to try to get a spot in Wizeline's Go Bootcamp. Please check the steps below in order to run the project.

-----------------

## Getting Started

### Pre-requesites
- Make sure you've got [GO](https://golang.org/dl/) installed in your local machine. 
- Althogh not necessary, having an API tester such as [Postman](https://www.postman.com/downloads/) will be useful when testing this API.

### Running the app
1. Download or clone this repo into your local machine:

    ```bash
    git clone  https://github.com/emamex98/golang_rest_lab.git
    ```

1. Navigate to the directory where the repo was cloned:

    ```bash
    cd golang_rest_lab
    ```

1. Install all dependencies:

    ```go
    go get ./...
    ```

1. Start the server:

    ```go
    go run api.go
    ```

1. If everything works as expected, you should be able to reach [http://localhost:10000/api](http://localhost:10000/api) and get a _Hello, World!_ message:

    ```json
    {
        "Message": "Hello, world!"
    }
    ```


### Testing components
Some test cases were included in this project to test each endpoint. Here's how to run the tests.

1. Make sure the server is up and running, and that your in the project's directory.

    ```go
    go run api.go
    ```

1. To run all tests with verobse:

    ```go
    go test -v
    ```

-----------------
## Endpoints


### 1. Hello World
Root of the API, returns a classic _Hello, World!_ message in JSON format.

**URL** : `/api`

**Method** : `GET`

**Auth Required** : `NO`

**Success code** : `200 OK`

**Response example** :


```json
{
    "Message": "Hello, world!"
}
```


### 2. City Codes
Get the list of all available city codes.

**URL** : `/api/cities`

**Method** : `GET`

**Auth Required** : `NO`

**Success code** : `200 OK`

**Response example** :


```json
[
    {
        "Code": "GDL"
    },
    {
        "Code": "MEX"
    }
]
```

### 2. Single City 
Get the details of a specific city, such as name and current date & time.

**URL** : `/api/cities/{code}`

**Method** : `GET`

**Auth Required** : `NO`

**Success code** : `200 OK`

**Response example** :


```json
{ 
    "Code": "GDL", 
    "Name": "Guadalajara", 
    "DateTime": "2021-09-02T14:54-05:00" 
}
```

