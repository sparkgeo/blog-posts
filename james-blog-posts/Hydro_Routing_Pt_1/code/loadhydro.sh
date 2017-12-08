#!/usr/bin/env bash

# To be run from the same directory as the zipped shapefiles

# Get user input
echo Please enter a Database name to connect to:
read databasename

# Unzip all compressed files in this directory
unzip '*.zip'

# Create the table scheme for the edges.
# Find the first occurrence of a shapefile to use as schema
files_edge=(*NLFLOW*.shp)
shp2pgsql -I -p -s 4140:26910 ${files_edge[0]} public.nhn_08_edges | psql -d $databasename

# Load Edges into database
for f in *NLFLOW*.shp;
do
    shp2pgsql -I -a -s 4140:26910 $f nhn_08_edges > `basename $f .shp`.sql
done;

for f in *.sql;
do
    psql -d $databasename -f $f
done;

# Cleanup SQL
rm *.sql

# Create the table scheme for the edges.
# Find the first occurrence of a shapefile to use as schema
files_node=(*HYDROJUNCT*.shp)
shp2pgsql -I -p -s 4140:26910 ${files_node[0]} public.nhn_08_nodes | psql -d $databasename

# Load Edges into database
for f in *HYDROJUNCT*.shp;
do
    shp2pgsql -I -a -s 4140:26910 $f nhn_08_nodes > `basename $f .shp`.sql
done;

for f in *.sql;
do
    psql -d $databasename -f $f
done;

# Cleanup SQL
rm *.sql

# Cleanup Shapefiles
rm *.shp
rm *.shx
rm *.dbf
rm *.prj
rm *.xml