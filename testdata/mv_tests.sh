#!/usr/bin/bash

tests_mv=$src/testdata/Mv

function reset_Mv {
  cd $tests_mv

  for i in {1..9}; do

    if [[ $i -lt 10 ]]; then
      c=case0$i
    else
      c=case$i
    fi

    rm -R $c
    mkdir $c

    mkdir $c/src
    mkdir $c/target

    for dir in $c/src $c/target; do
      echo $(basename $dir) > $dir/file
      mkdir $dir/dir
      touch $dir/dir/file
      ln -s $tests_mv/$dir/file $dir/symlink
    done

  done

  cd $src
}

function evaluate_Mv {
  cd $tests_mv

  c=case01
  if [[ -a $c/src/file ]]; then
    fail "Mv" $c 1
  elif [[ ! -d $c/src/dir ]]; then
    fail "Mv" $c 2
  elif [[ ! -L $c/src/symlink ]]; then
    fail "Mv" $c 3
  elif [[ ! -f $c/src/dir/file ]]; then
    fail "Mv" $c 4
  elif [[ ! -f $c/target/file ]]; then
    fail "Mv" $c 5
  elif [[ ! $(cat $c/target/file) = "target" ]]; then
    fail "Mv" $c 6
  elif [[ ! -f $c/target/moved ]]; then
    fail "Mv" $c 7
  elif [[ ! $(cat $c/target/moved) = "src" ]]; then
    fail "Mv" $c 8
  elif [[ ! -d $c/target/dir ]]; then
    fail "Mv" $c 9
  elif [[ ! -L $c/target/symlink ]]; then
    fail "Mv" $c 10
  elif [[ ! -f $c/target/dir/file ]]; then
    fail "Mv" $c 11
  fi

  c=case02
  if [[ ! -f $c/src/file ]]; then
    fail "Mv" $c 1
  elif [[ ! -d $c/src/dir ]]; then
    fail "Mv" $c 2
  elif [[ -a $c/src/symlink ]]; then
    fail "Mv" $c 3
  elif [[ ! -f $c/src/dir/file ]]; then
    fail "Mv" $c 4
  elif [[ ! -f $c/target/file ]]; then
    fail "Mv" $c 5
  elif [[ ! -L $c/target/moved ]]; then
    fail "Mv" $c 6
  elif [[ ! $(readlink $c/target/moved) = "$tests_mv/$c/src/file" ]]; then
    fail "Mv" $c 7
  elif [[ ! -d $c/target/dir ]]; then
    fail "Mv" $c 8
  elif [[ ! -L $c/target/symlink ]]; then
    fail "Mv" $c 9
  elif [[ ! $(readlink $c/target/symlink) = "$tests_mv/$c/target/file" ]]; then
    fail "Mv" $c 10
  elif [[ ! -f $c/target/dir/file ]]; then
    fail "Mv" $c 11
  fi

  unchanged_Mv case03
  unchanged_Mv case04

  c=case05
  if [[ ! -f $c/src/file ]]; then
    fail "Mv" $c 1
  elif [[ -a $c/src/dir ]]; then
    fail "Mv" $c 2
  elif [[ ! -L $c/src/symlink ]]; then
    fail "Mv" $c 3
  elif [[ -a $c/src/dir/file ]]; then
    fail "Mv" $c 4
  elif [[ ! -f $c/target/file ]]; then
    fail "Mv" $c 5
  elif [[ ! -d $c/target/dir ]]; then
    fail "Mv" $c 6
  elif [[ ! -d $c/target/moved ]]; then
    fail "Mv" $c 7
  elif [[ ! -f $c/target/moved/file ]]; then
    fail "Mv" $c 8
  elif [[ ! -L $c/target/symlink ]]; then
    fail "Mv" $c 9
  elif [[ ! -f $c/target/dir/file ]]; then
    fail "Mv" $c 10
  fi

  unchanged_Mv case06

  c=case07
  if [[ -a $c/src/file ]]; then
    fail "Mv" $c 1
  elif [[ ! -d $c/src/dir ]]; then
    fail "Mv" $c 2
  elif [[ ! -L $c/src/symlink ]]; then
    fail "Mv" $c 3
  elif [[ ! -f $c/src/dir/file ]]; then
    fail "Mv" $c 4
  elif [[ ! -f $c/target/file ]]; then
    fail "Mv" $c 5
  elif [[ ! -d $c/target/dir ]]; then
    fail "Mv" $c 6
  elif [[ ! -L $c/target/symlink ]]; then
    fail "Mv" $c 7
  elif [[ ! -f $c/target/dir/file ]]; then
    fail "Mv" $c 8
  fi

  unchanged_Mv case08

  c=case09
  if [[ -a $c/src/file ]]; then
    fail "Mv" $c 1
  elif [[ ! -d $c/src/dir ]]; then
    fail "Mv" $c 2
  elif [[ ! -L $c/src/symlink ]]; then
    fail "Mv" $c 3
  elif [[ ! -f $c/src/dir/file ]]; then
    fail "Mv" $c 4
  elif [[ ! -f $c/target/file ]]; then
    fail "Mv" $c 5
  elif [[ ! -d $c/target/dir ]]; then
    fail "Mv" $c 6
  elif [[ ! -L $c/target/symlink ]]; then
    fail "Mv" $c 7
  elif [[ ! -f $c/target/dir/file ]]; then
    fail "Mv" $c 8
  elif [[ ! -d $c/target/nested ]]; then
    fail "Mv" $c 9
  elif [[ ! -f $c/target/nested/moved ]]; then
    fail "Mv" $c 10
  fi

  cd $src
}

function unchanged_Mv {
  local c=$1

  if [[ ! -f $c/src/file ]]; then
    fail "Mv" $c 1
  elif [[ ! -d $c/src/dir ]]; then
    fail "Mv" $c 2
  elif [[ ! -L $c/src/symlink ]]; then
    fail "Mv" $c 3
  elif [[ ! -f $c/src/dir/file ]]; then
    fail "Mv" $c 4
  elif [[ ! -f $c/target/file ]]; then
    fail "Mv" $c 5
  elif [[ ! -d $c/target/dir ]]; then
    fail "Mv" $c 6
  elif [[ ! -L $c/target/symlink ]]; then
    fail "Mv" $c 7
  elif [[ ! -f $c/target/dir/file ]]; then
    fail "Mv" $c 8
  fi
}
