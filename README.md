# fakepack

I am not yet game to rewrite all my JCR6 utilites in Go. This may happen eventually, though, but the time is not now.
However, I do need to be able to use some compression methods working in JCR6 that that BlitzMax does not support natively. There are 3rd party modules out there, but I've bad experiences with them, tbh.

This fakepack module will have to suffice in the mean time. It will use Go to perform the actual packing, and BlitzMax just has to pick up the results. Since JCR6 is a MODULAR format, this workabout will work, as long as all tools have the fakepack module in place.

Please note if you are going to use the cli tools while this setup is still the standard, you will need to have Go installed.
The binaries of the cli tools will be able to put the GO file into a temp dir and Go can run Go programs without having to convert them into executables.

This is just a temp setup in my current plans.

Please note all .go files are programs on their own, so the directory they are in does not count as the package.
I know this setup is rather unusual for Go, but for this quick set up it was the easiest and most efficient way to go.
