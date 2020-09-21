package apache

type HtdigestFile struct {
}

func NewHtdigestFile(filepath, defaultRealm string, new bool, autosave bool, encoding string, autoload bool) *HtdigestFile {
	return nil
}

func NewHtdigestFileFromString(data string) *HtdigestFile {
	return nil
}
func (hdf *HtdigestFile) Load(path string, force bool) {

}

func (hdf *HtdigestFile) LoadIfChanged() {}

func (hdf *HtdigestFile) LoadString(data string) {}

func (hdf *HtdigestFile) Save(path string) {

}

func (hdf *HtdigestFile) ToString() string {
	return ""
}

func (hdf *HtdigestFile) Realms() string {
	return ""
}

func (hdf *HtdigestFile) Users() []string {
	return []string{}
}

func (hdf *HtdigestFile) CheckPassword(user, realm, password string) bool {
	return false
}

func (hdf *HtdigestFile) GetHash() string {
	return ""
}

func (hdf *HtdigestFile) SetPassword(user, realm, password string) {

}

func (hdf *HtdigestFile) Delete(user, realm string) {

}

func (hdf *HtdigestFile) DeleteRealm(realm string) {

}
