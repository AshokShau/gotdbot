package qrcode

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"reflect"
	"strings"
	"testing"
)

func TestQRCodeOutputAndQuietZone(t *testing.T) {
	q, err := NewQRCode("tg://login?token=hello")
	if err != nil {
		t.Fatal(err)
	}
	bitmap, err := q.Bitmap()
	if err != nil {
		t.Fatal(err)
	}
	if len(bitmap) != 45 {
		t.Fatalf("version 5 with border = %d", len(bitmap))
	}
	for y, row := range bitmap {
		if len(row) != 45 {
			t.Fatal("non-square bitmap")
		}
		for x, dark := range row {
			if (x < 4 || y < 4 || x >= 41 || y >= 41) && dark {
				t.Fatal("quiet zone contains dark module")
			}
		}
	}
	// Repeated encoding must not pad or modify the encoded data again.
	again, err := q.Bitmap()
	if err != nil || !reflect.DeepEqual(bitmap, again) {
		t.Fatal("repeated bitmap changed")
	}
	for _, tc := range []struct{ size, want int }{{0, 45}, {1, 45}, {90, 90}, {-2, 90}} {
		data, err := q.PNG(tc.size)
		if err != nil {
			t.Fatal(err)
		}
		img, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			t.Fatal(err)
		}
		if img.Bounds().Dx() != tc.want || img.Bounds().Dy() != tc.want {
			t.Fatalf("PNG size %d = %v", tc.size, img.Bounds())
		}
	}
	encoded, err := base64.StdEncoding.DecodeString(q.Base64PNG(90))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := png.Decode(bytes.NewReader(encoded)); err != nil {
		t.Fatal(err)
	}
	if text := q.ToSmallString(false); text == "" || text == q.ToSmallString(true) {
		t.Fatal("terminal rendering/inversion failed")
	}
	noBorder, err := NewQRCode(q.Content)
	if err != nil {
		t.Fatal(err)
	}
	noBorder.DisableBorder = true
	bare, err := noBorder.Bitmap()
	if err != nil || len(bare) != 37 {
		t.Fatalf("bare QR = %d/%v", len(bare), err)
	}
	for y := range bare {
		for x := range bare[y] {
			if bare[y][x] != bitmap[y+4][x+4] {
				t.Fatal("border changed modules")
			}
		}
	}
}

func TestQRCodeCapacity(t *testing.T) {
	for _, content := range []string{"hello", "😀日本語", "tg://login?token=abcdef"} {
		if _, err := NewQRCode(content); err != nil {
			t.Errorf("valid content %q: %v", content, err)
		}
	}
	if _, err := NewQRCode(strings.Repeat("x", 1024)); err == nil {
		t.Fatal("oversized content accepted")
	}
	if _, err := NewQRCode(""); err == nil {
		t.Fatal("empty content accepted")
	}
}

func TestBitsetRoundTripAndBounds(t *testing.T) {
	b := NewBitset()
	if err := b.AppendBytes([]byte{0xa5, 0x00, 0xff}); err != nil {
		t.Fatal(err)
	}
	if b.Len() != 24 {
		t.Fatalf("length = %d", b.Len())
	}
	for i, want := range []byte{0xa5, 0x00, 0xff} {
		if got := b.ByteAt(i * 8); got != want {
			t.Errorf("byte %d = %x", i, got)
		}
	}
	clone := CloneBitset(b)
	clone.AppendBools(true)
	if b.Len() != 24 || clone.Len() != 25 || !clone.At(24) {
		t.Fatal("clone mutated source")
	}
	part, err := b.Substr(8, 24)
	if err != nil || part.Len() != 16 || part.ByteAt(8) != 0xff {
		t.Fatalf("substring = (%v,%v)", part, err)
	}
	for _, bounds := range [][2]int{{-1, 1}, {2, 1}, {0, 25}} {
		if _, err := b.Substr(bounds[0], bounds[1]); err == nil {
			t.Errorf("invalid bounds %v accepted", bounds)
		}
	}
	if b.At(-1) || b.At(24) || b.ByteAt(-1) != 0 || b.ByteAt(24) != 0 {
		t.Fatal("out-of-range bitset access")
	}
	if err := b.AppendByte(1, 9); err == nil {
		t.Fatal("9-bit byte accepted")
	}
	if err := b.AppendUint32(1, 33); err == nil {
		t.Fatal("33-bit integer accepted")
	}
}

func TestGaloisFieldArithmetic(t *testing.T) {
	for a := range 256 {
		if gfAdd(gfElement(a), gfElement(a)) != 0 || gfMultiply(gfElement(a), 1) != gfElement(a) || gfMultiply(gfElement(a), 0) != 0 {
			t.Fatalf("field identity failed for %d", a)
		}
		for b := 1; b < 256; b++ {
			product := gfMultiply(gfElement(a), gfElement(b))
			if gfDivide(product, gfElement(b)) != gfElement(a) {
				t.Fatalf("multiply/divide inverse failed for %d/%d", a, b)
			}
		}
	}
}
