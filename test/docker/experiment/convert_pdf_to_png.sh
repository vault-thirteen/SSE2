#!/bin/bash


################################################################################
#
#  Copyright © 2021 by Vault Thirteen.
#
#  All rights reserved. No part of this publication may be reproduced,
#  distributed, or transmitted in any form or by any means, including
#  photocopying, recording, or other electronic or mechanical methods,
#  without the prior written permission of the publisher, except in the case
#  of brief quotations embodied in critical reviews and certain other
#  noncommercial uses permitted by copyright law. For permission requests,
#  write to the publisher, addressed “Copyright Protected Material” at the
#  address below.
#
#  Web Site:  'https://github.com/vault-thirteen'.
#  Author:     Vault Thirteen.
#  Web Site Address is an Address in the global Computer Internet Network.
#
################################################################################

# Exit on Error.
set -e

# Input file.
FILE_TO_PROCESS="$1"
if [ ! -f "$FILE_TO_PROCESS" ]
then
  echo "Input file '$FILE_TO_PROCESS' does not exist."
  exit 1
fi

# Output Folder.
OUTPUT_FOLDER="$2"
if [ ! -d "$OUTPUT_FOLDER" ]
then
  echo "Output folder '$OUTPUT_FOLDER' does not exist."
  exit 1
fi

# Convert the file.
soffice --convert-to png "$FILE_TO_PROCESS" --outdir "$OUTPUT_FOLDER" --headless
