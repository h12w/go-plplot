go-plplot
=========

Go binding and extentions to PLPlot library.

There are two layers, the functions under go-plplot/c are one-to-one mapping of
C function calls. The types under
go-plplot form a declarative interface but are not complete yet.

Note
----
C function wrappers are automatically generated by cwrap, so there is no need to
check them into the repo. Only additional Go extentions are included here. The
generator is under the "examples" directory of cwrap: https://h12.io/cwrap
The default output path is $GOPATH/src, so it is suggested that a symlink is
created under $GOPATH/src to link it here.

REFERENCE
---------
http://plplot.sourceforge.net
