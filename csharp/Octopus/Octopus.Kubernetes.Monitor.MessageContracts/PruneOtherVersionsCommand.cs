namespace Octopus.Kubernetes.Monitor.MessageContracts
{
    [OctopusServerInitiated]
    public partial class PruneOtherVersionsCommand : IHasClusterId;
}