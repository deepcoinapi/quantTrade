package dc

const (
	MASTER = "master"
	URL    = "https://api.deepcoin.com"

	WS_SWAP_ADDR = "wss://stream.deepcoin.com/public/ws"
	WS_SPOT_ADDR = "wss://stream.deepcoin.com/public/spotws"

	HTTP_METHOD_GET  string = "GET"
	HTTP_METHOD_POST string = "POST"

	SPOT string = "SPOT"
	SWAP string = "SWAP"

	CROSS    string = "cross"
	ISOLATED string = "isolated"
	MERGE    string = "merge"
	SPLIT    string = "split"
	CASH     string = "cash"

	SIDE_BUY  string = "buy"
	SIDE_SELL string = "sell"

	ORDER_TYPE_MARKET    = "market"
	ORDER_TYPE_LIMIT     = "limit"
	ORDER_TYPE_POST_ONLY = "post_only"
	ORDER_TYPE_IOC       = "ioc"

	POSITION_SIDE_LONG  = "long"
	POSITION_SIDE_SHORT = "short"

	ACCOUNT_BALANCE string = "/deepcoin/account/balances"
	ACCOUNT_BILLS   string = "/deepcoin/account/bills"
	SET_LEVERAGE    string = "/deepcoin/account/set-leverage"
	POSITIONS       string = "/deepcoin/account/positions"

	MARKET_CANDLES     string = "/deepcoin/market/candles"
	MARKET_TICKERS     string = "/deepcoin/market/tickers"
	MARKET_INSTRUMENTS string = "/deepcoin/market/instruments"

	TRADE_ORDER              string = "/deepcoin/trade/order"
	TRADE_CANCEL_ORDER       string = "/deepcoin/trade/cancel-order"
	TRADE_FILLS              string = "/deepcoin/trade/fills"
	TRADE_HISTORY_ORDER      string = "/deepcoin/trade/orders-history"
	TRADE_PENDING_ORDER      string = "/deepcoin/trade/orders-pending"
	TRADE_POSITION           string = "/deepcoin/account/positions"
	TRADE_ORDER_BY_ID        string = "/deepcoin/trade/orderByID"
	TRADE_FINISH_ORDER_BY_ID string = "/deepcoin/trade/finishOrderByID"
	TRADE_FUNDING_RATE       string = "/deepcoin/trade/funding-rate"
	TRADE_REPLACE_ORDER      string = "/deepcoin/trade/replace-order"
	TRADE_BATCH_CANCEL_ORDER string = "/deepcoin/trade/batch-cancel-order"
	TRADE_PENDING_ORDER_V2   string = "/deepcoin/trade/v2/orders-pending"
	TRADE_SWAP_CANCEL_ALL    string = "/deepcoin/trade/swap/cancel-all"
	TRADE_REPLACE_ORDER_SLTP string = "/deepcoin/trade/replace-order-sltp"
	TRADE_REPLACE_POS_SLTP   string = "/deepcoin/trade/replace-pos-sltp"

	COPYTRADING_LEADER_SETTINGS  string = "/deepcoin/copytrading/leader-settings"
	COPYTRADING_SUPPORT_CONTRACT string = "/deepcoin/copytrading/support-contracts"
	COPYTRADING_SET_CONTRACT     string = "/deepcoin/copytrading/set-contracts"
	COPYTRADING_LEADER_POSITION  string = "/deepcoin/copytrading/leader-position"
	COPYTRADING_ESTIMATE_PROFIT  string = "/deepcoin/copytrading/estimate-profit"
	COPYTRADING_HISTORY_PROFIT   string = "/deepcoin/copytrading/history-profit"
	COPYTRADING_FOLLOWER_RANK    string = "/deepcoin/copytrading/follower-rank"
	COPYTRADING_GET_ACCOUNTID    string = "/deepcoin/copytrading/get-accountIDs"

	ASSET_DEPOSIT_LIST  = "/deepcoin/asset/deposit-list"
	ASSET_WITHDARW_LIST = "/deepcoin/asset/withdraw-list"

	ListenKey       string = "/deepcoin/listenkey/acquire"
	ExtendListenKey string = "/deepcoin/listenkey/extend"

	INTERNAL_TRANSFER_SUPPORT string = "/deepcoin/internal-transfer/support"
	INTERNAL_TRANSFER         string = "/deepcoin/internal-transfer"
	INTERNAL_TRANSFER_HISTORY string = "/deepcoin/internal-transfer/history-order"
)
