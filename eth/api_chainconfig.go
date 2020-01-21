package eth

type PrivateChainConfigAPI struct {
	e *Ethereum
}

func NewPrivateChainConfigAPI(e *Ethereum) *PrivateChainConfigAPI {
	return &PrivateChainConfigAPI{e}
}


