package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func assets_templates_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xbc, 0x56,
		0x4b, 0x6f, 0x22, 0x39, 0x10, 0x3e, 0x2f, 0xbf, 0xa2, 0xe4, 0x45, 0x0a,
		0x91, 0x16, 0x7a, 0x0f, 0x39, 0x91, 0xa6, 0xa5, 0x48, 0x9b, 0xc3, 0x5e,
		0x72, 0xc9, 0xdc, 0x46, 0x73, 0x30, 0x6d, 0x43, 0x5b, 0x34, 0x76, 0xc7,
		0xae, 0x66, 0x06, 0x21, 0xfe, 0xfb, 0x94, 0xdd, 0x0f, 0x9a, 0x47, 0x87,
		0x64, 0x12, 0x25, 0x12, 0x69, 0x97, 0x1f, 0x9f, 0xbf, 0xfa, 0xaa, 0xca,
		0x76, 0xec, 0x70, 0x9b, 0xcb, 0x64, 0x80, 0x02, 0xb2, 0x3b, 0xd8, 0x0d,
		0x80, 0xfe, 0xd6, 0xdc, 0x2e, 0x95, 0x9e, 0xc2, 0xbf, 0xf7, 0x83, 0xfd,
		0x60, 0x92, 0x1a, 0x8d, 0x5c, 0x69, 0x69, 0xeb, 0xd1, 0x9f, 0x4a, 0x60,
		0x36, 0x05, 0x5e, 0xa2, 0xf1, 0xe3, 0xb4, 0xb0, 0xb0, 0xb2, 0x1e, 0x9b,
		0xf3, 0x74, 0xb5, 0xb4, 0xa6, 0xd4, 0x62, 0x0a, 0xda, 0x68, 0x79, 0x5f,
		0xf5, 0x1a, 0x2b, 0xa4, 0x6d, 0x7a, 0xf6, 0x83, 0x38, 0xaa, 0x37, 0x8d,
		0x5d, 0x6a, 0x55, 0x81, 0xc9, 0x60, 0x38, 0x5a, 0x94, 0x3a, 0x45, 0x65,
		0xf4, 0xe8, 0xb6, 0x86, 0x1a, 0x8e, 0xd8, 0x77, 0xc1, 0x91, 0x8f, 0xd1,
		0x2c, 0x97, 0xb9, 0x9c, 0xdd, 0xa0, 0x31, 0x39, 0xaa, 0xe2, 0xe6, 0x07,
		0xbb, 0x9d, 0xd4, 0xed, 0xd1, 0x2d, 0xc1, 0xd1, 0x8f, 0x00, 0x6b, 0xa0,
		0x58, 0xa8, 0x0d, 0xa4, 0x39, 0x77, 0x6e, 0xc6, 0x5a, 0xe2, 0x2c, 0x09,
		0x88, 0xf1, 0xc2, 0xd8, 0x35, 0xac, 0x25, 0x66, 0x46, 0xcc, 0x58, 0x61,
		0x1c, 0x36, 0x03, 0xc8, 0xe7, 0xb9, 0x6c, 0x96, 0x55, 0x46, 0xf8, 0x3f,
		0xae, 0x98, 0x4b, 0x51, 0x9b, 0x99, 0xd9, 0xb4, 0x68, 0xd5, 0xc2, 0x4c,
		0x72, 0x71, 0xb0, 0xff, 0x8a, 0xd1, 0x76, 0x2c, 0x32, 0xb3, 0xe4, 0x1b,
		0x77, 0xab, 0x38, 0xa2, 0xc6, 0x49, 0xff, 0xa3, 0xde, 0x28, 0x6b, 0xf4,
		0x5a, 0x6a, 0x3c, 0x1e, 0xee, 0x80, 0x27, 0xcf, 0xd2, 0x6e, 0x54, 0x2a,
		0x21, 0x76, 0x05, 0xd7, 0x0d, 0xc1, 0x65, 0xbe, 0x2d, 0x32, 0x45, 0xde,
		0x41, 0xdb, 0x1a, 0xbf, 0x94, 0xd2, 0x79, 0xf9, 0xc6, 0x4e, 0x2d, 0x35,
		0x83, 0xae, 0x70, 0xac, 0x16, 0xab, 0xee, 0x2d, 0x72, 0x9e, 0x4a, 0xbf,
		0xab, 0x1f, 0xa0, 0x4e, 0x54, 0xe8, 0x27, 0xd5, 0x3b, 0x39, 0xe0, 0x14,
		0x4c, 0x87, 0xdc, 0x22, 0x79, 0xed, 0x63, 0xbc, 0xe6, 0xa8, 0x52, 0x9e,
		0xe7, 0x5b, 0x96, 0x90, 0xcc, 0x44, 0x23, 0xb9, 0xe4, 0xcd, 0x03, 0x05,
		0x6f, 0xe3, 0x55, 0x3b, 0x75, 0xb6, 0xeb, 0xcd, 0x43, 0x88, 0xb0, 0x3b,
		0x59, 0x4f, 0x66, 0x47, 0x34, 0x3f, 0x78, 0xa4, 0x69, 0x8c, 0x73, 0x23,
		0xb6, 0x07, 0x7b, 0xb7, 0x03, 0xcb, 0xf5, 0x52, 0xc2, 0xc4, 0x2b, 0xeb,
		0x60, 0xbf, 0x1f, 0x1c, 0xef, 0x63, 0x8f, 0x37, 0x27, 0x7a, 0xe2, 0x9c,
		0x4e, 0x98, 0x9a, 0xdd, 0x25, 0x31, 0x87, 0xcc, 0xca, 0xc5, 0x8c, 0x45,
		0x9e, 0x7a, 0x44, 0xe0, 0x93, 0xff, 0x05, 0x61, 0xb2, 0xc4, 0x37, 0x9f,
		0xf8, 0x5a, 0x92, 0x11, 0x47, 0x9c, 0x9c, 0xa6, 0xd9, 0xe7, 0x5e, 0x45,
		0x97, 0xc0, 0xfd, 0x8e, 0xb4, 0x5c, 0x2d, 0x60, 0xd2, 0x89, 0xb2, 0x07,
		0xa2, 0x4a, 0x49, 0x5a, 0x07, 0x86, 0xab, 0x7f, 0x60, 0xb8, 0x81, 0xe9,
		0xec, 0x74, 0x1a, 0xcd, 0x18, 0xae, 0xe8, 0x0b, 0x33, 0xef, 0x2d, 0x4d,
		0x21, 0x1f, 0xa9, 0x21, 0xb5, 0x08, 0x64, 0x6a, 0x90, 0xc6, 0xec, 0x63,
		0x70, 0xd1, 0xe7, 0x8a, 0x96, 0x36, 0x08, 0x93, 0x26, 0xb7, 0x4e, 0x04,
		0x6c, 0x21, 0xe6, 0x25, 0x22, 0xe5, 0x98, 0x2f, 0x1a, 0x1e, 0x02, 0x77,
		0x26, 0x52, 0x24, 0xb5, 0x2f, 0x0b, 0xd6, 0x24, 0xe6, 0x1c, 0x35, 0xd0,
		0x6f, 0xfc, 0xcb, 0x85, 0x8f, 0x2b, 0x53, 0x4a, 0x28, 0x47, 0x79, 0x73,
		0x2d, 0x7b, 0x29, 0x27, 0x0f, 0xe9, 0x05, 0x8f, 0x01, 0x35, 0x8e, 0x2a,
		0x02, 0xbd, 0x7e, 0xc8, 0xdc, 0x7d, 0x88, 0xbc, 0x50, 0xee, 0x35, 0xf6,
		0xc2, 0xc7, 0xc8, 0xbe, 0x81, 0xbc, 0xf3, 0x65, 0xd4, 0x92, 0xff, 0xaf,
		0x42, 0xbd, 0xce, 0x3e, 0x44, 0xef, 0x1d, 0x09, 0xf5, 0x4a, 0x38, 0x27,
		0x55, 0xf5, 0xf9, 0x7a, 0xe8, 0x13, 0x24, 0xa0, 0xf4, 0xe4, 0x7a, 0x64,
		0x4b, 0x1d, 0xac, 0x03, 0x4c, 0x53, 0x04, 0x7f, 0x5f, 0xea, 0xf6, 0xe5,
		0x00, 0xc7, 0x03, 0x57, 0xb6, 0xed, 0x51, 0x58, 0x2e, 0x78, 0x99, 0x23,
		0x7b, 0x27, 0x2b, 0xba, 0x3c, 0x84, 0x29, 0xf1, 0x0d, 0x91, 0x59, 0x28,
		0x0a, 0x6f, 0x1b, 0x99, 0x6a, 0x99, 0x27, 0xff, 0xa5, 0x54, 0xa5, 0x7d,
		0x4b, 0x12, 0x9d, 0x51, 0xa5, 0x65, 0xbd, 0x54, 0xaf, 0x24, 0x7f, 0xf0,
		0x44, 0x25, 0x4f, 0x74, 0xdb, 0xc6, 0x91, 0xfa, 0x92, 0x0c, 0x94, 0x2f,
		0x74, 0x9e, 0x20, 0xc7, 0xd2, 0x01, 0x7b, 0xa6, 0x7a, 0x28, 0xa4, 0x60,
		0x1f, 0xa9, 0xcd, 0x70, 0x01, 0x7d, 0xf6, 0xb9, 0xf2, 0xec, 0x41, 0x3f,
		0x74, 0xac, 0xd4, 0xe5, 0x76, 0x38, 0x39, 0xc3, 0x73, 0xe3, 0x8f, 0x2f,
		0x5b, 0x58, 0x97, 0x0e, 0x61, 0x2e, 0xa1, 0x3e, 0x8a, 0xe8, 0x99, 0x61,
		0xc0, 0x1f, 0x27, 0xe1, 0x1e, 0x65, 0x49, 0x7f, 0x90, 0xba, 0x3a, 0x9e,
		0x91, 0x6a, 0xd0, 0x76, 0x3b, 0x5a, 0x4d, 0xb7, 0xc8, 0xeb, 0x42, 0x7b,
		0x5a, 0x9f, 0x7a, 0x02, 0xfa, 0xf8, 0x5f, 0x55, 0xf9, 0x44, 0x46, 0x3a,
		0x8c, 0x37, 0x57, 0xdc, 0x7d, 0x4f, 0xc6, 0x1e, 0x3f, 0x29, 0xce, 0x57,
		0xd2, 0xf8, 0xe1, 0x51, 0x41, 0x86, 0x97, 0xab, 0x31, 0xbc, 0x5a, 0xf4,
		0x8e, 0x0c, 0x8c, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x8f, 0xcd,
		0x18, 0x1d, 0x0b, 0x00, 0x00,
		},
		"assets/templates/index.html",
	)
}

func assets_templates_notfound_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x55,
		0x4d, 0x6f, 0xe3, 0x38, 0x0c, 0x3d, 0x37, 0xbf, 0x82, 0xd5, 0x5e, 0xab,
		0x18, 0xe9, 0x5e, 0xf6, 0x60, 0x1b, 0xd8, 0x2d, 0xf6, 0xb0, 0xb7, 0x3d,
		0x74, 0x81, 0xbd, 0xca, 0x16, 0x63, 0x2b, 0x95, 0x25, 0x57, 0xa2, 0x93,
		0xfa, 0xdf, 0x0f, 0xe5, 0x8f, 0x7c, 0x61, 0x8a, 0xe9, 0x1c, 0x26, 0x40,
		0x20, 0x89, 0x22, 0xf9, 0xc8, 0xf7, 0xc4, 0x24, 0x7f, 0xd4, 0xbe, 0xa6,
		0xb1, 0x47, 0x68, 0xa9, 0xb3, 0xe5, 0x26, 0x4f, 0x0b, 0x58, 0xe5, 0x9a,
		0x42, 0xa0, 0x13, 0xe5, 0xe6, 0x21, 0x6f, 0x51, 0x69, 0x5e, 0x1f, 0xf2,
		0x0e, 0x49, 0x41, 0xdd, 0xaa, 0x10, 0x91, 0x0a, 0x31, 0xd0, 0x5e, 0xfe,
		0xc1, 0x0e, 0xc0, 0x9f, 0xe5, 0xae, 0x25, 0xea, 0x25, 0xbe, 0x0f, 0xe6,
		0x58, 0x88, 0xff, 0xe5, 0x7f, 0x7f, 0xca, 0x17, 0xdf, 0xf5, 0x8a, 0x4c,
		0x65, 0x51, 0x40, 0xed, 0x1d, 0xa1, 0xe3, 0xc0, 0x7f, 0xfe, 0x2e, 0x50,
		0x37, 0x78, 0x1b, 0xea, 0x54, 0x87, 0x85, 0x38, 0x1a, 0x3c, 0xf5, 0x3e,
		0xd0, 0x95, 0xf7, 0xc9, 0x68, 0x6a, 0x0b, 0x8d, 0x47, 0x53, 0xa3, 0x9c,
		0x0e, 0x4f, 0x60, 0x9c, 0x21, 0xa3, 0xac, 0x8c, 0xb5, 0xb2, 0x58, 0xec,
		0x38, 0x53, 0x2a, 0x8f, 0x0c, 0x59, 0x2c, 0x4f, 0x58, 0xc5, 0x31, 0x6a,
		0xc0, 0x10, 0x7c, 0xc8, 0xb3, 0xd9, 0x38, 0xdd, 0x5b, 0xe3, 0xde, 0x20,
		0xa0, 0x2d, 0x44, 0xa4, 0xd1, 0x62, 0x6c, 0x11, 0x19, 0xa8, 0x0d, 0xb8,
		0x2f, 0x44, 0x56, 0xc7, 0x98, 0x69, 0xdc, 0xab, 0xc1, 0xd2, 0x96, 0xf7,
		0xa2, 0xfc, 0x61, 0x44, 0x56, 0x6b, 0x77, 0x88, 0xdb, 0xda, 0xfa, 0x41,
		0xef, 0xad, 0x0a, 0xb8, 0xad, 0x7d, 0x97, 0xa9, 0x83, 0xfa, 0xc8, 0xac,
		0xa9, 0x62, 0x46, 0x27, 0x43, 0x84, 0x41, 0x56, 0xde, 0x53, 0xa4, 0xa0,
		0xfa, 0xec, 0xf7, 0xed, 0x6e, 0xbb, 0x9b, 0x90, 0xce, 0xb6, 0x05, 0x2b,
		0x81, 0xc5, 0x3a, 0x98, 0x9e, 0x20, 0x86, 0xfa, 0x0b, 0xc9, 0x0f, 0xef,
		0x03, 0x86, 0x31, 0x7b, 0x9e, 0x32, 0xce, 0x87, 0x6d, 0x67, 0xdc, 0xf6,
		0xc0, 0xd9, 0xf2, 0x6c, 0x4e, 0x55, 0xfe, 0x74, 0xd6, 0xcf, 0x4a, 0x3e,
		0x5c, 0x57, 0x7c, 0x07, 0x91, 0x67, 0xcb, 0xfb, 0xc8, 0x2b, 0xaf, 0xc7,
		0x09, 0xd3, 0xa9, 0x23, 0xd4, 0x56, 0xc5, 0x58, 0x08, 0xde, 0x56, 0x2a,
		0xc0, 0xbc, 0xc8, 0x85, 0xe0, 0xf5, 0xb8, 0x37, 0x1f, 0xa8, 0x25, 0xf9,
		0x5e, 0x40, 0xf0, 0x2c, 0x64, 0xf2, 0x36, 0x0d, 0xbf, 0x17, 0x3f, 0x3d,
		0x3c, 0xce, 0xa4, 0xcd, 0x39, 0x53, 0x7a, 0x10, 0xca, 0x38, 0xae, 0x6e,
		0x6f, 0x07, 0xa3, 0x67, 0x07, 0x80, 0xfc, 0x51, 0x4a, 0xf8, 0x2b, 0x28,
		0xa7, 0x21, 0x7d, 0xc9, 0x37, 0x8d, 0x45, 0x68, 0x90, 0xa0, 0x09, 0x7e,
		0xe8, 0x51, 0xc3, 0xde, 0x07, 0xa8, 0x30, 0xf5, 0x05, 0x9d, 0xaf, 0x0c,
		0xdf, 0x6a, 0x13, 0x7b, 0xab, 0x46, 0x90, 0x72, 0x4d, 0x72, 0x85, 0xb3,
		0xd4, 0x96, 0x9a, 0xc2, 0xb0, 0xa2, 0xb0, 0x4b, 0x35, 0x10, 0x79, 0x07,
		0x69, 0x54, 0x0a, 0x31, 0x1f, 0xc4, 0x5d, 0xcc, 0x0c, 0x2e, 0x40, 0x2b,
		0x52, 0xcb, 0x21, 0xd5, 0x6d, 0xad, 0xea, 0xe3, 0xd9, 0xac, 0x42, 0x93,
		0xa6, 0xe7, 0xb7, 0x2a, 0x4a, 0xfc, 0x50, 0x5d, 0x6f, 0x51, 0x2e, 0xe1,
		0xab, 0xa7, 0xdc, 0x5d, 0x60, 0x19, 0x38, 0xf6, 0xca, 0xad, 0x40, 0x31,
		0x48, 0xef, 0xec, 0x28, 0xca, 0xd7, 0xb9, 0xcf, 0x0b, 0x61, 0xac, 0x08,
		0xfb, 0x7d, 0x16, 0x67, 0x98, 0x3d, 0xc9, 0x18, 0x93, 0x72, 0xbf, 0xd2,
		0x2f, 0xcf, 0x66, 0x6a, 0x2e, 0x06, 0x75, 0x47, 0x52, 0x95, 0xb4, 0x3a,
		0x0f, 0x91, 0x58, 0xe7, 0x35, 0xcf, 0xd4, 0x2a, 0x46, 0xc6, 0x6a, 0x4c,
		0x23, 0xb1, 0xca, 0xfb, 0xc2, 0xc4, 0x60, 0x4d, 0x40, 0xed, 0xd4, 0x31,
		0xa4, 0xb1, 0x8c, 0x4f, 0x49, 0xd8, 0x8e, 0x97, 0x24, 0xbb, 0xe7, 0xab,
		0xb0, 0xfe, 0x66, 0x4c, 0x8a, 0x4f, 0xf4, 0x1b, 0xd7, 0x7c, 0x5f, 0xe4,
		0x95, 0x6a, 0xb8, 0xa3, 0x5e, 0x80, 0xd1, 0xac, 0xee, 0x97, 0xa4, 0xc9,
		0x07, 0x7b, 0xd5, 0xda, 0x9a, 0x89, 0x97, 0x1b, 0xf5, 0xac, 0x29, 0x99,
		0x82, 0x73, 0xb7, 0x37, 0x2c, 0x36, 0x76, 0xec, 0xdb, 0x44, 0x25, 0x9c,
		0x77, 0xfc, 0x3e, 0xe2, 0x5b, 0x3c, 0xd3, 0x0a, 0xaf, 0xe9, 0x98, 0xb8,
		0xc9, 0x79, 0x42, 0xaf, 0x58, 0x1e, 0xec, 0x2d, 0x5b, 0x69, 0x58, 0xd6,
		0x5d, 0x9e, 0x71, 0x11, 0xbc, 0xd9, 0x7c, 0x32, 0x40, 0xcb, 0x6c, 0xb5,
		0xcf, 0xe5, 0xbf, 0xaa, 0x61, 0x0a, 0x7c, 0x62, 0x6c, 0x70, 0x2c, 0x01,
		0x9b, 0x36, 0x97, 0x44, 0xac, 0xe5, 0x34, 0xd1, 0x6c, 0x9f, 0xfe, 0x1a,
		0xbe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x60, 0x90, 0x31, 0x2b, 0x06,
		0x00, 0x00,
		},
		"assets/templates/notfound.html",
	)
}

func assets_templates_task_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xbc, 0x56,
		0x5f, 0x6f, 0xdb, 0x20, 0x10, 0x7f, 0x4e, 0x3e, 0x05, 0xf2, 0xfa, 0xb8,
		0xc6, 0x7b, 0x4e, 0x1d, 0x4b, 0x53, 0x57, 0x4d, 0x95, 0xa6, 0xaa, 0x5a,
		0xf7, 0xb4, 0x37, 0x62, 0x48, 0x8c, 0x62, 0x83, 0x07, 0xd8, 0x6b, 0x55,
		0xf5, 0xbb, 0xef, 0x0e, 0xfc, 0x87, 0x38, 0x89, 0x95, 0x6e, 0xd1, 0x22,
		0xd5, 0xc0, 0x71, 0xfc, 0xf8, 0xdd, 0x3f, 0xae, 0x89, 0xb1, 0x2f, 0x05,
		0x4f, 0xe7, 0x8b, 0x4c, 0x49, 0x4b, 0x85, 0xe4, 0x9a, 0xbc, 0xce, 0x09,
		0xfc, 0x7e, 0x0b, 0x66, 0xf3, 0x25, 0xa1, 0xb5, 0x55, 0x37, 0xf3, 0xb7,
		0xb9, 0x65, 0xa4, 0xd2, 0xbc, 0xdd, 0x5b, 0xd3, 0x6c, 0xb7, 0xd5, 0xaa,
		0x96, 0x6c, 0x49, 0xa4, 0x92, 0xfc, 0xc6, 0x4b, 0x95, 0x66, 0x5c, 0x87,
		0x92, 0x8a, 0x32, 0x26, 0xe4, 0x76, 0x49, 0x3e, 0x21, 0x44, 0x12, 0xb7,
		0x77, 0x25, 0x4c, 0x34, 0x24, 0x2b, 0xa8, 0x31, 0xab, 0xa8, 0xbf, 0x36,
		0x4a, 0xe7, 0xb3, 0x64, 0xa3, 0x74, 0x49, 0x4a, 0x6e, 0x73, 0xc5, 0x56,
		0x51, 0xa5, 0x8c, 0x75, 0x52, 0x4b, 0xd7, 0x05, 0xef, 0x0e, 0xf8, 0x85,
		0xfb, 0x5e, 0xc3, 0x61, 0xc6, 0xa5, 0xe1, 0x0c, 0xd5, 0x40, 0x4f, 0xe3,
		0x00, 0x63, 0x9e, 0xfe, 0xa0, 0x66, 0x97, 0xc4, 0x30, 0xf1, 0x02, 0x96,
		0xbe, 0xbe, 0x92, 0x05, 0x0a, 0x17, 0x0f, 0xb4, 0xe4, 0xe4, 0xed, 0x0d,
		0x36, 0x99, 0x3b, 0x14, 0xfb, 0x53, 0xe1, 0xe1, 0xbb, 0x67, 0x9e, 0x81,
		0xd9, 0x7a, 0x0f, 0x20, 0x01, 0xf3, 0x07, 0x94, 0xaf, 0xdc, 0x76, 0x5a,
		0x0e, 0x0c, 0x37, 0xa7, 0x20, 0x6f, 0x55, 0x59, 0x52, 0xc9, 0x26, 0x10,
		0x5b, 0x8d, 0xb3, 0xd0, 0xee, 0x64, 0x23, 0xb4, 0x92, 0x25, 0x97, 0x76,
		0x6c, 0xa4, 0xd8, 0xb4, 0x78, 0x81, 0x0e, 0x62, 0xb6, 0x77, 0x69, 0x2a,
		0xb7, 0x9c, 0x5c, 0xed, 0x3e, 0x92, 0xab, 0x86, 0x2c, 0x57, 0x47, 0x75,
		0x41, 0xed, 0x6a, 0x07, 0x23, 0x59, 0x11, 0x9c, 0x36, 0x30, 0x9d, 0xc3,
		0x84, 0x07, 0xe4, 0x82, 0xe5, 0x69, 0x96, 0x4f, 0x5c, 0x37, 0x22, 0xe3,
		0x7b, 0x0c, 0x71, 0x9c, 0x79, 0x9a, 0x52, 0xd9, 0xf6, 0xfa, 0x56, 0x11,
		0xef, 0x21, 0xa3, 0x5f, 0xb2, 0xae, 0xad, 0x55, 0x92, 0x60, 0x62, 0xd0,
		0xcc, 0x0a, 0x25, 0x57, 0x51, 0x6c, 0xe1, 0x50, 0xdc, 0x3b, 0xee, 0x1e,
		0x79, 0xc4, 0x5c, 0x62, 0x46, 0x44, 0x5d, 0x96, 0xac, 0xad, 0x24, 0xf0,
		0x77, 0xfd, 0x6c, 0xdc, 0x60, 0xea, 0x2c, 0xe3, 0xc6, 0x44, 0x69, 0x62,
		0x2a, 0x2a, 0x3b, 0xa5, 0x6d, 0xf1, 0x52, 0xe5, 0x02, 0x92, 0x88, 0xf4,
		0xb3, 0xeb, 0xaa, 0xa0, 0x2f, 0xa0, 0x16, 0xa3, 0x5e, 0x4a, 0xee, 0x1c,
		0x6a, 0x12, 0x7b, 0x16, 0xe9, 0x01, 0x3d, 0xf4, 0x43, 0x61, 0xfe, 0x8d,
		0x39, 0x13, 0x66, 0x8a, 0x3a, 0xc3, 0x88, 0xe9, 0x33, 0x98, 0x1b, 0xab,
		0xaa, 0x81, 0xf9, 0x17, 0x8f, 0x3a, 0x4d, 0xdd, 0x85, 0xd0, 0x85, 0x66,
		0x2a, 0x8c, 0x96, 0xda, 0xda, 0x8c, 0xa2, 0x88, 0x08, 0x5d, 0x20, 0xf9,
		0xaf, 0x2e, 0x8e, 0x4e, 0x93, 0x44, 0x4f, 0xc0, 0xa4, 0x82, 0xa2, 0x74,
		0xe0, 0x5e, 0xf3, 0x6c, 0x6f, 0x18, 0x4b, 0xb5, 0xbd, 0x74, 0x18, 0x9f,
		0x10, 0x74, 0x70, 0xc5, 0xc0, 0xbe, 0x0b, 0xde, 0xfb, 0x59, 0x82, 0xaf,
		0x2f, 0x1b, 0x30, 0x74, 0xda, 0x51, 0x8a, 0x67, 0x05, 0xe9, 0x33, 0xb0,
		0x6d, 0xf0, 0x59, 0x1c, 0x3d, 0x7b, 0x41, 0xbd, 0x79, 0xfe, 0x5e, 0x11,
		0xa7, 0x81, 0xdd, 0xb3, 0x84, 0x92, 0x5c, 0xf3, 0xcd, 0x09, 0x6b, 0x75,
		0x2d, 0x07, 0xd1, 0x00, 0xe0, 0x77, 0xa3, 0xf4, 0xc3, 0xc9, 0xbd, 0x24,
		0xa6, 0x29, 0x39, 0xb2, 0x3b, 0xba, 0xfa, 0x84, 0x1f, 0xf9, 0x86, 0xd6,
		0x05, 0xe4, 0xc2, 0xdf, 0x32, 0x83, 0x28, 0x31, 0x55, 0xdb, 0x33, 0x22,
		0xb1, 0x11, 0x50, 0x7f, 0x7d, 0x24, 0xfc, 0x31, 0xe4, 0xfe, 0x9f, 0x48,
		0x72, 0x7d, 0x4e, 0xba, 0x1c, 0x90, 0x84, 0x63, 0x21, 0xc9, 0xc3, 0x74,
		0x9e, 0x25, 0x22, 0x7d, 0x80, 0x66, 0x9c, 0xc4, 0xe2, 0x58, 0x46, 0x39,
		0xc1, 0x28, 0xa9, 0x60, 0xc0, 0x67, 0x23, 0x9d, 0x4f, 0xb5, 0x5d, 0xdf,
		0xe8, 0x39, 0x6b, 0x97, 0xb9, 0x6a, 0x7c, 0xfb, 0xde, 0xcb, 0xc8, 0xef,
		0xb5, 0x0c, 0x32, 0x31, 0x4f, 0x1f, 0x05, 0xdb, 0x5b, 0xbb, 0x9a, 0xe4,
		0x63, 0x99, 0x7b, 0x38, 0xf6, 0x64, 0xdf, 0xd4, 0xb6, 0x7f, 0x7c, 0x5a,
		0x92, 0xd8, 0x95, 0xd0, 0x83, 0xe0, 0xc0, 0xbe, 0x85, 0xdd, 0x7b, 0x9b,
		0xfa, 0x16, 0xe7, 0xa5, 0xf8, 0x01, 0x22, 0xa6, 0xb7, 0xd7, 0x33, 0xf4,
		0x96, 0x63, 0x0b, 0x3e, 0xc8, 0xfb, 0x0e, 0x38, 0x88, 0x5b, 0x98, 0xe6,
		0x7d, 0x5a, 0xb7, 0x7e, 0xeb, 0x91, 0x70, 0xf3, 0xb6, 0x64, 0x8b, 0x47,
		0xad, 0xf0, 0x89, 0x5a, 0x80, 0xb9, 0x43, 0x7b, 0x0c, 0xb5, 0xba, 0xce,
		0xd7, 0xda, 0xbf, 0xb8, 0x37, 0x3f, 0xb9, 0x56, 0xbe, 0xed, 0x76, 0x42,
		0xbf, 0xda, 0xeb, 0xb0, 0xc7, 0x21, 0x9c, 0xbb, 0xc6, 0x10, 0x4e, 0x38,
		0x05, 0xd1, 0x3f, 0x77, 0xef, 0xcd, 0xea, 0x53, 0xce, 0xb9, 0x5c, 0xa5,
		0x5d, 0x94, 0xd2, 0x45, 0xea, 0x2a, 0x70, 0xde, 0x90, 0x7e, 0x5d, 0x09,
		0xf5, 0xe5, 0x02, 0x33, 0x6c, 0x1a, 0xf0, 0xbf, 0x2d, 0xf4, 0xf3, 0x26,
		0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x06, 0xcb, 0x72, 0x56, 0x0b,
		0x00, 0x00,
		},
		"assets/templates/task.html",
	)
}

func assets_templates_error_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x55,
		0xcb, 0x6e, 0xdb, 0x3a, 0x10, 0x5d, 0xc7, 0x5f, 0xc1, 0xf0, 0x6e, 0x43,
		0x09, 0xce, 0xdd, 0xdc, 0x85, 0x24, 0xe0, 0x36, 0x08, 0x8a, 0xac, 0x9b,
		0x02, 0xdd, 0x52, 0xe2, 0x58, 0xa2, 0x43, 0x91, 0x0a, 0x39, 0xb2, 0x2d,
		0x04, 0xf9, 0xf7, 0x0e, 0xf5, 0xf0, 0x0b, 0x0d, 0x9a, 0x2e, 0x6a, 0xc0,
		0xe0, 0x6b, 0x66, 0xce, 0xf0, 0x9c, 0x19, 0x2a, 0xbb, 0x55, 0xae, 0xc2,
		0xa1, 0x03, 0xd6, 0x60, 0x6b, 0x8a, 0x55, 0x16, 0x07, 0x66, 0xa4, 0xad,
		0x73, 0x0e, 0x96, 0x17, 0xab, 0x9b, 0xac, 0x01, 0xa9, 0x68, 0xbc, 0xc9,
		0x5a, 0x40, 0xc9, 0xaa, 0x46, 0xfa, 0x00, 0x98, 0xf3, 0x1e, 0x37, 0xe2,
		0x3f, 0x32, 0x60, 0xf4, 0x9b, 0xcf, 0x1a, 0xc4, 0x4e, 0xc0, 0x6b, 0xaf,
		0x77, 0x39, 0xff, 0x21, 0xbe, 0xff, 0x2f, 0x1e, 0x5c, 0xdb, 0x49, 0xd4,
		0xa5, 0x01, 0xce, 0x2a, 0x67, 0x11, 0x2c, 0x39, 0x3e, 0x3d, 0xe6, 0xa0,
		0x6a, 0xb8, 0x74, 0xb5, 0xb2, 0x85, 0x9c, 0xef, 0x34, 0xec, 0x3b, 0xe7,
		0xf1, 0xcc, 0x7a, 0xaf, 0x15, 0x36, 0xb9, 0x82, 0x9d, 0xae, 0x40, 0x8c,
		0x8b, 0x3b, 0xa6, 0xad, 0x46, 0x2d, 0x8d, 0x08, 0x95, 0x34, 0x90, 0xaf,
		0x29, 0x52, 0x4c, 0x0f, 0x35, 0x1a, 0x28, 0xf6, 0x50, 0x86, 0x21, 0x28,
		0x06, 0xde, 0x3b, 0x9f, 0xa5, 0xd3, 0xe6, 0x78, 0x6e, 0xb4, 0x7d, 0x61,
		0x1e, 0x4c, 0xce, 0x03, 0x0e, 0x06, 0x42, 0x03, 0x40, 0x40, 0x8d, 0x87,
		0x4d, 0xce, 0xd3, 0x2a, 0x84, 0x54, 0xc1, 0x46, 0xf6, 0x06, 0x13, 0x9a,
		0xf3, 0xe2, 0xb7, 0x1e, 0x69, 0xa5, 0xec, 0x36, 0x24, 0x95, 0x71, 0xbd,
		0xda, 0x18, 0xe9, 0x21, 0xa9, 0x5c, 0x9b, 0xca, 0xad, 0x3c, 0xa4, 0x46,
		0x97, 0x21, 0xc5, 0xbd, 0x46, 0x04, 0x2f, 0x4a, 0xe7, 0x30, 0xa0, 0x97,
		0x5d, 0xfa, 0x6f, 0xb2, 0x4e, 0xd6, 0x23, 0xd2, 0x71, 0x6f, 0xc6, 0x8a,
		0x60, 0xa1, 0xf2, 0xba, 0x43, 0x16, 0x7c, 0xf5, 0x89, 0xe0, 0xdb, 0xd7,
		0x1e, 0xfc, 0x90, 0xde, 0x8f, 0x11, 0xa7, 0x45, 0xd2, 0x6a, 0x9b, 0x6c,
		0x29, 0x5a, 0x96, 0x4e, 0xa1, 0x8a, 0x3f, 0x8e, 0xfa, 0x51, 0xca, 0xdb,
		0xf3, 0x8c, 0xaf, 0x20, 0xb2, 0x74, 0xae, 0x8f, 0xac, 0x74, 0x6a, 0x18,
		0x31, 0xad, 0xdc, 0xb1, 0xca, 0xc8, 0x10, 0x72, 0x4e, 0xd3, 0x52, 0x7a,
		0x36, 0x0d, 0x62, 0x26, 0x78, 0x59, 0x6e, 0xf4, 0x01, 0x94, 0x40, 0xd7,
		0x71, 0xe6, 0x1d, 0x09, 0x19, 0xad, 0x75, 0x4d, 0xf5, 0xe2, 0xc6, 0xc2,
		0xa3, 0x48, 0x4a, 0x1f, 0x23, 0xc5, 0x82, 0x90, 0xda, 0x52, 0x76, 0x1b,
		0xd3, 0x6b, 0x35, 0x19, 0x30, 0x96, 0xdd, 0x0a, 0xc1, 0xbe, 0x78, 0x69,
		0x15, 0x8b, 0x7f, 0x74, 0x75, 0x6d, 0x80, 0xd5, 0x80, 0xac, 0xf6, 0xae,
		0xef, 0x40, 0xb1, 0x8d, 0xf3, 0xac, 0x84, 0x78, 0x2f, 0xd6, 0xba, 0x52,
		0xd3, 0xa9, 0xd2, 0xa1, 0x33, 0x72, 0x60, 0x42, 0x2c, 0x41, 0xce, 0x70,
		0xe6, 0xdc, 0xe2, 0xa5, 0xc0, 0x2f, 0x28, 0x64, 0x52, 0xf6, 0x88, 0xce,
		0xb2, 0xd8, 0x2a, 0x39, 0x9f, 0x16, 0xfc, 0xca, 0x67, 0x02, 0xe7, 0x4c,
		0x49, 0x94, 0xf3, 0x22, 0xe6, 0x6d, 0x8c, 0xec, 0xc2, 0x71, 0x5b, 0xfa,
		0x3a, 0x76, 0xcf, 0x3f, 0x65, 0x10, 0x70, 0x90, 0x6d, 0x67, 0x40, 0xcc,
		0xee, 0x8b, 0xa5, 0x58, 0x9f, 0x60, 0x09, 0x38, 0x74, 0xd2, 0x2e, 0x40,
		0xc1, 0x0b, 0x67, 0xcd, 0xc0, 0x8b, 0xe7, 0xe9, 0x9e, 0x27, 0xc2, 0x48,
		0x11, 0xb2, 0xfb, 0xc8, 0x4f, 0x13, 0x7b, 0x82, 0x30, 0x46, 0xe5, 0xfe,
		0xa6, 0x5d, 0x96, 0x4e, 0xd4, 0x9c, 0x36, 0xe4, 0x15, 0x49, 0x65, 0xd4,
		0xea, 0xd8, 0x44, 0x7c, 0xe9, 0xd7, 0x2c, 0x95, 0x8b, 0x18, 0x29, 0xa9,
		0x31, 0xb6, 0xc4, 0x22, 0xef, 0x03, 0x11, 0x03, 0x15, 0x32, 0x6c, 0xc6,
		0x1b, 0xb3, 0xd8, 0x96, 0xe1, 0x2e, 0x0a, 0xdb, 0xd2, 0x10, 0x65, 0x77,
		0x74, 0xe4, 0x97, 0x37, 0x63, 0x54, 0x7c, 0xa4, 0x5f, 0xdb, 0xfa, 0xd7,
		0x22, 0x2f, 0x54, 0xb3, 0x2b, 0xea, 0x39, 0xd3, 0x8a, 0xd4, 0xfd, 0x94,
		0x34, 0x59, 0x6f, 0xce, 0xae, 0xb6, 0x44, 0xa2, 0xe1, 0x42, 0x3d, 0xa3,
		0x0b, 0xa2, 0xe0, 0x78, 0xdb, 0x0b, 0x16, 0x6b, 0x33, 0x74, 0x4d, 0xa4,
		0x92, 0x1d, 0x67, 0x54, 0x1f, 0xe1, 0x25, 0x1c, 0x69, 0x65, 0xcf, 0x71,
		0x19, 0xb9, 0xc9, 0xa8, 0x43, 0xcf, 0x58, 0xee, 0xcd, 0x25, 0x5b, 0xb1,
		0x59, 0x96, 0x59, 0x96, 0x52, 0x12, 0x34, 0x59, 0x7d, 0xd0, 0x40, 0x73,
		0x6f, 0x35, 0xf7, 0xc5, 0xe3, 0xa1, 0x82, 0x2e, 0x56, 0x0f, 0x7b, 0x7b,
		0x4b, 0xbe, 0x3a, 0xa4, 0xb7, 0x3d, 0x79, 0x8c, 0xef, 0xe6, 0x93, 0x7a,
		0x7f, 0xa7, 0xc6, 0xbe, 0x9f, 0x2c, 0x3b, 0x0f, 0xc5, 0xc9, 0xe0, 0x1b,
		0xca, 0xea, 0x25, 0x1e, 0xc7, 0xed, 0xd5, 0x09, 0x97, 0xa4, 0x1f, 0x1f,
		0x00, 0xf2, 0x1b, 0xbf, 0x24, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f,
		0x95, 0x40, 0x4f, 0x5a, 0x06, 0x00, 0x00,
		},
		"assets/templates/error.html",
	)
}

func assets_templates_layout_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x56,
		0xdb, 0x8a, 0xe4, 0x36, 0x10, 0x7d, 0xde, 0xfd, 0x8a, 0x5a, 0xe5, 0x75,
		0x65, 0x31, 0xc9, 0x4b, 0x08, 0x76, 0x43, 0xb2, 0x04, 0xb2, 0x10, 0x92,
		0x10, 0x26, 0x90, 0x57, 0xd9, 0x2a, 0xdb, 0xea, 0x95, 0x25, 0x8f, 0x24,
		0xf7, 0x8c, 0x69, 0xfa, 0xdf, 0x53, 0xf2, 0xad, 0x2f, 0x99, 0x65, 0x7b,
		0x03, 0x19, 0x18, 0x74, 0xab, 0x3a, 0x75, 0x4e, 0x95, 0x4a, 0xee, 0xfc,
		0x9d, 0x72, 0x55, 0x1c, 0x7b, 0x84, 0x36, 0x76, 0x66, 0xf7, 0x36, 0x4f,
		0x03, 0x18, 0x69, 0x9b, 0x82, 0xa1, 0x65, 0xbb, 0xb7, 0x6f, 0xf2, 0x16,
		0xa5, 0xa2, 0xf1, 0x4d, 0xde, 0x61, 0x94, 0x50, 0xb5, 0xd2, 0x07, 0x8c,
		0x05, 0x1b, 0x62, 0xcd, 0xbf, 0x27, 0x03, 0xa0, 0xbf, 0xe5, 0xac, 0x8d,
		0xb1, 0xe7, 0xf8, 0x34, 0xe8, 0x43, 0xc1, 0xfe, 0xe6, 0x7f, 0xfd, 0xc8,
		0x3f, 0xb8, 0xae, 0x97, 0x51, 0x97, 0x06, 0x19, 0x54, 0xce, 0x46, 0xb4,
		0xe4, 0xf8, 0xf1, 0xe7, 0x02, 0x55, 0x83, 0xd7, 0xae, 0x56, 0x76, 0x58,
		0xb0, 0x83, 0xc6, 0xe7, 0xde, 0xf9, 0x78, 0x61, 0xfd, 0xac, 0x55, 0x6c,
		0x0b, 0x85, 0x07, 0x5d, 0x21, 0x9f, 0x16, 0xef, 0x41, 0x5b, 0x1d, 0xb5,
		0x34, 0x3c, 0x54, 0xd2, 0x60, 0xf1, 0x40, 0x48, 0x89, 0x5e, 0xd4, 0xd1,
		0xe0, 0xee, 0x78, 0xcc, 0x1e, 0xd3, 0xe4, 0x74, 0xca, 0xc5, 0xbc, 0x33,
		0x1d, 0x1a, 0x6d, 0x3f, 0x81, 0x47, 0x53, 0xb0, 0x10, 0x47, 0x83, 0xa1,
		0x45, 0xa4, 0x28, 0xad, 0xc7, 0xba, 0x60, 0xa2, 0x0a, 0x41, 0x28, 0xac,
		0xe5, 0x60, 0x62, 0x46, 0x73, 0xb6, 0xfb, 0xa2, 0x87, 0xa8, 0x94, 0xdd,
		0x87, 0xac, 0x32, 0x6e, 0x50, 0xb5, 0x91, 0x1e, 0xb3, 0xca, 0x75, 0x42,
		0xee, 0xe5, 0x8b, 0x30, 0xba, 0x0c, 0x22, 0x3e, 0xeb, 0x18, 0xd1, 0xf3,
		0xd2, 0xb9, 0x18, 0xa2, 0x97, 0xbd, 0xf8, 0x2e, 0x7b, 0xc8, 0x1e, 0xa6,
		0x48, 0xdb, 0xde, 0x12, 0x2b, 0x05, 0x0b, 0x95, 0xd7, 0x7d, 0x84, 0xe0,
		0xab, 0x3b, 0xc0, 0xf7, 0x4f, 0x03, 0xfa, 0x51, 0x7c, 0x3b, 0x21, 0xce,
		0x8b, 0xac, 0xd3, 0x36, 0xdb, 0x13, 0x5a, 0x2e, 0x66, 0xa8, 0xdd, 0x57,
		0xa3, 0x7e, 0x8e, 0xf2, 0xfe, 0x92, 0xf1, 0x4d, 0x88, 0x5c, 0x2c, 0x97,
		0x23, 0x2f, 0x9d, 0x1a, 0xa7, 0x98, 0x56, 0x1e, 0xa0, 0x32, 0x32, 0x84,
		0x82, 0xd1, 0xb4, 0x94, 0x1e, 0xe6, 0x81, 0x2f, 0x09, 0x5e, 0x97, 0xb5,
		0x7e, 0x41, 0xc5, 0xa3, 0xeb, 0x19, 0x78, 0x47, 0x55, 0x4c, 0xd6, 0xba,
		0xa1, 0xcb, 0xe2, 0xa6, 0x5b, 0x47, 0x48, 0x4a, 0x6f, 0x48, 0xe9, 0x36,
		0x48, 0x6d, 0x89, 0x5d, 0x6d, 0x06, 0xad, 0x66, 0x03, 0x80, 0xfc, 0x1d,
		0xe7, 0xf0, 0x93, 0x97, 0x56, 0x41, 0xfa, 0x8f, 0xae, 0x69, 0x0c, 0x42,
		0x83, 0x11, 0x1a, 0xef, 0x86, 0x1e, 0x15, 0xd4, 0xce, 0x43, 0x89, 0x49,
		0x17, 0x74, 0xae, 0xd4, 0x74, 0xaa, 0x74, 0xe8, 0x8d, 0x1c, 0x81, 0xf3,
		0x15, 0xe4, 0x22, 0xce, 0xc2, 0x2d, 0x89, 0x42, 0xbf, 0x46, 0x21, 0x93,
		0x72, 0x88, 0xd1, 0x59, 0x48, 0x7d, 0x52, 0xb0, 0x79, 0xc1, 0x6e, 0x7c,
		0xe6, 0xe0, 0x0c, 0x94, 0x8c, 0x72, 0x59, 0x24, 0xde, 0xc6, 0xc8, 0x3e,
		0x6c, 0xdb, 0xd2, 0x37, 0xa9, 0x75, 0xbe, 0x29, 0x03, 0xc7, 0x17, 0xd9,
		0xf5, 0x06, 0xf9, 0xe2, 0xbe, 0x5a, 0xf2, 0x87, 0x73, 0x58, 0x0a, 0x1c,
		0x7a, 0x69, 0xd7, 0x40, 0xc1, 0x73, 0x67, 0xcd, 0xc8, 0x76, 0x8f, 0xb3,
		0xce, 0x73, 0xc2, 0xa8, 0x22, 0x64, 0xf7, 0x39, 0x3f, 0x4d, 0xd9, 0xe3,
		0x14, 0x63, 0xaa, 0xdc, 0xff, 0x69, 0x97, 0x8b, 0x39, 0x35, 0xe7, 0x0d,
		0x79, 0x93, 0xa4, 0x32, 0xd5, 0x6a, 0x6b, 0x22, 0xb6, 0x7b, 0xc6, 0x32,
		0x8c, 0x41, 0xe5, 0x42, 0xae, 0xc5, 0x10, 0x54, 0x8d, 0xa9, 0x25, 0xd6,
		0xf2, 0x7e, 0xa0, 0xc4, 0x60, 0x15, 0x21, 0xb6, 0x93, 0x62, 0x48, 0x6d,
		0x19, 0xde, 0xa7, 0xc2, 0x76, 0x34, 0xa4, 0xb2, 0x3b, 0x3a, 0xf2, 0xeb,
		0x83, 0x31, 0x55, 0x7c, 0x4a, 0xbf, 0xb6, 0xcd, 0xeb, 0x45, 0x5e, 0x53,
		0x0d, 0x37, 0xa9, 0x67, 0xa0, 0x15, 0x55, 0xf7, 0xae, 0xd2, 0xe4, 0x83,
		0xb9, 0x90, 0xb6, 0x22, 0xd1, 0x70, 0x55, 0x3d, 0xa3, 0x8f, 0x47, 0xd0,
		0x35, 0xe0, 0x13, 0x64, 0x7f, 0xc8, 0x06, 0x81, 0x3d, 0xca, 0xf0, 0x29,
		0x30, 0x38, 0x9d, 0x56, 0x6f, 0x59, 0x45, 0x7d, 0x40, 0x76, 0x3c, 0xa2,
		0x55, 0xa7, 0xd3, 0x8e, 0x32, 0xb6, 0x25, 0xe7, 0x2a, 0xe9, 0x8d, 0x19,
		0xfb, 0x36, 0x65, 0x1e, 0xb6, 0x19, 0x5d, 0xa7, 0x04, 0xb6, 0x56, 0x01,
		0x26, 0xec, 0x94, 0xca, 0x9c, 0x1a, 0xfa, 0x82, 0xc6, 0xcc, 0x21, 0x4b,
		0xc7, 0x14, 0xf8, 0xbc, 0x9f, 0x9e, 0x38, 0xb8, 0x25, 0xf8, 0x8b, 0x0e,
		0xd1, 0xf9, 0x31, 0x51, 0xbc, 0x65, 0x08, 0x44, 0x11, 0xae, 0x38, 0x26,
		0x02, 0x82, 0x0e, 0x26, 0xec, 0xec, 0x63, 0x3a, 0xbd, 0x83, 0xb6, 0x92,
		0xa1, 0x2d, 0x9d, 0xf4, 0xea, 0x9a, 0xfa, 0x0f, 0xb0, 0x21, 0xfd, 0x46,
		0x1f, 0x03, 0x48, 0x0f, 0xf8, 0x6b, 0x5a, 0x66, 0x16, 0xaf, 0xca, 0xfb,
		0x73, 0xb0, 0x5f, 0x56, 0xb8, 0x18, 0xfe, 0x67, 0x85, 0xc2, 0x0f, 0x76,
		0xdb, 0x22, 0xa0, 0xbb, 0x75, 0xa7, 0x67, 0xe7, 0x2c, 0x99, 0x3c, 0xcf,
		0x8a, 0x37, 0x98, 0xaf, 0x91, 0x7c, 0x2d, 0xe9, 0xf7, 0x21, 0xf6, 0x43,
		0x64, 0xff, 0xd6, 0x7f, 0x2d, 0xf2, 0xac, 0xed, 0x0e, 0xc6, 0x35, 0x3d,
		0x98, 0x1b, 0xe3, 0x5f, 0x5d, 0xb3, 0x10, 0x4e, 0xbf, 0x16, 0xee, 0xa3,
		0x9a, 0x8b, 0xc1, 0x5c, 0x77, 0x76, 0x7a, 0xd8, 0xd7, 0x59, 0x2e, 0xa8,
		0x61, 0xa6, 0x66, 0xa7, 0x8f, 0xf6, 0x87, 0xb9, 0x7f, 0x93, 0x2f, 0xbd,
		0x23, 0xd3, 0xd7, 0x84, 0xbe, 0x2e, 0xd3, 0x6f, 0x92, 0x7f, 0x02, 0x00,
		0x00, 0xff, 0xff, 0x65, 0xd1, 0x28, 0x18, 0xa4, 0x08, 0x00, 0x00,
		},
		"assets/templates/layout.html",
	)
}

func assets_templates_taskrun_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x54,
		0xc1, 0x8e, 0xda, 0x30, 0x10, 0x3d, 0x6f, 0xbe, 0x62, 0x94, 0x72, 0x2c,
		0x49, 0xcf, 0x6c, 0x92, 0x43, 0xab, 0x55, 0xb5, 0x52, 0x55, 0xad, 0xda,
		0x9e, 0x7a, 0x33, 0xf1, 0x84, 0x44, 0x24, 0x76, 0x64, 0x3b, 0x74, 0x11,
		0xe2, 0xdf, 0x3b, 0xe3, 0x24, 0x10, 0xa0, 0x2c, 0xa8, 0xd2, 0x22, 0x81,
		0x9d, 0x99, 0x97, 0xf7, 0xe6, 0xd9, 0xcc, 0x24, 0xd6, 0x6d, 0x6b, 0xcc,
		0x82, 0x28, 0xd7, 0xca, 0x89, 0x4a, 0xa1, 0x81, 0x5d, 0x00, 0xf4, 0xf9,
		0x53, 0x49, 0x57, 0x2e, 0x40, 0x74, 0x4e, 0x3f, 0x06, 0xfb, 0xc0, 0x49,
		0x68, 0x0d, 0x0e, 0xb9, 0xa5, 0xc8, 0xd7, 0x2b, 0xa3, 0x3b, 0x25, 0x17,
		0xa0, 0xb4, 0xc2, 0xc7, 0x3e, 0xaa, 0x8d, 0x44, 0x33, 0x8d, 0xb4, 0x42,
		0xca, 0x4a, 0xad, 0x16, 0xf0, 0x89, 0x29, 0x92, 0x78, 0xd0, 0x4a, 0x64,
		0xb5, 0x81, 0xbc, 0x16, 0xd6, 0xa6, 0xe1, 0x41, 0x36, 0xcc, 0x82, 0x87,
		0xa4, 0xd0, 0xa6, 0x81, 0x06, 0x5d, 0xa9, 0x65, 0x1a, 0xb6, 0xda, 0x3a,
		0x1f, 0x75, 0x62, 0x59, 0xe3, 0xf8, 0x42, 0xff, 0xe0, 0x7f, 0xe7, 0xf4,
		0xb2, 0x44, 0x65, 0x51, 0x32, 0x8c, 0x70, 0x86, 0x17, 0x5a, 0xcb, 0xec,
		0x97, 0xb0, 0x6b, 0x30, 0x9d, 0x4a, 0x62, 0x7a, 0xe8, 0x83, 0x32, 0xdb,
		0xed, 0x20, 0xe2, 0x44, 0xf4, 0x5d, 0x34, 0x08, 0xfb, 0x3d, 0x7c, 0x18,
		0x23, 0x3f, 0x3a, 0x15, 0x3d, 0x4b, 0x0a, 0x11, 0x5e, 0x7a, 0xae, 0xb8,
		0x27, 0x9b, 0x72, 0x3e, 0xbd, 0x62, 0x4e, 0xa7, 0x61, 0x4e, 0x38, 0x13,
		0x3a, 0x95, 0x23, 0xf1, 0x57, 0x74, 0x23, 0xca, 0x93, 0x71, 0xf2, 0x2d,
		0xca, 0x2f, 0xba, 0x69, 0x84, 0x92, 0xd7, 0x19, 0xb9, 0xb0, 0x01, 0x74,
		0x17, 0xe1, 0x93, 0xda, 0x54, 0x46, 0xab, 0x06, 0x95, 0x3b, 0xb7, 0x5e,
		0x15, 0x47, 0xca, 0x09, 0x8c, 0x69, 0x07, 0x45, 0x23, 0xd4, 0x0a, 0x61,
		0xb6, 0xfe, 0x08, 0xb3, 0x0d, 0x2c, 0xd2, 0x6b, 0x70, 0x42, 0xce, 0xd6,
		0x7c, 0x7c, 0x29, 0xf0, 0x76, 0x43, 0xdb, 0x80, 0x36, 0x38, 0x29, 0x71,
		0xf2, 0x78, 0xbd, 0xd6, 0x9f, 0x4e, 0x18, 0x87, 0xf2, 0x1f, 0x75, 0x2a,
		0xed, 0x8e, 0xe2, 0x03, 0x2c, 0x7a, 0xb6, 0xbf, 0xd1, 0xe8, 0x5e, 0xff,
		0x3c, 0xd9, 0x47, 0xef, 0x91, 0xd4, 0x6d, 0x7b, 0x8f, 0xa4, 0x87, 0x5d,
		0x93, 0xf4, 0xc9, 0x7b, 0x25, 0x5f, 0xaa, 0x0b, 0xb9, 0xe3, 0xcd, 0x36,
		0x32, 0x7a, 0x31, 0x3a, 0x47, 0x6b, 0x23, 0xc2, 0xdd, 0x2c, 0x5e, 0xea,
		0xee, 0xf4, 0x5a, 0x79, 0x7d, 0x38, 0x2d, 0x8e, 0x31, 0x9f, 0xbb, 0x22,
		0xfa, 0x86, 0x8a, 0xef, 0x68, 0xb9, 0x75, 0x68, 0x3d, 0x2c, 0x11, 0x63,
		0x0b, 0x2d, 0x9d, 0x02, 0xfa, 0xce, 0x5f, 0xad, 0x5f, 0x24, 0x16, 0xa2,
		0xab, 0x5d, 0x08, 0xa5, 0xc1, 0x22, 0x0d, 0x63, 0x47, 0x5c, 0xf1, 0xe1,
		0x2f, 0xed, 0xdb, 0x22, 0xa6, 0x56, 0x8a, 0x2f, 0x9a, 0x85, 0xfa, 0x99,
		0xd5, 0xc2, 0x2c, 0xb1, 0xad, 0x50, 0x23, 0xfb, 0xaa, 0xde, 0xb6, 0x65,
		0x45, 0xad, 0x09, 0x87, 0xdd, 0xbc, 0xa8, 0x6a, 0x24, 0x58, 0xcc, 0xb8,
		0x0c, 0xec, 0x60, 0x44, 0xf4, 0x3e, 0xde, 0xb6, 0x8c, 0xc6, 0xdc, 0xb4,
		0x4c, 0x98, 0x4b, 0xcb, 0x3c, 0x80, 0xde, 0xcb, 0x35, 0x09, 0xfe, 0x8f,
		0x6b, 0xef, 0x85, 0x5d, 0xfb, 0xda, 0xce, 0x8c, 0xd3, 0xc2, 0x43, 0xcd,
		0xef, 0x78, 0x0e, 0xd2, 0x9c, 0x8c, 0x69, 0x50, 0x66, 0x7f, 0x03, 0x00,
		0x00, 0xff, 0xff, 0xd7, 0xef, 0xf2, 0xe7, 0xa2, 0x05, 0x00, 0x00,
		},
		"assets/templates/taskrun.html",
	)
}

func assets_templates_log_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x3c, 0x8d,
		0x41, 0xae, 0xc2, 0x30, 0x0c, 0x05, 0xd7, 0xdf, 0xa7, 0xb0, 0x7a, 0x80,
		0x76, 0xdf, 0xe6, 0xe7, 0x04, 0x48, 0x9c, 0x21, 0x24, 0x51, 0x89, 0xa8,
		0xec, 0xca, 0x71, 0x40, 0x28, 0xca, 0xdd, 0x09, 0xa2, 0xe0, 0xe5, 0x58,
		0x33, 0xcf, 0x64, 0x7d, 0x6e, 0xd1, 0xc2, 0xe8, 0x99, 0xd4, 0x25, 0x8a,
		0x82, 0x15, 0xb0, 0xdf, 0x23, 0x05, 0xbd, 0xce, 0xe8, 0x8a, 0xf2, 0x02,
		0x0d, 0x76, 0x89, 0xc7, 0xe3, 0xe2, 0xfc, 0x6d, 0x15, 0x2e, 0x14, 0x66,
		0x24, 0xa6, 0xb8, 0x7c, 0x28, 0x4b, 0x88, 0xf2, 0x25, 0x0d, 0xcc, 0x74,
		0x84, 0x4d, 0x48, 0x77, 0xf4, 0x9b, 0xcb, 0xf9, 0x7f, 0xf8, 0x6d, 0x0c,
		0x16, 0xfe, 0x4c, 0x4f, 0xda, 0x5a, 0x71, 0x3c, 0xf1, 0x7a, 0x2e, 0xba,
		0x17, 0xc5, 0xd6, 0xcc, 0xf4, 0xa6, 0xdd, 0xee, 0x96, 0x7d, 0x05, 0x00,
		0x00, 0xff, 0xff, 0x02, 0x5b, 0xd8, 0x2a, 0x9c, 0x00, 0x00, 0x00,
		},
		"assets/templates/log.html",
	)
}

func assets_css_default_css() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0xca,
		0x4f, 0xa9, 0x54, 0xa8, 0xe6, 0xe2, 0x2c, 0x48, 0x4c, 0x49, 0xc9, 0xcc,
		0x4b, 0xd7, 0x2d, 0xc9, 0x2f, 0xb0, 0x52, 0x30, 0x37, 0x28, 0xa8, 0xb0,
		0xe6, 0xaa, 0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x56, 0xf0, 0x95,
		0x1c, 0x00, 0x00, 0x00,
		},
		"assets/css/default.css",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"assets/templates/index.html": assets_templates_index_html,
	"assets/templates/notfound.html": assets_templates_notfound_html,
	"assets/templates/task.html": assets_templates_task_html,
	"assets/templates/error.html": assets_templates_error_html,
	"assets/templates/layout.html": assets_templates_layout_html,
	"assets/templates/taskrun.html": assets_templates_taskrun_html,
	"assets/templates/log.html": assets_templates_log_html,
	"assets/css/default.css": assets_css_default_css,

}
