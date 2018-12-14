/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var commonMessage_Eth_pb = require('./commonMessage_Eth_pb.js');
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    id: jspb.Message.getFieldWithDefault(msg, 2, ""),
    event: jspb.Message.getFieldWithDefault(msg, 3, ""),
    hash: jspb.Message.getFieldWithDefault(msg, 4, ""),
    address: jspb.Message.getFieldWithDefault(msg, 5, ""),
    confirmations: jspb.Message.getFieldWithDefault(msg, 6, 0),
    url: jspb.Message.getFieldWithDefault(msg, 7, ""),
    callbackErrors: jspb.Message.getFieldWithDefault(msg, 8, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest;
  return proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (reader.readEnum());
      msg.setNetwork(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setEvent(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setHash(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setConfirmations(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setUrl(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCallbackErrors(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetwork();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEvent();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getHash();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getConfirmations();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
  f = message.getUrl();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getCallbackErrors();
  if (f !== 0) {
    writer.writeInt32(
      8,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string id = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string event = 3;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getEvent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setEvent = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string hash = 4;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getHash = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setHash = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string address = 5;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setAddress = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional int32 confirmations = 6;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getConfirmations = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setConfirmations = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional string url = 7;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setUrl = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional int32 callback_errors = 8;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.getCallbackErrors = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.eth.CreateWebHookEndpointRequest.prototype.setCallbackErrors = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    token: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest;
  return proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (reader.readEnum());
      msg.setNetwork(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetwork();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string token = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.ListWebHooksEndpointRequest.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.repeatedFields_, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.displayName = 'proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.toObject = function(includeInstance, msg) {
  var f, obj = {
    eventList: jspb.Message.toObjectList(msg.getEventList(),
    commonMessage_Eth_pb.Event.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent;
  return proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new commonMessage_Eth_pb.Event;
      reader.readMessage(value,commonMessage_Eth_pb.Event.deserializeBinaryFromReader);
      msg.addEvent(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEventList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      commonMessage_Eth_pb.Event.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Event event = 1;
 * @return {!Array<!proto.fairwaycorp.blockchainprotobuf.eth.Event>}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.prototype.getEventList = function() {
  return /** @type{!Array<!proto.fairwaycorp.blockchainprotobuf.eth.Event>} */ (
    jspb.Message.getRepeatedWrapperField(this, commonMessage_Eth_pb.Event, 1));
};


/** @param {!Array<!proto.fairwaycorp.blockchainprotobuf.eth.Event>} value */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.prototype.setEventList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.Event=} opt_value
 * @param {number=} opt_index
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.Event}
 */
proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.prototype.addEvent = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.fairwaycorp.blockchainprotobuf.eth.Event, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.eth.ArrayEvent.prototype.clearEventList = function() {
  this.setEventList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    webhookid: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest;
  return proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (reader.readEnum());
      msg.setNetwork(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setWebhookid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetwork();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getWebhookid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string webhookid = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.prototype.getWebhookid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.WebHookIDEndpointRequest.prototype.setWebhookid = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    webhookid: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest;
  return proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (reader.readEnum());
      msg.setNetwork(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setWebhookid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetwork();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getWebhookid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string webhookid = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.prototype.getWebhookid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.eth.DeleteWebHookEndpointRequest.prototype.setWebhookid = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.fairwaycorp.blockchainprotobuf.eth);
