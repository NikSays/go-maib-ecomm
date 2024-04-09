set -euo pipefail

mkdir -p certs

echo Generating CA
openssl req -new -x509 -config openssl.cnf \
  -keyout certs/ca.key -out certs/ca.crt \
  -passout pass:password


echo Generating server key
openssl genrsa -out certs/server.key 4096

echo Generating server certificate request
openssl req -new -config openssl.cnf \
  -key certs/server.key -out certs/server.csr

echo Generating server certificate
openssl x509 -req -days 9999 \
  -CA certs/ca.crt -CAkey certs/ca.key \
  -extfile openssl.cnf -extensions v3_req \
  -passin pass:password \
  -in certs/server.csr -out certs/server.crt


echo Generating client key
openssl genrsa -out certs/client.key 4096

echo Generating client certificate request
openssl req -new -config openssl.cnf \
   -key certs/client.key -out certs/client.csr

echo Generating client certificate
openssl x509 -req -days 9999 \
  -CA certs/ca.crt -CAkey certs/ca.key \
  -passin pass:password \
  -in certs/client.csr -out certs/client.crt


echo Transforming into pkcs#12
echo You may need to add -legacy option if this fails.
openssl pkcs12 -export \
  -passout pass:password \
  -inkey certs/client.key -in certs/client.crt -out certs/client.pfx

echo Removing unneeded files
rm certs/client.crt certs/client.key certs/client.csr certs/server.csr
