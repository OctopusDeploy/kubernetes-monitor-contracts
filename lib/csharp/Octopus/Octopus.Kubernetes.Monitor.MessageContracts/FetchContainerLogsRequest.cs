namespace Octopus.Kubernetes.Monitor.MessageContracts
{
    [OctopusServerInitiated]
    public partial class FetchContainerLogsRequest : IHasClusterId;
}