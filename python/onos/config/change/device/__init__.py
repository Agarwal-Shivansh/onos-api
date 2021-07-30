# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/config/change/device/types.proto
# plugin: python-betterproto
from dataclasses import dataclass
from datetime import datetime
from typing import List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase


class ValueType(betterproto.Enum):
    """ValueType is the type for a value"""

    EMPTY = 0
    STRING = 1
    INT = 2
    UINT = 3
    BOOL = 4
    DECIMAL = 5
    FLOAT = 6
    BYTES = 7
    LEAFLIST_STRING = 8
    LEAFLIST_INT = 9
    LEAFLIST_UINT = 10
    LEAFLIST_BOOL = 11
    LEAFLIST_DECIMAL = 12
    LEAFLIST_FLOAT = 13
    LEAFLIST_BYTES = 14


@dataclass(eq=False, repr=False)
class DeviceChange(betterproto.Message):
    """DeviceChange is a stored configuration change for a single device"""

    # 'id' is the unique identifier of the change
    id: str = betterproto.string_field(1)
    # 'index' is a monotonically increasing, globally unique index of the change
    # The index is provided by the store, is static and unique for each unique
    # change identifier, and should not be modified by client code.
    index: int = betterproto.uint64_field(2)
    # 'revision' is the change revision number The revision number is provided by
    # the store and should not be modified by client code. Each unique state of
    # the change will be assigned a unique revision number which can be used for
    # optimistic concurrency control when updating or deleting the change state.
    revision: int = betterproto.uint64_field(3)
    # 'network_change' is a reference to the NetworkChange that created this
    # change
    network_change: "NetworkChangeRef" = betterproto.message_field(4)
    # 'change' is the change object
    change: "Change" = betterproto.message_field(5)
    # 'status' is the lifecycle status of the change
    status: "__change__.Status" = betterproto.message_field(6)
    # 'created' is the time at which the change was created
    created: datetime = betterproto.message_field(7)
    # 'updated' is the time at which the change was last updated
    updated: datetime = betterproto.message_field(8)


@dataclass(eq=False, repr=False)
class NetworkChangeRef(betterproto.Message):
    """
    NetworkChangeRef is a back reference to the NetworkChange that created a
    DeviceChange
    """

    # 'id' is the identifier of the network change from which this change was
    # created
    id: str = betterproto.string_field(1)
    # 'index' is the index of the network change from which this change was
    # created
    index: int = betterproto.uint64_field(2)


@dataclass(eq=False, repr=False)
class Change(betterproto.Message):
    """Change represents a configuration change to a single device"""

    # 'device_id' is the identifier of the device to which this change applies
    device_id: str = betterproto.string_field(1)
    # 'device_version' is an optional device version to which to apply this
    # change
    device_version: str = betterproto.string_field(2)
    # 'device_type' is an optional device type to which to apply this change
    device_type: str = betterproto.string_field(3)
    # 'values' is a set of change values to apply
    values: List["ChangeValue"] = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class TypedValue(betterproto.Message):
    """TypedValue is a value represented as a byte array"""

    # 'bytes' is the bytes array
    bytes_: bytes = betterproto.bytes_field(1)
    # 'type' is the value type
    type: "ValueType" = betterproto.enum_field(2)
    # 'type_opts' is a set of type options
    type_opts: List[int] = betterproto.int32_field(3)


@dataclass(eq=False, repr=False)
class ChangeValue(betterproto.Message):
    """
    ChangeValue is an individual Path/Value and removed flag combination in a
    Change
    """

    # 'path' is the path to change
    path: str = betterproto.string_field(1)
    # 'value' is the change value
    value: "TypedValue" = betterproto.message_field(2)
    # 'removed' indicates whether this is a delete
    removed: bool = betterproto.bool_field(3)


@dataclass(eq=False, repr=False)
class PathValue(betterproto.Message):
    """
    PathValue is an individual Path/Value combination - it is like ChangeValue
    above without the removed flag - it is not used in the DeviceChange store
    Instead it is useful for handling OpState and Snapshots where `removed` is
    not relevant
    """

    # 'path' is the path to change
    path: str = betterproto.string_field(1)
    # 'value' is the change value
    value: "TypedValue" = betterproto.message_field(2)


from ... import change as __change__
