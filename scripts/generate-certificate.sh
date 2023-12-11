#!/bin/bash

openssl genrsa -des3 -passout pass:x -out ssl/plutoengine.pass.key 2048
openssl rsa -passin pass:x -in ssl/plutoengine.pass.key -out ssl/plutoengine.key
rm ssl/plutoengine.pass.key
openssl req -new -key ssl/plutoengine.key -out ssl/plutoengine.csr \
    -subj "/C=IR/O=PlutoEngine/CN=example.com"
openssl x509 -req -days 365 -in ssl/plutoengine.csr -signkey ssl/plutoengine.key -out ssl/plutoengine.crt