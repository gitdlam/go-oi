package oi


import (
	"github.com/reiver/go-oi/test"

	"testing"
)


func TestLongWrite(t *testing.T) {

	tests := []struct{
		String string
	}{
		{
			String: "",
		},



		{
			String: "apple",
		},
		{
			String: "banana",
		},
		{
			String: "cherry",
		},



		{
			String: "Hello world!",
		},



		{
			String: "😁😂😃😄😅😆😉😊😋😌😍😏😒😓😔😖😘😚😜😝😞😠😡😢😣😤😥😨😩😪😫😭😰😱😲😳😵😷",
		},



		{
			String: "0123456789",
		},
		{
			String: "٠١٢٣٤٥٦٧٨٩", // Arabic-Indic Digits
		},
		{
			String: "۰۱۲۳۴۵۶۷۸۹", // Extended Arabic-Indic Digits
		},



		{
			String: "Ⅰ Ⅱ Ⅲ Ⅳ Ⅴ Ⅵ Ⅶ Ⅷ Ⅸ Ⅹ Ⅺ Ⅻ Ⅼ Ⅽ Ⅾ Ⅿ",
		},
		{
			String: "ⅰ ⅱ ⅲ ⅳ ⅴ ⅵ ⅶ ⅷ ⅸ ⅹ ⅺ ⅻ ⅼ ⅽ ⅾ ⅿ",
		},
		{
			String: "ↀ ↁ ↂ Ↄ ↄ ↅ ↆ ↇ ↈ",
		},
	}


	for testNumber, test := range tests {

		p := []byte(test.String)

		var writer oitest.ShortWriter
		n, err := LongWrite(&writer, p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q; for %q.", testNumber, err, err.Error(), test.String)
			continue
		}
		if expected, actual := int64(len([]byte(test.String))), n; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for %q.", testNumber, expected, actual, test.String)
			continue
		}
		if expected, actual := test.String, writer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}
	}
}
