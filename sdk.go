package main

import (
	"context"
	"tectone/tectone-desktop/internal/sdk"

	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SDK struct {
	ctx context.Context
}

func (s *SDK) GetAllExampleProjectNames(sdkPath string) []string {

	names, err := sdk.GetAllExampleProjectNames(sdkPath)
	if err != nil {
		runtime.LogErrorf(s.ctx, "could not get example projects from %s: %s", sdkPath, err)
		return nil
	}

	return names

}

func (s *SDK) GenerateAddr() string {
	account := crypto.GenerateAccount()

	return account.Address.String()
}

func (s *SDK) register(ctx context.Context) {
	s.ctx = ctx
}

func newSDK() *SDK {
	return &SDK{}
}
