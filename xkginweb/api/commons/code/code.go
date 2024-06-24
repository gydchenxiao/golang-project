package code

import "github.com/mojocn/base64Captcha"

// Captcha captcha basic information.
type Captcha struct {
	Driver base64Captcha.Driver
	Store  base64Captcha.Store
}

// NewCaptcha creates a captcha instance from driver and store
func NewCaptcha(driver base64Captcha.Driver, store base64Captcha.Store) *Captcha {
	return &Captcha{Driver: driver, Store: store}
}

// Generate generates a random id, base64 image string or an error if any
func (c *Captcha) Generate() (id, b64s string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}
	c.Store.Set(id, answer)
	b64s = item.EncodeB64string()
	return
}

// Verify by a given id key and remove the captcha value in store,
// return boolean value.
// if you has multiple captcha instances which share a same store.
// You may want to call `store.Verify` method instead.
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	match = c.Store.Get(id, clear) == answer
	return
}
