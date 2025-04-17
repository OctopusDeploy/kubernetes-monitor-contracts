// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: live_status_service.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Octopus.Kubernetes.Monitor.MessageContracts {
  /// <summary>
  /// Service to update the status of Kubernetes resources
  /// </summary>
  public static partial class LiveStatusService
  {
    static readonly string __ServiceName = "LiveStatusService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest> __Marshaller_UpdateMonitoredResourcesRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse> __Marshaller_UpdateMonitoredResourcesResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest> __Marshaller_ReplaceMonitoredResourcesRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse> __Marshaller_ReplaceMonitoredResourcesResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest> __Marshaller_DeleteChildMonitoredResourcesRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse> __Marshaller_DeleteChildMonitoredResourcesResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse> __Method_UpdateMonitoredResources = new grpc::Method<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UpdateMonitoredResources",
        __Marshaller_UpdateMonitoredResourcesRequest,
        __Marshaller_UpdateMonitoredResourcesResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse> __Method_ReplaceMonitoredResources = new grpc::Method<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "ReplaceMonitoredResources",
        __Marshaller_ReplaceMonitoredResourcesRequest,
        __Marshaller_ReplaceMonitoredResourcesResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse> __Method_DeleteChildMonitoredResources = new grpc::Method<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "DeleteChildMonitoredResources",
        __Marshaller_DeleteChildMonitoredResourcesRequest,
        __Marshaller_DeleteChildMonitoredResourcesResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Octopus.Kubernetes.Monitor.MessageContracts.LiveStatusServiceReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of LiveStatusService</summary>
    [grpc::BindServiceMethod(typeof(LiveStatusService), "BindService")]
    public abstract partial class LiveStatusServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse> UpdateMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse> ReplaceMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse> DeleteChildMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for LiveStatusService</summary>
    public partial class LiveStatusServiceClient : grpc::ClientBase<LiveStatusServiceClient>
    {
      /// <summary>Creates a new client for LiveStatusService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public LiveStatusServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for LiveStatusService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public LiveStatusServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected LiveStatusServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected LiveStatusServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse UpdateMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateMonitoredResources(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse UpdateMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UpdateMonitoredResources, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse> UpdateMonitoredResourcesAsync(global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UpdateMonitoredResourcesAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse> UpdateMonitoredResourcesAsync(global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UpdateMonitoredResources, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse ReplaceMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ReplaceMonitoredResources(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse ReplaceMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_ReplaceMonitoredResources, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse> ReplaceMonitoredResourcesAsync(global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return ReplaceMonitoredResourcesAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse> ReplaceMonitoredResourcesAsync(global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_ReplaceMonitoredResources, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse DeleteChildMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return DeleteChildMonitoredResources(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse DeleteChildMonitoredResources(global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_DeleteChildMonitoredResources, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse> DeleteChildMonitoredResourcesAsync(global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return DeleteChildMonitoredResourcesAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse> DeleteChildMonitoredResourcesAsync(global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_DeleteChildMonitoredResources, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override LiveStatusServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new LiveStatusServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(LiveStatusServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_UpdateMonitoredResources, serviceImpl.UpdateMonitoredResources)
          .AddMethod(__Method_ReplaceMonitoredResources, serviceImpl.ReplaceMonitoredResources)
          .AddMethod(__Method_DeleteChildMonitoredResources, serviceImpl.DeleteChildMonitoredResources).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, LiveStatusServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_UpdateMonitoredResources, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.UpdateMonitoredResourcesResponse>(serviceImpl.UpdateMonitoredResources));
      serviceBinder.AddMethod(__Method_ReplaceMonitoredResources, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.ReplaceMonitoredResourcesResponse>(serviceImpl.ReplaceMonitoredResources));
      serviceBinder.AddMethod(__Method_DeleteChildMonitoredResources, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesRequest, global::Octopus.Kubernetes.Monitor.MessageContracts.DeleteChildMonitoredResourcesResponse>(serviceImpl.DeleteChildMonitoredResources));
    }

  }
}
#endregion
