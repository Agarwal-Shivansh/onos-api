# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/mho/mho.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import Dict

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


class MhoParamType(betterproto.Enum):
    """MhoParamType is enumeration type of MHO parameters"""

    ALL = 0
    A3OFFSET = 1
    HYSTERESIS = 2
    TIMETOTRIGGER = 3


@dataclass(eq=False, repr=False)
class GetMhoParamRequest(betterproto.Message):
    # hoParamType is a type of handover parameter
    ho_param_type: "MhoParamType" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class GetMhoParamResponse(betterproto.Message):
    # hoParamType is a type of handover parameter
    ho_param_type: "MhoParamType" = betterproto.enum_field(1)
    # A3-Offset value
    a3_offset: int = betterproto.int32_field(2)
    # Hysteresis value
    hysteresis: int = betterproto.int32_field(3)
    # Time-to-Trigger value
    time_to_trigger: int = betterproto.int32_field(4)


@dataclass(eq=False, repr=False)
class SetMhoParamRequest(betterproto.Message):
    # hoParamType is a type of handover parameter
    ho_param_type: "MhoParamType" = betterproto.enum_field(1)
    # A3-Offset value
    a3_offset: int = betterproto.int32_field(2)
    # Hysteresis value
    hysteresis: int = betterproto.int32_field(3)
    # Time-to-Trigger value
    time_to_trigger: int = betterproto.int32_field(4)


@dataclass(eq=False, repr=False)
class SetMhoParamResponse(betterproto.Message):
    # success is a result whether MHO param is set successfully or not
    success: bool = betterproto.bool_field(1)


class MhoStub(betterproto.ServiceStub):
    async def get_mho_params(
        self, *, ho_param_type: "MhoParamType" = None
    ) -> "GetMhoParamResponse":

        request = GetMhoParamRequest()
        request.ho_param_type = ho_param_type

        return await self._unary_unary(
            "/onos.mho.Mho/GetMhoParams", request, GetMhoParamResponse
        )

    async def set_mho_params(
        self,
        *,
        ho_param_type: "MhoParamType" = None,
        a3_offset: int = 0,
        hysteresis: int = 0,
        time_to_trigger: int = 0,
    ) -> "SetMhoParamResponse":

        request = SetMhoParamRequest()
        request.ho_param_type = ho_param_type
        request.a3_offset = a3_offset
        request.hysteresis = hysteresis
        request.time_to_trigger = time_to_trigger

        return await self._unary_unary(
            "/onos.mho.Mho/SetMhoParams", request, SetMhoParamResponse
        )


class MhoBase(ServiceBase):
    async def get_mho_params(
        self, ho_param_type: "MhoParamType"
    ) -> "GetMhoParamResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def set_mho_params(
        self,
        ho_param_type: "MhoParamType",
        a3_offset: int,
        hysteresis: int,
        time_to_trigger: int,
    ) -> "SetMhoParamResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_get_mho_params(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "ho_param_type": request.ho_param_type,
        }

        response = await self.get_mho_params(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_set_mho_params(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "ho_param_type": request.ho_param_type,
            "a3_offset": request.a3_offset,
            "hysteresis": request.hysteresis,
            "time_to_trigger": request.time_to_trigger,
        }

        response = await self.set_mho_params(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.mho.Mho/GetMhoParams": grpclib.const.Handler(
                self.__rpc_get_mho_params,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetMhoParamRequest,
                GetMhoParamResponse,
            ),
            "/onos.mho.Mho/SetMhoParams": grpclib.const.Handler(
                self.__rpc_set_mho_params,
                grpclib.const.Cardinality.UNARY_UNARY,
                SetMhoParamRequest,
                SetMhoParamResponse,
            ),
        }
