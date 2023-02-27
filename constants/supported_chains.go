package constants

type ConfigChain struct {
	Name             string
	ContractMultical string
	ChainEndpoint    string
}

func (c ConfigChain) HasContractMultical() bool {
	return c.ContractMultical != ""
}

var (
	ConfigChains = map[string]ConfigChain{
		//"aptos":       {"aptos", ""},
		"arbitrum":  {"arbitrum", "0x7a7443f8c577d537f1d8cd4a629d40a3148dd7ee", "https://arb1.arbitrum.io/rpc"},
		"astar":     {"astar", "0x4a473406ec920b07687b25a7ee7630938ff89c27", "https://evm.astar.network"},
		"aurora":    {"aurora", "0x82d3dDaED9895db23BA59aa26a7858A546D8B6Af", "https://mainnet.aurora.dev/"},
		"avalanche": {"avalanche", "0xa00FB557AA68d2e98A830642DBbFA534E8512E5f", "https://rpc.ankr.com/avalanche"},
		//"binancecoin": {"binancecoin", ""},
		//"bitcoin":     {"bitcoin", ""},
		"bsc":       {"bsc", "0x1ee38d535d541c55c9dae27b12edf090c608e6fb", "https://rpc.ankr.com/bsc"},
		"celo":      {"celo", "0x75f59534dd892c1f8a7b172d639fa854d529ada3", "https://forno.celo.org/"},
		"coinex":    {"coinex", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://rpc.coinex.net/"},
		"cronos":    {"cronos", "0x0fA4d452693F2f45D28c4EC4d20b236C4010dA74", "https://evm.cronos.org/"},
		"dogechain": {"dogechain", "0xf38ffe881d0b1de39ef7a661f40841c0a5bd7d3c", "https://rpc.dogechain.dog"},
		//"dogecoin":    {"dogecoin", ""},
		"ethereum":  {"ethereum", "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696", "https://rpc.ankr.com/eth"},
		"fantom":    {"fantom", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://rpcapi.fantom.network/"},
		"fuse":      {"fuse", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://rpc.fuse.io"},
		"gnosis":    {"gnosis", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://rpc.gnosischain.com"},
		"harmony":   {"harmony", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://api.harmony.one"},
		"huobi":     {"huobi", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://http-mainnet.hecochain.com/"},
		"iotex":     {"iotex", "0x282B7f3234Cf13b62715DCFe1afCA3764ccA4244", "https://babel-api.mainnet.iotex.one"},
		"kava":      {"kava", "0x7ED7bBd8C454a1B0D9EdD939c45a81A03c20131C", "https://evm.kava.io/"},
		"kcc":       {"kcc", "0x90eee4d26c13b66088Aa5066A123fc52ED84eC44", "https://rpc-mainnet.kcc.network"},
		"klaytn":    {"klaytn", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://klaytn01.fandom.finance/"},
		"metis":     {"metis", "0xcA11bde05977b3631167028862bE2a173976CA11", "https://andromeda.metis.io/?owner=1088"},
		"moonbeam":  {"moonbeam", "0x6477204e12a7236b9619385ea453f370ad897bb2", "https://rpc.api.moonbeam.network"},
		"moonriver": {"moonriver", "0xc91bf5d84402c7271644d052fd4915ab2834157c", "https://rpc.api.moonriver.moonbeam.network"},
		//"near":        {"near", "", "https://bsc-dataseed2.defibit.io/"},
		"okc":      {"okc", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://exchainrpc.okex.org/"},
		"onus":     {"onus", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://rpc.onuschain.io/"},
		"optimism": {"optimism", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://rpc.ankr.com/optimism"},
		"polygon":  {"polygon", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72", "https://polygon-rpc.com"},
		// "solana":      {"solana", "", "https://api.mainnet-beta.solana.com"},
		// "sui":         {"sui", ""},
		// "zksync":      {"zksync", "", "https://api.zksync.io/jsrpc"},
	}
)

var (
	SupportedChains = map[string]string{
		"bsc":       "bsc",
		"ethereum":  "ethereum",
		"moonriver": "moonriver",
		"avalanche": "avalanche",
		"aurora":    "aurora",
		"moonbeam":  "moonbeam",
		"fantom":    "fantom",
		"arbitrum":  "arbitrum",
		"polygon":   "polygon",
		"gnosis":    "gnosis",
		"optimism":  "optimism",
		"harmony":   "harmony",
		"celo":      "celo",
		"fuse":      "fuse",
		"coinex":    "coinex",
		"iotex":     "iotex",
	}
	InternalKeysMap = map[string]string{
		"binance-smart-chain": "bsc",
		"ethereum":            "ethereum",
		"moonriver":           "moonriver",
		"avalanche":           "avalanche",
		"aurora":              "aurora",
		"moonbeam":            "moonbeam",
		"fantom":              "fantom",
		"arbitrum":            "arbitrum",
		"polygon":             "polygon",
		"gnosis":              "gnosis",
		"optimism":            "optimism",
		"harmony":             "harmony",
		"celo":                "celo",
		"fuse":                "fuse",
		"coinex":              "coinex",
		"iotex":               "iotex",
	}
)
