/*
 * @Description:  RFC2307
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-21 19:07:25
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-21 19:15:59
 * @FilePath: /passlib/hash/ldapHashes.go
 */
package hash

// Standard LDAP Schemes

type LDAPMD5 struct{}
type LDAPSha1 struct{}
type LDAPSaltedMD5 struct{}
type LDAPSaltedSha1 struct{}
type LDAPPlainText struct{}
type LDAPDESCrypt struct{}
type LDAPBSDiCrypt struct{}
type LDAPMD5Crypt struct{}
type LDAPBCrypt struct{}
type LDAPSha1Crypt struct{}
type LDAPSha256Crypt struct{}
type LDAPSha512Crypt struct{}

// Non-Standard LDAP Schemes

type LDAPHexMD5 struct{}
type LDAPHexSha1 struct{}
type LDAPPBKDF2Sha1 struct{}
type LDAPPBKDF2Sha256 struct{}
type LDAPPBKDF2Sha512 struct{}
type FSHP struct{}
type RoundupPlainText struct{}
type AtlassianPBKDF2Sha1 struct{}
