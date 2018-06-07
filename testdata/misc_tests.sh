#!/usr/bin/bash

tests_misc=$src/testdata/misc

function reset_misc {
  rm -R $tests_misc/*
  cd $tests_misc

  mkdir dir
  echo "contents" > a
  echo "contents" > b
  echo "other-contents-and-yadda-yadda" > c
  ln $(pwd)/a d
  ln -s $(pwd)/a e
  ln -s $(pwd)/e f

  # for i in {1..11}; do

  #   if [[ $i -lt 10 ]]; then
  #     c=case0$i
  #   else
  #     c=case$i
  #   fi

  #   rm -R $c
  #   mkdir $c

  #   mkdir $c/dir
  #   echo "contents" > $c/a
  #   echo "contents" > $c/b
  #   echo "other-contents" > $c/c
  #   ln $c/a $c/d

  # done

  cd $src
}

# function evaluate_misc {
#   cd $tests_misc

#   #case01
#   c=case01
#   if [[ ! -f $c/file ]]; then
#     fail "misc" $c 1
#   elif [[ ! $(cat $c/file)="a\nb\nc" ]]; then
#     fail "misc" $c 2
#   elif [[ ! -f $c/empty ]]; then
#     fail "misc" $c 3
#   elif [[ ! $(cat $c/empty)="written" ]]; then
#     fail "misc" $c 4
#   fi

#   #case02
#   c=case02
#   if [[ ! -f $c/file ]]; then
#     fail "misc" $c 1
#   elif [[ ! $(cat $c/file)="written" ]]; then
#     fail "misc" $c 2
#   elif [[ ! -f $c/empty ]]; then
#     fail "misc" $c 3
#   elif [[ ! $(cat $c/empty)="" ]]; then
#     fail "misc" $c 4
#   fi

#   #case03
#   c=case03
#   if [[ ! -f $c/file ]]; then
#     fail "misc" $c 1
#   elif [[ ! $(cat $c/file)="a\nb\nc" ]]; then
#     fail "misc" $c 2
#   elif [[ ! -f $c/empty ]]; then
#     fail "misc" $c 3
#   elif [[ -ne $(cat $c/empty) ]]; then
#     fail "misc" $c 4
#   # elif [[ ! -f $c/new ]]; then
#   #   fail "misc" $c 5
#   # elif [[ ! $(cat $c/new)="written" ]]; then
#   #   fail "misc" $c 6
#   fi


#   #case04
#   unchanged_misc case04

#   #case05
#   unchanged_misc case05

#   #case06
#   unchanged_misc case06

#   c=case07
#   if [[ ! -f $c/file ]]; then
#     fail "misc" $c 1
#   elif [[ ! $(cat $c/file)="a\nb\nc" ]]; then
#     fail "misc" $c 2
#   elif [[ ! -f $c/empty ]]; then
#     fail "misc" $c 3
#   elif [[ ! $(cat $c/empty)="appended" ]]; then
#     fail "misc" $c 4
#   fi

#   c=case08
#   if [[ ! -f $c/file ]]; then
#     fail "misc" $c 1
#   elif [[ ! $(cat $c/file)="a\nb\nc\nappended" ]]; then
#     fail "misc" $c 2
#   elif [[ ! -f $c/empty ]]; then
#     fail "misc" $c 3
#   elif [[ ! $(cat $c/empty)="" ]]; then
#     fail "misc" $c 4
#   fi

#   unchanged_misc case09

#   c=case10
#   if [[ ! -f $c/file ]]; then
#     fail "misc" $c 1
#   elif [[ ! $(cat $c/file)="a\nb\nc" ]]; then
#     fail "misc" $c 2
#   elif [[ ! -f $c/empty ]]; then
#     fail "misc" $c 3
#   elif [[ ! $(cat $c/empty)="a\nb\nc" ]]; then
#     fail "misc" $c 4
#   fi

#   unchanged_misc case11

#   cd $src
# }

# function unchanged_misc {
#   local c=$1

#   if [[ ! -f $c/file ]]; then
#     fail "misc" $c 1
#   elif [[ ! $(cat $c/file)="a\nb\nc" ]]; then
#     fail "misc" $c 2
#   elif [[ ! -f $c/empty ]]; then
#     fail "misc" $c 3
#   elif [[ -ne $(cat $c/empty) ]]; then
#     fail "misc" $c 4
#   fi
# }
