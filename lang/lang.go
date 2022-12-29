/********************************************************************************
* Temancode Lang Package                                            			*
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
********************************************************************************/

package lang

var LangMessage = map[string]string{}

type language struct {
	Name string
	Data map[string]string
}

type Language interface {
	GetName() string
	GetData() map[string]string
}

var StubStorage = map[string]func(){
	"ID": ID,
	"EN": EN,
}

func AddLang(initial string, f func()) {
	StubStorage[initial] = f
}

func NewLanguage(name string) *language {
	CallLang(StubStorage[name])
	return &language{Name: name, Data: LangMessage}
}

func (l *language) GetName() string {
	return l.Name
}

func (l *language) GetData() map[string]string {
	return l.Data
}

func (l *language) ChangeMessage(key string, message string) *language {
	l.Data[key] = message
	return l
}

// func NewLang(language string) map[string]string {
// 	CallLang(StubStorage[language])
// 	return LangMessage
// }

func CallLang(f func()) {
	f()
}

func ID() {
	LangMessage = LangID
}

func EN() {
	LangMessage = LangEN
}
