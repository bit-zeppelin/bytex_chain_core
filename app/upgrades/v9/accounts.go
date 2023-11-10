// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

// This accounts represent the affected accounts during the Claims record clawback incident on block 5074187
// with the respective balance on block 5074186.

// The accounts fell under the following conditions:
// - They were in the claims record or attestation accounts
// - They had a balance bigger than DUST (amount sent to claim record accounts on genesis to pay for gas) on block 5074186
// - They had an account sequence of 0 by block 5074186
// - They had a balance of 0 after block 5074187
// NOTE community and claims module account were removed from this list since they were not affected by this bug.

// The scripts that were used to find affected accounts were made public at https://github.com/evmos/claims_fixer
// with detail instructions on how to run them.

package v9

var Accounts = [0][2]string{
}
