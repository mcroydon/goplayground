# numberstation

I've been intrigued by [numbers stations](https://en.wikipedia.org/wiki/Numbers_station) for as long as I've
known about them. I thought it would be fun to put together a small server that broadcast random numbers (in
this case a random Uint64) over UDP Multicast.

## Installation

You can run numberstation from source:
	cd $GOPATH/src
	git clone git@github.com:mcroydon/goplayground.git
	cd github.com/mcroydon/goplayground/numberstation
	go run numberstation.go

## Usage

Numberstation takes three optional flags but provides sane defaults:

	-ip string
		IP to send UDP traffic to. (default "224.0.0.1")
	-port uint
		Port number to use. (default 4815)
	-seconds uint
		Number of seconds betweeen broadcast. (default 1)
