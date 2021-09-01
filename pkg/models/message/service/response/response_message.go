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

package message

import (
	"github.com/pkg/errors"
	"github.com/vault-thirteen/SSE2/pkg/models/convertor/result"
	kafkaRequestMessage "github.com/vault-thirteen/SSE2/pkg/models/message/kafka/request"
)

type ResponseMessage struct {
	KafkaMessage     *kafkaRequestMessage.RequestMessage
	ConversionResult *result.ConversionResult
}

var ErrResultNotSet = errors.New("result is not set")

func (rm *ResponseMessage) IsResultSuccess() (isResultSuccess bool) {
	if rm.ConversionResult == nil {
		return false
	}

	return rm.ConversionResult.IsSuccess()
}

func (rm *ResponseMessage) GetError() (err error) {
	if rm.ConversionResult == nil {
		return ErrResultNotSet
	}

	return rm.ConversionResult.Error
}
