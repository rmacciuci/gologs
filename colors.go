package rmlogger

import (
	"errors"
)

const reset = "\033[0m"
const Red = Color("\033[31m")
const Greed = Color("\033[32m")
const Yellow = Color("\033[33m")
const Blue = Color("\033[34m")
const Purple = Color("\033[35m")
const Cyan = Color("\033[36m")
const Gray = Color("\033[37m")
const White = Color("\033[97m")

type Color string

type colors struct {
	info    string
	warning string
	error   string
	fatal   string
}

func NewColors() *colors {
	return &colors{}
}

func (c *colors) SetError(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.error = string(newColor)

	return err
}

func (c *colors) SetWarn(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.warning = string(newColor)

	return err
}

func (c *colors) SetFatal(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.fatal = string(newColor)

	return err
}

func (c *colors) SetInfo(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.info = string(newColor)

	return err
}
