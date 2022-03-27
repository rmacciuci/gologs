package colors

import (
	"errors"
)

const reset = "\033[0m"
const red = Color("\033[31m")
const green = Color("\033[32m")
const yellow = Color("\033[33m")
const blue = Color("\033[34m")
const purple = Color("\033[35m")
const cyan = Color("\033[36m")
const gray = Color("\033[37m")
const white = Color("\033[97m")

type Color string

type Colors struct {
	Info    string
	Warning string
	Error   string
	Fatal   string
}

func GetReset() string {
	return reset
}

func GetBlue() Color {
	return Color(blue)
}
func GetPurple() Color {
	return Color(purple)
}
func GetCyan() Color {
	return Color(cyan)
}
func GetGray() Color {
	return Color(gray)
}
func GetWhite() Color {
	return Color(white)
}
func GetRed() Color {
	return Color(red)
}

func GetYellow() Color {
	return Color(yellow)
}

func GetGreen() Color {
	return Color(green)
}

func NewColors(info, warn, error, fatal Color) *Colors {
	return &Colors{
		Info:    string(info),
		Warning: string(warn),
		Error:   string(error),
		Fatal:   string(fatal),
	}
}

func InitDefault() *Colors {
	return NewColors(
		GetBlue(),
		GetYellow(),
		GetRed(),
		GetRed(),
	)
}

func (c *Colors) SetError(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("not defined color")
	}

	c.Error = string(newColor)

	return err
}

func (c *Colors) SetWarn(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("not defined color")
	}

	c.Warning = string(newColor)

	return err
}

func (c *Colors) SetFatal(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("not defined color")
	}

	c.Fatal = string(newColor)

	return err
}

func (c *Colors) SetInfo(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("not defined color")
	}

	c.Info = string(newColor)

	return err
}
