# Wildberries internship L2
## Task 3 - sort

Something like "sort" command in bash
## Usage:
```
$ go run ./ [options]... [filenames]...
```

Sorts lines in files and creates 
output files like "output_filename.txt"

### Options:

&nbsp;&nbsp;&nbsp;&nbsp; -k &nbsp;&nbsp;&nbsp;&nbsp; int

&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp; column number to sort by

&nbsp;&nbsp;&nbsp;&nbsp; -n &nbsp;&nbsp;&nbsp;&nbsp; sort by numeric value

&nbsp;&nbsp;&nbsp;&nbsp; -r &nbsp;&nbsp;&nbsp;&nbsp; reverse sorting order

&nbsp;&nbsp;&nbsp;&nbsp; -u &nbsp;&nbsp;&nbsp;&nbsp; do not output repeat lines

### Note:
&nbsp;&nbsp;&nbsp;&nbsp; every flag needs to be typed SEPARATELY

&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp; OK: -k 2 -n -u

&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp; NOT OK: -nk 2u