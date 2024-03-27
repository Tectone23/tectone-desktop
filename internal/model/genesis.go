package model

type TestNet struct {
	Genesis Genesis `json:"genesis"`
}
type DevNet struct {
	Genesis Genesis `json:"genesis"`
}
type MainNet struct {
	Genesis Genesis `json:"genesis"`
}
type Genesis struct {
	Alloc     []Allocation `json:"alloc"`
	Fees      string       `json:"fees"`
	ID        string       `json:"id"`
	Network   string       `json:"network"`
	Proto     string       `json:"proto,omitempty"`
	Rwd       string       `json:"rwd"`
	Timestamp int64        `json:"timestamp,omitempty"`
}
type Allocation struct {
	Addr    string `json:"addr"`
	Comment string `json:"comment"`
	State   State  `json:"state"`
}

type State struct {
	Algo    int64  `json:"algo"`
	Onl     int64  `json:"onl,omitempty"`
	Sel     string `json:"sel,omitempty"`
	Vote    string `json:"vote,omitempty"`
	VoteKD  int64  `json:"voteKD,omitempty"`
	VoteLst int64  `json:"voteLst,omitempty"`
}

type BankNetwork struct {
	RecaptchaSiteKey string `json:"recaptcha_sitekey"`
	RecaptchaSecret  string `json:"recaptcha_secret"`
	Amount           int64  `json:"amount"`
	Fee              int64  `json:"fee"`
	Wallet           string `json:"wallet"`
	DataDir          string `json:"data_dir"`
}
