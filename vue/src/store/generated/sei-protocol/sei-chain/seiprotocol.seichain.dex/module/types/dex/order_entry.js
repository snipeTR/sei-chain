/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "seiprotocol.seichain.dex";
const baseOrderEntry = {
    price: "",
    quantity: "",
    priceDenom: "",
    assetDenom: "",
};
export const OrderEntry = {
    encode(message, writer = Writer.create()) {
        if (message.price !== "") {
            writer.uint32(10).string(message.price);
        }
        if (message.quantity !== "") {
            writer.uint32(18).string(message.quantity);
        }
        for (const v of message.allocations) {
            Allocation.encode(v, writer.uint32(26).fork()).ldelim();
        }
        if (message.priceDenom !== "") {
            writer.uint32(34).string(message.priceDenom);
        }
        if (message.assetDenom !== "") {
            writer.uint32(42).string(message.assetDenom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseOrderEntry };
        message.allocations = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.price = reader.string();
                    break;
                case 2:
                    message.quantity = reader.string();
                    break;
                case 3:
                    message.allocations.push(Allocation.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.priceDenom = reader.string();
                    break;
                case 5:
                    message.assetDenom = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseOrderEntry };
        message.allocations = [];
        if (object.price !== undefined && object.price !== null) {
            message.price = String(object.price);
        }
        else {
            message.price = "";
        }
        if (object.quantity !== undefined && object.quantity !== null) {
            message.quantity = String(object.quantity);
        }
        else {
            message.quantity = "";
        }
        if (object.allocations !== undefined && object.allocations !== null) {
            for (const e of object.allocations) {
                message.allocations.push(Allocation.fromJSON(e));
            }
        }
        if (object.priceDenom !== undefined && object.priceDenom !== null) {
            message.priceDenom = String(object.priceDenom);
        }
        else {
            message.priceDenom = "";
        }
        if (object.assetDenom !== undefined && object.assetDenom !== null) {
            message.assetDenom = String(object.assetDenom);
        }
        else {
            message.assetDenom = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.price !== undefined && (obj.price = message.price);
        message.quantity !== undefined && (obj.quantity = message.quantity);
        if (message.allocations) {
            obj.allocations = message.allocations.map((e) => e ? Allocation.toJSON(e) : undefined);
        }
        else {
            obj.allocations = [];
        }
        message.priceDenom !== undefined && (obj.priceDenom = message.priceDenom);
        message.assetDenom !== undefined && (obj.assetDenom = message.assetDenom);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseOrderEntry };
        message.allocations = [];
        if (object.price !== undefined && object.price !== null) {
            message.price = object.price;
        }
        else {
            message.price = "";
        }
        if (object.quantity !== undefined && object.quantity !== null) {
            message.quantity = object.quantity;
        }
        else {
            message.quantity = "";
        }
        if (object.allocations !== undefined && object.allocations !== null) {
            for (const e of object.allocations) {
                message.allocations.push(Allocation.fromPartial(e));
            }
        }
        if (object.priceDenom !== undefined && object.priceDenom !== null) {
            message.priceDenom = object.priceDenom;
        }
        else {
            message.priceDenom = "";
        }
        if (object.assetDenom !== undefined && object.assetDenom !== null) {
            message.assetDenom = object.assetDenom;
        }
        else {
            message.assetDenom = "";
        }
        return message;
    },
};
const baseAllocation = { orderId: 0, quantity: "", account: "" };
export const Allocation = {
    encode(message, writer = Writer.create()) {
        if (message.orderId !== 0) {
            writer.uint32(8).uint64(message.orderId);
        }
        if (message.quantity !== "") {
            writer.uint32(18).string(message.quantity);
        }
        if (message.account !== "") {
            writer.uint32(26).string(message.account);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseAllocation };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.orderId = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.quantity = reader.string();
                    break;
                case 3:
                    message.account = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseAllocation };
        if (object.orderId !== undefined && object.orderId !== null) {
            message.orderId = Number(object.orderId);
        }
        else {
            message.orderId = 0;
        }
        if (object.quantity !== undefined && object.quantity !== null) {
            message.quantity = String(object.quantity);
        }
        else {
            message.quantity = "";
        }
        if (object.account !== undefined && object.account !== null) {
            message.account = String(object.account);
        }
        else {
            message.account = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.orderId !== undefined && (obj.orderId = message.orderId);
        message.quantity !== undefined && (obj.quantity = message.quantity);
        message.account !== undefined && (obj.account = message.account);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseAllocation };
        if (object.orderId !== undefined && object.orderId !== null) {
            message.orderId = object.orderId;
        }
        else {
            message.orderId = 0;
        }
        if (object.quantity !== undefined && object.quantity !== null) {
            message.quantity = object.quantity;
        }
        else {
            message.quantity = "";
        }
        if (object.account !== undefined && object.account !== null) {
            message.account = object.account;
        }
        else {
            message.account = "";
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
