"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

from lte.protos.policydb_pb2 import RatingGroup

from magma.common.redis.client import get_default_client
from magma.common.redis.containers import RedisHashDict
from magma.common.redis.serializers import get_proto_deserializer, \
    get_proto_serializer


class RatingGroupsDict(RedisHashDict):
    """
    RatingGroupDict uses the RedisHashDict collection to store a mapping of
    rating group ids to RatingGroups. Setting and deleting items in the
    dictionary syncs with Redis automatically
    """
    _DICT_HASH = "policydb:rating_groups"
    _NOTIFY_CHANNEL = "policydb:rating_groups:stream_update"

    def __init__(self):
        client = get_default_client()
        super().__init__(
            client,
            self._DICT_HASH,
            get_proto_serializer(),
            get_proto_deserializer(RatingGroup))

    def send_update_notification(self):
        """
        Use Redis pub/sub channels to send notifications. Subscribers can listen
        to this channel to know when an update is done to the policy store
        """
        self.redis.publish(self._NOTIFY_CHANNEL, "Stream Update")

    def __missing__(self, key):
        """Instead of throwing a key error, return None when key not found"""
        return None
