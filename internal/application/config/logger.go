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

import "github.com/kelseyhightower/envconfig"

type Logger struct {
	IsDebugEnabled bool `split_words:"true" default:"false"`
}

func NewLogger(envPrefix string) (cfg *Logger, err error) {
	cfg = new(Logger)
	err = envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (l *Logger) IsValid() (bool, error) {
	return true, nil
}

func GetLoggerConfig() (loggerConfig *Logger, err error) {
	loggerConfig, err = NewLogger(EnvironmentVariablePrefixApplication)
	if err != nil {
		return nil, err
	}

	_, err = loggerConfig.IsValid()
	if err != nil {
		return nil, err
	}

	return loggerConfig, nil
}
