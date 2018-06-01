#!/usr/bin/bash

tests_rm=$src/testdata/Rm

function reset_Rm {
  cd $tests_rm

  for i in {1..5}; do

    if [[ $i -lt 10 ]]; then
      c=case0$i
    else
      c=case$i
    fi

    rm -R $c
    mkdir $c

    touch $c/file
    mkdir $c/dir
    touch $c/dir/file

    ln -s $c/file $c/symlink

  done

  cd $src
}

function evaluate_Rm {
  cd $tests_rm

  #case01
  c=case01
  if [[ -a $c/file ]]; then
    fail "Rm" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Rm" $c 2
  elif [[ ! -L $c/symlink ]]; then
    fail "Rm" $c 3
  elif [[ ! -f $c/dir/file ]]; then
    fail "Rm" $c 4
  fi

  #case02
  c=case02
  if [[ ! -f $c/file ]]; then
    fail "Rm" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Rm" $c 2
  elif [[ -a $c/symlink ]]; then
    fail "Rm" $c 3
  elif [[ ! -f $c/dir/file ]]; then
    fail "Rm" $c 4
  fi

  #case03
  unchanged_Rm case03

  #case04
  unchanged_Rm case04

  #case05
  c=case05
  if [[ ! -f $c/file ]]; then
    fail "Rm" $c 1
  elif [[ -a $c/dir ]]; then
    fail "Rm" $c 2
  elif [[ ! -L $c/symlink ]]; then
    fail "Rm" $c 3
  elif [[ -a $c/dir/file ]]; then
    fail "Rm" $c 4
  fi

  cd $src
}

function unchanged_Rm {
  local c=$1

  if [[ ! -f $c/file ]]; then
    fail "Rm" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Rm" $c 2
  elif [[ ! -L $c/symlink ]]; then
    fail "Rm" $c 3
  elif [[ ! -f $c/dir/file ]]; then
    fail "Rm" $c 4
  fi
}
