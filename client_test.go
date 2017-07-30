package telegraf

import (
	"fmt"
	"net"
	"testing"
)

func BenchmarkUDPWriteRaw(b *testing.B) {
	conn, err := net.Dial("udp", "127.0.0.1:8094")
	if err != nil {
		b.Fatal("failed to connect: %s", err)
	}
	defer conn.Close()
	text := "weather,location=us-midwest temperature=82\n"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, err := fmt.Fprintf(conn, text); err != nil {
			b.Fatal("failed to write: %s", err)
		}
	}
}
