#!/usr/bin/env bash

# This script fetches a file from a remote git repo
# Arguments are:
# 1. The URL of the git repo
# 2. The name of the branch to fetch
# 3. The name of the specification file that is used to whitelist files

# Check if the number of arguments is correct
if [ "$#" -ne 3 ]; then
    echo "Illegal number of parameters"
    exit 1
fi

# Create variables
URL=$1
BRANCH=$2
SPEC_FILE=$3


# Create dummy repo
git init

# Enable sparse Checkout
git config core.sparseCheckout true

# add whitelist to sparse-checkout
cat $SPEC_FILE >> .git/info/sparse-checkout

# add remote repo
git remote add origin $URL

# git pull from remote repo
git pull origin $BRANCH

# remove .git folder
rm -rf .git

