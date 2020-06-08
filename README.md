<!--
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-06-08 15:43:39
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-06-08 15:56:21
 * @FilePath: /passlib/README.md
--> 
# passlib
Port of https://passlib.readthedocs.io/en/stable/

+ Apache
    + Htpasswd Files
    + Htdigest Files

+ Hash
    + Unix Hash
        + bcrypt - BCrypt
        + sha256_crypt - SHA-256 Crypt
        + sha512_crypt - SHA-512 Crypt
        + unix_disabled - Unix Disabled Account Helper
        + bsd_nthash - FreeBSD’s MCF-compatible encoding of nthash digests
        + md5_crypt - MD5 Crypt
        + sha1_crypt - SHA-1 Crypt
        + sun_md5_crypt - Sun MD5 Crypt
        + des_crypt - DES Crypt
        + bsdi_crypt - BSDi Crypt
        + bigcrypt - BigCrypt
        + crypt16 - Crypt16
    + Modular Crypt” Hashes
        + argon2 - Argon2
        + bcrypt_sha256 - BCrypt+SHA256
        + phpass - PHPass’ Portable Hash
        + pbkdf2_digest - Generic PBKDF2 Hashes
        + scram - SCRAM Hash
        + scrypt - SCrypt
        + apr_md5_crypt - Apache’s MD5-Crypt variant
        + cta_pbkdf2_sha1 - Cryptacular’s PBKDF2 hash
        + dlitz_pbkdf2_sha1 - Dwayne Litzenberger’s PBKDF2 hash
    + LDAP / RFC2307 Hashes
        + ldap_md5 - MD5 digest
        + ldap_sha1 - SHA1 digest
        + ldap_salted_md5 - salted MD5 digest
        + ldap_salted_sha1 - salted SHA1 digest
        + ldap_crypt - LDAP crypt() Wrappers
        + ldap_plaintext - LDAP-Aware Plaintext Handler
        + ldap_hex_md5 - Hex-encoded MD5 Digest
        + ldap_hex_sha1 - Hex-encoded SHA1 Digest
        + ldap_pbkdf2_digest - Generic PBKDF2 Hashes
        + atlassian_pbkdf2_sha1 - Atlassian’s PBKDF2-based Hash
        + fshp - Fairly Secure Hashed Password
        + roundup_plaintext - Roundup-specific LDAP Plaintext Handler
    + SQL Database Hashes
        + mssql2000 - MS SQL 2000 password hash
        + mssql2005 - MS SQL 2005 password hash
        + mysql323 - MySQL 3.2.3 password hash
        + mysql41 - MySQL 4.1 password hash
        + postgres_md5 - PostgreSQL MD5 password hash
        + oracle10 - Oracle 10g password hash
        + oracle11 - Oracle 11g password hash
    + MS Windows Hashes
        + lmhash - LanManager Hash
        + nthash - Windows’ NT-HASH
        + msdcc - Windows’ Domain Cached Credentials
        + msdcc2 - Windows’ Domain Cached Credentials v2
    + Cisco Hashes
        + cisco_pix
        + cisco_asa
    + Other
        + django_digest - Django-specific Hashes
        + grub_pbkdf2_sha512 - Grub’s PBKDF2 Hash
        + hex_digest - Generic Hexadecimal Digests
        + plaintext - Plaintext
