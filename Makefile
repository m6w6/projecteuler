.EXPORT:
.SUFFIXES:

all: 1 2 3 4 5 6 7 8 9 10 11 12

1:
	php $@.php

2 3 4 5 6 7 8 9 10 11 12:
	GOPATH=$$(pwd) go run src/$@.go
