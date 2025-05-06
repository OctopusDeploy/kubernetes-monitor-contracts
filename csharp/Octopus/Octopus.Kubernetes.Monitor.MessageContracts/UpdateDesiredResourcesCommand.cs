using System;

namespace Octopus.Kubernetes.Monitor.MessageContracts
{
    [OctopusServerInitiated]
    public partial class UpdateDesiredResourcesCommand : IHasClusterId;
}
