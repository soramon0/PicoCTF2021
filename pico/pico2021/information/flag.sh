#/bin/sh

strings -n 20 cat.jpg | grep license | cut -d"'" -f 2
