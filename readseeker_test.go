package oi_test

import (
	"github.com/reiver/go-oi"

	"io"
	"strings"

	"testing"
)

func TestReadSeeker(t *testing.T) {

	data := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var readerAt io.ReaderAt = strings.NewReader(data)

	readSeeker := oi.ReadSeeker(readerAt)
	if nil == readSeeker {
		t.Errorf("nil io.ReadSeeker: %#v", readSeeker)
		return
	}

	{
		var buffer [10]byte
		var p []byte = buffer[:]

		n, err := readSeeker.Read(p)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := len(buffer), n; expected != actual {
			t.Errorf("Number of bytes actually read is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}

		if expected, actual := "0123456789", string(buffer[:]); expected != actual {
			t.Errorf("What was actually read (into the buffer) is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}

	{
		var buffer [26]byte
		var p []byte = buffer[:]

		n, err := readSeeker.Read(p)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := len(buffer), n; expected != actual {
			t.Errorf("Number of bytes actually read is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}

		if expected, actual := "ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(buffer[:]); expected != actual {
			t.Errorf("What was actually read (into the buffer) is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}

	{
		offset := int64(-1) * int64(len("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))

		absolute, err := readSeeker.Seek(offset, io.SeekCurrent)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := int64(len("0123456789")), absolute; expected != actual {
			t.Errorf("The actual resulting Seek()'ed absolute offset is not what was not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		var buffer [3]byte
		var p []byte = buffer[:]

		n, err := readSeeker.Read(p)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := len(buffer), n; expected != actual {
			t.Errorf("Number of bytes actually read is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}

		if expected, actual := "ABC", string(buffer[:]); expected != actual {
			t.Errorf("What was actually read (into the buffer) is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}

	}

	{
		offset := int64(5)

		absolute, err := readSeeker.Seek(offset, io.SeekStart)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := offset, absolute; expected != actual {
			t.Errorf("The actual resulting Seek()'ed absolute offset is not what was not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		var buffer [1]byte
		var p []byte = buffer[:]

		n, err := readSeeker.Read(p)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := len(buffer), n; expected != actual {
			t.Errorf("Number of bytes actually read is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}

		if expected, actual := "5", string(buffer[:]); expected != actual {
			t.Errorf("What was actually read (into the buffer) is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}

	}

	{
		offset := int64(-3)

		absolute, err := readSeeker.Seek(offset, io.SeekEnd)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := int64(len("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"))+offset, absolute; expected != actual {
			t.Errorf("The actual resulting Seek()'ed absolute offset is not what was not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		var buffer [3]byte
		var p []byte = buffer[:]

		n, err := readSeeker.Read(p)
		if nil != err {
			t.Error("Did not expect an error, but actually got one.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := len(buffer), n; expected != actual {
			t.Errorf("Number of bytes actually read is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}

		if expected, actual := "xyz", string(buffer[:]); expected != actual {
			t.Errorf("What was actually read (into the buffer) is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}

	}

	{
		var buffer [1]byte
		var p []byte = buffer[:]

		n, err := readSeeker.Read(p)
		if expected, actual := io.EOF, err; expected != actual {
			t.Error("Did not actually get the errror that was expected.")
			t.Logf("EXPECTED ERROR: (%T) %q", expected, expected)
			t.Logf("ACTUAL ERROR:   (%T) %q", actual, actual)
			return
		}

		if expected, actual := 0, n; expected != actual {
			t.Errorf("Number of bytes actually read is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		offset := int64(-5)

		absolute, err := readSeeker.Seek(offset, io.SeekStart)
		if nil == err {
			t.Errorf("Expected to get an error, but did not actually get one: %#v", err)
			return
		}
		if _, casted := err.(oi.InvalidOffset); !casted {
			t.Error("Expected to get an error of type oi.InvalidOffset, but did not actually get it.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			return
		}

		if expected, actual := int64(0), absolute; expected != actual {
			t.Errorf("The actual resulting Seek()'ed absolute offset is not what was not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}
}
