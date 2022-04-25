#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
ethdir="$workspace/src/github.com/MeshBoxTech"
if [ ! -L "$ethdir/mesh-chain" ]; then
    mkdir -p "$ethdir"
    cd "$ethdir"
    ln -s ../../../../../. mesh-chain
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$ethdir/mesh-chain"
PWD="$ethdir/mesh-chain"

# Launch the arguments with the configured environment.
exec "$@"
