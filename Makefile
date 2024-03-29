root	:=		$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: clean build deploy

clean:
	rm -rfv bin

build:
	mkdir -p bin
	scripts/create_template.sh
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/bootstrap

deploy:
	sam package --output-template-file "${root}"/packaged.yml --s3-bucket "${bucket}"
	sam deploy --stack-name "${stack}" --capabilities CAPABILITY_IAM --template-file "${root}/packaged.yml"
