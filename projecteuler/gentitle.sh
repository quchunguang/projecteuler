#!/bin/bash

cat <<_END
package projecteuler

_END

for i in `seq $1 $2`; do
    title=$(curl -s https://projecteuler.net/problem=$i | grep -oPm1 "(?<=<h2>)[^<]+")
    cat <<DOCUMENTATIONXX
// Problem $i - $title
//
func PE$i() (ret int) {
    return
}

DOCUMENTATIONXX
done
