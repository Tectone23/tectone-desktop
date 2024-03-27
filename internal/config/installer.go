package config

import (
	"context"
	"encoding/json"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	ExampleInstallerConfig = `{
    "Version": 33,
    "AccountUpdatesStatsInterval": 5000000000,
    "AccountsRebuildSynchronousMode": 1,
    "AgreementIncomingBundlesQueueLength": 15,
    "AgreementIncomingProposalsQueueLength": 50,
    "AgreementIncomingVotesQueueLength": 20000,
    "AnnounceParticipationKey": true,
    "Archival": false,
    "BaseLoggerDebugLevel": 4,
    "BlockDBDir": "",
    "BlockServiceCustomFallbackEndpoints": "",
    "BlockServiceMemCap": 500000000,
    "BroadcastConnectionsLimit": -1,
    "CadaverDirectory": "",
    "CadaverSizeTarget": 0,
    "CatchpointDir": "",
    "CatchpointFileHistoryLength": 365,
    "CatchpointInterval": 10000,
    "CatchpointTracking": 0,
    "CatchupBlockDownloadRetryAttempts": 1000,
    "CatchupBlockValidateMode": 0,
    "CatchupFailurePeerRefreshRate": 10,
    "CatchupGossipBlockFetchTimeoutSec": 4,
    "CatchupHTTPBlockFetchTimeoutSec": 4,
    "CatchupLedgerDownloadRetryAttempts": 50,
    "CatchupParallelBlocks": 16,
    "ColdDataDir": "",
    "ConnectionsRateLimitingCount": 60,
    "ConnectionsRateLimitingWindowSeconds": 1,
    "CrashDBDir": "",
    "DNSBootstrapID": "<network>.algorand.network?backup=<network>.algorand.net&dedup=<name>.algorand-<network>.(network|net)",
    "DNSSecurityFlags": 1,
    "DeadlockDetection": 0,
    "DeadlockDetectionThreshold": 30,
    "DisableAPIAuth": false,
    "DisableLedgerLRUCache": false,
    "DisableLocalhostConnectionRateLimit": true,
    "DisableNetworking": false,
    "DisableOutgoingConnectionThrottling": false,
    "EnableAccountUpdatesStats": false,
    "EnableAgreementReporting": false,
    "EnableAgreementTimeMetrics": false,
    "EnableAssembleStats": false,
    "EnableBlockService": false,
    "EnableDeveloperAPI": false,
    "EnableExperimentalAPI": false,
    "EnableFollowMode": false,
    "EnableGossipBlockService": true,
    "EnableGossipService": true,
    "EnableIncomingMessageFilter": false,
    "EnableLedgerService": false,
    "EnableMetricReporting": false,
    "EnableOutgoingNetworkMessageFiltering": true,
    "EnableP2P": false,
    "EnablePingHandler": true,
    "EnableProcessBlockStats": false,
    "EnableProfiler": false,
    "EnableRequestLogger": false,
    "EnableRuntimeMetrics": false,
    "EnableTopAccountsReporting": false,
    "EnableTxBacklogAppRateLimiting": true,
    "EnableTxBacklogRateLimiting": true,
    "EnableTxnEvalTracer": false,
    "EnableUsageLog": false,
    "EnableVerbosedTransactionSyncLogging": false,
    "EndpointAddress": "127.0.0.1:0",
    "FallbackDNSResolverAddress": "",
    "ForceFetchTransactions": false,
    "ForceRelayMessages": false,
    "GossipFanout": 4,
    "HeartbeatUpdateInterval": 600,
    "HotDataDir": "",
    "IncomingConnectionsLimit": 2400,
    "IncomingMessageFilterBucketCount": 5,
    "IncomingMessageFilterBucketSize": 512,
    "LedgerSynchronousMode": 2,
    "LogArchiveDir": "",
    "LogArchiveMaxAge": "",
    "LogArchiveName": "node.archive.log",
    "LogFileDir": "",
    "LogSizeLimit": 1073741824,
    "MaxAPIBoxPerApplication": 100000,
    "MaxAPIResourcesPerAccount": 100000,
    "MaxAcctLookback": 4,
    "MaxBlockHistoryLookback": 0,
    "MaxCatchpointDownloadDuration": 43200000000000,
    "MaxConnectionsPerIP": 15,
    "MinCatchpointFileDownloadBytesPerSecond": 20480,
    "NetAddress": "",
    "NetworkMessageTraceServer": "",
    "NetworkProtocolVersion": "",
    "NodeExporterListenAddress": ":9100",
    "NodeExporterPath": "./node_exporter",
    "OptimizeAccountsDatabaseOnStartup": false,
    "OutgoingMessageFilterBucketCount": 3,
    "OutgoingMessageFilterBucketSize": 128,
    "P2PPersistPeerID": false,
    "P2PPrivateKeyLocation": "",
    "ParticipationKeysRefreshInterval": 60000000000,
    "PeerConnectionsUpdateInterval": 3600,
    "PeerPingPeriodSeconds": 0,
    "PriorityPeers": {},
    "ProposalAssemblyTime": 500000000,
    "PublicAddress": "",
    "ReconnectTime": 60000000000,
    "ReservedFDs": 256,
    "RestConnectionsHardLimit": 2048,
    "RestConnectionsSoftLimit": 1024,
    "RestReadTimeoutSeconds": 15,
    "RestWriteTimeoutSeconds": 120,
    "RunHosted": false,
    "StateproofDir": "",
    "StorageEngine": "sqlite",
    "SuggestedFeeBlockHistory": 3,
    "SuggestedFeeSlidingWindowSize": 50,
    "TLSCertFile": "",
    "TLSKeyFile": "",
    "TelemetryToLog": true,
    "TrackerDBDir": "",
    "TransactionSyncDataExchangeRate": 0,
    "TransactionSyncSignificantMessageThreshold": 0,
    "TxBacklogAppTxPerSecondRate": 100,
    "TxBacklogAppTxRateLimiterMaxSize": 1048576,
    "TxBacklogRateLimitingCongestionPct": 50,
    "TxBacklogReservedCapacityPerPeer": 20,
    "TxBacklogServiceRateWindowSeconds": 10,
    "TxBacklogSize": 26000,
    "TxIncomingFilterMaxSize": 500000,
    "TxIncomingFilteringFlags": 1,
    "TxPoolExponentialIncreaseFactor": 2,
    "TxPoolSize": 75000,
    "TxSyncIntervalSeconds": 60,
    "TxSyncServeResponseSize": 1000000,
    "TxSyncTimeoutSeconds": 30,
    "UseXForwardedForAddressField": "",
    "VerifiedTranscationsCacheSize": 150000
}`
)

type InstallerConfig struct {
	Version                                 int64  `json:"Version"`
	AccountUpdatesStatsInterval             int64  `json:"AccountUpdatesStatsInterval"`
	AccountsRebuildSynchronousMode          int64  `json:"AccountsRebuildSynchronousMode"`
	AgreementIncomingBundlesQueueLength     int64  `json:"AgreementIncomingBundlesQueueLength"`
	AgreementIncomingProposalsQueueLength   int64  `json:"AgreementIncomingProposalsQueueLength"`
	AgreementIncomingVotesQueueLength       int64  `json:"AgreementIncomingVotesQueueLength"`
	AnnounceParticipationKey                bool   `json:"AnnounceParticipationKey"`
	Archival                                bool   `json:"Archival"`
	BaseLoggerDebugLevel                    int64  `json:"BaseLoggerDebugLevel"`
	BlockDBDir                              string `json:"BlockDBDir"`
	BlockServiceCustomFallbackEndpoints     string `json:"BlockServiceCustomFallbackEndpoints"`
	BlockServiceMemCap                      int64  `json:"BlockServiceMemCap"`
	BroadcastConnectionsLimit               int64  `json:"BroadcastConnectionsLimit"`
	CadaverDirectory                        string `json:"CadaverDirectory"`
	CadaverSizeTarget                       int64  `json:"CadaverSizeTarget"`
	CatchpointDir                           string `json:"CatchpointDir"`
	CatchpointFileHistoryLength             int64  `json:"CatchpointFileHistoryLength"`
	CatchpointInterval                      int64  `json:"CatchpointInterval"`
	CatchpointTracking                      int64  `json:"CatchpointTracking"`
	CatchupBlockDownloadRetryAttempts       int64  `json:"CatchupBlockDownloadRetryAttempts"`
	CatchupBlockValidateMode                int64  `json:"CatchupBlockValidateMode"`
	CatchupFailurePeerRefreshRate           int64  `json:"CatchupFailurePeerRefreshRate"`
	CatchupGossipBlockFetchTimeoutSec       int64  `json:"CatchupGossipBlockFetchTimeoutSec"`
	CatchupHTTPBlockFetchTimeoutSec         int64  `json:"CatchupHTTPBlockFetchTimeoutSec"`
	CatchupLedgerDownloadRetryAttempts      int64  `json:"CatchupLedgerDownloadRetryAttempts"`
	CatchupParallelBlocks                   int64  `json:"CatchupParallelBlocks"`
	ColdDataDir                             string `json:"ColdDataDir"`
	ConnectionsRateLimitingCount            int64  `json:"ConnectionsRateLimitingCount"`
	ConnectionsRateLimitingWindowSeconds    int64  `json:"ConnectionsRateLimitingWindowSeconds"`
	CrashDBDir                              string `json:"CrashDBDir"`
	DNSBootstrapID                          string `json:"DNSBootstrapID"`
	DNSSecurityFlags                        int64  `json:"DNSSecurityFlags"`
	DeadlockDetection                       int64  `json:"DeadlockDetection"`
	DeadlockDetectionThreshold              int64  `json:"DeadlockDetectionThreshold"`
	DisableAPIAuth                          bool   `json:"DisableAPIAuth"`
	DisableLedgerLRUCache                   bool   `json:"DisableLedgerLRUCache"`
	DisableLocalhostConnectionRateLimit     bool   `json:"DisableLocalhostConnectionRateLimit"`
	DisableNetworking                       bool   `json:"DisableNetworking"`
	DisableOutgoingConnectionThrottling     bool   `json:"DisableOutgoingConnectionThrottling"`
	EnableAccountUpdatesStats               bool   `json:"EnableAccountUpdatesStats"`
	EnableAgreementReporting                bool   `json:"EnableAgreementReporting"`
	EnableAgreementTimeMetrics              bool   `json:"EnableAgreementTimeMetrics"`
	EnableAssembleStats                     bool   `json:"EnableAssembleStats"`
	EnableBlockService                      bool   `json:"EnableBlockService"`
	EnableDeveloperAPI                      bool   `json:"EnableDeveloperAPI"`
	EnableExperimentalAPI                   bool   `json:"EnableExperimentalAPI"`
	EnableFollowMode                        bool   `json:"EnableFollowMode"`
	EnableGossipBlockService                bool   `json:"EnableGossipBlockService"`
	EnableGossipService                     bool   `json:"EnableGossipService"`
	EnableIncomingMessageFilter             bool   `json:"EnableIncomingMessageFilter"`
	EnableLedgerService                     bool   `json:"EnableLedgerService"`
	EnableMetricReporting                   bool   `json:"EnableMetricReporting"`
	EnableOutgoingNetworkMessageFiltering   bool   `json:"EnableOutgoingNetworkMessageFiltering"`
	EnableP2P                               bool   `json:"EnableP2P"`
	EnablePingHandler                       bool   `json:"EnablePingHandler"`
	EnableProcessBlockStats                 bool   `json:"EnableProcessBlockStats"`
	EnableProfiler                          bool   `json:"EnableProfiler"`
	EnableRequestLogger                     bool   `json:"EnableRequestLogger"`
	EnableRuntimeMetrics                    bool   `json:"EnableRuntimeMetrics"`
	EnableTopAccountsReporting              bool   `json:"EnableTopAccountsReporting"`
	EnableTxBacklogAppRateLimiting          bool   `json:"EnableTxBacklogAppRateLimiting"`
	EnableTxBacklogRateLimiting             bool   `json:"EnableTxBacklogRateLimiting"`
	EnableTxnEvalTracer                     bool   `json:"EnableTxnEvalTracer"`
	EnableUsageLog                          bool   `json:"EnableUsageLog"`
	EnableVerbosedTransactionSyncLogging    bool   `json:"EnableVerbosedTransactionSyncLogging"`
	EndpointAddress                         string `json:"EndpointAddress"`
	FallbackDNSResolverAddress              string `json:"FallbackDNSResolverAddress"`
	ForceFetchTransactions                  bool   `json:"ForceFetchTransactions"`
	ForceRelayMessages                      bool   `json:"ForceRelayMessages"`
	GossipFanout                            int64  `json:"GossipFanout"`
	HeartbeatUpdateInterval                 int64  `json:"HeartbeatUpdateInterval"`
	HotDataDir                              int64  `json:"HotDataDir"`
	IncomingConnectionsLimit                int64  `json:"IncomingConnectionsLimit"`
	IncomingMessageFilterBucketCount        int64  `json:"IncomingMessageFilterBucketCount"`
	IncomingMessageFilterBucketSize         int64  `json:"IncomingMessageFilterBucketSize"`
	LedgerSynchronousMode                   int64  `json:"LedgerSynchronousMode"`
	LogArchiveDir                           string `json:"LogArchiveDir"`
	LogArchiveMaxAge                        string `json:"LogArchiveMaxAge"`
	LogArchiveName                          string `json:"LogArchiveName"`
	LogFileDir                              string `json:"LogFileDir"`
	LogSizeLimit                            int64  `json:"LogSizeLimit"`
	MaxAPIBoxPerApplication                 int64  `json:"MaxAPIBoxPerApplication"`
	MaxAPIResourcesPerAccount               int64  `json:"MaxAPIResourcesPerAccount"`
	MaxAcctLookback                         int64  `json:"MaxAcctLookback"`
	MaxBlockHistoryLookback                 int64  `json:"MaxBlockHistoryLookback"`
	MaxCatchpointDownloadDuration           int64  `json:"MaxCatchpointDownloadDuration"`
	MaxConnectionsPerIP                     int64  `json:"MaxConnectionsPerIP"`
	MinCatchpointFileDownloadBytesPerSecond int64  `json:"MinCatchpointFileDownloadBytesPerSecond"`
	NetAddress                              string `json:"NetAddress"`
	NetworkMessageTraceServer               string `json:"NetworkMessageTraceServer"`
	NetworkProtocolVersion                  string `json:"NetworkProtocolVersion"`
	NodeExporterListenAddress               string `json:"NodeExporterListenAddress"`
	NodeExporterPath                        string `json:"NodeExporterPath"`
	OptimizeAccountsDatabaseOnStartup       bool   `json:"OptimizeAccountsDatabaseOnStartup"`
	OutgoingMessageFilterBucketCount        int64  `json:"OutgoingMessageFilterBucketCount"`
	OutgoingMessageFilterBucketSize         int64  `json:"OutgoingMessageFilterBucketSize"`
	P2PPersistPeerID                        bool   `json:"P2PPersistPeerID"`
	P2PPrivateKeyLocation                   string `json:"P2PPrivateKeyLocation"`
	ParticipationKeysRefreshInterval        int64  `json:"ParticipationKeysRefreshInterval"`
	PeerConnectionsUpdateInterval           int64  `json:"PeerConnectionsUpdateInterval"`
	PeerPingPeriodSeconds                   int64  `json:"PeerPingPeriodSeconds"`
	// PriorityPeers": {},
	ProposalAssemblyTime                       int64  `json:"ProposalAssemblyTime"`
	PublicAddress                              string `json:"PublicAddress"`
	ReconnectTime                              int64  `json:"ReconnectTime"`
	ReservedFDs                                int64  `json:"ReservedFDs"`
	RestConnectionsHardLimit                   int64  `json:"RestConnectionsHardLimit"`
	RestConnectionsSoftLimit                   int64  `json:"RestConnectionsSoftLimit"`
	RestReadTimeoutSeconds                     int64  `json:"RestReadTimeoutSeconds"`
	RestWriteTimeoutSeconds                    int64  `json:"RestWriteTimeoutSeconds"`
	RunHosted                                  bool   `json:"RunHosted"`
	StateproofDir                              string `json:"StateproofDir"`
	StorageEngine                              string `json:"StorageEngine"`
	SuggestedFeeBlockHistory                   int64  `json:"SuggestedFeeBlockHistory"`
	SuggestedFeeSlidingWindowSize              int64  `json:"SuggestedFeeSlidingWindowSize"`
	TLSCertFile                                string `json:"TLSCertFile"`
	TLSKeyFile                                 string `json:"TLSKeyFile"`
	TelemetryToLog                             bool   `json:"TelemetryToLog"`
	TrackerDBDir                               string `json:"TrackerDBDir"`
	TransactionSyncDataExchangeRate            int64  `json:"TransactionSyncDataExchangeRate"`
	TransactionSyncSignificantMessageThreshold int64  `json:"TransactionSyncSignificantMessageThreshold"`
	TxBacklogAppTxPerSecondRate                int64  `json:"TxBacklogAppTxPerSecondRate"`
	TxBacklogAppTxRateLimiterMaxSize           int64  `json:"TxBacklogAppTxRateLimiterMaxSize"`
	TxBacklogRateLimitingCongestionPct         int64  `json:"TxBacklogRateLimitingCongestionPct"`
	TxBacklogReservedCapacityPerPeer           int64  `json:"TxBacklogReservedCapacityPerPeer"`
	TxBacklogServiceRateWindowSeconds          int64  `json:"TxBacklogServiceRateWindowSeconds"`
	TxBacklogSize                              int64  `json:"TxBacklogSize"`
	TxIncomingFilterMaxSize                    int64  `json:"TxIncomingFilterMaxSize"`
	TxIncomingFilteringFlags                   int64  `json:"TxIncomingFilteringFlags"`
	TxPoolExponentialIncreaseFactor            int64  `json:"TxPoolExponentialIncreaseFactor"`
	TxPoolSize                                 int64  `json:"TxPoolSize"`
	TxSyncIntervalSeconds                      int64  `json:"TxSyncIntervalSeconds"`
	TxSyncServeResponseSize                    int64  `json:"TxSyncServeResponseSize"`
	TxSyncTimeoutSeconds                       int64  `json:"TxSyncTimeoutSeconds"`
	UseXForwardedForAddressField               string `json:"UseXForwardedForAddressField"`
	VerifiedTranscationsCacheSize              int64  `json:"VerifiedTranscationsCacheSize"`
}

func GetExampleInstallerConfig(ctx context.Context) *InstallerConfig {
	var c *InstallerConfig
	if err := json.Unmarshal([]byte(ExampleInstallerConfig), c); err != nil {
		runtime.LogErrorf(ctx, "error could not unmarshal example installer config json: %s\n", err)
		return c
	}
	return c
}
