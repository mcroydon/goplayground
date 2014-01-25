Similarity
==========

The beginnings of a simple similarity engine, similar to the recommendation engine created
in Chapter 2 of `Programming Collective Intelligence <http://shop.oreilly.com/product/9780596529321.do>`_.

Installing
----------

::

    $ go get github.com/mcroydon/goplayground/similarity

Testing
-------

::

   $ go test github.com/mcroydon/goplayground/similarity

Using
-----

Adding items::

    sim := NewSimilarity()
    sim.Add("mcroydon", Item{"http://golang.org", 5.0}
    fmt.Print(sim.Keys())
    [mcroydon]
