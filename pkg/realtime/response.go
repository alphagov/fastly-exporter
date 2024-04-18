package realtime

// Response models the response from rt.fastly.com. It can get quite large; when
// there are lots of services being monitored, unmarshaling to this type is the
// CPU bottleneck of the program.
type Response struct {
	Timestamp      uint64 `json:"Timestamp"`
	AggregateDelay int64  `json:"AggregateDelay"`
	Data           []struct {
		Datacenter map[string]Datacenter `json:"datacenter"`
		Aggregated Datacenter            `json:"aggregated"`
		Recorded   uint64                `json:"recorded"`
	} `json:"Data"`
	Error string `json:"error"`
}

// Datacenter models the per-datacenter portion of the rt.fastly.com response.
type Datacenter struct {
	AttackBlockedReqBodyBytes            uint64            `json:"attack_blocked_req_body_bytes"`
	AttackBlockedReqHeaderBytes          uint64            `json:"attack_blocked_req_header_bytes"`
	AttackLoggedReqBodyBytes             uint64            `json:"attack_logged_req_body_bytes"`
	AttackLoggedReqHeaderBytes           uint64            `json:"attack_logged_req_header_bytes"`
	AttackPassedReqBodyBytes             uint64            `json:"attack_passed_req_body_bytes"`
	AttackPassedReqHeaderBytes           uint64            `json:"attack_passed_req_header_bytes"`
	AttackReqBodyBytes                   uint64            `json:"attack_req_body_bytes"`
	AttackReqHeaderBytes                 uint64            `json:"attack_req_header_bytes"`
	AttackRespSynthBytes                 uint64            `json:"attack_resp_synth_bytes"`
	BackendReqBodyBytes                  uint64            `json:"bereq_body_bytes"`
	BackendReqHeaderBytes                uint64            `json:"bereq_header_bytes"`
	Blacklisted                          uint64            `json:"blacklist"`
	BodySize                             uint64            `json:"body_size"`
	BotChallengeCompleteTokensChecked    uint64            `json:"bot_challenge_complete_tokens_checked"`
	BotChallengeCompleteTokensDisabled   uint64            `json:"bot_challenge_complete_tokens_disabled"`
	BotChallengeCompleteTokensFailed     uint64            `json:"bot_challenge_complete_tokens_failed"`
	BotChallengeCompleteTokensIssued     uint64            `json:"bot_challenge_complete_tokens_issued"`
	BotChallengeCompleteTokensPassed     uint64            `json:"bot_challenge_complete_tokens_passed"`
	BotChallengeStarts                   uint64            `json:"bot_challenge_starts"`
	BotChallengesFailed                  uint64            `json:"bot_challenges_failed"`
	BotChallengesIssued                  uint64            `json:"bot_challenges_issued"`
	BotChallengesSucceeded               uint64            `json:"bot_challenges_succeeded"`
	ComputeBackendReqBodyBytesTotal      uint64            `json:"compute_bereq_body_bytes"`
	ComputeBackendReqErrorsTotal         uint64            `json:"compute_bereq_errors"`
	ComputeBackendReqHeaderBytesTotal    uint64            `json:"compute_bereq_header_bytes"`
	ComputeBackendReqTotal               uint64            `json:"compute_bereqs"`
	ComputeBackendRespBodyBytesTotal     uint64            `json:"compute_beresp_body_bytes"`
	ComputeBackendRespHeaderBytesTotal   uint64            `json:"compute_beresp_header_bytes"`
	ComputeExecutionTimeMilliseconds     uint64            `json:"compute_execution_time_ms"`
	ComputeGlobalsLimitExceededTotal     uint64            `json:"compute_globals_limit_exceeded"`
	ComputeGuestErrorsTotal              uint64            `json:"compute_guest_errors"`
	ComputeHeapLimitExceededTotal        uint64            `json:"compute_heap_limit_exceeded"`
	ComputeRAMUsed                       uint64            `json:"compute_ram_used"`
	ComputeReqBodyBytesTotal             uint64            `json:"compute_req_body_bytes"`
	ComputeReqHeaderBytesTotal           uint64            `json:"compute_req_header_bytes"`
	ComputeRequestTimeBilledMilliseconds uint64            `json:"compute_request_time_billed_ms"`
	ComputeRequestTimeMilliseconds       uint64            `json:"compute_request_time_ms"`
	ComputeRequests                      uint64            `json:"compute_requests"`
	ComputeResourceLimitExceedTotal      uint64            `json:"compute_resource_limit_exceeded"`
	ComputeRespBodyBytesTotal            uint64            `json:"compute_resp_body_bytes"`
	ComputeRespHeaderBytesTotal          uint64            `json:"compute_resp_header_bytes"`
	ComputeRespStatus1xx                 uint64            `json:"compute_resp_status_1xx"`
	ComputeRespStatus2xx                 uint64            `json:"compute_resp_status_2xx"`
	ComputeRespStatus3xx                 uint64            `json:"compute_resp_status_3xx"`
	ComputeRespStatus4xx                 uint64            `json:"compute_resp_status_4xx"`
	ComputeRespStatus5xx                 uint64            `json:"compute_resp_status_5xx"`
	ComputeRuntimeErrorsTotal            uint64            `json:"compute_runtime_errors"`
	ComputeStackLimitExceededTotal       uint64            `json:"compute_stack_limit_exceeded"`
	DDOSActionBlackhole                  uint64            `json:"ddos_action_blackhole"`
	DDOSActionClose                      uint64            `json:"ddos_action_close"`
	DDOSActionDowngrade                  uint64            `json:"h2o_ddos_action_downgrade"`
	DDOSActionDowngradedConnections      uint64            `json:"h2o_ddos_action_downgraded_connections"`
	DDOSActionLimitStreamsConnections    uint64            `json:"ddos_action_limit_streams_connections"`
	DDOSActionLimitStreamsRequests       uint64            `json:"ddos_action_limit_streams_requests"`
	DDOSActionTarpit                     uint64            `json:"ddos_action_tarpit"`
	DDOSActionTarpitAccept               uint64            `json:"ddos_action_tarpit_accept"`
	DeliverSubCount                      uint64            `json:"deliver_sub_count"`
	DeliverSubTime                       uint64            `json:"deliver_sub_time"`
	Edge                                 uint64            `json:"edge_requests"`
	EdgeHitRequests                      uint64            `json:"edge_hit_requests"`
	EdgeHitRespBodyBytes                 uint64            `json:"edge_hit_resp_body_bytes"`
	EdgeHitRespHeaderBytes               uint64            `json:"edge_hit_resp_header_bytes"`
	EdgeMissRequests                     uint64            `json:"edge_miss_requests"`
	EdgeMissRespBodyBytes                uint64            `json:"edge_miss_resp_body_bytes"`
	EdgeMissRespHeaderBytes              uint64            `json:"edge_miss_resp_header_bytes"`
	EdgeRespBodyBytes                    uint64            `json:"edge_resp_body_bytes"`
	EdgeRespHeaderBytes                  uint64            `json:"edge_resp_header_bytes"`
	ErrorSubCount                        uint64            `json:"error_sub_count"`
	ErrorSubTime                         uint64            `json:"error_sub_time"`
	Errors                               uint64            `json:"errors"`
	FanoutBackendReqBodyBytes            uint64            `json:"fanout_bereq_body_bytes"`
	FanoutBackendReqHeaderBytes          uint64            `json:"fanout_bereq_header_bytes"`
	FanoutBackendRespBodyBytes           uint64            `json:"fanout_beresp_body_bytes"`
	FanoutBackendRespHeaderBytes         uint64            `json:"fanout_beresp_header_bytes"`
	FanoutConnTimeMs                     uint64            `json:"fanout_conn_time_ms"`
	FanoutRecvPublishes                  uint64            `json:"fanout_recv_publishes"`
	FanoutReqBodyBytes                   uint64            `json:"fanout_req_body_bytes"`
	FanoutReqHeaderBytes                 uint64            `json:"fanout_req_header_bytes"`
	FanoutRespBodyBytes                  uint64            `json:"fanout_resp_body_bytes"`
	FanoutRespHeaderBytes                uint64            `json:"fanout_resp_header_bytes"`
	FanoutSendPublishes                  uint64            `json:"fanout_send_publishes"`
	FetchSubCount                        uint64            `json:"fetch_sub_count"`
	FetchSubTime                         uint64            `json:"fetch_sub_time"`
	HTTP2                                uint64            `json:"http2"`
	HTTP3                                uint64            `json:"http3"`
	HashSubCount                         uint64            `json:"hash_sub_count"`
	HashSubTime                          uint64            `json:"hash_sub_time"`
	HeaderSize                           uint64            `json:"header_size"`
	HitRespBodyBytes                     uint64            `json:"hit_resp_body_bytes"`
	HitSubCount                          uint64            `json:"hit_sub_count"`
	HitSubTime                           uint64            `json:"hit_sub_time"`
	Hits                                 uint64            `json:"hits"`
	HitsTime                             float64           `json:"hits_time"`
	IPv6                                 uint64            `json:"ipv6"`
	ImgOpto                              uint64            `json:"imgopto"`
	ImgOptoRespBodyBytes                 uint64            `json:"imgopto_resp_body_bytes"`
	ImgOptoRespHeaderBytes               uint64            `json:"imgopto_resp_header_bytes"`
	ImgOptoShield                        uint64            `json:"imgopto_shield"`
	ImgOptoShieldRespBodyBytes           uint64            `json:"imgopto_shield_resp_body_bytes"`
	ImgOptoShieldRespHeaderBytes         uint64            `json:"imgopto_shield_resp_header_bytes"`
	ImgOptoTransform                     uint64            `json:"imgopto_transforms"`
	ImgOptoTransformRespBodyBytes        uint64            `json:"imgopto_transform_resp_body_bytes"`
	ImgOptoTransformRespHeaderBytes      uint64            `json:"imgopto_transform_resp_header_bytes"`
	ImgVideo                             uint64            `json:"imgvideo"`
	ImgVideoFrames                       uint64            `json:"imgvideo_frames"`
	ImgVideoRespBodyBytes                uint64            `json:"imgvideo_resp_body_bytes"`
	ImgVideoRespHeaderBytes              uint64            `json:"imgvideo_resp_header_bytes"`
	ImgVideoShield                       uint64            `json:"imgvideo_shield"`
	ImgVideoShieldFrames                 uint64            `json:"imgvideo_shield_frames"`
	ImgVideoShieldRespBodyBytes          uint64            `json:"imgvideo_shield_resp_body_bytes"`
	ImgVideoShieldRespHeaderBytes        uint64            `json:"imgvideo_shield_resp_header_bytes"`
	KVStoreClassAOperations              uint64            `json:"kv_store_class_a_operations"`
	KVStoreClassBOperations              uint64            `json:"kv_store_class_b_operations"`
	LogBytes                             uint64            `json:"log_bytes"`
	Logging                              uint64            `json:"logging"`
	MissHistogram                        map[string]uint64 `json:"miss_histogram"`
	MissRespBodyBytes                    uint64            `json:"miss_resp_body_bytes"`
	MissSubCount                         uint64            `json:"miss_sub_count"`
	MissSubTime                          uint64            `json:"miss_sub_time"`
	MissTime                             float64           `json:"miss_time"`
	Misses                               uint64            `json:"miss"`
	OTFP                                 uint64            `json:"otfp"`
	OTFPDeliverTime                      uint64            `json:"otfp_deliver_time"`
	OTFPManifest                         uint64            `json:"otfp_manifests"`
	OTFPRespBodyBytes                    uint64            `json:"otfp_resp_body_bytes"`
	OTFPRespHeaderBytes                  uint64            `json:"otfp_resp_header_bytes"`
	OTFPShield                           uint64            `json:"otfp_shield"`
	OTFPShieldRespBodyBytes              uint64            `json:"otfp_shield_resp_body_bytes"`
	OTFPShieldRespHeaderBytes            uint64            `json:"otfp_shield_resp_header_bytes"`
	OTFPShieldTime                       uint64            `json:"otfp_shield_time"`
	OTFPTransform                        uint64            `json:"otfp_transforms"`
	OTFPTransformRespBodyBytes           uint64            `json:"otfp_transform_resp_body_bytes"`
	OTFPTransformRespHeaderBytes         uint64            `json:"otfp_transform_resp_header_bytes"`
	OTFPTransformTime                    uint64            `json:"otfp_transform_time"`
	ObjectSize100k                       uint64            `json:"object_size_100k"`
	ObjectSize100m                       uint64            `json:"object_size_100m"`
	ObjectSize10k                        uint64            `json:"object_size_10k"`
	ObjectSize10m                        uint64            `json:"object_size_10m"`
	ObjectSize1g                         uint64            `json:"object_size_1g"`
	ObjectSize1k                         uint64            `json:"object_size_1k"`
	ObjectSize1m                         uint64            `json:"object_size_1m"`
	ObjectSizeOther                      uint64            `json:"object_size_other"`
	OriginCacheFetchRespBodyBytes        uint64            `json:"origin_cache_fetch_resp_body_bytes"`
	OriginCacheFetchRespHeaderBytes      uint64            `json:"origin_cache_fetch_resp_header_bytes"`
	OriginCacheFetches                   uint64            `json:"origin_cache_fetches"`
	OriginFetchBodyBytes                 uint64            `json:"origin_fetch_body_bytes"`
	OriginFetchHeaderBytes               uint64            `json:"origin_fetch_header_bytes"`
	OriginFetchRespBodyBytes             uint64            `json:"origin_fetch_resp_body_bytes"`
	OriginFetchRespHeaderBytes           uint64            `json:"origin_fetch_resp_header_bytes"`
	OriginFetches                        uint64            `json:"origin_fetches"`
	OriginRevalidations                  uint64            `json:"origin_revalidations"`
	PCI                                  uint64            `json:"pci"`
	PassRespBodyBytes                    uint64            `json:"pass_resp_body_bytes"`
	PassSubCount                         uint64            `json:"pass_sub_count"`
	PassSubTime                          uint64            `json:"pass_sub_time"`
	PassTime                             float64           `json:"pass_time"`
	Passes                               uint64            `json:"pass"`
	Pipe                                 uint64            `json:"pipe"`
	PipeSubCount                         uint64            `json:"pipe_sub_count"`
	PipeSubTime                          uint64            `json:"pipe_sub_time"`
	PredeliverSubCount                   uint64            `json:"predeliver_sub_count"`
	PredeliverSubTime                    uint64            `json:"predeliver_sub_time"`
	PrehashSubCount                      uint64            `json:"prehash_sub_count"`
	PrehashSubTime                       uint64            `json:"prehash_sub_time"`
	RecvSubCount                         uint64            `json:"recv_sub_count"`
	RecvSubTime                          uint64            `json:"recv_sub_time"`
	ReqBodyBytes                         uint64            `json:"req_body_bytes"`
	ReqHeaderBytes                       uint64            `json:"req_header_bytes"`
	Requests                             uint64            `json:"requests"`
	RespBodyBytes                        uint64            `json:"resp_body_bytes"`
	RespHeaderBytes                      uint64            `json:"resp_header_bytes"`
	Restart                              uint64            `json:"restarts"`
	SegBlockOriginFetches                uint64            `json:"segblock_origin_fetches"`
	SegBlockShieldFetches                uint64            `json:"segblock_shield_fetches"`
	Shield                               uint64            `json:"shield"`
	ShieldCacheFetches                   uint64            `json:"shield_cache_fetches"`
	ShieldFetchBodyBytes                 uint64            `json:"shield_fetch_body_bytes"`
	ShieldFetchHeaderBytes               uint64            `json:"shield_fetch_header_bytes"`
	ShieldFetchRespBodyBytes             uint64            `json:"shield_fetch_resp_body_bytes"`
	ShieldFetchRespHeaderBytes           uint64            `json:"shield_fetch_resp_header_bytes"`
	ShieldFetches                        uint64            `json:"shield_fetches"`
	ShieldHitRequests                    uint64            `json:"shield_hit_requests"`
	ShieldHitRespBodyBytes               uint64            `json:"shield_hit_resp_body_bytes"`
	ShieldHitRespHeaderBytes             uint64            `json:"shield_hit_resp_header_bytes"`
	ShieldMissRequests                   uint64            `json:"shield_miss_requests"`
	ShieldMissRespBodyBytes              uint64            `json:"shield_miss_resp_body_bytes"`
	ShieldMissRespHeaderBytes            uint64            `json:"shield_miss_resp_header_bytes"`
	ShieldRespBodyBytes                  uint64            `json:"shield_resp_body_bytes"`
	ShieldRespHeaderBytes                uint64            `json:"shield_resp_header_bytes"`
	ShieldRevalidations                  uint64            `json:"shield_revalidations"`
	Status1xx                            uint64            `json:"status_1xx"`
	Status200                            uint64            `json:"status_200"`
	Status204                            uint64            `json:"status_204"`
	Status206                            uint64            `json:"status_206"`
	Status2xx                            uint64            `json:"status_2xx"`
	Status301                            uint64            `json:"status_301"`
	Status302                            uint64            `json:"status_302"`
	Status304                            uint64            `json:"status_304"`
	Status3xx                            uint64            `json:"status_3xx"`
	Status400                            uint64            `json:"status_400"`
	Status401                            uint64            `json:"status_401"`
	Status403                            uint64            `json:"status_403"`
	Status404                            uint64            `json:"status_404"`
	Status406                            uint64            `json:"status_406"`
	Status416                            uint64            `json:"status_416"`
	Status429                            uint64            `json:"status_429"`
	Status4xx                            uint64            `json:"status_4xx"`
	Status500                            uint64            `json:"status_500"`
	Status501                            uint64            `json:"status_501"`
	Status502                            uint64            `json:"status_502"`
	Status503                            uint64            `json:"status_503"`
	Status504                            uint64            `json:"status_504"`
	Status505                            uint64            `json:"status_505"`
	Status5xx                            uint64            `json:"status_5xx"`
	Synths                               uint64            `json:"synth"`
	TLS                                  uint64            `json:"tls"`
	TLSHandshakeSentBytes                uint64            `json:"tls_handshake_sent_bytes"`
	TLSv10                               uint64            `json:"tls_v10"`
	TLSv11                               uint64            `json:"tls_v11"`
	TLSv12                               uint64            `json:"tls_v12"`
	TLSv13                               uint64            `json:"tls_v13"`
	Uncacheable                          uint64            `json:"uncacheable"`
	Video                                uint64            `json:"video"`
	WAFBlocked                           uint64            `json:"waf_blocked"`
	WAFLogged                            uint64            `json:"waf_logged"`
	WAFPassed                            uint64            `json:"waf_passed"`
	WebsocketBackendReqBodyBytes         uint64            `json:"websocket_bereq_body_bytes"`
	WebsocketBackendReqHeaderBytes       uint64            `json:"websocket_bereq_header_bytes"`
	WebsocketBackendRespBodyBytes        uint64            `json:"websocket_beresp_body_bytes"`
	WebsocketBackendRespHeaderBytes      uint64            `json:"websocket_beresp_header_bytes"`
	WebsocketConnTimeMs                  uint64            `json:"websocket_conn_time_ms"`
	WebsocketReqBodyBytes                uint64            `json:"websocket_req_body_bytes"`
	WebsocketReqHeaderBytes              uint64            `json:"websocket_req_header_bytes"`
	WebsocketRespBodyBytes               uint64            `json:"websocket_resp_body_bytes"`
	WebsocketRespHeaderBytes             uint64            `json:"websocket_resp_header_bytes"`
}
