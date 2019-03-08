package es

import (
	"testing"

	"github.com/septemhill/fion"
)

func TestSubstring(t *testing.T) {
	str1 := fion.Red("Septem") + "跟" + fion.Blue("Nicole") + "說你好"
	es1 := NewEscapeString(str1)

	substr := es1.Substring(0, 6)
	if substr != fion.Red("Septem") {
		t.Fatal("got errors")
	}

	substr = es1.Substring(3, 7)
	if substr != fion.Red("tem")+"跟" {
		t.Fatal("got errors")
	}

	substr = es1.Substring(7, 10)
	if substr != fion.Blue("Nic") {
		t.Fatal("got errors")
	}

	str2 := "Midori" + fion.Green("と") + "Asolia" + fion.Cyan("all my friends")
	es2 := NewEscapeString(str2)

	substr = es2.Substring(2, 7)
	if substr != "dori"+fion.Green("と") {
		t.Fatal("got errors")
	}

	substr = es2.Substring(1, 14)
	if substr != "idori"+fion.Green("と")+"Asolia"+fion.Cyan("a") {
		t.Fatal("got errors")
	}

	str3 := "This pure string, without any escape character"
	es3 := NewEscapeString(str3)

	substr = es3.Substring(2, 12)
	if substr != "is pure st" {
		t.Fatal("got errors")
	}
}

func TestSubstringByWidth(t *testing.T) {
	str := "AあBいCうえお" + fion.Red("Hi, 世界")
	es := NewEscapeString(str)

	substr := es.SubstringByWidth(2, 5)
	if substr != "BいC" {
		t.Fatal("got errors")
	}

	substr = es.SubstringByWidth(0, 8)
	if substr != "AあBいC" {
		t.Fatal("got errors")
	}

	substr = es.SubstringByWidth(0, 9)
	if substr != "AあBいCう" {
		t.Fatal("got errors")
	}

	substr = es.SubstringByWidth(6, 7)
	if substr != "えお"+fion.Red("Hi,") {
		t.Fatal("got errors")
	}
}

func TestElement(t *testing.T) {
	str := "Yellow" + fion.Red("Green") + "科學知識" + fion.BCyan("放假")
	es := NewEscapeString(str)

	if es.Element(6) != fion.Red("G") {
		t.Fatal("got errors")
	}

	if es.Element(3) != "l" {
		t.Fatal("got errors")
	}

	if es.Element(11) != "科" {
		t.Fatal("got errors")
	}

	if es.Element(15) != fion.BCyan("放") {
		t.Fatal("got errors")
	}
}

func TestLen(t *testing.T) {
	str1 := "Hi, Septem. Ur so cooooooooool"
	es1 := NewEscapeString(str1)

	if es1.Len() != 30 {
		t.Fatal("got errors")
	}

	str2 := "你好世界，お腹がすいている？"
	es2 := NewEscapeString(str2)

	if es2.Len() != 14 {
		t.Fatal("got errors")
	}

	str3 := "Hi, 世界。>///< T口T <(__ __)> お腹がすいている"
	es3 := NewEscapeString(str3)
	if es3.Len() != 35 {
		t.Fatal("got errors")
	}
}

func TestWidth(t *testing.T) {
	str1 := "Hi, Nicole. Did u take my note?"
	es1 := NewEscapeString(str1)

	if es1.Width() != 31 {
		t.Fatal("got errors")
	}

	str2 := "下班惹，大家可以回家了"
	es2 := NewEscapeString(str2)

	if es2.Width() != 22 {
		t.Fatal("got errors")
	}

	str3 := "Time's up. 每個人把" + "携帯電話" + "收起來"
	es3 := NewEscapeString(str3)

	if es3.Width() != 33 {
		t.Fatal("got errors")
	}
}
