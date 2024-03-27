package model

type Project struct {
	KeyHash     string      `json:"key_hash"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	SDKPath     string      `json:"sdk_path"`
	SandboxPath string      `json:"sandbox_path"`
	Containers  []Container `json:"containers"`
	TestNet     TestNet     `json:"testnet"`
	DevNet      DevNet      `json:"devnet"`
	MainNet     MainNet     `json:"mainnet"`

	CreatedAt int64 `json:"created_at"`
}
