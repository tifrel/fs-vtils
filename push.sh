#!/usr/bin/bash

bash test.sh

if [ $? -ne 0 ]; then
  echo "Pushing to git aborted because of failing tests."
  exit 1
fi

echo "# Commiting changes"
git add .
echo -e "\t$1\n"
git commit -m "$1"
if [ $? -ne 0 ]; then
  echo "Failed to commit changes."
  exit 1
fi

echo "# Pushing to git repo"
git push
if [ $? -ne 0 ]; then
  echo "Failed to push onto remote repo."
  exit 1
else
  echo "Commit and push successful"
  exit 0
fi
