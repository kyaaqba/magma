#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional, Dict
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from ..fragment.customer import CustomerFragment, QUERY as CustomerFragmentQuery

QUERY: List[str] = CustomerFragmentQuery + ["""
query CustomersQuery {
  customers {
    edges {
      node {
        ...CustomerFragment
      }
    }
  }
}

"""]

@dataclass
class CustomersQuery(DataClassJsonMixin):
    @dataclass
    class CustomersQueryData(DataClassJsonMixin):
        @dataclass
        class CustomerConnection(DataClassJsonMixin):
            @dataclass
            class CustomerEdge(DataClassJsonMixin):
                @dataclass
                class Customer(CustomerFragment):
                    pass

                node: Optional[Customer]

            edges: List[CustomerEdge]

        customers: Optional[CustomerConnection]

    data: CustomersQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient) -> Optional[CustomersQueryData.CustomerConnection]:
        # fmt: off
        variables: Dict[str, Any] = {}
        try:
            network_start = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            decode_start = perf_counter()
            res = cls.from_json(response_text).data
            decode_time = perf_counter() - decode_start
            network_time = decode_start - network_start
            client.reporter.log_successful_operation("CustomersQuery", variables, network_time, decode_time)
            return res.customers
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "CustomersQuery",
                variables,
            )
