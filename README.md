multigrep
==========

Powerful text search/grep program.

Matcher/Expression to control the way of serach
----

You can search/grep by using following matcher/expression.

- Plain String Matcher (default)

    Matches if the target string contains string of this matcher. 
    (Has poor functionality but it must be fastest.)

    This matcher is used if you omit the prefix.
    You can also explicitly specify this matcher by using prefix `s:`.

- Regular Expression Matcher

     You can use regular expression matcher by adding prefix `r:`
     and `i:` for case insensitive regular expression.

- Migemo Matcher

    Search Japanese text using latin-1 text by adding prefix `m:`.
    (powered by github.com/koron/gomigemo)

- Not Expression

    Prefix `!` inverts the match result.
    It should be placed before any of other matcher prefix.

Installation (WIP)
------------

Please visit our releases page, and download the appropriate version/platform:

[https://github.com/peco/migemogrep/releases](https://github.com/peco/migemogrep/releases)

Or, if you are on OS X and are using homebrew

    $ brew tap peco/peco
    $ brew install migemogrep

And finally, if you want the latest bleeding edge version:

    $ go get github.com/peco/migemogrep

Usage
-----

```sh
$ multigrep <pattern> <file>
$ cat file.txt | multigrep <pattern>
```

Example
----

1. To search string "some" from all golang source code in current directory.

        multigrep some *.go

2. To search string "Some" or "some".

        multigrep r:[S|s]ome *.go

    or

        multigrep i:some *.go

3. To search Japanese string "何".

        multigrep m:nani *.go

4. To search string which does not contain "some" or "Some".

        multigrep !i:some *.go

5. To search string which does not contain "some" or "Some" but "何"

        multigrep "!i:some m:nani" *.go

