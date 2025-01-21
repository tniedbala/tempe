#!/bin/bash

rm -rf $WORKSPACE/bin/grammar

antlr -o $WORKSPACE/bin/grammar -visitor $WORKSPACE/grammar/TemplateLexer.g4
javac $WORKSPACE/bin/grammar/*.java

antlr -o $WORKSPACE/bin/grammar -lib $WORKSPACE/bin/grammar \
    -visitor $WORKSPACE/grammar/TemplateParser.g4
javac $WORKSPACE/bin/grammar/*.java
