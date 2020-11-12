package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/dodo-cli/dodo-core/pkg/decoder"
)

const ErrPortFormat FormatError = "invalid publish format"

func (port *Port) FromString(spec string) error {
	var target, published, protocol, hostip string
	switch values := strings.SplitN(spec, ":", 3); len(values) {
	case 0:
		return fmt.Errorf("%s: %w", spec, ErrPortFormat)
	case 1:
		target = values[0]
	case 2:
		published = values[0]
		target = values[1]
	case 3:
		hostip = values[0]
		published = values[1]
		target = values[2]
	default:
		return fmt.Errorf("%s: %w", spec, ErrPortFormat)
	}

	switch values := strings.SplitN(target, "/", 2); len(values) {
	case 1:
		target = values[0]
	case 2:
		target = values[0]
		protocol = values[1]
	default:
		return fmt.Errorf("%s: %w", spec, ErrPortFormat)
	}

	if p, err := strconv.ParseInt(target, 10, 32); err != nil {
		return fmt.Errorf("%s: %w", target, err)
	} else {
		port.Target = int32(p)
	}

	if p, err := strconv.ParseInt(published, 10, 32); err != nil {
		return fmt.Errorf("%s: %w", published, err)
	} else {
		port.Published = int32(p)
	}

	switch protocol {
	case "tcp", "udp", "sctp":
		port.Protocol = protocol
	default:
		return fmt.Errorf("%s: %w", spec, ErrPortFormat)
	}

	port.HostIp = hostip

	return nil
}

func NewPort() decoder.Producer {
	return func() (interface{}, decoder.Decoding) {
		target := &Port{}
		return &target, DecodePort(&target)
	}
}

func DecodePort(target interface{}) decoder.Decoding {
	// TODO: wtf this cast
	port := *(target.(**Port))

	return decoder.Kinds(map[reflect.Kind]decoder.Decoding{
		reflect.Map: decoder.Keys(map[string]decoder.Decoding{
			"target":    decoder.String(&port.Target),
			"published": decoder.String(&port.Published),
			"protocol":  decoder.String(&port.Protocol),
			"host_ip":   decoder.String(&port.HostIp),
		}),
		reflect.String: func(d *decoder.Decoder, config interface{}) {
			var decoded string
			decoder.String(&decoded)(d, config)
			if err := port.FromString(decoded); err != nil {
				d.Error(err)
			}
		},
	})
}
