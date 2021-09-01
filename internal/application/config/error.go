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

import "github.com/pkg/errors"

var (
	ErrHost                              = errors.New("host is empty")
	ErrPort                              = errors.New("port is empty")
	ErrAddress                           = errors.New("address is empty")
	ErrAccessKey                         = errors.New("access key is empty")
	ErrSecret                            = errors.New("secret is empty")
	ErrRegion                            = errors.New("region is empty")
	ErrLocalFilesFolder                  = errors.New("local files folder is not set")
	ErrConsumerGroupID                   = errors.New("consumer group id is empty")
	ErrBrokerAddressListEmpty            = errors.New("broker address list is empty")
	ErrTopicListEmpty                    = errors.New("topic list is empty")
	ErrWorkersCount                      = errors.New("workers count is wrong")
	ErrPathToConverterExecutable         = errors.New("path to converter executable is not set")
	ErrLargePngImageMaximumSideDimension = errors.New("large png image maximum side dimension is not valid")
	ErrSmallPngImageMaximumSideDimension = errors.New("small png image maximum side dimension is not valid")
	ErrFileSizeLimitSettingsFile         = errors.New("file size limit settings file is not set")
)
