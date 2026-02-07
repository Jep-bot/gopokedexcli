#!/bin/bash

cover() {
  local t=$(mktemp -t coverXXXXXX)
  go test $COVERFLAGS -coverprofile=$t $@ &&
    go tool cover -func=$t &&
    unlink $t
}

cover
