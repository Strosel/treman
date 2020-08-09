#!/bin/bash

gogio -target android -minsdk 29 .
apksigner sign --ks ~/.android/sign.keystore treman.apk