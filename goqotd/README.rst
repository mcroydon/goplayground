GoQOTD
======

An implementation of the `Quote of the Day Protocol (RFC 865) <http://tools.ietf.org/html/rfc865>`_.  Surprisingly
enough I don't think I've implemented this one before, though gowl is probably RFC 865 complaint as long as no owls
exceed 512 bytes.

Installation
============

Install via go install::

    $ go install github.com/mcroydon/goplayground/goqotd

Running
=======

.. DANGER::
   The QOTD protocol runs on port 17 which requires superuser access.  Running this program in anything other than
   a controlled environment behind a firewall should scare you.

You will need to run goqotd as a superuser using ``su``, ``sudo`` or similar::

    $ goqotd
    2014/02/06 10:39:00 goqotd listening for connections on port 17.

If you run goqotd as a regular user, you will likely receive an error like this::

    $ goqotd
    2014/02/06 10:38:51 listen tcp :17: bind: permission denied
    exit status 1

Using
=====

You can interact with goqotd using telnet::

    $ telnet localhost 17
    Trying ::1...
    Connected to localhost.
    Escape character is '^]'.
    "In the immortal words of Jean Paul Sartre, 'Au revoir, gopher'."
    Connection closed by foreign host.

Extending
=========

There are several things that can be done to make goqotd more useful.  In no particular order:

* Support more than one quote.
* Support a random quote per request (see gowl).
* Load quotes from a file.
* Use `flag <http://golang.org/pkg/flag/>`_ to configure support for other ports, a quote to use, or a file
  to load quotes from.
