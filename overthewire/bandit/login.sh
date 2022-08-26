#!/bin/bash

# ls -v — natural sort of (version) numbers within text
files=`ls -v | grep -i "^bandit*"`
last_file=`echo "$files" | tail -1`

echo "ssh to $last_file@bandit.labs.overthewire.org"
sshpass -p `cat "$last_file"` ssh $last_file@bandit.labs.overthewire.org -p 2220
 
