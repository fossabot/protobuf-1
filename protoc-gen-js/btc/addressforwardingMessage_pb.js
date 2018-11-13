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

var commonMessage_pb = require('./commonMessage_pb.js');
goog.object.extend(proto, commonMessage_pb);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse', null, global);

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
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.repeatedFields_, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.repeatedFields_ = [12];



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
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    id: jspb.Message.getFieldWithDefault(msg, 2, ""),
    token: jspb.Message.getFieldWithDefault(msg, 3, ""),
    destination: jspb.Message.getFieldWithDefault(msg, 4, ""),
    inputAddress: jspb.Message.getFieldWithDefault(msg, 5, ""),
    processFeesAddress: jspb.Message.getFieldWithDefault(msg, 6, ""),
    processFeesSatoshis: jspb.Message.getFieldWithDefault(msg, 7, 0),
    processFeesPercent: +jspb.Message.getFieldWithDefault(msg, 8, 0.0),
    callbackUrl: jspb.Message.getFieldWithDefault(msg, 9, ""),
    enableConfirmations: jspb.Message.getFieldWithDefault(msg, 10, false),
    miningFeesSatoshis: jspb.Message.getFieldWithDefault(msg, 11, 0),
    txsList: jspb.Message.getRepeatedField(msg, 12)
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
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest;
  return proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} */ (reader.readEnum());
      msg.setNetwork(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDestination(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setInputAddress(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setProcessFeesAddress(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setProcessFeesSatoshis(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setProcessFeesPercent(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setCallbackUrl(value);
      break;
    case 10:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setEnableConfirmations(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setMiningFeesSatoshis(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.addTxs(value);
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
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.serializeBinaryToWriter = function(message, writer) {
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
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getDestination();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getInputAddress();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getProcessFeesAddress();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getProcessFeesSatoshis();
  if (f !== 0) {
    writer.writeInt32(
      7,
      f
    );
  }
  f = message.getProcessFeesPercent();
  if (f !== 0.0) {
    writer.writeFloat(
      8,
      f
    );
  }
  f = message.getCallbackUrl();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getEnableConfirmations();
  if (f) {
    writer.writeBool(
      10,
      f
    );
  }
  f = message.getMiningFeesSatoshis();
  if (f !== 0) {
    writer.writeInt32(
      11,
      f
    );
  }
  f = message.getTxsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      12,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string id = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string token = 3;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string destination = 4;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getDestination = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setDestination = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string input_address = 5;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getInputAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setInputAddress = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string process_fees_address = 6;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getProcessFeesAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setProcessFeesAddress = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional int32 process_fees_satoshis = 7;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getProcessFeesSatoshis = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setProcessFeesSatoshis = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional float process_fees_percent = 8;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getProcessFeesPercent = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 8, 0.0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setProcessFeesPercent = function(value) {
  jspb.Message.setProto3FloatField(this, 8, value);
};


/**
 * optional string callback_url = 9;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getCallbackUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setCallbackUrl = function(value) {
  jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional bool enable_confirmations = 10;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getEnableConfirmations = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 10, false));
};


/** @param {boolean} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setEnableConfirmations = function(value) {
  jspb.Message.setProto3BooleanField(this, 10, value);
};


/**
 * optional int32 mining_fees_satoshis = 11;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getMiningFeesSatoshis = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setMiningFeesSatoshis = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * repeated string txs = 12;
 * @return {!Array<string>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.getTxsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 12));
};


/** @param {!Array<string>} value */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.setTxsList = function(value) {
  jspb.Message.setField(this, 12, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 */
proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.addTxs = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 12, value, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest.prototype.clearTxsList = function() {
  this.setTxsList([]);
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
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest';
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
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    start: jspb.Message.getFieldWithDefault(msg, 2, 0)
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
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest;
  return proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} */ (reader.readEnum());
      msg.setNetwork(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setStart(value);
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
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetwork();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getStart();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional int32 start = 2;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.prototype.getStart = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest.prototype.setStart = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
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
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.repeatedFields_, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.displayName = 'proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.repeatedFields_ = [1];



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
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    forwardsList: jspb.Message.toObjectList(msg.getForwardsList(),
    commonMessage_pb.AddressForward.toObject, includeInstance)
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
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse;
  return proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new commonMessage_pb.AddressForward;
      reader.readMessage(value,commonMessage_pb.AddressForward.deserializeBinaryFromReader);
      msg.addForwards(value);
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
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getForwardsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      commonMessage_pb.AddressForward.serializeBinaryToWriter
    );
  }
};


/**
 * repeated AddressForward forwards = 1;
 * @return {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.AddressForward>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.prototype.getForwardsList = function() {
  return /** @type{!Array<!proto.fairwaycorp.blockchainprotobuf.btc.AddressForward>} */ (
    jspb.Message.getRepeatedWrapperField(this, commonMessage_pb.AddressForward, 1));
};


/** @param {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.AddressForward>} value */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.prototype.setForwardsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.AddressForward=} opt_value
 * @param {number=} opt_index
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.AddressForward}
 */
proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.prototype.addForwards = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.fairwaycorp.blockchainprotobuf.btc.AddressForward, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse.prototype.clearForwardsList = function() {
  this.setForwardsList([]);
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
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest';
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
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    payid: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest;
  return proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} */ (reader.readEnum());
      msg.setNetwork(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPayid(value);
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
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetwork();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getPayid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string payid = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.prototype.getPayid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest.prototype.setPayid = function(value) {
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
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.displayName = 'proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse';
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
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.toObject = function(includeInstance, msg) {
  var f, obj = {

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
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse}
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse;
  return proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse}
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
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
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


goog.object.extend(exports, proto.fairwaycorp.blockchainprotobuf.btc);
