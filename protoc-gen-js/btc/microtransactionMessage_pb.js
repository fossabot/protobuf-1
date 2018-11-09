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
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.btc.MicroTX', null, global);
goog.exportSymbol('proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest', null, global);

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
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.repeatedFields_, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.btc.MicroTX, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.displayName = 'proto.fairwaycorp.blockchainprotobuf.btc.MicroTX';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.repeatedFields_ = [9,10,11,12];



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
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTX} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.toObject = function(includeInstance, msg) {
  var f, obj = {
    fromPubkey: jspb.Message.getFieldWithDefault(msg, 1, ""),
    fromPrivate: jspb.Message.getFieldWithDefault(msg, 2, ""),
    fromWif: jspb.Message.getFieldWithDefault(msg, 3, ""),
    toAddress: jspb.Message.getFieldWithDefault(msg, 4, ""),
    valueSatoshis: jspb.Message.getFieldWithDefault(msg, 5, 0),
    token: jspb.Message.getFieldWithDefault(msg, 6, ""),
    changeAddress: jspb.Message.getFieldWithDefault(msg, 7, ""),
    waitGuarantee: jspb.Message.getFieldWithDefault(msg, 8, false),
    tosignList: jspb.Message.getRepeatedField(msg, 9),
    signaturesList: jspb.Message.getRepeatedField(msg, 10),
    inputsList: jspb.Message.toObjectList(msg.getInputsList(),
    commonMessage_pb.TXInput.toObject, includeInstance),
    outputsList: jspb.Message.toObjectList(msg.getOutputsList(),
    commonMessage_pb.TXOutput.toObject, includeInstance),
    fees: jspb.Message.getFieldWithDefault(msg, 13, 0),
    hash: jspb.Message.getFieldWithDefault(msg, 14, "")
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
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTX}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.btc.MicroTX;
  return proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTX} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTX}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFromPubkey(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFromPrivate(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setFromWif(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setToAddress(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setValueSatoshis(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setChangeAddress(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setWaitGuarantee(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.addTosign(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.addSignatures(value);
      break;
    case 11:
      var value = new commonMessage_pb.TXInput;
      reader.readMessage(value,commonMessage_pb.TXInput.deserializeBinaryFromReader);
      msg.addInputs(value);
      break;
    case 12:
      var value = new commonMessage_pb.TXOutput;
      reader.readMessage(value,commonMessage_pb.TXOutput.deserializeBinaryFromReader);
      msg.addOutputs(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFees(value);
      break;
    case 14:
      var value = /** @type {string} */ (reader.readString());
      msg.setHash(value);
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
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTX} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFromPubkey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFromPrivate();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getFromWif();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getToAddress();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getValueSatoshis();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getChangeAddress();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getWaitGuarantee();
  if (f) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getTosignList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      9,
      f
    );
  }
  f = message.getSignaturesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      10,
      f
    );
  }
  f = message.getInputsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      11,
      f,
      commonMessage_pb.TXInput.serializeBinaryToWriter
    );
  }
  f = message.getOutputsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      12,
      f,
      commonMessage_pb.TXOutput.serializeBinaryToWriter
    );
  }
  f = message.getFees();
  if (f !== 0) {
    writer.writeInt32(
      13,
      f
    );
  }
  f = message.getHash();
  if (f.length > 0) {
    writer.writeString(
      14,
      f
    );
  }
};


/**
 * optional string from_pubkey = 1;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getFromPubkey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setFromPubkey = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string from_private = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getFromPrivate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setFromPrivate = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string from_wif = 3;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getFromWif = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setFromWif = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string to_address = 4;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getToAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setToAddress = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int64 value_satoshis = 5;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getValueSatoshis = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setValueSatoshis = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string token = 6;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string change_address = 7;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getChangeAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setChangeAddress = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional bool wait_guarantee = 8;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getWaitGuarantee = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 8, false));
};


/** @param {boolean} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setWaitGuarantee = function(value) {
  jspb.Message.setProto3BooleanField(this, 8, value);
};


/**
 * repeated string tosign = 9;
 * @return {!Array<string>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getTosignList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 9));
};


/** @param {!Array<string>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setTosignList = function(value) {
  jspb.Message.setField(this, 9, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.addTosign = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 9, value, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.clearTosignList = function() {
  this.setTosignList([]);
};


/**
 * repeated string signatures = 10;
 * @return {!Array<string>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getSignaturesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 10));
};


/** @param {!Array<string>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setSignaturesList = function(value) {
  jspb.Message.setField(this, 10, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.addSignatures = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 10, value, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.clearSignaturesList = function() {
  this.setSignaturesList([]);
};


/**
 * repeated TXInput inputs = 11;
 * @return {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXInput>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getInputsList = function() {
  return /** @type{!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXInput>} */ (
    jspb.Message.getRepeatedWrapperField(this, commonMessage_pb.TXInput, 11));
};


/** @param {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXInput>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setInputsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 11, value);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.TXInput=} opt_value
 * @param {number=} opt_index
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.TXInput}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.addInputs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 11, opt_value, proto.fairwaycorp.blockchainprotobuf.btc.TXInput, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.clearInputsList = function() {
  this.setInputsList([]);
};


/**
 * repeated TXOutput outputs = 12;
 * @return {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getOutputsList = function() {
  return /** @type{!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput>} */ (
    jspb.Message.getRepeatedWrapperField(this, commonMessage_pb.TXOutput, 12));
};


/** @param {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setOutputsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 12, value);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput=} opt_value
 * @param {number=} opt_index
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.addOutputs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 12, opt_value, proto.fairwaycorp.blockchainprotobuf.btc.TXOutput, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.clearOutputsList = function() {
  this.setOutputsList([]);
};


/**
 * optional int32 fees = 13;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getFees = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setFees = function(value) {
  jspb.Message.setProto3IntField(this, 13, value);
};


/**
 * optional string hash = 14;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.getHash = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 14, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTX.prototype.setHash = function(value) {
  jspb.Message.setProto3StringField(this, 14, value);
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
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.repeatedFields_, null);
};
goog.inherits(proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.displayName = 'proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.repeatedFields_ = [10,11,12,13];



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
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    network: jspb.Message.getFieldWithDefault(msg, 1, 0),
    fromPubkey: jspb.Message.getFieldWithDefault(msg, 2, ""),
    fromPrivate: jspb.Message.getFieldWithDefault(msg, 3, ""),
    fromWif: jspb.Message.getFieldWithDefault(msg, 4, ""),
    toAddress: jspb.Message.getFieldWithDefault(msg, 5, ""),
    valueSatoshis: jspb.Message.getFieldWithDefault(msg, 6, 0),
    token: jspb.Message.getFieldWithDefault(msg, 7, ""),
    changeAddress: jspb.Message.getFieldWithDefault(msg, 8, ""),
    waitGuarantee: jspb.Message.getFieldWithDefault(msg, 9, false),
    tosignList: jspb.Message.getRepeatedField(msg, 10),
    signaturesList: jspb.Message.getRepeatedField(msg, 11),
    inputsList: jspb.Message.toObjectList(msg.getInputsList(),
    commonMessage_pb.TXInput.toObject, includeInstance),
    outputsList: jspb.Message.toObjectList(msg.getOutputsList(),
    commonMessage_pb.TXOutput.toObject, includeInstance),
    fees: jspb.Message.getFieldWithDefault(msg, 14, 0),
    hash: jspb.Message.getFieldWithDefault(msg, 15, "")
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
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest;
  return proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setFromPubkey(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setFromPrivate(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setFromWif(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setToAddress(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setValueSatoshis(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setChangeAddress(value);
      break;
    case 9:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setWaitGuarantee(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.addTosign(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.addSignatures(value);
      break;
    case 12:
      var value = new commonMessage_pb.TXInput;
      reader.readMessage(value,commonMessage_pb.TXInput.deserializeBinaryFromReader);
      msg.addInputs(value);
      break;
    case 13:
      var value = new commonMessage_pb.TXOutput;
      reader.readMessage(value,commonMessage_pb.TXOutput.deserializeBinaryFromReader);
      msg.addOutputs(value);
      break;
    case 14:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFees(value);
      break;
    case 15:
      var value = /** @type {string} */ (reader.readString());
      msg.setHash(value);
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
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetwork();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getFromPubkey();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getFromPrivate();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getFromWif();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getToAddress();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getValueSatoshis();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getChangeAddress();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getWaitGuarantee();
  if (f) {
    writer.writeBool(
      9,
      f
    );
  }
  f = message.getTosignList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      10,
      f
    );
  }
  f = message.getSignaturesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      11,
      f
    );
  }
  f = message.getInputsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      12,
      f,
      commonMessage_pb.TXInput.serializeBinaryToWriter
    );
  }
  f = message.getOutputsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      13,
      f,
      commonMessage_pb.TXOutput.serializeBinaryToWriter
    );
  }
  f = message.getFees();
  if (f !== 0) {
    writer.writeInt32(
      14,
      f
    );
  }
  f = message.getHash();
  if (f.length > 0) {
    writer.writeString(
      15,
      f
    );
  }
};


/**
 * optional NetworkAllowingAlias network = 1;
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getNetwork = function() {
  return /** @type {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setNetwork = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string from_pubkey = 2;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getFromPubkey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setFromPubkey = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string from_private = 3;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getFromPrivate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setFromPrivate = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string from_wif = 4;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getFromWif = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setFromWif = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string to_address = 5;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getToAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setToAddress = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional int64 value_satoshis = 6;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getValueSatoshis = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setValueSatoshis = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional string token = 7;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string change_address = 8;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getChangeAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setChangeAddress = function(value) {
  jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional bool wait_guarantee = 9;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getWaitGuarantee = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 9, false));
};


/** @param {boolean} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setWaitGuarantee = function(value) {
  jspb.Message.setProto3BooleanField(this, 9, value);
};


/**
 * repeated string tosign = 10;
 * @return {!Array<string>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getTosignList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 10));
};


/** @param {!Array<string>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setTosignList = function(value) {
  jspb.Message.setField(this, 10, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.addTosign = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 10, value, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.clearTosignList = function() {
  this.setTosignList([]);
};


/**
 * repeated string signatures = 11;
 * @return {!Array<string>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getSignaturesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 11));
};


/** @param {!Array<string>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setSignaturesList = function(value) {
  jspb.Message.setField(this, 11, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.addSignatures = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 11, value, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.clearSignaturesList = function() {
  this.setSignaturesList([]);
};


/**
 * repeated TXInput inputs = 12;
 * @return {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXInput>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getInputsList = function() {
  return /** @type{!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXInput>} */ (
    jspb.Message.getRepeatedWrapperField(this, commonMessage_pb.TXInput, 12));
};


/** @param {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXInput>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setInputsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 12, value);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.TXInput=} opt_value
 * @param {number=} opt_index
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.TXInput}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.addInputs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 12, opt_value, proto.fairwaycorp.blockchainprotobuf.btc.TXInput, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.clearInputsList = function() {
  this.setInputsList([]);
};


/**
 * repeated TXOutput outputs = 13;
 * @return {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput>}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getOutputsList = function() {
  return /** @type{!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput>} */ (
    jspb.Message.getRepeatedWrapperField(this, commonMessage_pb.TXOutput, 13));
};


/** @param {!Array<!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput>} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setOutputsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 13, value);
};


/**
 * @param {!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput=} opt_value
 * @param {number=} opt_index
 * @return {!proto.fairwaycorp.blockchainprotobuf.btc.TXOutput}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.addOutputs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 13, opt_value, proto.fairwaycorp.blockchainprotobuf.btc.TXOutput, opt_index);
};


proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.clearOutputsList = function() {
  this.setOutputsList([]);
};


/**
 * optional int32 fees = 14;
 * @return {number}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getFees = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/** @param {number} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setFees = function(value) {
  jspb.Message.setProto3IntField(this, 14, value);
};


/**
 * optional string hash = 15;
 * @return {string}
 */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.getHash = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 15, ""));
};


/** @param {string} value */
proto.fairwaycorp.blockchainprotobuf.btc.MicroTXRequest.prototype.setHash = function(value) {
  jspb.Message.setProto3StringField(this, 15, value);
};


goog.object.extend(exports, proto.fairwaycorp.blockchainprotobuf.btc);
