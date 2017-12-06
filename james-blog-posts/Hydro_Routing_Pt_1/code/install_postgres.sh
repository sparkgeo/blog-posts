#!/usr/bin/env bash

sudo apt-get install postgresql-10 postgresql-10-contrib postgresql-10-postgis-2.4 postgresql-10-pgrouting

sudo su postgres
psql postgres