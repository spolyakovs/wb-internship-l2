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

&emsp;-k&emsp;int
<br/>&emsp;&emsp;&emsp;column number to sort by

&emsp;-n&emsp;sort by numeric value

&emsp;-r&emsp;reverse sorting order

&emsp;-u&emsp;do not output repeat lines

### Note:
&emsp;every flag needs to be typed SEPARATELY

&emsp;OK: -k 2 -n -u

&emsp;NOT OK: -nk 2u