#!/bin/bash

git pull https://github.com/cjsmocjsmo/ats-htmx-echo;


# check if /usr/share/ats-htmx-echo exists if not create it
if [ ! -d /usr/share/ats-htmx-echo ]; then
    sudo mkdir /usr/share/ats-htmx-echo;
    chmod 755 /usr/share/ats-htmx-echo;
    chown -R porthose_cjsmo_cjsmo:porthose_cjsmo_cjsmo /usr/share/ats-htmx-echo;
fi

if [ -d /usr/share/ats-htmx-echo/Uploads ]; then
    mkdir /usr/share/ats-htmx-echo/Uploads;
    chmod 755 /usr/share/ats-htmx-echo/Uploads;
fi

if [ -f /usr/share/ats-htmx-echo/ats.db ]; then
    touch /usr/share/ats-htmx-echo/ats.db;
fi

