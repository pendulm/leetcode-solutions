#!/bin/bash

problem_number=$1
problem_dir=$(leetcode show $1|head -1|sed 's/[][]//g;s/^[ \t]*//;s/[ \t]*$//'|tr ' ' '-')
src_file=${problem_dir#*-}
src_file=$(echo ${src_file//-/_} | tr '[A-Z]' '[a-z]')
src_file="$src_file.go"

fill_golang_template() {
    cat <<-EOF > $src_file
package main

import (
    "fmt"
)


func main() {
}
EOF
}

main() {
    mkdir -p $problem_dir
    cd $problem_dir
    if [[ ! -f $problem_dir/$src_file ]]; then
        fill_golang_template
        echo "create template $src_file"
    fi
}

main


