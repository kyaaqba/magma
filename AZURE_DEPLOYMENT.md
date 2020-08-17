# Orchestrator Deployment to Azure

This document lays out the steps to get Orchestrator deployed and configured in Azure. It is assumed that you have the following installed on a Windows machine (refer to DEV_ENV_SETUP.md):

- Azure CLI (az) for Windows
- Windows Subsystem for Linux (WSL)
- Azure CLI for Linux
- Kubernetes CLI (kubectl) for Windows
- PowerShell

## Azure CLI Authentication

In order to use the Azure CLI, we first need to connect and authenticate with the tenant. Run the command below replacing and enter your credentials when prompted:

`az login --tenant TENANT_ID`

## Key Vault

### Creating the Resource

There are a few items that need to be present before deployment of the pipeline.  This includes definition of both a Server application and Client application within Azure AD (AAD).  Kubernetes will acquire tokens for these applications by using secrets that will be stored in Azure Key Vault and deployed to AKS secrets during the ARM template deployment.
To create the Key Vault instance, create these AAD applications and store their secrets as Key Vault secrets, run the PowerShell script: *`orc8r\cloud\deploy\azure\create-identities.ps1`* as a user that has both Azure AD global admin right as well as at least a contributor role in the resource group where the Key Vault instance will be created.
Here is an example usage of the script with the required parameters:

`.\create-identities.ps1 -rgName sonar-prod-magma -keyVaultName sonar-prod-magma-kv-1 -aksName magma-aks`

---

*Note: If for some reason, you need to re-run the script, please make sure and delete any AAD app registrations that were created previously so that the script will attempt to re-create them.*

---

### Creating the Secret

From a bash shell, set the current directory to a local location of the certificates by replacing the SECRET_VALUE below with the password for the certificate and running the command.

`az keyvault secret set --name SSLKeyPassword --value SECRET_VALUE --vault-name sonar-prod-magma-01`

### Packaging the Files

After creating the application certificates using the *`create_application_certs.sh`* script as described in the [Orchestrator installation documentation](https://magma.github.io/magma/docs/orc8r/deploy_install#ssl-certificates), we need to package the resulting files in order to import them. From a bash shell, run the following commands:

`openssl pkcs12 -inkey certifier.key -in certifier.pem -export -out certifier.pfx`

`openssl pkcs12 -inkey fluentd.key -in fluentd.pem -export -out fluentd.pfx`

`openssl req -x509 -new -nodes -key bootstrapper.key -sha256 -days 3650 -out bootstrapper.pem -subj "/C=US/CN=bootstrapper.sonarlte.com"`

`openssl pkcs12 -inkey bootstrapper.key -in bootstrapper.pem -export -out bootstrapper.pfx`

### Importing the Certificates

Now that the certificate files are packaged, we can import them into the Key Vault by running these commands:

`az keyvault certificate import --file .\sonar.pfx --name sonarlte-com --vault-name sonar-prod-magma-01`

`az keyvault certificate import --file .\certifier.pfx --name Certifier --vault-name sonar-prod-magma-01`

`az keyvault certificate import --file .\fluentd.pfx --name Fluentd --vault-name sonar-prod-magma-01`

`az keyvault certificate import --file .\bootstrapper.pfx --name BootstrapperKey --vault-name sonar-prod-magma-01`

### Updating the Secret IDs

The pipeline runs a script called *`fetch_certs.sh`* to prepare the certificates for deployment. We need to ensure the secret ID for each certificate is correct in the script. Open the script and replace the values in the --id parameters of the `az keyvault secret download` commands with the values in the Azure portal under:

**`[Key Vault resource] > Certificates > [certificate name] > Current Version > Secret Identifier`**

---

*Note: This process to be automated/improved.*

---

## Architecture

Run the pipeline that deploys the architecture to your Azure tenant via the ARM templates.

Alternatively, the ARM templates can be deployed with the Azure CLI.

To deploy the net template, you can use the command below to validate and then change `validate` to `deploy` when ready to deploy:

`az deployment group validate --template-file .\magma\orc8r\cloud\deploy\azure\arm_templates\net_template.json --parameters .\magma\orc8r\cloud\deploy\azure\arm_templates\net_template.parameters.json -g sonar-prod-magma`

To deploy the generic template, use the command below to validate and then change `validate` to `deploy` when ready to deploy:

`az deployment group validate --template-file .\magma\orc8r\cloud\deploy\azure\arm_templates\gen_template.json --parameters .\magma\orc8r\cloud\deploy\azure\arm_templates\gen_template.parameters.json -g sonar-prod-magma`

Each of these should execute and build all appropriate items needed for each resource group.

## Container Registry

Run the command below to create the Azure Container Registry:

`az acr create --resource-group sonar-prod-magma --name SonarRegistry --sku Basic`

Then run the following command to connect the ACR to AKS:

`az aks update -n magma-aks -g sonar-prod-magma --attach-acr SonarRegistry`

## Kubernetes CLI Authentication

Run the command below to generate AKS administrator credentials and connect the kubectl CLI to the cluster:

`az aks get-credentials --name magma-aks --resource-group sonar-prod-magma --admin`

---

*Note: Admin credentials only recommended for non-production environments.*

---

Then run the following command to create the AKS namespace:

`kubectl create namespace magma-stage`

## Pipeline Configuration

### Creating the AKS Service Connection

We need to create a service connection in Azure DevOps to allow the pipeline to complete the tasks needed in the AKS cluster.

In the Azure DevOps portal, click **Project settings** in the lower left corner of the page and then click **Service connections** in the Project Settings blade.

Click the **New service connection** button in the upper right corner of the page and select **Kubernetes** from the connection type list. Click **Next**.

Ensure the authentication method selected is **Azure Subscription** and then select the target subscription from the dropdown. You may be prompted for credentials.

Once a subscription is selected, the Cluster dropdown will populate with available AKS clusters. Select the target cluster, then select the AKS namespace that we created in the above section from the dropdown, and check the box to use cluster admin credentials.

Give the service connection a name, check the box to grant permission to all pipelines, and then click **Save**.

### Creating the Pipeline

In Azure DevOps portal, click **Pipelines** in the left side of the page and then click the **New pipeline** button.

Click **Use the classic editor** on the next page. On the following page, select **GitHub** as the source and either add or select the appropriate connection. Select the correct repository and branch to deploy from and then click the **Continue** button.

On the template page, choose **YAML**.

On the next page, give the pipeline a name, then browse the path to the *`azure-container.yml`* file and select it:

**`orc8r/cloud/deploy/azure/container-pipeline.yml`**

On the Triggers tab, you can override and disable continuous integration if you only wish to have the pipeline run when triggered manually.

Click the **Save & queue** dropdown arrow and select **Save**.

## Deployment

At this point, we should be ready to run the pipeline. Open the pipeline in Azure DevOps and click the **Run pipeline** button.

From the parameters blade, we can define some values to be used during execution of the pipeline. Ensure that everything is correct and that the appropriate stages are selected in the ***Stages to run*** area.

---

*Time saving tip: The `Build_And_Publish` stage builds and publishes the docker images to the container registry so it may not be necessary to run this stage every time you trigger a deployment.*

---
