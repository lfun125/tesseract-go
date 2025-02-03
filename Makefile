gen: 
	cd proto && buf dep update && buf generate
docker:
	docker build . --platform linux/amd64 -t tesseract-go