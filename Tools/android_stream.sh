#!/bin/bash

gogio -target android -minsdk 29 .
if [ $? -eq 0 ]; then
    adb uninstall com.github.treman > /dev/null 2>&1
    adb install treman.apk
fi