package monitoring

type RedisKeyInfo struct {
	KeyName string
	KeyMetrics []MetricDetails
}
type MetricDetails struct {
	MetricName string
	MetricDescription string
	MetricValue string
}

func GetMetricDescription() map[string]string {

	return map[string]string{
		//memory
		"used_memory": "Total number of bytes allocated by Redis using its allocator (either standard libc, jemalloc, or an alternative allocator such as tcmalloc)",
		"used_memory_human": "Human readable representation of previous value",
		"used_memory_rss": "Number of bytes that Redis allocated as seen by the operating system (a.k.a resident set size). This is the number reported by tools such as top(1) and ps(1)",
		"used_memory_rss_human": "Human readable representation of previous value",
		"used_memory_peak": "Peak memory consumed by Redis (in bytes)",
		"used_memory_peak_human": "Human readable representation of previous value",
		"used_memory_peak_perc": "The percentage of used_memory_peak out of used_memory",
		"used_memory_overhead": "The sum in bytes of all overheads that the server allocated for managing its internal data structures",
		"used_memory_startup": "Initial amount of memory consumed by Redis at startup in bytes",
		"used_memory_dataset": "The size in bytes of the dataset (used_memory_overhead subtracted from used_memory)",
		"used_memory_dataset_perc": "The percentage of used_memory_dataset out of the net memory usage (used_memory minus used_memory_startup)",
		"total_system_memory": "The total amount of memory that the Redis host has",
		"total_system_memory_human": "Human readable representation of previous value",
		"used_memory_lua": "Number of bytes used by the Lua engine",
		"used_memory_lua_human": "Human readable representation of previous value",
		"used_memory_scripts": "Number of bytes used by cached Lua scripts",
		"used_memory_scripts_human": "Human readable representation of previous value",
		"maxmemory": "The value of the maxmemory configuration directive",
		"maxmemory_human": "Human readable representation of previous value",
		"maxmemory_policy": "The value of the maxmemory-policy configuration directive",
		"mem_fragmentation_ratio": "Ratio between used_memory_rss and used_memory",
		"mem_allocator": "Memory allocator, chosen at compile time",
		"active_defrag_running": "Flag indicating if active defragmentation is active",
		"lazyfree_pending_objects": "The number of objects waiting to be freed (as a result of calling UNLINK, or FLUSHDB and FLUSHALL with the ASYNC option)",

		//server
		"redis_version": "Version of the Redis server",
		"redis_git_sha1": "Git SHA1",
		"redis_git_dirty": "Git dirty flag",
		"redis_build_id": "The build id",
		"redis_mode": `The server's mode ("standalone", "sentinel" or "cluster")`,
		"os": "Operating system hosting the Redis server",
		"arch_bits": "Architecture (32 or 64 bits)",
		"multiplexing_api": "Event loop mechanism used by Redis",
		"atomicvar_api": "Atomicvar API used by Redis",
		"gcc_version": "Version of the GCC compiler used to compile the Redis server",
		"process_id": "PID of the server process",
		"run_id": "Random value identifying the Redis server (to be used by Sentinel and Cluster)",
		"tcp_port": "TCP/IP listen port",
		"server_time_in_usec": "Epoch-based system time with microsecond precision",
		"uptime_in_seconds": "Number of seconds since Redis server start",
		"uptime_in_days": "Same value expressed in days",
		"hz": "The server's current frequency setting",
		"configured_hz": "The server's configured frequency setting",
		"lru_clock": "Clock incrementing every minute, for LRU management",
		"executable": "The path to the server's executable",
		"config_file": "The path to the config file",

		//clent
		"connected_clients": "Number of client connections (excluding connections from replicas)",
		"cluster_connections": "An approximation of the number of sockets used by the cluster's bus",
		"maxclients": "The value of the maxclients configuration directive. This is the upper limit for the sum of connected_clients, connected_slaves and cluster_connections.",
		"client_longest_output_list": "Longest output list among current client connections",
		"client_biggest_input_buf": "Biggest input buffer among current client connections",
		"blocked_clients": "Number of clients pending on a blocking call (BLPOP, BRPOP, BRPOPLPUSH, BLMOVE, BZPOPMIN, BZPOPMAX)",
		"tracking_clients": "Number of clients being tracked (CLIENT TRACKING)",
		"clients_in_timeout_table": "Number of clients in the clients timeout table",
		"io_threads_active": "Flag indicating if I/O threads are active",

		//cpu
		"used_cpu_sys": "System CPU consumed by the Redis server, which is the sum of system CPU consumed by all threads of the server process (main thread and background threads)",
		"used_cpu_user": "User CPU consumed by the Redis server, which is the sum of user CPU consumed by all threads of the server process (main thread and background threads)",
		"used_cpu_sys_children": "System CPU consumed by the background processes",
		"used_cpu_user_children": "User CPU consumed by the background processes",
		"used_cpu_sys_main_thread": "System CPU consumed by the Redis server main thread",
		"used_cpu_user_main_thread": "User CPU consumed by the Redis server main thread",

		//cluster
		"cluster_enabled": "Indicate Redis cluster is enabled",
	}
}
