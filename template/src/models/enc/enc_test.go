package enc

import (
	"encoding/base64"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecrypt(t *testing.T) {
	key, _ := base64.StdEncoding.DecodeString("SSHujS2fda0rpBRb9ugPFQ==")
	iv, _ := base64.StdEncoding.DecodeString("O7yKcrf3ZCtyEdbm8q4mLA==")
	enc, _ := base64.StdEncoding.DecodeString("kmFxrMuOuEIuqOg3nknPduatp9vH3Eb+UmyHdO9DQZz7RTsJSsFX5Pdgea/BezpU9mWA78xtlp1dO5LDPr8fS6YfSzAp3NSQUSrBhYBzallcmAaWtzGn4JmZsKhMGXrI/eAeEF9HeX8hd7/P3MVgEbe/5PuLfSeZrVljNASeg0eS7U9T5Q9r6FW8qMmhMigrubyI2ge/ejMIQSyiq0KThg==")

	plaintext, err := decryptWXCBC(key, iv, enc)
	Convey("WeiXin Decryption should work", t, func() {
		So(err == nil, ShouldBeTrue)
		So(plaintext != nil, ShouldBeTrue)
	})
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Printf("%s\n", plaintext)
}
