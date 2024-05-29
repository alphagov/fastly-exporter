package realtime

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// Process updates the metrics with data from the API response.
func Process(response *Response, serviceID, serviceName, _ string, m *Metrics, aggregateOnly bool) {
	const aggregateDC = "aggregate"

	for _, d := range response.Data {
		if aggregateOnly {
			process(serviceID, serviceName, aggregateDC, d.Aggregated, m)

			continue
		}

		for datacenter, stats := range d.Datacenter {
			process(serviceID, serviceName, datacenter, stats, m)
		}
	}
}

func process(serviceID, serviceName, datacenter string, stats Datacenter, m *Metrics) {
	m.AttackBlockedReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackBlockedReqBodyBytes))
	m.AttackBlockedReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackBlockedReqHeaderBytes))
	m.AttackLoggedReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackLoggedReqBodyBytes))
	m.AttackLoggedReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackLoggedReqHeaderBytes))
	m.AttackPassedReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackPassedReqBodyBytes))
	m.AttackPassedReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackPassedReqHeaderBytes))
	m.AttackReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackReqBodyBytes))
	m.AttackReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackReqHeaderBytes))
	m.AttackRespSynthBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.AttackRespSynthBytes))
	m.BackendReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BackendReqBodyBytes))
	m.BackendReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BackendReqHeaderBytes))
	m.BlacklistedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Blacklisted))
	m.BodySizeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BodySize))
	m.BotChallengeCompleteTokensCheckedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengeCompleteTokensChecked))
	m.BotChallengeCompleteTokensDisabledTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengeCompleteTokensDisabled))
	m.BotChallengeCompleteTokensFailedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengeCompleteTokensFailed))
	m.BotChallengeCompleteTokensIssuedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengeCompleteTokensIssued))
	m.BotChallengeCompleteTokensPassedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengeCompleteTokensPassed))
	m.BotChallengeStartsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengeStarts))
	m.BotChallengesFailedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengesFailed))
	m.BotChallengesIssuedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengesIssued))
	m.BotChallengesSucceededTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.BotChallengesSucceeded))
	m.ComputeBackendReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeBackendReqBodyBytesTotal))
	m.ComputeBackendReqErrorsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeBackendReqErrorsTotal))
	m.ComputeBackendReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeBackendReqHeaderBytesTotal))
	m.ComputeBackendReqTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeBackendReqTotal))
	m.ComputeBackendRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeBackendRespBodyBytesTotal))
	m.ComputeBackendRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeBackendRespHeaderBytesTotal))
	m.ComputeExecutionTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeExecutionTimeMilliseconds) / 10000.0)
	m.ComputeGlobalsLimitExceededTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeGlobalsLimitExceededTotal))
	m.ComputeGuestErrorsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeGuestErrorsTotal))
	m.ComputeHeapLimitExceededTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeHeapLimitExceededTotal))
	m.ComputeRAMUsedBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeRAMUsed))
	m.ComputeReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeReqBodyBytesTotal))
	m.ComputeReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeReqHeaderBytesTotal))
	m.ComputeRequestTimeBilledTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeRequestTimeBilledMilliseconds) / 10000.0)
	m.ComputeRequestTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeRequestTimeMilliseconds) / 10000.0)
	m.ComputeRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeRequests))
	m.ComputeResourceLimitExceedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeResourceLimitExceedTotal))
	m.ComputeRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeRespBodyBytesTotal))
	m.ComputeRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeRespHeaderBytesTotal))
	m.ComputeRespStatusTotal.WithLabelValues(serviceID, serviceName, datacenter, "1xx").Add(float64(stats.ComputeRespStatus1xx))
	m.ComputeRespStatusTotal.WithLabelValues(serviceID, serviceName, datacenter, "2xx").Add(float64(stats.ComputeRespStatus2xx))
	m.ComputeRespStatusTotal.WithLabelValues(serviceID, serviceName, datacenter, "3xx").Add(float64(stats.ComputeRespStatus3xx))
	m.ComputeRespStatusTotal.WithLabelValues(serviceID, serviceName, datacenter, "4xx").Add(float64(stats.ComputeRespStatus4xx))
	m.ComputeRespStatusTotal.WithLabelValues(serviceID, serviceName, datacenter, "5xx").Add(float64(stats.ComputeRespStatus5xx))
	m.ComputeRuntimeErrorsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeRuntimeErrorsTotal))
	m.ComputeStackLimitExceededTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ComputeStackLimitExceededTotal))
	m.DDOSActionBlackholeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionBlackhole))
	m.DDOSActionCloseTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionClose))
	m.DDOSActionDowngradeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionDowngrade))
	m.DDOSActionDowngradedConnectionsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionDowngradedConnections))
	m.DDOSActionLimitStreamsConnectionsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionLimitStreamsConnections))
	m.DDOSActionLimitStreamsRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionLimitStreamsRequests))
	m.DDOSActionTarpitAcceptTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionTarpitAccept))
	m.DDOSActionTarpitTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DDOSActionTarpit))
	m.DeliverSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DeliverSubCount))
	m.DeliverSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.DeliverSubTime))
	m.EdgeHitRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeHitRequests))
	m.EdgeHitRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeHitRespBodyBytes))
	m.EdgeHitRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeHitRespHeaderBytes))
	m.EdgeMissRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeMissRequests))
	m.EdgeMissRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeMissRespBodyBytes))
	m.EdgeMissRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeMissRespHeaderBytes))
	m.EdgeRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeRespBodyBytes))
	m.EdgeRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.EdgeRespHeaderBytes))
	m.EdgeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Edge))
	m.ErrorSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ErrorSubCount))
	m.ErrorSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ErrorSubTime))
	m.ErrorsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Errors))
	m.FanoutBackendReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutBackendReqBodyBytes))
	m.FanoutBackendReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutBackendReqHeaderBytes))
	m.FanoutBackendRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutBackendRespBodyBytes))
	m.FanoutBackendRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutBackendRespHeaderBytes))
	m.FanoutConnTimeMsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutConnTimeMs))
	m.FanoutRecvPublishesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutRecvPublishes))
	m.FanoutReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutReqBodyBytes))
	m.FanoutReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutReqHeaderBytes))
	m.FanoutRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutRespBodyBytes))
	m.FanoutRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutRespHeaderBytes))
	m.FanoutSendPublishesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FanoutSendPublishes))
	m.FetchSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FetchSubCount))
	m.FetchSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.FetchSubTime))
	m.HTTPTotal.WithLabelValues(serviceID, serviceName, datacenter, "1").Add(float64(stats.Requests - (stats.HTTP2 + stats.HTTP3)))
	m.HTTPTotal.WithLabelValues(serviceID, serviceName, datacenter, "2").Add(float64(stats.HTTP2))
	m.HTTPTotal.WithLabelValues(serviceID, serviceName, datacenter, "3").Add(float64(stats.HTTP3))
	m.HTTP2Total.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HTTP2))
	m.HTTP3Total.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HTTP3))
	m.HashSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HashSubCount))
	m.HashSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HashSubTime))
	m.HeaderSizeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HeaderSize))
	m.HitRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HitRespBodyBytes))
	m.HitSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HitSubCount))
	m.HitSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HitSubTime))
	m.HitsTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.HitsTime))
	m.HitsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Hits))
	m.IPv6Total.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.IPv6))
	m.ImgOptoRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoRespBodyBytes))
	m.ImgOptoRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoRespHeaderBytes))
	m.ImgOptoShieldRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoShieldRespBodyBytes))
	m.ImgOptoShieldRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoShieldRespHeaderBytes))
	m.ImgOptoShieldTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoShield))
	m.ImgOptoTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOpto))
	m.ImgOptoTransformRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoTransformRespBodyBytes))
	m.ImgOptoTransformRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoTransformRespHeaderBytes))
	m.ImgOptoTransformTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgOptoTransform))
	m.ImgVideoFramesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideoFrames))
	m.ImgVideoRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideoRespBodyBytes))
	m.ImgVideoRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideoRespHeaderBytes))
	m.ImgVideoShieldFramesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideoShieldFrames))
	m.ImgVideoShieldRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideoShieldRespBodyBytes))
	m.ImgVideoShieldRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideoShieldRespHeaderBytes))
	m.ImgVideoShieldTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideoShield))
	m.ImgVideoTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ImgVideo))
	m.KVStoreClassAOperationsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.KVStoreClassAOperations))
	m.KVStoreClassBOperationsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.KVStoreClassBOperations))
	m.LogBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.LogBytes))
	m.LoggingTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Logging))
	m.MissRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.MissRespBodyBytes))
	m.MissSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.MissSubCount))
	m.MissSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.MissSubTime))
	m.MissTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.MissTime))
	m.MissesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Misses))
	m.OTFPDeliverTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPDeliverTime))
	m.OTFPManifestTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPManifest))
	m.OTFPRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPRespBodyBytes))
	m.OTFPRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPRespHeaderBytes))
	m.OTFPShieldRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPShieldRespBodyBytes))
	m.OTFPShieldRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPShieldRespHeaderBytes))
	m.OTFPShieldTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPShieldTime))
	m.OTFPShieldTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPShield))
	m.OTFPTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFP))
	m.OTFPTransformRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPTransformRespBodyBytes))
	m.OTFPTransformRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPTransformRespHeaderBytes))
	m.OTFPTransformTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPTransformTime))
	m.OTFPTransformTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OTFPTransform))
	m.OriginCacheFetchRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginCacheFetchRespBodyBytes))
	m.OriginCacheFetchRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginCacheFetchRespHeaderBytes))
	m.OriginCacheFetchesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginCacheFetches))
	m.OriginFetchBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginFetchBodyBytes))
	m.OriginFetchHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginFetchHeaderBytes))
	m.OriginFetchRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginFetchRespBodyBytes))
	m.OriginFetchRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginFetchRespHeaderBytes))
	m.OriginFetchesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginFetches))
	m.OriginRevalidationsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.OriginRevalidations))
	m.PCITotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PCI))
	m.PassRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PassRespBodyBytes))
	m.PassSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PassSubCount))
	m.PassSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PassSubTime))
	m.PassTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PassTime))
	m.PassesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Passes))
	m.Pipe.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Pipe))
	m.PipeSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PipeSubCount))
	m.PipeSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PipeSubTime))
	m.PredeliverSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PredeliverSubCount))
	m.PredeliverSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PredeliverSubTime))
	m.PrehashSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PrehashSubCount))
	m.PrehashSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.PrehashSubTime))
	m.RecvSubCountTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.RecvSubCount))
	m.RecvSubTimeTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.RecvSubTime))
	m.ReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ReqBodyBytes))
	m.ReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ReqHeaderBytes))
	m.RequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Requests))
	m.RespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.RespBodyBytes))
	m.RespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.RespHeaderBytes))
	m.RestartTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Restart))
	m.SegBlockOriginFetchesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.SegBlockOriginFetches))
	m.SegBlockShieldFetchesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.SegBlockShieldFetches))
	m.ShieldCacheFetchesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldCacheFetches))
	m.ShieldFetchBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldFetchBodyBytes))
	m.ShieldFetchHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldFetchHeaderBytes))
	m.ShieldFetchRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldFetchRespBodyBytes))
	m.ShieldFetchRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldFetchRespHeaderBytes))
	m.ShieldFetchesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldFetches))
	m.ShieldHitRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldHitRequests))
	m.ShieldHitRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldHitRespBodyBytes))
	m.ShieldHitRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldHitRespHeaderBytes))
	m.ShieldMissRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldMissRequests))
	m.ShieldMissRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldMissRespBodyBytes))
	m.ShieldMissRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldMissRespHeaderBytes))
	m.ShieldRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldRespBodyBytes))
	m.ShieldRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldRespHeaderBytes))
	m.ShieldRevalidationsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.ShieldRevalidations))
	m.ShieldTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Shield))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "200").Add(float64(stats.Status200))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "204").Add(float64(stats.Status204))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "206").Add(float64(stats.Status206))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "301").Add(float64(stats.Status301))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "302").Add(float64(stats.Status302))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "304").Add(float64(stats.Status304))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "400").Add(float64(stats.Status400))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "401").Add(float64(stats.Status401))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "403").Add(float64(stats.Status403))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "404").Add(float64(stats.Status404))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "406").Add(float64(stats.Status406))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "416").Add(float64(stats.Status416))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "429").Add(float64(stats.Status429))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "500").Add(float64(stats.Status500))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "501").Add(float64(stats.Status501))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "502").Add(float64(stats.Status502))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "503").Add(float64(stats.Status503))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "504").Add(float64(stats.Status504))
	m.StatusCodeTotal.WithLabelValues(serviceID, serviceName, datacenter, "505").Add(float64(stats.Status505))
	m.StatusGroupTotal.WithLabelValues(serviceID, serviceName, datacenter, "1xx").Add(float64(stats.Status1xx))
	m.StatusGroupTotal.WithLabelValues(serviceID, serviceName, datacenter, "2xx").Add(float64(stats.Status2xx))
	m.StatusGroupTotal.WithLabelValues(serviceID, serviceName, datacenter, "3xx").Add(float64(stats.Status3xx))
	m.StatusGroupTotal.WithLabelValues(serviceID, serviceName, datacenter, "4xx").Add(float64(stats.Status4xx))
	m.StatusGroupTotal.WithLabelValues(serviceID, serviceName, datacenter, "5xx").Add(float64(stats.Status5xx))
	m.SynthsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Synths))
	m.TLSTotal.WithLabelValues(serviceID, serviceName, datacenter, "1.0").Add(float64(stats.TLSv10))
	m.TLSTotal.WithLabelValues(serviceID, serviceName, datacenter, "1.1").Add(float64(stats.TLSv11))
	m.TLSTotal.WithLabelValues(serviceID, serviceName, datacenter, "1.2").Add(float64(stats.TLSv12))
	m.TLSTotal.WithLabelValues(serviceID, serviceName, datacenter, "1.3").Add(float64(stats.TLSv13))
	m.UncacheableTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Uncacheable))
	m.VclOnComputeHitRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.VclOnComputeHitRequests))
	m.VclOnComputeMissRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.VclOnComputeMissRequests))
	m.VclOnComputePassRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.VclOnComputePassRequests))
	m.VclOnComputeErrorRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.VclOnComputeErrorRequests))
	m.VclOnComputeSynthRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.VclOnComputeSynthRequests))
	m.VclOnComputeEdgeHitRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.VclOnComputeEdgeHitRequests))
	m.VclOnComputeEdgeMissRequestsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.VclOnComputeEdgeMissRequests))
	m.VideoTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.Video))
	m.WAFBlockedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WAFBlocked))
	m.WAFLoggedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WAFLogged))
	m.WAFPassedTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WAFPassed))
	m.WebsocketBackendReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketBackendReqBodyBytes))
	m.WebsocketBackendReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketBackendReqHeaderBytes))
	m.WebsocketBackendRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketBackendRespBodyBytes))
	m.WebsocketBackendRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketBackendRespHeaderBytes))
	m.WebsocketConnTimeMsTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketConnTimeMs))
	m.WebsocketReqBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketReqBodyBytes))
	m.WebsocketReqHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketReqHeaderBytes))
	m.WebsocketRespBodyBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketRespBodyBytes))
	m.WebsocketRespHeaderBytesTotal.WithLabelValues(serviceID, serviceName, datacenter).Add(float64(stats.WebsocketRespHeaderBytes))
	processHistogram(stats.MissHistogram, m.MissDurationSeconds.WithLabelValues(serviceID, serviceName, datacenter))
	processObjectSizes(stats.ObjectSize1k, stats.ObjectSize10k, stats.ObjectSize100k, stats.ObjectSize1m, stats.ObjectSize10m, stats.ObjectSize100m, stats.ObjectSize1g, m.ObjectSizeBytes.WithLabelValues(serviceID, serviceName, datacenter))
}

func processHistogram(src map[string]uint64, obs prometheus.Observer) {
	for str, count := range src {
		ms, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		s := float64(ms) / 1e3
		for i := 0; i < int(count); i++ {
			obs.Observe(s)
		}
	}
}

func processObjectSizes(n1k, n10k, n100k, n1m, n10m, n100m, n1g uint64, obs prometheus.Observer) {
	for v, n := range map[uint64]uint64{
		1 * 1024:           n1k,
		10 * 1024:          n10k,
		100 * 1024:         n100k,
		1 * 1000 * 1024:    n1m,
		10 * 1000 * 1024:   n10m,
		100 * 1000 * 1024:  n100m,
		1000 * 1000 * 1024: n1g,
	} {
		for i := uint64(0); i < n; i++ {
			obs.Observe(float64(v))
		}
	}
}
