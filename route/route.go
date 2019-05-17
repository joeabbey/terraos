/*
	Copyright (c) 2019 Stellar Project

	Permission is hereby granted, free of charge, to any person
	obtaining a copy of this software and associated documentation
	files (the "Software"), to deal in the Software without
	restriction, including without limitation the rights to use, copy,
	modify, merge, publish, distribute, sublicense, and/or sell copies
	of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be
	included in all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY,
	WHETHER IN AN ACTION OF CONTRACT,
	TORT OR OTHERWISE,
	ARISING FROM, OUT OF OR IN CONNECTION WITH
	THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package route

import (
	"net"
	"os/exec"

	"github.com/pkg/errors"
)

const Interface = "mvlan0"

func Create(iface, address string) (err error) {
	// don't create if it already exists
	if _, err := net.InterfaceByName(Interface); err == nil {
		return nil
	}
	defer func() {
		if err != nil {
			ip("link", "del", Interface)
		}
	}()
	if err := ip("link", "add", "link", iface, Interface, "type", "macvlan", "mode", "bridge"); err != nil {
		return err
	}
	if err := ip("address", "add", address, "dev", Interface); err != nil {
		return err
	}
	if err := ip("link", "set", "dev", Interface, "up"); err != nil {
		return err
	}
	return ip("route", "flush", "dev", Interface)
}

func Add(address string) error {
	return ip("route", "add", address, "dev", Interface, "metric", "0")
}

func Remove(address string) error {
	return ip("route", "del", address, "dev", Interface)
}

func ip(args ...string) error {
	out, err := exec.Command("ip", args...).CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(out))
	}
	return nil
}
