# Spectrum v1.1.1 Release Note

## Overview

v1.1.1 is a patch release containing build compatibility fixes and mainnet bootnodes updates.

---

## Changes

### 🐛 Build & Dependencies

- Upgraded `golang.org/x/tools` from `v0.1.6` to `v0.49.0`, along with related `x/mod`, `x/net`, `x/sys`, `x/crypto`, `x/text` dependencies.
- Bumped the `go` directive from `1.15` to `1.25`.
- Fixed compilation failures on Go 1.25+ caused by `go/types` API changes (`SetTypeParams`, `SetRecvTypeParams`, `IsConstraint`, `Inferred`, etc.).

### 🔗 Networking

- Updated mainnet bootnodes: adjusted the IP addresses of 5 bootstrap nodes and removed the former `cn : shanghai` node.

### 🐛 VRF Elliptic Curve Panic on Go 1.25

- Fixed a runtime panic in `ProofToHash` during FullSync: `crypto/elliptic: attempted operation on invalid point`.
- Go 1.25 added strict `panicIfNotOnCurve` validation to `elliptic.CurveParams.Add`, which the VRF code triggered when using the standard library's `Add` on intermediate points from the CGO secp256k1 implementation.
- Changed `params.Add` to `curve.Add` in `crypto/vrf/vrf.go` to use the `BitCurve`'s own addition, consistent with the rest of the VRF logic.

---

## Full Changes

For the complete commit history against `v1.1.0`, see: `git log v1.1.0..v1.1.1`
