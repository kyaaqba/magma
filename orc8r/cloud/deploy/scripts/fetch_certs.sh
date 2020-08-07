echo "Downloading pfx from Key vault with no password..."
az keyvault secret download -f sonar-nopass.pfx --vault-name $3 --id $2 --encoding base64

echo "Converting pfx (no password) to pem format..."
openssl pkcs12 -in sonar-nopass.pfx -out temp-nopass.pem -nodes -password pass:""

echo "Converting pem to pfx with the password..."
openssl pkcs12 -export -out sonar-lte.pfx -in temp-nopass.pem -password pass:"$1"

echo "Converting pfx (with password) to pem..."
openssl pkcs12 -in sonar-lte.pfx -out rootCA.pem -nodes -password pass:"$1"

echo "Extracting private key from pfx..."
openssl pkcs12 -in sonar-lte.pfx -nocerts -out controller.key -passin pass:"$1" -passout pass:"$1"

echo "Decrypting private key..."
openssl rsa -in sonar-lte.key -out controller-decrypted.key -passin pass:"$1"

echo "Extracting public certificate from pfx..."
openssl pkcs12 -in sonar-lte.pfx -clcerts -nokeys -out controller.crt -password pass:"$1"

# bash create_application_certs.sh $4