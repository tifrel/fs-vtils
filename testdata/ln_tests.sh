#!/usr/bin/bash

tests_ln=$src/testdata/Ln

function reset_Ln {
  cd $tests_ln

  for i in {1..11}; do

    if [[ $i -lt 10 ]]; then
      c=case0$i
    else
      c=case$i
    fi

    rm -R $c
    mkdir $c


    touch $c/file
    ln -s $tests_ln/$c/file $c/oldlink
    mkdir $c/dir

  done

  cd $src
}

function check_hardlink {
  # $1 = supposed link
  # $2 = target
  # $3 = caller (tested function)
  # $4 = case
  # $5 = condition
  if [[ ! -f $1 ]]; then
    fail $3 $4 $5
  elif [[ ! $1 -ef $2 ]]; then
    fail $3 $4 $5
  fi
}

function check_softlink {
  # $1 = supposed link
  # $2 = target
  # $3 = caller (tested function)
  # $4 = case
  # $5 = condition
  if [[ ! -L $1 ]]; then
    echo $1
    echo $(ls -la $(dirname $1))
    echo "is not a symlink"
    fail $3 $4 $5
  elif [[ ! $(readlink $1) -ef $2 ]]; then
    echo "wrong indirection"
    fail $3 $4 $5
  fi
}

function evaluate_Ln {
  cd $tests_ln

  c=case01
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_softlink $c/newlink $c/file "Ln" $c 4

  c=case02
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_softlink $c/newlink $c/oldlink "Ln" $c 4

  c=case03
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_softlink $c/newlink $c/dir "Ln" $c 4

  c=case04
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_hardlink $c/newlink $c/oldlink "Ln" $c 4

  c=case05
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_softlink $c/newlink $c/file "Ln" $c 4

  c=case06
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_softlink $c/newlink $c/file "Ln" $c 4

  c=case07
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_hardlink $c/newlink $c/file "Ln" $c 4

  unchanged_Ln case08

  c=case09
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file2 "Ln" $c 3

  unchanged_Ln case10

  c=case11
  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi
  check_softlink $c/oldlink $c/file "Ln" $c 3
  check_softlink $c/newdir/newlink $c/file "Ln" $c 3

  cd $src
}

function unchanged_Ln {
  local c=$1

  if [[ ! -f $c/file ]]; then
    fail "Ln" $c 1
  elif [[ ! -d $c/dir ]]; then
    fail "Ln" $c 2
  fi

  check_softlink $c/oldlink $c/file "Ln" $c 3
}
