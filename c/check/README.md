check
=====

It is a simple program to check the consistancy between header file and xml
document of PLPlot.

Here are minimal steps to run the program:

* Get and install the current stable version of Go language
    Please refer: http://golang.org/doc/install
* Set GOPATH environment variable
* Use "go get" tool to download this package:
    go get -u github.com/hailiang/go-plplot
* Build
    cd $GOPATH/github.com/hailiang/go-plplot/c/check
    go build
* Run it without argument to see the command line option available
    ./check

