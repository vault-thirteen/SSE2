////////////////////////////////////////////////////////////////////////////////
//
//  Copyright © 2021 by Vault Thirteen.
//
//  All rights reserved. No part of this publication may be reproduced,
//  distributed, or transmitted in any form or by any means, including
//  photocopying, recording, or other electronic or mechanical methods,
//  without the prior written permission of the publisher, except in the case
//  of brief quotations embodied in critical reviews and certain other
//  noncommercial uses permitted by copyright law. For permission requests,
//  write to the publisher, addressed “Copyright Protected Material” at the
//  address below.
//
//  Web Site:  'https://github.com/vault-thirteen'.
//  Author:     Vault Thirteen.
//  Web Site Address is an Address in the global Computer Internet Network.
//
////////////////////////////////////////////////////////////////////////////////

package settings

import (
	"encoding/xml"
	"strings"

	"github.com/vault-thirteen/SSE2/internal/helper"
)

type XmlSettings struct {
	XMLName xml.Name `xml:"Settings"`

	FileSizeLimiter XmlSettingsFileSizeLimiter `xml:"FileSizeLimiter"`
}

type XmlSettingsFileSizeLimiter struct {
	XMLName xml.Name `xml:"FileSizeLimiter"`

	MimeType []XmlSettingsFileSizeLimiterMimeType `xml:"MimeType"`
}

type XmlSettingsFileSizeLimiterMimeType struct {
	XMLName xml.Name `xml:"MimeType"`

	Name      string `xml:"name,attr"`
	SizeLimit int    `xml:"sizeLimit,attr"`
}

func NewXmlSettings(
	filePath string,
) (xmlConfig *XmlSettings, err error) {
	var cfgFileContents string
	cfgFileContents, err = helper.GetTextFileContents(filePath)
	if err != nil {
		return nil, err
	}

	var decoder = xml.NewDecoder(strings.NewReader(cfgFileContents))
	decoder.Strict = true

	xmlConfig = new(XmlSettings)

	err = decoder.Decode(xmlConfig)
	if err != nil {
		return nil, err
	}

	return xmlConfig, nil
}
