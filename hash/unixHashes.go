/*
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-21 18:25:02
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-21 18:53:22
 * @FilePath: /passlib/hash/unixHashes.go
 */
package hash

// Active Unix Hashes

// Bcrypt https://www.usenix.org/legacy/event/usenix99/provos/provos_html/
type Bcrypt struct {
}

func NewBcrypt(salt string, rounds int, ident string, truncateError bool, relaxed bool) *Bcrypt {
	return nil
}

// Sha256Crypt http://www.akkadia.org/drepper/sha-crypt.html
type Sha256Crypt struct {
}

func NewSha256Crypt(salt string, rounds int, implicitRounds bool, relaxed bool) *Sha256Crypt {
	return nil
}

// Sha512Crypt http://www.akkadia.org/drepper/sha-crypt.html
type Sha512Crypt struct {
}

func NewSha512Crypt(salt string, rounds int, implicitRounds bool, relaxed bool) *Sha512Crypt {
	return nil
}

type UnixDisabled struct {
}

func NewUnixDisabled(marker string) *UnixDisabled {
	return nil
}

type UnixFallback struct {
}

func NewUnixFallback() *UnixFallback {
	return nil
}

// Deprecated Unix Hashes

type BSDNtHash struct {
}

type MD5Crypt struct {
}

type Sha1Crypt struct {
}

type SunMD5Crypt struct {
}

// Archaic Unix Hashes

type DesCrypt struct{}

type BSDiCrypt struct{}

type BigCrypt struct{}

type Crypt16 struct{}
