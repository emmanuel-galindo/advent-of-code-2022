package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestProvisioning(t *testing.T) {
	t.Run("read calories data", func(t *testing.T) {
		data := "5000\n1000\n\n3000"

		tmpFile, cleanFile := createTempFile(t, data)
		defer cleanFile()

		got, err := GetData(tmpFile)

		assertNoError(t, err)
		assertStringEquals(t, string(got), data)
	})
	t.Run("get calories per Elf", func(t *testing.T) {
		data := []byte("5000\n1000\n\n3000\n")
		want := []int{6000, 3000}
		got, err := CaloriesPerElf(bytes.NewBuffer(data))

		assertNoError(t, err)
		assertArrayEqual(t, got, want)
	})
	t.Run("get Elf with more calories", func(t *testing.T) {
		data := []int{6000, 9000}
		want := 1

		got := MaxElf(data)

		assertIntEqual(t, got, want)
	})
	t.Run("get Elfes calories ordered ascending", func(t *testing.T) {
		data := []int{6000, 9000, 1000, 4000, 2100}
		want := []int{9000, 6000, 4000, 2100, 1000}

		got := OrderedCalories(data)

		assertArrayEqual(t, got, want)
	})
	t.Run("sum top 3", func(t *testing.T) {
		data := []int{9000, 6000, 4000, 2100, 1000}
		want := 19000

		got := SumTop3(data)

		assertIntEqual(t, want, got)

	})
}

func assertIntEqual(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertArrayEqual(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertStringEquals(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "inputs")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))
	tmpFile.Seek(0, 0)

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
