#!/usr/bin/bash

back=$(pwd)
src=/Users/Till/Code/go/src/fsv

function fail {
  echo "FAIL"
  echo -e "\t$1 failed at $2, condition $3"
  exit 1
}

source $src/testdata/mkdir_tests.sh
source $src/testdata/mkfile_tests.sh
source $src/testdata/rm_tests.sh
source $src/testdata/mv_tests.sh
source $src/testdata/ln_tests.sh
source $src/testdata/cp_tests.sh


echo -e "\n# Resetting testing directories"
reset_MkDir
reset_MkFile
reset_Rm
reset_Mv
reset_Ln
reset_Cp

echo -e "\n# Running go tests"
go test

echo -e "\n# Checking for correct fs mutations"
evaluate_MkDir
evaluate_MkFile
evaluate_Rm
evaluate_Mv
evaluate_Ln
evaluate_Cp

echo "PASS"

echo ""
cd $back
exit 0
