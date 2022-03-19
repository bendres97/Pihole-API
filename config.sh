#!/bin/bash

# Show env vars
grep -v '^#' .env

# Export env vars
export $(grep -v '^#' .env | xargs)

mkdir -p ${DATA_DIR}
id_file="${DATA_DIR}/id_rsa"
hosts_file="${DATA_DIR}/hosts.conf"

ssh-keygen -f $id_file -P ''

add=true

while $add; do
    read -p 'Hostname: ' hostname
    read -p 'Username: ' username

    ssh-copy-id -i $id_file "$username@$hostname"

    echo "$username@$hostname" >> $hosts_file

    read -p 'Would you like to add another host? (y/n):' more

    if [ $more != 'y' ]; then
        add=false
    fi
done