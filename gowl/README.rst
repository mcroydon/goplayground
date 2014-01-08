Gowl
====

Sometimes you really need some owls served over TCP.

Serving owls
============

Here you go::

	$ go run gowl.go
	2014/01/08 13:02:27 Serving owls on :3400.

Consuming owls
==============

Here you go::

	$ telnet localhost 3400
	Trying ::1...
	Connected to localhost.
	Escape character is '^]'.
	^(OvO)^
	Connection closed by foreign host.

	$ telnet localhost 3400
	Trying ::1...
	Connected to localhost.
	Escape character is '^]'.
	  ___
	 (o,o)
	<  .  >
	--"-"---
	Connection closed by foreign host.
