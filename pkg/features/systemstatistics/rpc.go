package systemstatistics

import "encoding/xml"

type SystemStatistics struct {
	XMLName    xml.Name   `xml:"rpc-reply"`
	//Text       string     `xml:",chardata"`
	//Junos      string     `xml:"junos,attr"`
	rpcReply    rpcReply  `xml:"rpc-reply"`
	Statistics   Statistics `xml:"statistics"`
		Cli        struct {
			Text   string `xml:",chardata"`
			Banner string `xml:"banner"`
		} `xml:"cli"`

}

type rpcReply struct {
	XMLName xml.Name `xml:"rpc-reply"`
}
type Statistics struct {
	Text string `xml:",chardata"`
	Tcp  TCP `xml:"tcp"`
	Udp UDP `xml:"udp"`
	Ip  IP `xml:"ip"`
	Icmp ICMP `xml:"icmp"`
	Arp  ARP `xml:"arp"`
	Ip6  IP6 `xml:"ip6"`
	Icmp6 ICMP6`xml:"icmp6"`
	Mpls MPLS`xml:"mpls"`
}

type IP struct {
	Text                                      string  `xml:",chardata"`
	PacketsReceived                           float64 `xml:"packets-received"`
	BadHeaderChecksums                        float64 `xml:"bad-header-checksums"`
	PacketsWithSizeSmallerThanMinimum         float64 `xml:"packets-with-size-smaller-than-minimum"`
	PacketsWithDataSizeLessThanDatalength     float64 `xml:"packets-with-data-size-less-than-datalength"`
	PacketsWithHeaderLengthLessThanDataSize   float64 `xml:"packets-with-header-length-less-than-data-size"`
	PacketsWithDataLengthLessThanHeaderlength float64 `xml:"packets-with-data-length-less-than-headerlength"`
	PacketsWithIncorrectVersionNumber         float64 `xml:"packets-with-incorrect-version-number"`
	PacketsDestinedToDeadNextHop              float64 `xml:"packets-destined-to-dead-next-hop"`
	FragmentsReceived                         float64 `xml:"fragments-received"`
	FragmentsDroppedDueToOutofspaceOrDup      float64 `xml:"fragments-dropped-due-to-outofspace-or-dup"`
	FragmentsDroppedDueToQueueoverflow        float64 `xml:"fragments-dropped-due-to-queueoverflow"`
	FragmentsDroppedAfterTimeout              float64 `xml:"fragments-dropped-after-timeout"`
	PacketsReassembledOk                      float64 `xml:"packets-reassembled-ok"`
	PacketsForThisHost                        float64 `xml:"packets-for-this-host"`
	PacketsForUnknownOrUnsupportedProtocol    float64 `xml:"packets-for-unknown-or-unsupported-protocol"`
	PacketsForwarded                          float64 `xml:"packets-forwarded"`
	PacketsNotForwardable                     float64 `xml:"packets-not-forwardable"`
	RedirectsSent                             float64 `xml:"redirects-sent"`
	PacketsSentFromThisHost                   float64 `xml:"packets-sent-from-this-host"`
	PacketsSentWithFabricatedIpHeader         float64 `xml:"packets-sent-with-fabricated-ip-header"`
	OutputPacketsDroppedDueToNoBufs           float64 `xml:"output-packets-dropped-due-to-no-bufs"`
	OutputPacketsDiscardedDueToNoRoute        float64 `xml:"output-packets-discarded-due-to-no-route"`
	OutputDatagramsFragmented                 float64 `xml:"output-datagrams-fragmented"`
	FragmentsCreated                          float64 `xml:"fragments-created"`
	DatagramsThatCanNotBeFragmented           float64 `xml:"datagrams-that-can-not-be-fragmented"`
	PacketsWithBadOptions                     float64 `xml:"packets-with-bad-options"`
	PacketsWithOptionsHandledWithoutError     float64 `xml:"packets-with-options-handled-without-error"`
	StrictSourceAndRecordRouteOptions         float64 `xml:"strict-source-and-record-route-options"`
	LooseSourceAndRecordRouteOptions          float64 `xml:"loose-source-and-record-route-options"`
	RecordRouteOptions                        float64 `xml:"record-route-options"`
	TimestampOptions                          float64 `xml:"timestamp-options"`
	TimestampAndAddressOptions                float64 `xml:"timestamp-and-address-options"`
	TimestampAndPrespecifiedAddressOptions    float64 `xml:"timestamp-and-prespecified-address-options"`
	OptionPacketsDroppedDueToRateLimit        float64 `xml:"option-packets-dropped-due-to-rate-limit"`
	RouterAlertOptions                        float64 `xml:"router-alert-options"`
	MulticastPacketsDropped                   float64 `xml:"multicast-packets-dropped"`
	PacketsDropped                            float64 `xml:"packets-dropped"`
	TransitRePacketsDroppedOnMgmtInterface    float64 `xml:"transit-re-packets-dropped-on-mgmt-interface"`
	PacketsUsedFirstNexthopInEcmpUnilist      float64 `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
	IncomingTtpoipPacketsReceived             float64 `xml:"incoming-ttpoip-packets-received"`
	IncomingTtpoipPacketsDropped              float64 `xml:"incoming-ttpoip-packets-dropped"`
	OutgoingTtpoipPacketsSent                 float64 `xml:"outgoing-ttpoip-packets-sent"`
	OutgoingTtpoipPacketsDropped              float64 `xml:"outgoing-ttpoip-packets-dropped"`
	IncomingRawipPacketsDroppedNoSocketBuffer float64 `xml:"incoming-rawip-packets-dropped-no-socket-buffer"`
	IncomingVirtualNodePacketsDelivered       float64 `xml:"incoming-virtual-node-packets-delivered"`
}

type IP6 struct {
	Text                                  string       `xml:",chardata"`
	TotalPacketsReceived                  float64      `xml:"total-packets-received"`
	Ip6PacketsWithSizeSmallerThanMinimum  float64      `xml:"ip6-packets-with-size-smaller-than-minimum"`
	PacketsWithDatasizeLessThanDataLength float64      `xml:"packets-with-datasize-less-than-data-length"`
	Ip6PacketsWithBadOptions              float64      `xml:"ip6-packets-with-bad-options"`
	Ip6PacketsWithIncorrectVersionNumber  float64      `xml:"ip6-packets-with-incorrect-version-number"`
	Ip6FragmentsReceived                  float64      `xml:"ip6-fragments-received"`
	DuplicateOrOutOfSpaceFragmentsDropped float64      `xml:"duplicate-or-out-of-space-fragments-dropped"`
	Ip6FragmentsDroppedAfterTimeout       float64      `xml:"ip6-fragments-dropped-after-timeout"`
	FragmentsThatExceededLimit            float64      `xml:"fragments-that-exceeded-limit"`
	Ip6PacketsReassembledOk               float64      `xml:"ip6-packets-reassembled-ok"`
	Ip6PacketsForThisHost                 float64      `xml:"ip6-packets-for-this-host"`
	Ip6PacketsForwarded                   float64      `xml:"ip6-packets-forwarded"`
	Ip6PacketsNotForwardable              float64      `xml:"ip6-packets-not-forwardable"`
	Ip6RedirectsSent                      float64      `xml:"ip6-redirects-sent"`
	Ip6PacketsSentFromThisHost            float64      `xml:"ip6-packets-sent-from-this-host"`
	Ip6PacketsSentWithFabricatedIpHeader  float64      `xml:"ip6-packets-sent-with-fabricated-ip-header"`
	Ip6OutputPacketsDroppedDueToNoBufs    float64      `xml:"ip6-output-packets-dropped-due-to-no-bufs"`
	Ip6OutputPacketsDiscardedDueToNoRoute float64      `xml:"ip6-output-packets-discarded-due-to-no-route"`
	Ip6OutputDatagramsFragmented          float64      `xml:"ip6-output-datagrams-fragmented"`
	Ip6FragmentsCreated                   float64      `xml:"ip6-fragments-created"`
	Ip6DatagramsThatCanNotBeFragmented    float64      `xml:"ip6-datagrams-that-can-not-be-fragmented"`
	PacketsThatViolatedScopeRules         float64      `xml:"packets-that-violated-scope-rules"`
	MulticastPacketsWhichWeDoNotJoin      float64      `xml:"multicast-packets-which-we-do-not-join"`
	Histogram                             string      `xml:"histogram"`
	Ip6nhTcp                              float64      `xml:"ip6nh-tcp"`
	Ip6nhUdp                              float64      `xml:"ip6nh-udp"`
	Ip6nhIcmp6                            float64      `xml:"ip6nh-icmp6"`
	PacketsWhoseHeadersAreNotContinuous   float64      `xml:"packets-whose-headers-are-not-continuous"`
	TunnelingPacketsThatCanNotFindGif     float64      `xml:"tunneling-packets-that-can-not-find-gif"`
	PacketsDiscardedDueToTooMayHeaders    float64      `xml:"packets-discarded-due-to-too-may-headers"`
	FailuresOfSourceAddressSelection      float64      `xml:"failures-of-source-address-selection"`
	HeaderType                            []HeaderType `xml:"header-type"`
	ForwardCacheHit                       float64      `xml:"forward-cache-hit"`
	ForwardCacheMiss                      float64      `xml:"forward-cache-miss"`
	Ip6PacketsDestinedToDeadNextHop       float64      `xml:"ip6-packets-destined-to-dead-next-hop"`
	Ip6OptionPacketsDroppedDueToRateLimit float64      `xml:"ip6-option-packets-dropped-due-to-rate-limit"`
	Ip6PacketsDropped                     float64      `xml:"ip6-packets-dropped"`
	PacketsDroppedDueToBadProtocol        float64      `xml:"packets-dropped-due-to-bad-protocol"`
	TransitRePacketDroppedOnMgmtInterface float64      `xml:"transit-re-packet-dropped-on-mgmt-interface"`
	PacketUsedFirstNexthopInEcmpUnilist   float64      `xml:"packet-used-first-nexthop-in-ecmp-unilist"`
}
type HeaderType struct {
	Text                            string  `xml:",chardata"`
	HeaderForSourceAddressSelection string  `xml:"header-for-source-address-selection"`
	LinkLocals                      float64 `xml:"link-locals"`
	Globals                         float64 `xml:"globals"`
	AddressScope                    float64 `xml:"address-scope"`
	HexValue                        float64 `xml:"hex-value"`
}

type UDP struct {
	Text                                              string  `xml:",chardata"`
	DatagramsReceived                                 float64 `xml:"datagrams-received"`
	DatagramsWithIncompleteHeader                     float64 `xml:"datagrams-with-incomplete-header"`
	DatagramsWithBadDatalengthField                   float64 `xml:"datagrams-with-bad-datalength-field"`
	DatagramsWithBadChecksum                          float64 `xml:"datagrams-with-bad-checksum"`
	DatagramsDroppedDueToNoSocket                     float64 `xml:"datagrams-dropped-due-to-no-socket"`
	BroadcastOrMulticastDatagramsDroppedDueToNoSocket float64 `xml:"broadcast-or-multicast-datagrams-dropped-due-to-no-socket"`
	DatagramsDroppedDueToFullSocketBuffers            float64 `xml:"datagrams-dropped-due-to-full-socket-buffers"`
	DatagramsNotForHashedPcb                          float64 `xml:"datagrams-not-for-hashed-pcb"`
	DatagramsDelivered                                float64 `xml:"datagrams-delivered"`
	DatagramsOutput                                   float64 `xml:"datagrams-output"`
}

type TCP struct {
	Text                                             string  `xml:",chardata"`
	PacketsSent                                      float64 `xml:"packets-sent"`
	SentDataPackets                                  float64 `xml:"sent-data-packets"`
	DataPacketsBytes                                 float64 `xml:"data-packets-bytes"`
	SentDataPacketsRetransmitted                     float64 `xml:"sent-data-packets-retransmitted"`
	RetransmittedBytes                               float64 `xml:"retransmitted-bytes"`
	SentDataUnnecessaryRetransmitted                 float64 `xml:"sent-data-unnecessary-retransmitted"`
	SentResendsByMtuDiscovery                        float64 `xml:"sent-resends-by-mtu-discovery"`
	SentAckOnlyPackets                               float64 `xml:"sent-ack-only-packets"`
	SentPacketsDelayed                               float64 `xml:"sent-packets-delayed"`
	SentUrgOnlyPackets                               float64 `xml:"sent-urg-only-packets"`
	SentWindowProbePackets                           float64 `xml:"sent-window-probe-packets"`
	SentWindowUpdatePackets                          float64 `xml:"sent-window-update-packets"`
	SentControlPackets                               float64 `xml:"sent-control-packets"`
	PacketsReceived                                  float64 `xml:"packets-received"`
	ReceivedAcks                                     float64 `xml:"received-acks"`
	AcksBytes                                        float64 `xml:"acks-bytes"`
	ReceivedDuplicateAcks                            float64 `xml:"received-duplicate-acks"`
	ReceivedAcksForUnsentData                        float64 `xml:"received-acks-for-unsent-data"`
	PacketsReceivedInSequence                        float64 `xml:"packets-received-in-sequence"`
	InSequenceBytes                                  float64 `xml:"in-sequence-bytes"`
	ReceivedCompletelyDuplicatePacket                float64 `xml:"received-completely-duplicate-packet"`
	DuplicateInBytes                                 float64 `xml:"duplicate-in-bytes"`
	ReceivedOldDuplicatePackets                      float64 `xml:"received-old-duplicate-packets"`
	ReceivedPacketsWithSomeDupliacteData             float64 `xml:"received-packets-with-some-dupliacte-data"`
	SomeDuplicateInBytes                             float64 `xml:"some-duplicate-in-bytes"`
	ReceivedOutOfOrderPackets                        float64 `xml:"received-out-of-order-packets"`
	OutOfOrderInBytes                                float64 `xml:"out-of-order-in-bytes"`
	ReceivedPacketsOfDataAfterWindow                 float64 `xml:"received-packets-of-data-after-window"`
	Bytes                                            float64 `xml:"bytes"`
	ReceivedWindowProbes                             float64 `xml:"received-window-probes"`
	ReceivedWindowUpdatePackets                      float64 `xml:"received-window-update-packets"`
	PacketsReceivedAfterClose                        float64 `xml:"packets-received-after-close"`
	ReceivedDiscardedForBadChecksum                  float64 `xml:"received-discarded-for-bad-checksum"`
	ReceivedDiscardedForBadHeaderOffset              float64 `xml:"received-discarded-for-bad-header-offset"`
	ReceivedDiscardedBecausePacketTooShort           float64 `xml:"received-discarded-because-packet-too-short"`
	ConnectionRequests                               float64 `xml:"connection-requests"`
	ConnectionAccepts                                float64 `xml:"connection-accepts"`
	BadConnectionAttempts                            float64 `xml:"bad-connection-attempts"`
	ListenQueueOverflows                             float64 `xml:"listen-queue-overflows"`
	BadRstWindow                                     float64 `xml:"bad-rst-window"`
	ConnectionsEstablished                           float64 `xml:"connections-established"`
	ConnectionsClosed                                float64 `xml:"connections-closed"`
	Drops                                            float64 `xml:"drops"`
	ConnectionsUpdatedRttOnClose                     float64 `xml:"connections-updated-rtt-on-close"`
	ConnectionsUpdatedVarianceOnClose                float64 `xml:"connections-updated-variance-on-close"`
	ConnectionsUpdatedSsthreshOnClose                float64 `xml:"connections-updated-ssthresh-on-close"`
	EmbryonicConnectionsDropped                      float64 `xml:"embryonic-connections-dropped"`
	SegmentsUpdatedRtt                               float64 `xml:"segments-updated-rtt"`
	Attempts                                         float64 `xml:"attempts"`
	RetransmitTimeouts                               float64 `xml:"retransmit-timeouts"`
	ConnectionsDroppedByRetransmitTimeout            float64 `xml:"connections-dropped-by-retransmit-timeout"`
	PersistTimeouts                                  float64 `xml:"persist-timeouts"`
	ConnectionsDroppedByPersistTimeout               float64 `xml:"connections-dropped-by-persist-timeout"`
	KeepaliveTimeouts                                float64 `xml:"keepalive-timeouts"`
	KeepaliveProbesSent                              float64 `xml:"keepalive-probes-sent"`
	KeepaliveConnectionsDropped                      float64 `xml:"keepalive-connections-dropped"`
	AckHeaderPredictions                             float64 `xml:"ack-header-predictions"`
	DataPacketHeaderPredictions                      float64 `xml:"data-packet-header-predictions"`
	SyncacheEntriesAdded                             float64 `xml:"syncache-entries-added"`
	Retransmitted                                    float64 `xml:"retransmitted"`
	Dupsyn                                           float64 `xml:"dupsyn"`
	Dropped                                          float64 `xml:"dropped"`
	Completed                                        float64 `xml:"completed"`
	BucketOverflow                                   float64 `xml:"bucket-overflow"`
	CacheOverflow                                    float64 `xml:"cache-overflow"`
	Reset                                            float64 `xml:"reset"`
	Stale                                            float64 `xml:"stale"`
	Aborted                                          float64 `xml:"aborted"`
	Badack                                           float64 `xml:"badack"`
	Unreach                                          float64 `xml:"unreach"`
	ZoneFailures                                     float64 `xml:"zone-failures"`
	CookiesSent                                      float64 `xml:"cookies-sent"`
	CookiesReceived                                  float64 `xml:"cookies-received"`
	SackRecoveryEpisodes                             float64 `xml:"sack-recovery-episodes"`
	SegmentRetransmits                               float64 `xml:"segment-retransmits"`
	ByteRetransmits                                  float64 `xml:"byte-retransmits"`
	SackOptionsReceived                              float64 `xml:"sack-options-received"`
	SackOpitionsSent                                 float64 `xml:"sack-opitions-sent"`
	SackScoreboardOverflow                           float64 `xml:"sack-scoreboard-overflow"`
	AcksSentInResponseButNotExactRsts                float64 `xml:"acks-sent-in-response-but-not-exact-rsts"`
	AcksSentInResponseToSynsOnEstablishedConnections float64 `xml:"acks-sent-in-response-to-syns-on-established-connections"`
	RcvPacketsDroppedDueToBadAddress                 float64 `xml:"rcv-packets-dropped-due-to-bad-address"`
	OutOfSequenceSegmentDrops                        float64 `xml:"out-of-sequence-segment-drops"`
	RstPackets                                       float64 `xml:"rst-packets"`
	IcmpPacketsIgnored                               float64 `xml:"icmp-packets-ignored"`
	SendPacketsDropped                               float64 `xml:"send-packets-dropped"`
	RcvPacketsDropped                                float64 `xml:"rcv-packets-dropped"`
	OutgoingSegmentsDropped                          float64 `xml:"outgoing-segments-dropped"`
	ReceivedSynfinDropped                            float64 `xml:"received-synfin-dropped"`
	ReceivedIpsecDropped                             float64 `xml:"received-ipsec-dropped"`
	ReceivedMacDropped                               float64 `xml:"received-mac-dropped"`
	ReceivedMinttlExceeded                           float64 `xml:"received-minttl-exceeded"`
	ListenstateBadflagsDropped                       float64 `xml:"listenstate-badflags-dropped"`
	FinwaitstateBadflagsDropped                      float64 `xml:"finwaitstate-badflags-dropped"`
	ReceivedDosAttack                                float64 `xml:"received-dos-attack"`
	ReceivedBadSynack                                float64 `xml:"received-bad-synack"`
	SyncacheZoneFull                                 float64 `xml:"syncache-zone-full"`
	ReceivedRstFirewallfilter                        float64 `xml:"received-rst-firewallfilter"`
	ReceivedNoackTimewait                            float64 `xml:"received-noack-timewait"`
	ReceivedNoTimewaitState                          float64 `xml:"received-no-timewait-state"`
	ReceivedRstTimewaitState                         float64 `xml:"received-rst-timewait-state"`
	ReceivedTimewaitDrops                            float64 `xml:"received-timewait-drops"`
	ReceivedBadaddrTimewaitState                     float64 `xml:"received-badaddr-timewait-state"`
	ReceivedAckoffInSynSentrcvd                      float64 `xml:"received-ackoff-in-syn-sentrcvd"`
	ReceivedBadaddrFirewall                          float64 `xml:"received-badaddr-firewall"`
	ReceivedNosynSynSent                             float64 `xml:"received-nosyn-syn-sent"`
	ReceivedBadrstSynSent                            float64 `xml:"received-badrst-syn-sent"`
	ReceivedBadrstListenState                        float64 `xml:"received-badrst-listen-state"`
	OptionMaxsegmentLength                           float64 `xml:"option-maxsegment-length"`
	OptionWindowLength                               float64 `xml:"option-window-length"`
	OptionTimestampLength                            float64 `xml:"option-timestamp-length"`
	OptionMd5Length                                  float64 `xml:"option-md5-length"`
	OptionAuthLength                                 float64 `xml:"option-auth-length"`
	OptionSackpermittedLength                        float64 `xml:"option-sackpermitted-length"`
	OptionSackLength                                 float64 `xml:"option-sack-length"`
	OptionAuthoptionLength                           float64 `xml:"option-authoption-length"`
}

type ARP struct {
	Text                                                     string `xml:",chardata"`
	DatagramsReceived                                        float64 `xml:"datagrams-received"`
	ArpRequestsReceived                                      float64 `xml:"arp-requests-received"`
	ArpRepliesReceived                                       float64 `xml:"arp-replies-received"`
	ResolutionRequestReceived                                float64 `xml:"resolution-request-received"`
	ResolutionRequestDropped                                 float64 `xml:"resolution-request-dropped"`
	UnrestrictedProxyRequests                                float64 `xml:"unrestricted-proxy-requests"`
	RestrictedProxyRequests                                  float64 `xml:"restricted-proxy-requests"`
	ReceivedProxyRequests                                    float64 `xml:"received-proxy-requests"`
	ProxyRequestsNotProxied                                  float64 `xml:"proxy-requests-not-proxied"`
	RestrictedProxyRequestsNotProxied                        float64 `xml:"restricted-proxy-requests-not-proxied"`
	DatagramsWithBogusInterface                              float64 `xml:"datagrams-with-bogus-interface"`
	DatagramsWithIncorrectLength                             float64 `xml:"datagrams-with-incorrect-length"`
	DatagramsForNonIpProtocol                                float64 `xml:"datagrams-for-non-ip-protocol"`
	DatagramsWithUnsupportedOpcode                           float64 `xml:"datagrams-with-unsupported-opcode"`
	DatagramsWithBadProtocolAddressLength                    float64 `xml:"datagrams-with-bad-protocol-address-length"`
	DatagramsWithBadHardwareAddressLength                    float64 `xml:"datagrams-with-bad-hardware-address-length"`
	DatagramsWithMulticastSourceAddress                      float64 `xml:"datagrams-with-multicast-source-address"`
	DatagramsWithMulticastTargetAddress                      float64 `xml:"datagrams-with-multicast-target-address"`
	DatagramsWithMyOwnHardwareAddress                        float64 `xml:"datagrams-with-my-own-hardware-address"`
	DatagramsForAnAddressNotOnTheInterface                   float64 `xml:"datagrams-for-an-address-not-on-the-interface"`
	DatagramsWithABroadcastSourceAddress                     float64 `xml:"datagrams-with-a-broadcast-source-address"`
	DatagramsWithSourceAddressDuplicateToMine                float64 `xml:"datagrams-with-source-address-duplicate-to-mine"`
	DatagramsWhichWereNotForMe                               float64 `xml:"datagrams-which-were-not-for-me"`
	PacketsDiscardedWaitingForResolution                     float64 `xml:"packets-discarded-waiting-for-resolution"`
	PacketsSentAfterWaitingForResolution                     float64 `xml:"packets-sent-after-waiting-for-resolution"`
	ArpRequestsSent                                          float64 `xml:"arp-requests-sent"`
	ArpRepliesSent                                           float64 `xml:"arp-replies-sent"`
	RequestsForMemoryDenied                                  float64 `xml:"requests-for-memory-denied"`
	RequestsDroppedOnEntry                                   float64 `xml:"requests-dropped-on-entry"`
	RequestsDroppedDuringRetry                               float64 `xml:"requests-dropped-during-retry"`
	RequestsDroppedDueToInterfaceDeletion                    float64 `xml:"requests-dropped-due-to-interface-deletion"`
	RequestsOnUnnumberedInterfaces                           float64 `xml:"requests-on-unnumbered-interfaces"`
	NewRequestsOnUnnumberedInterfaces                        float64 `xml:"new-requests-on-unnumbered-interfaces"`
	RepliesFromUnnumberedInterfaces                          float64 `xml:"replies-from-unnumbered-interfaces"`
	RequestsOnUnnumberedInterfaceWithNonSubnettedDonor       float64 `xml:"requests-on-unnumbered-interface-with-non-subnetted-donor"`
	RepliesFromUnnumberedInterfaceWithNonSubnettedDonor      float64 `xml:"replies-from-unnumbered-interface-with-non-subnetted-donor"`
	ArpPacketsRejectedAsFamilyIsConfiguredWithDenyArp        float64 `xml:"arp-packets-rejected-as-family-is-configured-with-deny-arp"`
	ArpResponsePacketsAreRejectedOnMcAeIclInterface          float64 `xml:"arp-response-packets-are-rejected-on-mc-ae-icl-interface"`
	ArpRepliesAreRejectedAsSourceAndDestinationIsSame        float64 `xml:"arp-replies-are-rejected-as-source-and-destination-is-same"`
	ArpProbeForProxyAddressReachableFromTheIncomingInterface float64 `xml:"arp-probe-for-proxy-address-reachable-from-the-incoming-interface"`
	ArpRequestDiscardedForVrrpSourceAddress                  float64 `xml:"arp-request-discarded-for-vrrp-source-address"`
	SelfArpRequestPacketReceivedOnIrbInterface               float64 `xml:"self-arp-request-packet-received-on-irb-interface"`
	ProxyArpRequestDiscardedAsSourceIpIsAProxyTarget         float64 `xml:"proxy-arp-request-discarded-as-source-ip-is-a-proxy-target"`
	ArpPacketsAreDroppedAsNexthopAllocationFailed            float64 `xml:"arp-packets-are-dropped-as-nexthop-allocation-failed"`
	ArpPacketsReceivedFromPeerVrrpRouterAndDiscarded         float64 `xml:"arp-packets-received-from-peer-vrrp-router-and-discarded"`
	ArpPacketsAreRejectedAsTargetIpArpResolveIsInProgress    float64 `xml:"arp-packets-are-rejected-as-target-ip-arp-resolve-is-in-progress"`
	GratArpPacketsAreIgnoredAsMacAddressIsNotChanged         float64 `xml:"grat-arp-packets-are-ignored-as-mac-address-is-not-changed"`
	ArpPacketsAreDroppedFromPeerVrrp                         float64 `xml:"arp-packets-are-dropped-from-peer-vrrp"`
	ArpPacketsAreDroppedAsDriverCallFailed                   float64 `xml:"arp-packets-are-dropped-as-driver-call-failed"`
	ArpPacketsAreDroppedAsSourceIsNotValidated               float64 `xml:"arp-packets-are-dropped-as-source-is-not-validated"`
	ArpSystemMax                                             float64 `xml:"arp-system-max"`
	ArpPublicMax                                             float64 `xml:"arp-public-max"`
	ArpIriMax                                                float64 `xml:"arp-iri-max"`
	ArpMgtMax                                                float64 `xml:"arp-mgt-max"`
	ArpPublicCnt                                             float64 `xml:"arp-public-cnt"`
	ArpIriCnt                                                float64 `xml:"arp-iri-cnt"`
	ArpMgtCnt                                                float64 `xml:"arp-mgt-cnt"`
	ArpSystemDrop                                            float64 `xml:"arp-system-drop"`
	ArpPublicDrop                                            float64 `xml:"arp-public-drop"`
	ArpIriDrop                                               float64 `xml:"arp-iri-drop"`
	ArpMgtDrop                                               float64 `xml:"arp-mgt-drop"`
}

type ICMP struct {
	Text                                       string `xml:",chardata"`
	DropsDueToRateLimit                        float64 `xml:"drops-due-to-rate-limit"`
	CallsToIcmpError                           float64 `xml:"calls-to-icmp-error"`
	ErrorsNotGeneratedBecauseOldMessageWasIcmp float64 `xml:"errors-not-generated-because-old-message-was-icmp"`
	Histogram                                  []ICMPHistogram `xml:"histogram"`
	MessagesWithBadCodeFields                                float64 `xml:"messages-with-bad-code-fields"`
	MessagesLessThanTheMinimumLength                         float64 `xml:"messages-less-than-the-minimum-length"`
	MessagesWithBadChecksum                                  float64 `xml:"messages-with-bad-checksum"`
	MessagesWithBadSourceAddress                             float64 `xml:"messages-with-bad-source-address"`
	MessagesWithBadLength                                    float64 `xml:"messages-with-bad-length"`
	EchoDropsWithBroadcastOrMulticastDestinatonAddress       float64 `xml:"echo-drops-with-broadcast-or-multicast-destinaton-address"`
	TimestampDropsWithBroadcastOrMulticastDestinationAddress float64 `xml:"timestamp-drops-with-broadcast-or-multicast-destination-address"`
	MessageResponsesGenerated                                float64 `xml:"message-responses-generated"`
}

type ICMPHistogram struct {
	Text                             string `xml:",chardata"`
	TypeOfHistogram                  string `xml:"type-of-histogram"`
	IcmpEchoReply                    float64 `xml:"icmp-echo-reply"`
	DestinationUnreachable           float64 `xml:"destination-unreachable"`
	IcmpEcho                         float64 `xml:"icmp-echo"`
	TimeStampReply                   float64 `xml:"time-stamp-reply"`
	TimeExceeded                     float64 `xml:"time-exceeded"`
	TimeStamp                        float64 `xml:"time-stamp"`
	AddressMaskRequest               float64 `xml:"address-mask-request"`
	AnEndpointChangedItsCookieSecret float64 `xml:"an-endpoint-changed-its-cookiesecret"`
}

type ICMP6 struct {
	Text                                            string `xml:",chardata"`
	ProtocolName                                    string `xml:"protocol-name"`
	CallsToIcmp6Error                               float64 `xml:"calls-to-icmp6-error"`
	ErrorsNotGeneratedBecauseOldMessageWasIcmpError float64 `xml:"errors-not-generated-because-old-message-was-icmp-error"`
	ErrorsNotGeneratedBecauseRateLimitation         float64 `xml:"errors-not-generated-because-rate-limitation"`
	OutputHistogram  ICMP6OutputHistogram `xml:"output-histogram"`
	Icmp6MessagesWithBadCodeFields float64 `xml:"icmp6-messages-with-bad-code-fields"`
	MessagesLessThanMinimumLength  float64 `xml:"messages-less-than-minimum-length"`
	BadChecksums                   float64 `xml:"bad-checksums"`
	Icmp6MessagesWithBadLength     float64 `xml:"icmp6-messages-with-bad-length"`
	InputHistogram  ICMP6InputHistogram `xml:"input-histogram"`
	HistogramOfErrorMessagesToBeGenerated string `xml:"histogram-of-error-messages-to-be-generated"`
	NoRoute                               float64 `xml:"no-route"`
	AdministrativelyProhibited            float64 `xml:"administratively-prohibited"`
	BeyondScope                           float64 `xml:"beyond-scope"`
	AddressUnreachable                    float64 `xml:"address-unreachable"`
	PortUnreachable                       float64 `xml:"port-unreachable"`
	PacketTooBig                          float64 `xml:"packet-too-big"`
	TimeExceedTransit                     float64 `xml:"time-exceed-transit"`
	TimeExceedReassembly                  float64 `xml:"time-exceed-reassembly"`
	ErroneousHeaderField                  float64 `xml:"erroneous-header-field"`
	UnrecognizedNextHeader                float64 `xml:"unrecognized-next-header"`
	UnrecognizedOption                    float64 `xml:"unrecognized-option"`
	Redirect                              float64 `xml:"redirect"`
	Unknown                               float64 `xml:"unknown"`
	Icmp6MessageResponsesGenerated        float64 `xml:"icmp6-message-responses-generated"`
	MessagesWithTooManyNdOptions          float64 `xml:"messages-with-too-many-nd-options"`
	NdSystemMax                           float64 `xml:"nd-system-max"`
	NdPublicMax                           float64 `xml:"nd-public-max"`
	NdIriMax                              float64 `xml:"nd-iri-max"`
	NdMgtMax                              float64 `xml:"nd-mgt-max"`
	NdPublicCnt                           float64 `xml:"nd-public-cnt"`
	NdIriCnt                              float64 `xml:"nd-iri-cnt"`
	NdMgtCnt                              float64 `xml:"nd-mgt-cnt"`
	NdSystemDrop                          float64 `xml:"nd-system-drop"`
	NdPublicDrop                          float64 `xml:"nd-public-drop"`
	NdIriDrop                             float64 `xml:"nd-iri-drop"`
	NdMgtDrop                             float64 `xml:"nd-mgt-drop"`
	Nd6NdpProxyRequests                   float64 `xml:"nd6-ndp-proxy-requests"`
	Nd6DadProxyRequests                   float64 `xml:"nd6-dad-proxy-requests"`
	Nd6NdpProxyResponses                  float64 `xml:"nd6-ndp-proxy-responses"`
	Nd6DadProxyConflicts                  float64 `xml:"nd6-dad-proxy-conflicts"`
	Nd6DupProxyResponses                  float64 `xml:"nd6-dup-proxy-responses"`
	Nd6NdpProxyResolveCnt                 float64 `xml:"nd6-ndp-proxy-resolve-cnt"`
	Nd6DadProxyResolveCnt                 float64 `xml:"nd6-dad-proxy-resolve-cnt"`
	Nd6DadProxyEqmacDrop                  float64 `xml:"nd6-dad-proxy-eqmac-drop"`
	Nd6DadProxyNomacDrop                  float64 `xml:"nd6-dad-proxy-nomac-drop"`
	Nd6NdpProxyUnrRequests                float64 `xml:"nd6-ndp-proxy-unr-requests"`
	Nd6DadProxyUnrRequests                float64 `xml:"nd6-dad-proxy-unr-requests"`
	Nd6NdpProxyUnrResponses               float64 `xml:"nd6-ndp-proxy-unr-responses"`
	Nd6DadProxyUnrConflicts               float64 `xml:"nd6-dad-proxy-unr-conflicts"`
	Nd6DadProxyUnrResponses               float64 `xml:"nd6-dad-proxy-unr-responses"`
	Nd6NdpProxyUnrResolveCnt              float64 `xml:"nd6-ndp-proxy-unr-resolve-cnt"`
	Nd6DadProxyUnrResolveCnt              float64 `xml:"nd6-dad-proxy-unr-resolve-cnt"`
	Nd6DadProxyUnrEqportDrop              float64 `xml:"nd6-dad-proxy-unr-eqport-drop"`
	Nd6DadProxyUnrNomacDrop               float64 `xml:"nd6-dad-proxy-unr-nomac-drop"`
	Nd6RequestsDroppedOnEntry             float64 `xml:"nd6-requests-dropped-on-entry"`
	Nd6RequestsDroppedDuringRetry         float64 `xml:"nd6-requests-dropped-during-retry"`
}

type ICMP6OutputHistogram struct {
	Text                    string `xml:",chardata"`
	Style                   string `xml:"style,attr"`
	HistogramType           string `xml:"histogram-type"`
	UnreachableIcmp6Packets float64 `xml:"unreachable-icmp6-packets"`
	Icmp6Echo               float64 `xml:"icmp6-echo"`
	Icmp6EchoReply          float64 `xml:"icmp6-echo-reply"`
	NeighborSolicitation    float64 `xml:"neighbor-solicitation"`
	NeighborAdvertisement   float64 `xml:"neighbor-advertisement"`
}

type ICMP6InputHistogram struct {
	Text                           string `xml:",chardata"`
	Style                          string `xml:"style,attr"`
	HistogramType                  string `xml:"histogram-type"`
	UnreachableIcmp6Packets        float64 `xml:"unreachable-icmp6-packets"`
	PacketTooBig                   float64 `xml:"packet-too-big"`
	TimeExceededIcmp6Packets       float64 `xml:"time-exceeded-icmp6-packets"`
	Icmp6Echo                      float64 `xml:"icmp6-echo"`
	Icmp6EchoReply                 float64 `xml:"icmp6-echo-reply"`
	RouterSolicitationIcmp6Packets float64 `xml:"router-solicitation-icmp6-packets"`
	NeighborSolicitation           float64 `xml:"neighbor-solicitation"`
	NeighborAdvertisement          float64 `xml:"neighbor-advertisement"`
}

type MPLS struct {
	Text                                      string `xml:",chardata"`
	TotalMplsPacketsReceived                  float64 `xml:"total-mpls-packets-received"`
	PacketsForwarded                          float64 `xml:"packets-forwarded"`
	PacketsDropped                            float64 `xml:"packets-dropped"`
	PacketsWithHeaderTooSmall                 float64 `xml:"packets-with-header-too-small"`
	AfterTaggingPacketsCanNotFitLinkMtu       float64 `xml:"after-tagging-packets-can-not-fit-link-mtu"`
	PacketsWithIpv4ExplicitNullTag            float64 `xml:"packets-with-ipv4-explicit-null-tag"`
	PacketsWithIpv4ExplicitNullChecksumErrors float64 `xml:"packets-with-ipv4-explicit-null-checksum-errors"`
	PacketsWithRouterAlertTag                 float64 `xml:"packets-with-router-alert-tag"`
	LspPingPackets                            float64 `xml:"lsp-ping-packets"`
	PacketsWithTtlExpired                     float64 `xml:"packets-with-ttl-expired"`
	PacketsWithTagEncodingError               float64 `xml:"packets-with-tag-encoding-error"`
	PacketsDiscardedDueToNoRoute              float64 `xml:"packets-discarded-due-to-no-route"`
	PacketsUsedFirstNexthopInEcmpUnilist      float64 `xml:"packets-used-first-nexthop-in-ecmp-unilist"`
	PacketsDroppedDueToIflDown                float64 `xml:"packets-dropped-due-to-ifl-down"`
	PacketsDroppedAtMplsSocketSend            float64 `xml:"packets-dropped-at-mpls-socket-send"`
	PacketsForwardedAtMplsSocketSend          float64 `xml:"packets-forwarded-at-mpls-socket-send"`
	PacketsDroppedAtP2mpCnhOutput             float64 `xml:"packets-dropped-at-p2mp-cnh-output"`
}