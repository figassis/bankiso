#!/bin/bash
echo "" > iso20022.go
for i in ../iso/*.go; do cat "$i" >> iso20022.go;done