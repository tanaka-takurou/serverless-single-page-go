# serverless-single-page kit
Simple kit for serverless single web page using AWS Lambda.


## Dependence
- aws-lambda-go


## Requirements
- AWS (Lambda, API Gateway)
- aws-cli
- golang environment
- MySQL environment


## Usage

### Management Page

#### Setting
Edit mysql-user, mysql-pass in "management/controller/db_handler.go".

#### Run
```
$ cd management
$ go run main.go
```
Then open http://localhost:8080
Using this page, You can add content, add category, deploy.
