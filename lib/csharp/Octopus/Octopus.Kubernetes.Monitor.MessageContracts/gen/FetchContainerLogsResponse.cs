// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: fetch_container_logs_response.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Octopus.Kubernetes.Monitor.MessageContracts {

  /// <summary>Holder for reflection information generated from fetch_container_logs_response.proto</summary>
  public static partial class FetchContainerLogsResponseReflection {

    #region Descriptor
    /// <summary>File descriptor for fetch_container_logs_response.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static FetchContainerLogsResponseReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiNmZXRjaF9jb250YWluZXJfbG9nc19yZXNwb25zZS5wcm90bxoLZXJyb3Iu",
            "cHJvdG8aDmxvZ19saW5lLnByb3RvGhBzZXNzaW9uX2lkLnByb3RvIpsBChpG",
            "ZXRjaENvbnRhaW5lckxvZ3NSZXNwb25zZRIpCgpzZXNzaW9uX2lkGAEgASgL",
            "MgouU2Vzc2lvbklkUglzZXNzaW9uSWQSJQoJbG9nX2xpbmVzGAIgAygLMggu",
            "TG9nTGluZVIIbG9nTGluZXMSIQoFZXJyb3IYAyABKAsyBi5FcnJvckgAUgVl",
            "cnJvcogBAUIICgZfZXJyb3JCkwFCH0ZldGNoQ29udGFpbmVyTG9nc1Jlc3Bv",
            "bnNlUHJvdG9QAVpAZ2l0aHViLmNvbS9vY3RvcHVzZGVwbG95L2t1YmVybmV0",
            "ZXMtbW9uaXRvci1jb250cmFjdHMvbGliL2dvL2dlbqoCK09jdG9wdXMuS3Vi",
            "ZXJuZXRlcy5Nb25pdG9yLk1lc3NhZ2VDb250cmFjdHNiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Octopus.Kubernetes.Monitor.MessageContracts.ErrorReflection.Descriptor, global::Octopus.Kubernetes.Monitor.MessageContracts.LogLineReflection.Descriptor, global::Octopus.Kubernetes.Monitor.MessageContracts.SessionIdReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Octopus.Kubernetes.Monitor.MessageContracts.FetchContainerLogsResponse), global::Octopus.Kubernetes.Monitor.MessageContracts.FetchContainerLogsResponse.Parser, new[]{ "SessionId", "LogLines", "Error" }, new[]{ "Error" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class FetchContainerLogsResponse : pb::IMessage<FetchContainerLogsResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<FetchContainerLogsResponse> _parser = new pb::MessageParser<FetchContainerLogsResponse>(() => new FetchContainerLogsResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<FetchContainerLogsResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Octopus.Kubernetes.Monitor.MessageContracts.FetchContainerLogsResponseReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public FetchContainerLogsResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public FetchContainerLogsResponse(FetchContainerLogsResponse other) : this() {
      sessionId_ = other.sessionId_ != null ? other.sessionId_.Clone() : null;
      logLines_ = other.logLines_.Clone();
      error_ = other.error_ != null ? other.error_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public FetchContainerLogsResponse Clone() {
      return new FetchContainerLogsResponse(this);
    }

    /// <summary>Field number for the "session_id" field.</summary>
    public const int SessionIdFieldNumber = 1;
    private global::Octopus.Kubernetes.Monitor.MessageContracts.SessionId sessionId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Octopus.Kubernetes.Monitor.MessageContracts.SessionId SessionId {
      get { return sessionId_; }
      set {
        sessionId_ = value;
      }
    }

    /// <summary>Field number for the "log_lines" field.</summary>
    public const int LogLinesFieldNumber = 2;
    private static readonly pb::FieldCodec<global::Octopus.Kubernetes.Monitor.MessageContracts.LogLine> _repeated_logLines_codec
        = pb::FieldCodec.ForMessage(18, global::Octopus.Kubernetes.Monitor.MessageContracts.LogLine.Parser);
    private readonly pbc::RepeatedField<global::Octopus.Kubernetes.Monitor.MessageContracts.LogLine> logLines_ = new pbc::RepeatedField<global::Octopus.Kubernetes.Monitor.MessageContracts.LogLine>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<global::Octopus.Kubernetes.Monitor.MessageContracts.LogLine> LogLines {
      get { return logLines_; }
    }

    /// <summary>Field number for the "error" field.</summary>
    public const int ErrorFieldNumber = 3;
    private global::Octopus.Kubernetes.Monitor.MessageContracts.Error error_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Octopus.Kubernetes.Monitor.MessageContracts.Error Error {
      get { return error_; }
      set {
        error_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as FetchContainerLogsResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(FetchContainerLogsResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(SessionId, other.SessionId)) return false;
      if(!logLines_.Equals(other.logLines_)) return false;
      if (!object.Equals(Error, other.Error)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (sessionId_ != null) hash ^= SessionId.GetHashCode();
      hash ^= logLines_.GetHashCode();
      if (error_ != null) hash ^= Error.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (sessionId_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(SessionId);
      }
      logLines_.WriteTo(output, _repeated_logLines_codec);
      if (error_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(Error);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (sessionId_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(SessionId);
      }
      logLines_.WriteTo(ref output, _repeated_logLines_codec);
      if (error_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(Error);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (sessionId_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(SessionId);
      }
      size += logLines_.CalculateSize(_repeated_logLines_codec);
      if (error_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Error);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(FetchContainerLogsResponse other) {
      if (other == null) {
        return;
      }
      if (other.sessionId_ != null) {
        if (sessionId_ == null) {
          SessionId = new global::Octopus.Kubernetes.Monitor.MessageContracts.SessionId();
        }
        SessionId.MergeFrom(other.SessionId);
      }
      logLines_.Add(other.logLines_);
      if (other.error_ != null) {
        if (error_ == null) {
          Error = new global::Octopus.Kubernetes.Monitor.MessageContracts.Error();
        }
        Error.MergeFrom(other.Error);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
      if ((tag & 7) == 4) {
        // Abort on any end group tag.
        return;
      }
      switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (sessionId_ == null) {
              SessionId = new global::Octopus.Kubernetes.Monitor.MessageContracts.SessionId();
            }
            input.ReadMessage(SessionId);
            break;
          }
          case 18: {
            logLines_.AddEntriesFrom(input, _repeated_logLines_codec);
            break;
          }
          case 26: {
            if (error_ == null) {
              Error = new global::Octopus.Kubernetes.Monitor.MessageContracts.Error();
            }
            input.ReadMessage(Error);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
      if ((tag & 7) == 4) {
        // Abort on any end group tag.
        return;
      }
      switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (sessionId_ == null) {
              SessionId = new global::Octopus.Kubernetes.Monitor.MessageContracts.SessionId();
            }
            input.ReadMessage(SessionId);
            break;
          }
          case 18: {
            logLines_.AddEntriesFrom(ref input, _repeated_logLines_codec);
            break;
          }
          case 26: {
            if (error_ == null) {
              Error = new global::Octopus.Kubernetes.Monitor.MessageContracts.Error();
            }
            input.ReadMessage(Error);
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
