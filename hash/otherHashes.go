/*
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-21 18:59:25
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-21 19:06:22
 * @FilePath: /passlib/hash/otherHashes.go
 */
package hash

// Django related

type DjangoArgon2 struct{} // 1.10

type DjangoBcryptSha256 struct{} // 1.6

type DjangoPBKDF2Sha256 struct{} //1.4

type DjangoPBKDF2Sha1 struct{} // 1.4

type DjangoBcrypt struct{} // 1.4

// grub

type GrubPBKDF2Sha512 struct{}

// Hexadecimal Digests

type HexMD4 struct{}
type HexMD5 struct{}
type HexSha1 struct{}
type HexSHa256 struct{}
type HexSha512 struct{}

// plaintext

type PlainText struct{}
