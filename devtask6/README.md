# Wildberries internship L2
## Task 6 - cut

Something like "grep" command in bash
## Usage:
```
$ cat input.txt | go run ./ [options]... pattern
```

Prints lines that match pattern

### Options:

&emsp;-f&emsp;int
<br/>&emsp;&emsp;&emsp;select  only these fields;  also print any line that contains no delimiter character, unless the -s option is specified

&emsp;-d&emsp;string
<br/>&emsp;&emsp;&emsp;specify field delimiter (TAB by default)

&emsp;-s&emsp;do not print lines not containing delimiters

### Note:
&emsp;every flag needs to be typed SEPARATELY

&emsp;OK: -k 2 -n -u

&emsp;NOT OK: -nk 2u