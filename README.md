# AGH Web Assets

This repository contains the generated production web bundle consumed by
`github.com/compozy/agh`.

The AGH main repository intentionally does not version generated frontend
assets. Release preparation builds `web/dist`, publishes this module, and then
the AGH CLI imports this module so `go install github.com/compozy/agh@latest`
ships the complete single binary with the embedded UI.

Do not edit `dist/` by hand. Regenerate it from the AGH main repository.
