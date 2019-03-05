/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./geoaltsvc_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.GeoAltClient =
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
proto.GeoAltPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.GeoAltClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.GeoAltClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.LoginReq,
 *   !proto.LoginResp>}
 */
const methodInfo_GeoAlt_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.LoginResp,
  /** @param {!proto.LoginReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.LoginResp.deserializeBinary
);


/**
 * @param {!proto.LoginReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.LoginResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.LoginResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GeoAltClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GeoAlt/Login',
      request,
      metadata,
      methodInfo_GeoAlt_Login,
      callback);
};


/**
 * @param {!proto.LoginReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.LoginResp>}
 *     The XHR Node Readable Stream
 */
proto.GeoAltPromiseClient.prototype.login =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.login(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.RegisterReq,
 *   !proto.RegisterResp>}
 */
const methodInfo_GeoAlt_Register = new grpc.web.AbstractClientBase.MethodInfo(
  proto.RegisterResp,
  /** @param {!proto.RegisterReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.RegisterResp.deserializeBinary
);


/**
 * @param {!proto.RegisterReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.RegisterResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.RegisterResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GeoAltClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GeoAlt/Register',
      request,
      metadata,
      methodInfo_GeoAlt_Register,
      callback);
};


/**
 * @param {!proto.RegisterReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.RegisterResp>}
 *     The XHR Node Readable Stream
 */
proto.GeoAltPromiseClient.prototype.register =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.register(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.AddAlertReq,
 *   !proto.AddAlertResp>}
 */
const methodInfo_GeoAlt_AddAlert = new grpc.web.AbstractClientBase.MethodInfo(
  proto.AddAlertResp,
  /** @param {!proto.AddAlertReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.AddAlertResp.deserializeBinary
);


/**
 * @param {!proto.AddAlertReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.AddAlertResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.AddAlertResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GeoAltClient.prototype.addAlert =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GeoAlt/AddAlert',
      request,
      metadata,
      methodInfo_GeoAlt_AddAlert,
      callback);
};


/**
 * @param {!proto.AddAlertReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.AddAlertResp>}
 *     The XHR Node Readable Stream
 */
proto.GeoAltPromiseClient.prototype.addAlert =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.addAlert(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.GetAlertsReq,
 *   !proto.GetAlertsResp>}
 */
const methodInfo_GeoAlt_GetAlerts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GetAlertsResp,
  /** @param {!proto.GetAlertsReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.GetAlertsResp.deserializeBinary
);


/**
 * @param {!proto.GetAlertsReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.GetAlertsResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.GetAlertsResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GeoAltClient.prototype.getAlerts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GeoAlt/GetAlerts',
      request,
      metadata,
      methodInfo_GeoAlt_GetAlerts,
      callback);
};


/**
 * @param {!proto.GetAlertsReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.GetAlertsResp>}
 *     The XHR Node Readable Stream
 */
proto.GeoAltPromiseClient.prototype.getAlerts =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getAlerts(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.GetAlertsReq,
 *   !proto.GetAlertsResp>}
 */
const methodInfo_GeoAlt_GetActiveAlerts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GetAlertsResp,
  /** @param {!proto.GetAlertsReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.GetAlertsResp.deserializeBinary
);


/**
 * @param {!proto.GetAlertsReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.GetAlertsResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.GetAlertsResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GeoAltClient.prototype.getActiveAlerts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GeoAlt/GetActiveAlerts',
      request,
      metadata,
      methodInfo_GeoAlt_GetActiveAlerts,
      callback);
};


/**
 * @param {!proto.GetAlertsReq} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.GetAlertsResp>}
 *     The XHR Node Readable Stream
 */
proto.GeoAltPromiseClient.prototype.getActiveAlerts =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getActiveAlerts(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto;

