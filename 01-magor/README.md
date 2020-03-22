# `magor`

I've been trying to use the [Mnemonic major system](https://en.wikipedia.org/wiki/Mnemonic_major_system)
in my own life tohelp with memorizing numbers. The basic idea with this system is t

While I have the system memorized, I'm still not 100% fluent with it (sometimes
I lack the imagination to come up with good words). I thought a program could
help me with this!

For this project, I had the idea that I could write a lil' script that takes
numberical strings as input (like someone's phone number - or anything you want
to memorize) and outputs possible words according to the major system.

The usage and output looks like this:
```
->% go run *.go 123-45-6789
[/t/ /d/] [/n/] [/m/] [/r/] [/l/] [/tʃ/ /dʒ/ /ʃ/ /ʒ/] [/k/ /g/] [/f/ /v/] [/p/ /b/]
```

While I knew the full scope of that project would be a lot more work (I would
need to go from numbers to the International Phonetic Alphabet then somehow to
words that have those consonants).

It turns out that there is no good API for word pronounciations (that's
free). [Princeton's WordNet database](https://wordnet.princeton.edu/) doesn't
have pronounciations either.

Potential future work could be to make the output nicer and more helpful!
