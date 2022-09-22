This example covers three Pipelines native steps:

- **CreateReleaseBundle**: The CreateReleaseBundle is a native step that produces a Release Bundle for distribution to an Artifactory Edge Node. The step can be used to create a signed or unsigned release bundle.
                           For more information, see https://www.jfrog.com/confluence/display/JFROG/CreateReleaseBundle.
- **SignReleaseBundle**: The SignReleaseBundle native step signs a Release Bundle in preparation for distributing it to Edge nodes.
                         For more information about SignReleaseBundle native step, see https://www.jfrog.com/confluence/display/JFROG/SignReleaseBundle.
- **DistributeReleaseBundle**: The DistributeReleaseBundle native step triggers the distribution of a Release Bundle to an Artifactory Edge Node. This step requires a signed release bundle and one or more distribution rules to successfully execute.
                               For more information, see https://www.jfrog.com/confluence/display/JFROG/DistributeReleaseBundle.

