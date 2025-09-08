框架设计
项目整体思路

核心目标：支持多种策略逻辑的扩展和组合。

设计模式：策略模式（Strategy Pattern），便于增加新策略。

模块划分：

数据层：行情数据、订单簿、链上数据。

策略层：趋势跟随、均值回归、套利、做市。

执行层：交易所 API 封装，下单/撤单。

风控层：仓位控制、止损止盈、资金管理。

回测层：历史数据回测，绩效评估。

应用层：实盘交易、回测、模拟交易入口。

quantTrade/
├── cmd/                          # 应用入口
│   ├── backtest/                 # 回测入口
│   │   └── main.go
│   ├── live/                     # 实盘入口
│   │   └── main.go
│   └── simulate/                 # 模拟交易入口
│       └── main.go
│
├── core/                          # 核心业务逻辑
│   ├── data/                     # 数据层
│   │   ├── feed.go               # 数据源接口（行情、K线、订单簿）
│   │   ├── historical.go         # 历史数据获取
│   │   └── realtime.go           # 实时数据订阅
│   │
│   ├── strategy/                 # 策略层
│   │   ├── base.go               # 策略接口定义
│   │   ├── trend/                # 趋势跟随
│   │   │   └── ma_crossover.go   # 均线交叉策略
│   │   ├── mean_reversion/       # 均值回归
│   │   │   └── bollinger.go      # 布林带策略
│   │   ├── arbitrage/            # 套利
│   │   │   ├── triangular.go     # 三角套利
│   │   │   └── future_spot.go    # 期现套利
│   │   └── market_making/        # 做市
│   │       └── basic_mm.go       # 简单做市策略
│   │
│   ├── execution/                # 执行层
│   │   ├── exchange.go           # 交易所接口（DeepCoin, OKX, etc）
│   │   ├── binance.go            # Binance 实现
│   │   └── mock.go               # 模拟交易实现
│   │
│   ├── risk/                     # 风控层
│   │   ├── position.go           # 仓位管理
│   │   ├── stoploss.go           # 止损止盈
│   │   └── riskmanager.go        # 风控接口
│   │
│   └── backtest/                 # 回测层
│       ├── engine.go             # 回测引擎
│       ├── metrics.go            # 绩效指标（Sharpe、最大回撤）
│       └── report.go             # 回测结果报告
│
├── config/                      # 配置文件
│   ├── config.yaml               # API key、参数等
│   └── strategy.yaml             # 策略参数配置
│
├── internal/                     # 内部工具
│   ├── logger/                   # 日志
│   └── utils/                    # 通用函数（时间、数学计算）
│
├── scripts/                      # 数据下载、清洗脚本
│   └── fetch_data.go
│
├── tests/                        # 单元测试
│   └── strategy_test.go
│
├── go.mod
└── README.md
