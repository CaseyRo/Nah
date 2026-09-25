# Deploy tasks

## 1. Repository

- [x] 1.1 `apps/server/compose.yaml`: build in place, `/data` volume, 30 second stop grace, port on nebula-1's Tailscale address
- [x] 1.2 Dated note on ADR-0014: build from git, registry at M7

## 2. Fleet

- [x] 2.1 Komodo stack on nebula-1 from `CaseyRo/Nah`, `apps/server/compose.yaml`
- [x] 2.2 GitHub webhook to the stack, the same secret on both sides; one delivery answered 200
- [x] 2.3 Tunnel ingress and DNS for `nah.casey.berlin`, no Access
- [x] 2.4 Gatus endpoint on `/healthz`

## 3. Verification

- [x] 3.1 `https://nah.casey.berlin/healthz` answers ok
- [ ] 3.2 The two-person test passes against it
- [ ] 3.3 A second deploy keeps a session signed in (ADR-0015)
- [x] 3.4 A push deploys by itself: `webhook_force_deploy` and `run_build` on, and a redelivered push produced a DeployStack by the webhook
