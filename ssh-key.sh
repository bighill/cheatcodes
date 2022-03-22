#!/bin/bash

read -r -p "Create ssh key for $(whoami)? " response
case "$response" in
  [yY][eE][sS]|[yY]) 
    ssh-keygen -t ed25519
    ;;
  *)
    exit 0
    ;;
esac

exit 0