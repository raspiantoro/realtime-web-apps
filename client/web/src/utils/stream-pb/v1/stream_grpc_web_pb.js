/**
 * @fileoverview gRPC-Web generated client stub for stream
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


// @ts-ignore 
const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.stream = require('./stream_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stream.StreamClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.stream.StreamPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.stream.StreamRequest,
 *   !proto.stream.StreamResponse>}
 */
const methodDescriptor_Stream_GetDataStream = new grpc.web.MethodDescriptor(
  '/stream.Stream/GetDataStream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.stream.StreamRequest,
  proto.stream.StreamResponse,
  /**
   * @param {!proto.stream.StreamRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stream.StreamResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.stream.StreamRequest,
 *   !proto.stream.StreamResponse>}
 */
// eslint-disable-next-line
const methodInfo_Stream_GetDataStream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.stream.StreamResponse,
  /**
   * @param {!proto.stream.StreamRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.stream.StreamResponse.deserializeBinary
);


/**
 * @param {!proto.stream.StreamRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.stream.StreamResponse>}
 *     The XHR Node Readable Stream
 */
proto.stream.StreamClient.prototype.getDataStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stream.Stream/GetDataStream',
      request,
      metadata || {},
      methodDescriptor_Stream_GetDataStream);
};


/**
 * @param {!proto.stream.StreamRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.stream.StreamResponse>}
 *     The XHR Node Readable Stream
 */
proto.stream.StreamPromiseClient.prototype.getDataStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/stream.Stream/GetDataStream',
      request,
      metadata || {},
      methodDescriptor_Stream_GetDataStream);
};


module.exports = proto.stream;

