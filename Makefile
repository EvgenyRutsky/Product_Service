check_install:
	GO111MODULE=off go get -u github.com/go-swagger/

swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models check