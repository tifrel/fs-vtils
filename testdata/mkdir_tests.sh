#!/usr/bin/bash

tests_mkdir=$src/testdata/MkDir


reset_MkDir() {

  cd $tests_mkdir

  for i in {1..9}; do
    if [[ $i -lt 10 ]]; then
      c=case0$i
    else
      c=case$i
    fi

    rm -R $c
    mkdir $c

    mkdir $c/parent
    mkdir $c/parent/dir

    touch $c/parentfile
    touch $c/parent/file

  done

  cd $src
}


evaluate_MkDir() {
  cd $tests_mkdir

  # case01
  unchanged_MkDir case01

  # case02
  if [[ ! -f case02/parentfile ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case02/parent ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case02/parent/dir ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case02/parent/new ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -f case02/parent/file ]]; then
    echo "FAIL" && exit 1
  fi

  # case03
  unchanged_MkDir case03

  # case04
  if [[ ! -f case04/parentfile ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case04/parent ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case04/parent/dir ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case04/parent/file ]]; then
    echo "FAIL" && exit 1
  fi

  # case05/06/07
  unchanged_MkDir case05
  unchanged_MkDir case06
  unchanged_MkDir case07

  # case08
  if [[ ! -d case08/parentfile ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case08/parentfile/new ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case08/parent ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case08/parent/dir ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -f case08/parent/file ]]; then
    echo "FAIL" && exit 1
  fi

  # case09
  if [[ ! -d case09/parentfile ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case09/parentfile/deeply/nested/new/dir ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case09/parent ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d case09/parent/dir ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -f case09/parent/file ]]; then
    echo "FAIL" && exit 1
  fi

  cd $src
}

unchanged_MkDir() {


  if [[ ! -f $1/parentfile ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d $1/parent ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -d $1/parent/dir ]]; then
    echo "FAIL" && exit 1
  elif [[ ! -f $1/parent/file ]]; then
    echo "FAIL" && exit 1
  fi

}
