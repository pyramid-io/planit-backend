#!/bin/sh

current_dir=$(dirname "$0")
cd "$current_dir"

git config core.hooksPath "$(pwd)/git/hooks/"

