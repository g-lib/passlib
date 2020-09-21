/*
 * @Description: Interface of PasswordHash
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-21 18:31:12
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-21 19:27:39
 * @FilePath: /passlib/hash/base.go
 */
package hash

type PasswordHash interface {
	Hash(string) string                   // generate new salt, return hash of password.
	Verify(password, hash string) bool    // verify password against existing hash.
	Using() *PasswordHash                 // create subclass with customized configuration.
	Identify() bool                       // check if hash belongs to this algorithm.
	Encrypt(string) string                // alias to Hash(string)string
	NeedsUpdate(hash, secret string) bool //
	Indentify(hash string) bool //check if hash belongs to this scheme, returns True/False
	
}

type PasswordHashBaseAttr struct {
	Name            string
	SettingKwds     map[string]string
	Contextkwds     map[string]string
	MaxSaltSize     int
	MinSaltSize     int
	DefaultSaltSize int
	SaltChars       string
	MaxRounds       int
	MinRounds       int
	DefaultRounds   int
	RoundsCost int
	 
}


