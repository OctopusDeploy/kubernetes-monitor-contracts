#!/usr/bin/env pwsh

<#
.SYNOPSIS
    Automates local development experience for Kubernetes Monitor Message Contracts.
.DESCRIPTION
    Generates SDKs, builds packages, and updates project references.
.PARAMETER Languages
    Comma-separated list of languages to generate SDKs for (default: go,csharp)
.PARAMETER OutputLocation
    Path where built packages should be moved to (default: ../LocalPackages)
.PARAMETER OctopusLocation
    Path to the Octopus Deploy project (default: ../OctopusDeploy)
.PARAMETER K8sMonitorLocation
    Path to the Kubernetes Monitor project (default: ../lobster-watcher)
.PARAMETER OctopusProjects
    Comma-separated list of specific Octopus projects to update
#>

param (
    [string]$Languages = "go,csharp",
    [string]$OutputLocation = "../LocalPackages",
    [string]$OctopusLocation = "../OctopusDeploy",
    [string]$K8sMonitorLocation = "../lobster-watcher",
    [string]$OctopusProjects = ""
)

# Colors for prettier output
$Green = [System.ConsoleColor]::Green
$Yellow = [System.ConsoleColor]::Yellow
$Red = [System.ConsoleColor]::Red
$Cyan = [System.ConsoleColor]::Cyan
$Blue = [System.ConsoleColor]::Blue
$DefaultColor = [Console]::ForegroundColor

# Create output directory if it doesn't exist
if (-not (Test-Path $OutputLocation)) {
    New-Item -ItemType Directory -Path $OutputLocation | Out-Null
}

# Get absolute path of current directory (repository root)
$RepoPath = (Get-Location).Path

# Split languages into array
$LanguageArray = $Languages.Split(",")

Write-Host "`n=== Generating SDKs ===" -ForegroundColor $Blue
Write-Host "Running buf generate..." -ForegroundColor $Yellow
try {
    buf generate
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Failed to generate SDKs" -ForegroundColor $Red
        exit 1
    }
    Write-Host "✅ SDKs generated successfully" -ForegroundColor $Green
} catch {
    Write-Host "❌ Failed to generate SDKs: $_" -ForegroundColor $Red
    exit 1
}

# Process each language
foreach ($lang in $LanguageArray) {
    switch ($lang) {
        "csharp" {
            Write-Host "`n=== Processing C# SDK ===" -ForegroundColor $Blue

            # Get the current version from csproj
            $CsprojPath = "csharp/Octopus/Octopus.Kubernetes.Monitor.MessageContracts/Octopus.Kubernetes.Monitor.MessageContracts.csproj"
            $CsprojContent = Get-Content $CsprojPath -Raw
            if ($CsprojContent -match '<Version>(.*?)</Version>') {
                $BaseVersion = $Matches[1]
            } else {
                $BaseVersion = "1.0.0"
            }

            # Get the current branch name
            $BranchName = (git branch --show-current) -replace "/", "-"

            # Generate a timestamp
            $Timestamp = Get-Date -Format "yyyyMMddHHmmss"

            # Create prerelease version string
            $PrereleaseVersion = "${BaseVersion}-${BranchName}-${Timestamp}"

            Write-Host "📦 Packing C# SDK with version: $PrereleaseVersion" -ForegroundColor $Yellow

            try {
                dotnet pack $CsprojPath -c Debug -p:Version=$PrereleaseVersion -o $OutputLocation
                if ($LASTEXITCODE -ne 0) {
                    Write-Host "❌ Failed to pack C# SDK" -ForegroundColor $Red
                    exit 1
                }
                Write-Host "✅ C# SDK packaged to $OutputLocation" -ForegroundColor $Green
            } catch {
                Write-Host "❌ Failed to pack C# SDK: $_" -ForegroundColor $Red
                exit 1
            }

            # Update dependent projects in Octopus Deploy repository
            if (Test-Path $OctopusLocation) {
                $DependentProjects = @()

                # Check if specific projects were specified
                if ($OctopusProjects) {
                    Write-Host "🎯 Updating specified projects..." -ForegroundColor $Yellow
                    $ProjectArray = $OctopusProjects.Split(",")

                    # Find csproj files for specified projects
                    foreach ($projectName in $ProjectArray) {
                        $ProjectPaths = Get-ChildItem -Path $OctopusLocation -Recurse -Filter "${projectName}.csproj" -ErrorAction SilentlyContinue
                        if ($ProjectPaths) {
                            $DependentProjects += $ProjectPaths
                        } else {
                            Write-Host "⚠️ Project not found: $projectName" -ForegroundColor $Yellow
                        }
                    }
                } else {
                    # Search for all projects that reference the package
                    Write-Host "🔍 Searching for projects referencing Octopus.Kubernetes.Monitor.MessageContracts..." -ForegroundColor $Yellow
                    $AllProjects = Get-ChildItem -Path $OctopusLocation -Recurse -Filter "*.csproj" -ErrorAction SilentlyContinue
                    foreach ($project in $AllProjects) {
                        $content = Get-Content $project.FullName -Raw
                        if ($content -match 'Octopus\.Kubernetes\.Monitor\.MessageContracts') {
                            $DependentProjects += $project
                        }
                    }
                }

                # Update each project file
                $UpdatedCount = 0
                foreach ($project in $DependentProjects) {
                    Write-Host "📝 Updating $($project.FullName) to use version $PrereleaseVersion" -ForegroundColor $Yellow

                    $content = Get-Content $project.FullName -Raw
                    $updatedContent = $content -replace '<PackageReference\s+Include="Octopus\.Kubernetes\.Monitor\.MessageContracts"\s+Version="[^"]*"',
                    "<PackageReference Include=`"Octopus.Kubernetes.Monitor.MessageContracts`" Version=`"$PrereleaseVersion`""

                    Set-Content -Path $project.FullName -Value $updatedContent
                    $UpdatedCount++
                }

                if ($UpdatedCount -gt 0) {
                    Write-Host "✅ Updated $UpdatedCount project(s) to use version $PrereleaseVersion" -ForegroundColor $Green
                } else {
                    Write-Host "ℹ️ No dependent projects found to update" -ForegroundColor $Cyan
                }
            } else {
                Write-Host "⚠️ Octopus Deploy repository not found at $OctopusLocation" -ForegroundColor $Yellow
            }
        }

        "go" {
            Write-Host "`n=== Processing Go SDK ===" -ForegroundColor $Blue
            Write-Host "🔄 Adding replace directive to Go module..." -ForegroundColor $Yellow
            $GoModFile = Join-Path -Path $K8sMonitorLocation -ChildPath "go.mod"

            if (-not (Test-Path $GoModFile)) {
                Write-Host "❌ go.mod not found at $K8sMonitorLocation" -ForegroundColor $Red
                exit 1
            }

            # Correct replace directive format
            $ReplaceDirective = "replace github.com/OctopusDeploy/kubernetes-monitor-contracts/go => $($RepoPath.Replace('\', '/'))/go"

            $GoModContent = Get-Content $GoModFile -Raw

            # Check if the replace directive already exists
            if ($GoModContent -match 'replace github\.com/OctopusDeploy/kubernetes-monitor-contracts/go =>' ) {
                # Update existing directive
                $GoModContent = $GoModContent -replace 'replace github\.com/OctopusDeploy/kubernetes-monitor-contracts/go =>.*', $ReplaceDirective
            } else {
                # Add new directive at the end of the file
                $GoModContent = $GoModContent + "`n$ReplaceDirective`n"
            }

            Set-Content -Path $GoModFile -Value $GoModContent
            Write-Host "✅ Added Go module replace directive to $GoModFile" -ForegroundColor $Green
        }

        default {
            Write-Host "⚠️ Unsupported language: $lang" -ForegroundColor $Yellow
        }
    }
}

Write-Host "`n✨ Local development setup complete!" -ForegroundColor $Green