package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialModel(t *testing.T) {
	m := initialModel()
	assert.NotNil(t, m.spinner)
	assert.False(t, m.quitting)
	assert.Nil(t, m.err)
}

func TestModelInit(t *testing.T) {
	m := initialModel()
	cmd := m.Init()
	assert.NotNil(t, cmd)
}

func TestModelUpdateErrMsg(t *testing.T) {
	m := initialModel()
	err := errMsg(fmt.Errorf("test error"))
	updatedModel, _ := m.Update(err)
	assert.Equal(t, "test error", updatedModel.(model).err.Error())
}

func TestModelViewWithError(t *testing.T) {
	m := initialModel()
	m.err = fmt.Errorf("test error")
	view := m.View()
	assert.Equal(t, "test error", view)
}

func TestModelViewWithoutError(t *testing.T) {
	m := initialModel()
	view := m.View()
	assert.Contains(t, view, "Loading awesomeness")
}

func TestShowAwesomeBanner(t *testing.T) {
	banner := showAwesomeBanner()
	assert.Contains(t, banner, "AWESOME! ;)")
}
