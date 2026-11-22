package rpc

import "testing"

func TestIsTransientTunErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{
			name: "leftover adapter both errors",
			msg:  "start service: start inbound: start inbound/tun[tun-in]: configure tun interface: (create adapter: Cannot create a file when that file already exists. | open existing adapter: Element not found.)",
			want: true,
		},
		{
			name: "already exists only",
			msg:  "start inbound/tun[tun-in]: configure tun interface: Cannot create a file when that file already exists.",
			want: true,
		},
		{
			name: "file not found",
			msg:  "start service: start inbound: start inbound/tun[tun-in]: configure tun interface: The system cannot find the file specified.",
			want: true,
		},
		{
			name: "access denied should not retry",
			msg:  "start inbound/tun[tun-in]: configure tun interface: Access is denied.",
			want: false,
		},
		{
			name: "unrelated error",
			msg:  "start service: listen tcp 127.0.0.1:2080: bind: address already in use",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isTransientTunErrorMessage(tc.msg); got != tc.want {
				t.Fatalf("isTransientTunErrorMessage(%q) = %v, want %v", tc.msg, got, tc.want)
			}
		})
	}
}
