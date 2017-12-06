#!/usr/bin/env bash

# Create the file /etc/apt/sources.list.d/pgdg.list, and add a line for the repository 
echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" >> /etc/apt/sources.list.d/pgdg.list

# Import the repository signing key, and update the package lists 
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add - 
sudo apt-get update