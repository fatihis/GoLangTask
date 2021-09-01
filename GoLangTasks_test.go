package main

import (
	"testing"
)
//TESTS : TestSortByNumberOfCharA
// TestSortByNumberOfCharA calls main.TestSortByNumberOfCharA with a array, reordering array inorder to number of A's in array
func TestSortByNumberOfCharAOne(t *testing.T) {

	input := [6]string{"jumbA", "jumbAa", "jumbAaa", "jumbAaaa", "jumbAaaas"}

	want :=[6]string{"jumbAaaas", "jumbAaaa", "jumbAaa", "jumbAa","jumbA"}

	output, _ := sortByNumberOfCharA(input)

	a, err := sortByNumberOfCharA(input)
	if a != want {
		t.Fatalf(`sortByNumberOfCharA([6]string %q) = %q, want match for %#q, nil %q`, input, want, output, err)
	}
}

func TestSortByNumberOfCharATwo(t *testing.T) {

	input := [6]string{"", "aaa", "", "aa", "a"}

	want :=[6]string{"aaa", "aa", "a", "",""}

	output, _ := sortByNumberOfCharA(input)
	a, err := sortByNumberOfCharA(input)
	if a != want {
		t.Fatalf(`sortByNumberOfCharA([6]string %q) = %q, want match for %#q, nil %q`, input, want, output, err)
	}
}

func TestSortByNumberOfCharAThree(t *testing.T) {

	input := [6]string{"", "", "", "A", "a"}

	want :=[6]string{"A", "a", "", "",""}

	output, _ := sortByNumberOfCharA(input)
	a, err := sortByNumberOfCharA(input)
	if a != want {
		t.Fatalf(`sortByNumberOfCharA([6]string %q) = %q, want match for %#q, nil %q`, input, want, output, err)
	}
}
// <<END>> TestSortByNumberOfCharA <<END>>

//TESTS : TestFindMostRepeatedCharacter
// TestfFindMostRepeatedCharacter calls main.findMostRepeatedCharacter with a array, parse array and returns most repeated char
func TestFindMostRepeatedCharacterOne(t *testing.T) {

	input := [7]string{"aa", "e", "aa", "e", "e","c","aa"}

	want := "aa"

	output:= findMostRepeatedCharacter(input)

	if output != want {
		t.Fatalf(`findMostRepeatedCharacter([7]string %q) = %q, want match for %#q `, input, want, output)
	}
}

func TestFindMostRepeatedCharacterTwo(t *testing.T) {

	input := [7]string{"", "", "", "", "","",""}

	want := ""

	output:= findMostRepeatedCharacter(input)

	if output != want {
		t.Fatalf(`findMostRepeatedCharacter([7]string %q) = %q, want match for %#q `, input, want, output)
	}
}

func TestFindMostRepeatedCharacterAThree(t *testing.T) {

	input := [7]string{"1", "2", "1", "11", "11","11","1"}

	want := "1"

	output:= findMostRepeatedCharacter(input)

	if output != want {
		t.Fatalf(`findMostRepeatedCharacter([7]string %q) = %q, want match for %#q `, input, want, output)
	}
}
// <<END>> TestFindMostRepeatedCharacter <<END>>

