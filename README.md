multigrep
==========

A powerful text search/grep program.

One of the intentions of this program is to use as a matcher of [peco](https://github.com/peco/peco).

Usage
-----

```sh
$ multigrep <pattern> <file>
$ cat file.txt | multigrep <pattern>
```

This program intoroduces kind of unique __prefix search control system__ to pattern expression.

Taste
----

1. To search string which contains "some" or "Some" from all text files in current directory.

        multigrep some *.txt

2. To search string which contains "some".

        multigrep r:some *.txt

    or

        multigrep s:some *.txt

3. To search string which contains Japanese "何".

        multigrep m:nani *.txt

4. To search string which does not contain "some" or "Some".

        multigrep !some *.txt

5. To search string which does not contain "some" but "何"

        multigrep "!s:some m:nani" *.txt

6. To search string which contains "some stuff"

        multigrep some\sstuff *.txt

Details
----

You can controll search/grep by using following stuffs.

- Plain String Matcher

    You can use string matcher by adding prefix `s:`.  
    Matches if the target string contains string of this matcher.   
    (Has poor functionality but it must be fastest.)

- Case Sensitive Regular Expression Matcher

    You can use case sensitive regular expression matcher by adding prefix `r:`.

- Case Insensitive Regular Expression Matcher __(default matcher)__

    You can use case insensitive regular expression matcher by adding prefix `i:`.  
    If you omit the prefix, this matcher is used.

- Migemo Matcher

    Search Japanese text using latin-1 text by adding prefix `m:`.

- Not

    Prefix `!` inverts the match result.  
    It should be placed before any of other matcher prefix.

- Delimiter

    ` `, space, always works as a delimiter of patterns.  
    Each sub-patterns are connected as AND and evaluated.

    Use \s If you want to search phrase which includes ` `.  
    (You need to use Regular Expression Matcher)

Installation
------------

Currently there is no binary version.

You need to install golang in advance, to build and install.
Then you can get with following command.

    $ go get github.com/msr1k/multigrep

Thanks
----

- Almost all of this program is copied from [migemogrep](https://github.com/peco/migemogrep).

- Migemo Matcher is implemented in [gomigemo](https://github.com/koron/gomigemo).

