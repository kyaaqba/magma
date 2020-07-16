param(
        [Parameter(Mandatory)]
        [string]$rgName,

        [Parameter(Mandatory)]
        [string]$keyVaultName,

        [Parameter(Mandatory)]
        [string]$principalId,

        [Parameter()]
        [string]$location = 'eastus'
    )

az keyvault create --location eastus --name $keyVaultName --resource-group $rgName
az keyvault set-policy --name $keyVaultName --object-id $principalId --secret-permissions get list set
az keyvault secret set --vault-name $keyVaultName --name 'ServicePrincipalId' --value '[id value]'
az keyvault secret set --vault-name $keyVaultName --name 'ServicePrincipalSecret' --value '[secret value]'
az keyvault secret set --vault-name $keyVaultName --name 'ServerId' --value '[id value]'
az keyvault secret set --vault-name $keyVaultName --name 'ServerSecret' --value '[secret value]'
az keyvault secret set --vault-name $keyVaultName --name 'ClientId' --value '[id value]'
az keyvault secret set --vault-name $keyVaultName --name 'ClientSecret' --value '[secret value]'