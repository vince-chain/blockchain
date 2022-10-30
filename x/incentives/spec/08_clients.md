<!--
order: 8
-->

# Clients

A user can query the `x/incentives` module using the CLI, JSON-RPC, gRPC or REST.

## CLI

Find below a list of `vinced` commands added with the `x/incentives` module. You can obtain the full list by using the `vinced -h` command.

### Queries

The `query` commands allow users to query `incentives` state.

**`incentives`**

Allows users to query all registered incentives.

```go
vinced query incentives incentives [flags]
```

**`incentive`**

Allows users to query an incentive for a given contract.

```go
vinced query incentives incentive [contract-address] [flags]
```

**`gas-meters`**

Allows users to query all gas meters for a given incentive.

```bash
vinced query incentives gas-meters [contract-address] [flags]
```

**`gas-meter`**

Allows users to query a gas meter for a given incentive and user.

```go
vinced query incentives gas-meter [contract-address] [participant-address] [flags]
```

**`params`**

Allows users to query incentives params.

```bash
vinced query incentives params [flags]
```

### Proposals

The `tx gov submit-proposal` commands allow users to query create a proposal using the governance module CLI:

**`register-incentive`**

Allows users to submit a `RegisterIncentiveProposal`.

```bash
vinced tx gov submit-proposal register-incentive [contract-address] [allocation] [epochs] [flags]
```

**`cancel-incentive`**

Allows users to submit a `CanelIncentiveProposal`.

```bash
vinced tx gov submit-proposal cancel-incentive [contract-address] [flags]
```

**`param-change`**

Allows users to submit a `ParameterChangeProposal``.

```bash
vinced tx gov submit-proposal param-change [proposal-file] [flags]
```

## gRPC

### Queries

| Verb   | Method                                                     | Description                                   |
| ------ | ---------------------------------------------------------- | --------------------------------------------- |
| `gRPC` | `vince.incentives.v1.Query/Incentives`                     | Gets all registered incentives                |
| `gRPC` | `vince.incentives.v1.Query/Incentive`                      | Gets incentive for a given contract           |
| `gRPC` | `vince.incentives.v1.Query/GasMeters`                      | Gets gas meters for a given incentive         |
| `gRPC` | `vince.incentives.v1.Query/GasMeter`                       | Gets gas meter for a given incentive and user |
| `gRPC` | `vince.incentives.v1.Query/AllocationMeters`               | Gets all allocation meters                    |
| `gRPC` | `vince.incentives.v1.Query/AllocationMeter`                | Gets allocation meter for a denom             |
| `gRPC` | `vince.incentives.v1.Query/Params`                         | Gets incentives params                        |
| `GET`  | `/vince/incentives/v1/incentives`                          | Gets all registered incentives                |
| `GET`  | `/vince/incentives/v1/incentives/{contract}`               | Gets incentive for a given contract           |
| `GET`  | `/vince/incentives/v1/gas_meters`                          | Gets gas meters for a given incentive         |
| `GET`  | `/vince/incentives/v1/gas_meters/{contract}/{participant}` | Gets gas meter for a given incentive and user |
| `GET`  | `/vince/incentives/v1/allocation_meters`                   | Gets all allocation meters                    |
| `GET`  | `/vince/incentives/v1/allocation_meters/{denom}`           | Gets allocation meter for a denom             |
| `GET`  | `/vince/incentives/v1/params`                              | Gets incentives params                        |
