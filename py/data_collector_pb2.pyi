from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class HelloRequest(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class HelloResponse(_message.Message):
    __slots__ = ["message"]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class TickData(_message.Message):
    __slots__ = ["instrument", "strike_price", "is_call_option", "ignore", "token", "best_buy_price", "best_buy_quantity", "best_sell_price", "best_sell_quantity", "ltp", "price_high", "price_low", "price_avg_trade", "price_close", "price_open", "price_net_change_percent", "sell_quantity", "buy_quantity", "trade_quantity", "open_interest", "total_trade_value", "last_trade_quantity", "last_trade_time", "net_change"]
    INSTRUMENT_FIELD_NUMBER: _ClassVar[int]
    STRIKE_PRICE_FIELD_NUMBER: _ClassVar[int]
    IS_CALL_OPTION_FIELD_NUMBER: _ClassVar[int]
    IGNORE_FIELD_NUMBER: _ClassVar[int]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    BEST_BUY_PRICE_FIELD_NUMBER: _ClassVar[int]
    BEST_BUY_QUANTITY_FIELD_NUMBER: _ClassVar[int]
    BEST_SELL_PRICE_FIELD_NUMBER: _ClassVar[int]
    BEST_SELL_QUANTITY_FIELD_NUMBER: _ClassVar[int]
    LTP_FIELD_NUMBER: _ClassVar[int]
    PRICE_HIGH_FIELD_NUMBER: _ClassVar[int]
    PRICE_LOW_FIELD_NUMBER: _ClassVar[int]
    PRICE_AVG_TRADE_FIELD_NUMBER: _ClassVar[int]
    PRICE_CLOSE_FIELD_NUMBER: _ClassVar[int]
    PRICE_OPEN_FIELD_NUMBER: _ClassVar[int]
    PRICE_NET_CHANGE_PERCENT_FIELD_NUMBER: _ClassVar[int]
    SELL_QUANTITY_FIELD_NUMBER: _ClassVar[int]
    BUY_QUANTITY_FIELD_NUMBER: _ClassVar[int]
    TRADE_QUANTITY_FIELD_NUMBER: _ClassVar[int]
    OPEN_INTEREST_FIELD_NUMBER: _ClassVar[int]
    TOTAL_TRADE_VALUE_FIELD_NUMBER: _ClassVar[int]
    LAST_TRADE_QUANTITY_FIELD_NUMBER: _ClassVar[int]
    LAST_TRADE_TIME_FIELD_NUMBER: _ClassVar[int]
    NET_CHANGE_FIELD_NUMBER: _ClassVar[int]
    instrument: str
    strike_price: int
    is_call_option: bool
    ignore: str
    token: str
    best_buy_price: float
    best_buy_quantity: int
    best_sell_price: float
    best_sell_quantity: int
    ltp: float
    price_high: float
    price_low: float
    price_avg_trade: float
    price_close: float
    price_open: float
    price_net_change_percent: float
    sell_quantity: int
    buy_quantity: int
    trade_quantity: int
    open_interest: int
    total_trade_value: float
    last_trade_quantity: int
    last_trade_time: str
    net_change: float
    def __init__(self, instrument: _Optional[str] = ..., strike_price: _Optional[int] = ..., is_call_option: bool = ..., ignore: _Optional[str] = ..., token: _Optional[str] = ..., best_buy_price: _Optional[float] = ..., best_buy_quantity: _Optional[int] = ..., best_sell_price: _Optional[float] = ..., best_sell_quantity: _Optional[int] = ..., ltp: _Optional[float] = ..., price_high: _Optional[float] = ..., price_low: _Optional[float] = ..., price_avg_trade: _Optional[float] = ..., price_close: _Optional[float] = ..., price_open: _Optional[float] = ..., price_net_change_percent: _Optional[float] = ..., sell_quantity: _Optional[int] = ..., buy_quantity: _Optional[int] = ..., trade_quantity: _Optional[int] = ..., open_interest: _Optional[int] = ..., total_trade_value: _Optional[float] = ..., last_trade_quantity: _Optional[int] = ..., last_trade_time: _Optional[str] = ..., net_change: _Optional[float] = ...) -> None: ...

class TickResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...
