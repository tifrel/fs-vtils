#!/usr/bin/bash

tests_cp=$src/testdata/Cp

# // case01: p is file => no error
# // case02: p is symlink & no d flag => no error
# // case03: p is symlink & d flag => no error
# // case04: p is dir & no r flag => error
# // case05: p is dir & r flag => no error
# // case06: p is symlink to dir & d flag & no r flag => no error
# // case07: p is symlink to dir & d flag & r flag => no error
# // case08: p is dir with symlinks & no d flag => no error
# // case09: p is dir with symlinks & d flag => no error
# //
# // case10: target exists & no f flag => error
# // case11: target exists & f flag => no error
# // case12: targetdir doesn't exist & no p flag => error
# // case13: targetdir doesn't exist & p flag => no error


function reset_Cp {
  cd $tests_cp

  for i in {1..13}; do

    if [[ $i -lt 10 ]]; then
      c=case0$i
    else
      c=case$i
    fi

    rm -R $c
    mkdir $c

    touch $c/file
    ln -s $tests_cp/$c/file $c/filelink
    mkdir $c/dir
    touch $c/dir/file
    ln -s $tests_cp/$c/dir $c/dirlink
    mkdir $c/dir2
    ln -s $tests_cp/$c/file $c/dir2/link
  done

  cd $src
}

function evaluate_Cp {
  cd $tests_cp

  c=case01
  unchanged_Cp $c
  if [[ ! -f $c/copied ]]; then
    fail "Cp" $c 11
  fi

  c=case02
  unchanged_Cp $c
  check_softlink $c/copied $c/file "Cp" $c 11

  c=case03
  unchanged_Cp $c
  if [[ ! -f $c/copied ]]; then
    fail "Cp" $c 11
  fi

  c=case04
  unchanged_Cp $c

  c=case05
  unchanged_Cp $c
  if [[ ! -d $c/copied ]]; then
    fail "Cp" $c 11
  elif [[ ! -f $c/copied/file ]]; then
    fail "Cp" $c 12
  fi

  c=case06
  unchanged_Cp $c

  c=case07
  unchanged_Cp $c
  if [[ ! -d $c/copied ]]; then
    fail "Cp" $c 11
  elif [[ ! -f $c/copied/file ]]; then
    fail "Cp" $c 12
  fi

  c=case08
  unchanged_Cp $c
  if [[ ! -d $c/copied ]]; then
    fail "Cp" $c 11
  fi
  check_softlink $c/copied/link $c/file "Cp" $c 12

  c=case09
  unchanged_Cp $c
  if [[ ! -d $c/copied ]]; then
    fail "Cp" $c 11
  elif [[ ! -f $c/copied/link ]]; then
    fail "Cp" $c 12
  fi

  c=case10
  unchanged_Cp $c

  c=case11
    if [[ ! -f $c/file ]]; then
    fail "Cp" $c 1
  elif [[ ! -L $c/filelink ]]; then
    fail "Cp" $c 2
  elif [[ ! -L $c/dirlink ]]; then
    fail "Cp" $c 3
  elif [[ ! -f $c/dir ]]; then
    fail "Cp" $c 4
  elif [[ ! -d $c/dir2 ]]; then
    fail "Cp" $c 6
  elif [[ ! -L $c/dir2/link ]]; then
    fail "Cp" $c 7
  fi
  check_softlink $c/filelink $c/file "Cp" $c 8
  check_softlink $c/dirlink $c/dir "Cp" $c 9
  check_softlink $c/dir2/link $c/file "Cp" $c 10

  c=case12
  unchanged_Cp $c

  c=case13
  unchanged_Cp $c
    if [[ ! -d $c/newdir ]]; then
    fail "Cp" $c 1
  elif [[ ! -f $c/newdir/copied ]]; then
    fail "Cp" $c 2
  fi



  cd $src
}

function unchanged_Cp {
  local c=$1

  if [[ ! -f $c/file ]]; then
    fail "Cp" $c 1
  elif [[ ! -L $c/filelink ]]; then
    fail "Cp" $c 2
  elif [[ ! -L $c/dirlink ]]; then
    fail "Cp" $c 3
  elif [[ ! -d $c/dir ]]; then
    fail "Cp" $c 4
  elif [[ ! -f $c/dir/file ]]; then
    fail "Cp" $c 5
  elif [[ ! -d $c/dir2 ]]; then
    fail "Cp" $c 6
  elif [[ ! -L $c/dir2/link ]]; then
    fail "Cp" $c 7
  fi
  check_softlink $c/filelink $c/file "Cp" $c 8
  check_softlink $c/dirlink $c/dir "Cp" $c 9
  check_softlink $c/dir2/link $c/file "Cp" $c 10

}
