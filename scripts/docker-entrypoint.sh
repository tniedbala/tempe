#!/bin/bash

shopt -s expand_aliases
. ~/.bashrc

case $1 in 
    go)
        case $2 in
            build)
                . scripts/build-go.sh
                ;;
            test)
                shift; shift; go run test $@
                ;;
        esac
    ;;
    antlr)
        case $2 in
            test)
                . scripts/build-java.sh
                shift; shift; antlr-test $@
                ;;
            copy)
                cp $ANTLR_JAR_PATH $WORKSPACE/bin
                ;;
        esac
    ;;
    java)
        case $2 in 
            build)
                . scripts/build-java.sh
            ;;
        esac
esac