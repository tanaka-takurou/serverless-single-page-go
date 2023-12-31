# serverless-single-page kit
Simple kit for serverless single web page using AWS Lambda.


## Dependence
- aws-lambda-go


## Requirements
- AWS (Lambda, API Gateway)
- aws-sam-cli
- golang environment


## Usage

### Edit View
##### HTML
- Edit templates/index.html

##### CSS
- Edit static/css/main.css

##### Javascript
- Edit static/js/main.js

##### Image
- Add image file into static/img/
- Edit templates/index.html like as 'sample.jpg'.

### Deploy
```bash
make clean build
AWS_PROFILE={profile} AWS_DEFAULT_REGION={region} make bucket={bucket} stack={stack name} deploy
```

### Example
[README](https://github.com/tanaka-takurou/serverless-single-page-go/tree/master/example)
