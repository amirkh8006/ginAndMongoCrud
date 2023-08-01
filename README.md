## Go Crud App With Gin & MongoDB In MVC Structure

**Run**
 * **$ go run ginMongo**

 Create User

* **URL**

  http://127.0.0.1:9090/v1/user/create

* **Method:**

  `POST`
* **Success Response:**
  

  * {"message": "success"}



Get All Users

* **URL**

  http://127.0.0.1:9090/v1/user/getall

* **Method:**

  `GET`
* **Success Response:**
  

  * [Array Of Users List]


Get User

* **URL**

  http://127.0.0.1:9090/v1/user/get/{username}

* **Method:**

  `GET`
* **Success Response:**
  

  * {User Details}


Update User

* **URL**

  http://127.0.0.1:9090/v1/user/update

* **Method:**

  `PATCH`
* **Success Response:**
  

  * {"message": "success"}

Delete User

* **URL**

  http://127.0.0.1:9090/v1/user/delete/{username}

* **Method:**

  `DELETE`
* **Success Response:**
  

  * {"message": "success"}
    
