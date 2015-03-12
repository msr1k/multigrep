multigrep
==========

A powerful text search/grep program.

Note
----

Almost all of this program is copied from [migemogrep](https://github.com/peco/migemogrep) great program.

This program works fine as a matcher of [peco](https://github.com/peco/peco)

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

    To this feature use [gomigemo](https://github.com/koron/gomigemo).

- Not Expression

    Prefix `!` inverts the match result.
    It should be placed before any of other matcher prefix.

- Delimiter Expression

    Space ` ` always works as a delimiter of patterns.
    Delimitted patterns are evaluated as AND condition.

    Use \s with Regular Expression Matcher,
    If you want to search phrase which includes space ` `.

Installation
------------

Currently there is no binary version.

You need to install golang in advance, to build and install.
Then you can get with following command.

    $ go get github.com/msr1k/multigrep

Usage
-----

```sh
$ multigrep <pattern> <file>
$ cat file.txt | multigrep <pattern>
```

Example
----

1. To search string "some" from all text files in current directory.

        multigrep some *.txt

2. To search string "Some" or "some".

        multigrep r:[S|s]ome *.txt

    or

        multigrep i:some *.txt

3. To search Japanese string "何".

        multigrep m:nani *.txt

4. To search string which does not contain "some" or "Some".

        multigrep !i:some *.txt

5. To search string which does not contain "some" or "Some" but "何"

        multigrep "!i:some m:nani" *.txt

6. To search string which contains "some food"

        multigrep r:some\sfood *.txt

