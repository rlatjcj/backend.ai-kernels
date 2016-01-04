package policy

var TracedSyscalls []string
var AllowedSyscalls []string

func init() {
	// Following syscalls are intercepted by our ptrace-based tracer.
	// The tracer will implement its own policies, optinally by inspecting
	// the arguments in the registers.
	TracedSyscalls = []string{
		"open",
		"stat",
		"lstat",
		"statfs",
		"access",
		"readlink",
		"creat",
		"rename",
		"unlink",
		// traced by ptrace exec/fork/clone
		"fork",
		"vfork",
		"clone",
		"execve",
	}

	// Following syscalls are blindly allowed.
	// IMPORTANT: ptrace MUST NOT be included!
	AllowedSyscalls = []string{
		// blindly allowed
		"read",
		"write",
		"close",
		"openat",
		"fstat",
		"fstatfs",
		"mmap",
		"mprotect",
		"munmap",
		"brk",
		"lseek",
		"getdents",
		"dup",
		"dup2",
		"chdir",
		"fchdir",
		"tkill",
		"tgkill",
		"rt_sigaction",
		"rt_sigprocmask",
		"sigaltstack",
		"arch_prctl",
		"prctl",
		"getrlimit",
		"set_tid_address",
		"clear_tid_address",
		"set_robust_list",
		"get_robust_list",
		"futex",
		"sched_getaffinity",
		"sched_setaffinity",
		"sched_getparam",
		"sched_getscheduler",
		"sched_setscheduler",
		"eventfd",
		"eventfd2",
		"setsockopt",
		"getsockname",
		"bind",
		"listen",
		"gettid",
		"getuid",
		"geteuid",
		"getgid",
		"getegid",
		"getcwd",
		"socket",
		"socketpair",
		"connect",
		"pipe",
		"ioctl",
		"fcntl",
		"select",
		"poll",
		"epoll_create",
		"epoll_create1",
		"epoll_wait",
		"epoll_ctl",
		"exit_group",
		"wait4",
		"uname",
		// potentially replaced with VDSO
		"getpid",
		"getppid",
		"getpgid",
		"getpgrp",
		"getsid",
		"clock_gettime",
		"clock_getres",
		"clock_nanosleep",
	}
}

// vim: ts=4 sts=4 sw=4 noet
