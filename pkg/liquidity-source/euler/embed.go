package euler

import _ "embed"

//go:embed abis/EulerSwap.json
var eulerSwapJSON []byte

//go:embed abis/EulerSwapFactory.json
var eulerSwapFactoryJSON []byte

//go:embed abis/EulerSwapPeriphery.json
var eulerSwapPeripheryJSON []byte
