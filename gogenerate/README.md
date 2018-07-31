
    cd ${GOPATH}/src/github.com/mozey/go-cookbook/gogenerate
    
    rm -f poem.txt & go generate
    
    go run main.go