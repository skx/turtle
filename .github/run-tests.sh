#!/bin/bash

# Install tools to test our code-quality.
go get -u golang.org/x/lint/golint
go get -u honnef.co/go/tools/cmd/staticcheck

# Run the static-check tool
echo "Running golint on $i"
cd $i
t=$(mktemp)
staticcheck -checks all ./... > $t
if [ -s $t ]; then
    echo "Found errors via 'staticcheck'"
    cat $t
    rm $t
    exit 1
fi
rm $t



# At this point failures cause aborts
set -e

# Run the linter-tool
echo "Running golint on $i"
golint -set_exit_status ./...

# Run the vet-tool
echo "Running go vet $i"
go vet ./...
done
