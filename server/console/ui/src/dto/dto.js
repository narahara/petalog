/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.petalog = (function() {

    /**
     * Namespace petalog.
     * @exports petalog
     * @namespace
     */
    var petalog = {};

    petalog.service = (function() {

        /**
         * Namespace service.
         * @memberof petalog
         * @namespace
         */
        var service = {};

        service.console = (function() {

            /**
             * Namespace console.
             * @memberof petalog.service
             * @namespace
             */
            var console = {};

            console.PingRequest = (function() {

                /**
                 * Properties of a PingRequest.
                 * @memberof petalog.service.console
                 * @interface IPingRequest
                 * @property {string|null} [message] PingRequest message
                 */

                /**
                 * Constructs a new PingRequest.
                 * @memberof petalog.service.console
                 * @classdesc Represents a PingRequest.
                 * @implements IPingRequest
                 * @constructor
                 * @param {petalog.service.console.IPingRequest=} [properties] Properties to set
                 */
                function PingRequest(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * PingRequest message.
                 * @member {string} message
                 * @memberof petalog.service.console.PingRequest
                 * @instance
                 */
                PingRequest.prototype.message = "";

                /**
                 * Creates a new PingRequest instance using the specified properties.
                 * @function create
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {petalog.service.console.IPingRequest=} [properties] Properties to set
                 * @returns {petalog.service.console.PingRequest} PingRequest instance
                 */
                PingRequest.create = function create(properties) {
                    return new PingRequest(properties);
                };

                /**
                 * Encodes the specified PingRequest message. Does not implicitly {@link petalog.service.console.PingRequest.verify|verify} messages.
                 * @function encode
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {petalog.service.console.IPingRequest} message PingRequest message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                PingRequest.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.message != null && Object.hasOwnProperty.call(message, "message"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.message);
                    return writer;
                };

                /**
                 * Encodes the specified PingRequest message, length delimited. Does not implicitly {@link petalog.service.console.PingRequest.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {petalog.service.console.IPingRequest} message PingRequest message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                PingRequest.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a PingRequest message from the specified reader or buffer.
                 * @function decode
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {petalog.service.console.PingRequest} PingRequest
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                PingRequest.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.petalog.service.console.PingRequest();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1: {
                                message.message = reader.string();
                                break;
                            }
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a PingRequest message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {petalog.service.console.PingRequest} PingRequest
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                PingRequest.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a PingRequest message.
                 * @function verify
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                PingRequest.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.message != null && message.hasOwnProperty("message"))
                        if (!$util.isString(message.message))
                            return "message: string expected";
                    return null;
                };

                /**
                 * Creates a PingRequest message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {petalog.service.console.PingRequest} PingRequest
                 */
                PingRequest.fromObject = function fromObject(object) {
                    if (object instanceof $root.petalog.service.console.PingRequest)
                        return object;
                    var message = new $root.petalog.service.console.PingRequest();
                    if (object.message != null)
                        message.message = String(object.message);
                    return message;
                };

                /**
                 * Creates a plain object from a PingRequest message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {petalog.service.console.PingRequest} message PingRequest
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                PingRequest.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    var object = {};
                    if (options.defaults)
                        object.message = "";
                    if (message.message != null && message.hasOwnProperty("message"))
                        object.message = message.message;
                    return object;
                };

                /**
                 * Converts this PingRequest to JSON.
                 * @function toJSON
                 * @memberof petalog.service.console.PingRequest
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                PingRequest.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                /**
                 * Gets the default type url for PingRequest
                 * @function getTypeUrl
                 * @memberof petalog.service.console.PingRequest
                 * @static
                 * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
                 * @returns {string} The default type url
                 */
                PingRequest.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                    if (typeUrlPrefix === undefined) {
                        typeUrlPrefix = "type.googleapis.com";
                    }
                    return typeUrlPrefix + "/petalog.service.console.PingRequest";
                };

                return PingRequest;
            })();

            console.PingResponse = (function() {

                /**
                 * Properties of a PingResponse.
                 * @memberof petalog.service.console
                 * @interface IPingResponse
                 * @property {string|null} [message] PingResponse message
                 */

                /**
                 * Constructs a new PingResponse.
                 * @memberof petalog.service.console
                 * @classdesc Represents a PingResponse.
                 * @implements IPingResponse
                 * @constructor
                 * @param {petalog.service.console.IPingResponse=} [properties] Properties to set
                 */
                function PingResponse(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * PingResponse message.
                 * @member {string} message
                 * @memberof petalog.service.console.PingResponse
                 * @instance
                 */
                PingResponse.prototype.message = "";

                /**
                 * Creates a new PingResponse instance using the specified properties.
                 * @function create
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {petalog.service.console.IPingResponse=} [properties] Properties to set
                 * @returns {petalog.service.console.PingResponse} PingResponse instance
                 */
                PingResponse.create = function create(properties) {
                    return new PingResponse(properties);
                };

                /**
                 * Encodes the specified PingResponse message. Does not implicitly {@link petalog.service.console.PingResponse.verify|verify} messages.
                 * @function encode
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {petalog.service.console.IPingResponse} message PingResponse message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                PingResponse.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.message != null && Object.hasOwnProperty.call(message, "message"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.message);
                    return writer;
                };

                /**
                 * Encodes the specified PingResponse message, length delimited. Does not implicitly {@link petalog.service.console.PingResponse.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {petalog.service.console.IPingResponse} message PingResponse message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                PingResponse.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a PingResponse message from the specified reader or buffer.
                 * @function decode
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {petalog.service.console.PingResponse} PingResponse
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                PingResponse.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.petalog.service.console.PingResponse();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1: {
                                message.message = reader.string();
                                break;
                            }
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a PingResponse message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {petalog.service.console.PingResponse} PingResponse
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                PingResponse.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a PingResponse message.
                 * @function verify
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                PingResponse.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.message != null && message.hasOwnProperty("message"))
                        if (!$util.isString(message.message))
                            return "message: string expected";
                    return null;
                };

                /**
                 * Creates a PingResponse message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {petalog.service.console.PingResponse} PingResponse
                 */
                PingResponse.fromObject = function fromObject(object) {
                    if (object instanceof $root.petalog.service.console.PingResponse)
                        return object;
                    var message = new $root.petalog.service.console.PingResponse();
                    if (object.message != null)
                        message.message = String(object.message);
                    return message;
                };

                /**
                 * Creates a plain object from a PingResponse message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {petalog.service.console.PingResponse} message PingResponse
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                PingResponse.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    var object = {};
                    if (options.defaults)
                        object.message = "";
                    if (message.message != null && message.hasOwnProperty("message"))
                        object.message = message.message;
                    return object;
                };

                /**
                 * Converts this PingResponse to JSON.
                 * @function toJSON
                 * @memberof petalog.service.console.PingResponse
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                PingResponse.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                /**
                 * Gets the default type url for PingResponse
                 * @function getTypeUrl
                 * @memberof petalog.service.console.PingResponse
                 * @static
                 * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
                 * @returns {string} The default type url
                 */
                PingResponse.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
                    if (typeUrlPrefix === undefined) {
                        typeUrlPrefix = "type.googleapis.com";
                    }
                    return typeUrlPrefix + "/petalog.service.console.PingResponse";
                };

                return PingResponse;
            })();

            return console;
        })();

        return service;
    })();

    return petalog;
})();

module.exports = $root;
