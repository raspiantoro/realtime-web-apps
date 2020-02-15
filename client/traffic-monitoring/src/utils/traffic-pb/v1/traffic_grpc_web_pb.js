/**
 * @fileoverview gRPC-Web generated client stub for traffic
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


// @ts-ignore
const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.traffic = require('./traffic_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.traffic.TrafficClient =
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
proto.traffic.TrafficPromiseClient =
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
 *   !proto.traffic.TrafficCountRequest,
 *   !proto.traffic.TrafficCountResponse>}
 */
const methodDescriptor_Traffic_GetTrafficCount = new grpc.web.MethodDescriptor(
  '/traffic.Traffic/GetTrafficCount',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.traffic.TrafficCountRequest,
  proto.traffic.TrafficCountResponse,
  /**
   * @param {!proto.traffic.TrafficCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.traffic.TrafficCountResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.traffic.TrafficCountRequest,
 *   !proto.traffic.TrafficCountResponse>}
 */
// eslint-disable-next-line
const methodInfo_Traffic_GetTrafficCount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.traffic.TrafficCountResponse,
  /**
   * @param {!proto.traffic.TrafficCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.traffic.TrafficCountResponse.deserializeBinary
);


/**
 * @param {!proto.traffic.TrafficCountRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.traffic.TrafficCountResponse>}
 *     The XHR Node Readable Stream
 */
proto.traffic.TrafficClient.prototype.getTrafficCount =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/traffic.Traffic/GetTrafficCount',
      request,
      metadata || {},
      methodDescriptor_Traffic_GetTrafficCount);
};


/**
 * @param {!proto.traffic.TrafficCountRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.traffic.TrafficCountResponse>}
 *     The XHR Node Readable Stream
 */
proto.traffic.TrafficPromiseClient.prototype.getTrafficCount =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/traffic.Traffic/GetTrafficCount',
      request,
      metadata || {},
      methodDescriptor_Traffic_GetTrafficCount);
};


module.exports = proto.traffic;

