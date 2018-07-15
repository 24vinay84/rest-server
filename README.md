# JIO Rest Server
User can **register (sign-up) and sign-in** to access this **jio store server**.
A basic online store API written to learn Go Programming Language

This API is a very basic implementation of an online(e-commerce) store.
- **Authentication** 
	- New user can sign-up using username and password.
	- Existing user can sign-in using username and password.
- **Store REST APIs**	
	- You can perform basic CRUD(CREATE, READ, UPDATE and DELETE) operations
	- SEARCH on a predefined database of products 
	- Only Authenticated users can Get, Search, Add, Update and Delete products from database
- **JWT**
	- Authentication is based on JWT(JSON web Tokens) Tokens
- **Database**	
	- API is backed by a predefined **MongoDB** database.
- **Distributed Rate Limit**
	- Distributed rate limit GET and POST APIs per **IP Address**.

See [API Documentation and Usage](#api-documentation-and-usage) below on how to use it.

## Directory Structure
```
rest-server/
    |- store/              - Contains main API logic files 
        |- controller.go  - Defines methods handling calls at various endpoints
        |- model.go       - User and Product models
        |- repository.go  - Methods interacting with the database
        |- router.go      - Defines routes and endpoints
		|- encryption.go  - Method for Password encryption/decryption before saving into Database
        |- rate-limit.go  - Distributed rate limit GET and POST APIs per IP Address
    |- vendor/             - Dependency packages, necessary for deployment
    |- README.md
    |- mockingData.js       - Script to populate local mongodb with dummy data
    |- main.go              - Entry point of the API
	|- rest-server.exe      - Executable for this server
	
```

## Setup

### Project setup

1. Clone the repository in your `$GOPATH/src/` directory. If you have used the bash script for setup, your `$GOPATH` variable should point to `$HOME/go`
2. To run project locally, Install Mongo DB - https://www.mongodb.com/download-center?jmp=nav#community
3. After installing Mongo DB, start it's server by typing `mongod` in Terminal.
4. Open a new tab in terminal and type `mongo < mockingData.js` to insert the dummmy product data.
5. Open file `store/repository.go`, find the `SERVER` variable and replace with your URL for remote machine. 

```
const SERVER = "http://localhost:27017"
```
6. Last thing required to run the project, install all the go dependencies.
```
// Library to handle jwt authentication 
$ go get "github.com/dgrijalva/jwt-go"

// Libraries to handle network routing
$ go get "github.com/gorilla/mux"
$ go get "github.com/gorilla/context"
$ go get "github.com/gorilla/handlers"

// mgo library for handling Mongo DB
$ go get "gopkg.in/mgo.v2"
```
Wow! Now we're ready to run the API :tada: <br>
8. Type `export PORT=8000` in Terminal and open http://localhost:8000 in your browser to see the products.

## API Documentation and Usage

It is **recommended** to install some extension to beautify JSON(like [JSON Formatter](https://chrome.google.com/webstore/detail/json-formatter/bcjindcccaagfpapjjmafapmmgkkhgoa)) if you're trying in a browser.

**Important** - Don't forget to define $PORT in your shell variables. <br>Example: `export PORT=8000`

```sh
BASE_URL = "http://localhost:$PORT"
```

For **Retreive**,**Adding**, **Updating** , **Search** and **Deleting** products from database you must send a JWT token in Authentication header.
### 1. Authentication ( Sign-in)

- **Endpoint Name** - `SignIn` <br>
- **Method** - `GET`           <br>
- **URL Pattern** - `/sign-in` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X GET \
    -H "Content-Type: application/json" \
    -d '{ username: "<YOUR_USERNAME>", password: "<RANDOM_PASSWORD>"}' \
    BASE_URL/sign-in
    ```
- **Expected Response** - A JWT Authentication Token

### 2. Authentication ( Sign-up)

- **Endpoint Name** - `SignIn` <br>
- **Method** - `POST`            <br>
- **URL Pattern** - `/sign-up` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ username: "<YOUR_USERNAME>", password: "<RANDOM_PASSWORD>"}' \
    BASE_URL/sign-up
    ```
- **Expected Response** - Sign-up successful without any error message. Check the logs in Terminal window which is running server.

### 3. View Single Product

- **Endpoint Name** - `GetProduct`    <br>
- **Method** - `GET`                  <br>
- **URL Pattern** - `/products/{id}`  <br>
- **Usage**
    - Open BASE_URL/products/{id} in browser
    - **Terminal/CURL**
```
curl -X GET BASE_URL/products/{id} 
```
- **Expected Response** - Product with the {id} in database
- **NOTE** - There are only six(6) ids in the database, so 1 <= {id} <= 6   

### 4. Search Product

- **Endpoint Name** - `SearchProduct`  <br>
- **Method** - `GET`                   <br>
- **URL Pattern** - `/search/{query}`  <br>
- **Usage** - Browser OR curl        
- **BROWSER**
    - Open BASE_URL/search/{query} in browser
    - **Terminal/CURL**
    ```sh
    curl -X GET BASE_URL/search/{query}
    ```
- **Expected Response** - Products matching the search query <br>


### 5. Add Product

- **Endpoint Name** - `AddProduct` <br>
- **Method** - `POST`              <br>
- **URL Pattern** - `/add-product`  <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X POST \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    -d '{ "_id": 11, 
        "title": "Memes",
        "image": "I am selling memes, hehe.",          
        "price": 1,
        "rating": 5
        }' \
    BASE_URL/add-product
    ```
- **Expected Response** - Addition successful without any error message. Check the logs in Terminal window which is running server. 

### 6. Update Product

- **Endpoint Name** - `UpdateProduct` <br>
- **Method** - `PUT`                  <br>
- **URL Pattern** - `/update-product`  <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X PUT \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    -d '{ "ID": 14, 
        "title": "Memes",
        "image": "I am not selling memes to you, hehe.",          
        "price": 1000,
        "rating": 5
        }' \
    BASE_URL/update-product
    ```
- **Expected Response** - Update successful without any error message. Check the logs in Terminal window which is running server. <br>

### 7. Delete Product

- **Endpoint Name** - `DeleteProduct` <br>
- **Method** - `DELETE` <br>
- **URL Pattern** - `/delete-product/{id}` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X DELETE \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    BASE_URL/delete-product/{id}
    ```
- **Expected Response** - Deletion successful without any error message. Check the logs in Terminal window which is running server. <br>


## TODO
* [ ] Write unit tests to test every method
* [ ] Improve the code by proper exception handling
* [ ] User and roles management
* [ ] Session management using JWT tokens
* [ ] **Distributed rate limit GET and POST APIs per subscriber ID**.



