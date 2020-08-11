mkdir orc8r/cloud/helm/orc8r/charts/secrets/certs
cd orc8r/cloud/helm/orc8r/charts/secrets/certs

echo "******Acquiring Sonarlte Certificate******"

echo "Downloading pfx (no password) from Key vault..."
az keyvault secret download -f sonar-nopass.pfx --vault-name $2 --id https://sonar-prod-magma-01.vault.azure.net/secrets/sonarlte-com/dabfd633beb74baea2dfefcd57a806cc --encoding base64

echo "Converting pfx (no password) to pem format..."
openssl pkcs12 -in sonar-nopass.pfx -out sonar-nopass.pem -nodes -password pass:""

echo "Converting pem to pfx with the password..."
openssl pkcs12 -export -out sonar-lte.pfx -in sonar-nopass.pem -password pass:"$1"

echo "Converting pfx (with password) to pem..."
openssl pkcs12 -in sonar-lte.pfx -out rootCA.pem -nodes -password pass:"$1"

echo "Extracting private key from pfx..."
openssl pkcs12 -in sonar-lte.pfx -nocerts -out controller-encrypted.key -passin pass:"$1" -passout pass:"$1"

echo "Decrypting private key..."
openssl rsa -in controller-encrypted.key -out controller.key -passin pass:"$1"

echo "Extracting public certificate from pfx..."
openssl pkcs12 -in sonar-lte.pfx -clcerts -nokeys -out controller.crt -password pass:"$1"

echo "*******************************************"

echo "******Acquiring Certifier Certificate******"

echo "Downloading pfx (no password) from Key vault..."
az keyvault secret download -f certifier-nopass.pfx --vault-name $2 --id https://sonar-prod-magma-01.vault.azure.net/secrets/Certifier/05b28b14abc348ccaac5602b7fb4b2be --encoding base64

echo "Converting pfx (no password) to pem format..."
openssl pkcs12 -in certifier-nopass.pfx -out certifier-nopass.pem -nodes -password pass:""

echo "Converting pem to pfx with the password..."
openssl pkcs12 -export -out certifier.pfx -in certifier-nopass.pem -password pass:"$1"

echo "Converting pfx (with password) to pem..."
openssl pkcs12 -in certifier.pfx -out certifier.pem -nodes -password pass:"$1"

echo "Extracting private key from pfx..."
openssl pkcs12 -in certifier.pfx -nocerts -out certifier-encrypted.key -passin pass:"$1" -passout pass:"$1"

echo "Decrypting private key..."
openssl rsa -in certifier-encrypted.key -out certifier.key -passin pass:"$1"

echo "*****************************************"

echo "******Acquiring Fluentd Certificate******"

echo "Downloading pfx (no password) from Key vault..."
az keyvault secret download -f fluentd-nopass.pfx --vault-name $2 --id https://sonar-prod-magma-01.vault.azure.net/secrets/Fluentd/c01a57ceb15749e899fc0a1a2adbb604 --encoding base64

echo "Converting pfx (no password) to pem format..."
openssl pkcs12 -in fluentd-nopass.pfx -out fluentd-nopass.pem -nodes -password pass:""

echo "Converting pem to pfx with the password..."
openssl pkcs12 -export -out fluentd.pfx -in fluentd-nopass.pem -password pass:"$1"

echo "Converting pfx (with password) to pem..."
openssl pkcs12 -in fluentd.pfx -out fluentd.pem -nodes -password pass:"$1"

echo "Extracting private key from pfx..."
openssl pkcs12 -in fluentd.pfx -nocerts -out fluentd-encrypted.key -passin pass:"$1" -passout pass:"$1"

echo "Decrypting private key..."
openssl rsa -in fluentd-encrypted.key -out fluentd.key -passin pass:"$1"

echo "**************************************"

echo "******Acquiring Bootstrapper Key******"

echo "Downloading pfx (no password) from Key vault..."
az keyvault secret download -f bootstrapper-nopass.pfx --vault-name $2 --id https://sonar-prod-magma-01.vault.azure.net/secrets/BoostrapperKey/c643eb0765ca49fd9f1b333938670d06 --encoding base64

echo "Converting pfx (no password) to pem format..."
openssl pkcs12 -in bootstrapper-nopass.pfx -out bootstrapper-nopass.pem -nodes -password pass:""

echo "Converting pem to pfx with the password..."
openssl pkcs12 -export -out bootstrapper.pfx -in bootstrapper-nopass.pem -password pass:"$1"

echo "Converting pfx (with password) to pem..."
openssl pkcs12 -in bootstrapper.pfx -out bootstrapper.pem -nodes -password pass:"$1"

echo "Extracting private key from pfx..."
openssl pkcs12 -in bootstrapper.pfx -nocerts -out bootstrapper-encrypted.key -passin pass:"$1" -passout pass:"$1"

echo "Decrypting private key..."
openssl rsa -in bootstrapper-encrypted.key -out bootstrapper.key -passin pass:"$1"

echo "*************************************"
