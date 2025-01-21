#!/bin/bash

GRAMMAR_DIR=$WORKSPACE/grammar
PACKAGE_DIR=$WORKSPACE/pkg/parser/base

rm -rf $PACKAGE_DIR
antlr -o $PACKAGE_DIR -Dlanguage=Go -package base -visitor -no-listener $GRAMMAR_DIR/TemplateLexer.g4
antlr -o $PACKAGE_DIR -Dlanguage=Go -package base -visitor -no-listener \
    -lib $WORKSPACE/bin/grammar -visitor $GRAMMAR_DIR/TemplateParser.g4

