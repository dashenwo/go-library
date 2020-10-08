package identifiers

import (
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
	"unsafe"
)

var (
	sequence      Uint32                                   // Sequence for unique purpose of current process.
	sequenceMax   = uint32(46655)                          // Sequence max("zzz").
	randomStrBase = "0123456789abcdefghijklmnopqrstuvwxyz" // Random chars string(36 bytes).
	macAddrStr    = "0000000"                              // MAC addresses hash result in 7 bytes.
	processIdStr  = "0000"                                 // Process id in 4 bytes.
	DefaultIdentifiers = new(uuid)
)

type uuid struct {
	opts Options
}
// 初始化方法
func init()  {
	// MAC addresses hash result in 7 bytes.
	macs, _ := getMacArray()
	if len(macs) > 0 {
		var macAddrBytes []byte
		for _, mac := range macs {
			macAddrBytes = append(macAddrBytes, []byte(mac)...)
		}
		b := []byte{'0', '0', '0', '0', '0', '0', '0'}
		s := strconv.FormatUint(uint64(hash(macAddrBytes)), 36)
		copy(b, s)
		macAddrStr = string(b)
	}
	// Process id in 4 bytes.
	{
		b := []byte{'0', '0', '0', '0'}
		s := strconv.FormatInt(int64(os.Getpid()), 36)
		copy(b, s)
		processIdStr = string(b)
	}
}

func NewUuid(opts ...Option) Identifiers {
	options := Options{
	}
	for _, o := range opts {
		o(&options)
	}
	return &uuid{
		opts: options,
	}
}

func (r *uuid) Init(opts ...Option) error {
	for _, o := range opts {
		o(&r.opts)
	}
	return nil
}

func (r *uuid) Generate() string {
	return generate()
}

func generate(data ...[]byte) string {
	var (
		b       = make([]byte, 32)
		nanoStr = strconv.FormatInt(time.Now().UnixNano(), 36)
	)
	if len(data) == 0 {
		copy(b, macAddrStr)
		copy(b[7:], processIdStr)
		copy(b[11:], nanoStr)
		copy(b[23:], getSequence())
		copy(b[26:], getRandomStr(6))
	} else if len(data) <= 2 {
		n := 0
		for i, v := range data {
			// Ignore empty data item bytes.
			if len(v) > 0 {
				copy(b[i*7:], getDataHashStr(v))
				n += 7
			}
		}
		copy(b[n:], nanoStr)
		copy(b[n+12:], getSequence())
		copy(b[n+12+3:], getRandomStr(32-n-12-3))
	} else {
		panic("data count too long, no more than 2")
	}
	return unsafeBytesToStr(b)
}

func getSequence() []byte {
	b := []byte{'0', '0', '0'}
	s := strconv.FormatUint(uint64(sequence.Add(1)%sequenceMax), 36)
	copy(b, s)
	return b
}

func getRandomStr(n int) []byte {
	bytes := []byte(randomStrBase)
	result := []byte{}
	rm := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[rm.Intn(len(bytes))])
	}
	return result
}

func getDataHashStr(data []byte) []byte {
	b := []byte{'0', '0', '0', '0', '0', '0', '0'}
	s := strconv.FormatUint(uint64(hash(data)), 36)
	copy(b, s)
	return b
}

func getMacArray() (macs []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macs = append(macs, macAddr)
	}
	return macs, nil
}

func unsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func hash(str []byte) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint32(str[i])
	}
	return hash
}