package encoding

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/grafana/tempo/tempodb/backend"
	"github.com/grafana/tempo/tempodb/backend/local"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV2Block(t *testing.T) {
	// known ids and objs written
	// ids
	ids := [][]byte{
		{0x0c, 0xdb, 0x24, 0xe0, 0x4e, 0x57, 0x8d, 0x7f, 0x37, 0x7c, 0x0c, 0xa5, 0x00, 0x0d, 0xf5, 0x58},
		{0x12, 0x4f, 0x2c, 0x9b, 0x97, 0xb9, 0x20, 0xe1, 0x17, 0x5c, 0xa7, 0x51, 0x04, 0x68, 0x2f, 0x02},
		{0x1c, 0xab, 0x29, 0x93, 0x6f, 0x47, 0xec, 0xb8, 0x0b, 0x1b, 0x07, 0xda, 0xf2, 0xf7, 0xcf, 0xc0},
		{0x1d, 0xc5, 0xac, 0x5f, 0x57, 0x72, 0xaf, 0xf0, 0xf3, 0x53, 0xb1, 0xc8, 0x3b, 0xb9, 0x27, 0x38},
		{0x24, 0x34, 0x83, 0x30, 0x58, 0xef, 0x73, 0xb1, 0xdb, 0xe9, 0xc8, 0xca, 0x2d, 0xc2, 0xd1, 0x5c},
		{0x5b, 0x36, 0x83, 0x13, 0x0b, 0x32, 0x34, 0x0c, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35},
		{0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x02, 0x0b, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9},
		{0x9b, 0xc2, 0x5c, 0x22, 0x81, 0x1b, 0x8b, 0xfd, 0x73, 0x0f, 0xfd, 0xa6, 0x94, 0xf7, 0x28, 0x37},
		{0xb4, 0xc7, 0x3c, 0x0e, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x07, 0x4e, 0xc8, 0x64},
		{0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x01, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc},
	}

	// objs
	objs := [][]byte{
		{0xa, 0xba, 0x1, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xc, 0xdb, 0x24, 0xe0, 0x4e, 0x57, 0x8d, 0x7f, 0x37, 0x7c, 0xc, 0xa5, 0x0, 0xd, 0xf5, 0x58, 0x12, 0x8, 0xdb, 0xae, 0xdd, 0xd0, 0x81, 0xfe, 0xb1, 0x9c, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xc, 0xdb, 0x24, 0xe0, 0x4e, 0x57, 0x8d, 0x7f, 0x37, 0x7c, 0xc, 0xa5, 0x0, 0xd, 0xf5, 0x58, 0x12, 0x8, 0xf, 0xf8, 0x88, 0xcd, 0xbe, 0x58, 0xca, 0xf5, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xc, 0xdb, 0x24, 0xe0, 0x4e, 0x57, 0x8d, 0x7f, 0x37, 0x7c, 0xc, 0xa5, 0x0, 0xd, 0xf5, 0x58, 0x12, 0x8, 0xb9, 0x4c, 0x9b, 0xd4, 0xb1, 0xcf, 0xa, 0x3d, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0x9c, 0x2, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x12, 0x4f, 0x2c, 0x9b, 0x97, 0xb9, 0x20, 0xe1, 0x17, 0x5c, 0xa7, 0x51, 0x4, 0x68, 0x2f, 0x2, 0x12, 0x8, 0xc5, 0x98, 0xb, 0xa3, 0x54, 0x84, 0x19, 0x2f, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x60, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x12, 0x4f, 0x2c, 0x9b, 0x97, 0xb9, 0x20, 0xe1, 0x17, 0x5c, 0xa7, 0x51, 0x4, 0x68, 0x2f, 0x2, 0x12, 0x8, 0xf5, 0x9c, 0x1, 0xe5, 0xfb, 0x89, 0xa6, 0x32, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x12, 0x4f, 0x2c, 0x9b, 0x97, 0xb9, 0x20, 0xe1, 0x17, 0x5c, 0xa7, 0x51, 0x4, 0x68, 0x2f, 0x2, 0x12, 0x8, 0x40, 0x83, 0xeb, 0xca, 0x49, 0x4f, 0x7, 0x28, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x12, 0x4f, 0x2c, 0x9b, 0x97, 0xb9, 0x20, 0xe1, 0x17, 0x5c, 0xa7, 0x51, 0x4, 0x68, 0x2f, 0x2, 0x12, 0x8, 0xc2, 0x75, 0xd5, 0xb0, 0xd4, 0x86, 0xba, 0x98, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x12, 0x4f, 0x2c, 0x9b, 0x97, 0xb9, 0x20, 0xe1, 0x17, 0x5c, 0xa7, 0x51, 0x4, 0x68, 0x2f, 0x2, 0x12, 0x8, 0xf3, 0xf5, 0xa1, 0x7b, 0x4c, 0xcf, 0xd0, 0xb1, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0x0},
		{0xa, 0x7c, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x1d, 0xc5, 0xac, 0x5f, 0x57, 0x72, 0xaf, 0xf0, 0xf3, 0x53, 0xb1, 0xc8, 0x3b, 0xb9, 0x27, 0x38, 0x12, 0x8, 0xe0, 0xc7, 0xd9, 0xe, 0xdb, 0x10, 0xd1, 0xe7, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x1d, 0xc5, 0xac, 0x5f, 0x57, 0x72, 0xaf, 0xf0, 0xf3, 0x53, 0xb1, 0xc8, 0x3b, 0xb9, 0x27, 0x38, 0x12, 0x8, 0x5d, 0x4a, 0xa5, 0x17, 0x8d, 0x72, 0x8e, 0x53, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0x62, 0x12, 0x60, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x24, 0x34, 0x83, 0x30, 0x58, 0xef, 0x73, 0xb1, 0xdb, 0xe9, 0xc8, 0xca, 0x2d, 0xc2, 0xd1, 0x5c, 0x12, 0x8, 0x61, 0xa1, 0xf8, 0xdb, 0xd7, 0x32, 0xb, 0xfb, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x24, 0x34, 0x83, 0x30, 0x58, 0xef, 0x73, 0xb1, 0xdb, 0xe9, 0xc8, 0xca, 0x2d, 0xc2, 0xd1, 0x5c, 0x12, 0x8, 0x4f, 0xd7, 0x25, 0x38, 0x2c, 0xcf, 0xfc, 0xf5, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0xa3, 0x3, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0x80, 0x65, 0x64, 0x92, 0x13, 0xf0, 0x93, 0xc6, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x84, 0x1, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0x6d, 0xeb, 0x5c, 0xab, 0xbd, 0x55, 0x23, 0x10, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0xd5, 0x74, 0xca, 0x3d, 0x2f, 0x53, 0xa1, 0x88, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0x23, 0xdd, 0xaf, 0x16, 0x8c, 0xb7, 0x80, 0x58, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0x7c, 0x58, 0x6e, 0x89, 0xed, 0x11, 0x20, 0x3d, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0x9b, 0x59, 0x80, 0x47, 0x49, 0x88, 0xe5, 0xc, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x60, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0x59, 0xed, 0xa6, 0x31, 0x9c, 0xd, 0xff, 0x2a, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x5b, 0x36, 0x83, 0x13, 0xb, 0x32, 0x34, 0xc, 0x9d, 0x8a, 0xa5, 0xe6, 0x1e, 0x43, 0xd5, 0x35, 0x12, 0x8, 0x61, 0x5f, 0x4b, 0x5c, 0x2e, 0x6f, 0xa7, 0x9e, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0xef, 0x2, 0x12, 0xcc, 0x1, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0x3c, 0x5b, 0x93, 0x57, 0xfe, 0xb0, 0x4, 0x9e, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0x31, 0xfd, 0x9b, 0x12, 0x50, 0x5b, 0xf8, 0xb0, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0x1b, 0xd0, 0xe, 0xea, 0x57, 0xcb, 0x87, 0x9c, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0x74, 0xf5, 0x8, 0xc9, 0x66, 0x62, 0x8d, 0x91, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0x85, 0x3a, 0xf4, 0x90, 0xaa, 0xea, 0x68, 0xec, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0xc0, 0x9c, 0xe9, 0xb1, 0xf8, 0x5b, 0x41, 0x7a, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x60, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0x5f, 0xab, 0xe3, 0x55, 0x45, 0x6b, 0x35, 0x2f, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x7a, 0xbf, 0xf6, 0x87, 0x13, 0x95, 0x1a, 0x4a, 0x2, 0xb, 0x53, 0xb7, 0x53, 0x2d, 0xff, 0xb9, 0x12, 0x8, 0x90, 0xf2, 0xe8, 0xb9, 0x81, 0x19, 0x98, 0x32, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0x87, 0x1, 0x12, 0x84, 0x1, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0x9b, 0xc2, 0x5c, 0x22, 0x81, 0x1b, 0x8b, 0xfd, 0x73, 0xf, 0xfd, 0xa6, 0x94, 0xf7, 0x28, 0x37, 0x12, 0x8, 0xed, 0x80, 0x7b, 0x4a, 0x9c, 0x47, 0x3a, 0x88, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x9b, 0xc2, 0x5c, 0x22, 0x81, 0x1b, 0x8b, 0xfd, 0x73, 0xf, 0xfd, 0xa6, 0x94, 0xf7, 0x28, 0x37, 0x12, 0x8, 0xf9, 0xd7, 0x69, 0xb1, 0xc, 0xff, 0x8e, 0x98, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0x9b, 0xc2, 0x5c, 0x22, 0x81, 0x1b, 0x8b, 0xfd, 0x73, 0xf, 0xfd, 0xa6, 0x94, 0xf7, 0x28, 0x37, 0x12, 0x8, 0x1f, 0xc8, 0xcb, 0x70, 0x50, 0xf1, 0x50, 0xb9, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0x94, 0x3, 0x12, 0x60, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0xdf, 0xf, 0x69, 0xcc, 0xb1, 0xff, 0x68, 0x60, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0x59, 0x6, 0xaa, 0x2d, 0x79, 0x3b, 0x48, 0x56, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0xa8, 0x1, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0x13, 0x4b, 0xb7, 0x34, 0xfb, 0xe3, 0xd3, 0x1a, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0xdb, 0x54, 0xf7, 0x68, 0x53, 0x3e, 0x59, 0x67, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0x7a, 0x0, 0xaf, 0xc5, 0x32, 0xcd, 0x84, 0xfe, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0x94, 0x3e, 0x2b, 0xff, 0xe3, 0xd1, 0xf8, 0x81, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x84, 0x1, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0xf4, 0x16, 0xa7, 0x17, 0x75, 0x7a, 0xc1, 0x6, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0xf0, 0x81, 0x28, 0xe6, 0xd8, 0x84, 0x50, 0x3c, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xb4, 0xc7, 0x3c, 0xe, 0xba, 0x1d, 0xce, 0xb5, 0xd7, 0x4c, 0x2d, 0x99, 0x7, 0x4e, 0xc8, 0x64, 0x12, 0x8, 0xd9, 0x6c, 0x69, 0xb2, 0x7e, 0xc8, 0xc5, 0xfa, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
		{0xa, 0x8a, 0x3, 0x12, 0x84, 0x1, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0x36, 0x46, 0x5a, 0xdc, 0x8, 0xcf, 0x9e, 0xd3, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0xd5, 0xe1, 0x3b, 0x85, 0x39, 0xaf, 0xf6, 0x7d, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0x41, 0x29, 0x97, 0x11, 0x5c, 0x88, 0xd2, 0x99, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0x20, 0xc1, 0x71, 0xbd, 0x3f, 0xa2, 0x47, 0x1d, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x84, 0x1, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0x1c, 0x45, 0xdc, 0x4, 0x68, 0x37, 0xe, 0xc4, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0x76, 0x94, 0x9, 0xc, 0x7b, 0x63, 0x93, 0xe, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0xc4, 0x93, 0xab, 0x2f, 0xc, 0x42, 0xc3, 0x98, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74, 0x12, 0x3c, 0xa, 0x16, 0xa, 0xd, 0x73, 0x75, 0x70, 0x65, 0x72, 0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x5, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12, 0x22, 0xa, 0x10, 0xd4, 0x8e, 0xcb, 0x34, 0x9d, 0x47, 0x1, 0x50, 0x23, 0x7d, 0x89, 0x86, 0x13, 0x82, 0x1d, 0xfc, 0x12, 0x8, 0xd3, 0x6d, 0x52, 0x24, 0x75, 0x7, 0xf0, 0xc7, 0x2a, 0x4, 0x74, 0x65, 0x73, 0x74},
	}

	meta := backend.NewBlockMeta("fake", uuid.MustParse("4cd3c468-6398-481b-b5ec-de56d1048427"), "v2", backend.EncSnappy, "")
	testLegacyBlock(t, ids, objs, meta, "./v2test")
}

func testLegacyBlock(t *testing.T, ids [][]byte, objs [][]byte, meta *backend.BlockMeta, path string) {
	r, _, _, err := local.New(&local.Config{
		Path: path,
	})
	require.NoError(t, err, "error creating backend")

	reader := backend.NewReader(r)
	meta, err = reader.BlockMeta(context.Background(), meta.BlockID, meta.TenantID)
	require.NoError(t, err, "error retrieving meta")

	backendBlock, err := NewBackendBlock(meta, reader)
	require.NoError(t, err, "error creating backendblock")

	// test Find
	for i, id := range ids {
		foundBytes, err := backendBlock.Find(context.Background(), id)
		assert.NoError(t, err)

		assert.Equal(t, objs[i], foundBytes)
	}

	// test Iterator
	iterator, err := backendBlock.Iterator(10)
	require.NoError(t, err, "error getting iterator")
	i := 0
	for {
		id, obj, err := iterator.Next(context.Background())
		if id == nil {
			break
		}

		assert.NoError(t, err)
		assert.Equal(t, []byte(id), ids[i])
		assert.Equal(t, obj, objs[i])
		i++
	}
	assert.Equal(t, len(ids), i)
}
