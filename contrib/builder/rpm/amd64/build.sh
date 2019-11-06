#!/bin/bash
set -e

cd "$(dirname "$(readlink -f "$BASH_SOURCE")")"

set -x
./generate.sh centos-7
for d in */; do
        if [[ $(basename "$d") = centos-7 ]];then
            echo "Hack: only Centos7"
            docker build -t "dockercore/builder-rpm:$(basename "$d")" "$d"
        fi
done
