/*
* @Author: Tacey Wong
* @Date:   2020-09-21 22:20:28
* @Last Modified by:   Tacey Wong
* @Last Modified time: 2020-09-21 22:22:19
*/
package pwd



type GenWordOpt struct{
	Entropy string
	Length int
	Charset string
	Chars string
	Returns string
}


// GenWord Generate one or more random passwords.
func GenWord(opt *GenWordOpt)string{
	return ""
}


func GenWordQuickly(length int)string{
	opt := GenWordOpt{
		Length=length
	}
	return GenWord(&opt)
}



type GenPhraseOpt struct {
	Entropy string
	Length int
	Wordset string
	Words string
	Sep string
	Returns string
}

// GenPhrase Generate one or more random password/passphrases 
func GenPhrase(opt *GenPhraseOpt)¶string{
	return ""
}


func GenPhraseQuickly(lenght int)string{
	opt := GenPhraseOpt{
		Length=lenght
	}
	return GenPhrase(&opt)
}




//PasswordStrength https://github.com/dropbox/zxcvbn
func PasswordStrength(password string, userInputs []string)int{
	return 0
}