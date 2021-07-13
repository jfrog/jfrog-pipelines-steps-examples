# Overview

This sample shows you how to build a simple Go application, publish the binary to Artifactory, and publish build-info.

This sample uses two Pipelines native steps:

- **GoBuild**: The GoBuild native step performs a build from Go (GoLang) source. For more information, see https://www.jfrog.com/confluence/display/JFROG/GoBuild.
- **GoPublishBinary**: The GoPublishBinary native step publishes the GO (GoLang) binaries built in a GoBuild step to Artifactory. For more information, see https://www.jfrog.com/confluence/display/JFROG/GoPublishBinary.

Detailed instructions on how to run this sample, as well as explanation of the configuration, is in the JFrog Pipelines Quickstart documentation: [Pipelines Example: Build a Go project](https://www.jfrog.com/confluence/display/JFROG/Pipeline+Example%3A+Go+Build).
