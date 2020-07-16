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
az keyvault set-policy --name $keyVaultName --object-id $principalId --secret-permissions get list

