#!/bin/sh

version=`printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"`
sed -i "s/<%VERSION%>/${version}/g" main.go
