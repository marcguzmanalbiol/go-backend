#!/bin/bash
if [ $1 == "psql" ]
then 
    docker exec -it cantor_postgres $1 -U admin cantor
else
    docker exec -it cantor_postgres $1
fi