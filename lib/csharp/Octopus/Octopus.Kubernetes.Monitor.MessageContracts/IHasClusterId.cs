namespace Octopus.Kubernetes.Monitor.MessageContracts
{
    /// <summary>
    /// Commands and requests should specify ClusterId so we can validate that
    /// 1. The Kubernetes Monitor is asking for information/sending updates for clusters it's associated with
    /// 2. Server is sending information for a cluster to the correct Kubernetes Monitor 
    /// </summary>
    public interface IHasClusterId
    {
        ClusterId ClusterId { get; }
    }
}
