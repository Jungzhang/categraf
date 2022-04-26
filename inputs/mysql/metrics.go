package mysql

var STATUS_VARS = map[string]struct{}{
	"uptime":                     {},
	"prepared_stmt_count":        {}, // command metrics
	"slow_queries":               {},
	"questions":                  {},
	"queries":                    {},
	"com_select":                 {},
	"com_insert":                 {},
	"com_update":                 {},
	"com_delete":                 {},
	"com_replace":                {},
	"com_load":                   {},
	"com_insert_select":          {},
	"com_update_multi":           {},
	"com_delete_multi":           {},
	"com_replace_select":         {},
	"connections":                {}, // connection metrics
	"max_used_connections":       {},
	"aborted_clients":            {},
	"aborted_connects":           {},
	"open_files":                 {}, // table cache metrics
	"open_tables":                {},
	"bytes_sent":                 {}, // network metrics
	"bytes_received":             {},
	"qcache_hits":                {}, // query cache metrics
	"qcache_inserts":             {},
	"qcache_lowmem_prunes":       {},
	"table_locks_waited":         {}, // table lock metrics
	"table_locks_waited_rate":    {},
	"created_tmp_tables":         {}, // temporary table metrics
	"created_tmp_disk_tables":    {},
	"created_tmp_files":          {},
	"threads_connected":          {}, // thread metrics
	"threads_running":            {},
	"key_buffer_bytes_unflushed": {}, // myisam metrics
	"key_buffer_bytes_used":      {},
	"key_read_requests":          {},
	"key_reads":                  {},
	"key_write_requests":         {},
	"key_writes":                 {},
}

var VARIABLES_VARS = map[string]struct{}{
	"key_buffer_size":         {},
	"key_cache_utilization":   {},
	"max_connections":         {},
	"max_prepared_stmt_count": {},
	"query_cache_size":        {},
	"table_open_cache":        {},
	"thread_cache_size":       {},
}

// auto compute
// "innodb_buffer_pool_bytes_dirty":       {},
// "innodb_buffer_pool_bytes_free":        {},
// "innodb_buffer_pool_bytes_used":        {},
// "innodb_buffer_pool_bytes_total":       {},
// "innodb_buffer_pool_pages_utilization": {},
var INNODB_VARS = map[string]struct{}{
	"innodb_buffer_pool_size":          {},
	"open_files_limit":                 {},
	"innodb_log_waits":                 {},
	"innodb_data_reads":                {},
	"innodb_data_writes":               {},
	"innodb_os_log_fsyncs":             {},
	"innodb_mutex_spin_waits":          {},
	"innodb_mutex_spin_rounds":         {},
	"innodb_mutex_os_waits":            {},
	"innodb_row_lock_waits":            {},
	"innodb_row_lock_time":             {},
	"innodb_row_lock_current_waits":    {},
	"innodb_current_row_locks":         {},
	"innodb_buffer_pool_read_requests": {},
	"innodb_buffer_pool_reads":         {},
}

var BINLOG_VARS = map[string]struct{}{
	"binlog_space_usage_bytes": {},
}

var OPTIONAL_STATUS_VARS = map[string]struct{}{
	"binlog_cache_disk_use":      {},
	"binlog_cache_use":           {},
	"handler_commit":             {},
	"handler_delete":             {},
	"handler_prepare":            {},
	"handler_read_first":         {},
	"handler_read_key":           {},
	"handler_read_next":          {},
	"handler_read_prev":          {},
	"handler_read_rnd":           {},
	"handler_read_rnd_next":      {},
	"handler_rollback":           {},
	"handler_update":             {},
	"handler_write":              {},
	"opened_tables":              {},
	"qcache_total_blocks":        {},
	"qcache_free_blocks":         {},
	"qcache_free_memory":         {},
	"qcache_not_cached":          {},
	"qcache_queries_in_cache":    {},
	"select_full_join":           {},
	"select_full_range_join":     {},
	"select_range":               {},
	"select_range_check":         {},
	"select_scan":                {},
	"sort_merge_passes":          {},
	"sort_range":                 {},
	"sort_rows":                  {},
	"sort_scan":                  {},
	"table_locks_immediate":      {},
	"table_locks_immediate_rate": {},
	"threads_cached":             {},
	"threads_created":            {},
	"table_open_cache_hits":      {}, // Status Vars added in Mysql 5.6.6
	"table_open_cache_misses":    {},
}

var OPTIONAL_INNODB_VARS = map[string]struct{}{
	"innodb_active_transactions":            {},
	"innodb_buffer_pool_bytes_data":         {},
	"innodb_buffer_pool_pages_data":         {},
	"innodb_buffer_pool_pages_dirty":        {},
	"innodb_buffer_pool_pages_flushed":      {},
	"innodb_buffer_pool_pages_free":         {},
	"innodb_buffer_pool_pages_total":        {},
	"innodb_buffer_pool_read_ahead":         {},
	"innodb_buffer_pool_read_ahead_evicted": {},
	"innodb_buffer_pool_read_ahead_rnd":     {},
	"innodb_buffer_pool_wait_free":          {},
	"innodb_buffer_pool_write_requests":     {},
	"innodb_checkpoint_age":                 {},
	"innodb_current_transactions":           {},
	"innodb_data_fsyncs":                    {},
	"innodb_data_pending_fsyncs":            {},
	"innodb_data_pending_reads":             {},
	"innodb_data_pending_writes":            {},
	"innodb_data_read":                      {},
	"innodb_data_written":                   {},
	"innodb_dblwr_pages_written":            {},
	"innodb_dblwr_writes":                   {},
	"innodb_hash_index_cells_total":         {},
	"innodb_hash_index_cells_used":          {},
	"innodb_history_list_length":            {},
	"innodb_ibuf_free_list":                 {},
	"innodb_ibuf_merged":                    {},
	"innodb_ibuf_merged_delete_marks":       {},
	"innodb_ibuf_merged_deletes":            {},
	"innodb_ibuf_merged_inserts":            {},
	"innodb_ibuf_merges":                    {},
	"innodb_ibuf_segment_size":              {},
	"innodb_ibuf_size":                      {},
	"innodb_lock_structs":                   {},
	"innodb_locked_tables":                  {},
	"innodb_locked_transactions":            {},
	"innodb_log_write_requests":             {},
	"innodb_log_writes":                     {},
	"innodb_lsn_current":                    {},
	"innodb_lsn_flushed":                    {},
	"innodb_lsn_last_checkpoint":            {},
	"innodb_mem_adaptive_hash":              {},
	"innodb_mem_additional_pool":            {},
	"innodb_mem_dictionary":                 {},
	"innodb_mem_file_system":                {},
	"innodb_mem_lock_system":                {},
	"innodb_mem_page_hash":                  {},
	"innodb_mem_recovery_system":            {},
	"innodb_mem_thread_hash":                {},
	"innodb_mem_total":                      {},
	"innodb_os_file_fsyncs":                 {},
	"innodb_os_file_reads":                  {},
	"innodb_os_file_writes":                 {},
	"innodb_os_log_pending_fsyncs":          {},
	"innodb_os_log_pending_writes":          {},
	"innodb_os_log_written":                 {},
	"innodb_pages_created":                  {},
	"innodb_pages_read":                     {},
	"innodb_pages_written":                  {},
	"innodb_pending_aio_log_ios":            {},
	"innodb_pending_aio_sync_ios":           {},
	"innodb_pending_buffer_pool_flushes":    {},
	"innodb_pending_checkpoint_writes":      {},
	"innodb_pending_ibuf_aio_reads":         {},
	"innodb_pending_log_flushes":            {},
	"innodb_pending_log_writes":             {},
	"innodb_pending_normal_aio_reads":       {},
	"innodb_pending_normal_aio_writes":      {},
	"innodb_queries_inside":                 {},
	"innodb_queries_queued":                 {},
	"innodb_read_views":                     {},
	"innodb_rows_deleted":                   {},
	"innodb_rows_inserted":                  {},
	"innodb_rows_read":                      {},
	"innodb_rows_updated":                   {},
	"innodb_s_lock_os_waits":                {},
	"innodb_s_lock_spin_rounds":             {},
	"innodb_s_lock_spin_waits":              {},
	"innodb_semaphore_wait_time":            {},
	"innodb_semaphore_waits":                {},
	"innodb_tables_in_use":                  {},
	"innodb_x_lock_os_waits":                {},
	"innodb_x_lock_spin_rounds":             {},
	"innodb_x_lock_spin_waits":              {},
}

var GALERA_VARS = map[string]struct{}{
	"wsrep_cluster_size":           {},
	"wsrep_local_recv_queue_avg":   {},
	"wsrep_flow_control_paused":    {},
	"wsrep_flow_control_paused_ns": {},
	"wsrep_flow_control_recv":      {},
	"wsrep_flow_control_sent":      {},
	"wsrep_cert_deps_distance":     {},
	"wsrep_local_send_queue_avg":   {},
	"wsrep_replicated_bytes":       {},
	"wsrep_received_bytes":         {},
	"wsrep_received":               {},
	"wsrep_local_state":            {},
	"wsrep_local_cert_failures":    {},
}

var PERFORMANCE_VARS = map[string]struct{}{
	"query_run_time_avg":                 {},
	"perf_digest_95th_percentile_avg_us": {},
}

var SCHEMA_VARS = map[string]struct{}{
	"information_schema_size": {},
}

var TABLE_VARS = map[string]struct{}{
	"information_table_index_size": {},
	"information_table_data_size":  {},
}

var REPLICA_VARS = map[string]struct{}{
	"seconds_behind_source": {},
	"seconds_behind_master": {},
	"replicas_connected":    {},
}

var GROUP_REPLICATION_VARS = map[string]struct{}{
	"transactions_count":                {},
	"transactions_check":                {},
	"conflict_detected":                 {},
	"transactions_row_validating":       {},
	"transactions_remote_applier_queue": {},
	"transactions_remote_applied":       {},
	"transactions_local_proposed":       {},
	"transactions_local_rollback":       {},
}

var SYNTHETIC_VARS = map[string]struct{}{
	"qcache_utilization":         {},
	"qcache_instant_utilization": {},
}
