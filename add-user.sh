#!/bin/bash

if [ -z $1 ]
then
  echo Please specify the user
  echo Exiting...
  exit 1
fi

adduser $1

read -r -p "Give $1 sudo? " response
case "$response" in
  [yY][eE][sS]|[yY]) 
    usermod -a -G sudo $1
    ;;
  *)
    exit 0
    ;;
esac

exit 0