/**
 * @fileoverview gRPC-Web generated client stub for fairwaycorp.blockchainprotobuf.btc
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')

var webhooksMessage_pb = require('./webhooksMessage_pb.js')

var commonMessage_pb = require('./commonMessage_pb.js')
const proto = {};
proto.fairwaycorp = {};
proto.fairwaycorp.blockchainprotobuf = {};
proto.fairwaycorp.blockchainprotobuf.btc = require('./webhooksService_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.fairwaycorp.blockchainprotobuf.btc.CreateWebHookEndpointRequest,
 *   !proto.fairwaycorp.blockchainprotobuf.btc.Event>}
 */
const methodInfo_WebHooksService_CreateWebHookEndpoint = new grpc.web.AbstractClientBase.MethodInfo(
  commonMessage_pb.Event,
  /** @param {!proto.fairwaycorp.blockchainprotobuf.btc.CreateWebHookEndpointRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  commonMessage_pb.Event.deserializeBinary
);


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.CreateWebHookEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.fairwaycorp.blockchainprotobuf.btc.Event)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.fairwaycorp.blockchainprotobuf.btc.Event>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServiceClient.prototype.createWebHookEndpoint =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fairwaycorp.blockchainprotobuf.btc.WebHooksService/CreateWebHookEndpoint',
      request,
      metadata,
      methodInfo_WebHooksService_CreateWebHookEndpoint,
      callback);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.CreateWebHookEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.fairwaycorp.blockchainprotobuf.btc.Event>}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServicePromiseClient.prototype.createWebHookEndpoint =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.createWebHookEndpoint(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.fairwaycorp.blockchainprotobuf.btc.ListWebHooksEndpointRequest,
 *   !proto.fairwaycorp.blockchainprotobuf.btc.ArrayEvent>}
 */
const methodInfo_WebHooksService_ListWebHooksEndpoint = new grpc.web.AbstractClientBase.MethodInfo(
  webhooksMessage_pb.ArrayEvent,
  /** @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListWebHooksEndpointRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  webhooksMessage_pb.ArrayEvent.deserializeBinary
);


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListWebHooksEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.fairwaycorp.blockchainprotobuf.btc.ArrayEvent)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.fairwaycorp.blockchainprotobuf.btc.ArrayEvent>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServiceClient.prototype.listWebHooksEndpoint =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fairwaycorp.blockchainprotobuf.btc.WebHooksService/ListWebHooksEndpoint',
      request,
      metadata,
      methodInfo_WebHooksService_ListWebHooksEndpoint,
      callback);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListWebHooksEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.fairwaycorp.blockchainprotobuf.btc.ArrayEvent>}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServicePromiseClient.prototype.listWebHooksEndpoint =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.listWebHooksEndpoint(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.fairwaycorp.blockchainprotobuf.btc.WebHookIDEndpointRequest,
 *   !proto.fairwaycorp.blockchainprotobuf.btc.Event>}
 */
const methodInfo_WebHooksService_WebHookIDEndpoint = new grpc.web.AbstractClientBase.MethodInfo(
  commonMessage_pb.Event,
  /** @param {!proto.fairwaycorp.blockchainprotobuf.btc.WebHookIDEndpointRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  commonMessage_pb.Event.deserializeBinary
);


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.WebHookIDEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.fairwaycorp.blockchainprotobuf.btc.Event)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.fairwaycorp.blockchainprotobuf.btc.Event>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServiceClient.prototype.webHookIDEndpoint =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fairwaycorp.blockchainprotobuf.btc.WebHooksService/WebHookIDEndpoint',
      request,
      metadata,
      methodInfo_WebHooksService_WebHookIDEndpoint,
      callback);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.WebHookIDEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.fairwaycorp.blockchainprotobuf.btc.Event>}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServicePromiseClient.prototype.webHookIDEndpoint =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.webHookIDEndpoint(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.fairwaycorp.blockchainprotobuf.btc.DeleteWebHookEndpointRequest,
 *   !proto.fairwaycorp.blockchainprotobuf.btc.NullValue>}
 */
const methodInfo_WebHooksService_DeleteWebHookEndpoint = new grpc.web.AbstractClientBase.MethodInfo(
  commonMessage_pb.NullValue,
  /** @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteWebHookEndpointRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  commonMessage_pb.NullValue.deserializeBinary
);


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteWebHookEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.fairwaycorp.blockchainprotobuf.btc.NullValue)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.fairwaycorp.blockchainprotobuf.btc.NullValue>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServiceClient.prototype.deleteWebHookEndpoint =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fairwaycorp.blockchainprotobuf.btc.WebHooksService/DeleteWebHookEndpoint',
      request,
      metadata,
      methodInfo_WebHooksService_DeleteWebHookEndpoint,
      callback);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteWebHookEndpointRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.fairwaycorp.blockchainprotobuf.btc.NullValue>}
 *     The XHR Node Readable Stream
 */
proto.fairwaycorp.blockchainprotobuf.btc.WebHooksServicePromiseClient.prototype.deleteWebHookEndpoint =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.deleteWebHookEndpoint(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.fairwaycorp.blockchainprotobuf.btc;

