package constants

type ConfigChain struct {
	Name             string
	ContractMultical string
}

func (c ConfigChain) HasContractMultical() bool {
	return c.ContractMultical != ""
}

var (
	ConfigChains = map[string]ConfigChain{
		"aptos":       {"aptos", ""},
		"arbitrum":    {"arbitrum", "0x7a7443f8c577d537f1d8cd4a629d40a3148dd7ee"},
		"astar":       {"astar", "0x4a473406ec920b07687b25a7ee7630938ff89c27"},
		"aurora":      {"aurora", "0x82d3dDaED9895db23BA59aa26a7858A546D8B6Af"},
		"avalanche":   {"avalanche", "0xa00FB557AA68d2e98A830642DBbFA534E8512E5f"},
		"binancecoin": {"binancecoin", ""},
		"bitcoin":     {"bitcoin", ""},
		"bsc":         {"bsc", "0x1ee38d535d541c55c9dae27b12edf090c608e6fb"},
		"celo":        {"celo", "0x75f59534dd892c1f8a7b172d639fa854d529ada3"},
		"coinex":      {"coinex", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"cronos":      {"cronos", "0x0fA4d452693F2f45D28c4EC4d20b236C4010dA74"},
		"dogechain":   {"dogechain", "0xf38ffe881d0b1de39ef7a661f40841c0a5bd7d3c"},
		"dogecoin":    {"dogecoin", ""},
		"ethereum":    {"ethereum", "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696"},
		"fantom":      {"fantom", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"fuse":        {"fuse", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"gnosis":      {"gnosis", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"harmony":     {"harmony", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"huobi":       {"huobi", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"iotex":       {"iotex", "0x282B7f3234Cf13b62715DCFe1afCA3764ccA4244"},
		"kava":        {"kava", "0x7ED7bBd8C454a1B0D9EdD939c45a81A03c20131C"},
		"kcc":         {"kcc", "0x90eee4d26c13b66088Aa5066A123fc52ED84eC44"},
		"klaytn":      {"klaytn", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"metis":       {"metis", "0xcA11bde05977b3631167028862bE2a173976CA11"},
		"moonbeam":    {"moonbeam", "0x6477204e12a7236b9619385ea453f370ad897bb2"},
		"moonriver":   {"moonriver", "0xc91bf5d84402c7271644d052fd4915ab2834157c"},
		"near":        {"near", ""},
		"okc":         {"okc", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"onus":        {"onus", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"optimism":    {"optimism", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"polygon":     {"polygon", "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72"},
		"solana":      {"solana", ""},
		"sui":         {"sui", ""},
		"zksync":      {"zksync", ""},
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
