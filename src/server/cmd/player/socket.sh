#!/usr/bin/env bash

echo '{ "command": ["get_property", "volume"] }' | socat - /tmp/mpvsocket
#echo '{ "command": ["set_property", "pause", false] }' | socat - /tmp/mpvsocket