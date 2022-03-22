#!/bin/bash

if [ -z $1 ]
then
  echo Please specify the user
  echo Exiting...
  exit 1
fi

read -r -p "Are you sure you want to delete user $1 and home directory? [y/N] " response
case "$response" in
  [yY][eE][sS]|[yY]) 
    killall -9 -u $1
    deluser --remove-home $1
    ;;
  *)
    exit 0
    ;;
esac

exit 0