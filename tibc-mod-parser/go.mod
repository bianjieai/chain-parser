module github.com/bianjieai/chain-parser/tibc-mod-parser

go 1.16

require (
	github.com/bianjieai/chain-parser/common-parser v0.0.0-20240522021027-ed1aec9da717
	github.com/bianjieai/tibc-go v0.3.1-0.20220412124234-9b75094ff868
	github.com/cosmos/cosmos-sdk v0.45.5-0.20220523154235-2921a1c3c918
	github.com/stretchr/testify v1.7.5
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
)
