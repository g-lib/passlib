/*
* @Author: Tacey Wong
* @Date:   2020-09-21 22:20:28
* @Last Modified by:   Tacey Wong
* @Last Modified time: 2020-09-21 22:35:04
 */
package pwd

import (
	"github.com/g-lib/passlib/utils"
)

const (
	// barest protection from throttled online attack
	EntropyUnsafe = 12
	// some protection from unthrottled online attack
	EntropyWeak = 24
	// some protection from offline attacks
	EntropyFair = 36
	// reasonable protection from offline attacks
	EntropyStrong = 48
	// very good protection from offline attacks
	EntropySecure = 60
)

// global dict of predefined characters sets
var DefaultCharsets = map[string]string{
	// ascii letters, digits, and some punctuation
	"ascii_72": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*?/",

	//ascii letters and digits
	"ascii_62": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",

	// ascii_50, without visually similar '1IiLl', '0Oo', '5S', '8B'
	"ascii_50": "234679abcdefghjkmnpqrstuvwxyzACDEFGHJKMNPQRTUVWXYZ",

	// lower case hexadecimal
	"hex": "0123456789abcdef",
}

// Need lazy-load
var DefaultWordsets = map[string]string{}

type GenWordOpt struct {
	Entropy string
	Length  int
	Charset string
	Chars   string
	Returns string
}

// GenWord Generate one or more random passwords.
func GenWord(opt *GenWordOpt) string {
	return utils.GetRandomStr(true, opt.Length)
}

func GenWordQuickly(length int) string {
	opt := GenWordOpt{
		Length: length,
	}
	return GenWord(&opt)
}

type GenPhraseOpt struct {
	Entropy string
	Length  int
	Wordset string
	Words   string
	Sep     string
	Returns string
}

// GenPhrase Generate one or more random password/passphrases
func GenPhrase(opt *GenPhraseOpt) string {
	return ""
}

func GenPhraseQuickly(lenght int) string {
	opt := GenPhraseOpt{
		Length: lenght,
	}
	return GenPhrase(&opt)
}

//PasswordStrength https://github.com/dropbox/zxcvbn
//=============================================================================
// strength measurement
//
// NOTE:
// for a little while, had rough draft of password strength measurement alg here.
// but not sure if there's value in yet another measurement algorithm,
// that's not just duplicating the effort of libraries like zxcbn.
// may revive it later, but for now, leaving some refs to others out there:
//    * NIST 800-63 has simple alg
//    * zxcvbn (https://tech.dropbox.com/2012/04/zxcvbn-realistic-password-strength-estimation/)
//      might also be good, and has approach similar to composite approach i was already thinking about,
//      but much more well thought out.
//    * passfault (https://github.com/c-a-m/passfault) looks thorough,
//      but may have licensing issues, plus porting to python looks like very big job :(
//    * give a look at running things through zlib - might be able to cheaply
//      catch extra redundancies.
//=============================================================================
func PasswordStrength(password string, userInputs []string) int {
	return 0
}
