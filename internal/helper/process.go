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

package helper

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"

	"github.com/rs/zerolog"
	"github.com/vault-thirteen/SSE2/internal/messages"
)

var (
	ErrNoProcess        = errors.New("process has not been created")
	ErrStdOutAlreadySet = errors.New("stdout is already set")
	ErrStdErrAlreadySet = errors.New("stderr is already set")
)

func ExecuteCommandAndGetOutput(
	ctx context.Context,
	logger *zerolog.Logger,
	command string,
	arguments []string,
) (processId *int, outputLines []string, err error) {
	cmd := exec.CommandContext(ctx, command, arguments...)

	if cmd.Stdout != nil {
		return nil, nil, ErrStdOutAlreadySet
	}
	if cmd.Stderr != nil {
		return nil, nil, ErrStdErrAlreadySet
	}

	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	cmd.Stderr = &buffer

	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}

	if cmd.Process == nil {
		return nil, nil, ErrNoProcess
	}
	processId = NewIntPointer(cmd.Process.Pid)

	defer func() {
		outputLines = strings.Split(buffer.String(), "\n")
	}()

	logger.Debug().Msgf(messages.MsgFProcessHasBeenCreated, cmd.Process.Pid)

	err = cmd.Wait()
	if err != nil {
		return processId, outputLines, err
	}

	logger.Debug().Msgf(messages.MsgFProcessHasFinished, cmd.Process.Pid)

	return processId, outputLines, nil
}
