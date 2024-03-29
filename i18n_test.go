package i18next_test

import (
	"fmt"
	"github.com/yuangwei/go-i18next"
	"testing"
)

func TestI18nFromJsonLocal(t *testing.T) {
	i18n, err := i18next.Init(i18next.I18nOptions{
		Lng:        []string{"en", "cn"},
		DefaultLng: "en",
		Ns:         "json",
		Backend: i18next.Backend{
			LoadPath: []string{
				"./examples/datas/{{.Ns}}/{{.Lng}}/home.json",
			},
		},
	})
	if err != nil {
		t.Fatalf("i18n Initial error %s", err)
	}
	val, err := i18n.T("title", struct {
		Name string
	}{Name: "Mike"})

	if err != nil {
		t.Fatalf("i18n Format error %s", err)
	}
	if val != "Hello, Mike" {
		t.Fatalf("i18n Format error")
	}
	fmt.Println(val)
}

func TestI18nChangeLanguage(t *testing.T) {
	i18n, err := i18next.Init(i18next.I18nOptions{
		Lng:        []string{"en", "cn"},
		DefaultLng: "en",
		Ns:         "json",
		Backend: i18next.Backend{
			LoadPath: []string{
				"./examples/datas/{{.Ns}}/{{.Lng}}/home.json",
			},
		},
	})
	if err != nil {
		t.Fatalf("i18n Initial error %s", err)
	}
	val, err := i18n.T("title", struct {
		Name string
	}{Name: "Mike"})

	if err != nil {
		t.Fatalf("i18n Format error %s", err)
	}
	if val != "Hello, Mike" {
		t.Fatalf("i18n Format error")
	}
	err = i18n.ChangeLanguage("cn")
	if err != nil {
		t.Fatalf("i18n changeLanguage error %s", err)
	}
	val, err = i18n.T("title", struct {
		Name string
	}{Name: "麦克"})
	if val != "你好, 麦克" {
		t.Fatalf("i18n Format error")
	}
	fmt.Println(val)
}

func TestI18nFromJsonNetwork(t *testing.T) {
	i18n, err := i18next.Init(i18next.I18nOptions{
		Lng:        []string{"en", "cn"},
		DefaultLng: "en",
		Ns:         "json",
		Backend: i18next.Backend{
			LoadPath: []string{
				"https://res.cloudinary.com/dvbgkny62/raw/upload/v1703728229/i18n/{{.Ns}}/{{.Lng}}/home.json",
			},
		},
	})
	if err != nil {
		t.Fatalf("i18n Initial error %s", err)
	}
	val, err := i18n.T("title", struct {
		Name string
	}{Name: "Mike"})

	if err != nil {
		t.Fatalf("i18n Format error %s", err)
	}
	if val != "Hello, Mike" {
		t.Fatalf("i18n Format error")
	}
	fmt.Println(val)
}
