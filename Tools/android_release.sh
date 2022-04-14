#!/bin/bash

gogio -target android -icon assets/meta/icon.png -version $1 -minsdk 29 .
$ANDROID_HOME/build-tools/29.0.2/apksigner sign --ks ~/.android/sign.keystore treman.apk
