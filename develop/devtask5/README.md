# Wildberries internship L2
## Task 5 - grep

Something like "grep" command in bash
## Usage:
```
$ go run ./ [options]... pattern [filenames]...
```

Prints lines that match pattern

### Options:

&emsp;-A&emsp;int
<br/>&emsp;&emsp;&emsp;lines to print after matching lines

&emsp;-B&emsp;int
<br/>&emsp;&emsp;&emsp;lines to print before matching lines

&emsp;-C&emsp;int
<br/>&emsp;&emsp;&emsp;lines to print after or before matching lines (if either -A or -B flag are not set, otherwise won't have any effect)

&emsp;-c&emsp;print number of matched lines in each file instead of lines themselves

&emsp;-i&emsp;ignore case (with this flag patterns that only differ in case will match)

&emsp;-v&emsp;inverse matching (match lines that DO NOT match pattern)

&emsp;-F&emsp;matches line ONLY if pattern matches whole line

&emsp;-n&emsp;also prints number of line in file

### Note:
&emsp;every flag needs to be typed SEPARATELY

&emsp;OK: -k 2 -n -u

&emsp;NOT OK: -nk 2u