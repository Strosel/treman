#!/bin/bash

gogio -target android -icon assets/meta/icon.png -version $1 -minsdk 29 .
apksigner sign --ks ~/.android/sign.keystore treman.apk