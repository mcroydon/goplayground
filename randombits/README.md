# randombits

Use [termui](https://github.com/gizak/termui) and [qrand](https://github.com/salviati/go-qrand/tree/master/qrand) to retrieve random bits from the [Quantum Random Bit Generator](http://random.irb.hr) and display them in your terminal.

This is really just an excuse to play with termui, which looks incredibly neat, and way less painful than the last time I wrote something to the terminal.

To run:

	$ go get github.com/mcroydon/goplayground/randombits
	$ $GOPATH/bin/randombits -u <yourusername> -p <yourpassword>

Once the pogram is running, you can hit q to quit or r for another random int64.
