#!/bin/bash

gogio -target android .
if [ $? -eq 0 ]; then
    adb uninstall com.github.treman > /dev/null 2>&1
    adb install treman.apk
fi