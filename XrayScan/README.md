# CI/CD Deep Dive with JFrog Pipelines

## Goal of this Lab

The goal of this lab is to get some hands on experience with the following Pipelines concepts:
* Simple YAML-based configuration with preconfigured steps (aka native steps)
* Usage of resources to create interconnected pipelines. i.e. “Pipeline of Pipelines”
* Using integrations to abstract secrets from pipeline configuration
* Templatizing your pipeline configuration to make it easy to reuse
* Creating a Docker-based pipeline to ship an application written in Go

This lab consists of 2 pipelines:

Pipeline 1: app_dev_pipeline
* Builds an Application’s Docker image
* Pushes the docker image to a DEV Repo in Artifactory
* Publishes the build information to Artifactory

Pipeline 2: app_prod_pipeline
* Promotes the docker image to a PROD Repo in Artifactory
* “Deploys” the image - for lab purposes, we will just download the image from PROD repo and run it

## Prerequisites

* Github account. Note the following and plan accordingly.
  * You would be forking a GitHub Repository to your account.
  * Webhooks would be automatically created on these Repositories when the pipeline is configured.
  * Create an API token (if you do not already have one) and keep it handy. Make sure the token you create has the permission scope as listed in here under the Token Permissions section.

* A SaaS Instance of JFrog. This will be provided as part of your enrollment to the Training class.

## Instructions

### Step 1: Fork required GitHub Repositories
The pipeline configuration and application code for this lab is in this repository:
https://github.com/jfrog/Swampup2021-su204-cicd-deep-dive

Please fork it to your Github account. Your forked repo’s Github URL will be **https://github.com/<your-account-name>/Swampup2021-su204-cicd-deep-dive** , where your-account-name is the account or organization name.

The repository consists of the following files:
* Go application: main_go , go_mod  
* Pipelines configuration: pipelines.yml , values.yml
* Dockerfile
* README.md

### Step 2: Create necessary Repositories in your Cloud Artifactory Server
In this step, you will create 2 Artifactory docker repositories where your pipeline will push images.
Create 2 local repositories of package type docker in Artifactory with the following Repository Key:
* su204-docker-staging-local
* su204-docker-prod-local

Instructions for creating these repositories are here: Local Docker repositories. You should see these repositories when you go to **Administration View | Repositories | Repositories and select the Local tab**.

![Artifactory Local Repositories](https://i.ibb.co/DgcvN1G/imageA.png)

### Step 3: Create necessary Integrations
In this step, you will add your Github and Artifactory credentials so that your pipeline can connect to these services during execution.

Go to **Administration View | Pipelines | Integrations** and create these Integrations.
* Github
  * Follow these instructions to create a Github integration. Enter github_int in the “Name” field. Please note that you will need to use the previously created API token to create this integration

* Artifactory
  * Follow these instructions to create an Artifactory integration. Enter artifactory_int in the “Name” field. Please note that you can click on “Get API Key” to automatically populate the API Key.

You will see these integrations in the table as shown below:

![Integrations](https://i.ibb.co/pvr3yhS/imageB.png)

### Step 4: Modify pipelines configuration to reflect your environment

The pipelines configuration for this lab is spread across 2 files:
* **pipelines.yml** contains the complete configuration for the pipelines. This config is written in template form, so **you do not have to make any changes to this file**.
* **values.yml** file contains the custom configuration that you will edit as part of this lab.

Please modify the **values.yml** file in your forked repository as follows:
* gitRepositoryPath: Modify this to reflect the account name where you forked the repo, E.g. if I fork the repo to github username manishas, the gitRepositoryPath is manishas/Swampup2021-su204-cicd-deep-dive. In general, the syntax is **<youraccountname>/Swampup2021-su204-cicd-deep-dive**
* artifactoryURL: Change this to reflect your SaaS server name. You can find this in your URL. Please do not include https:// in the server name the format for this entry is **<server-name>.jfrog.io**

The rest of the values do not need to be changed if you have followed instructions.

### Step 5: Create a Pipeline Source

Next, you will add your pipeline configuration so that it can be parsed to create your pipelines.

Go to **Administration View | Pipelines | Pipeline Sources** and create a new pipeline source. Click on **Add Pipeline Source**, choose **From YAML** and create a source as shown below:
* Choose **Single Branch**
* Choose **github_int** as SCM provider integration
* Choose your forked repository (**<youraccountname>/Swampup2021-su204-cicd-deep-dive**) for **Repository Full Name**
* Branch is **master**
* Pipeline config filter specifies which files contain pipeline configuration. This  should be set to **(pipelines|values).yml**

![Add a pipeline source](https://i.ibb.co/GJSdh3k/imageC.png)

To check the status of pipeline creation, go to **Administration View | Pipelines | Pipeline Sources** to view status.  If the status is **Success**, then your pipelines have been created successfully. If the status shows **Failed**, you can click on the Logs icon to view errors and debug the configuration.

![Pipeline source status](https://i.ibb.co/ZM45gy1/imageD.png)

When a repository is added as a pipeline source, webhooks are added to the source control repository and each time the configuration changes, your pipeline is immediately updated. This ensures that your pipeline is always in sync with the configuration committed to source control.

### Step 6: View your pipelines

You can now navigate to **Application View | Pipelines | MyPipelines** to see the pipelines that were created - **app_dev_pipeline_su204** and **app_prod_pipeline_su204**.

![View Pipelines](https://i.ibb.co/fk5Gy1K/imageEE.png)

### Step 7:  Execute the dev pipeline

To simulate a real world scenario, you should first execute the **app_dev_pipeline_su204**. This can be done in one of two ways:

* Manual trigger: Click on **app_dev_pipeline_su204** to navigate to the pipeline page. You can click on the first step - **app_build** - and click on the Play button to trigger the pipeline.

![Dev pipeline](https://i.ibb.co/TbGZC05/imageF.png)

* Automatic trigger: You can commit to your forked repository to trigger an automated run. For example, change the **README.md** file in your repository and commit - the pipeline will trigger automatically.

Each time this pipeline executes, it builds the application docker image and pushes it to the **su204-docker-staging-local** repository in Artifactory.  It also publishes a Build for each run.

You can navigate to **Application View | Artifactory | Builds** to see the builds created.

![Build view](https://i.ibb.co/3h1trZN/imageG.png)

You can click on the build and dig in further to find that it points to the docker image that was created by the pipeline.

![Build view](https://i.ibb.co/cXqjsK1/imageh.png)

You can also view the application docker image in the Artifactory tree view by navigating to **Application View | Artifactory | Artifacts**.

![Docker image in tree view](https://i.ibb.co/Kh7byGY/imagei.png)

The BuildInfo resource that is output by the **app_publish** step points to the Build created by the pipeline. This BuildInfo resource, just like any resource) can be used to connect this dev pipeline to a downstream pipeline, such as our production pipeline.

![Docker image in tree view](https://i.ibb.co/vj9T0W5/imagej.png)

### Step 6:  Execute the prod pipeline
Next, you will trigger the production pipeline, which promotes the Build to production repository - **su204-docker-prod-local** - and then “deploys” it. For our lab, deployment just means running the docker container on the build machine.

Navigate to **Application View | Pipelines | My Pipelines** and click on **app_prod_pipeline_su204**.

Note that this pipeline starts with the BuildInfo resource - **app_buildinfo** -  that was updated by the dev pipeline.

Another thing to note about the pipeline is that the line connecting **app_buildinfo** resource to the first step **app_promote_build** is dashed and not solid like other lines in the graph view. This is because in the pipeline configuration, we have configured trigger: false for the input resource. This ensures that the step is not automatically triggered when a new version of the resource is available, which makes sense for a production pipeline since you might not want to deploy each new build.

To execute the production pipeline, click on the first step - **app_promote_build** - and click on the Play button. This will promote your build - i,e, move your docker image to the production repository **su204-docker-prod-local**.  Next, the app_deploy step will download the docker image from the production repository and run the container.

![Execute prod pipeline](https://i.ibb.co/HYmDb6m/imagek.png)

You can click on any step in the pipeline and click on View Logs to switch to the logs view. You can also click on any run in the **Run History** table to view logs for a specific run,  

And that’s it! You have successfully shipped your application!

### Step 7:  Prepare for your Lab Challenge

Create a UI screenshot of the Pipeline of Pipelines view by clicking on **My Pipelines** in the left navbar and clicking on **Graph**. You screenshow should show your server name in the URL similar to the sample provided below.

Once done follow the instructions provided by the instructor to complete the challenge.

![Pipeline of Pipelines view](https://i.ibb.co/z2nHhd4/imageL.png)
