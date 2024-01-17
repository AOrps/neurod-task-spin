#!/bin/bash
set -x

OPT_FORCE=

# usage: shows how to use this script
function usage {
cat<<EOF
    $0 -[h|f]

       -h : Help
       	  Shows the usage of the program

       -f : Force
       	  Force build of the container

EOF
exit 0
}

function main {
    _version=$(grep -E '^LABEL version.*' Dockerfile | cut -d= -f2 | tr -d '"')
    # error handling on whether or not it can parse the version number
    [ -n "$_version" ] && version=$_version || echo "err: version number could not be found!"

    if [ $OPT_FORCE ]; then
	docker build -t spins:$version . -f Dockerfile --build-arg CACHEBUST=$(date '+%FT%T.%N%:z')
	[ $? -eq 1 ] && echo "err: could not force build docker container"
    else
	docker build -t spins:$version . -f Dockerfile
	[ $? -eq 1 ] && echo "err: could not build docker container"
    fi
    
    docker tag spins:$version spins:latest
    [ $? -eq 1 ] && echo "err: unsuccessful in tagging container"

}

while getopts 'hf' opt; do
    case "$opt" in
	# h) help
       	h) usage;;
	# f) Force
	f) OPT_FORCE=1;;
    esac
done

main
