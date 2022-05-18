# WEBPAGE-DRILL

This is a simple api server to drill websites.


## Setting up in a local machine
1. Install latest Go version >1.18
2. Clone the repo and goto root directory of the project.
3. Run the web server `go run main.go` 
    - Your console should look like ![Employee data](/screenshots/server-start.png)
4. Using client is completely up to you, you can use postman like tools or simply curl the api
    - curl command -> `curl --request GET 'http://127.0.0.1:8848/?url=www.facebook.com/'`
