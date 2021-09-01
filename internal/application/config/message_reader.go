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

package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/vault-thirteen/SSE2/internal/helper"
)

type MessageReader struct {
	KafkaConsumerGroupID   string   `split_words:"true"`
	KafkaBrokerAddressList []string `split_words:"true"`
	KafkaTopicList         []string `split_words:"true"`
}

func NewMessageReader(envPrefix string) (cfg *MessageReader, err error) {
	cfg = new(MessageReader)
	err = envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (mrc *MessageReader) IsValid() (bool, error) {
	if len(mrc.KafkaConsumerGroupID) < 1 {
		return false, ErrConsumerGroupID
	}

	if len(mrc.KafkaBrokerAddressList) < 1 {
		return false, ErrBrokerAddressListEmpty
	}

	if len(mrc.KafkaTopicList) < 1 {
		return false, ErrTopicListEmpty
	}

	return true, nil
}

func GetMessageReaderConfig() (readerConfig *MessageReader, err error) {
	envPrefix := helper.ConcatenateEnvVarPrefixes(
		EnvironmentVariablePrefixApplication,
		EnvironmentVariablePrefixKafkaInput,
	)

	readerConfig, err = NewMessageReader(envPrefix)
	if err != nil {
		return nil, err
	}

	_, err = readerConfig.IsValid()
	if err != nil {
		return nil, err
	}

	return readerConfig, nil
}
