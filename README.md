dict
=====

Searchs through a copy of the webster dictionary from project gutenberg for
the word you want and prints out a definition.

Download the dictionary from
http://www.gutenberg.org/cache/epub/29765/pg29765.txt and copy it to
/usr/share/dict/dict.txt

You will need golang to build.

Build and install to /usr/bin/dict

	$ go build dict.go 
	# install dict /usr/bin/dict
	$ dict halloo
