.PHONY: apisprout-lambda zip goupx

apisprout-lambda:
	go mod vendor
	GOOS=linux go build -ldflags "-w" -o out/apisprout-lambda

goupx:
	goupx out/apisprout-lambda

zip: apisprout-lambda goupx
	zip -r -j out/apisprout-lambda.zip out/apisprout-lambda
