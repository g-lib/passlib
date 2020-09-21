/*
 * @Description: https://foss.heptapod.net/python-libs/passlib/-/blob/branch/stable/passlib/exc.py
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-21 13:51:00
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-21 14:05:25
 * @FilePath: /passlib/exceptions/exceptions.go
 */

package exceptions

// UnknownBackendError Error raised if multi-backend handler doesn't recognize backend name.
//    Inherits from :exc:`ValueError`.
type UnknownBackendError struct {
}

// MissingBackendError Error raised if multi-backend handler has no available backends;
//    or if specifically requested backend is not available.
//    :exc:`!MissingBackendError` derives
//    from :exc:`RuntimeError`, since it usually indicates
//    lack of an external library or OS feature.
//    This is primarily raised by handlers which depend on
//    external libraries (which is currently just
//    :class:`~passlib.hash.bcrypt`).
type MissingBackendError struct {
}

// PasswordValueError Error raised if a password can't be hashed / verified for various reasons.
//    This exception derives from the builtin :exc:`!ValueError`.
//    May be thrown directly when password violates internal invariants of hasher
//    (e.g. some don't support NULL characters).  Hashers may also throw more specific subclasses,
//    such as :exc:`!PasswordSizeError`.
type PasswordValueError struct {
}

type PasswordSizeError struct {
}
type PasswordTruncateError struct{}
type PasslibSecurityError struct{}
type TokenError struct{}
type MalformedTokenError struct{}
type InvalidTokenError struct{}
type UsedTokenError struct{}
type UnknownHashErro struct{}
type PasslibWarning struct{}
type PasslibConfigWarning struct{}
type PasslibHashWarning struct{}
type PasslibRuntimeWarning struct{}
type PasslibSecurityWarning struct{}
type ExpectedTypeError struct{}
type ExpectedStringError struct{}
type MissingDigestError struct{}
type NullPasswordError struct{}
type InvalidHashError struct{}
type MalformedHashError struct{}
type ZeroPaddedRoundsError struct{}
type ChecksumSizeError struct{}
