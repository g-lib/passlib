/*
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-21 17:06:55
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-21 18:17:07
 * @FilePath: /passlib/apache/htpasswd.go
 */

package apache

type HtpasswdFile struct {
}

func NewHtpasswdFile(filePath string, new bool, autoSave bool, encoding string, defaultScheme string, context string) *HtpasswdFile {
	return nil
}

func NewHtpasswdFileFromString(data string) {}

func (hpf *HtpasswdFile) SetPassword(user, password string) {

}

func (hpf *HtpasswdFile) Load(path string, force bool) {

}

func (hpf *HtpasswdFile) LoadIfChanged() {}

func (hpf *HtpasswdFile) LoadString(data string) {}

func (hpf *HtpasswdFile) Save() {

}

func (hpf *HtpasswdFile) Users() []string {
	return []string{}
}
func (hpf *HtpasswdFile) CheckPassword(user, password string) bool {
	return false
}
func (hpf *HtpasswdFile) ToString() string {
	return ""
}

func (hpf *HtpasswdFile) GetHash(user string) string {
	return ""
}

func (hpf *HtpasswdFile) Delete(user string) {}
