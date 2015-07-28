package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

const INVALID_TIME_MARKER = "invalid time"

type BirthdayStorage struct {
	file      string
	birthdays BirthdaySet
}

func NewBirthdayStorage(file string) *BirthdayStorage {
	return &BirthdayStorage{file, make([]*Birthday, 0)}
}

func (bs *BirthdayStorage) Load() error {
	var (
		file *os.File
		err  error
	)

	if file, err = os.Open(bs.file); err != nil {
		return err
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&bs.birthdays); err != nil {
		return err
	}

	return nil
}

func (bs *BirthdayStorage) Save() error {
	var (
		file *os.File
		err  error
	)

	if file, err = os.Create(bs.file); err != nil {
		return err
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(bs.birthdays); err != nil {
		return err
	}

	return nil
}

func (bs *BirthdayStorage) GetBirthdaySet() BirthdaySet {
	return bs.birthdays
}

type BirthdaySet []*Birthday

func (bs BirthdaySet) FilterByName(name string) BirthdaySet {
	if len(name) == 0 {
		return bs
	}

	name = strings.ToLower(name)
	birthdays := make(BirthdaySet, 0)

	for _, bday := range bs {
		if strings.Contains(strings.ToLower(bday.Name), name) {
			birthdays = append(birthdays, bday)
		}
	}

	return birthdays
}

func (bs BirthdaySet) FilterByDuration(d time.Duration) BirthdaySet {
	if d == 0 {
		return []*Birthday{}
	}

	birthdays := make(BirthdaySet, 0)
	tnow := time.Now()

	for _, bday := range bs {
		tbday := bday.GetTime()
		btime := time.Date(tnow.Year(), tbday.Month(), tbday.Day(),
			0, 0, 0, 0, time.Local)
		diff := btime.Sub(tnow)

		if (diff >= 0 && diff < d) || (diff < 0 && diff > d) {
			birthdays = append(birthdays, bday)
		}
	}

	return birthdays
}

type Birthday struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

func (b *Birthday) GetTime() time.Time {
	if b.Time == INVALID_TIME_MARKER {
		return time.Now()
	}

	t, err := time.Parse(time.RFC822, b.Time)

	if err != nil {
		b.Time = INVALID_TIME_MARKER
		return time.Now()
	}

	return t
}

func (b *Birthday) GetAge() time.Duration {
	return time.Now().Sub(b.GetTime())
}
