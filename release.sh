#!/bin/bash

# Release script for AutoStack
# Usage: ./release.sh <version>
# Example: ./release.sh 1.0.0

set -e

if [ -z "$1" ]; then
    echo "Usage: ./release.sh <version>"
    echo "Example: ./release.sh 1.0.0"
    exit 1
fi

VERSION="$1"
TAG="v${VERSION}"

echo "Creating release ${TAG}..."

# Check if there are uncommitted changes
if [[ -n $(git status -s) ]]; then
    echo "Error: You have uncommitted changes. Please commit or stash them first."
    git status -s
    exit 1
fi

# Check if tag already exists
if git rev-parse "$TAG" >/dev/null 2>&1; then
    echo "Error: Tag ${TAG} already exists."
    echo "To delete it: git tag -d ${TAG} && git push origin :refs/tags/${TAG}"
    exit 1
fi

# Confirm with user
echo ""
echo "This will:"
echo "  1. Create tag ${TAG}"
echo "  2. Push to origin"
echo "  3. Trigger GitHub Actions to build and release"
echo ""
read -p "Continue? (y/n) " -n 1 -r
echo

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Release cancelled."
    exit 1
fi

# Create and push tag
git tag -a "$TAG" -m "Release ${TAG}"
git push origin "$TAG"

echo ""
echo "Release ${TAG} created successfully!"
echo ""
echo "GitHub Actions is now building the binaries."
echo "Check progress at: https://github.com/bait-py/autostack/actions"
echo "Release will be available at: https://github.com/bait-py/autostack/releases/tag/${TAG}"
