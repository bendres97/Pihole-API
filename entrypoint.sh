#!/bin/bash

HOSTS="/app/hosts.conf"

if test -f $HOSTS; then
    python3 main.py
else
    ssh-keygen -f /root/.ssh/id_rsa -P ''

    add=true

    while $add; do
        read -p 'Hostname: ' hostname
        read -p 'Username: ' username

        ssh-copy-id "$username@$hostname"

        echo "$username@$hostname" >>hosts.conf

        read -p 'Would you like to add another host? (y/n):' more

        if [ $more != 'y' ]; then
            add=false
        fi
    done

    echo "You are now ready for deployment. If you did not set this container to restart automatically, you will need to run 'docker start <container>'"
fi
