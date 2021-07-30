# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/e2sub/subscription/subscription.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterator, Dict, List

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


class Status(betterproto.Enum):
    """Status is a subscription status"""

    ACTIVE = 0
    PENDING_DELETE = 1
    FAILED = 2


class EventType(betterproto.Enum):
    """EventType is a subscription event type"""

    NONE = 0
    ADDED = 1
    UPDATED = 2
    REMOVED = 3


class Encoding(betterproto.Enum):
    """Encoding indicates a payload encoding"""

    ENCODING_ASN1 = 0
    ENCODING_PROTO = 1


class ActionType(betterproto.Enum):
    ACTION_TYPE_REPORT = 0
    ACTION_TYPE_INSERT = 1
    ACTION_TYPE_POLICY = 2


class SubsequentActionType(betterproto.Enum):
    SUBSEQUENT_ACTION_TYPE_CONTINUE = 0
    SUBSEQUENT_ACTION_TYPE_WAIT = 1


class TimeToWait(betterproto.Enum):
    TIME_TO_WAIT_ZERO = 0
    TIME_TO_WAIT_W1MS = 1
    TIME_TO_WAIT_W2MS = 2
    TIME_TO_WAIT_W5MS = 3
    TIME_TO_WAIT_W10MS = 4
    TIME_TO_WAIT_W20MS = 5
    TIME_TO_WAIT_W30MS = 6
    TIME_TO_WAIT_W40MS = 7
    TIME_TO_WAIT_W50MS = 8
    TIME_TO_WAIT_W100MS = 9
    TIME_TO_WAIT_W200MS = 10
    TIME_TO_WAIT_W500MS = 11
    TIME_TO_WAIT_W1S = 12
    TIME_TO_WAIT_W2S = 13
    TIME_TO_WAIT_W5S = 14
    TIME_TO_WAIT_W10S = 15
    TIME_TO_WAIT_W20S = 16
    TIME_TO_WAIT_W60S = 17


@dataclass(eq=False, repr=False)
class Lifecycle(betterproto.Message):
    """Lifecycle is the subscription lifecycle"""

    status: "Status" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class Event(betterproto.Message):
    """Event is a subscription event"""

    type: "EventType" = betterproto.enum_field(1)
    subscription: "Subscription" = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class ServiceModel(betterproto.Message):
    """ServiceModel is a service model definition"""

    name: str = betterproto.string_field(1)
    version: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class Subscription(betterproto.Message):
    """Subscription is a subscription state"""

    id: str = betterproto.string_field(1)
    revision: int = betterproto.uint64_field(2)
    app_id: str = betterproto.string_field(3)
    details: "SubscriptionDetails" = betterproto.message_field(4)
    lifecycle: "Lifecycle" = betterproto.message_field(5)


@dataclass(eq=False, repr=False)
class SubscriptionDetails(betterproto.Message):
    e2_node_id: str = betterproto.string_field(1)
    service_model: "ServiceModel" = betterproto.message_field(2)
    event_trigger: "EventTrigger" = betterproto.message_field(3)
    actions: List["Action"] = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class Payload(betterproto.Message):
    encoding: "Encoding" = betterproto.enum_field(1)
    data: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class EventTrigger(betterproto.Message):
    payload: "Payload" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class Action(betterproto.Message):
    id: int = betterproto.int32_field(1)
    type: "ActionType" = betterproto.enum_field(2)
    payload: "Payload" = betterproto.message_field(3)
    subsequent_action: "SubsequentAction" = betterproto.message_field(4)


@dataclass(eq=False, repr=False)
class SubsequentAction(betterproto.Message):
    """sequence from e2ap-v01.00.00.asn1:1132"""

    type: "SubsequentActionType" = betterproto.enum_field(1)
    time_to_wait: "TimeToWait" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class AddSubscriptionRequest(betterproto.Message):
    """AddSubscriptionRequest a subscription request"""

    subscription: "Subscription" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class AddSubscriptionResponse(betterproto.Message):
    """AddSubscriptionResponse a subscription response"""

    subscription: "Subscription" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class RemoveSubscriptionRequest(betterproto.Message):
    """RemoveSubscriptionRequest a subscription delete request"""

    id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class RemoveSubscriptionResponse(betterproto.Message):
    """RemoveSubscriptionResponse a subscription delete response"""

    pass


@dataclass(eq=False, repr=False)
class GetSubscriptionRequest(betterproto.Message):
    id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class GetSubscriptionResponse(betterproto.Message):
    subscription: "Subscription" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class ListSubscriptionsRequest(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class ListSubscriptionsResponse(betterproto.Message):
    subscriptions: List["Subscription"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class WatchSubscriptionsRequest(betterproto.Message):
    noreplay: bool = betterproto.bool_field(1)


@dataclass(eq=False, repr=False)
class WatchSubscriptionsResponse(betterproto.Message):
    event: "Event" = betterproto.message_field(1)


class E2SubscriptionServiceStub(betterproto.ServiceStub):
    async def add_subscription(
        self, *, subscription: "Subscription" = None
    ) -> "AddSubscriptionResponse":

        request = AddSubscriptionRequest()
        if subscription is not None:
            request.subscription = subscription

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/AddSubscription",
            request,
            AddSubscriptionResponse,
        )

    async def remove_subscription(
        self, *, id: str = ""
    ) -> "RemoveSubscriptionResponse":

        request = RemoveSubscriptionRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/RemoveSubscription",
            request,
            RemoveSubscriptionResponse,
        )

    async def get_subscription(self, *, id: str = "") -> "GetSubscriptionResponse":

        request = GetSubscriptionRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/GetSubscription",
            request,
            GetSubscriptionResponse,
        )

    async def list_subscriptions(self) -> "ListSubscriptionsResponse":

        request = ListSubscriptionsRequest()

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/ListSubscriptions",
            request,
            ListSubscriptionsResponse,
        )

    async def watch_subscriptions(
        self, *, noreplay: bool = False
    ) -> AsyncIterator["WatchSubscriptionsResponse"]:

        request = WatchSubscriptionsRequest()
        request.noreplay = noreplay

        async for response in self._unary_stream(
            "/onos.e2sub.subscription.E2SubscriptionService/WatchSubscriptions",
            request,
            WatchSubscriptionsResponse,
        ):
            yield response


class E2SubscriptionServiceBase(ServiceBase):
    async def add_subscription(
        self, subscription: "Subscription"
    ) -> "AddSubscriptionResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def remove_subscription(self, id: str) -> "RemoveSubscriptionResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def get_subscription(self, id: str) -> "GetSubscriptionResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def list_subscriptions(self) -> "ListSubscriptionsResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def watch_subscriptions(
        self, noreplay: bool
    ) -> AsyncIterator["WatchSubscriptionsResponse"]:
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_add_subscription(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "subscription": request.subscription,
        }

        response = await self.add_subscription(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_remove_subscription(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "id": request.id,
        }

        response = await self.remove_subscription(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_get_subscription(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "id": request.id,
        }

        response = await self.get_subscription(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_list_subscriptions(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.list_subscriptions(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_watch_subscriptions(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "noreplay": request.noreplay,
        }

        await self._call_rpc_handler_server_stream(
            self.watch_subscriptions,
            stream,
            request_kwargs,
        )

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/onos.e2sub.subscription.E2SubscriptionService/AddSubscription": grpclib.const.Handler(
                self.__rpc_add_subscription,
                grpclib.const.Cardinality.UNARY_UNARY,
                AddSubscriptionRequest,
                AddSubscriptionResponse,
            ),
            "/onos.e2sub.subscription.E2SubscriptionService/RemoveSubscription": grpclib.const.Handler(
                self.__rpc_remove_subscription,
                grpclib.const.Cardinality.UNARY_UNARY,
                RemoveSubscriptionRequest,
                RemoveSubscriptionResponse,
            ),
            "/onos.e2sub.subscription.E2SubscriptionService/GetSubscription": grpclib.const.Handler(
                self.__rpc_get_subscription,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetSubscriptionRequest,
                GetSubscriptionResponse,
            ),
            "/onos.e2sub.subscription.E2SubscriptionService/ListSubscriptions": grpclib.const.Handler(
                self.__rpc_list_subscriptions,
                grpclib.const.Cardinality.UNARY_UNARY,
                ListSubscriptionsRequest,
                ListSubscriptionsResponse,
            ),
            "/onos.e2sub.subscription.E2SubscriptionService/WatchSubscriptions": grpclib.const.Handler(
                self.__rpc_watch_subscriptions,
                grpclib.const.Cardinality.UNARY_STREAM,
                WatchSubscriptionsRequest,
                WatchSubscriptionsResponse,
            ),
        }
