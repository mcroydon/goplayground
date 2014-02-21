========
Timetoys
========

Playing with time in go.

Installing the `timetoys package <https://godoc.org/github.com/mcroydon/goplayground/timetoys>`_::

    $ go get github.com/mcroydon/goplayground/timetoys

There is currently a single function, ``HowFast()`` which continuously calls time.Now(), records the delta in
nanoseconds from the previous call, prints to stdout when a new maximum delta has been reached and prints
statistics every second.

As a convenience, the ``howfast`` command is a simple wrapper around ``HowFast()`` if you want to see how fast go can,
err, "go" on your system::

    $ go install github.com/mcroydon/goplayground/timetoys/howfast
    $ howfast
    Feb 20 22:30:07.759707395 New max delta 19139 since Feb 20 22:30:07.759688256
    Feb 20 22:30:07.759862271 New max delta 154876 since Feb 20 22:30:07.759707395
    Feb 20 22:30:08.177426976 New max delta 237491 since Feb 20 22:30:08.177189485
    Feb 20 22:30:08.763725265 processed 26385297 current delta 26 max delta 237491
    Feb 20 22:30:09.762292411 processed 52407912 current delta 38 max delta 237491
    Feb 20 22:30:10.765726246 processed 78265792 current delta 38 max delta 237491
    Feb 20 22:30:11.781065042 processed 104502132 current delta 33 max delta 237491
    Feb 20 22:30:12.176512596 New max delta 395212 since Feb 20 22:30:12.176117384
    Feb 20 22:30:12.775715636 processed 130058403 current delta 37 max delta 395212
    Feb 20 22:30:13.771520873 processed 155695735 current delta 39 max delta 395212
    Feb 20 22:30:14.772070841 processed 181513732 current delta 38 max delta 395212

This will run forever until you quit the program with Control-C.
