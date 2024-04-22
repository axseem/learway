package security_test

import (
	"testing"

	"github.com/axseem/learway/security"
)

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		desc     string
		password string
		valid    bool
	}{
		{
			desc:     "too short",
			password: "0Aaaaaa",
			valid:    false,
		},
		{
			desc:     "too long",
			password: "0Aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			valid:    false,
		},
		{
			desc:     "no number",
			password: "Aaaaaaaa",
			valid:    false,
		},
		{
			desc:     "no uppercase letter",
			password: "0aaaaaaa",
			valid:    false,
		},
		{
			desc:     "no lowercase letter",
			password: "0AAAAAAA",
			valid:    false,
		},
		{
			desc:     "no lowercase letter",
			password: "0AAAAAAA",
			valid:    false,
		},
		{
			desc:     "non-ascii characters",
			password: "とても-Καλός0密码",
			valid:    false,
		},
		{
			desc:     "all ascii characters",
			password: "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
			valid:    true,
		},
		{
			desc:     "valid password",
			password: "*UL2}fB!UV&9",
			valid:    true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := security.ValidatePassword(tC.password)
			if (err != nil) && tC.valid {
				t.Error(tC.desc)
			}
		})
	}
}
