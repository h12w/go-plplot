check
=====

This is a simple program for checking the consistancy between the header file and the
xml document of PLPlot.

Here are minimal steps to run the program:

* Get and install the current stable version of Go language (just install it and Set GOPATH environment variable)

   Please refer to: http://golang.org/doc/install

* Use "go get" tool to download this package

        go get -u h12.me/go-plplot

* Build

        cd $GOPATH/h12.me/go-plplot/c/check
        go build

* Run it without argument to see the command line arguments available

        ./check

