#!/bin/bash

if [ -z $1 ]
then
  echo Please specify the hostname
  echo Exiting...
  exit 1
fi

hostnamectl set-hostname $1

echo Now reboot to make it official

exit 0