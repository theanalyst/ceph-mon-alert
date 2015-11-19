package main

type ClusterReport struct {
	Auth struct {
		FirstCommitted int `json:"first_committed"`
		LastCommitted  int `json:"last_committed"`
		NumSecrets     int `json:"num_secrets"`
	} `json:"auth"`
	ClusterFingerprint string `json:"cluster_fingerprint"`
	Commit             string `json:"commit"`
	Crushmap           struct {
		Buckets []struct {
			Alg   string `json:"alg"`
			Hash  string `json:"hash"`
			ID    int    `json:"id"`
			Items []struct {
				ID     int `json:"id"`
				Pos    int `json:"pos"`
				Weight int `json:"weight"`
			} `json:"items"`
			Name     string `json:"name"`
			TypeID   int    `json:"type_id"`
			TypeName string `json:"type_name"`
			Weight   int    `json:"weight"`
		} `json:"buckets"`
		Devices []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"devices"`
		Rules []struct {
			MaxSize  int    `json:"max_size"`
			MinSize  int    `json:"min_size"`
			RuleID   int    `json:"rule_id"`
			RuleName string `json:"rule_name"`
			Ruleset  int    `json:"ruleset"`
			Steps    []struct {
				Item     int    `json:"item"`
				ItemName string `json:"item_name"`
				Op       string `json:"op"`
			} `json:"steps"`
			Type int `json:"type"`
		} `json:"rules"`
		Tunables struct {
			AllowedBucketAlgs        int    `json:"allowed_bucket_algs"`
			ChooseLocalFallbackTries int    `json:"choose_local_fallback_tries"`
			ChooseLocalTries         int    `json:"choose_local_tries"`
			ChooseTotalTries         int    `json:"choose_total_tries"`
			ChooseleafDescendOnce    int    `json:"chooseleaf_descend_once"`
			ChooseleafVaryR          int    `json:"chooseleaf_vary_r"`
			HasV2Rules               int    `json:"has_v2_rules"`
			HasV3Rules               int    `json:"has_v3_rules"`
			HasV4Buckets             int    `json:"has_v4_buckets"`
			LegacyTunables           int    `json:"legacy_tunables"`
			OptimalTunables          int    `json:"optimal_tunables"`
			Profile                  string `json:"profile"`
			RequireFeatureTunables   int    `json:"require_feature_tunables"`
			RequireFeatureTunables2  int    `json:"require_feature_tunables2"`
			RequireFeatureTunables3  int    `json:"require_feature_tunables3"`
			StrawCalcVersion         int    `json:"straw_calc_version"`
		} `json:"tunables"`
		Types []struct {
			Name   string `json:"name"`
			TypeID int    `json:"type_id"`
		} `json:"types"`
	} `json:"crushmap"`
	Health struct {
		Detail []interface{} `json:"detail"`
		Health struct {
			HealthServices []struct {
				Mons []struct {
					AvailPercent int    `json:"avail_percent"`
					Health       string `json:"health"`
					KbAvail      int    `json:"kb_avail"`
					KbTotal      int    `json:"kb_total"`
					KbUsed       int    `json:"kb_used"`
					LastUpdated  string `json:"last_updated"`
					Name         string `json:"name"`
					StoreStats   struct {
						BytesLog    int    `json:"bytes_log"`
						BytesMisc   int    `json:"bytes_misc"`
						BytesSst    int    `json:"bytes_sst"`
						BytesTotal  int    `json:"bytes_total"`
						LastUpdated string `json:"last_updated"`
					} `json:"store_stats"`
				} `json:"mons"`
			} `json:"health_services"`
		} `json:"health"`
		OverallStatus string        `json:"overall_status"`
		Summary       []interface{} `json:"summary"`
		Timechecks    struct {
			Epoch       int    `json:"epoch"`
			Round       int    `json:"round"`
			RoundStatus string `json:"round_status"`
		} `json:"timechecks"`
	} `json:"health"`
	Mdsmap struct {
		Compat struct {
			Compat   struct{} `json:"compat"`
			Incompat struct{} `json:"incompat"`
			RoCompat struct{} `json:"ro_compat"`
		} `json:"compat"`
		Created             string        `json:"created"`
		DataPools           []interface{} `json:"data_pools"`
		Enabled             bool          `json:"enabled"`
		Epoch               int           `json:"epoch"`
		Failed              []interface{} `json:"failed"`
		Flags               int           `json:"flags"`
		FsName              string        `json:"fs_name"`
		In                  []interface{} `json:"in"`
		Info                struct{}      `json:"info"`
		LastFailure         int           `json:"last_failure"`
		LastFailureOsdEpoch int           `json:"last_failure_osd_epoch"`
		MaxFileSize         int           `json:"max_file_size"`
		MaxMds              int           `json:"max_mds"`
		MetadataPool        int           `json:"metadata_pool"`
		Modified            string        `json:"modified"`
		Root                int           `json:"root"`
		SessionAutoclose    int           `json:"session_autoclose"`
		SessionTimeout      int           `json:"session_timeout"`
		Stopped             []interface{} `json:"stopped"`
		Tableserver         int           `json:"tableserver"`
		Up                  struct{}      `json:"up"`
	} `json:"mdsmap"`
	MdsmapFirstCommitted int `json:"mdsmap_first_committed"`
	MdsmapLastCommitted  int `json:"mdsmap_last_committed"`
	Monmap               struct {
		Created  string `json:"created"`
		Epoch    int    `json:"epoch"`
		Fsid     string `json:"fsid"`
		Modified string `json:"modified"`
		Mons     []struct {
			Addr string `json:"addr"`
			Name string `json:"name"`
			Rank int    `json:"rank"`
		} `json:"mons"`
	} `json:"monmap"`
	MonmapFirstCommitted int `json:"monmap_first_committed"`
	MonmapLastCommitted  int `json:"monmap_last_committed"`
	OsdMetadata          []struct {
		Arch              string `json:"arch"`
		BackAddr          string `json:"back_addr"`
		CephVersion       string `json:"ceph_version"`
		CPU               string `json:"cpu"`
		Distro            string `json:"distro"`
		DistroCodename    string `json:"distro_codename"`
		DistroDescription string `json:"distro_description"`
		DistroVersion     string `json:"distro_version"`
		FilestoreBackend  string `json:"filestore_backend"`
		FilestoreFType    string `json:"filestore_f_type"`
		FrontAddr         string `json:"front_addr"`
		HbBackAddr        string `json:"hb_back_addr"`
		HbFrontAddr       string `json:"hb_front_addr"`
		Hostname          string `json:"hostname"`
		ID                int    `json:"id"`
		KernelDescription string `json:"kernel_description"`
		KernelVersion     string `json:"kernel_version"`
		MemSwapKb         string `json:"mem_swap_kb"`
		MemTotalKb        string `json:"mem_total_kb"`
		Os                string `json:"os"`
		OsdData           string `json:"osd_data"`
		OsdJournal        string `json:"osd_journal"`
		OsdObjectstore    string `json:"osd_objectstore"`
	} `json:"osd_metadata"`
	Osdmap struct {
		Blacklist           []interface{} `json:"blacklist"`
		ClusterSnapshot     string        `json:"cluster_snapshot"`
		Created             string        `json:"created"`
		Epoch               int           `json:"epoch"`
		ErasureCodeProfiles struct {
			Default struct {
				Directory string `json:"directory"`
				K         string `json:"k"`
				M         string `json:"m"`
				Plugin    string `json:"plugin"`
				Technique string `json:"technique"`
			} `json:"default"`
		} `json:"erasure_code_profiles"`
		Flags    string `json:"flags"`
		Fsid     string `json:"fsid"`
		MaxOsd   int    `json:"max_osd"`
		Modified string `json:"modified"`
		OsdXinfo []struct {
			DownStamp        string  `json:"down_stamp"`
			Features         int     `json:"features"`
			LaggyInterval    int     `json:"laggy_interval"`
			LaggyProbability float64 `json:"laggy_probability"`
			OldWeight        int     `json:"old_weight"`
			Osd              int     `json:"osd"`
		} `json:"osd_xinfo"`
		Osds []struct {
			ClusterAddr        string   `json:"cluster_addr"`
			DownAt             int      `json:"down_at"`
			HeartbeatBackAddr  string   `json:"heartbeat_back_addr"`
			HeartbeatFrontAddr string   `json:"heartbeat_front_addr"`
			In                 int      `json:"in"`
			LastCleanBegin     int      `json:"last_clean_begin"`
			LastCleanEnd       int      `json:"last_clean_end"`
			LostAt             int      `json:"lost_at"`
			Osd                int      `json:"osd"`
			PrimaryAffinity    int      `json:"primary_affinity"`
			PublicAddr         string   `json:"public_addr"`
			State              []string `json:"state"`
			Up                 int      `json:"up"`
			UpFrom             int      `json:"up_from"`
			UpThru             int      `json:"up_thru"`
			UUID               string   `json:"uuid"`
			Weight             int      `json:"weight"`
		} `json:"osds"`
		PgTemp  []interface{} `json:"pg_temp"`
		PoolMax int           `json:"pool_max"`
		Pools   []struct {
			Auid                       int    `json:"auid"`
			CacheMinEvictAge           int    `json:"cache_min_evict_age"`
			CacheMinFlushAge           int    `json:"cache_min_flush_age"`
			CacheMode                  string `json:"cache_mode"`
			CacheTargetDirtyRatioMicro int    `json:"cache_target_dirty_ratio_micro"`
			CacheTargetFullRatioMicro  int    `json:"cache_target_full_ratio_micro"`
			CrashReplayInterval        int    `json:"crash_replay_interval"`
			CrushRuleset               int    `json:"crush_ruleset"`
			ErasureCodeProfile         string `json:"erasure_code_profile"`
			ExpectedNumObjects         int    `json:"expected_num_objects"`
			Flags                      int    `json:"flags"`
			FlagsNames                 string `json:"flags_names"`
			HitSetCount                int    `json:"hit_set_count"`
			HitSetParams               struct {
				Type string `json:"type"`
			} `json:"hit_set_params"`
			HitSetPeriod             int           `json:"hit_set_period"`
			LastChange               string        `json:"last_change"`
			LastForceOpResend        string        `json:"last_force_op_resend"`
			MinReadRecencyForPromote int           `json:"min_read_recency_for_promote"`
			MinSize                  int           `json:"min_size"`
			ObjectHash               int           `json:"object_hash"`
			PgNum                    int           `json:"pg_num"`
			PgPlacementNum           int           `json:"pg_placement_num"`
			Pool                     int           `json:"pool"`
			PoolName                 string        `json:"pool_name"`
			PoolSnaps                []interface{} `json:"pool_snaps"`
			QuotaMaxBytes            int           `json:"quota_max_bytes"`
			QuotaMaxObjects          int           `json:"quota_max_objects"`
			ReadTier                 int           `json:"read_tier"`
			RemovedSnaps             string        `json:"removed_snaps"`
			Size                     int           `json:"size"`
			SnapEpoch                int           `json:"snap_epoch"`
			SnapMode                 string        `json:"snap_mode"`
			SnapSeq                  int           `json:"snap_seq"`
			StripeWidth              int           `json:"stripe_width"`
			TargetMaxBytes           int           `json:"target_max_bytes"`
			TargetMaxObjects         int           `json:"target_max_objects"`
			TierOf                   int           `json:"tier_of"`
			Tiers                    []interface{} `json:"tiers"`
			Type                     int           `json:"type"`
			WriteTier                int           `json:"write_tier"`
		} `json:"pools"`
		PrimaryTemp []interface{} `json:"primary_temp"`
	} `json:"osdmap"`
	OsdmapFirstCommitted int `json:"osdmap_first_committed"`
	OsdmapLastCommitted  int `json:"osdmap_last_committed"`
	Paxos                struct {
		AcceptedPn     int `json:"accepted_pn"`
		FirstCommitted int `json:"first_committed"`
		LastCommitted  int `json:"last_committed"`
		LastPn         int `json:"last_pn"`
	} `json:"paxos"`
	Pgmap struct {
		FullRatio       float64 `json:"full_ratio"`
		LastOsdmapEpoch int     `json:"last_osdmap_epoch"`
		LastPgScan      int     `json:"last_pg_scan"`
		NearFullRatio   float64 `json:"near_full_ratio"`
		OsdStats        []struct {
			FsPerfStat struct {
				ApplyLatencyMs  int `json:"apply_latency_ms"`
				CommitLatencyMs int `json:"commit_latency_ms"`
			} `json:"fs_perf_stat"`
			HbIn            []int         `json:"hb_in"`
			HbOut           []interface{} `json:"hb_out"`
			Kb              int           `json:"kb"`
			KbAvail         int           `json:"kb_avail"`
			KbUsed          int           `json:"kb_used"`
			NumSnapTrimming int           `json:"num_snap_trimming"`
			OpQueueAgeHist  struct {
				Histogram  []interface{} `json:"histogram"`
				UpperBound int           `json:"upper_bound"`
			} `json:"op_queue_age_hist"`
			Osd              int `json:"osd"`
			SnapTrimQueueLen int `json:"snap_trim_queue_len"`
		} `json:"osd_stats"`
		OsdStatsSum struct {
			FsPerfStat struct {
				ApplyLatencyMs  int `json:"apply_latency_ms"`
				CommitLatencyMs int `json:"commit_latency_ms"`
			} `json:"fs_perf_stat"`
			HbIn            []interface{} `json:"hb_in"`
			HbOut           []interface{} `json:"hb_out"`
			Kb              int           `json:"kb"`
			KbAvail         int           `json:"kb_avail"`
			KbUsed          int           `json:"kb_used"`
			NumSnapTrimming int           `json:"num_snap_trimming"`
			OpQueueAgeHist  struct {
				Histogram  []interface{} `json:"histogram"`
				UpperBound int           `json:"upper_bound"`
			} `json:"op_queue_age_hist"`
			SnapTrimQueueLen int `json:"snap_trim_queue_len"`
		} `json:"osd_stats_sum"`
		PgStats []struct {
			Acting              []int         `json:"acting"`
			ActingPrimary       int           `json:"acting_primary"`
			BlockedBy           []interface{} `json:"blocked_by"`
			Created             int           `json:"created"`
			LastActive          string        `json:"last_active"`
			LastBecameActive    string        `json:"last_became_active"`
			LastBecamePeered    string        `json:"last_became_peered"`
			LastChange          string        `json:"last_change"`
			LastClean           string        `json:"last_clean"`
			LastCleanScrubStamp string        `json:"last_clean_scrub_stamp"`
			LastDeepScrub       string        `json:"last_deep_scrub"`
			LastDeepScrubStamp  string        `json:"last_deep_scrub_stamp"`
			LastEpochClean      int           `json:"last_epoch_clean"`
			LastFresh           string        `json:"last_fresh"`
			LastFullsized       string        `json:"last_fullsized"`
			LastPeered          string        `json:"last_peered"`
			LastScrub           string        `json:"last_scrub"`
			LastScrubStamp      string        `json:"last_scrub_stamp"`
			LastUndegraded      string        `json:"last_undegraded"`
			LastUnstale         string        `json:"last_unstale"`
			LogSize             int           `json:"log_size"`
			LogStart            string        `json:"log_start"`
			MappingEpoch        int           `json:"mapping_epoch"`
			OndiskLogSize       int           `json:"ondisk_log_size"`
			OndiskLogStart      string        `json:"ondisk_log_start"`
			Parent              string        `json:"parent"`
			ParentSplitBits     int           `json:"parent_split_bits"`
			Pgid                string        `json:"pgid"`
			ReportedEpoch       string        `json:"reported_epoch"`
			ReportedSeq         string        `json:"reported_seq"`
			StatSum             struct {
				NumBytes                   int `json:"num_bytes"`
				NumBytesHitSetArchive      int `json:"num_bytes_hit_set_archive"`
				NumBytesRecovered          int `json:"num_bytes_recovered"`
				NumDeepScrubErrors         int `json:"num_deep_scrub_errors"`
				NumKeysRecovered           int `json:"num_keys_recovered"`
				NumObjectClones            int `json:"num_object_clones"`
				NumObjectCopies            int `json:"num_object_copies"`
				NumObjects                 int `json:"num_objects"`
				NumObjectsDegraded         int `json:"num_objects_degraded"`
				NumObjectsDirty            int `json:"num_objects_dirty"`
				NumObjectsHitSetArchive    int `json:"num_objects_hit_set_archive"`
				NumObjectsMisplaced        int `json:"num_objects_misplaced"`
				NumObjectsMissingOnPrimary int `json:"num_objects_missing_on_primary"`
				NumObjectsOmap             int `json:"num_objects_omap"`
				NumObjectsRecovered        int `json:"num_objects_recovered"`
				NumObjectsUnfound          int `json:"num_objects_unfound"`
				NumRead                    int `json:"num_read"`
				NumReadKb                  int `json:"num_read_kb"`
				NumScrubErrors             int `json:"num_scrub_errors"`
				NumShallowScrubErrors      int `json:"num_shallow_scrub_errors"`
				NumWhiteouts               int `json:"num_whiteouts"`
				NumWrite                   int `json:"num_write"`
				NumWriteKb                 int `json:"num_write_kb"`
			} `json:"stat_sum"`
			State        string `json:"state"`
			StatsInvalid string `json:"stats_invalid"`
			Up           []int  `json:"up"`
			UpPrimary    int    `json:"up_primary"`
			Version      string `json:"version"`
		} `json:"pg_stats"`
		PgStatsDelta struct {
			Acting        int `json:"acting"`
			LogSize       int `json:"log_size"`
			OndiskLogSize int `json:"ondisk_log_size"`
			StatSum       struct {
				NumBytes                   int `json:"num_bytes"`
				NumBytesHitSetArchive      int `json:"num_bytes_hit_set_archive"`
				NumBytesRecovered          int `json:"num_bytes_recovered"`
				NumDeepScrubErrors         int `json:"num_deep_scrub_errors"`
				NumKeysRecovered           int `json:"num_keys_recovered"`
				NumObjectClones            int `json:"num_object_clones"`
				NumObjectCopies            int `json:"num_object_copies"`
				NumObjects                 int `json:"num_objects"`
				NumObjectsDegraded         int `json:"num_objects_degraded"`
				NumObjectsDirty            int `json:"num_objects_dirty"`
				NumObjectsHitSetArchive    int `json:"num_objects_hit_set_archive"`
				NumObjectsMisplaced        int `json:"num_objects_misplaced"`
				NumObjectsMissingOnPrimary int `json:"num_objects_missing_on_primary"`
				NumObjectsOmap             int `json:"num_objects_omap"`
				NumObjectsRecovered        int `json:"num_objects_recovered"`
				NumObjectsUnfound          int `json:"num_objects_unfound"`
				NumRead                    int `json:"num_read"`
				NumReadKb                  int `json:"num_read_kb"`
				NumScrubErrors             int `json:"num_scrub_errors"`
				NumShallowScrubErrors      int `json:"num_shallow_scrub_errors"`
				NumWhiteouts               int `json:"num_whiteouts"`
				NumWrite                   int `json:"num_write"`
				NumWriteKb                 int `json:"num_write_kb"`
			} `json:"stat_sum"`
			Up int `json:"up"`
		} `json:"pg_stats_delta"`
		PgStatsSum struct {
			Acting        int `json:"acting"`
			LogSize       int `json:"log_size"`
			OndiskLogSize int `json:"ondisk_log_size"`
			StatSum       struct {
				NumBytes                   int `json:"num_bytes"`
				NumBytesHitSetArchive      int `json:"num_bytes_hit_set_archive"`
				NumBytesRecovered          int `json:"num_bytes_recovered"`
				NumDeepScrubErrors         int `json:"num_deep_scrub_errors"`
				NumKeysRecovered           int `json:"num_keys_recovered"`
				NumObjectClones            int `json:"num_object_clones"`
				NumObjectCopies            int `json:"num_object_copies"`
				NumObjects                 int `json:"num_objects"`
				NumObjectsDegraded         int `json:"num_objects_degraded"`
				NumObjectsDirty            int `json:"num_objects_dirty"`
				NumObjectsHitSetArchive    int `json:"num_objects_hit_set_archive"`
				NumObjectsMisplaced        int `json:"num_objects_misplaced"`
				NumObjectsMissingOnPrimary int `json:"num_objects_missing_on_primary"`
				NumObjectsOmap             int `json:"num_objects_omap"`
				NumObjectsRecovered        int `json:"num_objects_recovered"`
				NumObjectsUnfound          int `json:"num_objects_unfound"`
				NumRead                    int `json:"num_read"`
				NumReadKb                  int `json:"num_read_kb"`
				NumScrubErrors             int `json:"num_scrub_errors"`
				NumShallowScrubErrors      int `json:"num_shallow_scrub_errors"`
				NumWhiteouts               int `json:"num_whiteouts"`
				NumWrite                   int `json:"num_write"`
				NumWriteKb                 int `json:"num_write_kb"`
			} `json:"stat_sum"`
			Up int `json:"up"`
		} `json:"pg_stats_sum"`
		PoolStats []struct {
			Acting        int `json:"acting"`
			LogSize       int `json:"log_size"`
			OndiskLogSize int `json:"ondisk_log_size"`
			Poolid        int `json:"poolid"`
			StatSum       struct {
				NumBytes                   int `json:"num_bytes"`
				NumBytesHitSetArchive      int `json:"num_bytes_hit_set_archive"`
				NumBytesRecovered          int `json:"num_bytes_recovered"`
				NumDeepScrubErrors         int `json:"num_deep_scrub_errors"`
				NumKeysRecovered           int `json:"num_keys_recovered"`
				NumObjectClones            int `json:"num_object_clones"`
				NumObjectCopies            int `json:"num_object_copies"`
				NumObjects                 int `json:"num_objects"`
				NumObjectsDegraded         int `json:"num_objects_degraded"`
				NumObjectsDirty            int `json:"num_objects_dirty"`
				NumObjectsHitSetArchive    int `json:"num_objects_hit_set_archive"`
				NumObjectsMisplaced        int `json:"num_objects_misplaced"`
				NumObjectsMissingOnPrimary int `json:"num_objects_missing_on_primary"`
				NumObjectsOmap             int `json:"num_objects_omap"`
				NumObjectsRecovered        int `json:"num_objects_recovered"`
				NumObjectsUnfound          int `json:"num_objects_unfound"`
				NumRead                    int `json:"num_read"`
				NumReadKb                  int `json:"num_read_kb"`
				NumScrubErrors             int `json:"num_scrub_errors"`
				NumShallowScrubErrors      int `json:"num_shallow_scrub_errors"`
				NumWhiteouts               int `json:"num_whiteouts"`
				NumWrite                   int `json:"num_write"`
				NumWriteKb                 int `json:"num_write_kb"`
			} `json:"stat_sum"`
			Up int `json:"up"`
		} `json:"pool_stats"`
		Stamp   string `json:"stamp"`
		Version int    `json:"version"`
	} `json:"pgmap"`
	PgmapFirstCommitted int    `json:"pgmap_first_committed"`
	PgmapLastCommitted  int    `json:"pgmap_last_committed"`
	Quorum              []int  `json:"quorum"`
	Tag                 string `json:"tag"`
	Timestamp           string `json:"timestamp"`
	Version             string `json:"version"`
}
