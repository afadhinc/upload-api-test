build:  
	docker build -t upload-api  .
run: 
	docker run --rm -p 8080:8080 upload-api