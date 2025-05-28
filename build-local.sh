#!/bin/bash

# Colors for prettier output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default values
LANGUAGES="go,csharp"
OUTPUT_LOCATION="../LocalPackages"
OCTOPUS_DEPLOY_LOCATION="../OctopusDeploy"
K8S_MONITOR_LOCATION="../lobster-watcher"
OCTOPUS_PROJECTS=""

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --languages=*)
      LANGUAGES="${1#*=}"
      shift
      ;;
    --output=*)
      OUTPUT_LOCATION="${1#*=}"
      shift
      ;;
    --octopus=*)
      OCTOPUS_DEPLOY_LOCATION="${1#*=}"
      shift
      ;;
    --k8s-monitor=*)
      K8S_MONITOR_LOCATION="${1#*=}"
      shift
      ;;
    --octopus-projects=*)
      OCTOPUS_PROJECTS="${1#*=}"
      shift
      ;;
    -h|--help)
      echo -e "${CYAN}Usage: $0 [options]${NC}"
      echo "Options:"
      echo "  --languages=LANGS       Comma-separated list of languages (default: go,csharp)"
      echo "  --output=PATH           Output location for built packages (default: ../LocalPackages)"
      echo "  --octopus=PATH          Octopus Deploy project location (default: ../OctopusDeploy)"
      echo "  --k8s-monitor=PATH      Kubernetes Monitor project location (default: ../lobster-watcher)"
      echo "  --octopus-projects=LIST Comma-separated list of specific Octopus projects to update"
      echo "  -h, --help              Show this help message"
      exit 0
      ;;
    *)
      echo -e "${RED}Unknown option: $1${NC}"
      echo "Use -h or --help for usage information"
      exit 1
      ;;
  esac
done

# Convert languages string to array
IFS=',' read -ra LANGUAGE_ARRAY <<< "$LANGUAGES"

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_LOCATION"

# Get absolute path of current directory (repository root)
REPO_PATH=$(pwd)

echo -e "\n${BLUE}=== Generating SDKs ===${NC}"
echo -e "${YELLOW}Running buf generate...${NC}"
buf generate || { echo -e "${RED}❌ Failed to generate SDKs${NC}"; exit 1; }
echo -e "${GREEN}✅ SDKs generated successfully${NC}\n"

# Process each language
for lang in "${LANGUAGE_ARRAY[@]}"; do
  case "$lang" in
    "csharp")
      echo -e "${BLUE}=== Processing C# SDK ===${NC}"

      # Get the current version from csproj
      CSPROJ_PATH="csharp/Octopus/Octopus.Kubernetes.Monitor.MessageContracts/Octopus.Kubernetes.Monitor.MessageContracts.csproj"
      BASE_VERSION=$(grep -o '<Version>[^<]*</Version>' "$CSPROJ_PATH" | sed 's/<Version>\(.*\)<\/Version>/\1/')

      if [ -z "$BASE_VERSION" ]; then
        BASE_VERSION="1.0.0"
      fi

      # Get the current branch name
      BRANCH_NAME=$(git branch --show-current | sed 's/\//-/g')

      # Generate a timestamp
      TIMESTAMP=$(date +"%Y%m%d%H%M%S")

      # Create prerelease version string
      PRERELEASE_VERSION="${BASE_VERSION}-${BRANCH_NAME}-${TIMESTAMP}"

      echo -e "${YELLOW}📦 Packing C# SDK with version: ${PRERELEASE_VERSION}${NC}"

      # Pack with debug configuration and prerelease version
      dotnet pack "$CSPROJ_PATH" -c Debug -p:Version="$PRERELEASE_VERSION" -o "$OUTPUT_LOCATION" || {
        echo -e "${RED}❌ Failed to pack C# SDK${NC}";
        exit 1;
      }
      echo -e "${GREEN}✅ C# SDK packaged to $OUTPUT_LOCATION${NC}\n"

      # Update dependent projects in Octopus Deploy repository
      if [ -d "$OCTOPUS_DEPLOY_LOCATION" ]; then
        DEPENDENT_PROJECTS=""
        
        # Check if specific projects were specified
        if [ -n "$OCTOPUS_PROJECTS" ]; then
          echo -e "${YELLOW}🎯 Updating specified projects...${NC}"
          # Convert projects string to array
          IFS=',' read -ra PROJECT_ARRAY <<< "$OCTOPUS_PROJECTS"
          
          # Find csproj files for specified projects
          for project_name in "${PROJECT_ARRAY[@]}"; do
            PROJECT_PATHS=$(find "$OCTOPUS_DEPLOY_LOCATION" -name "${project_name}.csproj" 2>/dev/null)
            if [ -n "$PROJECT_PATHS" ]; then
              # Add to the list of dependent projects if found
              DEPENDENT_PROJECTS="${DEPENDENT_PROJECTS}${PROJECT_PATHS}"$'\n'
            else
              echo -e "${YELLOW}⚠️ Project not found: $project_name${NC}"
            fi
          done
        else
          # Search for all projects that reference the package
          echo -e "${YELLOW}🔍 Searching for projects referencing Octopus.Kubernetes.Monitor.MessageContracts...${NC}"
          DEPENDENT_PROJECTS=$(find "$OCTOPUS_DEPLOY_LOCATION" -name "*.csproj" -exec grep -l "Octopus.Kubernetes.Monitor.MessageContracts" {} \;)
        fi
        
        # Trim trailing newline
        DEPENDENT_PROJECTS=$(echo "$DEPENDENT_PROJECTS" | sed '/^$/d')

        # Update each project file
        for project in $DEPENDENT_PROJECTS; do
          echo -e "${YELLOW}📝 Updating $project to use version $PRERELEASE_VERSION${NC}"

          # Different sed syntax for macOS vs Linux
          if [[ "$OSTYPE" == "darwin"* ]]; then
            # macOS version
            sed -i '' -E "s|<PackageReference Include=\"Octopus.Kubernetes.Monitor.MessageContracts\" Version=\"[^\"]*\"|<PackageReference Include=\"Octopus.Kubernetes.Monitor.MessageContracts\" Version=\"$PRERELEASE_VERSION\"|" "$project"
          else
            # Linux version
            sed -i -E "s|<PackageReference Include=\"Octopus.Kubernetes.Monitor.MessageContracts\" Version=\"[^\"]*\"|<PackageReference Include=\"Octopus.Kubernetes.Monitor.MessageContracts\" Version=\"$PRERELEASE_VERSION\"|" "$project"
          fi
        done

        if [ -n "$DEPENDENT_PROJECTS" ]; then
          echo -e "${GREEN}✅ Updated $(echo "$DEPENDENT_PROJECTS" | wc -l | tr -d ' ') project(s) to use version $PRERELEASE_VERSION${NC}\n"
        else
          echo -e "${CYAN}ℹ️ No dependent projects found to update${NC}\n"
        fi
      else
        echo -e "${YELLOW}⚠️ Octopus Deploy repository not found at $OCTOPUS_DEPLOY_LOCATION${NC}\n"
      fi
      ;;
    "go")
      echo -e "${BLUE}=== Processing Go SDK ===${NC}"
      echo -e "${YELLOW}🔄 Adding replace directive to Go module...${NC}"
      GO_MOD_FILE="$K8S_MONITOR_LOCATION/go.mod"

      if [ ! -f "$GO_MOD_FILE" ]; then
        echo -e "${RED}❌ go.mod not found at $K8S_MONITOR_LOCATION${NC}"
        exit 1
      fi

      # Correct replace directive format
      REPLACE_DIRECTIVE="replace github.com/OctopusDeploy/kubernetes-monitor-contracts/go => $REPO_PATH/go"

      # Check if the replace directive already exists
      if grep -q "replace github.com/OctopusDeploy/kubernetes-monitor-contracts/go" "$GO_MOD_FILE"; then
        # Update existing directive (macOS and Linux compatible)
        sed -i.bak "s|replace github.com/OctopusDeploy/kubernetes-monitor-contracts/go =>.*|$REPLACE_DIRECTIVE|" "$GO_MOD_FILE" && rm -f "${GO_MOD_FILE}.bak"
      else
        # Add new directive at the end of the file
        echo "" >> "$GO_MOD_FILE"
        echo "$REPLACE_DIRECTIVE" >> "$GO_MOD_FILE"
      fi
      echo -e "${GREEN}✅ Added Go module replace directive to $GO_MOD_FILE${NC}\n"
      ;;
    *)
      echo -e "${YELLOW}⚠️ Unsupported language: $lang${NC}\n"
      ;;
  esac
done

echo -e "${GREEN}✨ Local development setup complete!${NC}"