plplot-doc-check
================

This is a simple program for checking the consistancy between the header file and the
xml document of PLPlot.

Here are minimal steps to run the program:

* Get and install the current stable version of Go language (just install it and Set GOPATH environment variable)

   Please refer to: http://golang.org/doc/install

* Use "go get" tool to download this package

        go get -u h12.io/go-plplot/c/plplot-doc-check

* Build

        cd $GOPATH/h12.io/go-plplot/c/plplot-doc-check
        go build

* Run it without argument to see the command line arguments available

        ./plplot-doc-check

