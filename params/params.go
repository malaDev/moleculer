package params

import (
	. "github.com/moleculer-go/moleculer/common"
)

type Params struct {
}

func (params *Params) GetInt(name string) int {
	return 0
}

func (params *Params) GetString(name string) string {
	return ""
}

func (params *Params) Get(name string) string {
	return params.GetString(name)
}

func ParamsFromContext(ctx *Context) Params {
	return Params{}
}
