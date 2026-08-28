# Kubernetes Monitor Message Contracts
This repository contains the Protobuf contracts and the generated gRPC SDKs that are used for communication between Octopus Server and the Octopus Kubernetes Monitor.

# Guidelines

## General
- This repository is set up with [release please](https://github.com/googleapis/release-please) to handle versioning, changelogs and releasing. To ensure that the changelogs are generated correctly, please use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/#summary) when merging a PR into `main`.

## Protos
- All Protobuf definitions should be placed in the `protos` directory.
- Updates to existing definitions should be backward compatible.
- Follow the [1-1-1 rule](https://protobuf.dev/best-practices/1-1-1/): All proto definitions should have one top-level element and build target per file. A canonical example of this would be the `fetch_container_logs` contracts. 
  - [fetch_container_logs_request.proto](protos/fetch_container_logs_request.proto)
  - [fetch_container_logs_response.proto](protos/fetch_container_logs_response.proto)
  - [fetch_container_logs_service.proto](protos/fetch_container_logs_service.proto)
- Language specific proto properties shouldn't be added directly in the Proto file but rather in the [buf.gen.yaml](buf.gen.yaml) configuration under the key `managed.override`


# Prerequisites 
- [Buf CLI](https://buf.build/docs/cli/installation/)
- [.NET 8.0](https://dotnet.microsoft.com/en-us/download/dotnet/8.0)
- [Go 1.24](https://go.dev/doc/install)

Optionally
- [Buf account](https://buf.build/) - To avoid rate limiting issues when generating SDKs 

# How To
## Add a new Protobuf contract  
1. Create new `*.proto` file(s) within the `protos` directory - Depending on what you are adding you might need multiple proto files to follow the 1-1-1 rule
2. Run `buf format -w` in the root directory of the repository to automatically format the proto files
3. Run `buf lint` to ensure that the proto files meet the linting requirements. (If you don't run this locally it will run as a PR required check)
4. Run `buf generate` to generate the SDKs

## Update an existing Protobuf contract
1. Make changes to your desired Proto file, ensure the changes are backward compatible see: [Protobuf best practices](https://protobuf.dev/best-practices/dos-donts/)
2. Run `buf format -w` in the root directory of the repository to automatically format the proto files
3. Run `buf lint` to ensure that the proto files meet the linting requirements. (If you don't run this locally it will run as a PR required check)
4. Run `buf breaking --against .git#subdir=protos` to detect any breaking changes(If you don't run this locally it will run as a PR required check)
4. Run `buf generate` to generate the SDKs

## Releasing new version 
When a change has been merged into main, the [release Github action](.github/workflows/release.yaml) will kick off and create a release please PR with the title `chore: release main`, after approving and merging this PR a Github release and tag will be created automatically and cause the [build Github action](.github/workflows/build.yaml) to kick off to pack and push the generated SDKs to their associated package repositories.

## Testing changes locally

### Dotnet SDK
To use a local version of the Dotnet SDK you can use a local NuGet package source and add your locally built NuGet package there, then you will be able to install the local NuGet package in your Dotnet project.
```bash
# Add a local NuGet package source
dotnet nuget add source -n Local ~/dev/LocalPackages/

# Pack SDK into NuGet package 
dotnet pack lib/csharp/Octopus/Octopus.Kubernetes.Monitor.MessageContracts -o {PATH_TO_LOCAL_NUGET_PACKAGES}

# Example
dotnet pack lib/csharp/Octopus/Octopus.Kubernetes.Monitor.MessageContracts -o ../LocalPackages/
```

### Go SDK
You can use a local version of the Go SDK by using a replace directive in the Kubernete's monitor Go module file. 

```go
replace github.com/octopusdeploy/kubernetes-monitor-contracts/go => {PATH_TO_LOCAL_REPO}

// Example
replace github.com/octopusdeploy/kubernetes-monitor-contracts/go => ../kubernetes-monitor-contracts/go
```

### `build-local` script
Two scripts have been provided to automate generating and installing local packages. The scripts will do the following by default: 
1. Run `buf generate` to generate the gRPC SDKs from the Proto files - This will generate the SDKs for all of the languages
2. Add a replace directive to the Kubernetes Monitor `go.mod` file assuming that the Kubernetes Project folder is at `../lobster-watcher`
3. Update all the versions of any package references for to `Octopus.Kubernetes.Monitor.MessageContracts` to the local nuget package.

**This script assumes that a local nuget package source has been configured already, the default location is `../LocalPackages`**

#### Usage
```bash
# Bash
./build-local.sh
./build-local.sh --languages=go 

# For more detailed information about available parameters
./build-local.sh --help

# Powershell
./build-local.ps1 
./build-local.ps1 --languages=go 
# For more detailed information about available parameters
get-help .\build-local.ps1 -full  
```

# Configuration 

## Versioning 
Each generated SDK is versioned and released at the same time however the version number differ slightly depending on the language's version number requirements e.g. For version `0.1.2` of the contracts would result in two Github releases/tags
- Dotnet: `csharp: v0.1.2`
- Go: `go/v0.1.2` - Due to how Go modules work the version number is prefixed with `go/` to include the sub directory the module is in from the root of the repository 

For branch builds of the SDKs(Currently only the dotnet SDK is affected by this) the version number will be generated based on this following pattern
```
{CURRENT_VERSION}-{BRANCH_NAME}-{SHORT_COMMIT_SHA}

// Example 
0.1.2-exciting-new-feature-e1f2999
```

## Linting
Buf is configured to use the [BASIC](https://buf.build/docs/lint/rules/#basic) linting rules with some exceptions which can be seen in the [buf.yaml](buf.yaml) file.