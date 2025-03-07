package euler

const (
	DexType = "euler"
)

const ( // EulerSwap methods
	peripheryMethodGetLimits = "getLimits"
	peripheryMethodQuoteExactInput = "quoteExactInput"
	peripheryMethodQuoteExactOutput = "quoteExactOutput"
)

const (
	factoryMethodAllPools = "allPools"
	factoryMethodAllPoolsLength = "allPoolsLength"
)

const (
	pairMethodAsset0 = "asset0"
	pairMethodAsset1 = "asset1"
	pairMethodGetReserves = "getReserves"

	pairMethodCurve = "curve"
	pairMethodVault0 = "vault0"
	pairMethodVault1 = "vault1"

	pairMethodPriceX = "priceX"
	pairMethodPriceY = "priceY"
	pairMethodConcentrationX = "concentrationX"
	pairMethodConcentrationY = "concentrationY"
	pairMethodEquilibriumReserve0 = "equilibriumReserve0"
	pairMethodEquilibriumReserve1 = "equilibriumReserve1"

	pairMethodEulerAccount = "eulerAccount"
)

var (
	defaultGas = Gas{Swap: 400000}
)
