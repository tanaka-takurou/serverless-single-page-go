# serverless-single-page kit
Simple kit for serverless single web page using AWS Lambda.


## Dependence
- aws-lambda-go


## Requirements
- AWS (Lambda, API Gateway)
- aws-cli
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
- Edit templates/index.html like as 'sample_jpg'.
- Edit create_template.sh like as 'sample_jpg'. (add three lines)

### Deploy
Open scripts/deploy.sh and edit 'your_function_name' first.
Then run this command.

```
$ sh scripts/deploy.sh
```
