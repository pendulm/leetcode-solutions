#!/bin/bash

SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"

die() {
    echo "$@" >&2
    exit 1
}

usage() {
    die "scriptname.sh <problem ID> [Lang]"
}

if [[ -z $1 ]]; then
    usage
fi

script_lang=${2:-golang}
script_name=""

case $script_lang in
    go|golang) script_name=golang_template.sh;;
    rust) script_name=rust_template.sh;;
    python) die no python template;;
    java|c|c++) die no java, c, c++ template;;
    *) die no $script_name template ;;
esac


bash  "$SCRIPTPATH/$script_name" $1
