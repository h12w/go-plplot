check
=====

It is a simple program to check the consistancy between the header file and the
xml document of PLPlot.

Here are minimal steps to run the program:

* Get and install the current stable version of Go language

   Please refer to: http://golang.org/doc/install

* Set GOPATH environment variable

* Use "go get" tool to download this package:

    go get -u github.com/hailiang/go-plplot

* Build

    cd $GOPATH/github.com/hailiang/go-plplot/c/check
    go build

* Run it without argument to see the command line option available

    ./check

