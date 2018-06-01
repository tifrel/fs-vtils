#!/usr/bin/bash

tests_mkfile=$src/testdata/MkFile

# TODO: less duplication

reset_MkFile() {

  cd $tests_mkfile

  for i in {1..11}; do
    if [[ $i -lt 10 ]]; then
      c=case0$i
    else
      c=case$i
    fi

    rm -R $c
    mkdir $c

    mkdir $c/parent
    touch $c/parentfile

    mkdir $c/parent/dir
    touch $c/parent/file


  done

  cd $src
}

evaluate_MkFile() {
  cd $tests_mkfile

  # case01
  if [[ ! -f case01/parentfile ]]; then
    fail "MkFile" 01 1
  elif [[ ! -d case01/parent ]]; then
    fail "MkFile" 01 2
  elif [[ ! -d case01/parent/dir ]]; then
    fail "MkFile" 01 3
  elif [[ ! -f case01/parent/file ]]; then
    fail "MkFile" 01 4
  elif [[ ! -f case01/parent/new ]]; then
    fail "MkFile" 01 5
  fi


  # case02
  unchanged_MkFile case02

  # case03
  if [[ ! -f case03/parentfile ]]; then
    fail "MkFile" 03 1
  elif [[ ! -d case03/parent ]]; then
    fail "MkFile" 03 2
  elif [[ ! -d case03/parent/dir ]]; then
    fail "MkFile" 03 3
  elif [[ ! -f case03/parent/file ]]; then
    fail "MkFile" 03 4
  fi

  # case04
  unchanged_MkFile case04


  # case05
  if [[ ! -f case05/parentfile ]]; then
    fail "MkFile" 05 1
  elif [[ ! -d case05/parent ]]; then
    fail "MkFile" 05 2
  elif [[ ! -f case05/parent/dir ]]; then
    fail "MkFile" 05 3
  elif [[ ! -f case05/parent/file ]]; then
    fail "MkFile" 05 4
  fi

  # case06
  unchanged_MkFile case06

  # case07
  if [[ ! -f case07/parentfile ]]; then
    fail "MkFile" 07 1
  elif [[ ! -d case07/parent ]]; then
    fail "MkFile" 07 2
  elif [[ ! -d case07/parent/dir ]]; then
    fail "MkFile" 07 3
  elif [[ ! -f case07/parent/file ]]; then
    fail "MkFile" 07 4
  elif [[ ! -d case07/parent2 ]]; then
    fail "MkFile" 07 5
  elif [[ ! -f case07/parent2/new ]]; then
    fail "MkFile" 07 6
  fi

  #cases08/09/10
  unchanged_MkFile case08
  unchanged_MkFile case09
  unchanged_MkFile case10

  # case11
  if [[ ! -d case11/parentfile ]]; then
    fail "MkFile" 11 1
  elif [[ ! -f case11/parentfile/new ]]; then
    fail "MkFile" 11 2
  elif [[ ! -d case11/parent ]]; then
    fail "MkFile" 11 2
  elif [[ ! -d case11/parent/dir ]]; then
    fail "MkFile" 11 3
  elif [[ ! -f case11/parent/file ]]; then
    fail "MkFile" 11 4
  fi

  cd $src
}

unchanged_MkFile() {
  local case="${1:4:2}"

  if [[ ! -f $1/parentfile ]]; then
    fail "MkFile" $case 1
  elif [[ ! -d $1/parent ]]; then
    fail "MkFile" $case 2
  elif [[ ! -d $1/parent/dir ]]; then
    fail "MkFile" $case 3
  elif [[ ! -f $1/parent/file ]]; then
    fail "MkFile" $case 4
  fi

}
